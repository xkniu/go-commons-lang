# Go Commons Lang

Common utilities to extend go standard library.

Build a high-quality library, follow these coding principals:

- Provide only really common used utils instead of business related utils.
- Provide simple and delicate functions instead of complicated functions.
- It's suggested to combine multiple simple functions for difficult scenes for caller.
- Comments and test cases.

Attention please. Before you use this library, you should know that:

- This library will use panic instead of error. It indicates a function is not always safe, you should wrap it by yourself if calling quietly is important to you.
