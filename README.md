# **Sessy CLI Tool**

Sessy is a Go-based CLI tool to manage and list session-related tasks efficiently.

---

## **Installation**

```bash
# 1. Clone the repository
git clone https://github.com/a-kumar5/sessy.git
cd sessy

# 2. Build the binary
go build -o sessy

# 3. Add the binary to your PATH (optional)
mv sessy /usr/local/bin/

### **List Files Command**

The `list` (or `ls`) command is used to display a list of files managed by Sessy.

#### **Usage**
```bash
sessy list

# **Create a YAML File**

### **Usage**

```bash
# Create a YAML file with a custom name
sessy yaml my-config

# Or use the alias
sessy yl my-config

# Create a default YAML file
sessy yaml