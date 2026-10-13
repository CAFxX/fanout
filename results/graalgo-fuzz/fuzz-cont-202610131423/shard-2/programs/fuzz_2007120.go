package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }
func genFn0(p0 int64, p1 int32, p2 bool) int16 {
    var v4 float64 = -2.25
    if (792 != 4294967295) {
        v5 := []int{127}
        var v6 uint = 18446744073709551615
        _ = v5
        _ = v6
    } else {
        var v7 int32 = 1
        _ = v7
    }
    fmt.Println(v4, v4, p0)
    v8 := []int{-1, 599}
    if (!(!(238 == 216))) {
        var v9 uint32 = 2
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return 32767
}
func genFn1(p0 bool) uint64 {
    var v10 float32 = 0.1
    if false {
        var v11 uint64 = 1
        var v12 uint32 = uint32(v11)
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v10
    return 352
}

func main() {
    var v1 int16 = 121
    var v2 int64 = 756
    var v3 int16 = (v1 ^ v1)
    v3 = v1
    if true {
        var v13 uint16 = 2
        _ = v13
    } else {
        fmt.Println(v2, v1, v1)
        var v14 float64 = (0.0 * (float64(float32(0.0)) + -2.25))
        v14 = v14
        _ = v14
    }
    var v15 int64 = -1
    for fi16 := 0; fi16 < 20; fi16++ {
        switch fi16 {
        case -2147483648:
            var v17 bool = (!(255 > 474))
            _ = v17
        case -120:
            switch 127 {
            case -777:
                if true {
                    switch "foo bar" {
                    case "":
                        var v18 int32 = 2147483647
                        _ = v18
                    default:
                        v15 = (v15 >> ((65535) & 63))
                    }
                    var v19 int = fi16
                    _ = v19
                }
            case -900:
                v20 := GenS1{F0: 13, F1: 1, F2: 342}
                _ = v20
            case -62:
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := []int{(-(fi16))}
                    _ = fi21
                    _ = v22
                }
            }
        case 128:
            v1 = (^(v3))
        default:
            var v23 int = fi16
            _ = v23
        }
        _ = fi16
    }
    v24 := []int{32767, int(v1)}
    fmt.Println(v3, v1)
    if true {
        var v25 uint64 = genFn1(true)
        v3 -= -697
        _ = v25
    } else {
        for fi26 := 0; fi26 < 4; fi26++ {
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            switch fi26 {
            case -51:
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            }
            _ = fi26
        }
    }
    fmt.Println(v1)
    if true {
        v15 = 1
        var v27 int32 = -719
        var v28 int64 = v2
        _ = v27
        _ = v28
    } else {
        if len(v24) > 0 {
            v24[0] = 128
        }
        fmt.Println(v1, v15, v15)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
}
