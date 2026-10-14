package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 float32 = (0.0 * (1.5 + (0.1 + 0.0)))
    var v2 float64 = 3.14159
    var v3 string = ("hello" + "")
    v4 := []int{128, -456, 42}
    for fi5 := 0; fi5 < 18; fi5++ {
        switch "foo bar" {
        case "hello":
            if len(v4) > 0 {
                v4[0] = 32767
            }
        case "a":
            v6 := GenS1{F0: 1, F1: 2147483647}
            _ = v6
        case "":
            v7 := GenS2{F0: 42}
            _ = v7
        default:
            var v8 uint8 = 255
            _ = v8
        }
        switch (fi5 + (fi5 << ((797) & 63))) {
        case 128:
            var v9 uint = 1
            _ = v9
        case -1:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi5
    }
    if (35 != 0) {
        if len(v4) > 0 {
            v4[0] = 42
        }
        var v10 int64 = -9223372036854775808
        _ = v10
    }
    var v11 bool = (!true)
    v2 = float64((v1 + 1e10))
    if len(v4) > 0 {
        v4[0] = 1
    }
    var v12 uint = 18446744073709551615
    if false {
        for fi13 := 0; fi13 < 14; fi13++ {
            v1 -= 1e10
            _ = fi13
        }
        var v14 float32 = 0.1
        _ = v14
    } else {
        if len(v4) > 0 {
            v4[0] = 128
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
