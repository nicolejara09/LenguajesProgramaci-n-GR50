/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.Ejercicio1;

/**
 *
 * @author Nicole Jara
 */
public class Cuadrado extends Figura {

    int lado;
    
    public Cuadrado(String elemento, int lado) {

        super(elemento);
        this.lado = lado;
    }

    public int getArea() {
        return lado * lado;
    }

    public String toString() {
        return "Cuadrado{" + "  elemento='" + elemento + '\'' + ", lado=" + lado + ", altura=" + altura + ", largo=" + largo + ", ancho=" + ancho + ", vertices=" + vertices + ", perimetro=" + perimetro +'}';
    }
}