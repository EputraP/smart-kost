import 'package:flutter/material.dart';
import 'package:location/location.dart';

class LocationService {
  LocationService.init();
  static LocationService instance = LocationService.init();

  Location _location = Location();
  Future<bool> checkForServiceAvaibility() async {
    bool isEnabled = await _location.serviceEnabled();
    if (isEnabled) {
      return Future.value(true);
    }
    isEnabled = await _location.requestService();

    if (isEnabled) {
      return Future.value(true);
    }
    return Future.error("Location service is not enabled");
  }

  Future<bool> getLocation() async {
    PermissionStatus status = await _location.hasPermission();

    if (status == PermissionStatus.denied) {
      status = await _location.requestPermission();
      if (status == PermissionStatus.granted) {
        return true;
      }
      return false;
    }
    if (status == PermissionStatus.deniedForever) {
      return false;
    }

    return false;
  }
}
