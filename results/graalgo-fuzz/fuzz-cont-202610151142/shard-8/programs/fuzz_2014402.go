package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint = 255
    var v2 bool = (3.14159 != (-2.25 + 0.0))
    var v3 int = int(v1)
    v4 := []uint32{2147483648}
    v1 -= 42
    fmt.Println(v3, v1)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: -890}
        _ = fi5
        _ = v6
    }
    fmt.Println(v2, v3, v2)
    switch "" {
    case "hello":
        var v7 uint16 = 0
        _ = v7
    case "a":
        v3 += (v3 % (((^(v3))) | 1))
    }
    var v8 int8 = 0
    _ = len(v4)
    var v9 int16 = 1
    v3 = (v3 % (((v3 * (v3 % (((v3 / ((v3) | 1))) | 1)))) | 1))
    v10 := GenS1{F0: v9}
    v2 = v2
    var v11 int16 = v9
    switch 650 {
    case 32768:
        v1 += v1
    case -967:
        var v12 uint64 = 2
        _ = v12
    case -9223372036854775808:
        fmt.Println(v9, v3)
    }
    if (255 >= 599) {
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS1{F0: (^(v9))}
            _ = fi13
            _ = v14
        }
    } else {
        var v15 uint16 = 1
        v16 := GenS1{F0: (-(v9))}
        _ = v15
        _ = v16
    }
    var v17 int8 = -1
    v1 = (v1 & (v1 * v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
