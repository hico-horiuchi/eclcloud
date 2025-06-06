package certificates

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

/*
List Certificates
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the certificate attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - Name of the resource
	// - This field accepts UTF-8 characters up to 3 bytes
	Name string `q:"name"`

	// - Description of the resource
	// - This field accepts UTF-8 characters up to 3 bytes
	Description string `q:"description"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`

	// - If `true` is set, information of the certificate file are displayed
	Details bool `q:"details"`

	// - CA certificate file upload status of the certificate
	CACertStatus string `q:"ca_cert_status"`

	// - SSL certificate file upload status of the certificate
	SSLCertStatus string `q:"ssl_cert_status"`

	// - SSL key file upload status of the certificate
	SSLKeyStatus string `q:"ssl_key_status"`
}

// ToCertificateListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToCertificateListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToCertificateListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of certificates.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToCertificateListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return CertificatePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Create Certificate
*/

// CreateOpts represents options used to create a new certificate.
type CreateOpts struct {

	// - Name of the certificate
	// - This field accepts UTF-8 characters up to 3 bytes
	Name string `json:"name,omitempty"`

	// - Description of the certificate
	// - This field accepts UTF-8 characters up to 3 bytes
	Description string `json:"description,omitempty"`

	// - Tags of the certificate
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// ToCertificateCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToCertificateCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "certificate")
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToCertificateCreateMap() (map[string]interface{}, error)
}

// Create accepts a CreateOpts struct and creates a new certificate using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCertificateCreateMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Post(createURL(c), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Show Certificate
*/

// Show retrieves a specific certificate based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string) (r ShowResult) {
	_, r.Err = c.Get(showURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Certificate
*/

// UpdateOpts represents options used to update a existing certificate.
type UpdateOpts struct {

	// - Name of the certificate
	// - This field accepts UTF-8 characters up to 3 bytes
	Name *string `json:"name,omitempty"`

	// - Description of the certificate
	// - This field accepts UTF-8 characters up to 3 bytes
	Description *string `json:"description,omitempty"`

	// - Tags of the certificate
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// ToCertificateUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToCertificateUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "certificate")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToCertificateUpdateMap() (map[string]interface{}, error)
}

// Update accepts a UpdateOpts struct and updates a existing certificate using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToCertificateUpdateMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Patch(updateURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Delete Certificate
*/

// Delete accepts a unique ID and deletes the certificate associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Upload Certificate File
*/

// UploadFileOpts represents options used to upload a file to a existing certificate.
type UploadFileOpts struct {

	// - Type of the certificate file to be uploaded
	// - Can be uploaded only once for each type
	Type string `json:"type"`

	// - Content of the certificate file to be uploaded
	// - The content must be Base64 encoded
	//   - The file size before encoding must be less than or equal to 16KB
	//   - The file format before encoding must be PEM
	//     - DER can be converted to PEM by using OpenSSL command
	// - The following key algorithms are supported
	//   - RSA 1024, 2048, 3072 and 4096 bits
	//   - ECDSA P-256 (prime256v1, secp256r1), P-384 (secp384r1) and P-521 (secp521r1)
	// - The content of `"ssl-cert"` and the content of `"ssl-key"` must be a pair (must be matched correctly)
	Content string `json:"content"`

	// - Passphrase of the certificate file to be uploaded
	// - This parameter can be set when 'type' is `"ssl-key"`
	Passphrase string `json:"passphrase,omitempty"`
}

// ToCertificateUploadFileMap builds a request body from UploadFileOpts.
func (opts UploadFileOpts) ToCertificateUploadFileMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// UploadFileOptsBuilder allows extensions to add additional parameters to the UploadFile request.
type UploadFileOptsBuilder interface {
	ToCertificateUploadFileMap() (map[string]interface{}, error)
}

// UploadFile accepts a UploadFileOpts struct and uploads a file to a existing certificate using the values provided.
func UploadFile(c *eclcloud.ServiceClient, id string, opts UploadFileOptsBuilder) (r UploadFileResult) {
	b, err := opts.ToCertificateUploadFileMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Post(uploadFileURL(c, id), b, nil, &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
