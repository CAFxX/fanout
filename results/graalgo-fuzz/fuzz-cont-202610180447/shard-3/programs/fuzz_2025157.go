package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1
    var v3 bool = true
    v4 := make(map[uint32]bool)
    v5 := []uint32{617}
    var v6 bool = (127 <= 42)
    v2 -= 42
    var v7 string = ("" + "\n\t")
    v8 := GenS1{F0: 32767}
    var v9 int8 = 42
    v9 = 1
    fmt.Println(v3)
    v10 := GenS2{F0: v9, F1: -902}
    fmt.Println(v9)
    v2 += v2
    var v11 float32 = ((float32(float64(3.14159)) - ((float32(0.1) / (-2.25 * 0.1)) / 3.14159)) * ((3.14159 / (-2.25 / 1e10)) - 0.1))
    _ = len(v1)
    if v3 {
        if ((v11 - (float32(1e10) - 0.1)) < 3.14159) {
            v2 -= v2
        }
        fmt.Println(v7, v2, v3)
    } else {
        switch v7 {
        case "foo bar":
            var v12 int64 = 42
            _ = v12
        case "a":
            if v3 {
                if ((2147483647 != -731) && (0 != 4294967295)) {
                    var v13 uint8 = 218
                    _ = v13
                }
                v14 := GenS3{F0: 4294967295, F1: 9223372036854775807}
                for fi15 := 0; fi15 < 20; fi15++ {
                    if ((((fi15 | fi15) >= 42) || true) || (80 == 0)) {
                        var v16 uint64 = 808
                        for fi17 := 0; fi17 < 20; fi17++ {
                            v18 := v4[1]
                            if v18 {
                                if len(v5) > 0 {
                                    v5[0] = 42
                                }
                                v7 = "hello"
                            } else {
                                if v6 {
                                    v4[1] = (!v18)
                                } else {
                                    var v19 bool = false
                                    _ = v19
                                }
                                if true {
                                    var v20 uint64 = v16
                                    var v21 int16 = 766
                                    _ = v20
                                    _ = v21
                                }
                                fmt.Println(v2)
                                v11 -= 0.1
                            }
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    }
                    v3 = v3
                    _ = fi15
                }
                _ = v14
            }
        case "":
            var v22 uint16 = 713
            _ = v22
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 float64 = 0.1
            for fi25 := 0; fi25 < 20; fi25++ {
                _ = len(v1)
                v5 = append(v5, 2147483648)
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
    }
    switch v7 {
    case "\n\t":
        v6 = ((!(2 <= 125)) || (255 != 65535))
    case "foo bar":
        v26 := GenS1{F0: 65535}
        _ = v26
    default:
        if (3.14159 <= -2.25) {
            var v27 int8 = (v9 & v9)
            var v28 float64 = -2.25
            _ = v27
            _ = v28
        }
    }
    v29 := GenS4{F0: 42}
    if ((("" + v7) <= v7) || v6) {
        var v30 uint32 = 2147483647
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
