package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int32; F1 uint; F2 uint64 }
func genFn0() int16 {
    if (!false) {
        v5 := []int{-2147483648}
        _ = v5
    } else {
        var v6 uint16 = 899
        v7 := GenS1{F0: 113, F1: 113}
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(fi8)
            switch "" {
            case "foo bar":
                v10 := GenS2{F0: -291, F1: 1000}
                _ = v10
            case "hello":
                var v11 uint32 = 355
                _ = v11
            }
            fmt.Println(fi8, fi9)
            _ = fi9
        }
        _ = fi8
    }
    var v12 float64 = -2.25
    var v13 int32 = 42
    var v14 uint16 = 32767
    _ = v12
    _ = v13
    _ = v14
    return 8
}

func main() {
    v1 := []uint32{42, 806}
    var v2 int8 = 0
    var v3 int = 9223372036854775807
    var v4 uint8 = 127
    var v15 int64 = -438
    var v16 float32 = (1e10 - 0.0)
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS1{F0: 65535, F1: 127}
        if ((42 <= 0) && true) {
            if ((v4 - 1) < v4) {
                v19 := GenS3{F0: -1, F1: 65535, F2: 42}
                fmt.Println(v4)
                _ = v19
            } else {
                if (370 > (fi17 >> ((4294967295) & 63))) {
                    v3 = v3
                    v20 := GenS1{F0: 65535, F1: v4}
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 uint32 = 2147483648
                        v16 += float32(0.1)
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                }
            }
        }
        fmt.Println(v16, v4, v15)
        _ = fi17
        _ = v18
    }
    v15 = (v15 >> ((42) & 63))
    var v23 uint8 = v4
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 int8 = 0
        _ = fi24
        _ = v25
    }
    v23 += (^(v23))
    v1 = append(v1, 4294967295)
    for fi26 := 0; fi26 < 18; fi26++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v27 int16 = 42
        _ = fi26
        _ = v27
    }
    switch "" {
    case "":
        for fi28 := 0; fi28 < 9; fi28++ {
            if (!(42 <= 42)) {
                v4 -= (v23 / ((v4) | 1))
            } else {
                v29 := GenS2{F0: v15, F1: 238}
                v30 := GenS2{F0: (v15 | v15), F1: 1000}
                _ = len(v1)
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                fmt.Println(v23)
                _ = fi31
            }
            _ = fi28
        }
    }
    v4 = v4
    var v32 int = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
}
