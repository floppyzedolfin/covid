#########################
# 1/ Le Test
#########################

- Ce test technique a vocation à comprendre comment vous abordez un problème et comment vous faites les choix techniques pour y répondre.
- Lors de vos explications, n'hésitez donc pas à citer des noms de technos ou à décrire une architecture, mais prenez soin d'expliquer vos choix.
- Il n'est absolument pas requis de faire l'ensemble du test : il faut y passer environ 1h30 à 2h. Une attention particulière sera portée aux réponses de la partie 4 si vous avez pris le temps de vous y attarder.

Il s'agit de développer une petite API exploitant les données du gouvernement concernant les tests Covid.
Le dataset est le suivant : https://www.data.gouv.fr/fr/datasets/r/406c6a23-e283-4300-9484-54e78c8ae675
La description des headers se trouve en annexe de ce fichier.

#########################
# 2/ Consignes
#########################

Une liberté complète est laissée pour la technique, hormis les contraintes suivantes :
- l'API doit être codée en Go
- les échanges avec l'API doivent être réalisés au format JSON
- l'API doit être capable de récupérer les données depuis le site data.gouv.fr

Des explications (2/3 phrases) sont attendues :
- sur les choix de conception de l'API
- sur la gestion et la structuration des données

#########################
# 3/ Besoins
#########################

Catégorie "ADMIN" :
- récupérer les données depuis le site data.gouv.fr

Catégorie "API" :
- récupérer la liste des départements pour lesquels nous disposons de données
- récupérer les données dont nous disposons pour 1 département et 1 date
- récupérer les données dont nous disposons pour X département et un range de dates

Catégorie "ANALYTICS" :
- récupérer au niveau national, pour un range de dates, un historique des métriques suivantes : nombre de tests, nombre de tests positifs et ratio de tests positifs.
- récupérer un résumé pour 1 département. Le résumé comprend : le jour avec le plus de tests, le jour avec le plus de gens positifs et le jour avec le taux le plus élevé (parmis les jours eligibles, c'est à dire ceux dont le nombre de tests > 10).

#########################
# 4/ En bonus
#########################

Nouveaux besoins en Catégorie "API" :
- récupérer la date la plus vieille et la date la plus récente pour lesquelles nous disposons de données
- récupérer les catégories d'âge pour lesquelles nous disposons de données

Nouveaux besoins en Catégorie "ANALYTICS" :
- récupérer pour un jour donné : le top 5 des départements avec les ratio de tests positifs les plus élevés, par catégorie d'age (pour les départements et catégories d'âge éligibles, c'est à dire ceux dont le nombre de tests > 10).

Quelques questions :
- comment automatiser la mise à jour régulière des données ?
- quelle features devrions-nous ajouter à l'application pour une intégration front-end ?
- quelles seraient les modifications à apporter si nous souhaitons industrialiser cette API, par exemple sous la forme d'une application mobile ?
- avec-vous une idée sur une optimisation, une techno de BDD qui serait adaptée ?
- avez-vous des remarques ou des points que vous souhaitez aborder ?



============

API
- curl localhost:8405/departements => renvoie la liste des départements avec de la donnée
- curl localhost:8405/departement/${DEPARTEMENT_ID}/${DATE} => renvoie les données pour ce département à la date voulue
- curl localhost:8405/departements/dates/${MIN_DATE}/${MAX_DATE} => renvoie les données des départements entre les dates voulues
