#!/usr/bin/gnuplot
set style line 1 lt 1 lw 2 pt 1 linecolor rgb "red"
set style line 2 lt 2 lw 1 pt 1 linecolor rgb "green"
set style line 3 lt 1 lw 1 pt 1 linecolor rgb "blue"
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output "./tex/img/z1.1/refuses.pdf"
set ylabel "Отказы"
set xlabel "Количество операторов"
plot "./dat/refuses.dat" using 1:2 smooth csplines title 'Отказы' w l ls 1
##########################
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output "./tex/img/z1.1/koeffzagrop.pdf"
set ylabel "Загрузка операторов"
set xlabel "Количество операторов"
plot "./dat/refuses.dat" using 1:3 smooth csplines title 'Коэффициент загрузки' w l ls 1