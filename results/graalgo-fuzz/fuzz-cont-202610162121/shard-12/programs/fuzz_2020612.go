package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32) uint32 {
    v5 := GenS1{F0: 0, F1: 42}
    if false {
        if false {
            p0 -= p0
            p0 = p0
        }
        var v6 int8 = 0
        _ = v6
    }
    _ = p0
    _ = v5
    return 2
}
func genFn1(p0 uint64) uint8 {
    var v7 int16 = 1
    p0 -= 1
    _ = p0
    _ = v7
    return 181
}

func main() {
    var v1 int16 = 1000
    var v2 string = ("" + "hello")
    var v3 int64 = -276
    v4 := make(map[string]uint32)
    var v8 uint8 = 151
    var v9 int64 = (v3 % ((0) | 1))
    var v10 int32 = 0
    switch 127 {
    case 42:
        v11 := GenS2{F0: (false && ((v1 < v1) && (3.14159 > 1.5))), F1: v1}
        _ = v11
    case -9223372036854775808:
        var v12 int32 = v10
        _ = v12
    case 32768:
        v13 := []int{128}
        _ = v13
    }
    if false {
        var v14 int64 = 261
        var v15 uint64 = 0
        if (!false) {
            v16 := []int{42, 2147483647, -346, 1}
            var v17 uint16 = 1
            _ = v16
            _ = v17
        } else {
            var v18 int64 = (v9 - v9)
            if false {
                fmt.Println(v3)
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS3{F0: genFn1(v15), F1: 0, F2: 0}
                v21 := GenS2{F0: false, F1: 32767}
                v1 += (v1 + v1)
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v18
        }
        _ = v14
        _ = v15
    }
    var v22 int16 = v1
    fmt.Println(v1)
    v1 = (^(v22))
    switch 9223372036854775807 {
    case -2147483648:
        var v23 uint32 = 65535
        _ = v23
    case -1:
        fmt.Println(v2)
    case 128:
        if false {
            var v24 uint32 = 770
            _ = v24
        } else {
            _ = len(v4)
            if true {
                for fi25 := 0; fi25 < 5; fi25++ {
                    switch (-(fi25)) {
                    case -635:
                        var v26 bool = true
                        _ = v26
                    case 4294967295:
                        if false {
                            var v27 float64 = 1e10
                            _ = v27
                        }
                    default:
                        v28 := v4[v2]
                        _ = v28
                    }
                    var v29 int32 = (v10 >> ((4294967295) & 63))
                    _ = fi25
                    _ = v29
                }
                v10 = v10
            }
        }
    default:
        v30 := GenS4{F0: (v8 + (v8 - (^(v8)))), F1: -926}
        _ = v30
    }
    v8 = 128
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS5{F0: 1, F1: v22}
        _ = fi31
        _ = v32
    }
    if (32767 == uint16(v8)) {
        var v33 int32 = 320
        v34 := []int{2147483647, 2147483647, 128, 417}
        _ = v33
        _ = v34
    }
    if ((v2 + v2) < "") {
        if true {
            v22 = (v1 | 0)
            v35 := GenS3{F0: 128, F1: 2147483647, F2: uint16(v8)}
            _ = v35
        }
        switch 2147483647 {
        case -1:
            var v36 uint8 = genFn1(581)
            _ = v36
        default:
            v37 := []int{-409, 0, 2147483647, 127}
            _ = v37
        }
    } else {
        v38 := []int{0, -2147483648}
        var v39 int8 = 42
        _ = v38
        _ = v39
    }
    v40 := GenS6{F0: 0}
    v41 := GenS5{F0: 0, F1: (v1 * (v22 & (v1 + (v1 | v1))))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v8
    _ = v9
}
