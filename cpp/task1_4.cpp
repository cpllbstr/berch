#include <iostream>
#include <math.h>  
#include <cstdio>

using namespace std;

FILE *scrn,*datan,*scrn1,*datan1; 

int N = 1;
int N_max = 11;
int Ts = 229;
int Tc = 45;
int Tw = 553;
int queue = 30;
int q2 = 1;

double lamb = (double)1/Tc;
double mu = (double)1/Ts;
double v = (double)1/Tw;

int channels[11];
double queue_length[11];
double load_coef[11];
double wait_time[11];

int queue_ar[65];
double p0[65];


long int fact(int x)
{
    if(x < 0)
        return 0;
    if (x == 0)
        return 1;
    else
        return x * fact(x - 1);
}

int main()
{ 
	double* P;
	int n = N;
	for(int i = 0; i < (N_max - N + 1); i++)
	{
		channels[i] = n;
		queue_length[i] = 0;
		load_coef[i] = 0;
		wait_time[i] = 0;
		n = n + 1;
	}
    for(int i = 0; i < 100; i++)
    {
        queue_ar[i] = i+1;
    }

	n = N;
	double Pn;
	double inf_part;
	int i,k;
	double sum;
	double znam;
	double znam1;
    double ql;
    double N_av;
	datan = fopen("datan.txt","w");
    datan1 = fopen("datan1.txt","w");

	while(n < (N_max + 1))
	{
		P = new double[queue+n+1];
		for(int i = 0; i < queue+n+1; i++)
		{
			P[i] = 0;
		}
		i = 0;
		sum = 0;
		znam = 1;
		znam1 = 1;

		Pn = pow(lamb,n)/(fact(n)*pow(mu,n));
    	inf_part = 0;

    	//-------------------находим P0---------------------
        //--------------------------------------------------
    	while(i < (n+1))
    	{
    		sum  = sum + pow(lamb,i)/(fact(i)*pow(mu,i));
    		i += 1;
    	}

    	k = 1;
    	while(k < (queue+1))
    	{
            znam = 1;
    		for(int j = 1; j < (k+1); j++)
    		{
    			znam = znam*(n*mu+j*v);
    		}
    		inf_part = inf_part + pow(lamb,k)/znam;
    		k += 1;
    	}
    	P[0] = (double)1/(sum+Pn*inf_part);

        //---------Доп задание Берчуна------------

        if(n == 7)
        {
            while(q2 < 66)
            {
                k = 1;
                while(k < (q2+1))
                {
                    znam = 1;
                    for(int j = 1; j < (k+1); j++)
                    {
                        znam = znam*(n*mu+j*v);
                    }
                    inf_part = inf_part + pow(lamb,k)/znam;
                    k += 1;
                }
                p0[q2-1] = (double)1/(sum+Pn*inf_part);
                ++q2;
            }
            q2 = 1;
        }
        //--------------------------------------------------
        //--------------------------------------------------


    	//-находим предельные вероятности состояний системы-
        //--------------------------------------------------
    	i = 1;
    	while(i < n+1)
    	{
        	P[i] = pow(lamb,i)/(fact(i)*pow(mu,i))*P[0];
        	i += 1;
        }
        //--------------------------------------------------
        //--------------------------------------------------


        //-----находим вероятности состояний с очередью-----
        //--------------------------------------------------
        i = 1;
        while(i < queue+1)
        {
            znam1 = 1;
        	for(int j = 1; j < (i+1); j++)
    		{
    			znam1 = znam1*(n*mu+j*v);
    		}
        	P[n+i] = (pow(lamb,i)*P[n])/znam1;
        	i += 1;
        }
        //--------------------------------------------------
        //--------------------------------------------------

        ql = 0;
        N_av = 0;
        for(int i = 1; i < (queue+ n + 1); i++)
        {
            if(i < (n+1))
                N_av += i*P[i];
            else
                N_av += n*P[i]; 
        } 
        for(int i = 1; i < (queue+1); i++)
        {
        	ql = ql + i*P[n+i];
        }
        queue_length[n-N] = ql;
        wait_time[n-N] = queue_length[n-N]/lamb;
        load_coef[n-N] = N_av/n;

        fprintf(datan,"%d\t%lf\t%lf\t%lf\n",channels[n-N],queue_length[n-N], wait_time[n-N], load_coef[n-N]);
        
        if(n == 7)
        {
            for(int i = 0; i < 65; i++)
            {
                fprintf(datan1,"%d\t%lf\n",queue_ar[i],p0[i]);
            }
        }

		n += 1;
		delete[] P;
	}

	scrn = fopen("scriptn.dat","w");
    fprintf(scrn, "set xrange [1:11] \n");
	fprintf(scrn, "set xtics 1,1,11\n");
	fprintf(scrn, "set grid\n");
	fprintf(scrn,"plot 'datan.txt' using 1:2 w l lt rgb 'blue' lw 2 , 'datan.txt' u 1:2 w p pt 5 lc rgb 'red'\n");
	fprintf(scrn, "pause -1\n");
    fprintf(scrn,"plot 'datan.txt' using 1:3 w l lt rgb 'blue' lw 2 , 'datan.txt' u 1:3 w p pt 5 lc rgb 'red'\n");
    fprintf(scrn, "pause -1\n");
    fprintf(scrn,"plot 'datan.txt' using 1:4 w l lt rgb 'blue' lw 2 , 'datan.txt' u 1:4 w p pt 5 lc rgb 'red'\n");
    fprintf(scrn, "pause -1\n");

    scrn1 = fopen("scriptn1.dat","w");
    fprintf(scrn1, "set xrange [1:65] \n");
    fprintf(scrn1, "set xtics 1,1,65\n");
    fprintf(scrn1, "set grid\n");
    fprintf(scrn1,"plot 'datan1.txt' using 1:2 w l lt rgb 'blue' lw 2 , 'datan1.txt' u 1:2 w p pt 5 lc rgb 'red'\n");
    fprintf(scrn1, "pause -1\n");
}