//Testing some javascript for the first time! Don't judge me...


function goDirection(button){
    var x = button.id;
    switch (x){
        case 'left':
            document.getElementById("demo").innerHTML="Going Left!";
            break;
        case 'up':
            document.getElementById("demo").innerHTML="Going Up!";
            break;
        case 'down':
            document.getElementById("demo").innerHTML="Going Down!";
            break;
        case 'right':
            document.getElementById("demo").innerHTML="Going Right!";
            break;
        default:
            document.getElementById("demo").innerHTML="Something is wrong!";


    }
}

