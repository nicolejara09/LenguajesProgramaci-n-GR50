/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.Ejercicio1;

/**
 *
 * @author Nicole Jara
 */

public class Circulo extends Figura {
    
    double pi;
    int radio;

    public Circulo(String elemento, int radio) {
        super(elemento);
        this.pi = 3.14;        
        this.radio = radio;
    }

    public int getArea() {
        return 0;
    }

    public String toString() {
        return "Circulo{" + "  elemento='" + elemento + '\'' + ", pi=" + pi + ", radio=" + radio + ", altura=" + altura + ", largo=" + largo + ", ancho=" + ancho + ", vertices=" + vertices + ", perimetro=" + perimetro +'}';
    }
}
