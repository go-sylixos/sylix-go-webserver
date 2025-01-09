package main

const (
	StatusDeploying = iota
	StatusDeployed
	StatusError
)

var listData = []App{
	{
		Name:       "app1",
		Status:     StatusDeploying,
		Cluster:    "cluster1",
		DeployTime: "2020-01-11 10:03:05",
		UpdateBy:   "admin",
	},
	{
		Name:       "app2",
		Status:     StatusDeploying,
		Cluster:    "cluster2",
		DeployTime: "2020-01-11 10:53:55",
		UpdateBy:   "admin",
	},
	{
		Name:       "app3",
		Status:     StatusDeployed,
		Cluster:    "cluster3",
		DeployTime: "2020-01-11 08:30:50",
		UpdateBy:   "admin",
	},
	{
		Name:       "app4",
		Status:     StatusDeployed,
		Cluster:    "cluster4",
		DeployTime: "2020-01-11 11:00:03",
		UpdateBy:   "admin",
	},
	{
		Name:       "app5",
		Status:     StatusDeployed,
		Cluster:    "cluster5",
		DeployTime: "2020-01-11 17:59:33",
		UpdateBy:   "admin",
	},
	{
		Name:       "app6",
		Status:     StatusError,
		Cluster:    "cluster6",
		DeployTime: "2020-01-11 17:55:58",
		UpdateBy:   "admin",
	},
	{
		Name:       "app7",
		Status:     StatusDeploying,
		Cluster:    "cluster7",
		DeployTime: "2020-01-11 13:01:01",
		UpdateBy:   "admin",
	},
	{
		Name:       "app8",
		Status:     StatusError,
		Cluster:    "cluster8",
		DeployTime: "2020-01-11 08:09:51",
		UpdateBy:   "admin",
	},
	{
		Name:       "app9",
		Status:     StatusDeployed,
		Cluster:    "cluster9",
		DeployTime: "2020-01-11 15:15:15",
		UpdateBy:   "admin",
	},
	{
		Name:       "app10",
		Status:     StatusDeployed,
		Cluster:    "cluster10",
		DeployTime: "2020-01-11 13:13:13",
		UpdateBy:   "admin",
	},
	{
		Name:       "app11",
		Status:     StatusDeploying,
		Cluster:    "cluster11",
		DeployTime: "2020-01-11 11:00:01",
		UpdateBy:   "admin",
	},
	{
		Name:       "app12",
		Status:     StatusError,
		Cluster:    "cluster12",
		DeployTime: "2020-01-11 13:01:01",
		UpdateBy:   "admin",
	},
	{
		Name:       "app13",
		Status:     StatusDeploying,
		Cluster:    "cluster13",
		DeployTime: "2020-01-11 11:00:01",
		UpdateBy:   "admin",
	},
	{
		Name:       "app14",
		Status:     StatusDeployed,
		Cluster:    "cluster14",
		DeployTime: "2020-01-11 09:09:59",
		UpdateBy:   "admin",
	},
	{
		Name:       "app15",
		Status:     StatusDeployed,
		Cluster:    "cluster15",
		DeployTime: "2020-01-11 10:00:00",
		UpdateBy:   "admin",
	},
}
