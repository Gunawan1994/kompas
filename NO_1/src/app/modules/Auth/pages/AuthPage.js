/* eslint-disable jsx-a11y/anchor-is-valid */
import React from "react";
import {Link, Switch, Redirect} from "react-router-dom";
import {toAbsoluteUrl} from "../../../../_metronic/_helpers";
import {ContentRoute} from "../../../../_metronic/layout"
import Login from "./Login";
// import ResetPassword from "./ResetPassword";
// import Registration from "./Registration";
// import ForgotPassword from "./ForgotPassword";
import "../../../../_metronic/_assets/sass/pages/login/classic/login-1.scss";

export function AuthPage() {
  return (
      <>
        <div className="d-flex flex-column flex-root">
          {/*begin::Login*/}
          <div className="login login-1 login-signin-on d-flex flex-column flex-lg-row flex-row-fluid bg-white">

                <ContentRoute path="/auth/login" component={Login}/>

          </div>
          {/*end::Login*/}
        </div>
      </>
  );
}
