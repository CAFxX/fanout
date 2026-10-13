package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int32 }
func genFn0(p0 uint8, p1 int8) uint32 {
    fmt.Println(p0, p1)
    p1 = (p1 | 30)
    switch 0 {
    case 32767:
        for fi4 := 0; fi4 < 17; fi4++ {
            var v5 int8 = 0
            fmt.Println(fi4, p1)
            v5 -= p1
            _ = fi4
            _ = v5
        }
    }
    fmt.Println(p0)
    switch 200 {
    case 128:
        p0 = p0
    case 335:
        var v6 uint16 = 65535
        _ = v6
    case 32768:
        var v7 float64 = 1e10
        _ = v7
    default:
        var v8 int32 = 1
        _ = v8
    }
    _ = p0
    _ = p1
    return 29
}
func genFn1(p0 bool) int {
    var v9 int = -850
    var v10 uint8 = 2
    v11 := []int{}
    if p0 {
        v11 = append(v11, v9)
        var v12 string = ("\n\t" + ("a" + "hello"))
        p0 = false
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return v9
}

func main() {
    v1 := []uint32{}
    var v2 float32 = ((1e10 / (((-2.25 - 1.5) / (float32(1.5) * float32(3.14159))) / (-2.25 - 3.14159))) + 0.1)
    var v3 int32 = 260
    fmt.Println(v3, v3)
    var v13 uint64 = 714
    var v14 int16 = 32767
    var v15 int32 = (v3 >> ((177) & 63))
    v1 = append(v1, 65535)
    v16 := GenS1{F0: v13, F1: (v14 ^ -1), F2: (v15 >> ((4294967295) & 63))}
    if true {
        fmt.Println(v13, v14, v3)
    }
    v1 = append(v1, genFn0(255, 0))
    var v17 int64 = 230
    v18 := GenS1{F0: v13, F1: v14, F2: v15}
    switch "foo bar" {
    case "":
        if (-23 < -1) {
            v19 := GenS1{F0: v13, F1: (v14 << ((1) & 63)), F2: v15}
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint16 = 345
                _ = fi20
                _ = v21
            }
            _ = v19
        } else {
            for fi22 := 0; fi22 < 15; fi22++ {
                fmt.Println(v3, v17)
                switch 32767 {
                case -580:
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 string = "a"
                        _ = len(v1)
                        v17 = (v17 + (v17 << ((944) & 63)))
                        _ = fi23
                        _ = v24
                    }
                }
                _ = fi22
            }
        }
    case "a":
        switch ("" + "") {
        case "foo bar":
            var v25 int64 = v17
            _ = v25
        default:
            switch 42 {
            case 4294967295:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -9223372036854775808:
                fmt.Println(v15, v3, v17)
            case 32768:
                fmt.Println(v13)
            default:
                fmt.Println(v15, v17)
            }
        }
    case "hello":
        fmt.Println(v3, v15, v15)
    }
    fmt.Println(v15, v17)
    var v26 uint = 2
    fmt.Println(v3)
    v27 := GenS1{F0: v13, F1: 32767, F2: (v3 * (v3 - v3))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
