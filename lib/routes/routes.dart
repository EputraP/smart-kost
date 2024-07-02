import 'package:get/get.dart';
import 'package:mobile_app/pages/params.dart';
import 'package:mobile_app/pages/setting.dart';
import '../pages/home.dart';

class AppPage {
  static List<GetPage> routes = [
    GetPage(
      name: home,
      page: () => const Home(),
    ),
    GetPage(
      name: params,
      page: () => const Params(),
    ),
    GetPage(
      name: setting,
      page: () => const Setting(),
    ),
  ];

  static gethome() => home;
  static getparams() => params;
  static getsetting() => setting;

  static String home = '/home';
  static String params = '/params';
  static String setting = '/setting';
}
