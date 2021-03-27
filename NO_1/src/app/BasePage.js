import React, { Suspense, lazy } from "react";
import { Redirect, Switch, Route } from "react-router-dom";
import { LayoutSplashScreen, ContentRoute } from "../_metronic/layout";
import { DashboardPage } from "./pages/Dashboard";

export default function BasePage() {

  const dataUser = {
    role: localStorage.getItem("role")
  }

  return (
    <Suspense fallback={<LayoutSplashScreen />}>
      <Switch>
        {dataUser.role === "Super Admin" ?
          <ContentRoute path="/dashboard" component={DashboardPage} /> : <></>}
      </Switch>
    </Suspense>
  );
}
