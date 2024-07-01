import 'package:flutter/material.dart';

class Params extends StatefulWidget {
  const Params({Key? key}) : super(key: key);

  @override
  _ParamsState createState() => _ParamsState();
}

class _ParamsState extends State<Params> {
  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Text("Params"),
      ),
    );
  }
}
