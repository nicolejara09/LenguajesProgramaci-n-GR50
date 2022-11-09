/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package com.mycompany.ejercicios2;

/**
 *
 * @author Nicole Jara
 */
public class main {
    public static void main(String[] args) {

        Biblioteca laBiblioteca = new Biblioteca("Biblioteca1");
        
        //AGREGANDO USUARIOS Y LIBROS.

        laBiblioteca.listarLibro(new Libro(8, "Juanita la del barrio", "El chavo del ocho", false));
        laBiblioteca.listarSocio(new Socio(5, "Margarita", "Guanajuato", 9));


        laBiblioteca.listarLibro(new Libro(9, "Me duele la lengua", "Pablito Clavo un clavito", true));
        laBiblioteca.listarSocio(new Socio(6, "Fermin", "Rio Cuarto", 3));

        laBiblioteca.listarLibro(new Libro(10, "La llorona", "Jose María Figueres", true));
        laBiblioteca.listarSocio(new Socio(7, "Hellen", "Tokio", 8));
        
        laBiblioteca.listarLibro(new Libro(11, "Las mañanitas", "Mario Bros", true));
        laBiblioteca.listarSocio(new Socio(15, "Rose", "Seul", 1));
        
        System.out.println(laBiblioteca.toString());

        //AGREGANDO PRESTAMOS.
        laBiblioteca.realizarPrestamo(new Prestamo(1, 5, 8, "27 de octubre 2022")); //EL LIBRO NO SE ENCUENTRA DISPONIBLE, POR LO TANTO NO HAY PRESTAMO.
        laBiblioteca.realizarPrestamo(new Prestamo(2, 6, 9, "27 de octubre 2022")); //EL SOCIO NO SE ENCUENTRA, POR LO TANTO NO HAY PRESTAMO.
        laBiblioteca.realizarPrestamo(new Prestamo(3, 7, 10, "27 de octubre 2022")); //EL LIBRO NO SE ENCUENTRA, POR LO TANTO NO HAY PRESTAMO.
        laBiblioteca.realizarPrestamo(new Prestamo(4, 15, 11, "27 de octubre 2022")); //EL LIBRO NO SE ENCUENTRA, POR LO TANTO NO HAY PRESTAMO.

        System.out.println(laBiblioteca.toString());
        
        //BUSCANDO LOS SOCIOS QUE TENGAN MAS DE 3 LIBROS PRESTADOS.
        System.out.println(laBiblioteca.sociosMaximaCantidad());
    }
}
