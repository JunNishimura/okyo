# okyo - お経

<p align='left'>
  <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/JunNishimura/okyo">
  <img alt="GitHub" src="https://img.shields.io/github/license/JunNishimura/okyo">
  <a href="https://github.com/JunNishimura/okyo/actions/workflows/test.yml"><img src="https://github.com/JunNishimura/okyo/actions/workflows/test.yml/badge.svg" alt="test"></a>
  <a href="https://goreportcard.com/report/github.com/JunNishimura/okyo"><img src="https://goreportcard.com/badge/github.com/JunNishimura/okyo" alt="Go Report Card"></a>
</p>

## 📖 Overview
`okyo` is a CLI tool to chant a hannya shingyo(般若心経).

## 💻 How to use
### 1. Install
#### Homebrew Tap
```
brew install JunNishimura/tap/okyo
```

#### go intall
```
go install github.com/JunNishimura/okyo@latest
```

### 2. Execute Command
```
$ okyo
```

## 🔨 Options
### `--count` or `-c`
Flag to set how many times to listen to the sutra repeatedly. Default value is 1 and negative value will repeat indefinitely.
```
$ okyo --count 2
```

### `--speed` or `-s`
Flag to set how fast to listen to the sutra. Default value is 1.
```
$ okyo --speed 1.5
```

## Thanks
The audio file for the `hannya shingyo(般若心経)` is the one read by the abbot of Jigenji(慈眼寺), which is available at the following website.

http://www.sakado-jigenji.jp/download.html


## 🪧 License
okyo is released under MIT License. See [MIT](https://raw.githubusercontent.com/JunNishimura/okyo/main/LICENSE)
