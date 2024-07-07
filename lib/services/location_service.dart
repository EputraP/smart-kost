import 'package:flutter/material.dart';
import 'package:location/location.dart';
import 'package:permission_handler/permission_handler.dart' as handler;

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

  Future<bool> checkForPermission() async {
    PermissionStatus status = await _location.hasPermission();

    if (status == PermissionStatus.denied) {
      status = await _location.requestPermission();
      if (status == PermissionStatus.granted) {
        return true;
      }
      return false;
    }
    if (status == PermissionStatus.deniedForever) {
      SnackBar(
        content: const Text(
            "Permission Needed, We use permission to get your location"),
        action: SnackBarAction(
          label: "Setting",
          onPressed: () {
            handler.openAppSettings();
          },
          // onPressed: () {},
        ),
      );
      return false;
    }

    return Future.value(true);
  }

  Future<void> getUserLocation() async {
    if (!(await checkForServiceAvaibility())) {
      return;
    }
    if (!(await checkForPermission())) {
      return;
    }
    final LocationData data = await _location.getLocation();
  }
}
