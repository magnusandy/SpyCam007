function doCall() {
  startTime = (document.getElementById("firstDate").valueAsNumber)/1000 //divide by 1000 to convert from miliseconds to seconds
  endTime = (document.getElementById("secondDate").valueAsNumber)/1000
    var query = "?startTime=" + startTime + "&endTime=" + endTime;
    $.ajax({
        url: "/pictures/" + query,
        success: function(results) {
            linksDIV = document.getElementById("links");
            while (linksDIV.firstChild) {
                linksDIV.removeChild(linksDIV.firstChild);
            }
            console.log(results)
            for(index in results)
            {
                pic = document.createElement("a");
                pic.href = results[index].Url;
                pic.title = results[index].Timestamp.substring(0,10);
                linksDIV.appendChild(pic);
            }
            blueimp.Gallery(
              document.getElementById('links').getElementsByTagName('a'),
              {
                  container: '#blueimp-gallery-carousel',
                  carousel: true
              }
          );
        },
        error: function(xhr,status,error) {
            console.error("Error: " + status + " " + error);
        }
    });
}
