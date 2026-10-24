package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    var v2 string = "hello"
    var v3 int32 = 538
    var v4 int16 = -32768
    var v5 int32 = v3
    var v6 int = 473
    delete(v1, -2147483648)
    v3 = -1
    v3 = (^(v5))
    var v7 uint16 = 954
    switch ("hello" + v2) {
    case "a":
        v8 := []int{(v6 % (((v6 * (v6 >> ((1) & 63)))) | 1))}
        _ = v8
    case "foo bar":
        fmt.Println(v2, v2)
    }
    v1[(-(v6))] = v2
    v9 := GenS1{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
