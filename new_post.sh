#!/bin/bash
if [ -n "$1" ];then 
	dtime=`date '+%Y-%m-%d %H:%M'`
	ddate=`date '+%Y-%m-%d'`
	ndate=`date +%Y-%m-%d`
	file_name="${ndate}_${1}_jlab_tech.md"
	echo $file_name
	touch content/post/$file_name
echo "---
title: ${1}
date: ${ddate}
author: jhfnetboy
draft: false
comments: true
categories: [Creative_Innovation]
---
# ${1}
+ say something here">>content/post/$file_name

else
	echo "Pls input like ./new_post.sh \"your post title\""
fi
