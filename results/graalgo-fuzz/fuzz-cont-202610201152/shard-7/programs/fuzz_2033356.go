package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0(p0 int8) uint16 {
    v5 := GenS1{F0: 42, F1: false, F2: -1}
    if ((!true) || ((-2.25 / -2.25) < 3.14159)) {
        var v6 int8 = (p0 >> ((42) & 63))
        var v7 float32 = ((((float32(3.14159) - 1.5) * (1.5 + 3.14159)) + 0.0) + -2.25)
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: -655, F1: 128}
    var v9 int64 = 1
    v10 := []int{42}
    _ = p0
    _ = v10
    _ = v5
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint32, p1 uint64) int16 {
    fmt.Println(p1, p1)
    if (2147483647 <= -2147483648) {
        v11 := GenS3{F0: 42}
        fmt.Println(p1, p1)
        _ = v11
    } else {
        var v12 int8 = 42
        v12 = (^(v12))
        var v13 string = "hello"
        _ = v12
        _ = v13
    }
    p1 = p1
    p0 += (^(p0))
    v14 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = v14
    return -813
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 42
    var v3 string = ""
    v4 := make(map[int]uint32)
    for fi15 := 0; fi15 < 15; fi15++ {
        delete(v4, (fi15 | -9223372036854775808))
        var v16 int16 = -379
        for fi17 := 0; fi17 < 5; fi17++ {
            if ((float32(((((0.0 - (0.0 - 3.14159)) - 1e10) / -2.25) / 1.5)) * 3.14159) < (0.1 / 0.1)) {
                v18 := GenS2{F0: (v2 >> ((0) & 63)), F1: 128}
                v3 = v3
                _ = v18
            }
            _ = len(v4)
            switch (fi15 | (fi17 & (fi17 - -1))) {
            case 0:
                v2 -= (v2 * -2147483648)
            }
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    v19 := []int{42}
    v2 = (v2 + (v2 ^ v2))
    _ = len(v19)
    v3 = (v3 + ("\n\t" + v3))
    switch v3 {
    case "":
        v2 -= v2
    }
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS4{F0: 903, F1: 42, F2: 335}
            for fi22 := 0; fi22 < 1; fi22++ {
                fmt.Println(fi22, v3, v2)
                var v23 uint64 = 517
                if len(v19) > 0 {
                    v19[0] = fi20
                }
                _ = fi22
                _ = v23
            }
            _ = fi20
            _ = v21
        }
    } else {
        var v24 int8 = -1
        _ = v24
    }
    var v25 float64 = (0.1 + 0.0)
    v26 := GenS3{F0: 0}
    var v27 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
