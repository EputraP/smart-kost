import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';

typedef CallbackFunction = void Function(dynamic data);

class CurvedBottomNavigator extends StatefulWidget {
  const CurvedBottomNavigator(
      {super.key, required this.bgColor, required this.callbackFunction});
  final Color bgColor;
  final CallbackFunction callbackFunction;

  @override
  _CurvedBottomNavigatorState createState() => _CurvedBottomNavigatorState();
}

class _CurvedBottomNavigatorState extends State<CurvedBottomNavigator> {
  // final int _page = 1;
  // Color bgColor = Colors.red;
  final GlobalKey<CurvedNavigationBarState> _bottomNavigationKey = GlobalKey();

  final List<Widget> _navigationItem = [
    const Icon(Icons.psychology),
    const Icon(Icons.home),
    const Icon(Icons.settings),
  ];

  @override
  Widget build(BuildContext context) {
    return CurvedNavigationBar(
      key: _bottomNavigationKey,
      backgroundColor: widget.bgColor,
      items: _navigationItem,
      index: 1,
      animationDuration: const Duration(milliseconds: 300),
      onTap: (index) {
        // if (index == 0) {
        //   bgColor = Colors.blue;
        // } else if (index == 1) {
        //   bgColor = Colors.yellow;
        // } else if (index == 2) {
        //   bgColor = Colors.red;
        // } else if (index == 3) {
        //   bgColor = Colors.green;
        // } else if (index == 4) {
        //   bgColor = Colors.teal;
        // }
        widget.callbackFunction(index);
      },
      letIndexChange: (index) => true,
    );
  }
}
