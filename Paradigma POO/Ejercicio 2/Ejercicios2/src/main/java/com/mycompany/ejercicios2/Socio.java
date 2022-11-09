/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicios2;

/**
 *
 * @author Nicole Jara
 */
public class Socio {

    private int idSocio;
    private String nombreSocio;
    private String direccionSocio;
    private int librosPrestados;

    public Socio(int idSocio, String nombreSocio, String direccionSocio, int librosPrestados) {
        this.idSocio = idSocio;
        this.nombreSocio = nombreSocio;
        this.direccionSocio = direccionSocio;
        this.librosPrestados = librosPrestados;
    }
    
    public void setLibrosPrestados(int librosPrestados) {
        this.librosPrestados = librosPrestados;
    }

    public int getLibrosPrestados() {
        return librosPrestados;
    }

    public int getIdSocio() {
        return idSocio;
    }

    public String toString() {
        return "Socio{" + "idSocio=" + idSocio + ", nombreSocio='" + nombreSocio + '\'' + ", direccionSocio='" + direccionSocio + '\'' + ", librosPrestados=" + librosPrestados + '}';
    }
}
