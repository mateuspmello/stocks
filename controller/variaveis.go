package controller

import "text/template"

//Modelos armazenam os modelos de paginas que serão manipulados pelos controllers
var Modelos = template.Must(template.ParseFiles("html/ticker.html"))
