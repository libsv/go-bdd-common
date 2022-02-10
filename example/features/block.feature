Feature: Bitcoin Blocks
  Should be able to read block information

  Scenario: Block 0 should be retrievable
    When I make a GET request to "https://api.whatsonchain.com/v1/bsv/main/block/height/0"
    Then the HTTP response code should be 200
    And the data should match JSON "block-0"