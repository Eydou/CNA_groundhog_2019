set terminal pngcairo
set output 'ground1.png'
set grid
set title 'Grounghog by Edouard Touch'
plot 'data.csv' u 1:4 w lp title 'standard deviation,', 'data.csv' u 1:3 w lp title 'relative temperature evolution', 'data.csv' u 1:2 w lp title 'temperature increase average'