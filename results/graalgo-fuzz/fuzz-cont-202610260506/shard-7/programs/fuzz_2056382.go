package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int32) uint32 {
    v7 := []int{4294967295}
    if (-36 != -89) {
        var v8 uint64 = 631
        v9 := GenS1{F0: 65535}
        _ = v8
        _ = v9
    } else {
        p0 = 812
    }
    _ = p0
    _ = v7
    return 299
}
func genFn1(p0 uint8, p1 uint16, p2 int64) int8 {
    fmt.Println(p2, p1)
    v10 := GenS1{F0: 2}
    for fi11 := 0; fi11 < 11; fi11++ {
        if true {
            p2 -= 0
        }
        p1 = (^(p1))
        _ = fi11
    }
    p2 = (p2 | (p2 % (((p2 >> ((816) & 63))) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = -531
    var v3 int = 515
    v4 := make(map[uint32]uint32)
    var v5 int = v3
    v6 := make(map[uint32]int)
    v12 := []int{}
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int8 = 59
        _ = fi13
        _ = v14
    }
    var v15 float64 = 1.5
    v1 = (v1 ^ v1)
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int32 = 718
            switch "hello" {
            case "foo bar":
                v18 += 120
            default:
                var v19 string = "foo bar"
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        var v20 float32 = 1.5
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi16
        _ = v20
    }
    var v21 float64 = ((v15 / 3.14159) + 1.5)
    v22 := GenS1{F0: 792}
    var v23 string = ("foo bar" + "")
    var v24 bool = true
    var v25 int64 = (v1 + (v1 - (v1 << ((855) & 63))))
    var v26 uint32 = genFn0(-1)
    var v27 int16 = 1000
    v27 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
