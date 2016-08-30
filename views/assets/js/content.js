summary_noimg = 350;
summary_img = 375;
img_thumb_height = 150;
img_thumb_width = 200;

jQuery(document).ready(function () {

    prettyPrint();

    jQuery(".join-form form").on("submit", function (e) {
        var email = jQuery(this).serializeArray()[0].value;
        if (jQuery.trim(email).length < 3){
            e.preventDefault();
            jQuery(".join-form form div input").css({"border-color":"red"});
        }
    });

    jQuery(".join-form form div input").on("focus", function () {
        jQuery(".join-form form div input").css({"border-color":"#ccc"});
    });

    jQuery("form#login").on("submit", function (e) {
        e.preventDefault();
        alert("This feature is under developemnt. Keep patience.")
    });
    
    jQuery("form#registrationForm").on("submit", function (e) {
        e.preventDefault();
        alert("This feature is under developemnt. Keep patience.")
    });

});

function isValidEmail(email) {
    var regex = /^([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    return regex.test(email);
}


function removeHtmlTag(strx, chop) {
    if (strx.indexOf("<") != -1) {
        var s = strx.split("<");
        for (var i = 0; i < s.length; i++) {
            if (s[i].indexOf(">") != -1) {
                s[i] = s[i].substring(s[i].indexOf(">") + 1, s[i].length);
            }
        }
        strx = s.join("");
    }
    chop = (chop < strx.length - 1) ? chop : strx.length - 2;
    while (strx.charAt(chop - 1) != ' ' && strx.indexOf(' ', chop) != -1) chop++;
    console.log()
    strx = strx.substring(0, chop - 1);
    return strx + '...';
}

function createSummaryAndThumb(pID) {
    var div = document.getElementById(pID);
    var imgtag = "";
    var img = div.getElementsByTagName("img");
    var summ = summary_noimg;
    if (img.length >= 1) {
        imgtag = '<img src="' + img[0].src + '" class="pbtthumbimg"/>';
        summ = summary_img;
    }

    var summary = imgtag + '<div>' + removeHtmlTag(div.innerHTML, summ) + '</div>';
    div.innerHTML = summary;
}
