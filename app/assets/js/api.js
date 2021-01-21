$(document).ready(function(){
    $(".error-display").hide()

    let APIName = "api_thirdparty"
    let APIKey = "AAhzclm9jHdypqdmEQx"
    // 
    // Redirect if cookie seted
    // 
    token = Cookies.get("access_token")
    if (token != "" && window.location.pathname == "/" ) {
        window.location = "/dashboard"
    }

    if (token == "" && window.location.pathname != "/" ) {
        window.location = "/"
    }
    


    var apiEndpoint = "http://localhost:9091/api"

    $("#login_button").on("click", function(){
        let username = $(".field-username").val()
        let password = $(".field-password").val()

        var form = new FormData();
        form.append("username", username);
        form.append("password", password);

        var settings = {
            "url": apiEndpoint + "/login",
            "method": "POST",
            "timeout": 0,
            "headers": {
                "X-API-KEY": APIKey,
                "X-API-NAME": APIName,
            },
            "processData": false,
            "mimeType": "multipart/form-data",
            "contentType": false,
            "data": form
        };

        $.ajax(settings).done(function (response) {
            data = JSON.parse(response)
            if (data.status == "success") {
                Cookies.set('access_token', data.data["token"]);
                window.location = "/dashboard"
            }

            if (data.status == "error") {
                $(".error-display").html(data.message)
                $(".error-display").show()
            } 
        });
    })

    // 
    // Upcoming Meetings
    // 
    if ($(".upcoming-meetings").length > 0 ){
        var form = new FormData();

        var settings = {
            "url": apiEndpoint + "/list-meetings/upcoming",
            "method": "POST",
            "timeout": 0,
            "headers": {
                "X-API-KEY": APIKey,
                "X-API-NAME": APIName,
                "X-ACCESS-TOKEN" : token,
            },
            "processData": false,
            "mimeType": "multipart/form-data",
            "contentType": false,
            "data": form
        };

        $.ajax(settings).done(function (response) {
            data = JSON.parse(response)
            console.log(data)
            if (data.status == "success") {
                let html_tag = template__meetings(data.data)
                $(".upcoming-meetings .upcoming-meetings--table").html(html_tag)

                var $tooltip = $('[data-toggle="tooltip"]');
                $tooltip.tooltip();
            }

            if (data.status == "error") {
                $(".error-display").html(data.message)
                $(".error-display").show()
            } 
        });
    } 

     // 
    // Upcoming Meetings
    // 
    if ($(".recent-meetings").length > 0 ){
        var form = new FormData();

        var settings = {
            "url": apiEndpoint + "/list-meetings/recent",
            "method": "POST",
            "timeout": 0,
            "headers": {
                "X-API-KEY": APIKey,
                "X-API-NAME": APIName,
                "X-ACCESS-TOKEN" : token,
            },
            "processData": false,
            "mimeType": "multipart/form-data",
            "contentType": false,
            "data": form
        };

        $.ajax(settings).done(function (response) {
            data = JSON.parse(response)
            if (data.status == "success") {
                let html_tag = template__meetings(data.data)
                $(".recent-meetings .recent-meetings--table").html(html_tag)

                var $tooltip = $('[data-toggle="tooltip"]');
                $tooltip.tooltip();

            }

            if (data.status == "error") {
                $(".error-display").html(data.message)
                $(".error-display").show()
            } 
        });
    } 
    

    function template__meetings( data ) {
        let html = ""
        $.each(data, function(index,val) {
            meeting__id = (val.meeting_id != null) ? val.meeting_id : "-"
            html += '<tr> \
                        <th scope="row">\
                        <div class="media align-items-center">\
                            <div class="media-body">\
                            <span class="name mb-0 text-sm">'+meeting__id+'</span>\
                            </div>\
                        </div>\
                        </th>\
                        <td>\
                        '+val.title+'\
                        </td>\
                        <td>\
                        <div class="avatar-group">\
                        '
                        
            $.each(val.participants, function(id,username){
                html += '\
                            <a href="#" class="avatar avatar-sm rounded-circle" data-toggle="tooltip" data-original-title="'+username+'">\
                            <img alt="Image placeholder" src="assets/img/user.png">\
                            </a>\
                '
            })
            html += '</div>\
                        </td>\
                        <td>\
                        '+new Date(val.start_time*1000).toLocaleString()+' - '+new Date((val.end_time*1000)).toLocaleString()+'\
                        </td>\
                    </tr>'
        })

       
        return html
    }
    // 
    // List users
    // 
    if ($(".field-schedule--participants").length > 0) {
        var form = new FormData();
        var settings = {
            "url": apiEndpoint + "/get-users",
            "method": "POST",
            "timeout": 0,
            "headers": {
                "X-API-KEY": APIKey,
                "X-API-NAME": APIName,
                "X-ACCESS-TOKEN" : token,
            },
            "processData": false,
            "mimeType": "multipart/form-data",
            "contentType": false,
            "data": form
        };

        $.ajax(settings).done(function (response) {
            data = JSON.parse(response)
            if (data.status == "success") {
                let html__option = ""
                $.each(data.data, function(key,val){
                    html__option += '<option value="'+val.id+'">'+val.username+'</option>'
                });

                $(".field-schedule--participants").html(html__option)
            }

            if (data.status == "error") {
                $(".error-display").html(data.message)
                $(".error-display").show()
            } 
        });
    }

    // 
    // On schedule 
    // 
    $(".schedule-meeting-submit").on("click", function(){
        var form = new FormData();
        form.append("Title", $(".field-schedule--title").val());
        form.append("Description", $(".field-schedule--description").val());
        form.append("StartTime", $(".field-schedule--starttime").val());
        form.append("EndTime", $(".field-schedule--endtime").val());
        form.append("Participants", $(".field-schedule--participants").val());

        startTime = new Date($(".field-schedule--starttime").val())
        endTime = new Date($(".field-schedule--endtime").val())

        if(startTime >= endTime){
            $(".error-display").html("Start time should be less than end time")
            $(".error-display").show()
            return
        }
        var settings = {
            "url": apiEndpoint + "/schedule-meeting",
            "method": "POST",
            "timeout": 0,
            "headers": {
                "X-API-KEY": APIKey,
                "X-API-NAME": APIName,
                "X-ACCESS-TOKEN" : token,
            },
            "processData": false,
            "mimeType": "multipart/form-data",
            "contentType": false,
            "data": form
        };

        $.ajax(settings).done(function (response) {
            data = JSON.parse(response)
            if (data.status == "success") {
                window.location = "/dashboard"
            }

            if (data.status == "error") {
                $(".error-display").html(data.message)
                $(".error-display").show()
            } 
        });
    })

    $(".logout").on("click", function(){
        Cookies.set("access_token","")
        window.location = '/'
    })
})