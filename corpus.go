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
		300: {
			300,
			"Multiple Choice",
			"The request has more than one possible response. The user-agent or user should choose one of them. (There is no standardized way of choosing one of the responses, but HTML links to the possibilities are recommended so the user can pick.)",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/300",
		},
		301: {
			301,
			"Moved Permanently",
			"The URL of the requested resource has been changed permanently. The new URL is given in the response.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/301",
		},
		302: {
			302,
			"Found",
			"This response code means that the URI of requested resource has been changed temporarily. Further changes in the URI might be made in the future. Therefore, this same URI should be used by the client in future requests.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/302",
		},
		303: {
			303,
			"See Other",
			"The server sent this response to direct the client to get the requested resource at another URI with a GET request.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/303",
		},
		304: {
			304,
			"Not Modified",
			"This is used for caching purposes. It tells the client that the response has not been modified, so the client can continue to use the same cached version of the response.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/304",
		},
		305: {
			305,
			"Use Proxy",
			"Defined in a previous version of the HTTP specification to indicate that a requested response must be accessed by a proxy. It has been deprecated due to security concerns regarding in-band configuration of a proxy.",
			"",
		},
		306: {
			306,
			"unused",
			"This response code is no longer used; it is just reserved. It was used in a previous version of the HTTP/1.1 specification.",
			"",
		},
		307: {
			307,
			"Temporary Redirect",
			"The server sends this response to direct the client to get the requested resource at another URI with same method that was used in the prior request. This has the same semantics as the 302 Found HTTP response code, with the exception that the user agent must not change the HTTP method used: If a POST was used in the first request, a POST must be used in the second request.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/307",
		},
		308: {
			308,
			"Permanent Redirect",
			"This means that the resource is now permanently located at another URI, specified by the Location: HTTP Response header. This has the same semantics as the 301 Moved Permanently HTTP response code, with the exception that the user agent must not change the HTTP method used: If a POST was used in the first request, a POST must be used in the second request.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/308",
		},
		400: {
			400,
			"Bad Request",
			"The server could not understand the request due to invalid syntax.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400",
		},
		401: {
			401,
			"Unauthorized",
			`Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated". That is, the client must authenticate itself to get the requested response.`,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401",
		},
		402: {
			402,
			"Payment Required",
			"This response code is reserved for future use. The initial aim for creating this code was using it for digital payment systems, however this status code is used very rarely and no standard convention exists.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/402",
		},
		403: {
			403,
			"Forbidden",
			"The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource. Unlike 401, the client's identity is known to the server.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403",
		},
		404: {
			404,
			"Not Found",
			"The server can not find the requested resource. In the browser, this means the URL is not recognized. In an API, this can also mean that the endpoint is valid but the resource itself does not exist. Servers may also send this response instead of 403 to hide the existence of a resource from an unauthorized client. This response code is probably the most famous one due to its frequent occurrence on the web.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404",
		},
		405: {
			405,
			"Method Not Allowed",
			"The request method is known by the server but has been disabled and cannot be used. For example, an API may forbid DELETE-ing a resource. The two mandatory methods, GET and HEAD, must never be disabled and should not return this error code.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405",
		},
		406: {
			406,
			"Not Acceptable",
			"This response is sent when the web server, after performing server-driven content negotiation, doesn't find any content that conforms to the criteria given by the user agent.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406",
		},
		407: {
			407,
			"Proxy Authentication Required",
			"This is similar to 401 but authentication is needed to be done by a proxy.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/407",
		},
		408: {
			408,
			"Request Timeout",
			"This response is sent on an idle connection by some servers, even without any previous request by the client. It means that the server would like to shut down this unused connection. This response is used much more since some browsers, like Chrome, Firefox 27+, or IE9, use HTTP pre-connection mechanisms to speed up surfing. Also note that some servers merely shut down the connection without sending this message.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408",
		},
		409: {
			409,
			"Conflict",
			"This response is sent when a request conflicts with the current state of the server.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409",
		},
		410: {
			410,
			"Gone",
			`This response is sent when the requested content has been permanently deleted from server, with no forwarding address. Clients are expected to remove their caches and links to the resource. The HTTP specification intends this status code to be used for "limited-time, promotional services". APIs should not feel compelled to indicate resources that have been deleted with this status code.`,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/410",
		},
		411: {
			411,
			"Length Required",
			"Server rejected the request because the Content-Length header field is not defined and the server requires it.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/411",
		},
		412: {
			412,
			"Precondition Failed",
			"The client has indicated preconditions in its headers which the server does not meet.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/412",
		},
		413: {
			413,
			"Payload Too Large",
			"Request entity is larger than limits defined by server; the server might close the connection or return an Retry-After header field.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413",
		},
		414: {
			414,
			"URI Too Long",
			"The URI requested by the client is longer than the server is willing to interpret.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/414",
		},
		415: {
			415,
			"Unsupported Media Type",
			"The media format of the requested data is not supported by the server, so the server is rejecting the request.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415",
		},
		416: {
			416,
			"Range Not Satisfiable",
			"The range specified by the Range header field in the request can't be fulfilled; it's possible that the range is outside the size of the target URI's data.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/416",
		},
		417: {
			417,
			"Expectation Failed",
			"This response code means the expectation indicated by the Expect request header field can't be met by the server.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/417",
		},
		418: {
			418,
			"I'm a teapot",
			"The server refuses the attempt to brew coffee with a teapot.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/418",
		},
		421: {
			421,
			"Misdirected Request",
			"The request was directed at a server that is not able to produce a response. This can be sent by a server that is not configured to produce responses for the combination of scheme and authority that are included in the request URI.",
			"",
		},
		422: {
			422,
			"Unprocessable Entity",
			"The request was well-formed but was unable to be followed due to semantic errors.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422",
		},
		423: {
			423,
			"Locked",
			"The resource that is being accessed is locked.",
			"",
		},
		424: {
			424,
			"Failed Dependency",
			"The request failed due to failure of a previous request.",
			"",
		},
		425: {
			425,
			"Too Early",
			"Indicates that the server is unwilling to risk processing a request that might be replayed.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/425",
		},
		426: {
			426,
			"Upgrade Required",
			"The server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol. The server sends an Upgrade header in a 426 response to indicate the required protocol(s).",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/426",
		},
		428: {
			428,
			"Precondition Required",
			"The origin server requires the request to be conditional. This response is intended to prevent the 'lost update' problem, where a client GETs a resource's state, modifies it, and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflict.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/428",
		},
		429: {
			429,
			"Too Many Requests",
			`The user has sent too many requests in a given amount of time ("rate limiting").`,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429",
		},
		431: {
			431,
			"Request Header Fields Too Large",
			"The server is unwilling to process the request because its header fields are too large. The request may be resubmitted after reducing the size of the request header fields.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/431",
		},
		451: {
			451,
			"Unavailable For Legal Reasons",
			"The user-agent requested a resource that cannot legally be provided, such as a web page censored by a government.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/451",
		},
		500: {
			500,
			"Internal Server Error",
			"The server has encountered a situation it doesn't know how to handle.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500",
		},
		501: {
			501,
			"Not Implemented",
			"The request method is not supported by the server and cannot be handled. The only methods that servers are required to support (and therefore that must not return this code) are GET and HEAD.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501",
		},
		502: {
			502,
			"Bad Gateway",
			"This error response means that the server, while working as a gateway to get a response needed to handle the request, got an invalid response.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502",
		},
		503: {
			503,
			"Service Unavailable",
			"The server is not ready to handle the request. Common causes are a server that is down for maintenance or that is overloaded. Note that together with this response, a user-friendly page explaining the problem should be sent. This responses should be used for temporary conditions and the Retry-After: HTTP header should, if possible, contain the estimated time before the recovery of the service. The webmaster must also take care about the caching-related headers that are sent along with this response, as these temporary condition responses should usually not be cached.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503",
		},
		504: {
			504,
			"Gateway Timeout",
			"This error response is given when the server is acting as a gateway and cannot get a response in time.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504",
		},
		505: {
			505,
			"HTTP Version Not Supported",
			"The HTTP version used in the request is not supported by the server.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/505",
		},
		506: {
			506,
			"Variant Also Negotiates",
			"The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/506",
		},
		507: {
			507,
			"Insufficient Storage",
			"The method could not be performed on the resource because the server is unable to store the representation needed to successfully complete the request.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/507",
		},
		508: {
			508,
			"Loop Detected",
			"The server detected an infinite loop while processing the request.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/508",
		},
		510: {
			510,
			"Not Extended",
			"Further extensions to the request are required for the server to fulfil it.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/510",
		},
		511: {
			511,
			"Network Authentication Required",
			"The 511 status code indicates that the client needs to authenticate to gain network access.",
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/511",
		},
	}
)

var errCodeNotFound = errors.New("code not found")

func getDetails(code int) (*httpcode, error) {
	httpCode, ok := corpus[code]
	if !ok {
		return nil, errCodeNotFound
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
