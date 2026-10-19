package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int32; F2 uint16 }
type GenS3 struct { F0 bool; F1 int64; F2 int }
func genFn0() uint8 {
    if (1 <= 1) {
        var v7 int = -94
        _ = v7
    }
    var v8 string = "a"
    var v9 bool = false
    _ = v8
    _ = v9
    return 181
}
func genFn1(p0 string, p1 int) uint64 {
    if true {
        p0 = ((p0 + "a") + (("\n\t" + p0) + p0))
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        p0 = ""
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS1{F0: 0}
            _ = fi11
            _ = v12
        }
        p0 = p0
        _ = fi10
    }
    v13 := []int{}
    if ((false && (!false)) || false) {
        p1 = (^(p1))
    } else {
        v14 := GenS2{F0: 1, F1: 42, F2: 0}
        var v15 uint = 226
        var v16 uint8 = 2
        _ = v14
        _ = v15
        _ = v16
    }
    if true {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        if (float32(3.14159) >= -2.25) {
            if len(v13) > 0 {
                v13[0] = (-(p1))
            }
            if len(v13) > 0 {
                v13[0] = (p1 * (-(p1)))
            }
        } else {
            var v17 bool = (genFn0() == 42)
            if len(v13) > 0 {
                v13[0] = (-(p1))
            }
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v13
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 int16 = 1
    var v4 int32 = -347
    v5 := make(map[string]int)
    var v6 int16 = (v3 - v3)
    v18 := GenS1{F0: 2}
    var v19 bool = (false && true)
    v20 := GenS3{F0: (0 < (^(v3))), F1: 23, F2: -9223372036854775808}
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 uint = 44
        v4 = -236
        v19 = v19
        _ = fi21
        _ = v22
    }
    var v23 uint64 = 377
    v5["a"] = 32767
    v4 += (v4 & v4)
    v3 = -767
    v3 += -299
    v24 := GenS2{F0: 9223372036854775807, F1: (-(v4)), F2: 2}
    delete(v1, 2147483647)
    switch 68 {
    case 491:
        var v25 int = 128
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
