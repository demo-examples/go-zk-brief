<?php

$post_string = array('key'=>'1122-3434', 'destName' => 'hello.test.unit', 'zkidc'=>'qa');
#$url = '192.168.35.141:9090/servicelist';
$url = '127.0.0.1:9090/delservice';


$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $url);
curl_setopt($ch, CURLOPT_POSTFIELDS, $post_string);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 10);
curl_setopt($ch, CURLOPT_TIMEOUT, 30);
$result = curl_exec($ch);  
var_dump($result);
curl_close($ch);
