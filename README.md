To install **Go (Golang)** on **Ubuntu Linux**, follow these commands:

---

### ✅ Method 1: Install from Ubuntu's official package (may not be the latest version)

```bash
sudo apt update
sudo apt install golang-go -y
```

* ✅ Easy to install
* ❌ Might not be the latest version

---

### ✅ Method 2: Install the latest version manually (recommended)

#### 1. Download the latest Go binary from the official website:

```bash
wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
```

> Replace `go1.24.4` with the latest version from [https://go.dev/dl](https://go.dev/dl)

#### 2. Remove any previous Go installation (optional but recommended):

```bash
sudo rm -rf /usr/local/go
```

#### 3. Extract the downloaded archive to `/usr/local`:

```bash
sudo tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz
```

#### 4. Add Go to your system PATH:

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

> If you're using `zsh`:

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc
```

---

### ✅ Verify Go installation:

```bash
go version
```

> Example output:

```
go version go1.24.4 linux/amd64
```
