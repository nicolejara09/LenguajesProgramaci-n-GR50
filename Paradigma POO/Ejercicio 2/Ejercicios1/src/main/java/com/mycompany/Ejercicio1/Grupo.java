/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.Ejercicio1;

/**
 *
 * @author Nicole Jara
 */

import java.util.ArrayList;

public class Grupo {

    private String elemento;
    private ArrayList<Figura> figuras;

    public Grupo(String elemento) {
        this.elemento = elemento;
        this.figuras = new ArrayList<>();
    }

    public ArrayList<Figura> getFiguras() {
        return figuras;
    }

    public void setFiguras(ArrayList<Figura> figuras) {
        this.figuras = figuras;
    }

    public void addFigura(Figura figuras1) {

        this.figuras.add(figuras1);
    }

    public String toString() {
        return "Grupo{" + "elemento='" + elemento + '\'' + ", figuras=" + figuras + '}';
    }
}
