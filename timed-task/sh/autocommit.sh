#!/usr/local/bin/zsh

function cdgpacp() {
    cd $1
	git pull && git add .
	if [[ $2 != "" ]]; then
		git commit -m $2
	else 
		git commit -m "Push without message"
	fi	
	git push
    cd -
}

WORKSPACE="/Users/linonon/Workspace"

PATH_101=$WORKSPACE/101
Cloud_native_gocamp=$WORKSPACE/Cloud-native-gocamp
GO_Advanced_training_camp=$WORKSPACE/GO-Advanced-training-camp
Learn_Go_with_Github=$WORKSPACE/Learn-Go-with-Github
MyMacConfig=$WORKSPACE/MyMacConfig
leetcode_go_TDD=$WORKSPACE/leetcode-go-TDD

cdgpacp 101PATH