package utils

import "time"

var (
	defaultTZ = time.UTC
	pgFormat  = "2006-01-02 15:04:05.000000"
)

//go:generate mockgen -destination=../test_utils/mocks/utils/time.go . ITimeService
type ITimeService interface {
	GetTimeNow() time.Time
	GetUTCTimeNow() time.Time
	GetTimeISONow() string
	GetUTCTimeISONow() string
	GetTimeFromUnixMilliEpoch(t int64) time.Time
	GetUnixEpochMilliFromTime(t time.Time) int64
}

type TimeService struct {
	TZ      *time.Location
	Pattern string
}

type TimeServiceOption func(*TimeService)

func WithTimeZone(tz *time.Location) TimeServiceOption {
	return func(ts *TimeService) {
		ts.TZ = tz
	}
}

func NewTimeService(opts ...TimeServiceOption) TimeService {
	ts := &TimeService{
		TZ: defaultTZ,
	}

	// Loop through each option
	for _, opt := range opts {
		opt(ts)
	}

	return *ts
}

func (ts TimeService) GetTimeNow() time.Time {
	return time.Now().In(ts.TZ)
}

func (ts TimeService) GetUTCTimeNow() time.Time {
	return time.Now().UTC()
}

func (ts TimeService) GetUTCTimeNowMS() time.Time {
	return time.Now().UTC().Truncate(time.Millisecond)
}

func (ts TimeService) GetTimeISONow() string {
	return ts.GetTimeNow().Format(time.RFC3339)
}

func (ts TimeService) GetUTCTimeISONow() string {
	return ts.GetUTCTimeNow().Format(time.RFC3339)
}

func (ts TimeService) GetTimeFromUnixMilliEpoch(t int64) time.Time {
	return time.Unix(t/1000, 0).In(ts.TZ)
}

func (ts TimeService) GetUnixEpochMilliFromTime(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func ParseStringPGFormat(val string) (time.Time, error) {
	return time.Parse(pgFormat, val)
}
