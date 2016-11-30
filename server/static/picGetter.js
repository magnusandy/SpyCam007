function doCallSetCurrentDates()
{
var currentDate = new Date();
var tomorrowDate = new Date();
tomorrowDate.setDate(tomorrowDate.getDate() + 1);
document.getElementById("firstDate").valueAsNumber = currentDate.getTime();
document.getElementById("secondDate").valueAsNumber = tomorrowDate.getTime();
doCall();
}

function doCall() {
  startTime = (Date.parse(document.getElementById("firstDate").value)/1000)+21600 //divide by 1000 to convert from miliseconds to seconds then add 21600 which is 6 hours in seconds to align unix time with local time
  endTime = (Date.parse(document.getElementById("secondDate").value)/1000)+21600
    var query = "?startTime=" + startTime + "&endTime=" + endTime;
    $.ajax({
        url: "/pictures/" + query,
        success: function(results) {
            linksDIV = document.getElementById("links");
            while (linksDIV.firstChild) {
                linksDIV.removeChild(linksDIV.firstChild);
            }
            console.log(results);
            for(index in results)
            {
                pic = document.createElement("a");
                pic.href = results[index].Url;
                pic.title = results[index].Timestamp.substring(0,10);
                linksDIV.appendChild(pic);
            }
            //gallery code
            initializeGallery();
            document.getElementById('links').onclick = function (event) {
                event = event || window.event;
                var target = event.target || event.srcElement,
                    link = target.src ? target.parentNode : target,
                    options = {index: link, event: event},
                    links = this.getElementsByTagName('a');
                blueimp.Gallery(links, options);
              };
            blueimp.Gallery(
              document.getElementById('links').getElementsByTagName('a'),
              {
                  container: '#blueimp-gallery-carousel',
                  carousel: true
              }
          );
            var galleryDiv = document.getElementById("longgallery");
            while (galleryDiv.firstChild) {
                galleryDiv.removeChild(galleryDiv.firstChild);
            }
            for (index in results) {
                var pic = document.createElement("img");
                pic.className += " longpic col-xs-8 col-xs-offset-2 col-md-offset-0 col-md-3";
                galleryDiv.appendChild(pic);
                pic.src=results[index].Url;
            }
        },
        error: function(xhr,status,error) {
            console.error("Error: " + status + " " + error);
        }
    });
}
