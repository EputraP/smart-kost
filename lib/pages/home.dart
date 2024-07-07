import 'package:flutter/material.dart';
import 'package:mobile_app/components/current_location.dart';

class Home extends StatefulWidget {
  const Home({Key? key}) : super(key: key);

  @override
  _HomeState createState() => _HomeState();
}

class _HomeState extends State<Home> {
  dynamic test = MyComponent.create();
  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Text("home"),
      ),
    );
  }
}
