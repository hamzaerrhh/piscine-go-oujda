curl -s "https://learn01oujda.ma/assets/superhero/all.json" | jq -r '.[] | select(.id==170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'
