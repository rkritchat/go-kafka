package rule

import validation "github.com/go-ozzo/ozzo-validation/v4"

var FistName = []validation.Rule{validation.Required, validation.Length(5, 20)}
var LastName = []validation.Rule{validation.Required, validation.Length(5, 20)}
var Age = []validation.Rule{validation.Required, validation.Min(1), validation.Max(100)}