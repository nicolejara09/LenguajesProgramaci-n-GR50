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

public class EditorGrafico{


    private String elemento;    
    private ArrayList<Figura> figuras;
    private ArrayList<Texto> textos;
    private ArrayList<Grupo> grupos;

    public EditorGrafico(String elemento) {
        this.elemento = elemento;
        this.figuras = new ArrayList<>();
        this.textos = new ArrayList<>();
        this.grupos = new ArrayList<>();
    }

    public ArrayList<Figura> getFiguras() {
        return figuras;
    }

    public ArrayList<Texto> getTextos() {
        return textos;
    }

    public ArrayList<Grupo> getGrupos() {
        return grupos;
    }

    public void addFigura(Figura figura) {
        figuras.add(figura);
    }

    public void addTexto(Texto texto) {
        textos.add(texto);
    }

    public void addGrupo(Grupo grupo) {
        grupos.add(grupo);
    }

    public String toString() {
        return "EditorGrafico {\n" + "elemento='" + elemento + '\'' + ",\n figuras=" + figuras + ",\n textos=" + textos + ",\n grupos=" + grupos + "\n" + '}';
    }
}
