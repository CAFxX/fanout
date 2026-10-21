package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 string; F2 int }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int = fi4
        v6 := GenS1{F0: 19}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 float32 = 0.0
    _ = v7
    return 1
}
func genFn1(p0 int) uint16 {
    p0 += p0
    p0 = (p0 % ((-2147483648) | 1))
    var v8 int = -58
    v9 := GenS1{F0: 2147483647}
    _ = p0
    _ = v8
    _ = v9
    return 756
}

func main() {
    v1 := []uint32{4294967295, 539, 2147483648, 2}
    var v2 int16 = 42
    var v3 int8 = 0
    if false {
        for fi10 := 0; fi10 < 9; fi10++ {
            v3 = v3
            _ = fi10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            switch "foo bar" {
            case "":
                v3 += (^(v3))
            case "hello":
                var v12 string = ""
                _ = v12
            default:
                var v13 uint64 = 2
                _ = v13
            }
            var v14 uint = 2
            v3 += -106
            _ = fi11
            _ = v14
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v3 = (v3 >> ((18446744073709551615) & 63))
        var v15 float32 = (3.14159 / 1.0)
        _ = v15
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v2, v3)
    default:
        var v16 bool = ((false && true) || (float64(0.1) > 1e10))
        _ = v16
    }
    var v17 int64 = int64(v3)
    switch ("hello" + ("hello" + "hello")) {
    case "\n\t":
        var v18 float32 = float32((((float64((float32(0.1) * (0.0 * float32(-2.25)))) / (0.1 * -2.25)) + 1.5) + 0.1))
        _ = v18
    }
    var v19 int32 = -2147483648
    var v20 string = "foo bar"
    if (true || (!false)) {
        v3 += v3
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        v3 += (v3 % (((v3 | -128)) | 1))
        if true {
            v19 -= 1
            switch "" {
            case "":
                for fi22 := 0; fi22 < 14; fi22++ {
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                    if (!false) {
                        var v23 int16 = 32767
                        var v24 int8 = v3
                        _ = v23
                        _ = v24
                    }
                    var v25 uint16 = 913
                    _ = fi22
                    _ = v25
                }
            default:
                switch fi21 {
                case 127:
                    var v26 int8 = 42
                    _ = v26
                case 1:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -1:
                    v3 = 127
                default:
                    var v27 uint16 = 295
                    _ = v27
                }
            }
        } else {
            fmt.Println(v17)
            v3 += v3
            if (!true) {
                fmt.Println(v2, fi21)
            } else {
                for fi28 := 0; fi28 < 1; fi28++ {
                    v29 := GenS2{F0: v20, F1: v20, F2: (fi21 | fi21)}
                    _ = fi28
                    _ = v29
                }
            }
        }
        v30 := GenS3{F0: 32767, F1: false}
        _ = fi21
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
