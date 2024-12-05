<?php

$filename = "input.txt";
$rawData = file_get_contents($filename);

$run = true;
$lastPos = 0;

$value=0;

while(true) {
	$start=strpos($rawData,"mul(",$lastPos);
	if($start === false) {
		break;
	}
	$end=strpos($rawData,")",$start);
	if($end === false) {
		break;
	}
	$contents=substr($rawData, $start+4, $end-$start-4);
	$nums=explode(",",$contents);
	if(count($nums) != 2) {
		$lastPos=$start+1;	
		continue;
	}
	if(!is_numeric($nums[0]) || !is_numeric($nums[1]) ||
		$nums[0] > 999 || $nums[1] > 999 ||
		$nums[0] <= 0 || $nums[1] <= 0) { //paranoia

		$lastPos=$start+1;
		continue;
	}
	echo "$contents\n";
	$value+=$nums[0]*$nums[1];
	$lastPos=$start+1;
}

echo "found $value\n";
