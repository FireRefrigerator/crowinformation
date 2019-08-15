package dao


type DaoAPI interface {
	CreateN(n interface{})
	GetNById(id string) interface{}
	DeleteNById(id string)
	UpdateN(n interface{})
	Connect(om interface{})
}

var daoAPI DaoAPI

func SetDaoAPI(d DaoAPI) {
	daoAPI = d
}

func GetDaoAPI() DaoAPI {
	return daoAPI
}

func Connect(om interface{}) {
	GetCurrentDaoAPI()
	daoAPI.Connect(om)
}

func GetCurrentDaoAPI() {
	if daoAPI == nil {
		daoAPI = &MysqlDao{}
	}
}

func CreateNews(n interface{}) {
	GetCurrentDaoAPI()
	daoAPI.CreateN(n)
}

func UpdateNews(n interface{}) {
	GetCurrentDaoAPI()
	daoAPI.UpdateN(n)
}

func DeleteNewsByid(id string) {
	GetCurrentDaoAPI()
	daoAPI.DeleteNById(id)
}

func getNewsById(id string)  {
	GetCurrentDaoAPI()
	daoAPI.GetNById(id)
}
