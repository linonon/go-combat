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

101PATH=$WORKSPACE/101
Cloud-native-gocamp=$WORKSPACE"/Cloud-native-gocamp"
GO-Advanced-training-camp=$WORKSPACE"/GO-Advanced-training-camp"
Learn-Go-with-Github=$WORKSPACE"/Learn-Go-with-Github"
MyMacConfig=$WORKSPACE"/MyMacConfig"
leetcode-go-TDD=$WORKSPACE"/leetcode-go-TDD"

cdgpacp 101PATH