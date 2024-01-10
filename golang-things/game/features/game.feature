Feature: 猜數字
  猜數字遊戲， 可以設定1..100的數字，跟最多猜幾次。
  猜對了變回傳「恭喜過關」並且結束遊戲
  如果猜錯就回傳「過大」、「過小」來提示玩家
  如果猜到最後一次都猜錯，就回傳「失敗」，並且結束遊戲。
  已結束的遊戲，總是回傳「遊戲已結束」

  Scenario: 剩餘次數大於零時猜錯
    Given game a target is 80
    And has 3 more time to guess
    When player guess 85
    Then game return message "過大"

Scenario Outline: 猜數字
    Given game a target is <target>
    And has <chances> more time to guess
    When player guess <guess>
    Then game return message <returns>

  Examples:
    | target | chances | guess |returns |
    |    80 |   3 |    85 |    "過大" |
    |    80 |   1 |   85 |    "失敗" |
    |    80 |   0 |   80 |    "失敗" |
