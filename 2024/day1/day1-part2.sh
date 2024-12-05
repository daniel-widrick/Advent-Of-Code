#/bin/bash

## Simple problems are solved with simple scripts

#Read input.txt | print the first column | sprt the list numerically
col1=$(cat input.txt | awk '{print $1}' | sort -n)
#Same thing second group
col2=$(cat input.txt | awk '{print $2}' | sort -n | uniq -c)


lineCount=$(cat input.txt | wc -l)
i=1
similaritySum=0

#loop over each line
while [ $i -le $lineCount ]
do
	#Extract the <i> line from each dataset
	col1Value=$(echo "$col1" | sed -n "${i}p")
	#Get Occurences from second list
	occurence=$(echo "$col2" | grep -E " $col1Value\$" || echo 0)
	count=$(echo "$occurence" | awk '{print $1}')
	similaritySum=$(( col1Value * count + similaritySum ))
	echo -e "$col1Value\t$count\t$similaritySum"
	i=$((i + 1))
done

echo "Final sum: $similaritySum"
