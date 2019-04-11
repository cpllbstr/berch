#!/usr/bin/gnuplot
set style line 1 lt 1 lw 2 pt 1 linecolor rgb "red"
set style line 2 lt 2 lw 1 pt 1 linecolor rgb "green"
set style line 3 lt 1 lw 1 pt 1 linecolor rgb "blue"
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
set term svg enhanced background rgb 'white'
set output "./img/z1.1/refuses.svg"
set ylabel "Отказы"
set xlabel "Количество операторов"
plot "./dat/refuses.dat" using 1:2 smooth csplines title 'Отказы' w l ls 1
##########################
set term svg enhanced background rgb 'white'
set output "./img/z1.1/koeffzagrop.svg"
set ylabel "Загрузка операторов"
set xlabel "Количество операторов"
plot "./dat/refuses.dat" using 1:3 smooth csplines title 'Коэффициент загрузки' w l ls 1