
import React, { useState , useEffect} from "react";
import {useForm,Controller} from "react-hook-form";
import Alert from '@material-ui/lab/Alert';
import { ErrorMessage } from '@hookform/error-message';
import {customAxios} from '../../../services/customAxios';
import qs from 'qs';


export default function ResetPassword(){
    const { control,handleSubmit,errors } = useForm();
    const onSubmit = data => {
        if(data.password !== data.confirm_password){
            setMsg("Password tidak sama");
            setShowMsg(true);
        }else{
            customAxios({
                url:`/password_reset`,
                method:"PUT",
                data:qs.stringify(data),
                headers:{
                  'Content-Type':'application/x-www-form-urlencoded;charset=utf-8'
                }
              }).then(response=>{
                if(response.status === 200){
                    setSuccess(true)
                }
              }).
              catch(error=>{console.log("error submit",error)});
        }
    }

    const [showMsg,setShowMsg] = useState(false);
    const [msg,setMsg] = useState('');
    const [success,setSuccess] = useState(false);
    const [passShow,setPassShow] = useState(false);
    const [passConfirmShow,setPassConfirmShow] = useState(false);

    const showPass = ()=>{
        setPassShow(!passShow);
    }
    const showConfirmPass = ()=>{
        setPassConfirmShow(!passConfirmShow);
    }

    useEffect(()=>{
        if(showMsg){
          setTimeout(()=>{setShowMsg(false)},5000);
        }
        if(success){
          setTimeout(()=>{
            window.location.href = '/auth/login';
          },2000);
        }
      },[showMsg,success]);

    useEffect(()=>{
      const urlParams = new URLSearchParams(window.location.search);
      const token = urlParams.get('token');
      localStorage.setItem("token",token);
      if(token){
      }
    },[]);

    return(
        <>
        {success && <Alert variant="filled" severity="success" >
        Password berhasil diperbaharui
      </Alert>}
        {!success && (
        <div className="login-form login-forgot" style={{ display: "block" }}>
          <div className="text-center mb-10 mb-lg-20">
            <h3 className="font-size-h1">SILAHKAN MASUKKAN PASSWORD BARU ANDA</h3>
          </div>
          <form
            onSubmit={handleSubmit(onSubmit)}
            className="form fv-plugins-bootstrap fv-plugins-framework animated animate__animated animate__backInUp"
          >
            {showMsg && (
              <div className="mb-10 alert alert-custom alert-light-danger alert-dismissible">
                <div className="alert-text font-weight-bold">
                  {msg}
                </div>
              </div>
            )}
            <div className="form-group fv-plugins-icon-container">
                <div class="input-group mb-3">
                        <Controller 
                        as={
                        <input type={passShow?"text":"password"} 
                            className="form-control form-control-solid h-auto py-5 px-6" 
                            placeholder={`Password`}
                            />
                        }
                      name="password"
                      control={control}
                      defaultValue=''
                      rules={{required:true}}
                    />
                    <div class="input-group-append">
                        <span class="input-group-text border-0 cursor-pointer" 
                            onClick={showPass}
                        >{passShow?"Sembunyi":"Lihat"}</span>
                    </div>
                </div>
                <ErrorMessage
                    errors={errors}
                    name="password"
                    message = {<p className="error-message">Wajib Diisi</p>}
                  />
            </div>
            <div className="form-group fv-plugins-icon-container">
                <div class="input-group mb-3">
                        <Controller 
                        as={
                        <input type={passConfirmShow?"text":"password"} 
                            className="form-control form-control-solid h-auto py-5 px-6" 
                            placeholder={`Konfirmasi Password`}
                            />
                        }
                      name="confirm_password"
                      control={control}
                      defaultValue=''
                      rules={{required:true}}
                    />
                    <div class="input-group-append">
                        <span class="input-group-text border-0 cursor-pointer" 
                            onClick={showConfirmPass}
                        >{passConfirmShow?"Sembunyi":"Lihat"}</span>
                    </div>
                </div>
                <ErrorMessage
                    errors={errors}
                    name="confirm_password"
                    message = {<p className="error-message">Wajib Diisi</p>}
                  />
            </div>
            <div className="form-group d-flex flex-wrap flex-center">
              <button
                id="kt_login_forgot_submit"
                type="submit"
                className="btn btn-primary font-weight-bold px-9 py-4 my-3 mx-4"
                onClick={handleSubmit(onSubmit)}
              >
                Ubah Password
              </button>
            </div>
          </form>
        </div>)}
        </>
    )
}