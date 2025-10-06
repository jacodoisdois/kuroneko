# Kuronek- **📁 Smart Folder Detection**: Automatically detects and processes manga chapter folders 🐾

A manga CBZ creator inspired by HakuNeko. Transform your downloaded manga chapters into organized CBZ files.

<img src="assets/images/logo.jpg" alt="Kuroneko Logo" width="200" height="200">

## ✨ Features

- **🎨 Modern UI**: Circular logo and intuitive design
- **� Smart Folder Detection**: Automatically detects and processes manga chapter folders
- **📊 Progress Tracking**: Real-time progress bar during CBZ creation
- **💾 Memory Persistence**: Remembers your last selected folder between sessions
- **🖼️ Asset Bundling**: All assets embedded in the binary for easy distribution
- **🍎 macOS Integration**: Native "Open in Finder" functionality
- **🔄 Cross-Platform**: Works on Windows, macOS, and Linux

## 🚀 Installation

### Prerequisites
- Go 1.19 or later
- Fyne framework (automatically installed via go.mod)

### Build from Source
```bash
# Clone the repository
git clone https://github.com/jacodoisdois/kuroneko.git
cd kuroneko

# Install dependencies
go mod download

# Build the application
go build .

# Run the application
./kuroneko
```

### macOS App Bundle
```bash
# Build macOS app bundle
fyne package -os darwin -icon assets/icons/paw-icon-2.png
```

## 📖 Usage

1. **Launch Kuroneko**: Run the application
2. **Select Folder**: Click "Browse" to choose a folder containing manga chapters
3. **Create CBZ**: Click "Create CBZ" to process all chapter folders
4. **Monitor Progress**: Watch the progress bar as files are created
5. **Access Files**: Use "Open in Finder" to view your CBZ files

### Folder Structure
Kuroneko expects a folder structure like:
```
MangaFolder/
├── chapter1/
│   ├── page1.jpg
│   ├── page2.png
│   └── ...
├── chapter2/
│   ├── page1.jpg
│   └── ...
└── chapter3/
    └── ...
```

## 🛠️ Development

### Prerequisites
- Go 1.19+
- Fyne v2.4.5+

### Setup Development Environment
```bash
# Clone and setup
git clone https://github.com/jacodoisdois/kuroneko.git
cd kuroneko

# Install dependencies
go mod tidy

# Run in development mode
go run .
```

### Project Structure
```
kuroneko/
├── assets/                 # Source assets
│   ├── images/            # Logo and images
│   ├── icons/             # Icons (SVG/PNG)
│   └── fonts/             # Custom fonts
├── internal/
│   ├── app.go             # Application entry point
│   ├── handler/           # CBZ processing logic
│   └── ui/                # User interface
│       ├── bundled.go     # Generated bundled assets
│       ├── main_window.go # Main UI components
│       └── theme.go       # Custom theme
├── cmd/                   # Command line interface
├── test-data/             # Sample test data
└── main.go               # Main entry point
```

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines
- Follow Go best practices
- Add tests for new features
- Update documentation
- Ensure cross-platform compatibility

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by [HakuNeko](https://github.com/manga-download/hakuneko)
- Built with [Fyne](https://fyne.io/) - Cross-platform GUI toolkit
- Icons and assets designed for the project

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/jacodoisdois/kuroneko/issues)
- **Discussions**: [GitHub Discussions](https://github.com/jacodoisdois/kuroneko/discussions)

---

**🐾 Happy reading with Kuroneko!**
