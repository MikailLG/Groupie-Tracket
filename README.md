# ⚓ One Piece Tracker - Groupie Tracker

## Présentation du Projet
Ce projet est une application web développée en **Go** permettant de naviguer dans l'univers de One Piece. L'objectif est d'fficher un maximum d'informations des pirates, leurs primes et leurs équipages.

**Fonctionnalités clés :**
* **Recherche dynamique :** Trouver un pirate par son nom.
* **Filtres avancés :** Filtrer par statut (Vivant/Mort), équipage, ou présence d'une prime.
* **Pagination :** Navigation optimisée par lots de 10 personnages.
* **Détails des ressources :** Fiche complète pour chaque pirate (Fruit du démon, rôle, âge, etc.).

---

## Installation et Lancement

1. **Prérequis :** Avoir [Go](https://go.dev/dl/) installé sur votre machine.
2. **Clonage / Téléchargement :** Placez-vous dans le dossier `Groupie-Tracker`.
3. **Initialisation :**
   ```bash
   go mod tidy

Accès : Ouvrez votre navigateur sur http://localhost:8080

**Détail des routes :**
/ : Page d'accueil présentant l'univers et le thème du site.
/collection : Page collection/recherche affichant la liste paginée et les filtres.
/details?id=X : Page ressource affichant les détails complets d'un pirate spécifique.
/about : Page à-propos contenant la documentation et la gestion de projet.

**API et Endpoints :**

Ce projet exploite une API dédiée à One Piece. Les endpoints utilisés sont :
GET /characters : Récupération de l'intégralité des personnages.
GET /characters/{id} : (Simulé via filtrage interne) pour les détails spécifiques.