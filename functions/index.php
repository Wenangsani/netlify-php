<?php
   require 'vendor/autoload.php';

use Aws\S3\S3Client;

// Instantiate an Amazon S3 client.
$s3 = new S3Client([
    'version' => 'latest',
    'region'  => 'us-west-2'
]);

print_r($s3);
?>
