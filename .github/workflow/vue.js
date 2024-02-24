name: Deploy Vue.js

on:
  push:
    branches:
      - frontend # Trigger the workflow on push events to the frontend branch.

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v2
      with:
        persist-credentials: false

    - name: Set up Node.js
      uses: actions/setup-node@v1
      with:
        node-version: '21' # Set this to the node version you are using

    - name: Install dependencies
      run: npm install
      working-directory: ./webapp # Change this to the directory of your Vue.js project

    - name: Build
      run: npm run build && echo $PWD && ls -la
      working-directory: ./webapp # Adjust if your build script is located elsewhere


    - name: Upload artifact
      uses: actions/upload-artifact@v2
      with:
          name: webapp
          path: ./webapp/dist
          
  deploy:
    runs-on: ubuntu-latest
    needs: build
    steps:
      - name: Download artifact
        uses: actions/download-artifact@v2
        with:
           name: webapp
           path: ./webapp
           
      - name: Deploy to GitHub Pages
        uses: JamesIves/github-pages-deploy-action@releases/v3
        with:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          BRANCH: gh-pages # The branch the action should deploy to.
          FOLDER: ./webapp # The folder the action should deploy.
          CLEAN: true # Automatically remove deleted files from the deployment