
<?php

/* @var $this yii\web\View */

$this->title = 'My Yii Application';
 $url         = \yii\helpers\Url::to( [ 'site/api' ] );


$script = <<< JS
var foo = [];

for (var i = 1; i <= 25; i++) {
   foo.push(i);
}
var ctx = document.getElementById('myChart').getContext('2d');

 var chart = new Chart(ctx, {
        type: 'line',
        data: {},
        options: {
            title: {
                display: true,
                text: 'World population per region (in millions)'
            }
        }
    });
console.log('chart is ready');
var  dataset;
  dataset = {
            labels: foo,
             datasets: [
                 {
                data: {},
                label: "Pressure",
                borderColor: "#3e95cd",
                fill: false
            }, 
            {
                data: {},
                label: "Humidity",
                borderColor: "#8e5ea2",
                fill: false
            },
            //  {
            //     data: dataa['TemperatureR'],
            //     label: "TemperatureR",
            //     borderColor: "#3cba9f",
            //     fill: false
            // },
            //  {
            //     data: dataa['TemperatureA'],
            //     label: "TemperatureA",
            //     borderColor: "#e8c3b9",
            //     fill: false
            // }, {
            //     data: dataa['pH'],
            //     label: "pH",
            //     borderColor: "#c45850",
            //     fill: false
            // }
            ]
        };
   chart.data = dataset;
 
   function the_func(response) {
       console.log('ajax send wait ....');
     dataa =  JSON.parse(response);
      
   
    
    
          
        chart.data.datasets[0].data = dataa['Pressure'];
        chart.data.datasets[1].data = dataa['Humidity'];
    console.log(chart.data.datasets);
      chart.update();
      setTimeout(worker(the_func), 5000);
  }
 
function worker(the_func) {
       
    $.ajax({
        type: "POST",
        // dataType: 'json',
       // async: false,

        url: "$url",
        contentType: "application/json",
        success:  the_func,
        error: function (data) {
            console.log(data);
        },
    });
  
   // setTimeout(worker(the_func), 30000);
    }
    worker(the_func);
  
    
JS;
$this->registerJs( $script );


?>

<div class="site-index">
    <div class="row">
        <div class="col-md-12">
            <canvas id="myChart" width="800" height="400"></canvas>
        </div>

    </div>
</div>
<script>


</script>