

// global array variables instantiated on start
var idArray;
var valueArray;
var solvedArray;
var current;

function testFunction(x, y, z){
    idArray = x;
    valueArray = y;
    solvedArray = z;
    
    startGame();
}
// Uses the array variables to set the game's starting values
function startGame(){
    // alert(idArray);
     for (var i =0; i< idArray.length; i++){
          if (valueArray[i] != 0){
             document.getElementById(idArray[i]).value = valueArray[i]+"";
            document.getElementById(idArray[i]).type = 'input1';
            document.getElementById(idArray[i]).disabled = true;
                                    
        }
     }    
}

// updates the board to show the completed solved puzzle
function solve(){
    for (var i =0; i< idArray.length; i++){
    if (game[i] == 0){
        document.getElementById(idArray[i]).value = solvedArray[i]+"";
        document.getElementById(idArray[i]).type = 'solved';
        document.getElementById(idArray[i]).disabled = true; 
        }   
    }
}

// clears the board
function clearBoard(){
    for (var i =0; i< idArray.length; i++){
        document.getElementById(idArray[i]).value = "";
        document.getElementById(idArray[i]).type = 'solved';
        document.getElementById(idArray[i]).disabled = false; 
    }
    
}

function getCurrent(clicked_id){

    current = clicked_id;
    // alert(clicked_id);
}

function solveCell(){

    if (current == null){
        alert("Please select a Cell");
    }
    //alert(current);
    document.getElementById(current).value = "F";
    for (var i =0; i< idArray.length; i++){
    if (current == idArray[i]){
        document.getElementById(idArray[i]).value = solvedArray[i]+"";
        document.getElementById(idArray[i]).type = 'solved';
        document.getElementById(idArray[i]).disabled = false;
        document.getElementById(idArray[i]).className = 'current';
        current = null; 
        
        }   
    }
    
    
}
function resetBoard(){
    for (var i =0; i< idArray.length; i++){
          if (valueArray[i] != 0){
             document.getElementById(idArray[i]).value = valueArray[i]+"";
            document.getElementById(idArray[i]).type = 'input1';
            document.getElementById(idArray[i]).disabled = true;
                                    
        }
        else{
            document.getElementById(idArray[i]).value = "";
        document.getElementById(idArray[i]).type = 'solved';
        document.getElementById(idArray[i]).disabled = false; 

        }
     }    
}