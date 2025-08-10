package usecase

import (
	"backend/internal/url/entity"
	"backend/pkg/constant"
	"context"
	"net/url"
	"strings"
)

type urlUsecase struct {
}

func NewURLUsecase() URLUsecase {
	return &urlUsecase{}
}

func (uc *urlUsecase) CleanupURL(ctx context.Context, input entity.URL) (entity.URL, error) {
	u, err := url.Parse(input.InitialURL)
	if err != nil {
		return entity.URL{}, err
	}

	if input.Operation == constant.URLOperationRedirection ||
		input.Operation == constant.URLOperationAll {

		// Normalize scheme e.g. http:// or https://
		if u.Scheme != "http" && u.Scheme != "https" {
			u.Scheme = "https"
		}

		// Normalize host e.g. www.byfood.com
		if u.Host != "" {
			uc.normalizeHost(u, u.Host)
		}

		// Normalize path e.g. /home
		if !strings.HasPrefix(u.Path, "/") { // host + path
			urlSplit := strings.Split(u.Path, "/")

			uc.normalizeHost(u, urlSplit[0]) // Extract host from path.

			if len(urlSplit) > 1 {
				u.Path = "/" + strings.ToLower(strings.Join(urlSplit[1:], "/"))
			}

		} else { // path only
			u.Path = strings.ToLower(u.Path)
		}

		// Normalize query e.g. ?query=abc
		if u.RawQuery != "" {
			u.RawQuery = strings.ToLower(u.RawQuery)
		}

		// Normalize fragment e.g. #fragment
		if u.Fragment != "" {
			u.Fragment = strings.ToLower(u.Fragment)
		}
	}

	if input.Operation == constant.URLOperationCanonical ||
		input.Operation == constant.URLOperationAll {

		// Remove query and fragment.
		u.RawQuery = ""
		u.Fragment = ""

		// Remove trailing slash.
		u.Path = strings.TrimSuffix(u.Path, "/")
	}

	return entity.URL{
		ProcessedURL: u.String(),
	}, nil
}

func (uc *urlUsecase) normalizeHost(u *url.URL, host string) {
	host = strings.ToLower(host)
	if !strings.HasPrefix(host, "www.") {
		host = "www." + host
	}
	u.Host = host
}
