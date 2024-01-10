## 現況與需求
1. 產品經理A不定時更新所有待辦事項到File Server的一個Excel檔
2. 各成員沒有到File Server查看檔案的習慣
3. 產品經理A希望能有一個工具程式可以每週提醒各成員還未完成的事項。


## 辨認組織與流程
1. 公司最高層為總經理(P)，下轄業務、財務、產品、維運五個部門。
2. 產品部共有4個產品，Prod-1、Prod-2、Prod-3、Prod-4。
3. 產品部有兩位產品經理，PM-A、PM-B，三位開發人員，Dev-A、Dev-B、Dev-C。
4. 維運部門除了有一位系統管理員(MIS-A)，另有兩位開發人員(IT-A、IT-B)負責公司內部系統開發，但也會支援產品開發。
5. 每一家公司的組織架構、企業文化皆不相同。


## 業務流程內的角色、職責
1. 產品經理: 不定時編輯待辦事項。 
2. 開發組員: 每週收到一次個人的未結案事項通知。
3. 其他潛在角色: File Server管理者(MIS-A)、開發組員的組織主管(P、PM-A、PM-B)等
4. 其他部門例如財務部、業務部等，是否有相同或相似的流程或需求？


## 非功能需求
1. File Server可以是FTP Server、Google drive、NAS等
2. 待辦事項可以為Excel、CSV或其他線上編輯文件
3. 通知方式可以是Email、Telegram、APP Push Notify


## 定義資料字典
1. 待辦事項
2. File Server
3. 定期通知
4. 角色： 系統管理者、流程管理者、接收者


## 決定架構
1. File Server: FTP
2. 待辦事項: Excel
3. 通知方式: Email
4. 系統運行在公司內部Server，需提供安裝檔由MIS-A安裝、設定、管理。
5. 需有介面讓MIS-A設定Mail Server、FTP Server等存取權限。


## 辨認需要的技術
1. FTP: 
* https://tools.ietf.org/html/rfc959 
* github.com/jlaffaye/ftp 

2. Excel: 
* github.com/tealeg/xlsx/v3
* https://en.wikipedia.org/wiki/Office_Open_XML

3. Email:
* gopkg.in/gomail.v2
* SMTP: https://tools.ietf.org/html/rfc5321
* Address Spec: https://tools.ietf.org/html/rfc5322


## 開發方法
1. MVP: Minimum Viable Product。
2. TDD: Test-Driven Development

