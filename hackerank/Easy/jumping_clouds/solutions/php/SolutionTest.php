<?php declare(strict_types=1);
use PHPUnit\Framework\TestCase;
require_once('solution.php');

final class SolutionTest extends TestCase
{
    public function testJumpingClouds(): void
    {
        $stack1 = [0, 0, 1, 0, 0, 1, 0];
        $stack2 = [0, 0, 0, 1, 0, 0];
        $this->assertSame(4, jumpingOnClouds($stack1));
        $this->assertSame(3, jumpingOnClouds($stack2));
    }
}