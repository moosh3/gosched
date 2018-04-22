// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gosched

import "errors"

// AnError is an error instance useful for testing purposes. If
// the code does not care about error specifics, this error should
// be used to make the test code more readable
var AnError = errors.New("General error for testing")
