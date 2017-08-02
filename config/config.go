package config

import "time"

type Multiline struct {
	Pattern string
	Negate  bool
	Match   string
}

type Prospector struct {
	Id         string     `config:"id"`
	GroupNames []string   `config:"groupnames"`
	Multiline  *Multiline `config:"multiline"`
}

type Config struct {
	GroupRefreshFrequency          time.Duration `config:"group_refresh_frequency"`
	StreamRefreshFrequency         time.Duration `config:"stream_refresh_frequency"`
	ReportFrequency                time.Duration `config:"report_frequency"`
	GroupNames                     []string      `config:"groupnames"`
	S3BucketName                   string        `config:"s3_bucket_name"`
	AWSRegion                      string        `config:"aws_region"`
	Prospectors                    []Prospector  `config:"prospectors"`
	HotStreamEventHorizon          time.Duration `config:"hot_stream_event_horizon"`
	StreamLastEventHorizon         time.Duration `config:"stream_last_event_horizon"`
	HotStreamEventRefreshFrequency time.Duration `config:"hot_stream_event_refresh_frequency"`
	StreamEventRefreshFrequency    time.Duration `config:"hot_stream_event_refresh_frequency"`
}
