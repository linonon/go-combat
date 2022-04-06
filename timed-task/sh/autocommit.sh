#!/usr/local/bin/zsh

function cdgpacp() {
    echo ""
    cd $1
    pwd
	git pull && git add .
	if [[ $2 != "" ]]; then
		git commit -m $2
	else 
		git commit -m "Push without message"
	fi	
	git push
    cd -
}

TIMESTAMP=$(echo "定時提交: $(date)")

WORKSPACE="/Users/linonon/Workspace"

Go_Combat=$WORKSPACE/go-combat
PATH_101=$WORKSPACE/101
Cloud_native_gocamp=$WORKSPACE/Cloud-native-gocamp
GO_Advanced_training_camp=$WORKSPACE/GO-Advanced-training-camp
Learn_Go_with_Github=$WORKSPACE/Learn-Go-with-Github
MyMacConfig=$WORKSPACE/MyMacConfig
leetcode_go_TDD=$WORKSPACE/leetcode-go-TDD

cdgpacp $Go_Combat $TIMESTAMP
cdgpacp $PATH_101 $TIMESTAMP
cdgpacp $Cloud_native_gocamp $TIMESTAMP
cdgpacp $GO_Advanced_training_camp $TIMESTAMP
cdgpacp $Learn_Go_with_Github $TIMESTAMP
cdgpacp $MyMacConfig $TIMESTAMP
cdgpacp $leetcode_go_TDD $TIMESTAMP