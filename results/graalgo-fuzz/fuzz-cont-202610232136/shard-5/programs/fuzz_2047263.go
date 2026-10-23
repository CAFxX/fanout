package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 int64, p1 uint32, p2 uint16) bool {
    var v4 int = 956
    var v5 uint = 933
    v6 := []int{}
    if (!(uint(p1) == v5)) {
        fmt.Println(v5, v4, p0)
        var v7 int = v4
        _ = v7
    }
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (v5 <= v5)
}
func genFn1(p0 int8, p1 bool) int16 {
    v8 := GenS1{F0: 2147483647, F1: 127}
    v9 := GenS1{F0: 42, F1: 174}
    switch -546 {
    case -1:
        var v10 int16 = -32768
        _ = v10
    case 9223372036854775807:
        p0 -= (p0 + p0)
    case 2147483647:
        fmt.Println(p1)
    default:
        v11 := GenS1{F0: 42, F1: 128}
        _ = v11
    }
    var v12 uint16 = 32767
    _ = p0
    _ = p1
    _ = v12
    _ = v8
    _ = v9
    return 484
}

func main() {
    var v1 int32 = 0
    var v2 uint8 = 0
    var v3 float64 = (-2.25 * 1e10)
    var v13 bool = false
    var v14 uint8 = 141
    v15 := []int{24}
    var v16 string = ("hello" + "")
    v13 = (0.0 >= (v3 / v3))
    var v17 uint8 = (^(v14))
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            switch 372 {
            case 42:
                var v20 int32 = (v1 | v1)
                _ = v20
            case -1:
                switch v16 {
                case "\n\t":
                    v2 = (v14 << ((213) & 63))
                default:
                    var v21 float32 = 1.5
                    _ = v21
                }
            case 2147483647:
                v13 = false
            }
            v1 -= (-(v1))
            v1 = v1
            _ = fi19
        }
        var v22 uint16 = 446
        _ = fi18
        _ = v22
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v14 -= (v2 & v14)
    var v23 int64 = 42
    v1 = v1
    v23 += 955
    v2 -= v2
    var v24 uint16 = 492
    v2 -= v17
    var v25 bool = (v13 || false)
    v16 = (((((v16 + v16) + "a") + v16) + (v16 + v16)) + v16)
    v15 = append(v15, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
