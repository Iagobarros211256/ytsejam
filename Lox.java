package com.craftinginterpreters.lox;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

//lox class definition
// here basically we are verifing if is some code to read
// if its not, we run a prompt to get the code from the user
public class Lox {
  static boolean hadError = false;
  public static void main(String[] args) throws IOException {
    if (args.length > 1) {
      System.out.println("Usage: jlox [script]");
      System.exit(64); 
    } else if (args.length == 1) {
      runFile(args[0]);
    } else {
      runPrompt();
    }
  }
//if we have a file with code to read this will be executed
  private static void runFile(String path) throws IOException {
    byte[] bytes = Files.readAllBytes(Paths.get(path));
    run(new String(bytes, Charset.defaultCharset()));
       // Indicate an error in the exit code.
    if (hadError) System.exit(65);
  }
// if whe dont have any file to read we will run an prompt
// to get the code from the user
  private static void runPrompt() throws IOException {
    InputStreamReader input = new InputStreamReader(System.in);
    BufferedReader reader = new BufferedReader(input);

    for (;;) { 
      System.out.print("> ");
      String line = reader.readLine();
      if (line == null) break;
      run(line);
      hadError = false;
    }
  }

  // this run method runs your code
  // transforms it in a string scans it
  // and print its tokens

  private static void run(String source) {
    Scanner scanner = new Scanner(source);
    List<Token> tokens = scanner.scanTokens();

    // For now, just print the tokens.
    for (Token token : tokens) {
      System.out.println(token);
    }
  }
// a general error handling method,if we encounter an error 
// save it
  static void error(int line, String message) {
    report(line, "", message);
  }
// and print it to the screen
  private static void report(int line, String where,String message){
    System.err.println(
        "[line " + line + "] Error" + where + ": " + message);
    hadError = true;
  }
}