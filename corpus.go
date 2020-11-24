package main

import (
	"errors"
)

var (
	corpus = map[int]*httpcode{
		100: {
			100,
			"Continue",
			"This interim response indicates that everything so far is OK and that the client should continue the request, or ignore the response if the request is already finished.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/100",
		},
		101: {
			101,
			"Switching Protocols",
			"This code is sent in response to an Upgrade request header from the client, and indicates the protocol the server is switching to.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/101",
		},
		102: {
			102,
			"Processing",
			"This code indicates that the server has received and is processing the request, but no response is available yet.",
			"",
		},
		103: {
			103,
			"Early Hints",
			"This status code is primarily intended to be used with the Link header, letting the user agent start preloading resources while the server prepares a response.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/103",
		},
		200: {
			200,
			"OK",
			`The request has succeeded. The meaning of the success depends on the HTTP method:
			GET: The resource has been fetched and is transmitted in the message body.
			HEAD: The entity headers are in the message body.
			PUT or POST: The resource describing the result of the action is transmitted in the message body.
			TRACE: The message body contains the request message as received by the server.`,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200",
		},
		201: {
			201,
			"Created",
			"The request has succeeded and a new resource has been created as a result. This is typically the response sent after POST requests, or some PUT requests.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201",
		},
		202: {
			202,
			"Accepted",
			"The request has been received but not yet acted upon. It is noncommittal, since there is no way in HTTP to later send an asynchronous response indicating the outcome of the request. It is intended for cases where another process or server handles the request, or for batch processing.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/202",
		},
		203: {
			203,
			"Non-Authoritative Information",
			`This response code means the returned meta-information is not exactly the same as is available from the origin server, but is collected from a local or a third-party copy. This is mostly used for mirrors or backups of another resource. Except for that specific case, the "200 OK" response is preferred to this status.`,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/203",
		},
		204: {
			204,
			"No Content",
			"There is no content to send for this request, but the headers may be useful. The user-agent may update its cached headers for this resource with the new ones.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204",
		},
		205: {
			205,
			"Reset Content",
			"Tells the user-agent to reset the document which sent this request.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/205",
		},
		206: {
			206,
			"Partial Content",
			"This response code is used when the Range header is sent from the client to request only part of a resource.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/206",
		},
		207: {
			207,
			"Multi-Status",
			"Conveys information about multiple resources, for situations where multiple status codes might be appropriate.",
			"",
		},
		208: {
			208,
			"Already Reported",
			"Used inside a <dav:propstat> response element to avoid repeatedly enumerating the internal members of multiple bindings to the same collection.",
			"",
		},
		226: {
			226,
			"IM Used",
			"The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.",
			"",
		},
	}
)

func getDetails(code int) (*httpcode, error) {
	httpCode, ok := corpus[code]
	if !ok {
		return nil, errors.New("code not found")
	}
	return httpCode, nil
}

func getReasonPhrase(code int) (string, error) {
	httpCode, err := getDetails(code)
	if err != nil {
		return "", err
	}
	return httpCode.reasonPhrase, nil
}
