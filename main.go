package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"sort"
	"strings"
	"sync"
)

var repoURLs = []string{
	"https://github.com/gpakosz/.tmux",
	"https://github.com/tony/tmux-config",
	"https://github.com/brandur/tmux-extra",
	"https://github.com/sriramkandukuri/automux",
	"https://github.com/zdcthomas/dmux",
	"https://github.com/tmux-python/libtmux",
	"https://github.com/powerline/powerline",
	"https://github.com/ivaaaan/smug",
	"https://github.com/ryandotsmith/tat",
	"https://github.com/remi/teamocil",
	"https://github.com/evnp/tmex",
	"https://github.com/zinic/tmux-cssh",
	"https://github.com/MunifTanjim/tmux-suspend",
	"https://github.com/jamesottaway/tmux-up",
	"https://github.com/nkh/tmuxake",
	"https://github.com/jimeh/tmuxifier",
	"https://github.com/tmuxinator/tmuxinator",
	"https://github.com/oxidane/tmuxomatic",
	"https://github.com/tmux-python/tmuxp",
	"https://github.com/goerz/tmuxpair",
	"https://github.com/christoomey/vim-tmux-navigator",
	"https://github.com/jatap/tmux-base16-statusline", 
	"https://github.com/chriskempson/base16-shell",
	"https://github.com/seebi/tmux-colors-solarized",
	"https://github.com/wfxr/tmux-power",
	"https://github.com/jimeh/tmux-themepack",
	"https://github.com/edouard-lopez/tmux-tomorrow/",
	"https://github.com/chriskempson/tomorrow-theme", 
	"https://github.com/dracula/tmux", 
	"https://github.com/arcticicestudio/nord-tmux",
	"https://github.com/egel/tmux-gruvbox",
	"https://github.com/o0th/tmux-nova",
	"https://github.com/darko-mesaros/aws-tmux",
	"https://github.com/arl/gitmux",
	"https://github.com/tmux-plugins/tmux-battery",
	"https://github.com/vascomfnunes/tmux-clima",
	"https://github.com/Determinant/tmux-colortag",
	"https://github.com/jdxcode/tmux-cpu-info",
	"https://github.com/tmux-plugins/tmux-cpu",
	"https://github.com/tassaron/tmux-df",
	"https://github.com/vascomfnunes/tmux-kripto",
	"https://github.com/tmux-plugins/tmux-maildir-counter",
	"https://github.com/thewtex/tmux-mem-cpu-load",
	"https://github.com/MunifTanjim/tmux-mode-indicator",
	"https://github.com/Feqzz/tmux-mpv-info",
	"https://github.com/jaclu/tmux-mullvad",
	"https://github.com/xamut/tmux-network-bandwidth",
	"https://github.com/maxrodrigo/tmux-nordvpn",
	"https://github.com/tmux-plugins/tmux-online-status",
	"https://github.com/Brutuski/tmux-piavpn",
	"https://github.com/olimorris/tmux-pomodoro-plus",
	"https://github.com/tmux-plugins/tmux-prefix-highlight",
	"https://github.com/jdxcode/tmux-spotify-info", 
	"https://github.com/Feqzz/tmux-spotify-info",
	"https://github.com/jdxcode/tmux-weather",
	"https://github.com/xamut/tmux-weather",
	"https://github.com/ofirgall/tmux-window-name",
	"https://github.com/Feqzz/tmux-weather-info-yr",
	"https://github.com/alexanderjeurissen/tmux-world-clock",
	"https://github.com/tweekmonster/tmux2html",
	"https://github.com/NHDaly/tmux-better-mouse-mode",
	"https://github.com/laktak/extrakto",
	"https://github.com/bjesus/muxile",
	"https://github.com/b0o/tmux-autoreload",
	"https://github.com/Alkindi42/tmux-bitwarden",
	"https://github.com/ofirgall/tmux-browser",
	"https://github.com/CrispyConductor/tmux-copy-toolkit",
	"https://github.com/IngoMeyer441/tmux-easy-motion",
	"https://github.com/Morantron/tmux-fingers",
	"https://github.com/roosta/tmux-fuzzback",
	"https://github.com/wfxr/tmux-fzf-url",
	"https://github.com/jaclu/tmux-menus",
	"https://github.com/whame/tmux-modal",
	"https://github.com/jaclu/tmux-mouse-swipe",
	"https://github.com/ChanderG/tmux-notify",
	"https://github.com/fcsonline/tmux-thumbs",
	"https://github.com/yardnsm/tmux-1password",
	"https://github.com/schasse/tmux-jump",
	"https://github.com/jaclu/tmux-power-zoom",
	"https://github.com/kristijanhusak/tmux-simple-git-status",
	"https://github.com/xamut/tmux-spotify",
	"https://github.com/artemave/tmux_super_fingers",
	"https://github.com/jabirali/tmux-tilish", 
	"https://github.com/gcla/tmux-wormhole",
	"https://github.com/tmux-plugins",
	"https://github.com/tmux-plugins/tmux-continuum",
	"https://github.com/tmux-plugins/tmux-copycat",
	"https://github.com/tmux-plugins/tmux-fpp",
	"https://github.com/tmux-plugins/tmux-logging",
	"https://github.com/tmux-plugins/tmux-open",
	"https://github.com/tmux-plugins/tmux-pain-control",
	"https://github.com/tmux-plugins/tmux-resurrect",
	"https://github.com/tmux-plugins/tmux-sessionist",
	"https://github.com/tmux-plugins/tmux-sidebar",
	"https://github.com/tmux-plugins/tpm",
	"https://github.com/tmux-plugins/tmux-urlview",
	"https://github.com/tmux-plugins/tmux-yank",
	"https://github.com/tmux-plugins/tmux-example-plugin",
	"https://github.com/tmux-plugins/tmux-test",
	"https://github.com/csdvrx/sixel-tmux",
	"https://github.com/huntie/sublime-tmux",
	"https://github.com/tmux-plugins/vim-tmux",
	"https://github.com/mapio/tmux-tail-f",
}

func main() {
	urlSet := map[string]struct{}{}
	for _, url := range repoURLs {
		urlSet[url] = struct{}{}
	}

	ch := make(chan string, len(urlSet))
	for url := range urlSet {
		ch <- url
	}
	close(ch)

	number := 50
	var wg sync.WaitGroup
	wg.Add(number)
	var mu sync.Mutex
	var repositories []*github.Repository
	for i := 0; i < number; i++ {
		go func() {
			defer wg.Done()
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: os.Args[1]},
			)
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)

			for url := range ch {
				segments := strings.Split(url, "/")
				if len(segments) != 5 {
					continue
				}

				owner := segments[3]
				repo := segments[4]

				repository, _, err := client.Repositories.Get(context.Background(), owner, repo)
				if err != nil {
					fmt.Println(err)
					continue
				}

				mu.Lock()
				repositories = append(repositories, repository)
				// fmt.Printf("repo: %v, star count: %v\n", *repository.HTMLURL, *repository.StargazersCount)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].GetStargazersCount() < repositories[j].GetStargazersCount()
	})

	fmt.Printf("------------------------------------\n\n\n\n\n")
	for _, repo := range repositories {
		fmt.Printf("repo: [%v](%v), star count: %v \n", *repo.HTMLURL, *repo.HTMLURL, *repo.StargazersCount)
		fmt.Println("")
	}
}
