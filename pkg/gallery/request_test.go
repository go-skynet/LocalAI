package gallery_test

import (
	. "github.com/go-skynet/LocalAI/pkg/gallery"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gallery API tests", func() {
	Context("requests", func() {
		It("parses github with a branch", func() {
			req := GalleryModel{URL: "github:go-skynet/model-gallery/gpt4all-j.yaml@main"}
			str, err := req.DecodeURL()
			Expect(err).ToNot(HaveOccurred())
			Expect(str).To(Equal("https://raw.githubusercontent.com/go-skynet/model-gallery/main/gpt4all-j.yaml"))
		})
		It("parses github without a branch", func() {
			req := GalleryModel{URL: "github:go-skynet/model-gallery/gpt4all-j.yaml"}
			str, err := req.DecodeURL()
			Expect(err).ToNot(HaveOccurred())
			Expect(str).To(Equal("https://raw.githubusercontent.com/go-skynet/model-gallery/main/gpt4all-j.yaml"))
		})
		It("parses URLS", func() {
			req := GalleryModel{URL: "https://raw.githubusercontent.com/go-skynet/model-gallery/main/gpt4all-j.yaml"}
			str, err := req.DecodeURL()
			Expect(err).ToNot(HaveOccurred())
			Expect(str).To(Equal("https://raw.githubusercontent.com/go-skynet/model-gallery/main/gpt4all-j.yaml"))
		})
	})
})