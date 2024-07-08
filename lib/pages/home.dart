import 'package:flutter/material.dart';
import 'package:mobile_app/components/current_location.dart';
import 'package:mobile_app/controller/location_controller.dart';
import 'package:get/get.dart';
import 'package:mobile_app/services/location_service.dart';

class Home extends StatefulWidget {
  const Home({Key? key}) : super(key: key);

  @override
  _HomeState createState() => _HomeState();
}

class _HomeState extends State<Home> {
  final LocationController locationController =
      Get.put<LocationController>(LocationController());

  @override
  void initState() {
    LocationService.instance.getUserLocation(controller: locationController);
    super.initState();
  }

  dynamic test = MyComponent.create();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Obx(() {
          return locationController.isAccessingLocation.value
              ? const Column(
                  children: [
                    CircularProgressIndicator(),
                    Text("Accessing Location")
                  ],
                )
              : locationController.errorDescription.value.isNotEmpty ||
                      locationController.userLocation.value == null
                  ? Column(
                      children: [
                        Text(locationController.errorDescription.value),
                      ],
                    )
                  : Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Text(
                          "Latitude: ${locationController.userLocation.value?.latitude} ",
                          style: Theme.of(context).textTheme.headlineSmall,
                        ),
                        Text(
                          "Longitude :  ${locationController.userLocation.value?.longitude} ",
                          style: Theme.of(context).textTheme.headlineSmall,
                        )
                      ],
                    );
        }),
      ),
    );
  }
}
