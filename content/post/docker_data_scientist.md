jupyter/datascience-notebook
Source on GitHub | Dockerfile commit history | Docker Hub image tags

jupyter/datascience-notebook includes libraries for data analysis from the Julia, Python, and R communities.

Everything in the jupyter/scipy-notebook and jupyter/r-notebook images, and their ancestor images
The Julia compiler and base environment
IJulia to support Julia code in Jupyter notebooks
HDF5, Gadfly, and RDatasets packages
docker run --rm -p 10000:8888 -e JUPYTER_ENABLE_LAB=yes -v "$PWD":/Users/mac/Library/Mobile Documents/com~apple~CloudDocs/log jupyter/datascience-notebook:latest
docker run --rm -p 10000:8888 -e JUPYTER_ENABLE_LAB=yes -v "/Users/mac/Library/Mobile Documents/com~apple~CloudDocs/log" jupyter/datascience-notebook:latest

