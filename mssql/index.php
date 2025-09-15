<?php

$con = new sqlsrv_connect('127.0.0.1', 'sa', 'vira16011999!', 'sqlsrv');

if ($con) {
    echo 'Connected';
} else {
    echo 'Not Connected';
}