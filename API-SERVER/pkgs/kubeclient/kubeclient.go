package kubeclient


import (
	"API-SERVER/pkgs/types"
	"net/http"
	"io/ioutil"
	 "encoding/json"
	"github.com/astaxie/beego"
	"fmt"
	"bytes"
)

type Kubeops struct{


}


const(

	namespaceapiversion = "v1"
	buildnamespace = "chandraspace"

)



func CheckNamespace(namecheck *types.Namespace,kubernetesendpoint string)(value bool,err error){
	req,err := http.NewRequest("GET",kubernetesendpoint+"/api/v1/namespaces/"+buildnamespace,nil)
	if err!=nil{
		beego.Info("Http request creation Failed",err)
		return true,err
	}

	res, err:= http.DefaultClient.Do(req)
	if err!=nil{
		beego.Info("Http request Failed",err)
		return true,err
	}

	Namespaceres := new(types.Namespace)
	body,_:= ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body,&Namespaceres)
	if err!=nil{
		beego.Info("Json unmarshal failed",err)
		return true,err
	}
	if res.StatusCode == 404 {
                beego.Info("Namespace already present")
		return false,nil
	}else{
		return true,nil
	}
}


func (k *Kubeops)CreateNamespace(kubernetes_endpoint string)(err error){

	 BuilderNamespace := new(types.Namespace)
	 BuilderNamespace.TypeMeta.APIVersion = namespaceapiversion
	 BuilderNamespace.TypeMeta.Kind = "Namespace"
	 BuilderNamespace.Metadata.Name =  buildnamespace
	 Namespacestatus,err := CheckNamespace(BuilderNamespace,kubernetes_endpoint)
   	   if err != nil {
		fmt.Println(err)
		return err

	}
	if  Namespacestatus == false{
		body, err := json.Marshal(BuilderNamespace)
		if err != nil {
			beego.Info("Unable to Marshal",err)
			return err
		}
		body_io := bytes.NewReader(body)
		fmt.Println("POST", kubernetes_endpoint + "/api/v1/namespaces/", body_io)
		req, err := http.NewRequest("POST", kubernetes_endpoint + "/api/v1/namespaces/", body_io)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return err

		}
		fmt.Println(res)


	}
       return nil
}













































