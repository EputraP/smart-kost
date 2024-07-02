import 'package:flutter/material.dart';
import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:get/get.dart';
import 'package:mobile_app/controller/controller.dart';

class CurvedBottomNavigation extends StatefulWidget {
  const CurvedBottomNavigation({super.key});

  @override
  _CurvedBottomNavigationState createState() => _CurvedBottomNavigationState();
}

class _CurvedBottomNavigationState extends State<CurvedBottomNavigation> {
  Color bgColor = Colors.red;

  final List<Widget> _navigationItem = [
    const Icon(Icons.psychology),
    const Icon(Icons.home),
    const Icon(Icons.settings),
  ];

  final controller = Get.put(NavBarController());
  @override
  Widget build(BuildContext context) {
    return CurvedNavigationBar(
      // key: controller.tabIndex,
      // backgroundColor: bgColor,
      items: _navigationItem,
      index: 1,
      animationDuration: const Duration(milliseconds: 300),
      onTap: controller.changeTabIndex,
      letIndexChange: (index) => true,
    );
  }
}
