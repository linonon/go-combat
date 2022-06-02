#!/usr/local/bin/zsh

function cdgpacp() {
    cd $1
    echo "cd to `pwd`"
	git pull && git add .
	git commit -m "定時提交: $(date '+%Y-%m-%d %H:%M:%S')"
	git push
	echo "Sync done in $1"
	echo ""
}

CURRENT_DIR=$(pwd)

DATE="定時提交:"$(date '+%Y-%m-%d+%H:%M:%S')
TIMESTAMP=$DATE

WS="/Users/linonon/Workspace"

go_combat=$WS/go-combat
PATH_101=$WS/101
cloud_native_gocamp=$WS/Cloud-native-gocamp
GO_Advanced_training_camp=$WS/GO-Advanced-training-camp
Learn_Go_with_Github=$WS/Learn-Go-with-Github
MyMacConfig=$WS/MyMacConfig
leetcode_go_TDD=$WS/leetcode-go-TDD

cdgpacp $go_combat
cdgpacp $PATH_101
cdgpacp $cloud_native_gocamp
cdgpacp $GO_Advanced_training_camp
cdgpacp $Learn_Go_with_Github
cdgpacp $MyMacConfig
cdgpacp $leetcode_go_TDD

cd $CURRENT_DIR
