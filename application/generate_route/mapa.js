

function initMap() {
    const garca = { lat: -22.217813, lng: -49.657082 };
    const map = new google.maps.Map(document.getElementById("map"), {
        zoom:14,
        center: garca,
    });
    
    const drawingManager = new google.maps.drawing.DrawingManager({
        drawingMode: google.maps.drawing.OverlayType.MARKER,
        drawingControl: true,
        drawingControlOptions: {
        position: google.maps.ControlPosition.TOP_CENTER,
        drawingModes: [
            google.maps.drawing.OverlayType.MARKER,
            google.maps.drawing.OverlayType.CIRCLE,
            google.maps.drawing.OverlayType.POLYGON,
            google.maps.drawing.OverlayType.POLYLINE,
            google.maps.drawing.OverlayType.RECTANGLE,
        ],
        },
        markerOptions: {
        icon: "https://developers.google.com/maps/documentation/javascript/examples/full/images/beachflag.png",
        },
        circleOptions: {
        fillColor: "#ffff00",
        fillOpacity: 1,
        strokeWeight: 5,
        clickable: false,
        editable: true,
        zIndex: 1,
        },
    });

    drawingManager.setMap(map);


    google.maps.event.addListener(map, "click", function (e) {
        var marker = new google.maps.Marker({
            position: e["latLng"],
            title: "Hello World!",
        });
        console.log(e["latLng"].lat(),e["latLng"].lng());
        marker.setMap(map);
    });
    
    google.maps.event.addListener(drawingManager, 'polygoncomplete', function(polygon) {
        const coords = polygon.getPath().getArray().map(coord => {
            return {
            lat: coord.lat(),
            lng: coord.lng()
            }
        });

        console.log(JSON.stringify(coords, null, 1));

        // SAVE COORDINATES HERE
    });
}
window.initMap = initMap;
