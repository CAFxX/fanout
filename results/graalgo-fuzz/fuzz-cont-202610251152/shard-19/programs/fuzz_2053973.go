package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }

func main() {
    v1 := []int{2147483647}
    v2 := make(map[uint32]bool)
    var v3 string = ""
    fmt.Println(v3)
    v2[1] = true
    v3 = ("a" + (v3 + ""))
    var v4 bool = true
    switch ("hello" + v3) {
    case "\n\t":
        _ = len(v2)
    }
    var v5 bool = false
    v5 = (1 >= 9223372036854775807)
    var v6 float64 = 0.1
    fmt.Println(v4, v6, v3)
    v3 = v3
    switch "" {
    case "":
        var v7 bool = (178 > 2)
        _ = v7
    }
    var v8 int8 = 127
    v8 += (v8 << ((65535) & 63))
    switch v3 {
    case "\n\t":
        switch v3 {
        case "\n\t":
            var v9 int64 = 42
            _ = v9
        default:
            v10 := v2[2147483647]
            _ = v10
        }
    case "a":
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(v5)
            v3 = ((v3 + "foo bar") + ((v3 + v3) + v3))
            v12 := v2[1]
            _ = fi11
            _ = v12
        }
    default:
        fmt.Println(v4, v6)
    }
    v5 = v5
    fmt.Println(v6)
    switch "a" {
    case "":
        var v13 bool = (((v3 + ("\n\t" + v3)) + "hello") >= (v3 + v3))
        _ = v13
    case "foo bar":
        var v14 uint = 735
        _ = v14
    case "hello":
        v15 := GenS1{F0: -392, F1: 0}
        _ = v15
    default:
        var v16 int16 = 1
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v4 = (v4 && (true || (!v5)))
        _ = fi17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
