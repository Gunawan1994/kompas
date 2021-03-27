import React, { useMemo } from "react";
import objectPath from "object-path";
import { useHtmlClassService } from "../../_core/MetronicLayout";
import { SearchDropdown } from "../extras/dropdowns/search/SearchDropdown";
import { UserNotificationsDropdown } from "../extras/dropdowns/UserNotificationsDropdown";
import { QuickUserToggler } from "../extras/QuiclUserToggler";
import { useHistory } from "react-router-dom";

export function Topbar() {
  const uiService = useHtmlClassService();
  const history = useHistory();

  const logoutClick = () => {
      const toggle = document.getElementById("kt_quick_user_toggle");
      if (toggle) {
        toggle.click();
      }
      history.push("/logout");
  };

  const layoutProps = useMemo(() => {
    return {
      viewSearchDisplay: objectPath.get(
        uiService.config,
        "extras.search.display"
      ),
      viewNotificationsDisplay: objectPath.get(
        uiService.config,
        "extras.notifications.display"
      ),
      viewUserDisplay: objectPath.get(uiService.config, "extras.user.display"),
    };
  }, [uiService]);

  return (
    <div className="topbar">
      {/* {layoutProps.viewSearchDisplay && <SearchDropdown />} */}

      {/* {layoutProps.viewNotificationsDisplay && <UserNotificationsDropdown />} */}

      {/* {layoutProps.viewUserDisplay && <QuickUserToggler />} */}
       <div className="topbar-item">
            <div className="btn btn-icon w-auto btn-clean d-flex align-items-center btn-lg px-2"
                 id="kt_quick_user_toggle">
              <>

                <button className="btn btn-light-primary btn-bold" onClick={logoutClick}>Keluar</button>
              </>
            </div>
          </div>

    </div>
  );
}
