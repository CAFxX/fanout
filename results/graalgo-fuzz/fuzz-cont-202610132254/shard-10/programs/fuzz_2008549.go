package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int16 }
func genFn0() uint64 {
    v6 := []int{42}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = v6
    return 1
}
func genFn1(p0 int32) uint {
    var v7 uint32 = 2
    fmt.Println(p0, p0)
    p0 -= 42
    p0 = (p0 ^ 1)
    _ = p0
    _ = v7
    return 578
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-603, 781}
    var v3 uint8 = 56
    v4 := make(map[uint32]uint32)
    var v5 int = -2147483648
    fmt.Println(v5)
    fmt.Println(v5, v3)
    var v8 float32 = (float32(3.14159) - ((-2.25 / 1e10) / ((1e10 / 0.1) - (3.14159 - -2.25))))
    fmt.Println(v5)
    if len(v2) > 0 {
        v2[0] = (v5 >> ((genFn1(0)) & 63))
    }
    var v9 int8 = 42
    var v10 uint = genFn1(173)
    if false {
        v11 := GenS1{F0: -29, F1: 325}
        v5 += v5
        v8 = v8
        _ = v11
    } else {
        v12 := GenS1{F0: v9, F1: 897}
        _ = v12
    }
    var v13 int16 = 1
    v13 = v13
    if (!true) {
        v14 := GenS2{F0: (^(v10)), F1: (v3 - v3), F2: 0}
        fmt.Println(v5, v9)
        _ = v14
    }
    var v15 int32 = 42
    _ = len(v4)
    var v16 uint16 = 65535
    var v17 float64 = 1e10
    if (float64(((v8 + (v8 + float32(0.1))) / (v8 + (v8 + 0.0)))) >= (v17 + 0.1)) {
        v18 := GenS1{F0: (-(v9)), F1: 315}
        if true {
            fmt.Println(v16, v8)
            for fi19 := 0; fi19 < 2; fi19++ {
                fmt.Println(fi19)
                _ = fi19
            }
            v20 := GenS3{F0: 1000}
            _ = v20
        }
        v1[42] = (-(v5))
        _ = v18
    }
    v5 -= v5
    var v21 string = ("a" + ("" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
