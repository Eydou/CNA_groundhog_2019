set terminal pngcairo
set output 'ground2.png'
set multiplot layout 3, 1 title "Groundhog by Edouard Touch" font ",14"
set tmargin 2
set title 'standard deviation'
unset key
set grid
plot 'data.csv' u 1:4 w lp title 'standard deviation,'
set title 'relative temperature evolution'
plot 'data.csv' u 1:3 w lp title 'relative temperature evolution'
set title 'temperature increase average'
plot 'data.csv' u 1:2 w lp title 'temperature increase average'
unset multiplot