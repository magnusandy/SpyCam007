function doCall() {
  startTime = (document.getElementById("firstDate").valueAsNumber)/1000; //divide by 1000 to convert from miliseconds to seconds
  endTime = (document.getElementById("secondDate").valueAsNumber)/1000;
    var query = "?startTime=" + startTime + "&endTime=" + endTime;
    $.ajax({
        url: "/pictures/" + query,
        success: function(results) {
            linksDIV = document.getElementById("links");
            console.log(results)
        },
        error: function(xhr,status,error) {
            console.error("Error: " + status + " " + error);
        }
    });
}
