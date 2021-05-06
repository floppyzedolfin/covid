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

API:
- curl localhost:8405/api/departements => renvoie la liste des départements avec de la donnée
- curl localhost:8405/api/departements/${DEPARTEMENT_ID}/dates/${DATE} => renvoie les données pour ce département à la date voulue
- curl localhost:8405/api/departements/${DEPARTEMENT_IDS}/dates/${MIN_DATE}/${MAX_DATE} => renvoie les données des départements entre les dates voulues -- "${DEPARTEMENT_IDS}" est une liste séparée par des virgules (`01,02,75`)

ADMIN:
- curl localhost:8405/admin/load => charge le document depuis le site du gouvernement

ANALYTICS:
- curl localhost:8405/analytics/national/dates/${MIN_DATE}/${MAX_DATE} => renvoie des informations niveau national
- curl localhost:8405/analytics/departements/${DEPARTEMENT_ID}/brief => renvoie un résumé d'un département
- curl localhost:8405/analytics/dates/${DATE}/top5 => renvoie les 5 départements ayant le plus de tests postitifs par catégorie d'age


Choix architecturaux:
- Passer par fiber pour avoir du code rapide à écrire (ma premiere piste était d'utiliser des protobuff, pour un gos coût et pas un grand gain)
- Exposer des endpoints avec des paths "intuitifs"


- Comment automatiser la mise à jour des données : on peut utiliser un `timer.Tick()` qui ira régulièrement mettre à jour le contenu de la base du server en le retéléchargeant depuis la source.
- Pour une intégration front-end, ce code écrit en 2h demande pas mal de gestion d'erreur actuellement absente, de couverture de tests. Par ailleurs, télécharger un fichier de 10mo tous les jours n'est pas forcément le meilleur moyen, il doit y avoir une interface exposée quelque part.
- pour des BDDs, on peut utiliser godog (de cucumber)
- J'ai commencé ce test en voulant utiliser gRPC, mais ça s'est avéré trop coûteux en temps (30 min). Ce test est sous pression, et plus long que la plupart des autres tests que j'ai réalisés. Il est néanmoins intéressant car réalisable dans le temps imparti. Cependant, la majorité des endpoints se ressemble sans grande différence, on se retrouve à écrire souvent la même chose.
L'endpoint des différents ages est le plus intéressant une fois qu'on en a fait 2 précédents

