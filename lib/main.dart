import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';
import 'package:mobile_app/services/current_location.dart';
import 'package:mobile_app/component/curved_bottom_navigator.dart';
import 'package:mobile_app/component/splash_screen.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      debugShowCheckedModeBanner: false,
      home: SplashScreen(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _page = 1;
  final Color bgColor = Colors.red;
  GlobalKey<CurvedNavigationBarState> _bottomNavigationKey = GlobalKey();
  // CurrentLocation currentLocationInstance =
  //     CurrentLocation(age: 10, name: 'Bob');
  // String test = currentLocationInstance.get("te");
  // log('data: $testVal');
  void callbackFunction(value) {
    setState(() {
      _page = value;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        color: bgColor,
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Text(_page.toString(), style: TextStyle(fontSize: 160)),
              // Text(testVal),
            ],
          ),
        ),
      ),
      bottomNavigationBar: CurvedBottomNavigator(
        bgColor: bgColor,
        callbackFunction: callbackFunction,
      ),
    );
  }
}
