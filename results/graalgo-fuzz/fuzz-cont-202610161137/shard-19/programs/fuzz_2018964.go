package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 uint8 = 255
    delete(v1, "hello")
    var v4 int64 = 716
    switch "hello" {
    case "foo bar":
        v5 := v1["foo bar"]
        _ = v5
    case "\n\t":
        if false {
            var v6 uint16 = 42
            _ = v6
        } else {
            v4 -= v4
        }
    default:
        var v7 string = "hello"
        _ = v7
    }
    var v8 int8 = -128
    v3 += 10
    v1["a"] = "a"
    v9 := GenS1{F0: 4294967295, F1: v4, F2: 637}
    fmt.Println(v8, v4, v4)
    v3 -= v3
    var v10 int16 = 1000
    v11 := GenS1{F0: 24, F1: v4, F2: v4}
    for fi12 := 0; fi12 < 10; fi12++ {
        switch "hello" {
        case "":
            v4 -= 9223372036854775807
        case "hello":
            var v13 float32 = 3.14159
            _ = v13
        case "\n\t":
            var v14 uint16 = 377
            _ = v14
        default:
            v15 := GenS1{F0: 643, F1: (v4 % (((^(v4))) | 1)), F2: (-(v4))}
            _ = v15
        }
        var v16 int = fi12
        _ = fi12
        _ = v16
    }
    switch -851 {
    case 128:
        if (!true) {
            if len(v2) > 0 {
                v2[0] = (("" + "") + "")
            }
            v3 = v3
        }
    default:
        v3 = v3
    }
    v8 = 1
    fmt.Println(v8, v8, v4)
    var v17 int64 = (-(v4))
    var v18 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
