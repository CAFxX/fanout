package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
func genFn0() int16 {
    var v5 int16 = -614
    switch -593 {
    case 2147483647:
        for fi6 := 0; fi6 < 3; fi6++ {
            v5 -= 1000
            v7 := []int{fi6}
            _ = fi6
            _ = v7
        }
    case -909:
        v8 := GenS1{F0: 833, F1: 1}
        _ = v8
    }
    v5 -= v5
    _ = v5
    return 82
}
func genFn1(p0 uint32, p1 int, p2 uint16) uint64 {
    if true {
        v9 := []int{32768, (p1 + (p1 - p1)), (p1 ^ p1), (^(p1))}
        var v10 int = p1
        _ = v10
        _ = v9
    } else {
        var v11 int = 9223372036854775807
        fmt.Println(p2, v11, p1)
        _ = v11
    }
    p1 += p1
    var v12 bool = true
    var v13 uint32 = p0
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int8 = 45
    var v4 uint32 = 0
    fmt.Println(v3, v4)
    fmt.Println(v4, v4)
    v4 += v4
    v14 := GenS1{F0: uint16(v4), F1: -343}
    v15 := v1[v4]
    if (true || (-947 < 0)) {
        switch -2147483648 {
        case -9223372036854775808:
            v4 += (v15 ^ v4)
        case -597:
            if (-1 != int64(v15)) {
                v16 := v1[65535]
                _ = v16
            }
        case 128:
            if false {
                if false {
                    var v17 uint8 = 127
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 uint8 = (^(v17))
                        _ = fi18
                        _ = v19
                    }
                    v17 -= v17
                    _ = v17
                } else {
                    v20 := GenS1{F0: 32767, F1: -558}
                    _ = v20
                }
                v21 := []int{311, 2147483647}
                _ = v21
            }
        }
    }
    v22 := []int{-646, 4294967295}
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 int8 = 0
        if true {
            v25 := GenS1{F0: 793, F1: -1}
            v26 := GenS2{F0: -32768, F1: (v4 / (((v15 >> ((255) & 63))) | 1)), F2: 127}
            switch 120 {
            case -9223372036854775808:
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 uint32 = 512
                    var v29 uint8 = 242
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            case 32767:
                if false {
                    v3 = (v3 / ((100) | 1))
                    var v30 uint16 = 65535
                    switch -205 {
                    case 32767:
                        var v31 string = "foo bar"
                        _ = v31
                    case 42:
                        v1[(v4 * (v15 ^ (v4 + 965)))] = (^(v4))
                    default:
                        var v32 int16 = -32768
                        _ = v32
                    }
                    _ = v30
                } else {
                    delete(v2, (v4 + v4))
                }
            case 312:
                fmt.Println(v15)
            }
            _ = v25
            _ = v26
        } else {
            var v33 int32 = 0
            var v34 string = (((("foo bar" + "a") + ("a" + "hello")) + "") + "\n\t")
            _ = v33
            _ = v34
        }
        _ = fi23
        _ = v24
    }
    var v35 uint64 = 2
    v15 += (v15 - (v15 >> ((2) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v4
}
