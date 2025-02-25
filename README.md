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

## Alternative Installation

Falls du den GO-Lang Compiler und die benötigten Bibliotheken nicht installieren möchtest, kannst du vorkompilierte Binärdateien für den Server herunterladen. Der Vorteil dieser Methode ist, dass du keine zusätzlichen Tools oder Abhängigkeiten benötigst. Allerdings bist du mit dieser Methode weniger flexibel.

### Vorkompilierte Binärdateien herunterladen

Lade die vorkompilierte Binärdatei mit folgendem Befehl herunter:

```bash
wget https://github.com/JanikGilzer/ZooRheine/releases/download/v1.0.0/go_build_ZooDaBa -O ./build/go_build_ZooDaBa
```

### Docker-Container bauen und starten

Nachdem du die Binärdatei heruntergeladen hast, kannst du die Docker-Container wie gewohnt bauen und starten:

```bash
sudo make docker-build
```

Diese Methode ermöglicht es dir, die Anwendung ohne den GO-Lang Compiler und die benötigten Bibliotheken auszuführen.

---
