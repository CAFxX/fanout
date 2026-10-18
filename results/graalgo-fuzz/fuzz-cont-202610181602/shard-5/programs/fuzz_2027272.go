package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }

func main() {
    var v1 uint = 813
    var v2 int8 = 127
    v3 := []string{"hello", "a", "", (("a" + ("\n\t" + "foo bar")) + ("hello" + ""))}
    var v4 float64 = 0.0
    var v5 uint16 = 255
    fmt.Println(v1, v4, v1)
    v5 += (v5 / ((2) | 1))
    v5 = v5
    v6 := GenS1{F0: 1, F1: -884, F2: 2}
    var v7 int16 = -1
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint32 = 2147483647
        v1 = 18446744073709551615
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int = (fi10 - fi10)
        _ = fi10
        _ = v11
    }
    var v12 int64 = 1
    for fi13 := 0; fi13 < 10; fi13++ {
        v1 = v1
        _ = fi13
    }
    v5 = (v5 | v5)
    if (v4 != float64(0.1)) {
        if false {
            var v14 uint32 = 940
            _ = v14
        } else {
            if false {
                v7 += (v7 | v7)
            }
        }
        fmt.Println(v7, v4, v5)
        v5 = (v5 >> ((v1) & 63))
    }
    switch -1 {
    case -9223372036854775808:
        switch "\n\t" {
        case "a":
            v2 += (v2 - v2)
        }
    default:
        v12 = (v12 & -9223372036854775808)
    }
    switch ("\n\t" + "hello") {
    case "a":
        v15 := GenS1{F0: 468, F1: (-(v12)), F2: (v5 + v5)}
        _ = v15
    case "":
        var v16 int8 = (v2 + v2)
        _ = v16
    }
    fmt.Println(v4, v2, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
