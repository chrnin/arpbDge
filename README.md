# Extraction JSON de la Foire Aux Questions à destination des entreprises dans le contexte COVID-19

## Objet
Ce dépot contient les Avances Remboursables et Prêts Bonifiés accordés aux entreprises dans le cadre des aides Covid-19.

## Format des données
Le fichier `arpb.csv` est au format CSV, séparé par les `,` et avec les retours de lignes `LF`.

Les données sont remises à jour quotidiennement à partir du dispositif [Démarches Simplifiées](https://www.demarches-simplifiees.fr/) au sein duquel se déroule l'instruction des dossiers.

Structure du fichier:
```
siret: identifiant de l'entreprise
regionCRPInstructeur: région dans laquelle le CRP a instruit le dossier
typeAide: Prêt à taux bonifié ou Avance remboursable
dateDecision: date de la décision au format textuel
dureeAide: durée accordée pour le remboursement de l'aide
montant: montant de l'aide exprimé en euros
```

## License accordée pour les données
Les données contenues dans le fichier `arpb.csv` sont soumises à la Licence Ouverte v2.0
https://www.etalab.gouv.fr/wp-content/uploads/2017/04/ETALAB-Licence-Ouverte-v2.0.pdf

## License accordée pour le programme d'extraction
Copyright © July, 4th 2021, Christophe Ninucci

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

The Software is provided “as is”, without warranty of any kind, express or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose and noninfringement. In no event shall the authors or copyright holders X be liable for any claim, damages or other liability, whether in an action of contract, tort or otherwise, arising from, out of or in connection with the software or the use or other dealings in the Software.
Except as contained in this notice, the name of Christophe Ninucci shall not be used in advertising or otherwise to promote the sale, use or other dealings in this Software without prior written authorization from Christophe Ninucci.