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

	//init object globally
	var objImage= null;
	function init(){
		objImage=document.getElementById("spaceship");				
		objImage.style.position='relative';
		objImage.style.left='0px';
		objImage.style.top='0px';
	}
	function getKeyAndMove(e){				
		var key_code=e.which||e.keyCode;
		switch(key_code){
			case 37: //left arrow key
				moveLeft();
				break;
			case 38: //Up arrow key
				moveUp();
				break;
			case 39: //right arrow key
				moveRight();
				break;
			case 40: //down arrow key
				moveDown();
				break;						
		}
	}
	function moveLeft(){
		objImage.style.left=parseInt(objImage.style.left)-5 +'vw';
	}
	function moveUp(){
		objImage.style.top=parseInt(objImage.style.top)-5 +'vw';
	}
	function moveRight(){
		objImage.style.left=parseInt(objImage.style.left)+5 +'vw';
	}
	function moveDown(){
		objImage.style.top=parseInt(objImage.style.top)+5 +'vw';
	}
	window.onload=init;

