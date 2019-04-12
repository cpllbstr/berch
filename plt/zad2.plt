#!/usr/bin/gnuplot
set style line 2 lt 1 lw 2 pt 1 linecolor rgb "#e28cb3"
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
#############
set term svg enhanced background rgb 'white'
set output "./img/z2/nozh.svg"
set ylabel "Матожидание"
set xlabel "Число наладчиков"
plot "./dat/zad2.dat"  using 1:2 smooth csplines title  'Ожидающие наладки станки' w l ls 2
#############
set term svg enhanced background rgb 'white'
set output "./img/z2/nprost.svg"
set ylabel "Матожидание"
set xlabel "Число наладчиков"
plot "./dat/zad2.dat"  using 1:3 smooth csplines title  'Простаивающие станки' w l ls 2
#############
set term svg enhanced background rgb 'white'
set output "./img/z2/mz.svg"
set ylabel "Матожидание"
set xlabel "Число наладчиков"
plot "./dat/zad2.dat"  using 1:4 smooth csplines title  'Занятые наладчики' w l ls 2
#############
set term svg enhanced background rgb 'white'
set output "./img/z2/mzdn.svg"
set ylabel "Коэффициент занятости"
set xlabel "Число наладчиков"
plot "./dat/zad2.dat"  using 1:5 smooth csplines title  'Занятые наладчики' w l ls 2