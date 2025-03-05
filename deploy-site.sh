#!/bin/sh

# Deploy configuration
DEPLOY_HOST=$(grep DEPLOY_HOST .env | cut -d '=' -f2)
DEPLOY_DIR=$(grep DEPLOY_DIR .env | cut -d '=' -f2)
GITHUB_USER=$(grep GITHUB_USER .env | cut -d '=' -f2)
CR_PAT=$(grep CR_PAT .env | cut -d '=' -f2)
TAG=$(git rev-parse --short HEAD)
DEPLOY_CMD="TAG=$TAG docker-compose pull && TAG=$TAG docker-compose up -d"

# Color codes
YELLOW='\033[0;33m'
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

warnmsg() {
    printf "${YELLOW}warning: %s${NC}\n" "$1" >&2
}

debugmsg() {
    printf "${RED}debug: %s${NC}\n" "$1" >&2
}

infomsg() {
    printf "${GREEN}info: %s${NC}\n" "$1" >&2
}

confirm_prompt() {
    while true; do
        warnmsg "Do you want to proceed? (y/n) "
        read choice
        case "$choice" in
        y | Y)
            echo "Proceeding..."
            return 0
            ;;
        n | N)
            echo "Exiting..."
            return 1
            ;;
        *)
            warnmsg "Please answer y or n."
            ;;
        esac
    done
}

main() {
    if [ ! -f .env ]; then
        warnmsg "Local .env file not found!"
        exit 1
    fi

    # Login to GitHub Container Registry on remote server
    infomsg "Logging into GitHub Container Registry on '$DEPLOY_HOST'."
    ssh "$DEPLOY_HOST" "echo $CR_PAT | docker login ghcr.io -u $GITHUB_USER --password-stdin"

    infomsg "Syncing files to '$DEPLOY_DIR' on '$DEPLOY_HOST'."
    rsync -avz --delete docker-compose.yaml Caddyfile "$DEPLOY_HOST:$DEPLOY_DIR"

    case "${1:-}" in
    -f | --force)
        warnmsg "Running '$0' with the force [-f|--force] option will permanently"
        warnmsg "remove docker-compose data under '$DEPLOY_DIR' on '$DEPLOY_HOST'."
        warnmsg "Proceed with caution!"
        confirm_prompt && ssh "$DEPLOY_HOST" -t "cd $DEPLOY_DIR && TAG=$TAG docker-compose down -v"
        ;;
    *)
        infomsg "Deploying docker-compose services on '$DEPLOY_HOST' with tag: $TAG"
        ssh "$DEPLOY_HOST" -t "cd $DEPLOY_DIR && $DEPLOY_CMD"
        ;;
    esac
}

main "$@"
