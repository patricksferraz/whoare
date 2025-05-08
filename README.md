# WhoAre - Employee Management System

<div align="center">

![Go](https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-2.0+-000000?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-24.0+-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/Kubernetes-1.28+-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/whoare?style=for-the-badge)](https://goreportcard.com/report/github.com/patricksferraz/whoare)

</div>

## 📝 Description

WhoAre is a modern employee management system built with Go, designed to help organizations manage their workforce effectively. It provides a robust platform for employee registration, profile management, and skill tracking.

## ✨ Features

- 🔍 Employee Search and Discovery
- 👤 Employee Profile Management
- 🎯 Skill Tracking and Management
- 🔐 Secure Authentication
- 📱 Responsive Web Interface
- 🚀 High Performance with Fiber Framework
- 🐳 Docker and Kubernetes Support

## 🚀 Getting Started

### Prerequisites

- Go 1.17 or higher
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/patricksferraz/whoare.git
cd whoare
```

2. Copy the environment file and configure it:
```bash
cp .env.example .env
```

3. Run with Docker Compose:
```bash
docker-compose up -d
```

Or run locally:
```bash
make run
```

## 🛠️ Development

### Project Structure

```
.
├── app/            # Application layer
│   └── front/      # Frontend handlers and views
├── config/         # Configuration files
├── domain/         # Domain models and services
├── infra/          # Infrastructure layer
├── k8s/            # Kubernetes manifests
└── utils/          # Utility functions
```

### Available Make Commands

```bash
make run        # Run the application
make build      # Build the application
make test       # Run tests
make lint       # Run linters
```

## 🐳 Docker Support

The application is containerized and can be run using Docker Compose:

```bash
docker-compose up -d
```

## ☸️ Kubernetes Deployment

Kubernetes manifests are available in the `k8s/` directory. Deploy using:

```bash
kubectl apply -f k8s/
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Patrick Ferraz** - *Initial work* - [patricksferraz](https://github.com/patricksferraz)

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber) - Fast and efficient web framework
- [GORM](https://gorm.io/) - The fantastic ORM library for Go
- [Air](https://github.com/cosmtrek/air) - Live reload for Go applications
