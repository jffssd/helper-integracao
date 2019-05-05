package lib

import (
    "fmt"
    "os"
    "gopkg.in/ini.v1"
)

func LerINI(filepath string) {
    cfg, err := ini.Load(filepath)
    if err != nil {
        fmt.Printf("Falha ao tentar ler o arquivo: %v", err)
        os.Exit(1)
    }

    // Classic read of values, default section can be represented as empty string
    fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
    fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

    // Let's do some candidate value limitation
    fmt.Println("Server Protocol:",
        cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
    // Value read that is not in candidates will be discarded and fall back to given default value
    fmt.Println("Email Protocol:",
        cfg.Section("server").Key("email-protocol").In("smtp", []string{"imap", "smtp"}))

    // Try out auto-type conversion
    fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(5432))
    fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))
}

func EscreverINI(filepath string) {
    cfg, err := ini.Load(filepath)
    if err != nil {
        fmt.Printf("Falha ao tentar ler o arquivo: %v", err)
        os.Exit(1)
    }

    // Now, make some changes and save it
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.Section("server").Key("http_port").SetValue("8080")
	cfg.Section("config").Key("host").SetValue("http://localhost/")

    cfg.SaveTo("application/controller/my.ini")
}