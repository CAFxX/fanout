package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 bool; F1 int; F2 int32 }
func genFn0(p0 uint32, p1 uint8) int {
    if true {
        fmt.Println(p1, p1)
    } else {
        p0 = p0
        v7 := []int{2147483647, 2147483647}
        _ = v7
    }
    var v8 uint = 2
    if true {
        v9 := GenS1{F0: 169, F1: uint16(p0)}
        _ = v9
    } else {
        var v10 int32 = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 int, p1 int64, p2 bool) int64 {
    v11 := GenS1{F0: 2, F1: 2}
    v12 := []int{p0, p0, (p0 * 128), 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return int64(p0)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 24
    var v3 int64 = 1
    var v4 float64 = 3.14159
    var v5 int16 = 1000
    var v6 int64 = v3
    v3 += (v3 / ((v3) | 1))
    v4 -= (v4 + v4)
    v13 := GenS2{F0: 613, F1: 923, F2: 0}
    var v14 uint8 = 128
    fmt.Println(v3, v4)
    for fi15 := 0; fi15 < 15; fi15++ {
        v1[42] = 1
        if ((false && true) || false) {
            if (!(v3 != v3)) {
                for fi16 := 0; fi16 < 20; fi16++ {
                    _ = len(v1)
                    _ = fi16
                }
                v17 := GenS3{F0: (true && true), F1: (fi15 >> ((65535) & 63)), F2: -1}
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 float64 = v4
                fmt.Println(v19, v3)
                _ = fi18
                _ = v19
            }
            v5 = (v5 % (((v5 ^ 32767)) | 1))
        }
        _ = fi15
    }
    v2 += v2
    v20 := GenS1{F0: 2, F1: 42}
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 float32 = 1e10
        _ = fi21
        _ = v22
    }
    var v23 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
