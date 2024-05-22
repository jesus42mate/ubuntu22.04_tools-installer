package main

import (
  "bufio"
  "fmt"
  "os"
  "golang.org/x/term"
)

func main() {
  var stdin *os.File = os.Stdin
  var scanner *bufio.Scanner = bufio.NewScanner(stdin)
  var terminal *term.Terminal = term.NewTerminal(stdin, "λ ")
  var FD int = int(stdin.Fd())

  userHomeDir, err := os.UserHomeDir()

  if err != nil {
    fmt.Printf("%s", err)
  } else {
    fmt.Println("User's home dir is:", userHomeDir)
  }

  err = os.Chdir("/home")
  if err != nil {
    fmt.Printf("%s", err)
  } else {
    ShellExec(terminal, "pwd")
    ChangeDir(userHomeDir)
    ShellExec(terminal, "mkdir", "Jammer")
  }

  terminal.Write([]byte("\nJammer: Helloo world! Im alive!\n"))

  fmt.Println("\nWelcome to Jammer, what would you like to do?")

  prevState, error := term.MakeRaw(FD)
  formalPanicNeeds := FormalPanicNeeds{ FD, prevState }
  if error != nil {
    FormalPanic(formalPanicNeeds, error)
  } else {
    terminal.Write([]byte("\nJammer: Terminal RAW MODE ACTIVATED (in a robot voice)\n"))
  }

  choices, exit := TermChoice([]string{
    "Update and Upgrade apt packet manager.",
    "Install Neovim from source(LTS).",
    "Install NVM (Node Version Manager).",
  }, terminal, "")

  if exit != nil {
    FormalPanic(formalPanicNeeds, exit)
  }

  if choices[0] {
    AptUpdate(terminal)
    AptUpgrade(terminal)
  }
  if (choices[1]) {
    InstallNeovim(scanner, terminal, userHomeDir)
  }

  defer term.Restore(FD, prevState)
  terminal.Write([]byte("\nGraceful shutdown.\n"))
  terminal.Write([]byte("Thanks for using Jammer!\n"))
  terminal.Write([]byte("----------------------\n"))
}


