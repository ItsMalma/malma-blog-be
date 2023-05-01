package utilfile

import (
	"mime"
	"net/http"
)

func GetExtensionFromBytes(bytes []byte) (string, error) {
	contentType := http.DetectContentType(bytes)
	extensions, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", err
	}
	if len(extensions) < 1 {
		return "", nil
	}
	return extensions[0], nil
}
