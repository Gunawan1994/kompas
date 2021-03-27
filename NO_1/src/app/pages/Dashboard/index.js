import React,{useState,useEffect} from 'react';
import qs from 'qs';
import {
  Card,
  CardBody,
  CardHeader,
  CardHeaderToolbar
} from "../../../_metronic/_partials/controls";
import {IconButton,makeStyles,InputBase,InputLabel,Button} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import BootstrapTable from 'react-bootstrap-table-next';
import paginationFactory from 'react-bootstrap-table2-paginator';
import {Modal} from 'react-bootstrap';
import {useForm,Controller} from "react-hook-form";
import {customAxios} from "../../services/customAxios";
import Alert from '@material-ui/lab/Alert';

const useStyles = makeStyles((theme)=>({
      input:{
        marginLeft: theme.spacing(1),
        flex: 1,
      },
      iconButton: {
        padding: 10,
      },
}))

export function DashboardPage(){
  const {control,handleSubmit} = useForm();
  
  const onSubmit = data => {
      customAxios({
        url:'/artikel',
        method:"POST",
        data:qs.stringify(data),
        headers:{
          'Content-Type':'application/x-www-form-urlencoded;charset=utf-8'
        }
      }).then(response=>{
        if(response.status === 200){
          setLoadData(true);
          setShow(false);
        }
      }).
      catch(error=>{console.log("error submit",error)});
  };

  const [show, setShow] = useState(false);
  const handleClose = () => {setShow(false);};

  const [page, setPage] = useState(1);
  const [searchData, setSearchData] = useState('');
  const [searchText, setSearchText] = useState('');
  const [loadData,setLoadData] = useState(true);
  const [dataTable, setDataTable] = useState([]);
  const [showMsg,setShowMsg] = useState(false);
  const [msg,setMsg] = useState('');
  const [msgType,setMsgtype] = useState('success');
  

  useEffect(()=>{
    if(loadData){
      customAxios({
        url:`/artikel`,
        method:"GET",
      }).then(response=>{
        if(response.status === 200 && response.data.error.status === false){
          setDataTable(response.data.data);
          setLoadData(false);
        }else{
          console.log("error get list",response);
          setMsgtype('error');
          setMsg(`${response.data.error.msg} (code : ${response.data.error.code.toString()})`);
          setShowMsg(true);
          setLoadData(false);
        }
      }).catch(error=>{
        console.log("error get list",error);
          setMsgtype('error');
          setMsg(error.message);
          setShowMsg(true);
          setLoadData(false);
      }
     );
    }
  },[page,searchData,loadData])
  
  const columns = [{
    dataField: 'judul',
    text: 'Judul'
  }, {
    dataField: 'body',
    text: 'Body'
  }, {
    dataField: 'author',
    text: 'author'
  }];

  const classes = useStyles();
  const search = ()=>{
    return (<div>
    <h3 className="py-5">List Artikel</h3>
    <InputBase
        className={classes.input}
        placeholder="Cari"
        inputProps={{ 'aria-label': 'Cari' }}
        onChange={(e)=>{setSearchText(e.target.value)}}
        onKeyDown={(e)=>{
          if(e.key == 'Enter'){
            setSearchData(e.target.value);
            setLoadData(true);
          }
        }}
        value={searchText}
      />
      <IconButton type="submit" onClick={()=>{
        setSearchData(searchText);
        setLoadData(true);
        }}
        className={classes.iconButton} aria-label="search">
        <SearchIcon />
      </IconButton>
  </div>)
  }

  const modal = ()=>{
    return(
      <Modal show={show} onHide={handleClose} size="lg">
        <Modal.Header closeButton className="bg-dark">
          <Modal.Title className="text-white">Tambah artikel</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <form onSubmit={handleSubmit(onSubmit)}>
              <div className="form-group row">
                  <div className="col-lg-6">
                  </div>
                  <div className="col-lg-6">
                  </div>
              </div>
              <div className="form-group row">
                  <div className="col-lg-6">
                  <InputLabel>Judul</InputLabel>
                   <Controller
                    as={<input type="text" className="form-control border-thick" id="judul" />}
                    name="judul"
                    control={control}
                  />
                  </div>
                  </div>
                  <div className="form-group row">
                  <div className="col-lg-6">
                  <InputLabel>Body</InputLabel>
                   <Controller
                    as={<input type="text" className="form-control border-thick" id="body" />}
                    name="body"
                    control={control}
                  />
                  </div>
                  </div>
                  <div className="form-group row">
                  <div className="col-lg-6">
                  <InputLabel>Author</InputLabel>
                   <Controller
                    as={<input type="text" className="form-control border-thick" id="author" />}
                    name="author"
                    control={control}
                  />
                  </div>
              </div>
          </form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Batal
          </Button>
          <Button variant="primary" color="primary" onClick={handleSubmit(onSubmit)} className="btn btn-success">
            Simpan
          </Button>
        </Modal.Footer>
      </Modal>
    )
  }

  return(<>
      <Card>
        {showMsg?<Alert variant="filled" severity={msgType}>
          {msg}
        </Alert>:''}
        <CardHeader title={search()}>
          <CardHeaderToolbar>
            <Button onClick={()=>{
          setShow(true);
                }} variant="contained" color="primary" className="mr-10 text-white">Buat Baru</Button>
          </CardHeaderToolbar>
        </CardHeader>
        <CardBody>
          <BootstrapTable
                  wrapperClasses="table-responsive"
                  bordered={false}
                  classes="table table-head-custom table-vertical-center overflow-hidden"
                  bootstrap4
                  remote
                  keyField="id"
                  data={ dataTable } columns={ columns } pagination={ paginationFactory() } />
        </CardBody>
        {modal()}
    </Card>

  </>)
}