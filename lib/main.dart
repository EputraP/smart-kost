import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:mobile_app/controller/location_controller.dart';
import 'package:mobile_app/navigation/app_navigation.dart';
import 'package:mobile_app/services/location_service.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      debugShowCheckedModeBanner: false,
      routerConfig: AppNavigation.router,
    );
  }
}
