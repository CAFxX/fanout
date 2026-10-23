package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int; F1 int16; F2 uint }

func main() {
    v1 := []int{42, 1}
    var v2 uint = 18446744073709551615
    var v3 uint32 = 4294967295
    var v4 string = ("hello" + "foo bar")
    var v5 uint64 = 567
    v2 = (v2 + v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch (v4 + v4) {
    case "foo bar":
        v2 = v2
    case "hello":
        v3 = v3
    case "a":
        var v6 int32 = 0
        _ = v6
    }
    var v7 uint = (v2 | uint(v5))
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 string = "hello"
        var v10 float64 = (-2.25 - float64((float32(-2.25) - 1e10)))
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        v11 := GenS1{F0: int8(v3), F1: 39, F2: 0}
        var v12 int64 = 1
        var v13 float64 = 0.1
        _ = v11
        _ = v12
        _ = v13
    } else {
        v1 = append(v1, int(v5))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v3, v4)
    }
    v14 := GenS1{F0: -1, F1: 128, F2: 1}
    if (v5 > (^(v5))) {
        switch "\n\t" {
        case "a":
            v4 = v4
        case "":
            var v15 uint = (v7 - v7)
            _ = v15
        default:
            switch -413 {
            case -550:
                var v16 int32 = -2147483648
                _ = v16
            case 32768:
                v2 -= 769
            case -222:
                var v17 bool = true
                _ = v17
            default:
                var v18 string = v4
                _ = v18
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        _ = len(v1)
    }
    if false {
        var v19 int32 = -2147483648
        switch 2147483647 {
        case 993:
            if (v19 == -455) {
                _ = len(v1)
                v3 += (v3 >> (((v2 / ((uint(v5)) | 1))) & 63))
            }
        case 32768:
            v5 = v5
        }
        _ = v19
    }
    v3 = (^(v3))
    if false {
        var v20 uint64 = v5
        v21 := GenS2{F0: 127, F1: 679, F2: (v7 * v7)}
        _ = v20
        _ = v21
    }
    var v22 string = ("hello" + v4)
    v23 := GenS2{F0: 32768, F1: -1, F2: v7}
    v7 = (^(v7))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
