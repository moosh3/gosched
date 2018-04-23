// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package context provides a context that allows for an abstracted equivilant of
context switching between processes; in this case, Jobs
*/
package context

import (
	"context"
	"time"
)

var timeSlice int

func WithTimeout(ctx context.Context, timeout time.Time) context.Context {
	return context.WithValue(ctx, timeSlice, timeout)
}
