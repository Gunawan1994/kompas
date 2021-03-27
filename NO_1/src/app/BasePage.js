import React, { Suspense, lazy } from "react";
import { Redirect, Switch, Route } from "react-router-dom";
import { LayoutSplashScreen, ContentRoute } from "../_metronic/layout";
import { DashboardPage } from "./pages/Dashboard";
// import { DaftarKunjungan } from "./pages/DaftarKunjungan";
// import { HistoryKunjungan } from "./pages/HistoryKunjungan";
// import { ArsipKunjungan } from "./pages/ArsipKunjungan";
// import { DaftarOperator } from "./pages/DaftarOperator";
// import { DataMaster } from "./pages/DataMaster";
// import { DaftarTahanan } from "./pages/DaftarTahanan";
// import { Calendar } from "./pages/Calendar";
// import { Log } from "./pages/Log";

// const Laporan = lazy(() =>
//   import("./pages/Laporan/ReportPage")
// );

export default function BasePage() {
  // useEffect(() => {
  //   console.log('Base page');
  // }, []) // [] - is required if you need only one call
  // https://reactjs.org/docs/hooks-reference.html#useeffect

  const dataUser = {
    role: localStorage.getItem("role")
  }

  return (
    <Suspense fallback={<LayoutSplashScreen />}>
      <Switch>
        {/* {dataUser.role === "Operator" ?
          <Redirect exact from="/" to="/satwil" /> : <Redirect exact from="/" to="/dashboard" />} */}
        {dataUser.role === "Super Admin" ?
          <ContentRoute path="/dashboard" component={DashboardPage} /> : <></>}
        {/* {dataUser.role === "Super Admin" ?
          <ContentRoute path="/daftar-kunjungan" component={DaftarKunjungan} /> : <></>}
        {dataUser.role === "Super Admin" ?
          <ContentRoute path="/history-kunjungan" component={HistoryKunjungan} /> : <></>}
        {dataUser.role === "Super Admin" ?
          <ContentRoute path="/arsip-kunjungan" component={ArsipKunjungan} /> : <></>}
        {dataUser.role === "Super Admin" ?
          <ContentRoute path="/daftar-operator" component={DaftarOperator} /> : <></>}
        {dataUser.role === "Super Admin" ?
          <ContentRoute path="/catatan-aktivitas" component={Log} /> : <></>}
        {dataUser.role === "Super Admin" ?
          <ContentRoute path="/data-master" component={DataMaster} /> : <></>}
        {dataUser.role === "Super Admin" ?
          <ContentRoute path="/daftar-tahanan" component={DaftarTahanan} /> : <></>}
        {dataUser.role === "Super Admin" ?
          <Route path="/laporan" component={Laporan} /> : <></>}
        {dataUser.role === "Super Admin" ?
          <ContentRoute path="/calendar" component={Calendar} /> : <></>}
        <Redirect to="error/error-v1" /> */}
      </Switch>
    </Suspense>
  );
}
