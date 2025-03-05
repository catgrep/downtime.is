# downtime.is

[![Go Coverage](https://codecov.io/gh/bosdhill/downtime.is/branch/main/graph/badge.svg)](https://codecov.io/gh/bosdhill/downtime.is)

[![Go Report Card](https://goreportcard.com/badge/github.com/bosdhill/downtime.is)](https://goreportcard.com/report/github.com/bosdhill/downtime.is)

A simple SLA calculator inspired by [uptime.is](https://uptime.is). Calculate SLA percentages for various reporting periods based on downtime duration.

## Usage

Visit [downtime.is](https://downtime.is) followed by a duration to calculate SLA percentages:

- `downtime.is/100` - 100 seconds of downtime
- `downtime.is/1h` - 1 hour of downtime
- `downtime.is/30m` - 30 minutes of downtime
- `downtime.is/2d12h` - 2 days and 12 hours of downtime

### Example Output

For 1 hour of downtime:
- Daily reporting: 95.8333%
- Weekly reporting: 99.4048%
- Monthly reporting: 99.8620%
- Quarterly reporting: 99.9540%
- Yearly reporting: 99.9885%

## Development

### Prerequisites
- Go 1.24 or higher
- Docker and Docker Compose

### Local Development
```bash
make server  # Run the server locally
make test    # Run tests
```
