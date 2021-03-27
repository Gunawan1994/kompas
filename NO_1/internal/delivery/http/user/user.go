package user

// Handler class
// type Handler struct {
// 	userSvc IUserService
// 	rstSvc  IResetService
// }

// // IUserService is the abstraction of user service
// type IUserService interface {
// 	GetNirp(nirp uint64) (user.UserModel, error)
// 	AddUser(ctx context.Context, nirp uint64, username string, email string, password string, role string, idSatwil uint64) (user.UserModel, error)
// 	EditUser(ctx context.Context, nirp uint64, username string, email string, role string, idSatwil uint64) (user.UserModel, error)
// 	DelUser(ctx context.Context, nirp uint64) (user.UserModel, error)
// 	GetAllData(search string, page int, limit int) ([]user.UserModel, error)
// 	GetAllCount(search string) (int, error)
// 	Operator(ctx context.Context) (user.Operator, error)
// }

// // IResetService Interface
// type IResetService interface {
// 	EditPass(email string, password string) (reset.ResetModel, error)
// }

// // New will create object for class Handler
// func New(u IUserService, r IResetService) Handler {
// 	return Handler{
// 		userSvc: u,
// 		rstSvc:  r,
// 	}
// }

// //Get user data handler
// func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
// 	resp := &response.Response{}
// 	defer resp.RenderJSON(w, r)

// 	vars := mux.Vars(r)
// 	var id string
// 	id = vars["nirp"]

// 	u64, err := strconv.ParseUint(id, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	conv := uint64(u64)

// 	data, err := h.userSvc.GetNirp(conv)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
// 		return
// 	}

// 	resp.Data = data
// 	resp.Error.Msg = SucGet
// 	resp.Error.Code = http.StatusOK
// 	return
// }

// // Post user data handler
// func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
// 	resp := &response.Response{}
// 	defer resp.RenderJSON(w, r)

// 	var (
// 		nirp, username, email, password, role, idSatwil string
// 	)
// 	nirp = r.PostFormValue("nirp")
// 	username = r.PostFormValue("username")
// 	email = r.PostFormValue("email")
// 	password = r.PostFormValue("password")
// 	role = r.PostFormValue("role")
// 	idSatwil = r.PostFormValue("id_satwil")

// 	if idSatwil == "" {
// 		idSatwil = "0"
// 	}

// 	hashPass, err := encrypt.HashPassword(password)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrAdd), http.StatusBadRequest)
// 		return
// 	}

// 	u64nirp, err := strconv.ParseUint(nirp, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrAdd), http.StatusBadRequest)
// 		return
// 	}
// 	u64satwil, err := strconv.ParseUint(idSatwil, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrAdd), http.StatusBadRequest)
// 		return
// 	}

// 	data, err := h.userSvc.AddUser(r.Context(), uint64(u64nirp), username, email, hashPass, role, uint64(u64satwil))
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrAdd), http.StatusBadRequest)
// 		return
// 	}

// 	resp.Data = data
// 	resp.Error.Msg = SucAdd
// 	resp.Error.Code = http.StatusOK
// 	return
// }

// // Put user data handler
// func (h Handler) Put(w http.ResponseWriter, r *http.Request) {
// 	resp := &response.Response{}
// 	defer resp.RenderJSON(w, r)

// 	vars := mux.Vars(r)
// 	var (
// 		nirp, username, email, password, role, idSatwil string
// 	)
// 	nirp = vars["nirp"]
// 	username = r.PostFormValue("username")
// 	email = r.PostFormValue("email")
// 	password = r.PostFormValue("password")
// 	role = r.PostFormValue("role")
// 	idSatwil = r.PostFormValue("id_satwil")

// 	if idSatwil == "" {
// 		idSatwil = "0"
// 	}

// 	hashPass, err := encrypt.HashPassword(password)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrEdit), http.StatusBadRequest)
// 		return
// 	}

// 	u64nirp, err := strconv.ParseUint(nirp, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrEdit), http.StatusBadRequest)
// 		return
// 	}
// 	u64satwil, err := strconv.ParseUint(idSatwil, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrEdit), http.StatusBadRequest)
// 		return
// 	}

// 	data, err := h.userSvc.EditUser(r.Context(), uint64(u64nirp), username, email, role, uint64(u64satwil))
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrEdit), http.StatusBadRequest)
// 		return
// 	}

// 	if password != "" {
// 		_, err := h.rstSvc.EditPass(email, hashPass)
// 		if err != nil {
// 			log.Println(err)
// 			resp.SetError(errors.New(ErrEdit), http.StatusBadRequest)
// 			return
// 		}
// 	}

// 	resp.Data = data
// 	resp.Error.Msg = SucEdit
// 	resp.Error.Code = http.StatusOK
// 	return
// }

// // Del user data handler
// func (h Handler) Del(w http.ResponseWriter, r *http.Request) {
// 	resp := &response.Response{}
// 	defer resp.RenderJSON(w, r)

// 	vars := mux.Vars(r)
// 	var nirp string
// 	nirp = vars["nirp"]

