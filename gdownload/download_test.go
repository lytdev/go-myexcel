package gdownload

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	onWatch := func(current, total int, percentage float64) {
		fmt.Printf("\r当前已下载大小 %f MB, 下载进度：%.2f%%, 总大小 %f MB",
			float64(current)/1024/1024,
			percentage,
			float64(total)/1024/1024,
		)
	}
	url := "https://playback-tc.videocc.net/polyvlive/76490dba387702307790940685/f0.mp4"
	downloader := NewWithSingle()

	err := downloader.Download(url, "../testdata/example2.mp4", true, onWatch)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestDownloadSingle(t *testing.T) {
	wc := new(WriteCounter)
	wc.onWatch = func(current, total int, percentage float64) {
		fmt.Printf("\r当前已下载大小 %f MB, 下载进度：%.2f%%, 总大小 %f MB",
			float64(current)/1024/1024,
			percentage,
			float64(total)/1024/1024,
		)
	}
	url := "https://playback-tc.videocc.net/polyvlive/76490dba387702307790940685/f0.mp4"
	downloader := NewWithSingle()

	err := downloader.SingleDownload(wc, url, "../testdata/example2.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestDownloadMulti(t *testing.T) {
	url := "https://playback-tc.videocc.net/polyvlive/76490dba387702307790940685/f0.mp4"
	resp, err := http.Head(url)
	if err != nil {
		t.Error(err)
	}
	wc := new(WriteCounter)
	wc.onWatch = func(current, total int, percentage float64) {
		fmt.Printf("\r当前已下载大小 %f MB, 下载进度：%.2f%%, 总大小 %f MB",
			float64(current)/1024/1024,
			percentage,
			float64(total)/1024/1024,
		)
	}
	downloader := NewWithMulti(12)
	err = downloader.MultiDownload(wc, url, "../testdata/example2.mp4", int(resp.ContentLength))
	if err != nil {
		fmt.Println(err)
		return
	}
}
