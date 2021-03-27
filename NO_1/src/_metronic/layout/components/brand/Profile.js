/* eslint-disable no-script-url,jsx-a11y/anchor-is-valid,no-undef */
import React, { useState, useEffect } from "react";
import { toAbsoluteUrl } from "../../../_helpers";
import { customAxios } from "../../../../app/services/customAxios";

export function Profile() {
  const userData = {
    role: localStorage.getItem("role"),
    username : localStorage.getItem("username"),
    nirp: localStorage.getItem("unirp")
  }

  const [imgSrc, setImgSrc] = useState('');

  const [loadImage, setLoadImage] = useState(true);

  useEffect(() => {
    if (loadImage) {
      customAxios({
        url: `/users/profil/image`,
        method: "GET",
        responseType: "blob",
      }).then(response => {
        if (response.status === 200) {
          setImgSrc(URL.createObjectURL(response.data))
          setLoadImage(false)
        } else {
          console.log("error get image profile", response);
        }
      }).catch(error => {
        console.log("error get image profile", error);
      });
    }
  }, [loadImage])

  return (
    <div
      className="aside-menu-wrapper flex-column-fluid"
    >
      <div className="mx-10 d-flex align-items-center mt-5">
        <div
          className="symbol symbol-100 mr-5"
        >
          <div className="symbol-label" style={{
            backgroundImage: `url(${toAbsoluteUrl(
              `${imgSrc}`
            )})`
          }} />
          <i className="symbol-badge bg-success" />
        </div>
        <div className="d-flex flex-column">
          <a
            href="#"
            className="text-color-greenlight font-weight-bold font-size-h5 text-hover-primary"
          >
            {userData.username}
          </a>
          <div className="text-muted mt-1">{userData.role}</div>
          <div className="navi mt-2">
            <a href="#" className="navi-item">
              <span className="navi-link p-0 pb-2">
                <span className="navi-icon mr-1">
                  <span className="svg-icon-lg svg-icon-primary text-color-greenlight">
                    NIRP :
                    </span>
                </span>
                <span className="navi-text text-muted text-hover-primary">
                  {userData.nirp}
                </span>
              </span>
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}
