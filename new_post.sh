#!/bin/bash
if [ -n "$1" ];then 
	dtime=`date '+%Y-%m-%d %H:%M'`
	ndate=`date +%Y-%m-%d`
	num=$(date +%s)
	file_name="${ndate}_jlab_blog${num}.md"
	echo $file_name
	touch content/post/$file_name
echo "---
title: ${1}
date: ${dtime}
author: jhfnetboy
draft: true
comments: true
categories: [Creative_Innovation]
---">>content/post/$file_name

else
	echo "Pls input like ./new_post.sh \"your post title\""
fi
