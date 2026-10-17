package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []uint32{2, 0}
    v2 := []string{"\n\t", ("\n\t" + ("" + ("hello" + "")))}
    v3 := []string{}
    var v4 uint16 = 480
    var v5 int = 376
    var v6 uint32 = 2147483647
    v6 -= v6
    v4 = v4
    switch "a" {
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        fmt.Println(v6)
    default:
        if len(v3) > 0 {
            v3[0] = "a"
        }
    }
    v6 -= v6
    if (false || false) {
        switch ((("hello" + "\n\t") + "foo bar") + "foo bar") {
        case "a":
            v4 = 255
        case "foo bar":
            var v7 uint64 = 541
            _ = v7
        case "hello":
            v5 = (v5 & v5)
        }
    } else {
        v8 := GenS1{F0: 42}
        v4 = 340
        _ = v8
    }
    fmt.Println(v5)
    var v9 string = ("hello" + "hello")
    var v10 float64 = 0.1
    v10 = v10
    var v11 int64 = 1
    v10 = (float64(1e10) / v10)
    if true {
        var v12 int32 = -2147483648
        var v13 bool = ((^(v6)) != v6)
        v10 = (0.0 * v10)
        _ = v12
        _ = v13
    } else {
        fmt.Println(v10, v10, v4)
        if (1 <= 295) {
            var v14 float64 = (v10 - ((-2.25 / (1e10 + v10)) - v10))
            _ = v14
        }
        v4 = (^(v4))
    }
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
