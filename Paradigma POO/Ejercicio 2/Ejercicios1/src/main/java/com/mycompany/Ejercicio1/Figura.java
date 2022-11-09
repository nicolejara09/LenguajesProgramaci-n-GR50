/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.Ejercicio1;

/**
 *
 * @author Nicole Jara
 */

public abstract class Figura {

    String elemento;
    int altura;
    int largo;
    int ancho;
    int vertices;
    int perimetro;

    public Figura(String elemento ) {
        this.elemento = elemento;
        this.altura = 0;        
        this.largo = 0;
        this.ancho = 0;
        this.vertices = 0;
        this.perimetro = 0;
    }

    public abstract int getArea();

    public abstract String toString();
}