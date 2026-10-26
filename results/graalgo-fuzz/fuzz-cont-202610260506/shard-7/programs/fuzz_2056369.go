package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int; F1 bool }
func genFn0() bool {
    v4 := GenS1{F0: -915, F1: 66}
    if ((!false) && (!(255 >= 872))) {
        var v5 bool = false
        _ = v5
    }
    _ = v4
    return (!(!true))
}

func main() {
    var v1 uint = 127
    var v2 uint32 = 65535
    v3 := []uint32{}
    fmt.Println(v2, v1)
    switch "\n\t" {
    case "a":
        if false {
            if len(v3) > 0 {
                v3[0] = v2
            }
            _ = len(v3)
            v3 = append(v3, v2)
        }
    }
    if (!false) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v6 := GenS2{F0: 32768, F1: true}
        _ = len(v3)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            _ = len(v3)
            _ = fi7
        }
        v1 = v1
        _ = len(v3)
    }
    var v8 float32 = 3.14159
    var v9 uint16 = 735
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v10 uint32 = (v2 + 147)
    v11 := GenS1{F0: 2147483647, F1: 758}
    var v12 float32 = 0.0
    _ = len(v3)
    v2 = 0
    switch 128 {
    case 9223372036854775807:
        var v13 uint16 = 65535
        _ = v13
    }
    var v14 float32 = (float32(1e10) / (3.14159 / 0.1))
    v12 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
