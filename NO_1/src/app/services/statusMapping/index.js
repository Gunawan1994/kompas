const statusLabel = (code)=>{
    switch (code) {
        case 0:
            return "Terima";
        case 1:
            return "Batal";
        case 2:
            return "Selesai";
        case 4:
            return "Online";
        default:
            return "";
    }
}
export {statusLabel};