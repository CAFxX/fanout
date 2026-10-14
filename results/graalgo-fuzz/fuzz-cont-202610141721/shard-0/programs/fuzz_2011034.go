package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int64; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint; F2 int32 }
func genFn0(p0 uint32, p1 int32) uint8 {
    v4 := []int{32767, -1, 820}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            switch (fi6 + (fi5 & 724)) {
            case 4294967295:
                if (!true) {
                    p1 += p1
                    p0 = 2147483648
                }
                fmt.Println(fi5, fi6)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = len(v4)
                if false {
                    p0 = 2147483647
                    var v7 int32 = (p1 << ((341) & 63))
                    p1 -= -1
                    _ = v7
                }
            case 42:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                var v8 uint8 = 2
                _ = v8
            }
            v9 := GenS1{F0: 255, F1: uint8(p0)}
            var v10 float64 = float64(1.5)
            _ = fi6
            _ = v10
            _ = v9
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            if false {
                switch -325 {
                case -15:
                    var v11 uint16 = 0
                    _ = v11
                case 1:
                    fmt.Println(p1, fi5, fi5)
                case 0:
                    fmt.Println(p1, p0, p0)
                default:
                    var v12 bool = true
                    _ = v12
                }
                p1 += (p1 << ((748) & 63))
                p1 += p1
            }
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 127
}

func main() {
    var v1 bool = (1000 < 0)
    var v2 uint32 = 1
    v3 := []uint32{v2, (v2 & (v2 - 42))}
    v13 := GenS1{F0: 834, F1: 2}
    v1 = ((-2.25 < ((3.14159 + float64((1.5 / 1.0))) * 1.5)) && v1)
    _ = len(v3)
    fmt.Println(v1, v2)
    if false {
        for fi14 := 0; fi14 < 14; fi14++ {
            v1 = (-128 >= 0)
            _ = fi14
        }
        if v1 {
            var v15 int64 = 1
            _ = v15
        }
        var v16 uint64 = 992
        _ = v16
    }
    var v17 int32 = 42
    var v18 float64 = ((0.0 - ((3.14159 - ((1e10 + (float64((0.0 + -2.25)) + 1e10)) * -2.25)) + -2.25)) - 0.1)
    v19 := GenS2{F0: 276}
    v1 = v1
    if v1 {
        v1 = (9223372036854775807 == 0)
    } else {
        for fi20 := 0; fi20 < 9; fi20++ {
            v2 += (^(v2))
            v21 := GenS3{F0: (!true), F1: 42, F2: v1}
            fmt.Println(v17)
            _ = fi20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        fmt.Println(v1, v1, fi22)
        var v23 bool = v1
        _ = fi22
        _ = v23
    }
    var v24 uint32 = v2
    _ = len(v3)
    v25 := GenS1{F0: 255, F1: 2}
    switch 625 {
    case 2147483647:
        v26 := GenS4{F0: 1}
        _ = v26
    case 42:
        switch "\n\t" {
        case "foo bar":
            var v27 uint64 = 2
            _ = v27
        }
    case 922:
        v28 := GenS1{F0: 255, F1: 0}
        _ = v28
    default:
        var v29 int64 = 829
        _ = v29
    }
    fmt.Println(v17, v1)
    if v1 {
        if (!true) {
            v1 = (!((!v1) && false))
        }
        v18 = float64(float32(((v18 - (v18 / 1.5)) + v18)))
    } else {
        if len(v3) > 0 {
            v3[0] = (v2 | (^(v24)))
        }
        if len(v3) > 0 {
            v3[0] = 65535
        }
        v30 := GenS5{F0: 42, F1: 65535, F2: (v17 - (v17 | v17))}
        _ = v30
    }
    var v31 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
}
