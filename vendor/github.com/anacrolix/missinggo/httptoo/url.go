package httptoo

import (
	"net/http"
	"net/url"
	"path"
)

// Deep copies a URL.
func CopyURL(u *url.URL) (ret *url.URL) {
	ret = new(url.URL)
	*ret = *u
	if u.User != nil {
		ret.User = new(url.Userinfo)
		*ret.User = *u.User
	}
	return
}

// Reconstructs the URL that would have produced the given Request.
// Request.URLs are not fully populated in http.Server handlers.
func RequestedURL(r *http.Request) (ret *url.URL) {
	ret = CopyURL(r.URL)
	ret.Host = r.Host
	ret.Scheme = OriginatingProtocol(r)
	return
}

// 	Scheme     string
// 	Opaque     string    // encoded opaque data
// 	User       *Userinfo // username and password information
// 	Host       string    // host or host:port
// 	Path       string
// 	RawPath    string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
// 	ForceQuery bool   // append a query ('?') even if RawQuery is empty
// 	RawQuery   string // encoded query values, without '?'
// 	Fragment   string // fragment for references, without '#'
// }

// Return the first URL extended with elements of the second, in the manner
// that occurs throughout my projects.
func AppendURL(u, v *url.URL) *url.URL {
	u = CopyURL(u)
	clobberString(&u.Scheme, v.Scheme)
	clobberString(&u.Host, v.Host)
	u.Path = path.Join(u.Path, v.Path)
	q := u.Query()
	for k, v := range v.Query() {
		q[k] = append(q[k], v...)
	}
	u.RawQuery = q.Encode()
	return u
}

func clobberString(s *string, value string) {
	if value != "" {
		*s = value
	}
}
