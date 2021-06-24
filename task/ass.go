package main

import (
	"log"
	// "encoding/json"
	// "io/ioutil"

	"github.com/pkg/browser"
	"github.com/leaanthony/clir"
)

type subject struct{
	name string
	link string
}

func main() {

	// Create new cli
	cli := clir.NewCli("Ass", "assistant for school productivity", "v0.0.1")

	// school
	school:=cli.NewSubCommand("school","opens the school section")

	// to set subjects
	var s string
	var link string
	g:= subject{}
	set:=cli.NewSubCommand("set","set subjects")
	set.StringFlag("sub","sets the subject",&s)
	set.StringFlag("link","pass a link to subject",&link)

	set.Action(func() error{
		if s!=""{
			g.name=s
			g.link=link
			log.Println(g)
		}
		return nil
	})


	// select subjects
	var sub string
	school.StringFlag("sub","enter one of the four subjects",&sub)

	school.Action(func() error {
		if sub=="main"{
			browser.OpenURL("https://teams.microsoft.com/_?culture=en-us&country=US&lm=deeplink&lmsrc=homePageWeb&cmpid=WebSignIn#/school/conversations/General?threadId=19:aXxr_nW304fsSPMK8ykXQ5fD-iqHPTOVv3DJsfINgxo1@thread.tacv2&ctx=channel")
		}

		if sub!=""{
			switch sub {
			case "chem":
				// log.Println("chem reached")
				browser.OpenURL("https://teams.microsoft.com/_?culture=en-us&country=US&lm=deeplink&lmsrc=homePageWeb&cmpid=WebSignIn#/school/conversations/General?threadId=19:_j-UGaRjRrwJ_pbRMa3CUzGx-9u6_1g9Uizbvz_eeco1@thread.tacv2&ctx=channel")
			case "bus":
				// log.Println("bus reached")
				browser.OpenURL("https://teams.microsoft.com/_?culture=en-us&country=US&lm=deeplink&lmsrc=homePageWeb&cmpid=WebSignIn#/school/conversations/General?threadId=19:uU-ZGdYhrtHej6rJcp_i1ngsCwybUjRiMWeSBBk5bS41@thread.tacv2&ctx=channel")
			case "compsci":
				// log.Println("compsci reached")
				browser.OpenURL("https://teams.microsoft.com/_?culture=en-us&country=US&lm=deeplink&lmsrc=homePageWeb&cmpid=WebSignIn#/school/conversations/General?threadId=19:8_nRUhI2FHe9IPOwukKWYN6pEBUxsFqqz1EvljqVUz41@thread.tacv2&ctx=channel")
			case "bio":
				// log.Println("bio reached")
				browser.OpenURL("https://teams.microsoft.com/_?culture=en-us&country=US&lm=deeplink&lmsrc=homePageWeb&cmpid=WebSignIn#/school/conversations/General?threadId=19:wFSL2ET5O8KjBqxJZ3cEuahVenVv-Hi9nn7Nw4J5we41@thread.tacv2&ctx=channel")
			}
			
		}

		return nil
	})

	cli.Run()

}