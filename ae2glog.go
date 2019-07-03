package ae2glog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type httpRequest struct { // https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#HttpRequest
	RequestMethod string `json:"requestMethod"`
	RequestURL    string `json:"requestUrl"`
	RequestSize   string `json:"requestSize"`
	UserAgent     string `json:"userAgent"`
	Referer       string `json:"referer"`
}

type operation struct { // https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogEntryOperation
	ID       string `json:"id"`
	Producer string `json:"producer"`
	First    bool   `json:"first"`
	Last     bool   `json:"last"`
}

// JSONPayload is json payload.
type JSONPayload struct { // https://cloud.google.com/logging/docs/agent/configuration#special-fields
	TraceID     string      `json:"traceId"`
	Trace       string      `json:"logging.googleapis.com/trace"`
	Message     string      `json:"message"`
	Severity    string      `json:"severity"`
	SpanID      string      `json:"logging.googleapis.com/spanId"`
	Operation   operation   `json:"logging.googleapis.com/operation"`
	HTTPRequest httpRequest `json:"httpRequest"`
}

type contextKey string

const tokenContextKey contextKey = "AppEngine2ndGenerationLogger-JsonPayload"

// NewContext makes context.
func NewContext(req *http.Request) (ctx context.Context) {
	return AddContext(req.Context(), req)
}

// AddContext send HTTP Request log.
func AddContext(origctx context.Context, req *http.Request) (ctx context.Context) {
	projectID := os.Getenv(`GOOGLE_CLOUD_PROJECT`)
	ctcs := strings.SplitN(strings.SplitN(req.Header.Get("X-Cloud-Trace-Context"), `;`, 2)[0], `/`, 2)
	if len(ctcs) != 0 {
		if len(ctcs) == 1 {
			ctcs = append(ctcs, ``)
		}

		e := JSONPayload{
			TraceID: ctcs[0],
			Trace:   `projects/` + projectID + `/traces/` + ctcs[0],
			SpanID:  ctcs[1],
			Operation: operation{
				ID:       req.Header.Get("X-Appengine-Request-Log-Id"),
				Producer: "appengine.googleapis.com/request_id",
			},
		}

		ctx = context.WithValue(origctx, tokenContextKey, e)

		payload := e
		payload.HTTPRequest = httpRequest{
			RequestMethod: req.Method,
			RequestURL:    req.URL.String(),
			RequestSize:   strconv.Itoa(int(req.ContentLength)),
			UserAgent:     req.UserAgent(),
			Referer:       req.Referer(),
		}

		json.NewEncoder(os.Stderr).Encode(payload)
	}
	return
}

// Defaultf send Application log.
func Defaultf(ctx context.Context, format string, a ...interface{}) {
	payload, ok := ctx.Value(tokenContextKey).(JSONPayload)
	if !ok {
		fmt.Printf("Default: "+format, a...)
		return
	}
	payload.Message = fmt.Sprintf(format, a...)
	payload.Severity = "DEFAULT"
	json.NewEncoder(os.Stdout).Encode(payload)
}

// Debugf send Application log.
func Debugf(ctx context.Context, format string, a ...interface{}) {
	payload, ok := ctx.Value(tokenContextKey).(JSONPayload)
	if !ok {
		fmt.Printf("Debug: "+format, a...)
		return
	}
	payload.Message = fmt.Sprintf(format, a...)
	payload.Severity = "DEBUG"

	json.NewEncoder(os.Stdout).Encode(payload)
}

// Infof send Application log.
func Infof(ctx context.Context, format string, a ...interface{}) {
	payload, ok := ctx.Value(tokenContextKey).(JSONPayload)
	if !ok {
		fmt.Printf("Info: "+format, a...)
		return
	}
	payload.Message = fmt.Sprintf(format, a...)
	payload.Severity = "INFO"

	json.NewEncoder(os.Stdout).Encode(payload)
}

// Noticef send Application log.
func Noticef(ctx context.Context, format string, a ...interface{}) {
	payload, ok := ctx.Value(tokenContextKey).(JSONPayload)
	if !ok {
		fmt.Printf("Notice: "+format, a...)
		return
	}
	payload.Message = fmt.Sprintf(format, a...)
	payload.Severity = "NOTICE"

	json.NewEncoder(os.Stdout).Encode(payload)
}

// Warningf send Application log.
func Warningf(ctx context.Context, format string, a ...interface{}) {
	payload, ok := ctx.Value(tokenContextKey).(JSONPayload)
	if !ok {
		fmt.Printf("Warning: "+format, a...)
		return
	}
	payload.Message = fmt.Sprintf(format, a...)
	payload.Severity = "WARNING"

	json.NewEncoder(os.Stdout).Encode(payload)
}

// Errorf send Application log.
func Errorf(ctx context.Context, format string, a ...interface{}) {
	payload, ok := ctx.Value(tokenContextKey).(JSONPayload)
	if !ok {
		fmt.Printf("Error: "+format, a...)
		return
	}
	payload.Message = fmt.Sprintf(format, a...)
	payload.Severity = "ERROR"

	json.NewEncoder(os.Stdout).Encode(payload)
}

// Criticalf send Application log.
func Criticalf(ctx context.Context, format string, a ...interface{}) {
	payload, ok := ctx.Value(tokenContextKey).(JSONPayload)
	if !ok {
		fmt.Printf("Critical: "+format, a...)
		return
	}
	payload.Message = fmt.Sprintf(format, a...)
	payload.Severity = "CRITICAL"

	json.NewEncoder(os.Stdout).Encode(payload)
}

// Alertf send Application log.
func Alertf(ctx context.Context, format string, a ...interface{}) {
	payload, ok := ctx.Value(tokenContextKey).(JSONPayload)
	if !ok {
		fmt.Printf("Alert: "+format, a...)
		return
	}
	payload.Message = fmt.Sprintf(format, a...)
	payload.Severity = "ALERT"

	json.NewEncoder(os.Stdout).Encode(payload)
}

// Emergencyf send Application log.
func Emergencyf(ctx context.Context, format string, a ...interface{}) {
	payload, ok := ctx.Value(tokenContextKey).(JSONPayload)
	if !ok {
		fmt.Printf("Emergency: "+format, a...)
		return
	}
	payload.Message = fmt.Sprintf(format, a...)
	payload.Severity = "EMERGENCY"

	json.NewEncoder(os.Stdout).Encode(payload)
}
