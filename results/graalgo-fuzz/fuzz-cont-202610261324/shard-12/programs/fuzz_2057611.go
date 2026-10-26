package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint64 = 826
    var v2 uint32 = 94
    v3 := make(map[int]int)
    var v4 int32 = 916
    v2 = (v2 - v2)
    if true {
        var v5 float64 = 1e10
        v6 := GenS1{F0: v4, F1: int64(v1), F2: 182}
        v2 = 2147483647
        _ = v5
        _ = v6
    } else {
        v7 := GenS2{F0: 0}
        fmt.Println(v1, v2)
        _ = v7
    }
    if (3.14159 != (0.1 - 1e10)) {
        var v8 int = 4294967295
        _ = v8
    }
    var v9 int32 = v4
    v10 := []int{-1}
    for fi11 := 0; fi11 < 20; fi11++ {
        switch "a" {
        case "foo bar":
            var v12 bool = (!(false || (!false)))
            _ = v12
        case "a":
            var v13 uint64 = v1
            _ = v13
        case "\n\t":
            v9 -= v9
        default:
            fmt.Println(v9)
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 string = ("\n\t" + "")
            v4 = (v4 ^ v9)
            if true {
                v2 = (v2 / (((v2 & v2)) | 1))
                v4 -= v9
                fmt.Println(v9, v2, fi14)
            }
            var v16 int8 = -1
            if true {
                if (!(false || true)) {
                    v15 = (v15 + (v15 + (((("\n\t" + v15) + "") + "") + ("\n\t" + v15))))
                    switch v15 {
                    case "\n\t":
                        switch -1 {
                        case 9223372036854775807:
                            fmt.Println(v9, v4)
                        case -9223372036854775808:
                            v17 := GenS2{F0: 1}
                            _ = v17
                        }
                    case "hello":
                        if (v16 != v16) {
                            v9 += (v9 - (v9 >> ((65535) & 63)))
                        } else {
                            v2 = v2
                        }
                    }
                }
                switch (^(fi14)) {
                case 4294967295:
                    v18 := GenS3{F0: -32768}
                    _ = v18
                case 32768:
                    if ((v1 & (v1 + (v1 | (v1 >> ((2) & 63))))) <= uint64(v2)) {
                        v19 := GenS3{F0: 1}
                        _ = v19
                    } else {
                        v1 -= 1
                        var v20 int8 = v16
                        if (false && (float32((-2.25 - 1.5)) < 3.14159)) {
                            v2 += v2
                        } else {
                            var v21 int = (fi14 * (fi14 & (-(fi14))))
                            _ = v21
                        }
                        _ = v20
                    }
                default:
                    var v22 int = (fi11 | 315)
                    _ = v22
                }
            } else {
                var v23 float32 = ((1.5 + 0.0) + 3.14159)
                if false {
                    var v24 uint8 = 2
                    _ = v24
                } else {
                    var v25 float64 = 1e10
                    _ = v25
                }
                _ = v23
            }
            var v26 int64 = -9223372036854775808
            _ = fi14
            _ = v15
            _ = v16
            _ = v26
        }
        _ = fi11
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v28 := GenS4{F0: 0}
        var v29 uint8 = 140
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 string = "\n\t"
    if false {
        var v31 int16 = 1000
        _ = v31
    }
    var v32 int16 = 768
    v2 -= v2
    var v33 bool = false
    var v34 int64 = 9223372036854775807
    delete(v3, 1)
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v35 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v9
}
