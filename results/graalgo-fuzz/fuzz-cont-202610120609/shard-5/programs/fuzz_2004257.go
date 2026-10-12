package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 int; F2 bool }
func genFn0(p0 uint32, p1 int16) bool {
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    var v6 uint16 = 65535
    _ = p0
    _ = p1
    _ = v6
    return (65535 <= 0)
}
func genFn1() int64 {
    var v7 float32 = 3.14159
    var v8 int8 = 127
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{236}
    v3 := make(map[int]bool)
    v4 := []uint32{42, 4294967295, 42}
    if len(v2) > 0 {
        v2[0] = 100
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        switch "foo bar" {
        case "foo bar":
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 int32 = 1
                v11 += v11
                _ = fi10
                _ = v11
            }
        case "hello":
            var v12 int8 = -1
            _ = v12
        default:
            var v13 int32 = -1
            _ = v13
        }
        fmt.Println(fi9)
        _ = fi9
    }
    var v14 float32 = (-2.25 / 0.1)
    v15 := GenS1{F0: 98}
    var v16 uint8 = 228
    if false {
        var v17 float32 = v14
        _ = v17
    }
    if (1 != 65535) {
        if (true || (false && false)) {
            if false {
                var v18 int8 = 1
                _ = v18
            } else {
                v19 := GenS1{F0: 0}
                for fi20 := 0; fi20 < 11; fi20++ {
                    v16 -= (v16 ^ 128)
                    var v21 float32 = ((((v14 - v14) - float32(1e10)) - v14) + (0.1 * (((1e10 - 0.1) - v14) * (((-2.25 * v14) + (v14 - v14)) - 0.0))))
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
            v16 = (v16 % ((v16) | 1))
        }
    } else {
        v22 := GenS2{F0: 65535, F1: (-86 <= 64)}
        for fi23 := 0; fi23 < 20; fi23++ {
            if (float32(1e10) > float32(((float64(v14) * 0.0) * ((((1e10 - 0.0) - 0.1) + 1e10) - 1.5)))) {
                for fi24 := 0; fi24 < 16; fi24++ {
                    v25 := GenS3{F0: 0, F1: 127, F2: (v16 <= v16)}
                    _ = fi24
                    _ = v25
                }
                switch ("" + "") {
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "foo bar":
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v16 = (v16 >> ((0) & 63))
                        var v27 int = fi23
                        var v28 int32 = -1
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                default:
                    fmt.Println(v14)
                }
            } else {
                var v29 int8 = -1
                if false {
                    switch fi23 {
                    case 128:
                        fmt.Println(fi23, v14, v29)
                    case 2147483647:
                        for fi30 := 0; fi30 < 11; fi30++ {
                            var v31 uint8 = v16
                            _ = fi30
                            _ = v31
                        }
                    default:
                        v32 := GenS1{F0: v29}
                        _ = v32
                    }
                    fmt.Println(v14)
                } else {
                    v16 -= v16
                    fmt.Println(fi23)
                    var v33 int16 = 1
                    _ = v33
                }
                switch "\n\t" {
                case "foo bar":
                    var v34 string = (("hello" + ("" + ("foo bar" + "\n\t"))) + ("\n\t" + "a"))
                    _ = v34
                case "":
                    var v35 string = ""
                    _ = v35
                case "hello":
                    v16 -= v16
                default:
                    v3[(fi23 % ((fi23) | 1))] = (true || ((v16 % (((v16 ^ 237)) | 1)) == v16))
                }
                _ = v29
            }
            if false {
                fmt.Println(v16, fi23, v16)
                v4 = append(v4, 689)
                v36 := GenS2{F0: 895, F1: true}
                _ = v36
            }
            _ = fi23
        }
        _ = v22
    }
    var v37 uint32 = 65535
    v14 += v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
