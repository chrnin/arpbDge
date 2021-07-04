package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/machinebox/graphql"
)

type Annotations []struct {
	Label       string `json:"label"`
	StringValue string `json:"stringValue"`
}

func (a Annotations) get(key string) string {
	for _, v := range a {
		if v.Label == key {
			return v.StringValue
		}
	}
	return "n/c"
}

type Demandeur struct {
	Siret string `json:"siret"`
}

type GroupeInstructeur struct {
	Label string `json:"label"`
}

type Node struct {
	Annotations       Annotations       `json:"annotations"`
	Demandeur         Demandeur         `json:"demandeur"`
	GroupeInstructeur GroupeInstructeur `json:"groupeInstructeur"`
}

type graphqlResponse struct {
	Demarche struct {
		Id       string `json:"id"`
		Dossiers struct {
			PageInfo struct {
				EndCursor   string `json:"endCursor"`
				HasNextPage bool   `json:"hasNextPage"`
			} `json:"pageInfo"`
			Nodes []Node `json:"nodes"`
		} `json:"dossiers"`
	} `json:"demarche"`
}

type ARPB [][]string

func (a ARPB) csv() {
	os.Stdout.WriteString("siret,regionCRPInstructeur,typeAide,dateDecision,dureeAide,montant\n")
	writer := csv.NewWriter(os.Stdout)
	writer.WriteAll(a)
}
func main() {
	getARPB().csv()
}

func getARPB() ARPB {
	var arpb ARPB
	response := request("")
	for {
		for _, node := range response.Demarche.Dossiers.Nodes {
			// On ignore silencieusement les lignes qui ne comportent pas un montant lisible
			if _, err := strconv.Atoi(node.Annotations.get("Montant du prêt")); err != nil {
				continue
			}
			a := []string{
				node.Demandeur.Siret,
				node.GroupeInstructeur.Label,
				node.Annotations.get("Quelle forme prend l'aide ?"),
				node.Annotations.get("Date de la décision"),
				node.Annotations.get("Durée du prêt"),
				node.Annotations.get("Montant du prêt"),
			}
			arpb = append(arpb, a)
		}

		if response.Demarche.Dossiers.PageInfo.HasNextPage {
			response = request(response.Demarche.Dossiers.PageInfo.EndCursor)
		} else {
			break
		}
	}
	return arpb
}

func request(cursor string) graphqlResponse {
	dsKey := os.Getenv("DS_KEY")
	graphqlClient := graphql.NewClient("https://www.demarches-simplifiees.fr/api/v2/graphql")
	graphqlRequest := graphql.NewRequest(query(cursor))
	graphqlRequest.Header.Add("Authorization", "Bearer "+dsKey)
	graphqlRequest.Header.Add("Content-Type", "application/json")
	var response graphqlResponse

	if err := graphqlClient.Run(context.Background(), graphqlRequest, &response); err != nil {
		fmt.Println(graphqlRequest)
		panic(err)
	}
	return response
}

func query(cursor string) string {
	var cursorString string
	if cursor != "" {
		cursorString = fmt.Sprintf(", after:\"%s\"", cursor)
	}
	return fmt.Sprintf(`{
    demarche(number: 30928) {
			dossiers(state:accepte, first:100%s) { 
				pageInfo {
					endCursor
					hasNextPage
				}
				nodes {
						annotations {
								label
								stringValue
						}
						demandeur { 
								... on PersonneMorale { 
										siret
								}
						}
						groupeInstructeur {
								label
	}}}}}`, cursorString)
}
