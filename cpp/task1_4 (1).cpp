#include <iostream>
#include <math.h>  
using namespace std;
FILE *scrn,*datan; 

int N = 1;
int N_max = 20;
int Ts = 229;
int Tc = 43;
int Tw = 553;
int queue = 10;

double lamb = (double)1/Tc;
double mu = (double)1/Ts;
double v = (double)1/Tw;

int channels[20];
double queue_length[20];
double load_coef[20];
double wait_time[20];


long int fact(int x)
{
    if(x < 0) // если пользователь ввел отрицательное число
        return 0; // возвращаем ноль
    if (x == 0) // если пользователь ввел ноль,
        return 1; // возвращаем факториал от нуля
    else // Во всех остальных случаях
        return x * fact(x - 1); // делаем рекурсию.
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

	n = N;
	double Pn;
	double inf_part;
	int i,k,ind;
	double sum;
	double znam;
	double znam1;
    double ql;
	datan = fopen("datan.txt","w");

	while(n < (N_max + 1))
	{
		P = new double[queue+n+1];
		i = 0;
		k = 1;
		sum = 0;
		znam = 1;
		znam1 = 1;

		Pn = (double)pow(lamb,n)/(fact(n)*pow(mu,n));
    	inf_part = 0;

    	//находим P0
    	while(i < n+1)
    	{
    		sum  = (double)sum + (double)pow(lamb,i)/(fact(i)*pow(mu,i));
    		i += 1;
    	}
    	while(k < 11)
    	{
    		for(int j = 1; j < k+1; j++)
    			{
    				znam = (double)znam*(n*mu+j*v);
    			}
    		inf_part = (double)inf_part + pow(lamb,k)/znam;
    		k += 1;
    	}
    	P[0] = (double)1/(sum + Pn*inf_part);

    	//находим предельные вероятности состояний системы
    	i = 1;
    	while(i < n+1)
    	{
        	P[i] = (double)pow(lamb,i)/(fact(i)*pow(mu,i))*P[0];
        	i += 1;
        }

        //находим вероятности состояний с очередью
        while(i < queue+n+1)
        {
        	for(int j = 1; j < i+1; j++)
    			{
    				znam1 = (double)znam1*(n*mu+j*v);
    			}
        	P[i] = (double)(pow(lamb,i)*P[n])/znam1;
        	i += 1;
        }
        
        //double a = (double)lamb/(n*mu);
        ql = 0;
        ind = 1;
        for(int t = 0; t < queue; t++)
        {
        	ql = (double)ql + (double)ind*P[n+t+1];
        	ind++;
        }

        queue_length[n-N] = ql;//(double)(P[n]*a)/pow(1-a,2);
        //wait_time[n-N] = (double)queue_length[n-N]/lamb;
        //load_coef[n-N] = (double)lamb/(mu*n);

        fprintf(datan,"%d\t%lf\n",channels[n-N],queue_length[n-N]);

		n += 1;
		delete[] P;
	}

	scrn = fopen("scriptn.dat","w");
	fprintf(scrn, "set grid ytics mytics\n");
	fprintf(scrn, "set grid\n");
	fprintf(scrn,"plot 'datan.txt' using 1:2 smooth bezier with line title 'Мат. ожидание длины очереди'\n");
	fprintf(scrn, "pause -1\n");
}