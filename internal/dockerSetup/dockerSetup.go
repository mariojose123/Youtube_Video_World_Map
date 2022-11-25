package Dockersetup

import "os/exec"

type Docker struct {
}

func initDocker() {
	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
}

/*sudo docker run  -d  -e POSTGRES_PASSWORD=123 -e POSTGRES_USER=user -e POSTGRES_DB=dbname -p 5432:5432 postgres
sudo docker cp  worldcities.csv e7e18eab08f5:/data/worldcities.csv*/
