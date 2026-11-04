# LAB 6: F# Programs

This lab contains three standalone F# scripts:

1. `marks_division.fsx` reads four subject marks, calculates their average percentage, and classifies it with pattern matching.
2. `prime_check.fsx` uses a recursive divisor check, demonstrates five test values, and prompts for an integer.
3. `string_processing.fsx` demonstrates string reversal and palindrome checking, alphabetic/numeric/mixed transformations, and basic email-format validation. Each task includes three test inputs.

## Run

Install the .NET SDK, then run each script from this directory:

```powershell
dotnet fsi .\marks_division.fsx
dotnet fsi .\prime_check.fsx
dotnet fsi .\string_processing.fsx
```

The marks entered for each subject should be whole numbers from 0 to 100. The string transformation duplicates numeric strings by concatenating them with themselves. Palindrome checks are exact (case- and whitespace-sensitive), and the email regex checks a basic address format rather than every address permitted by the email standards.
