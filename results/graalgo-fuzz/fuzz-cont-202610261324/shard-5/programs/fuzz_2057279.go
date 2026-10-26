package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 int16, p1 uint8) uint16 {
    p1 += 42
    p0 = (p0 >> ((2) & 63))
    var v4 uint8 = (p1 << ((255) & 63))
    var v5 int32 = -57
    var v6 int = 32767
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 bool, p1 int8) uint32 {
    v7 := []int{0, -1, 128}
    var v8 int32 = 42
    var v9 int8 = 44
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 uint64 = 0
    var v10 uint64 = v3
    v11 := GenS1{F0: -722, F1: 1, F2: 1}
    var v12 uint = 42
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            if true {
                var v15 int16 = 32767
                _ = v15
            } else {
                v10 -= v3
                fmt.Println(v12, fi14)
            }
            if true {
                var v16 int8 = -1
                _ = v16
            }
            var v17 float64 = 0.0
            _ = fi14
            _ = v17
        }
        _ = fi13
    }
    fmt.Println(v12)
    v1["foo bar"] = ("a" + ("hello" + ("foo bar" + "")))
    if (true || false) {
        var v18 uint64 = uint64(v12)
        _ = v18
    }
    v19 := GenS2{F0: -550, F1: (false && true)}
    switch 42 {
    case -9223372036854775808:
        fmt.Println(v10, v3)
    default:
        if false {
            for fi20 := 0; fi20 < 9; fi20++ {
                if (-77 < -32768) {
                    var v21 int16 = -1
                    _ = v21
                } else {
                    v22 := GenS2{F0: fi20, F1: false}
                    _ = v22
                }
                v10 = uint64(v12)
                switch -9223372036854775808 {
                case 32768:
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v24 := []int{(fi20 + fi20), (fi20 | fi20), 0, (-(fi20))}
                        _ = fi23
                        _ = v24
                    }
                case -1:
                    v25 := GenS3{F0: fi20, F1: 23}
                    _ = v25
                }
                _ = fi20
            }
            if false {
                var v26 bool = (false && (1000 > 127))
                _ = v26
            } else {
                v27 := GenS4{F0: (v12 - (v12 | v12))}
                v3 = 0
                _ = v27
            }
            var v28 bool = false
            _ = v28
        }
    }
    var v29 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v29
    _ = v3
}
