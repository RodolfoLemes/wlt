package timehelper

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "when parsing a valid date, it should return the expected time",
			args: args{
				date: "2023-12-21",
			},
			want:    time.Date(2023, 12, 21, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name: "when parsing an invalid date, it should return an error",
			args: args{
				date: "invalid-date",
			},
			want:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("ParseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDateTime(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "when parsing a valid date-time, it should return the expected time",
			args: args{
				date: "2023-12-21T12:34:56Z",
			},
			want:    time.Date(2023, 12, 21, 12, 34, 56, 0, time.UTC),
			wantErr: false,
		},
		{
			name: "when parsing an invalid date-time, it should return an error",
			args: args{
				date: "invalid-date-time",
			},
			want:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateTime(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("ParseDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDateTimePostgres(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "when parsing a valid date-time postgres, it should return the expected time",
			args: args{
				date: "2023-12-21 12:34:56",
			},
			want:    time.Date(2023, 12, 21, 12, 34, 56, 0, time.UTC),
			wantErr: false,
		},
		{
			name: "when parsing an invalid date-time postgres, it should return an error",
			args: args{
				date: "invalid-date-time-postgres",
			},
			want:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateTimePostgres(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateTimePostgres() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("ParseDateTimePostgres() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDate(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "when formatting a valid date, it should return the expected string",
			args: args{
				date: time.Date(2023, 12, 21, 0, 0, 0, 0, time.UTC),
			},
			want: "2023-12-21",
		},
		{
			name: "when formatting another valid date, it should return the expected string",
			args: args{
				date: time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
			},
			want: "2020-01-15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDate(tt.args.date); got != tt.want {
				t.Errorf("FormatDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDateTime(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "when formatting a valid date-time, it should return the expected string",
			args: args{
				date: time.Date(2023, 12, 21, 12, 34, 56, 0, time.UTC),
			},
			want: "2023-12-21T12:34:56Z",
		},
		{
			name: "when formatting another valid date-time, it should return the expected string",
			args: args{
				date: time.Date(2020, 1, 15, 8, 45, 30, 0, time.UTC),
			},
			want: "2020-01-15T08:45:30Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDateTime(tt.args.date); got != tt.want {
				t.Errorf("FormatDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDateTimePostgres(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "when formatting a valid date-time postgres, it should return the expected string",
			args: args{
				date: time.Date(2023, 12, 21, 12, 34, 56, 0, time.UTC),
			},
			want: "2023-12-21 12:34:56",
		},
		{
			name: "when formatting another valid date-time postgres, it should return the expected string",
			args: args{
				date: time.Date(2020, 1, 15, 8, 45, 30, 0, time.UTC),
			},
			want: "2020-01-15 08:45:30",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDateTimePostgres(tt.args.date); got != tt.want {
				t.Errorf("FormatDateTimePostgres() = %v, want %v", got, tt.want)
			}
		})
	}
}
