#!/bin/bash
echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"

msg="rebuilding jlab site `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi

# Push Hugo content 
git add -A
git commit -m "$msg"
git push origin master


# Build the project. 
hugo   --baseUrl="https://www.jlab.tech/" 
# if using a theme, replace by `hugo --theme  <yourtheme>`

# Go To Public folder
cd public
# Add changes to git.
git add -A

# Commit changes.

git commit -m "$msg"

# Push source and build repos.
git push origin master

# Come Back
cd ..
