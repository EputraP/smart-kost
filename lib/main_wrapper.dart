import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:mobile_app/component/curved_bottom_navigator.dart';

class MainWrapper extends StatefulWidget {
  const MainWrapper({super.key, required this.navigationShell});

  final StatefulNavigationShell navigationShell;

  @override
  _MainWrapperState createState() => _MainWrapperState();
}

class _MainWrapperState extends State<MainWrapper> {
  int _page = 1;
  int trek = 1;
  Color bgColor = Colors.red;

  GlobalKey<CurvedNavigationBarState> _bottomNavigationKey = GlobalKey();

  void _goBranch(int index) {
    widget.navigationShell.goBranch(
      index,
      initialLocation: index == widget.navigationShell.currentIndex,
    );
  }

  void callbackFunction(value) {
    setState(() {
      _page = value["index"];
      bgColor = value["bgColor"];
      _goBranch(value["index"]);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SizedBox(
        width: double.infinity,
        height: double.infinity,
        child: widget.navigationShell,
      ),
      bottomNavigationBar: CurvedBottomNavigator(
        bgColor: bgColor,
        callbackFunction: callbackFunction,
      ),
    );
  }
}
