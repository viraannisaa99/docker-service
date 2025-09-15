<?php

$con = new mysqli('mysql_db', 'root', 'root', 'mysql');

if ($con) {
    echo 'Connected';
} else {
    echo 'Not Connected';
}