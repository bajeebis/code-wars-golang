/* Problem description
Is the string uppercase?
Task
Create a method IsUpperCase to see whether the string is ALL CAPS. For example:

type MyString string
MyString("c").IsUpperCase() == false
MyString("C").IsUpperCase() == true
MyString("hello I AM DONALD").IsUpperCase() == false
MyString("HELLO I AM DONALD").IsUpperCase() == true
MyString("ACSKLDFJSgSKLDFJSKLDFJ").IsUpperCase() == false
MyString("ACSKLDFJSGSKLDFJSKLDFJ").IsUpperCase() == true
In this Kata, a string is said to be in ALL CAPS whenever it does not contain any lowercase letter so any string containing no letters at all is trivially considered to be in ALL CAPS.
*/

package kata
//package main

import (
  //"fmt"
  "unicode"
)

type MyString string

func (s MyString) IsUpperCase() bool {
  // Your code here!
  char := []rune(s)
  for  i := 0; i < len(char); i++{
    if unicode.IsUpper(char[i]) == true || char[i] == ' ' {
      //fmt.Printf("Character %c is at position %d is valid\n", char, pos)
      continue
    } else {
      //fmt.Printf("Not an uppercase letter or whitespace\n")
      return false
    }
  }
  return true
}
