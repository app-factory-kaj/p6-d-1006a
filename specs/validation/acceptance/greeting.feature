Feature: Greeting

  @story-1
  Rule: Requesting a greeting for a name returns a greeting addressed to that name

    Scenario: Greeting a named caller
      Given the greeter service is running
      When an API Consumer calls GET /hello with name "Ada"
      Then the response is a greeting addressed to "Ada"

  @story-2
  Rule: Requesting a greeting without a name returns a default greeting instead of an error

    Scenario: Greeting with no name supplied
      Given the greeter service is running
      When an API Consumer calls GET /hello with no name
      Then the response is a default greeting
      And the response is not an error
