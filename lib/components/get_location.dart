import 'package:location/location.dart';
import 'package:flutter/material.dart';

mixin GetLocation {
  static String lat = "";
  static String long = "";
  static Location location = new Location();

  // static Future<void> checkServiceEnabled() async {
  //   var serviceEnabled = await location.serviceEnabled();
  //   if (!serviceEnabled) {
  //     serviceEnabled = await location.requestService();
  //   }
  //   if (!serviceEnabled) {
  //     return;
  //   }
  // }
}
