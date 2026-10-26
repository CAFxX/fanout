package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int8 }
func genFn0() uint64 {
    v5 := GenS1{F0: 864}
    if (100 < 100) {
        var v6 uint32 = 65535
        var v7 int16 = 880
        _ = v6
        _ = v7
    }
    _ = v5
    return 0
}
func genFn1(p0 uint32, p1 uint32, p2 uint32) int32 {
    v8 := GenS2{F0: 813, F1: 1}
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 bool = (0 > 0)
        fmt.Println(fi9)
        _ = fi9
        _ = v10
    }
    if (325 == int64(p1)) {
        p1 = (^(p2))
    }
    if false {
        v11 := GenS3{F0: (p2 - p0), F1: 226, F2: p0}
        if false {
            fmt.Println(p1, p2, p2)
            v12 := GenS2{F0: 0, F1: 32768}
            _ = v12
        } else {
            var v13 int16 = -1
            var v14 int64 = -9223372036854775808
            if (!true) {
                for fi15 := 0; fi15 < 18; fi15++ {
                    v16 := []int{-375, (fi15 * fi15), (fi15 ^ 127), fi15}
                    _ = fi15
                    _ = v16
                }
            } else {
                for fi17 := 0; fi17 < 2; fi17++ {
                    fmt.Println(v14, v14)
                    p1 += p2
                    _ = fi17
                }
            }
            _ = v13
            _ = v14
        }
        _ = v11
    } else {
        p2 = p1
        v18 := []int{128, 543, 528, 2147483647}
        _ = v18
    }
    var v19 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v8
    return -2147483648
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{905, 2147483648, 42, 2}
    v3 := make(map[string]bool)
    var v4 bool = true
    var v20 int8 = 127
    var v21 string = "\n\t"
    if v4 {
        fmt.Println(v4)
        if v4 {
            var v22 uint16 = 2
            v20 = v20
            var v23 uint16 = (^(v22))
            _ = v22
            _ = v23
        } else {
            v24 := GenS4{F0: 1, F1: 122}
            _ = v24
        }
    }
    var v25 uint8 = 0
    delete(v3, "a")
    v26 := GenS5{F0: v20}
    for fi27 := 0; fi27 < 16; fi27++ {
        for fi28 := 0; fi28 < 10; fi28++ {
            v2 = append(v2, 4294967295)
            delete(v3, "a")
            _ = fi28
        }
        _ = fi27
    }
    v29 := GenS5{F0: v20}
    delete(v3, v21)
    if v4 {
        v25 -= v25
        var v30 float32 = 3.14159
        _ = v30
    } else {
        for fi31 := 0; fi31 < 10; fi31++ {
            fmt.Println(v4, v20, v4)
            _ = fi31
        }
    }
    v32 := GenS2{F0: 65535, F1: 713}
    if ((!false) || (!(0 != -32768))) {
        if true {
            v25 += (v25 << ((42) & 63))
        }
        var v33 int8 = v20
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
