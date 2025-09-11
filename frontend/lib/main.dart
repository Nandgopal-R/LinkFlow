import 'package:flutter/material.dart';
import 'package:frontend/pages/home_page.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false ,
      theme: ThemeData(
        brightness: Brightness.dark,
        colorScheme: ColorScheme.dark(
          surface: Colors.grey.shade900,
          primary: Colors.grey.shade800,
          secondary: Colors.grey.shade700,
          inversePrimary: Colors.grey.shade300,
        ),
        // focusColor: Colors.transparent,
        // highlightColor: Colors.transparent,
        // hoverColor: Colors.transparent,
        // splashColor: Colors.transparent,
      ),
      home:HomePage(),
    );
  }
}