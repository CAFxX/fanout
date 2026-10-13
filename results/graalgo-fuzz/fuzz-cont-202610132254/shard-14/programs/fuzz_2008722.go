package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int8; F1 int32; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint16; F2 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int32, p1 uint16) uint8 {
    if false {
        if (true || (p1 == (^(p1)))) {
            var v7 float32 = (-2.25 / 1.5)
            v7 = 0.0
            _ = v7
        } else {
            var v8 uint16 = p1
            _ = v8
        }
        var v9 uint32 = 2147483647
        _ = v9
    } else {
        var v10 float64 = 1.5
        if true {
            p0 -= p0
            v11 := GenS1{F0: p1, F1: 0, F2: 4294967295}
            var v12 int8 = 0
            _ = v11
            _ = v12
        } else {
            v13 := GenS2{F0: -1, F1: p0, F2: 0}
            v10 += v10
            _ = v13
        }
        if ((v10 - v10) < v10) {
            p0 -= p0
            var v14 int32 = p0
            v15 := GenS3{F0: 744}
            _ = v14
            _ = v15
        } else {
            v16 := []int{-276, -9223372036854775808, 1, -888}
            _ = v16
        }
        _ = v10
    }
    var v17 uint16 = (p1 & 32767)
    var v18 float64 = 1e10
    p1 -= (v17 / ((v17) | 1))
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 128
}
func genFn1() int {
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            if ((2 >= 2) || false) {
                var v21 bool = true
                v21 = ((1e10 == ((float32((1e10 * ((0.0 + 1.5) + 3.14159))) + (-2.25 - 1.5)) + 0.1)) || (!v21))
                switch fi19 {
                case 128:
                    v21 = v21
                case 2147483647:
                    fmt.Println(fi19, fi20, fi20)
                }
                _ = v21
            }
            fmt.Println(fi20, fi20)
            if ((((true || (2147483648 == 42)) || true) || true) && ((42 > 127) && (!(1 != 0)))) {
                var v22 uint16 = 2
                var v23 float64 = float64((1.5 - -2.25))
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        v24 := GenS4{F0: "a", F1: 32767, F2: 42}
        var v25 int64 = 9223372036854775807
        _ = fi19
        _ = v24
        _ = v25
    }
    switch 127 {
    case -1:
        var v26 uint = 42
        _ = v26
    }
    return -656
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (1000 < -1)
    v3 := make(map[uint32]uint32)
    var v4 int8 = -1
    v5 := []uint32{42}
    var v6 bool = false
    if len(v5) > 0 {
        v5[0] = 2147483648
    }
    fmt.Println(v6, v4)
    v6 = v2
    v27 := GenS5{F0: 42}
    v6 = false
    fmt.Println(v4, v2)
    v28 := GenS4{F0: "\n\t", F1: 255, F2: 9223372036854775807}
    var v29 float32 = 0.1
    var v30 uint = 0
    v6 = (935 < 463)
    if false {
        var v31 uint64 = 42
        _ = v31
    }
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
