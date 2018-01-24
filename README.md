# Derision SMTP

[![GoDoc](https://godoc.org/github.com/efritz/derision-smtp?status.svg)](https://godoc.org/github.com/efritz/derision-smtp)
[![Build Status](https://secure.travis-ci.org/efritz/derision-smtp.png)](http://travis-ci.org/efritz/derision-smtp)
[![Code Coverage](http://codecov.io/github/efritz/derision-smtp/coverage.svg?branch=master)](http://codecov.io/github/efritz/derision-smtp?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/efritz/derision-smtp)](https://goreportcard.com/report/github.com/efritz/derision-smtp)

A mock SMTP server which reflects its payloads to a Derision API.

## Usage

Simply run the following command to get a (rather fragile) SMTP server
(which has currently only been tested against python's SMTP server
sendmail function) and a Derision API. The email endpoint has the following
fields in its payload.


| Name         | Description |
| ------------ | ----------- |
| Recipient    | The recipient's email |
| Subject      | The email subject (pulled from body) |
| Body		   | The email body (raw text) |

## License

Copyright (c) 2018 Eric Fritz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
