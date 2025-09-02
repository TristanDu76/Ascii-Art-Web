# Ascii-Art-Web

## Description

Ascii-Art-Web is a web application written in Go that lets users convert text into ASCII art using one of three banner styles: `standard`, `shadow`, or `thinkertoy`. The interface allows you to:

- Enter custom text
- Choose a banner style
- Submit and view the ASCII art
- Download the result as a `.txt` file

The app runs a local HTTP server, renders pages using HTML templates, and handles form submissions using standard Go packages.

## Usage

1. Run the server:
   ```bash
   go run main.go
