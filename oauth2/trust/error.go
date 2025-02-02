// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package trust

import (
	"net/http"

	"github.com/ory/fosite"
)

var ErrMissingRequiredParameter = &fosite.RFC6749Error{
	DescriptionField: "One of the required parameters is missing. Check your request parameters.",
	ErrorField:       "missing_required_parameter",
	CodeField:        http.StatusBadRequest,
}
