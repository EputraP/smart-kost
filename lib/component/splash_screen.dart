import 'package:animated_splash_screen/animated_splash_screen.dart';
import 'package:flutter/material.dart';
import 'package:lottie/lottie.dart';
import 'package:mobile_app/pages/home.dart';

class SplashScreen extends StatelessWidget {
  const SplashScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AnimatedSplashScreen(
      splash: LottieBuilder.asset("assets/animation/home_splash_screen.json"),
      nextScreen: const Home(),
      splashIconSize: 300,
      backgroundColor: Colors.blue,
    );
  }
}
