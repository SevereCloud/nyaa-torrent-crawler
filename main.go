package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/SevereCloud/nyaa-torrent-crawler/crawler"
	"github.com/SevereCloud/nyaa-torrent-crawler/downloader"
	"github.com/SevereCloud/nyaa-torrent-crawler/subscription"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		showHelp()
	} else {
		switch args[1] {
		case "subscribe":
			subscribe(args[2], args[3])
		case "crawl":
			crawl()
		case "list":
			list()
		case "unsubscribe":
			removeSubscription(args)
		default:
			showHelp()
		}
	}
}

func subscribe(args1 string, args2 string) {
	intargs, _ := strconv.Atoi(args2)
	subscription.InsertNewSubscription(args1, intargs)
}

func showHelp() {
	fmt.Println("usage: ")
	fmt.Println("nyaa-torrent-crawler <option>")
	fmt.Println("\noption list:")
	fmt.Println("subscribe <keyword> <current episode> | subscribe anime based on nyaa.si search keyword")
	fmt.Println("crawl | start crawling based on susbcribe list")
	fmt.Println("list | show subscribe list")
	fmt.Println("removesubscribe | remove subscribe from list")
}

func crawl() {
	listSubscribe := subscription.GetListSubscription()
	for i := range listSubscribe {
		keyword, eps := subscription.GetSubscriptionInfo(i)
		isSuccess, torrentUrl := crawler.StartCrawling(keyword, eps)
		if isSuccess {
			isSuccessDownload := downloader.DownloadTorrent(torrentUrl)
			if isSuccessDownload {
				subscription.UpdateSubscriptionEpisode(i)
			}
		}
	}
}

func list() {
	fmt.Printf("|%-2s|%s|%-2s|\n", "Index", "Keyword", "Current Episode")
	listSubscribe := subscription.GetListSubscription()
	for i := range listSubscribe {
		keyword, eps := subscription.GetSubscriptionInfo(i)
		fmt.Printf("|%-2d|%s|%-2d|\n", i, keyword, eps)
	}
}

func removeSubscription(args []string) {
	var index int
	if len(args) != 3 {
		list()
		fmt.Printf("which index? ")
		fmt.Scanln(&index)
	} else {
		index, _ = strconv.Atoi(args[2])
	}
	subscription.RemoveSubscription(index)
}
