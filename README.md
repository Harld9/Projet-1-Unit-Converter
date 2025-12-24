# Projet 1 Convertisseur d'unités

Petit projet scolaire pour apprendre et revoir des notions de Go. C'est un convertisseur qui permet de passer des km en miles, des kg en livres et des degrés Celsius en Fahrenheit.

## Phase 1 - Convertisseur d'unités dans le terminal

C'est un programme simple qu'on lance dans la console.

### Comment l'utiliser

1. On lance le programme.
2. On choisit ce qu'on veut convertir dans le menu (1, 2 ou 3).
3. On tape le nombre à convertir.
4. Le résultat s'affiche.

Pour quitter, il suffit de taper 4 dans le menu.

cd "Phase 1 - Version Go (Terminal)"
go run main.go

Infos sur le code
J'ai utilisé des boucles pour que le programme ne s'arrête pas tout seul.
J'ai géré les erreurs (si on tape des lettres au lieu de chiffres).
Les calculs sont précis, je n'ai pas arrondi les résultats.

## Phase 2 - Convertisseur d'unités avec interface web