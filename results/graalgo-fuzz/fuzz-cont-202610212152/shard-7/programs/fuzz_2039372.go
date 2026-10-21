package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int32 }

func main() {
    var v1 float64 = 1.5
    v2 := []string{}
    v3 := make(map[int]uint32)
    v4 := []string{}
    var v5 string = "foo bar"
    var v6 int64 = 1
    v6 = 0
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            if ((v6 >> ((301) & 63)) <= (v6 << ((316) & 63))) {
                if false {
                    var v9 uint16 = 255
                    var v10 float32 = 1.5
                    _ = v10
                    _ = v9
                }
                var v11 int16 = 1000
                fmt.Println(v11, v1, v11)
                _ = v11
            } else {
                _ = len(v3)
                var v12 int32 = 1
                v3[127] = 1
                _ = v12
            }
            _ = fi8
        }
        if ("a" != ("foo bar" + v5)) {
            v5 = "a"
        }
        _ = fi7
    }
    if true {
        fmt.Println(v1, v6)
        var v13 int16 = -887
        _ = v13
    }
    v6 = 1
    if true {
        var v14 uint64 = 4294967295
        v15 := GenS1{F0: 0, F1: 1, F2: 18446744073709551615}
        v6 = (v6 >> ((42) & 63))
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 6; fi16++ {
            v6 -= 0
            for fi17 := 0; fi17 < 10; fi17++ {
                if false {
                    v18 := GenS2{F0: (1.5 > 0.0)}
                    if false {
                        v5 = ("hello" + v5)
                        var v19 float64 = (float64((1.5 + 3.14159)) * (float64(((float32(-2.25) * 1e10) / -2.25)) - v1))
                        fmt.Println(v19)
                        _ = v19
                    }
                    var v20 int16 = 1
                    _ = v18
                    _ = v20
                }
                v6 = (-(v6))
                var v21 uint8 = 0
                _ = fi17
                _ = v21
            }
            v22 := GenS1{F0: 671, F1: 42, F2: 67}
            _ = fi16
            _ = v22
        }
        fmt.Println(v6, v6, v6)
    }
    if true {
        for fi23 := 0; fi23 < 15; fi23++ {
            switch fi23 {
            case -36:
                v5 = v5
            default:
                _ = len(v4)
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                v6 = v6
                fmt.Println(v1, fi24, fi23)
                _ = fi24
            }
            _ = fi23
        }
        var v25 uint32 = 65535
        _ = v25
    }
    var v26 int64 = (-(v6))
    v6 = v6
    if true {
        v27 := GenS3{F0: 127, F1: 42}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
