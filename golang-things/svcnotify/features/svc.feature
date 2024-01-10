Feature: Notify System
  需由系統管理員(MIS-A)安裝、設定。
  MIS-A要能編輯config.json，將檔案放置在正確的路徑。

  Scenario: 應用程式被Scheduler呼叫時，先讀取並解析config.json
    Given service trigger by schduler
    When config path is "./../config.json"
    Then config content is
    """
    {
        "FTP": {
            "Host": "localhost",
            "Account": "guessme"
        },
        "Email": {
            "Host":"example.mail.server"
        }
    }
    """


