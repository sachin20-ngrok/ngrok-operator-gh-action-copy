package main

import "time"

type Config struct {
	Plans []Plan `yaml:"plans"`
}

type Plan struct {
	Name    string `yaml:"name"`
	Install struct {
		Helm *HelmInstall `yaml:"helm,omitempty"`
	} `yaml:"install"`
	Tests []*Test `yaml:"tests,omitempty"`
}

type HelmInstall struct {
	Chart       string            `yaml:"chart"`
	ReleaseName string            `yaml:"release-name"`
	Namespace   string            `yaml:"namespace"`
	Sets        map[string]string `yaml:"set,omitempty"`
	ValuesFiles []string          `yaml:"values-files,omitempty"`

	Results struct {
		StartTime time.Time     `yaml:"start_time"`
		Took      time.Duration `yaml:"took"`
		Err       error         `yaml:"error,omitempty"`
		Stdout    string        `yaml:"stdout"`
	} `yaml:"results,omitempty"`
}

type WaitURLReady struct {
	URL                string         `yaml:"url"`
	Timeout            *time.Duration `yaml:"timeout"`
	Interval           time.Duration  `yaml:"interval"`
	Retries            *int           `yaml:"retries,omitempty"`
	ExpectedStatusCode int            `yaml:"expected-status-code"`
	Results            struct {
		Requests struct {
			Total   int
			Success int
		}

		Timings struct {
			// TTA is the time to availability, the time it took for the URL to be available.
			TTA *time.Duration

			Latency struct {
				Max *time.Duration
				Min *time.Duration
			}
		}
	} `yaml:"results,omitempty"`
}

type Test struct {
	WaitURLReady *WaitURLReady `yaml:"wait-url-ready,omitempty"`
}
