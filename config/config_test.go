package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/api/config"
	"io/ioutil"
	"os"
	"path/filepath"
)

var exampleData = []byte(`
---
default_backend_url: http://www.google.com:80
port: 3000
db:
  database: "sqlite://tmp/api.db"
app_packages:
  provider: "local"
  filepath: "/tmp/api/app_packages"
  access_key_id: "access_key_id"
  access_key_secret: "access_key_secret"
  host: "s3.amazon.com"
  bucket_name: "app_packages"
`)

var expectedConfig = config.Config{
	DefaultBackendURL: "http://www.google.com:80",
	Port:              3000,
	DB: config.DbConfig{
		URI: "sqlite://tmp/api.db",
	},

	AppPackages: config.BlobstoreConfig{
		Provider:        "local",
		Filepath:        "/tmp/api/app_packages",
		AccessKeyId:     "access_key_id",
		AccessKeySecret: "access_key_secret",
		Host:            "s3.amazon.com",
		BucketName:      "app_packages",
	},
}

var _ = Describe("Configuration", func() {
	Context("When creating a config file from a byte slice", func() {
		It("Contains all of the set values", func() {
			c, err := config.New(exampleData)
			Expect(err).ToNot(HaveOccurred())
			Expect(c).To(Equal(expectedConfig))
		})
	})

	Context("When creating a config file from a file", func() {
		var filePath string

		BeforeEach(func() {
			dir, err := ioutil.TempDir("", "api_test")
			Expect(err).ToNot(HaveOccurred())

			filePath = filepath.Join(dir, "config.yml")

			err = ioutil.WriteFile(filePath, exampleData, os.FileMode(0600))
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.Remove(filePath)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Contains all of the set values", func() {
			c, err := config.NewFromFile(filePath)
			Expect(err).ToNot(HaveOccurred())
			Expect(c).To(Equal(expectedConfig))
		})
	})
})
