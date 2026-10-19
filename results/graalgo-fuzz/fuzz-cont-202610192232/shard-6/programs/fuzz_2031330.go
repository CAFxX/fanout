package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
func genFn0() uint32 {
    var v5 int = -2147483648
    switch "" {
    case "\n\t":
        var v6 int = (v5 % (((v5 & v5)) | 1))
        _ = v6
    case "a":
        v7 := []int{(v5 + (v5 + -1)), v5}
        _ = v7
    case "hello":
        v5 -= 2147483647
    default:
        v8 := []int{(v5 & 9223372036854775807)}
        _ = v8
    }
    var v9 float32 = (1e10 / 0.1)
    _ = v5
    _ = v9
    return 2147483647
}

func main() {
    var v1 float32 = (float32(((1.5 - -2.25) * 0.0)) * -2.25)
    v2 := []int{127, 640, -597}
    v3 := []uint32{}
    var v4 int32 = 207
    fmt.Println(v1, v1)
    fmt.Println(v4, v1)
    switch ("hello" + "a") {
    case "a":
        var v10 uint8 = 140
        _ = v10
    case "":
        var v11 uint64 = 18446744073709551615
        _ = v11
    default:
        var v12 int32 = -1
        _ = v12
    }
    var v13 int8 = 57
    var v14 int8 = (v13 >> ((207) & 63))
    for fi15 := 0; fi15 < 8; fi15++ {
        v13 -= v13
        _ = fi15
    }
    v16 := GenS1{F0: -21, F1: "hello"}
    if len(v2) > 0 {
        v2[0] = 443
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int64 = 256
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 float64 = (1.5 + (0.0 - 0.1))
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v13, v13)
    for fi21 := 0; fi21 < 10; fi21++ {
        switch -1 {
        case -649:
            for fi22 := 0; fi22 < 4; fi22++ {
                if true {
                    v23 := GenS1{F0: 1, F1: ("\n\t" + "\n\t")}
                    v14 += (v13 - (v14 >> ((629) & 63)))
                    switch (fi21 % ((fi21) | 1)) {
                    case 32767:
                        var v24 int64 = -529
                        _ = v24
                    default:
                        var v25 int = (fi21 / (((fi21 * fi21)) | 1))
                        _ = v25
                    }
                    var v26 uint16 = 32767
                    if (!true) {
                        v4 = 2147483647
                    }
                    _ = v23
                    _ = v26
                } else {
                    var v27 bool = (0 != 42)
                    v3 = append(v3, 1)
                    switch "" {
                    case "a":
                        v28 := GenS1{F0: -1, F1: ("\n\t" + "\n\t")}
                        _ = v28
                    default:
                        _ = len(v3)
                        var v29 int64 = 9223372036854775807
                        if len(v2) > 0 {
                            v2[0] = fi22
                        }
                        _ = v29
                    }
                    if v27 {
                        var v30 int = -977
                        _ = v30
                    }
                    switch (^(fi21)) {
                    case 127:
                        var v31 bool = false
                        _ = v31
                    case -1:
                        fmt.Println(v13, v27, fi21)
                    default:
                        var v32 bool = true
                        _ = v32
                    }
                    v13 = v14
                    if len(v2) > 0 {
                        v2[0] = fi21
                    }
                    _ = v27
                }
                v13 = (v13 & (^(v14)))
                v33 := GenS2{F0: true}
                switch "\n\t" {
                case "a":
                    v34 := GenS3{F0: -810}
                    _ = v34
                default:
                    if (!(65535 == 844)) {
                        var v35 uint32 = 42
                        _ = v35
                    } else {
                        var v36 uint16 = 32767
                        v36 = v36
                        _ = v36
                    }
                }
                _ = fi22
                _ = v33
            }
        default:
            var v37 bool = (2147483647 != 42)
            _ = v37
        }
        var v38 int64 = -1
        _ = fi21
        _ = v38
    }
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 int64 = 151
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
