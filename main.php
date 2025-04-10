<?php

class Main
{

    public function add(int $x, int $y): int
    {
        return ($x + $y);
    }
    public function main()
    {
        $message = "Hello, World!";

        $num1 = 10;

        $num2 = 5;

        $sum = $this->add($num1, $num2);

        $product = ($num1 * $num2);

        echo $message . "\n" ;

        echo "Sum:" . $sum . "\n" ;

        echo "Product:" . $product . "\n" ;

    }
}
$main = new Main();
$main->main();
