package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 uint8 }
func genFn0() string {
    var v5 bool = false
    switch 127 {
    case 2147483647:
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := []int{(fi6 << ((182) & 63)), fi6, fi6, 32767}
            fmt.Println(fi6, fi6)
            fmt.Println(v5, fi6)
            _ = fi6
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 12; fi8++ {
            switch fi8 {
            case -2147483648:
                for fi9 := 0; fi9 < 13; fi9++ {
                    fmt.Println(fi8)
                    v10 := []int{127, (fi9 * fi8)}
                    _ = fi9
                    _ = v10
                }
            case 4294967295:
                if (0 == -1) {
                    v5 = v5
                }
            case -9223372036854775808:
                v11 := GenS1{F0: -97, F1: 442, F2: 2}
                _ = v11
            }
            if false {
                var v12 float64 = -2.25
                v5 = v5
                _ = v12
            }
            v13 := GenS1{F0: 42, F1: 4294967295, F2: 247}
            _ = fi8
            _ = v13
        }
    }
    _ = v5
    return "\n\t"
}
func genFn1(p0 int64) uint32 {
    var v14 uint32 = 4294967295
    v14 = v14
    v15 := []int{32768, -2147483648, 617, 1}
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(fi16, p0)
        v14 -= 42
        if true {
            var v17 string = "foo bar"
            _ = v17
        } else {
            switch -578 {
            case 128:
                if len(v15) > 0 {
                    v15[0] = (fi16 >> ((418) & 63))
                }
            default:
                v18 := GenS1{F0: 100, F1: 2, F2: 0}
                _ = v18
            }
            for fi19 := 0; fi19 < 3; fi19++ {
                if true {
                    p0 += p0
                    fmt.Println(fi16)
                } else {
                    v20 := GenS1{F0: 100, F1: 446, F2: uint8(v14)}
                    p0 = -9223372036854775808
                    _ = v20
                }
                _ = fi19
            }
            fmt.Println(fi16, fi16)
        }
        _ = fi16
    }
    fmt.Println(v14, p0)
    _ = p0
    _ = v14
    _ = v15
    return (v14 / ((v14) | 1))
}

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 680
    var v3 bool = (v2 != 282)
    v4 := []int{-1, 0, 4294967295}
    v3 = true
    v2 -= 42
    v21 := GenS2{F0: "\n\t", F1: 617}
    switch 9223372036854775807 {
    case -295:
        var v22 bool = ("hello" <= "hello")
        _ = v22
    case -2147483648:
        v23 := GenS1{F0: -128, F1: v2, F2: 255}
        _ = v23
    case -1:
        var v24 float32 = (0.1 + ((0.1 * (1e10 / 3.14159)) - ((0.1 * 0.0) + 1e10)))
        _ = v24
    default:
        v3 = v3
    }
    if (1 != 977) {
        for fi25 := 0; fi25 < 9; fi25++ {
            v3 = v3
            for fi26 := 0; fi26 < 2; fi26++ {
                v2 = (v2 % ((798) | 1))
                v27 := GenS3{F0: 2, F1: 42, F2: 100}
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    }
    v28 := GenS4{F0: 42}
    v29 := GenS1{F0: 16, F1: v2, F2: 0}
    switch -786 {
    case -27:
        fmt.Println(v1)
    case 32768:
        _ = len(v4)
    case 2147483647:
        v30 := GenS5{F0: 2}
        _ = v30
    }
    v1 = 1.5
    for fi31 := 0; fi31 < 16; fi31++ {
        for fi32 := 0; fi32 < 18; fi32++ {
            if len(v4) > 0 {
                v4[0] = fi32
            }
            if v3 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                var v33 int8 = 100
                _ = v33
            }
            _ = fi32
        }
        v2 += v2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi31
    }
    switch -1 {
    case 4294967295:
        fmt.Println(v2)
    default:
        var v34 int8 = -128
        _ = v34
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3)
    if false {
        switch 127 {
        case 1:
            var v35 uint8 = 127
            _ = v35
        case -9223372036854775808:
            v3 = false
        default:
            v36 := GenS2{F0: "a", F1: 792}
            _ = v36
        }
        v4 = append(v4, 42)
    } else {
        var v37 uint = 1
        for fi38 := 0; fi38 < 9; fi38++ {
            v2 -= v2
            v39 := GenS6{F0: 76, F1: 127}
            _ = fi38
            _ = v39
        }
        v2 += 42
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
