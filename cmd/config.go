package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Culqi struct {
	MerchantCode string `yaml:"merchant_code"`
	ApiKey       string `yaml:"api_key"`
}

type MailJet struct {
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	FromName  string `yaml:"from_name"`
	FromEmail string `yaml:"from_email"`
}

type Store struct {
	Filepath string `yaml:"filepath"`
}

type Company struct {
	Name             string `yaml:"name"`
	PhoneSupport     string `yaml:"phone_support"`
	EmailSupport     string `yaml:"email_support"`
	TemplateFilename string `yaml:"template_filename"`
	MaxAttempts      int    `yaml:"max_attempts"`
}

type UCulqiConfiguration struct {
	Culqi   *Culqi   `yaml:"culqi"`
	MailJet *MailJet `yaml:"mailjet"`
	Store   *Store   `yaml:"store"`
	Company *Company `yaml:"company"`
}

func parseConfig(filename string) (*UCulqiConfiguration, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := new(UCulqiConfiguration)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
