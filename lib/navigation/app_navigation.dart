import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:mobile_app/main_wrapper.dart';
import 'package:mobile_app/pages/home.dart';
import 'package:mobile_app/pages/params.dart';
import 'package:mobile_app/pages/setting.dart';

class AppNavigation {
  AppNavigation._();

  static final _rootNavigatorKey = GlobalKey<NavigatorState>();
  static final _shellNavigatorHome =
      GlobalKey<NavigatorState>(debugLabel: 'shellHome');
  static final _shellNavigatorParams =
      GlobalKey<NavigatorState>(debugLabel: 'shellParams');
  static final _shellNavigatorSetting =
      GlobalKey<NavigatorState>(debugLabel: 'shellSetting');

  static final GoRouter router = GoRouter(
    initialLocation: "/home",
    debugLogDiagnostics: true,
    navigatorKey: _rootNavigatorKey,
    routes: [
      StatefulShellRoute.indexedStack(
        builder: (context, state, navigationShell) {
          return MainWrapper(navigationShell: navigationShell);
        },
        branches: <StatefulShellBranch>[
          StatefulShellBranch(
            navigatorKey: _shellNavigatorParams,
            routes: <RouteBase>[
              GoRoute(
                path: "/params",
                name: "Params",
                builder: (BuildContext context, GoRouterState state) =>
                    const Params(),
              )
            ],
          ),
          StatefulShellBranch(
            navigatorKey: _shellNavigatorHome,
            routes: <RouteBase>[
              GoRoute(
                path: "/home",
                name: "Home",
                builder: (BuildContext context, GoRouterState state) =>
                    const Home(),
              )
            ],
          ),
          StatefulShellBranch(
            navigatorKey: _shellNavigatorSetting,
            routes: <RouteBase>[
              GoRoute(
                path: "/setting",
                name: "Setting",
                builder: (BuildContext context, GoRouterState state) =>
                    const Setting(),
              )
            ],
          ),
        ],
      ),
      // GoRoute(
      //   parentNavigatorKey: _rootNavigatorKey,
      //   path: "/splash_screen",
      //   name: "Splash Screen",
      //   builder: (context, state) => const SplashScreen(),
      // ),
    ],
  );
}
