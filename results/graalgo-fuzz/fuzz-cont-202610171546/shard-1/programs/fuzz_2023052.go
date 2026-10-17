package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 int16) uint64 {
    fmt.Println(p0)
    var v5 uint32 = 839
    v6 := GenS1{F0: -1, F1: 127}
    switch ("hello" + (("" + "") + ("hello" + "a"))) {
    case "\n\t":
        p0 = 32767
    case "":
        if (1 <= int32(p0)) {
            p0 = p0
            var v7 bool = (!((v5 / (((v5 ^ 42)) | 1)) < 2))
            _ = v7
        } else {
            if (1 < 421) {
                v5 = 4
                v5 -= v5
                var v8 uint16 = 32767
                _ = v8
            }
            fmt.Println(p0, v5)
        }
    case "hello":
        p0 -= (p0 ^ (-(p0)))
    default:
        var v9 int = -596
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    return 330
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -1
    v3 := []uint32{}
    var v4 uint16 = 1
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 string = "\n\t"
        _ = fi10
        _ = v11
    }
    if (!true) {
        v2 = (v2 | v2)
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint8 = 0
            _ = len(v1)
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi14
        }
    }
    v3 = append(v3, 687)
    var v15 uint64 = 0
    v15 = (v15 % ((uint64(v4)) | 1))
    switch 711 {
    case 9223372036854775807:
        v2 += -2147483648
    default:
        v16 := GenS1{F0: 32767, F1: -68}
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        if (fi17 < 32767) {
            var v18 int32 = 73
            var v19 uint = 42
            _ = v18
            _ = v19
        }
        v4 -= (v4 | (v4 >> ((691) & 63)))
        var v20 int8 = 0
        _ = fi17
        _ = v20
    }
    if (v4 <= v4) {
        for fi21 := 0; fi21 < 19; fi21++ {
            _ = len(v3)
            if false {
                fmt.Println(v2, v2)
                var v22 uint8 = 255
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v22
            }
            v23 := GenS1{F0: fi21, F1: 48}
            _ = fi21
            _ = v23
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            if ((!(1 == 65535)) && (((("hello" + "hello") + "hello") + "hello") >= "a")) {
                switch "hello" {
                case "foo bar":
                    v15 = v15
                case "":
                    if len(v3) > 0 {
                        v3[0] = 708
                    }
                }
            } else {
                v25 := GenS2{F0: v15, F1: 42}
                _ = v25
            }
            var v26 int64 = -128
            var v27 string = ("foo bar" + "hello")
            _ = fi24
            _ = v26
            _ = v27
        }
        switch 32768 {
        case 32767:
            v28 := GenS1{F0: 2147483647, F1: 127}
            _ = v28
        default:
            v29 := GenS3{F0: (^(v2))}
            _ = v29
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        v4 = 255
        var v31 int32 = -1
        _ = fi30
        _ = v31
    }
    fmt.Println(v15)
    var v32 uint64 = (v15 & (v15 + v15))
    var v33 uint = 2
    if ((42 <= 2147483648) && true) {
        var v34 uint64 = 18446744073709551615
        _ = v34
    }
    if len(v3) > 0 {
        v3[0] = uint32(v33)
    }
    if ((-92 == 100) || ((v33 ^ (v33 / ((0) | 1))) == 4294967295)) {
        v35 := GenS3{F0: (v2 ^ (-(v2)))}
        var v36 uint16 = v4
        v15 = v15
        _ = v35
        _ = v36
    }
    var v37 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v4
}
