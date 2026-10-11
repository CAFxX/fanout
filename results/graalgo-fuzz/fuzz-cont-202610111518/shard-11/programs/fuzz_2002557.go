package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint64, p2 int16) uint32 {
    if true {
        var v4 uint64 = 555
        _ = v4
    } else {
        fmt.Println(p2, p2)
        var v5 uint64 = p1
        _ = v5
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1() uint8 {
    var v6 int16 = 1000
    v6 -= -544
    _ = v6
    return 128
}

func main() {
    v1 := []int{212, 32768, 792}
    var v2 string = ("foo bar" + "a")
    var v3 float64 = 1.5
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: -47}
        _ = fi7
        _ = v8
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = 105
    }
    fmt.Println(v3)
    var v9 int32 = -990
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 uint = 4294967295
    v11 := GenS2{F0: v9, F1: 0, F2: 2147483647}
    switch v2 {
    case "foo bar":
        v12 := GenS1{F0: 127}
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS3{F0: 95}
        for fi15 := 0; fi15 < 16; fi15++ {
            v2 = "a"
            _ = len(v1)
            fmt.Println(v3, v2)
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
