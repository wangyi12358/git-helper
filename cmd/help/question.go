package help

import "github.com/wangyi12358/go-tools/array"

type Command struct {
	Title string
	Short string
}

var Commands = []*Command{
	{
		Title: "修改最近一次提交的信息",
		Short: "git commit --amend",
	},
	{
		Title: "撤销指定的提交，会生成一个新的 commit",
		Short: "git revert <commit>",
	},
	{
		Title: "退回到指定 commit，不会生成新的 commit",
		Short: "git reset --hard <commit_id>",
	},
	{
		Title: "设置|修改 远程分支地址",
		Short: "git remote add|set-url origin <url>",
	},
	{
		Title: "查看远程分支",
		Short: "git remote -v",
	},
	{
		Title: "还原文件到暂存区的状态 | 还原所有文件到暂存区的状态",
		Short: "git restore --staged <file> | git restore .",
	},
}

var Titles = array.Map(Commands, func(command *Command) string {
	return command.Title
})
