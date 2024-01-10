
## 使用Git
當我們寫程式時往往分很多階段撰寫，必須要有一個像是書籤的工具，方便我們過一段時間後，仍然可以接續著工作。<br>
另外也會遇到單人同時進行多個專案，或者多人同時修改一個檔案的這類更為複雜的情況。<br>
使用Git可以幫助我們管理程式專案編輯的歷程。<br>



### 下載、安裝Git
```sh
https://git-scm.com/downloads
```

### 參考線上書籍學習Git指令
https://gitbook.tw/

## 使用GitHub作練習Git的工具
假設我們需要將文件放在網路上與特定或不特定的網友共享，並且可以共同編輯、追蹤此文件的版本變化歷程。需要確認以下面步驟進行：<br>
- [ ] 1. 申請一個GitHub帳號
- [ ] 2. 在GitHub建立一個私人的遠端倉
- [ ] 3. 設定個人GitHub帳號的SSH Key
- [ ] 4. 在本地端完成一個Git目錄
- [ ] 5. 將本地端完成的Git目錄，推送到在GitHub上的私人遠端倉


### 1. 申請一個GitHub帳號

### 2. 在GitHub建立一個私人的遠端倉
![在https://github.com/new建立一個private repository](./imgs/2-1_Create%20a%20new%20repository.png)

### 3. 設定個人GitHub帳號的SSH Key
* SSH Key是一對所謂的加密公私鑰，產生了之後，把公鑰放到GitHub上，在我們的電腦端用私鑰把資料加密，再把加密後的資料，傳輸給Github<br>
GitHub再用你在帳號設定的公鑰，將該加密資料解密。

1. 進入個人設定頁面  `https://github.com/settings/keys`
2. 參考`https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent`設定SSH Key
3. 在terminal輸入指令 `ssh-keygen -t ed25519`
4. Enter file in which to save the key：是說要把ssh key存在哪個檔案，在Mac環境，預設是存在`~/.ssh/`路徑，檔案名稱可以自訂，如果要用作業系統預設的檔名，直接按Enter即可。
5. Enter passphrase：直接按Enter即可。
6. Enter same passphrase again:直接按Enter即可。

![3-1_Generate SSH Key](./imgs/3-1_Generate%20SSH%20Key.png)

### 4. 在本地端完成一個Git目錄
1. 建立一個本地端資料夾，例如lesson-go。
```sh
mkdir lesson-go && cd lesson-go
```
2. 初始化一個git專案，並且把編輯好要上傳到GitHub上的資料commit
```sh 
git init
git add .
git commit -m "first commit"
```

### 5. 將本地端完成的Git目錄，推送到在GitHub上的私人遠端倉
* {MYREPO} = 在GitHub建立的私人的遠端倉路徑
```sh
git remote add origin {MYREPO}
git branch -M main
git push -u origin main
``` 

* 成功的話會有類似下方的訊息
```sh
➜  lesson-go git:(main) git push -u origin main  
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 8 threads
Compressing objects: 100% (9/9), done.
Writing objects: 100% (9/9), 328.26 KiB | 1.82 MiB/s, done.
Total 9 (delta 1), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (1/1), done.
To github.com:shodawu/lesson-go.git
 * [new branch]      main -> main
branch 'main' set up to track 'origin/main'.
```

* 重整頁面，確認資料都已經上傳。