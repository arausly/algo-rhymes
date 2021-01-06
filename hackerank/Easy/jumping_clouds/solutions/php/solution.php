<?php

// I had solved this earlier, I think i did a great job.
function jumpingOnClouds($c) {
    $count = count($c);
    $steps = 0;
    for($i = 0; $i< $count; $i++){
        if(isset($c[($i + 2)]) && $c[($i + 2)] != 1){
            $steps++;
            $i = $i + 1;
        }elseif(isset($c[($i + 1)]) && $c[($i + 1)] != 1){
            $steps++;
        }
    }
    return $steps;
}