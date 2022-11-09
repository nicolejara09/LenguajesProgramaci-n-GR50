/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicios2;

/**
 *
 * @author Nicole Jara
 */

public class Libro {
    
    int id;
    String titulo;
    String autor;
    boolean disponibilidad;

    public Libro(int id, String titulo, String autor, boolean disponibilidad) {
        this.id = id;
        this.titulo = titulo;
        this.autor = autor;
        this.disponibilidad = disponibilidad;
    }


    public int getId() {
        return id;
    }

    public String getTitulo() {
        return titulo;
    }

    public String getAutor() {
        return autor;
    }

    public boolean isDisponibilidad() {
        return disponibilidad;
    }

    public void setDisponibilidad(boolean disponibilidad) {
        this.disponibilidad = disponibilidad;
    }

    public String toString() {
        return "Libro{" + "id=" + id + ", titulo=" + titulo + ", autor=" + autor + ", disponibilidad=" + disponibilidad + '}';
    }
    
}
