# ZooRheine

## Installation

Um das Projekt erfolgreich zu installieren und auszuführen, benötigst du den GO-Lang Compiler sowie Docker (inklusive Docker Compose). Stelle sicher, dass beide Tools auf deinem System installiert sind.

### Voraussetzungen

- **[GO-Lang](https://go.dev/)** – Die Programmiersprache, in der das Projekt entwickelt wurde.
- **[Docker](https://www.docker.com/)** – Für die Containerisierung und einfache Bereitstellung der Anwendung.

### Abhängigkeiten installieren

Bevor du das Projekt kompilieren kannst, müssen einige GO-Pakete installiert werden. Führe dazu folgenden Befehl aus:

```bash
go get github.com/golang-jwt/jwt/v5 github.com/go-sql-driver/mysql golang.org/x/crypto/bcrypt
```

### Projekt kompilieren und Docker-Container starten

Sobald alle Abhängigkeiten installiert sind, kannst du das Projekt kompilieren und den Docker-Container starten. Verwende dazu den folgenden Befehl:

```bash
sudo make
```

Dieser Befehl kompiliert das Projekt und startet die notwendigen Docker-Container, um die Anwendung auszuführen.

---
