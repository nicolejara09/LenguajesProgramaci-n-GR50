/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicios2;

/**
 *
 * @author Nicole Jara
 */
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

public class Biblioteca {
    
    private String nombre;
    private LinkedList<Libro> libros;
    private LinkedList<Socio> socios;
    private LinkedList<Prestamo> prestamos;
    private int maximaCantidadLibrosSocios=3;


    public Biblioteca(String nombre) {
        this.nombre = nombre;

        this.libros = new LinkedList<Libro>();
        this.socios = new LinkedList<Socio>();
        this.prestamos = new LinkedList<Prestamo>();
    }

    public void listarLibro(Libro libros){
        this.libros.add(libros);

    }

    public void listarSocio(Socio socios) {
        this.socios.add(socios);
    }

    public void realizarPrestamo(Prestamo realizarprestamo) {
        boolean revisarLibro=true;
        for (Socio s : this.socios){
            if(realizarprestamo.getIdSocioPrestamo()==s.getIdSocio()){
            for (Libro l : this.libros){
            if ((l.getId() == realizarprestamo.getIdLibroPrestamo()) && (revisarLibro == l.isDisponibilidad())){
                this.prestamos.add(realizarprestamo);
                
                s.setLibrosPrestados(s.getLibrosPrestados()+1);
                
                l.setDisponibilidad(false);
                
                System.out.println("El prestamo con el codPrestamo "+realizarprestamo.getIdPrestamo()+" se realizo exitosamente" +"\n");
                return;
            }
        }
                
            }
        }
        return;
    }


    public String sociosMaximaCantidad(){
        List<Socio> socios = this.socios.stream()
        .filter(socio -> socio.getLibrosPrestados()> this.maximaCantidadLibrosSocios)
        .collect(Collectors.toList());
        
        final StringBuffer sb = new StringBuffer("\n");
        for (Socio s : socios) {
            sb.append(s.toString() + "\n");
        }
        return sb.toString();
    }

    public String printDatos() {
        final StringBuffer sb = new StringBuffer("\n");
        for (Socio socios : this.socios) {
            sb.append(socios.toString() + "\n");
        }
        for (Libro crearLibro : this.libros) {
            sb.append(crearLibro.toString() + "\n");
        }
        for (Prestamo crearPrestamo : this.prestamos) {
            sb.append(crearPrestamo.toString() + "\n");
        }
        return sb.toString();
    }

    public String toString() {
        final StringBuffer sb = new StringBuffer("Biblioteca=");
        sb.append(this.nombre);
        sb.append(printDatos());
        return sb.toString();
    }

}