package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
func genFn0(p0 int64, p1 uint) int64 {
    var v5 int = -883
    p0 = (p0 | (^(p0)))
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 string = "foo bar"
            _ = fi7
            _ = v8
        }
        var v9 uint64 = 315
        fmt.Println(v5, fi6, v9)
        _ = fi6
        _ = v9
    }
    v10 := GenS1{F0: 502, F1: true}
    p0 = p0
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return (p0 - p0)
}

func main() {
    var v1 float32 = 0.0
    var v2 bool = true
    v3 := make(map[int]int)
    v4 := []uint32{0, 875}
    v1 += float32(((0.0 * 0.0) / 1.0))
    var v11 bool = (v2 || ((!(0 == 139)) && (0 == 353)))
    switch "\n\t" {
    case "foo bar":
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int64 = 1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi12
            _ = v13
        }
    case "":
        v2 = false
    case "\n\t":
        var v14 uint = 255
        _ = v14
    }
    fmt.Println(v11, v1)
    v1 -= 1.5
    var v15 uint8 = 1
    var v16 int64 = 42
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 uint = 4294967295
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 2, F1: false}
    v4 = append(v4, 0)
    v15 = (v15 % (((v15 * 0)) | 1))
    v11 = (!v11)
    v20 := GenS1{F0: 18446744073709551615, F1: (true || (0 != -128))}
    var v21 uint8 = v15
    v22 := GenS1{F0: 2, F1: v11}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
