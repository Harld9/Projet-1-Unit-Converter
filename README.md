# Projet 1 : Convertisseur d'unités

Ce projet consiste en la création d'un outil de conversion multi-systèmes développé en Go. Il se divise en deux étapes : une interface en ligne de commande (CLI) et une extension web complète.

---

## Phase 1 : Version Terminal (CLI)

La première phase se concentre sur la logique métier et l'interaction via la console.

### Objectif
Valider les concepts fondamentaux : structures de contrôle, fonctions, gestion des types numériques et validation des entrées utilisateur. L'accent est mis sur la fiabilité et la précision des calculs.

### Fonctionnalités
* **Conversions disponibles** :
    * Kilomètres (km) vers Miles (mi)
    * Kilogrammes (kg) vers Livres (lb)
    * Celsius (°C) vers Fahrenheit (°F)
* **Gestion des erreurs** : Le programme vérifie la validité des saisies (ex: rejette le texte au lieu des nombres) sans crash.
* **Persistance** : Boucle de contrôle permettant d'effectuer plusieurs calculs à la suite.
* **Précision** : Utilisation de flottants haute précision sans arrondi arbitraire.

### Installation et Lancement
1. Accéder au répertoire :
```bash
cd "Phase 1 - Version Go (Terminal)"
```
2. Exécuter le programme :
```bash
go run main.go
```

### Utilisation
1. Sélectionner l'option dans le menu (1, 2 ou 3).
2. Saisir la valeur numérique.
3. Lire le résultat affiché.
4. Choisir l'option 4 pour quitter.

---

## Phase 2 : Extension Web (HTML + Go)

Cette phase porte la logique de conversion vers une interface web interactive en introduisant l'architecture client-serveur.

### Objectif
Développer une interface graphique accessible via navigateur tout en séparant strictement la logique métier de la couche de présentation.

### Fonctionnalités
* **Interface unifiée** : Formulaires dédiés pour chaque conversion sur une page unique.
* **Affichage dynamique** : Les résultats et erreurs sont générés dynamiquement.
* **Architecture MVC** : Organisation du code en modules (`router`, `logic`, `controller`).
* **Serveur HTTP** : Utilisation du package standard `net/http` de Go.

### Installation et Lancement
1. Accéder au répertoire :
```bash
cd "Phase 2 - Extension Web (HTML + Go)"
```
2. Démarrer le serveur :
```bash
go run main.go
```
3. Accéder à l'application :
Ouvrir un navigateur à l'adresse : `http://localhost:8080`

### Structure du Projet Web
* **main.go** : Initialisation du serveur.
* **controller/** : Traitement des requêtes HTTP.
* **logic/** : Fonctions mathématiques réutilisables.
* **router/** : Définition des routes URL.
* **template/** : Fichiers HTML dynamiques.
* **static/** : Feuilles de style CSS et assets.

---

## Annexes : Formules de conversion

Le projet applique les standards de conversion suivants :
* Distance : 1 km = 0.621371 miles
* Masse : 1 kg = 2.20462 livres
* Température : °F = (°C * 1.8) + 32