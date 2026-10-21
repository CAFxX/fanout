package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 uint16; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16 }
func genFn0() int64 {
    var v5 uint = 284
    var v6 int64 = 1
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 bool) int32 {
    p0 = (true && (false || (!p0)))
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = (("a" != "a") && p0)
        _ = fi7
    }
    _ = p0
    return 9
}

func main() {
    v1 := []string{("hello" + ("" + ("foo bar" + ""))), ("" + "")}
    var v2 int16 = -1
    v3 := make(map[int]string)
    var v4 uint64 = 0
    var v8 uint16 = uint16(v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 := GenS1{F0: 0, F1: 42, F2: (-1 >= 42)}
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(v8, fi10)
        var v11 uint8 = 235
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "\n\t":
        if false {
            v12 := GenS2{F0: (true || true), F1: v2}
            for fi13 := 0; fi13 < 12; fi13++ {
                switch (fi13 + (^(fi13))) {
                case -515:
                    v8 -= (v8 & v8)
                case 32768:
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                default:
                    for fi14 := 0; fi14 < 14; fi14++ {
                        switch 128 {
                        case 354:
                            switch (fi13 % ((fi14) | 1)) {
                            case 2147483647:
                                v15 := GenS3{F0: 2147483647, F1: v8, F2: -1}
                                _ = v15
                            case -2147483648:
                                v16 := GenS2{F0: (0 != 255), F1: v2}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v3[128] = "foo bar"
                                _ = v16
                            default:
                                var v17 bool = (float32(float64(-2.25)) >= (1.5 - -2.25))
                                _ = v17
                            }
                        case -573:
                            var v18 int = (fi14 / (((fi13 << ((1) & 63))) | 1))
                            _ = v18
                        default:
                            switch "" {
                            case "\n\t":
                                v1 = append(v1, (("\n\t" + ("foo bar" + ("a" + "\n\t"))) + "\n\t"))
                            default:
                                var v19 uint64 = (v4 % (((v4 | 1)) | 1))
                                _ = v19
                            }
                        }
                        var v20 uint64 = 492
                        fmt.Println(v2)
                        _ = fi14
                        _ = v20
                    }
                }
                var v21 int16 = 1000
                _ = fi13
                _ = v21
            }
            _ = v12
        } else {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v8, v2)
            }
        }
    case "":
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 float64 = 1.5
            fmt.Println(v2, fi22)
            switch 42 {
            case 127:
                if false {
                    v2 = 32767
                    delete(v3, fi22)
                    v24 := GenS4{F0: v2}
                    _ = v24
                }
            default:
                v2 = v2
            }
            _ = fi22
            _ = v23
        }
    }
    v4 = v4
    switch "a" {
    case "foo bar":
        if true {
            v4 = v4
            fmt.Println(v8, v2)
            delete(v3, 1)
        } else {
            var v25 int = 42
            v2 -= v2
            _ = v25
        }
    case "hello":
        v26 := GenS5{F0: 32767}
        _ = v26
    default:
        var v27 string = ("" + "\n\t")
        _ = v27
    }
    v8 -= 0
    if false {
        var v28 uint32 = uint32(v4)
        fmt.Println(v8, v8, v4)
        _ = v28
    } else {
        var v29 uint = 255
        var v30 float64 = (0.1 * ((0.1 + ((-2.25 / 0.1) * 3.14159)) - 0.0))
        v4 -= (v4 << ((v29) & 63))
        _ = v29
        _ = v30
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v31 float64 = 3.14159
    v8 = (v8 & 0)
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 float64 = 0.0
        var v34 int8 = 1
        if ((v2 * (v2 ^ v2)) > 777) {
            fmt.Println(v4, v31, fi32)
            _ = len(v1)
        } else {
            switch (-(fi32)) {
            case 32768:
                var v35 int8 = 100
                _ = v35
            case 2147483647:
                v36 := v3[(fi32 >> ((0) & 63))]
                _ = v36
            default:
                v37 := GenS2{F0: (!((v34 | (^(v34))) <= int8(v4))), F1: (v2 % (((v2 ^ (^(v2)))) | 1))}
                _ = v37
            }
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println(v8, v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v8
    _ = v9
}
