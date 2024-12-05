#/bin/bash

## Simple problems are solved with simple scripts

#Read input.txt | print the first column | sprt the list numerically
col1=$(cat input.txt | awk '{print $1}' | sort -n)
#Same thing second group
col2=$(cat input.txt | awk '{print $2}' | sort -n)


lineCount=$(cat input.txt | wc -l)
i=1
distanceSum=0

#loop over each line
while [ $i -le $lineCount ]
do
	#Extract the <i> line from each dataset
	col1Value=$(echo "$col1" | sed -n "${i}p")
	col2Value=$(echo "$col2" | sed -n "${i}p")
	#Get the distance
	distance=$((col1Value - col2Value))
	#Get the abs distance
	abs=$((distance < 0 ? -distance: distance))
	sum=$((sum + abs))
	echo -e "$col1Value\t$col2Value\t$distance\t$abs\t$sum"
	i=$((i + 1))
done

echo "Final sum: $sum"
