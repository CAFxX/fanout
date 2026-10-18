package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]string)
    var v3 int = -790
    switch (v3 >> ((255) & 63)) {
    case -192:
        var v4 uint16 = 32767
        _ = v4
    case 954:
        v3 = v3
    case 127:
        var v5 uint32 = 2147483647
        _ = v5
    default:
        fmt.Println(v3, v3)
    }
    var v6 string = "a"
    if (32767 != 65535) {
        var v7 float64 = float64(0.0)
        _ = v7
    }
    fmt.Println(v1, v6)
    switch "foo bar" {
    case "\n\t":
        var v8 int64 = 42
        _ = v8
    case "":
        v9 := []int{(v3 - (v3 % ((v3) | 1))), v3}
        _ = v9
    default:
        v10 := GenS1{F0: v1}
        _ = v10
    }
    fmt.Println(v1)
    if true {
        if false {
            for fi11 := 0; fi11 < 20; fi11++ {
                if (false && true) {
                    v1 = v1
                } else {
                    var v12 bool = ((fi11 / ((fi11) | 1)) != fi11)
                    _ = v12
                }
                _ = fi11
            }
        }
    } else {
        v1 += (-(v1))
    }
    var v13 uint32 = 65535
    v14 := []int{(v3 + (v3 & 1)), v3, (^(v3))}
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
}
