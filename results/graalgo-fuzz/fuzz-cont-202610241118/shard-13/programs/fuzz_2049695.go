package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int8; F1 uint32; F2 string }
type GenS3 struct { F0 uint32; F1 int; F2 uint32 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := []string{"", "hello"}
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    var v4 float32 = 0.1
    v5 := GenS1{F0: 42, F1: -1}
    v4 = (v4 / float32(((1e10 - (float64((1.5 * ((float32(((3.14159 + (0.0 / 0.1)) / 1e10)) - float32((1e10 + float64((v4 - v4))))) + v4))) - -2.25)) * 0.0)))
    v6 := GenS2{F0: -38, F1: 833, F2: (("" + "hello") + "a")}
    fmt.Println(v4)
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(fi7, v4)
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(fi8, fi7)
            _ = fi8
        }
        if (100 < 27) {
            if true {
                for fi9 := 0; fi9 < 13; fi9++ {
                    if false {
                        v10 := GenS3{F0: 2147483648, F1: fi7, F2: 4294967295}
                        fmt.Println(fi9)
                        _ = v10
                    }
                    v4 -= ((v4 * 3.14159) * v4)
                    _ = fi9
                }
                var v11 int64 = -1
                _ = v11
            } else {
                if (4294967295 < 387) {
                    v12 := GenS2{F0: -57, F1: 2, F2: "hello"}
                    fmt.Println(v4, fi7)
                    if ((930 == -1) || false) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = v12
                }
                var v13 uint = 65535
                fmt.Println(v4, v13, fi7)
                _ = v13
            }
            if (false || true) {
                v14 := GenS1{F0: -9223372036854775808, F1: -313}
                var v15 uint64 = 821
                switch (("hello" + "\n\t") + ("" + "\n\t")) {
                case "a":
                    var v16 uint64 = (v15 - (v15 - v15))
                    _ = v16
                }
                _ = v14
                _ = v15
            } else {
                fmt.Println(v4, fi7)
                v17 := GenS4{F0: -9223372036854775808, F1: 32767}
                var v18 uint16 = 1
                _ = v17
                _ = v18
            }
        }
        _ = fi7
    }
    var v19 uint64 = 736
    v20 := GenS5{F0: 1}
    if (285 != 1) {
        var v21 int = 42
        _ = v21
    } else {
        for fi22 := 0; fi22 < 14; fi22++ {
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 uint32 = 0
                v4 += (float32(float64(((v4 + 0.0) - ((((float32((3.14159 * 3.14159)) / v4) * v4) / 1.0) - (v4 / 0.1))))) * 3.14159)
                if true {
                    v19 -= 4294967295
                    v25 := GenS2{F0: -44, F1: (v24 % ((v24) | 1)), F2: "\n\t"}
                    _ = v25
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 int = fi26
        _ = fi26
        _ = v27
    }
    var v28 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