// 	u64, err := strconv.ParseUint(nirp, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	data, err := h.userSvc.DelUser(r.Context(), uint64(u64))
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrDel), http.StatusBadRequest)
// 		return
// 	}
// 	resp.Data = data
// 	resp.Error.Msg = SucDel
// 	resp.Error.Code = http.StatusOK
// 	return
// }

// // GetAll user data handler
// func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
// 	resp := &response.Response{}
// 	defer resp.RenderJSON(w, r)
// 	var (
// 		search, limit, page string
// 	)
// 	search = r.FormValue("search")
// 	limit = r.FormValue("limit")
// 	page = r.FormValue("page")
// 	if limit == "" {
// 		limit = "10"
// 	}
// 	if page == "" {
// 		page = "1"
// 	}
// 	limitInt, err := strconv.Atoi(limit)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	pageInt, err := strconv.Atoi(page)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	data, err := h.userSvc.GetAllData(search, pageInt, limitInt)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
// 		return
// 	}
// 	count, err := h.userSvc.GetAllCount(search)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
// 		return
// 	}

// 	resp.Data = data
// 	resp.Pagination = Pagination{
// 		NextPage:     pageInt + 1,
// 		TotalData:    count,
// 		PreviousPage: pageInt - 1,
// 	}
// 	resp.Error.Msg = SucGet
// 	resp.Error.Code = http.StatusOK

// 	return
// }

// // Upload user data handler
// func (h Handler) Upload(w http.ResponseWriter, r *http.Request) {
// 	resp := &response.Response{}
// 	defer resp.RenderJSON(w, r)

// 	// nirp := r.Context().Value("nirp")
// 	nirp := r.FormValue("nirp")

// 	// Maximum upload of 10 MB files
// 	r.ParseMultipartForm(10 << 20)

// 	// Get handler for filename, size and headers
// 	file, handler, err := r.FormFile("imageFile")
// 	if err != nil {
// 		log.Println(errors.Wrap(err))
// 		return
// 	}

// 	defer file.Close()

// 	fileExtension := filepath.Ext(handler.Filename)
// 	if fileExtension != ".jpg" {
// 		resp.SetError(errors.New(ErrExt), http.StatusBadRequest)
// 		return
// 	}

// 	dir := "/opt/visol/profile"

// 	err = os.MkdirAll(dir, os.ModePerm)
// 	if err != nil {
// 		log.Println(errors.Wrap(err))
// 		resp.SetError(err, http.StatusInternalServerError)
// 		return
// 	}

// 	// Create a new file in the uploads directory
// 	dst, err := os.Create(fmt.Sprintf("%s/%s.jpg", dir, nirp))
// 	if err != nil {
// 		log.Println(errors.Wrap(err))
// 		resp.SetError(err, http.StatusInternalServerError)
// 		return
// 	}

// 	defer dst.Close()

// 	_, err = io.Copy(dst, file)
// 	if err != nil {
// 		log.Println(errors.Wrap(err))
// 		resp.SetError(err, http.StatusInternalServerError)
// 		return
// 	}

// 	resp.Error.Msg = SucUpload
// 	resp.Error.Code = http.StatusOK
// 	return
// }

// // GetImage user data handler
// func (h Handler) GetImage(w http.ResponseWriter, r *http.Request) {
// 	nirp := r.Context().Value("nirp")
// 	str := fmt.Sprintf("%v", nirp)
// 	w.Header().Set("Content-Type", "image/jpeg")
// 	var url = "/opt/visol/profile/" + str + ".jpg"
// 	_, err := os.Open(url)
// 	if err != nil {
// 		log.Println(errors.Wrap(err), "fallback to default image")
// 		url = "/opt/visol/profile/default.jpg"
// 	}
// 	http.ServeFile(w, r, url)
// 	return
// }

// //UserProfil user data handler
// func (h Handler) UserProfil(w http.ResponseWriter, r *http.Request) {
// 	x := r.Context().Value("nirp")
// 	str := fmt.Sprintf("%v", x)
// 	resp := &response.Response{}
// 	defer resp.RenderJSON(w, r)
// 	u64, err := strconv.ParseUint(str, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	conv := uint64(u64)

// 	data, err := h.userSvc.GetNirp(conv)
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
// 		return
// 	}
// 	// Example direct return the data
// 	resp.Data = Profil{
// 		Id:       data.Nirp,
// 		Role:     data.Role,
// 		Username: data.Username,
// 	}
// 	resp.Error.Msg = SucGet
// 	resp.Error.Code = http.StatusOK
// 	return
// }

// //Operator handler
// func (h Handler) Operator(w http.ResponseWriter, r *http.Request) {
// 	resp := &response.Response{}
// 	defer resp.RenderJSON(w, r)

// 	data, err := h.userSvc.Operator(r.Context())
// 	if err != nil {
// 		log.Println(err)
// 		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
// 		return
// 	}

// 	resp.Data = data
// 	resp.Error.Msg = SucAdd
// 	resp.Error.Code = http.StatusOK
// 	return
// }
