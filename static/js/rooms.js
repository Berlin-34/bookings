function setupRoomAvailability(roomID, csrfToken) {
    document.getElementById("check-availability-button").addEventListener("click", function () {
        let html = `
        <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
            <div class="d-flex flex-row justify-content-evenly" id="reservation-dates-modal">
                <div class="p-2">
                    <input type="text" disabled name="start" id="start" class="form-control" placeholder="Arrival" required>
                </div>
                <div class="p-2">
                    <input type="text" disabled name="end" id="end" class="form-control" placeholder="Departure" required>
                </div>
            </div>
        </form>
        `

        attention.custom({
            msg: html,
            title: "Choose your dates",

            willOpen: () => {
                const elem = document.getElementById("reservation-dates-modal");
                const rp = new DateRangePicker(elem, {
                    format: 'yyyy-mm-dd',
                    showOnFocus: true,
                    buttonClass: 'btn',
                    orientation: 'top',
                    minDate: new Date()
                })
            },

            didOpen: () => {
                document.getElementById("start").removeAttribute("disabled")
                document.getElementById("end").removeAttribute("disabled")
            },

            callback: function (result) {
                if (result) {
                    let form = document.getElementById("check-availability-form");
                    let formData = new FormData(form);
                    formData.append("csrf_token", csrfToken)
                    formData.append("room_id", roomID)

                    fetch("/search-availability-json", {
                        method: "post",
                        body: formData,
                    })
                        .then(response => response.json())
                        .then(data => {
                            if (data.ok) {
                                attention.custom({
                                    icon: "success",
                                    msg: "<p>Room is available</p>"
                                        + "<p><a href='/book-room?id="
                                        + data.room_id
                                        + "&s="
                                        + data.start_date
                                        + "&e="
                                        + data.end_date
                                        + "' class='btn btn-primary'>"
                                        + "Book Now!</a></p>",
                                    showConfirmButton: false
                                })
                            } else {
                                attention.error({
                                    msg: "No availability",
                                })
                            }
                        })
                }
            }
        })
    })
}
