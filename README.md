# Go_training

<p align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_LightBlue.svg" />  
</p>


# Link

- test

# Tutorial (Go)

https://go.dev/doc/tutorial/getting-started

## Go Playground 
https://go.dev/play/


# WorkSpace 만들어주기

```

- go.work // go version 관리
- tools  // LSP관련자료

// 작업공간 workspace
- go_lang
├── Go_training
│   ├── 01_Go_Tutorial
│   │   ├── 01_Hello_Go
│   │   │   ├── main.go
│   │   │   └── README.md
│   │   └── README.md
│   └── README.md
└── test_go
    └── main.go


```

https://github.com/golang/tools



# Go LSP Server

- LSP 설치하기 https://pkg.go.dev/golang.org/x/tools/gopls#section-readme

- https://github.com/golang/tools/tree/master/gopls

- https://langserver.org/

# Go WorkSpace 만들어주기


https://github.com/golang/tools/blob/master/gopls/doc/workspace.md

# ```.gitignore``` Go Lang

```
# macOS
.DS_Store

# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
```
