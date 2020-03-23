package executor

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v2"
)

// create Config type to tbe used for accessing config information
type Config struct {
    Provider struct {
        Name string `yaml:"name"`
        Version string `yaml:"version"`
    } `yaml:"provider"`
    Job struct {
        Name string `yaml:"name"`
        Description string `yaml:"description"`
        Language string `yaml:"language"`
        File string `yaml:"file"`
        Schedule string `yaml:"schedule"`
    } `yaml:"job"`
}

// this function is used to handle an error and exit upon doing so
// parameters: err (an error to print out)
// returns: nil
func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}

// this function is used to read a file and return it's contents
// parameters: filename (a string of the path to the file)
// returns: *os.File (the file descriptor)
func readFile(filename string) *os.File {
    file, err := os.Open(filename)
    if err != nil {
        defer file.Close()
        processError(err)
        var tmp *os.File
        return tmp
    }
    contents := file
    return contents
}

// this function is used to parse a given YAML config
// parameters: cfg (a *Config to be decoded into), file, (am *os.File to build the decoder on)
// returns: boolean (true if parseable, false if otherwise)
func parseYaml(cfg *Config, file *os.File) bool {
    decoder:= yaml.NewDecoder(file)
    err := decoder.Decode(cfg)
    if err != nil {
        processError(err)
        return false
    }
    return true
}

// this function is used to get the YAML of a config
// parameters: filename (a string of the path to the file))
// returns: two strings (the language and file used by the job)
func getYaml(filename string) (string, string) {
    var cfg Config
    yamlFile := readFile(filename)
    successfulParse := parseYaml(&cfg, yamlFile)
    if successfulParse {
        return cfg.Job.Language, cfg.Job.File
    } else {
        return "Failed to read file.", ""
    }
}
