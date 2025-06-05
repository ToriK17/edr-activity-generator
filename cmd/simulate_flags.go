package cmd

import "time"

var (
	count    int           // number of times to repeat the activity
	delay    time.Duration // delay between each repetition
	stream   time.Duration // duration to stream activity (used in --stream mode)
	duration time.Duration // future use: may support time-limited stress tests
)
