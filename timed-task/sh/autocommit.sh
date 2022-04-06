#!/usr/local/bin/zsh

function cdgpacp() {
    echo ""
    cd $1
    pwd
	CURRENT_DIR=$(pwd)
	git pull && git add .
	if [[ $2 != "" ]]; then
		git commit -m $2
	else 
		git commit -m "Push without message"
	fi	
	git push
    cd CURRENT_DIR
	echo ""
}

DATE="定時提交:"$(date +%Y%m%d)
TIMESTAMP=$DATE

WORKSPACE="/Users/linonon/Workspace"

go_combat=$WORKSPACE/go-combat
PATH_101=$WORKSPACE/101
cloud_native_gocamp=$WORKSPACE/Cloud-native-gocamp
GO_Advanced_training_camp=$WORKSPACE/GO-Advanced-training-camp
Learn_Go_with_Github=$WORKSPACE/Learn-Go-with-Github
MyMacConfig=$WORKSPACE/MyMacConfig
leetcode_go_TDD=$WORKSPACE/leetcode-go-TDD

cdgpacp $go_combat $TIMESTAMP
cdgpacp $PATH_101 $TIMESTAMP
cdgpacp $cloud_native_gocamp $TIMESTAMP
cdgpacp $GO_Advanced_training_camp $TIMESTAMP
cdgpacp $Learn_Go_with_Github $TIMESTAMP
cdgpacp $MyMacConfig $TIMESTAMP
cdgpacp $leetcode_go_TDD $TIMESTAMP