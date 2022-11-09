/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicios2;

/**
 *
 * @author Nicole Jara
 */
public class Prestamo {
    
    private int idPrestamo;
    private int idSocio;
    private int idLibro;
    private String fechaPrestamo;

    public Prestamo(int idPrestamo, int idSocio, int idLibro, String fechaPrestamo) {
        this.idPrestamo = idPrestamo;
        this.idSocio = idSocio;
        this.idLibro = idLibro;
        this.fechaPrestamo = fechaPrestamo;
    }

    public int getIdPrestamo() {
        return idPrestamo;
    }

    public int getIdSocioPrestamo() {
        return idSocio;
    }

    public int getIdLibroPrestamo() {
        return idLibro;
    }

    public String toString() {
        return "Prestamo{" + " idPrestamo=" + idPrestamo + ", idSocio=" + idSocio + ", idLibro=" + idLibro + ", fechaPrestamo='" + fechaPrestamo + '\'' + '}';
    }
}