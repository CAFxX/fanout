package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 bool; F1 int; F2 bool }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int64, p1 uint32, p2 int16) int {
    v5 := GenS1{F0: false, F1: 127}
    v6 := []int{32768, 128, -9223372036854775808, -1}
    var v7 int64 = (p0 | (p0 * p0))
    switch 148 {
    case 4294967295:
        if (true || true) {
            v8 := GenS1{F0: (1.5 == (0.0 + (1.5 * -2.25))), F1: -1}
            _ = v8
        }
    case 769:
        if len(v6) > 0 {
            v6[0] = 0
        }
    case -9223372036854775808:
        p0 -= v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint, p1 bool) int8 {
    p1 = (!false)
    if p1 {
        v9 := GenS2{F0: true, F1: 628, F2: true}
        if p1 {
            var v10 uint8 = 223
            var v11 int32 = 42
            v12 := []int{}
            _ = v10
            _ = v11
            _ = v12
        }
        switch "foo bar" {
        case "\n\t":
            if p1 {
                var v13 int64 = 0
                _ = v13
            }
        default:
            var v14 int16 = -934
            _ = v14
        }
        _ = v9
    }
    v15 := GenS2{F0: (4294967295 < -567), F1: 753, F2: (!(!(((((float32(float64(1.5)) * ((1e10 / -2.25) + ((0.0 - ((-2.25 / 1.0) / (1e10 + -2.25))) / -2.25))) - 0.0) - 1e10) == 0.1) && (!p1))))}
    v16 := []int{9223372036854775807, 128, int(p0), 42}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return -19
}

func main() {
    var v1 float32 = 0.1
    var v2 uint8 = 2
    v3 := make(map[uint32]bool)
    var v4 uint64 = 42
    delete(v3, 0)
    switch 32767 {
    case 42:
        fmt.Println(v4, v1)
    case 128:
        if (v2 <= v2) {
            v2 = (v2 >> ((4294967295) & 63))
        }
    }
    var v17 uint8 = (v2 | (v2 + (v2 & (v2 & (^(v2))))))
    v18 := GenS2{F0: true, F1: -284, F2: (!(!(v17 != v17)))}
    v19 := GenS1{F0: (!true), F1: -9223372036854775808}
    var v20 uint32 = 42
    var v21 uint64 = (v4 >> ((247) & 63))
    v22 := GenS2{F0: true, F1: genFn0(int64(v21), (v20 & v20), -934), F2: true}
    var v23 int8 = 100
    if true {
        var v24 uint32 = (v20 | v20)
        var v25 int16 = -32768
        var v26 int32 = 2147483647
        _ = v24
        _ = v25
        _ = v26
    } else {
        v27 := []int{}
        _ = v27
    }
    if (v17 >= (v17 * (v2 / ((v2) | 1)))) {
        v28 := GenS1{F0: ((255 == 42) || ((true && (255 >= 255)) || (159 == 2147483647))), F1: -9223372036854775808}
        var v29 uint = 65535
        var v30 int32 = 994
        _ = v28
        _ = v29
        _ = v30
    } else {
        fmt.Println(v17)
        v31 := v3[v20]
        switch 4294967295 {
        case 0:
            for fi32 := 0; fi32 < 9; fi32++ {
                if v31 {
                    fmt.Println(fi32, v21)
                } else {
                    v33 := GenS3{F0: (v2 + (v2 / ((v17) | 1))), F1: (-2147483648 >= 218)}
                    for fi34 := 0; fi34 < 8; fi34++ {
                        var v35 string = "foo bar"
                        _ = fi34
                        _ = v35
                    }
                    delete(v3, v20)
                    _ = v33
                }
                v36 := GenS4{F0: (v21 * uint64(v20))}
                fmt.Println(v17, v31)
                _ = fi32
                _ = v36
            }
        case 32767:
            v37 := GenS5{F0: 32767}
            _ = v37
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
