#!/usr/local/bin/zsh

function cdgpacp() {
    echo ""
    cd $1
    pwd
	git pull && git add .
	if [[ $2 != "" ]]; then
		git commit -m "定時提交:"$(date '+%Y-%m-%d+%H:%M:%S')
	else 
		git commit -m "Push without message"
	fi	
	git push
	echo ""
}

CURRENT_DIR=$(pwd)

DATE="定時提交:"$(date '+%Y-%m-%d+%H:%M:%S')
TIMESTAMP=$DATE

WORKSPACE="/Users/linonon/Workspace"

go_combat=$WORKSPACE/go-combat
PATH_101=$WORKSPACE/101
cloud_native_gocamp=$WORKSPACE/Cloud-native-gocamp
GO_Advanced_training_camp=$WORKSPACE/GO-Advanced-training-camp
Learn_Go_with_Github=$WORKSPACE/Learn-Go-with-Github
MyMacConfig=$WORKSPACE/MyMacConfig
leetcode_go_TDD=$WORKSPACE/leetcode-go-TDD

cdgpacp $go_combat
cdgpacp $PATH_101
cdgpacp $cloud_native_gocamp
cdgpacp $GO_Advanced_training_camp
cdgpacp $Learn_Go_with_Github
cdgpacp $MyMacConfig
cdgpacp $leetcode_go_TDD

cd $CURRENT_DIR