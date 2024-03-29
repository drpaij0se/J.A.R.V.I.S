package main

import (
	"log"
	"testing"

	"github.com/paij0se/J.A.R.V.I.S/src/cli"
	"github.com/paij0se/J.A.R.V.I.S/src/tools"
)

func TestOpenAI(t *testing.T) {
	var err error
	var config map[string]string
	if err = cli.CreateConfigDirectory(); err != nil {
		log.Fatal(err)
	}

	if config, err = cli.ReadYml(); err != nil {
		log.Fatal(err)
	}

	tools.SendTextToOPenAI("el rap esta en contra del racismo?", "gpt-3.5-turbo", config["auth"])
}

func TestMicrophone(t *testing.T) {
	filename := tools.RecordAudio()
	if filename == "" {
		t.Error("Error al grabar el audio")
	}
}

func TestTTS(t *testing.T) {
	text := `
Hello, This is a test

`
	filename := tools.TTS(text, "Lupe")
	tools.PlayAudio(filename)
}
