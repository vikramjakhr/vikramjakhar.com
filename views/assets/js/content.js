summary_noimg = 350;
summary_img = 375;
img_thumb_height = 150;
img_thumb_width = 200;
hostName = "http://vikramjakhar.com";

jQuery(document).ready(function () {

    prettyPrint();
    jQuery("time.timeago").timeago();

    jQuery(".join-form form").on("submit", function (e) {
        var email = jQuery(this).serializeArray()[0].value;
        if (jQuery.trim(email).length < 3) {
            e.preventDefault();
            jQuery(".join-form form div input").css({"border-color": "red"});
        }
    });

    jQuery(".join-form form div input").on("focus", function () {
        jQuery(".join-form form div input").css({"border-color": "#ccc"});
    });

    jQuery("form#login").on("submit", function (e) {
        e.preventDefault();
        alert("This feature is under developemnt. Keep patience.")
    });

    jQuery("form#registrationForm").on("submit", function (e) {
        e.preventDefault();
        alert("This feature is under developemnt. Keep patience.")
    });

    jQuery("#scroll-up").hide();
    jQuery(function () {
        jQuery(window).scroll(function () {
            if (jQuery(this).scrollTop() > 200) {
                jQuery('#scroll-up').fadeIn();
            } else {
                jQuery('#scroll-up').fadeOut();
            }
        });
        jQuery('a#scroll-up').click(function () {
            jQuery('body,html').animate({
                scrollTop: 0
            }, 500);
            return false;
        });
    });
    jQuery('.commentTxtArea').on('paste input', function () {
        if ($(this).outerHeight() > this.scrollHeight) {
            $(this).height(40)
        }
        while ($(this).outerHeight() < this.scrollHeight + parseFloat($(this).css("borderTopWidth")) + parseFloat($(this).css("borderBottomWidth"))) {
            $(this).height($(this).height() + 1)
        }
    });

    jQuery('a.comment-link').click(function () {
        jQuery('body,html').animate({
            scrollTop: jQuery("#comments").offset().top
        }, 600);
        return false;
    });

    jQuery(function () {
        var title = jQuery("h3.post-title a").text();
        var name = jQuery("h3.post-title a").attr("href");
        var picture = jQuery("div.post-body p:nth-child(1) a img").attr("src");
        var description = "&description=" + jQuery(".post-body p:nth-child(2)").text();
        var href = "https://www.facebook.com/sharer/sharer.php?u=" + hostName + name + "\
        &src=sdkpreparse\
        &title=" + title + "\
        &link=" + hostName + name + "\
        &picture=" + hostName + picture + "\
        &description=" + description;
        jQuery(".fb-sharer").attr("href", href);
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
