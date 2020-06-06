package main

var page string = `<!DOCTYPE html>
<html>
<!--here's an html comment -->
<head>
    <script>
        // here's a one liner outside of a function 
        $(document).ready(function () { // here's one at the end
            $('body').bootstrapMaterialDesign();

            $('#mode').change(function () {
                var mode = $('#mode').val();
                console.log(mode);
                if ("input" == mode) {
                    $('#valueG').show();
                } else if ("output" == mode) {
                    $('#valueG').hide();
                }
            });

            $("form").submit(function (event) {
                /**
                 * big comment
                 * big big comment
                 **/
                event.preventDefault();
                var actionFile = $(this).attr("action");
                var formValues = $(this).serialize();
                var req = formValues;
                console.log(req);
                $("#response").html(req);
                $.get(actionFile, req, function (data) {
                    $("#response").html(data);
                });
            });
        });
    </script>
</head>
<!--here's an html comment between the body and head ends -->

</body>

</html>
`
