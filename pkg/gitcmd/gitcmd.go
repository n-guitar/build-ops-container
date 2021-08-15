package gitcmd

import (
	"bytes"
	"os"
	"os/exec"
)

func GitCloneCmd(gitrepo string) string {
	git_token := os.Getenv("GOB_GIT_TOKEN")
	if len(git_token) == 0 {
		// fmt.Println("no GOB_GIT_TOKEN environment variable set")
		return "no GOB_GIT_TOKEN environment variable set"
	}
	cmd := exec.Command("git", "clone", "--progress", gitrepo, "gitclone_tmp")

	var stdout bytes.Buffer
	// なんとgit cloneなどの出力は標準エラー(stderr)である
	// プログラムの実行結果の出力ではなく、実行の過程での出力だから」という理由らしい
	cmd.Stderr = &stdout
	err := cmd.Run()

	if err != nil {
		// fmt.Println(err.Error())
		return string(err.Error())
	}
	// fmt.Printf("gitrepo %s\n", gitrepo)
	// fmt.Printf("stdout %s\n", stdout.String())
	return string(stdout.String())
}
