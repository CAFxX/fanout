package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint }

func main() {
    var v1 int16 = -1
    v2 := []int{32767}
    v3 := make(map[uint32]uint32)
    v4 := []int{32767, 632}
    if false {
        var v5 int32 = 795
        if (1.5 <= 1.5) {
            if ((v5 | (v5 % ((v5) | 1))) < -2147483648) {
                if false {
                    fmt.Println(v1)
                    var v6 float64 = 0.1
                    for fi7 := 0; fi7 < 15; fi7++ {
                        var v8 int64 = -9223372036854775808
                        v9 := GenS1{F0: (245 >= 182), F1: 65535, F2: -416}
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    _ = v6
                } else {
                    if (-722 >= -1) {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v1, v1)
                        }
                        v5 = (v5 ^ (v5 ^ -2))
                    }
                    delete(v3, 478)
                    v1 = (v1 << ((925) & 63))
                }
            }
        } else {
            if (!true) {
                var v10 int32 = v5
                _ = v10
            }
            switch 128 {
            case -654:
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 uint64 = 4294967295
                    var v13 uint32 = 2
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 1:
                var v14 string = "\n\t"
                _ = v14
            case 32767:
                if false {
                    fmt.Println(v5, v5)
                } else {
                    fmt.Println(v1, v1)
                }
            }
        }
        fmt.Println(v1)
        _ = v5
    } else {
        _ = len(v4)
        var v15 int64 = -678
        _ = v15
    }
    v16 := GenS2{F0: -158, F1: 1}
    v1 += v1
    v17 := GenS3{F0: ((1e10 >= 3.14159) || true), F1: 196, F2: 4294967295}
    var v18 uint64 = 817
    for fi19 := 0; fi19 < 13; fi19++ {
        fmt.Println(fi19)
        var v20 string = ("a" + "hello")
        _ = fi19
        _ = v20
    }
    var v21 int = 0
    var v22 bool = false
    var v23 int = v21
    var v24 uint = 1
    v18 -= v18
    fmt.Println(v23, v21, v22)
    for fi25 := 0; fi25 < 2; fi25++ {
        v4 = append(v4, v21)
        v26 := GenS1{F0: (1 != 42), F1: 368, F2: v23}
        v3[274] = 386
        _ = fi25
        _ = v26
    }
    delete(v3, 354)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
