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
"https://github.com/andreasvc/vim-256noir",
"https://github.com/jdsimcoe/abstract.vim",
"https://github.com/danilo-augusto/vim-afterglow",
"https://github.com/AlessandroYorba/Alduin",
"https://github.com/romainl/Apprentice",
"https://github.com/Badacadabra/vim-archery",
"https://github.com/gilgigilgil/anderson.vim",
"https://github.com/zacanger/angr.vim",
"https://github.com/gregsexton/Atom",
"https://github.com/ayu-theme/ayu-vim",
"https://github.com/nightsense/carbonized",
"https://github.com/challenger-deep-theme/vim",
"https://github.com/tyrannicaltoucan/vim-deep-space",
"https://github.com/ajmwagar/vim-deus",
"https://github.com/wadackel/vim-dogrun",
"https://github.com/romainl/flattened",
"https://github.com/chase/focuspoint-vim",
"https://github.com/jaredgorski/fogbell.vim",
"https://github.com/endel/vim-github-colorscheme",
"https://github.com/whatyouhide/vim-gotham",
"https://github.com/morhetz/gruvbox",
"https://github.com/yorickpeterse/happy_hacking.vim",
"https://github.com/NLKNguyen/papercolor-theme",
"https://github.com/keith/parsec.vim",
"https://github.com/scheakur/vim-scheakur",
"https://github.com/w0ng/vim-hybrid",
"https://github.com/kristijanhusak/vim-hybrid-material",
"https://github.com/cocopon/iceberg.vim",
"https://github.com/nanotech/jellybeans.vim",
"https://github.com/wimstefan/Lightning",
"https://github.com/cseelus/vim-colors-lucid",
"https://github.com/jonathanfilip/vim-lucius",
"https://github.com/mkarmona/materialbox",
"https://github.com/christophermca/meta5",
"https://github.com/dikiaap/minimalist",
"https://github.com/tomasr/molokai",
"https://github.com/fmoralesc/molokayo",
"https://github.com/co1ncidence/mountaineer",
"https://github.com/arcticicestudio/nord-vim",
"https://github.com/mhartington/oceanic-next",
"https://github.com/hardcoreplayers/oceanic-material",
"https://github.com/rakr/vim-one",
"https://github.com/joshdick/onedark.vim",
"https://github.com/sonph/onehalf",
"https://github.com/fcpg/vim-orbital",
"https://github.com/owickstrom/vim-colors-paramount",
"https://github.com/sts10/vim-pink-moon",
"https://github.com/kyoz/purify",
"https://github.com/vim-scripts/pyte",
"https://github.com/rakr/vim-colors-rakr",
"https://github.com/vim-scripts/rdark-terminal2.vim",
"https://github.com/junegunn/seoul256.vim",
"https://github.com/AlessandroYorba/Sierra",
"https://github.com/lifepillar/vim-solarized8",
"https://github.com/sainnhe/sonokai",
"https://github.com/liuchengxu/space-vim-dark",
"https://github.com/jaredgorski/SpaceCamp",
"https://github.com/nikolvs/vim-sunbather",
"https://github.com/jacoborus/tender.vim",
"https://github.com/marcopaganini/termschool-vim-theme",
"https://github.com/vim-scripts/twilight256.vim",
"https://github.com/rakr/vim-two-firewatch",
"https://github.com/vim-scripts/wombat256.vim",
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
