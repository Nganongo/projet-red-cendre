# Terre de Cendres

Mini-jeu de rôle en ligne de commande (CLI) écrit en Go, dans un univers post-apocalyptique.

Projet réalisé dans le cadre du Projet RED — Ynov Campus Aix.

**Équipe :** Christ, Alexandre, Jordan

## Présentation du projet

2189, douze ans après la Chute. Le joueur incarne un survivant qui sort d'un bunker pour rejoindre la dernière cité libre, dans un monde de ferraille et de radiations. Il doit gérer son inventaire, s'équiper auprès du marchand et du forgeron, apprendre des techniques de combat, et affronter des créatures mutées au combat tour par tour.

### Fonctionnalités principales

- Création de personnage : choix du nom et de la classe (Survivant, Mutant, Blindé), chacune avec ses propres PV max
- Gestion de l'inventaire (limité à 10 emplacements, extensible chez le marchand)
- Marchand : achat de potions, matériaux et sorts contre de l'or
- Forgeron : fabrication d'équipement à partir de matériaux récupérés, avec bonus de PV
- Combat tour par tour contre un monstre (chien-cendre), avec attaque de base et objets d'inventaire

### Fonctionnalités bonus

- Initiative : détermine qui commence le combat selon la rapidité de chacun
- Système d'expérience et de montée de niveau (avec report de l'XP excédentaire et bonus de PV)
- Sorts offensifs (coup de poing, boule de feu) alimentés par un système de mana
- Habillage visuel en ASCII art (écran de titre, monstre, victoire/défaite)

## Prérequis

- [Go](https://go.dev/dl/) version 1.22 ou supérieure

## Installation

```bash
git clone https://github.com/Nganongo/projet-red-personnage-pokemon.git
cd projet-red-personnage-pokemon
```

## Lancement

Depuis la racine du dépôt :

```bash
go run ./src
```

## Structure du projet

```
src/
├── main.go              # point d'entrée, menu principal
└── character/           # logique du jeu
    ├── character.go         # structure Personnage et affichage
    ├── charactercreation.go # création du personnage par le joueur
    ├── accessInventory.go   # gestion et utilisation de l'inventaire
    ├── marchand.go          # boutique
    ├── forgeron.go          # fabrication d'équipement
    ├── equipement.go        # structure Equipement
    ├── monstre.go           # structure Monstre
    ├── goblinpattern.go     # comportement du monstre en combat
    ├── characterturn.go     # menu de combat du joueur
    ├── trainingfight.go     # boucle de combat
    ├── spellbook.go         # apprentissage des sorts
    └── ascii.go             # visuels ASCII art
```
