#!/bin/bash

middleSum=0

ruleSet=$(grep '|' input.txt)


for line in $(grep -v '|' input.txt)
do
	includedRules=''
	numbers=$(echo "$line" | tr ',' '\n')
	for number in $(echo "$numbers"); do
		for n in $(echo "$numbers"); do
			if [ $number -eq $n ]
			then
				continue;
			fi
			includedRules+=$(echo "$ruleSet" | grep -E "$number\|$n" | awk -F'|' '{print $1, $2}')
			includedRules+=$'\n'
			includedRules+=$(echo "$ruleSet" | grep -E "$n\|$number" | awk -F'|' '{print $1, $2}')
			includedRules+=$'\n'
		done
	done
	includedRules=$(printf "%s" "$includedRules" | grep -vE "^$")
	sorted=$(echo "$includedRules" | sort | uniq | tsort)
	sortedLine=$(echo "$sorted" | while read -r number; do
		echo $line | tr ',' '\n' | (grep -w $number || true)
	done | paste -sd ','
	)
	echo "---"
	echo "$line"
	echo "$sortedLine"
	if [ "$line" == "$sortedLine" ]; then
		IFS=',' read -r -a array <<< "$sortedLine"
		middleIndex=$(( (${#array[@]} -1) / 2))
		middleNumber=${array[$middleIndex]}
		echo "mid: $middleNumber"
		middleSum=$((middleSum + middleNumber))
		echo "running Sum: $middleSum"
	fi
done

echo "final: $middleSum"

