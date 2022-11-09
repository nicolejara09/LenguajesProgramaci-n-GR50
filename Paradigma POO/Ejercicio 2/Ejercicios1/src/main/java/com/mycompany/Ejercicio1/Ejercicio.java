/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.Ejercicio1;

/**
 *
 * @author Nicole Jara
 */
public class Ejercicio {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {

        
        Circulo circuloDatos = new Circulo("Circulo", 55);

        Cuadrado cuadradoDatos = new Cuadrado("Cuadrado", 15);

        Texto textoDatos = new Texto("Stan Dreamcatcher");
        
        Texto textoDatos1 = new Texto("Ojalá pasar lenguajes :c");

        Grupo grupoDatos = new Grupo("Grupo Mutual La Vivienda");

        grupoDatos.addFigura(circuloDatos);

        EditorGrafico editorGraficoDatos = new EditorGrafico("Creación de Datos");

        editorGraficoDatos.addFigura(circuloDatos);
        editorGraficoDatos.addFigura(cuadradoDatos);
        editorGraficoDatos.addGrupo(grupoDatos);
        editorGraficoDatos.addTexto(textoDatos);
        editorGraficoDatos.addTexto(textoDatos1);

        System.out.println("Elementos creados" + editorGraficoDatos);
    }
}