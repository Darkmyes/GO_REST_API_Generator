var proyectos = [];
var tables = [];
var proyect = {};
var table = {};
var currentView = null;

// DOM Elements
view = document.getElementById("view");
home = document.getElementById("home");
proyects = document.getElementById("proyects");
configuration = document.getElementById("configuration");
help = document.getElementById("help"); 

// Functions
CleanView = (v)=>{
    for (let i=(v.childElementCount-1); i>=0; i--){
        v.removeChild(v.childNodes[i]);
    }
}

HideChilds = (v)=>{
    for (let i=(v.childElementCount-1); i>=0; i--){
        v.childNodes[i].style.display = "none";
    }
}

// Navigation
moveTo = (option)=>{
    HideChilds(view);
    switch (option){
        case "home" : 
            home.style.display = "block";
            break;
        case "proyects" : 
            proyects.style.display = "block";
            break;
        case "configuration" : 
            configuration.style.display = "block";
            break;
        case "help" : 
            help.style.display = "block";
            break;
    }
}

// Home

// Proyects
newProyect = ()=>{

}

saveProyect = ()=>{

}

listProyects = ()=>{

}

// Configuration


// Help

