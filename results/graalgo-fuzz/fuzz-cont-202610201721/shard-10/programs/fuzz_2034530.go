package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int8 = 0
    var v2 bool = (!(!(758 <= 4294967295)))
    var v3 int32 = -575
    var v4 uint16 = 893
    v5 := []uint32{2, 1}
    var v6 uint16 = v4
    var v7 int32 = v3
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            v1 += int8(v7)
            v10 := GenS1{F0: 33, F1: 2147483648}
            v11 := GenS1{F0: 127, F1: 65535}
            _ = fi9
            _ = v10
            _ = v11
        }
        v1 += (v1 * (v1 | (^(v1))))
        _ = fi8
    }
    var v12 float64 = float64(1.5)
    var v13 uint16 = 389
    var v14 int = -74
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v15 uint = 42
    var v16 int = v14
    var v17 int16 = -1
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            v5 = append(v5, 605)
            var v20 int16 = (v17 / ((496) | 1))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v21 float32 = (0.0 / (float32(v12) / (-2.25 - (-2.25 * ((1.5 / -2.25) * ((((3.14159 * (3.14159 * (1e10 * 1.5))) + -2.25) / 1.5) + 0.1))))))
            v17 = v17
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    fmt.Println(v2, v7)
    for fi22 := 0; fi22 < 4; fi22++ {
        if v2 {
            v16 = fi22
            v14 -= (-(fi22))
        } else {
            switch "a" {
            case "hello":
                v15 -= 2
            case "a":
                var v23 uint16 = v4
                _ = v23
            case "":
                fmt.Println(v14)
            default:
                for fi24 := 0; fi24 < 11; fi24++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v3 = 0
                    _ = fi24
                }
            }
            if len(v5) > 0 {
                v5[0] = 428
            }
        }
        fmt.Println(v4, v15)
        _ = fi22
    }
    v1 -= v1
    v25 := GenS2{F0: 1}
    var v26 bool = (v2 && v2)
    fmt.Println(v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
