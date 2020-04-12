#!/usr/bin/env bats

@test "test error no argument" {
    run bash -c "echo $? | ./groundhog"
    [ "$status" -eq 84 ]
}

@test "test error not enough number" {
    run bash -c "echo $? | ./groundhog 1"
    [ "$status" -eq 84 ]
}

@test "test STOP 0 number" {
    run bash -c "echo STOP | ./groundhog 3"
    [ "$status" -eq 84 ]
}

@test "test error arg" {
    run bash -c "echo $? | cat filetest/fileerror.txt | ./groundhog 3"
    [ "$status" -eq 84 ]
}

@test "test file basic" {
    run bash -c "cat filetest/filepdf.txt | ./groundhog 7 > .Bastest.txt"
    run  bash -c "diff filetest/filepdf_test.txt .Bastest.txt"
    [ "$status" -eq 0 ]
    run  bash -c "rm .Bastest.txt"
}

@test "test file advance" {
    run bash -c "cat filetest/file.txt | ./groundhog 3 > .Advtest.txt"
    run  bash -c "diff filetest/file_test.txt .Advtest.txt"
    [ "$status" -eq 0 ]
    run  bash -c "rm .Advtest.txt"
}