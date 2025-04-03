# SAAS Template - Go Backend & Next.js Frontend

Ce projet est un template de base pour une application SAAS (Software as a Service) avec :
- Backend en Go
- Frontend en Next.js
- Base de données MySQL
- Architecture Dockerisée

## Architecture

### Backend (Go)
- API REST avec Gorilla Mux
- Gestion d'authentification
- Middleware de sécurité
- Connexion à MySQL
- Gestion des sessions utilisateur

### Frontend (Next.js)
- Interface utilisateur moderne
- Authentification
- Dashboard utilisateur
- Affichage des données en temps réel
- Design responsive avec Tailwind CSS

### Base de données (MySQL)
- Stockage des utilisateurs
- Gestion des sessions
- Données de trading

## Fonctionnalités

- Inscription et connexion utilisateur
- Gestion de profil
- Système d'upgrade utilisateur
- Affichage de données de paires de trading crypto en temps réel
- Interface utilisateur moderne et responsive

## Prérequis

- Docker
- Docker Compose

## Installation et Lancement

1. Cloner le repository :
```bash
git clone [URL_DU_REPO]
cd SAAS-TEMPLATE-GO
```

2. Lancer l'application :
```bash
docker-compose up
```

L'application sera disponible sur :
- Frontend : http://localhost:3000
- Backend : http://localhost:8080
- Base de données : localhost:3306

## Variables d'Environnement

Les variables d'environnement sont configurées dans le `docker-compose.yaml` :

- `DB_HOST`: Adresse de la base de données
- `DB_USER`: Utilisateur MySQL
- `DB_PASSWORD`: Mot de passe MySQL
- `DB_NAME`: Nom de la base de données
- `DB_PORT`: Port MySQL

## Licence

Ce projet est sous licence AGPL-3.0. Voir le fichier `LICENSE` pour plus de détails.
