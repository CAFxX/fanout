package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int16; F1 int }
func genFn0(p0 uint32) int16 {
    fmt.Println(p0)
    var v4 uint = 0
    switch 1 {
    case -1:
        p0 = p0
    case 32768:
        var v5 uint32 = (^(p0))
        _ = v5
    default:
        if true {
            var v6 uint16 = 32767
            _ = v6
        } else {
            v7 := GenS1{F0: 9223372036854775807}
            v4 = v4
            var v8 bool = true
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = v4
    return 1
}
func genFn1() uint8 {
    var v9 int = 128
    v10 := []int{(v9 >> ((161) & 63)), (^(v9)), v9}
    _ = v10
    _ = v9
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -989
    v3 := []int{2147483647, v2, 586, 42}
    var v11 int16 = genFn0(1)
    var v12 uint = 8
    var v13 uint32 = 2147483647
    v13 = 2147483647
    v11 = 42
    v2 = 372
    var v14 uint = (v12 % ((528) | 1))
    v15 := GenS2{F0: 17, F1: (v13 >> ((v12) & 63))}
    v16 := GenS3{F0: genFn0((v13 | v13)), F1: v2}
    v11 += v11
    for fi17 := 0; fi17 < 1; fi17++ {
        v2 = 32767
        if ((42 >= 2) || false) {
            v14 -= (^(v14))
            var v18 float64 = 0.1
            _ = v18
        } else {
            fmt.Println(fi17)
        }
        _ = fi17
    }
    if len(v3) > 0 {
        v3[0] = 127
    }
    fmt.Println(v12, v13, v2)
    var v19 uint16 = 641
    v13 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
