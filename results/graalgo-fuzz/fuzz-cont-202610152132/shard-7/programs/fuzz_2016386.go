package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint64, p1 uint64) int32 {
    v4 := GenS1{F0: -453, F1: 255, F2: 1}
    switch "foo bar" {
    case "foo bar":
        var v5 uint16 = 453
        _ = v5
    case "hello":
        v6 := GenS1{F0: 734, F1: 55, F2: 2147483647}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}
func genFn1() bool {
    var v7 int32 = -2147483648
    v8 := GenS1{F0: (-(v7)), F1: 1, F2: 2147483647}
    v7 = 1
    var v9 bool = (1 < -233)
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 int8 = -22
    var v2 int64 = 42
    var v3 int = 32767
    v10 := []int{(v3 / (((v3 >> ((964) & 63))) | 1))}
    v10 = append(v10, 955)
    var v11 int16 = 1
    fmt.Println(v2)
    v12 := GenS2{F0: v11}
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            v2 -= (v2 >> ((0) & 63))
            if ((-2147483648 != -983) && false) {
                var v15 uint64 = 670
                _ = v15
            } else {
                if true {
                    if (1.5 <= -2.25) {
                        var v16 int = (v3 >> ((478) & 63))
                        v10 = append(v10, v3)
                        _ = v16
                    } else {
                        var v17 uint16 = 32767
                        var v18 uint16 = v17
                        if (v2 < (v2 - (v2 << ((65535) & 63)))) {
                            _ = len(v10)
                            fmt.Println(v18, fi13)
                            var v19 uint = 1
                            var v20 int = fi14
                            _ = v19
                            _ = v20
                        }
                        _ = v17
                        _ = v18
                    }
                    v3 -= (fi13 - v3)
                }
                var v21 int32 = 1
                fmt.Println(v2)
                _ = v21
            }
            v1 = 109
            _ = fi14
        }
        _ = fi13
    }
    v2 -= (^(v2))
    var v22 uint16 = 65535
    fmt.Println(v11, v11)
    v23 := GenS2{F0: (v11 % ((-1) | 1))}
    var v24 uint32 = 2147483648
    v25 := GenS1{F0: 2147483647, F1: 2, F2: 42}
    for fi26 := 0; fi26 < 12; fi26++ {
        v3 -= (fi26 / (((v3 % (((-(fi26))) | 1))) | 1))
        v27 := GenS3{F0: v24, F1: -869}
        _ = fi26
        _ = v27
    }
    switch "a" {
    case "hello":
        v28 := GenS1{F0: genFn0(351, 1), F1: 255, F2: v24}
        _ = v28
    }
    switch 720 {
    case 733:
        var v29 uint = 835
        _ = v29
    default:
        switch "a" {
        case "a":
            v30 := GenS3{F0: (v24 % (((v24 | (^(v24)))) | 1)), F1: -1}
            _ = v30
        }
    }
    var v31 string = "foo bar"
    var v32 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
}
