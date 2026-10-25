package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int8; F2 int64 }
type GenS4 struct { F0 int32; F1 bool; F2 uint64 }
type GenS5 struct { F0 int; F1 int32; F2 uint64 }
func genFn0(p0 bool, p1 int32, p2 uint16) uint32 {
    if (((true && (p0 || false)) && true) || (!p0)) {
        v5 := GenS1{F0: 1}
        p1 = p1
        _ = v5
    } else {
        if false {
            v6 := []int{128}
            var v7 float64 = (0.0 - 1e10)
            _ = v6
            _ = v7
        } else {
            v8 := GenS2{F0: 1, F1: 42, F2: 65535}
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 bool = p0
                if (127 <= 0) {
                    fmt.Println(p0, fi9, p0)
                    for fi11 := 0; fi11 < 2; fi11++ {
                        var v12 int16 = 486
                        if (v10 && true) {
                            v13 := GenS3{F0: 2147483648, F1: 42, F2: 706}
                            _ = v13
                        } else {
                            v14 := []int{730, (^(fi11))}
                            _ = v14
                        }
                        p1 -= (p1 << ((1) & 63))
                        _ = fi11
                        _ = v12
                    }
                    if p0 {
                        var v15 bool = (466 <= 786)
                        _ = v15
                    }
                }
                v16 := GenS4{F0: p1, F1: false, F2: 0}
                _ = fi9
                _ = v10
                _ = v16
            }
            _ = v8
        }
        fmt.Println(p2, p2, p0)
        var v17 int8 = -123
        _ = v17
    }
    var v18 int32 = (p1 ^ (p1 | (^(p1))))
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 282
}

func main() {
    var v1 int64 = -1
    var v2 int = 4294967295
    var v3 uint16 = 241
    var v4 int64 = v1
    v19 := GenS3{F0: 2147483647, F1: -51, F2: (v4 + v1)}
    v2 += -923
    for fi20 := 0; fi20 < 9; fi20++ {
        v3 += 255
        v21 := []int{(fi20 | (v2 >> ((592) & 63))), -277}
        v21 = append(v21, v2)
        _ = fi20
        _ = v21
    }
    v1 = 9223372036854775807
    var v22 uint8 = 2
    v23 := GenS5{F0: (^(v2)), F1: 2147483647, F2: 0}
    v4 -= v1
    var v24 uint64 = 2
    v25 := []int{v2, v2, (^(v2))}
    var v26 string = ("hello" + ("a" + ""))
    if len(v25) > 0 {
        v25[0] = int(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
