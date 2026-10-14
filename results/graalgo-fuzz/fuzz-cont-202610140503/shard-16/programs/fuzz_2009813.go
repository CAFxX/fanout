package main

import (
    "fmt"
)


func main() {
    var v1 int32 = -1
    var v2 float64 = 0.0
    var v3 float32 = 0.0
    v4 := []int{}
    fmt.Println(v3)
    v1 = 2147483647
    v3 = v3
    switch ((("" + ("foo bar" + ("a" + "a"))) + ("\n\t" + "")) + "") {
    case "":
        v3 = (v3 / float32(v2))
    case "foo bar":
        var v5 uint8 = 255
        _ = v5
    default:
        fmt.Println(v3)
    }
    var v6 uint16 = 65535
    if (-1 < -12) {
        v2 += v2
    } else {
        var v7 int64 = 44
        var v8 int32 = -2147483648
        var v9 uint32 = 1
        _ = v7
        _ = v8
        _ = v9
    }
    v4 = append(v4, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
