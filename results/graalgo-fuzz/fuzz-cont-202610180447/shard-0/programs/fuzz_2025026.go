package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0(p0 int) int16 {
    if (!("foo bar" == "hello")) {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{(fi7 >> ((42) & 63)), (fi7 * (fi7 ^ fi7))}
            for fi9 := 0; fi9 < 1; fi9++ {
                p0 = (p0 - fi7)
                _ = fi9
            }
            v10 := GenS1{F0: 32768, F1: 255}
            _ = fi7
            _ = v10
            _ = v8
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: fi11, F1: 571}
            _ = fi11
            _ = v12
        }
        v13 := []int{p0, (p0 / ((127) | 1))}
        _ = v13
    } else {
        v14 := GenS1{F0: p0, F1: 65535}
        _ = v14
    }
    v15 := GenS2{F0: 637, F1: 1}
    _ = p0
    _ = v15
    return -973
}

func main() {
    var v1 int16 = -821
    var v2 int64 = 611
    v3 := []uint32{612}
    var v4 uint8 = 210
    v5 := make(map[int]int)
    var v6 uint64 = 1
    var v16 uint64 = (v6 * 674)
    if (42 < 0) {
        v17 := GenS1{F0: -426, F1: uint16(v16)}
        v4 -= (v4 & v4)
        v6 -= v16
        _ = v17
    }
    var v18 uint8 = 72
    fmt.Println(v18, v6)
    v16 -= 626
    fmt.Println(v2, v18)
    fmt.Println(v1, v4, v2)
    switch ("\n\t" + "a") {
    case "":
        if true {
            var v19 int32 = 42
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(v1)
                if (v6 < 4294967295) {
                    v2 += 586
                }
                _ = fi20
            }
            var v21 float64 = (3.14159 * 0.0)
            _ = v19
            _ = v21
        } else {
            v2 -= v2
            var v22 int = -2147483648
            for fi23 := 0; fi23 < 12; fi23++ {
                v18 = v4
                switch v22 {
                case -18:
                    var v24 uint32 = 4294967295
                    _ = v24
                case 128:
                    if (("a" + ("hello" + "")) == (("hello" + "hello") + "foo bar")) {
                        for fi25 := 0; fi25 < 14; fi25++ {
                            v1 -= int16(v16)
                            fmt.Println(v1)
                            switch 128 {
                            case 2147483647:
                                var v26 int32 = int32(v2)
                                _ = v26
                            default:
                                v1 = (-(v1))
                            }
                            _ = fi25
                        }
                        switch "" {
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = 859
                            }
                        case "":
                            _ = len(v3)
                        }
                    } else {
                        fmt.Println(v4, v16)
                        var v27 uint8 = v18
                        var v28 int64 = (v2 % (((v2 >> ((4294967295) & 63))) | 1))
                        _ = v27
                        _ = v28
                    }
                case 32767:
                    if (!false) {
                        var v29 int8 = -128
                        _ = v29
                    } else {
                        v22 = fi23
                        fmt.Println(v2)
                        delete(v5, fi23)
                    }
                default:
                    fmt.Println(v6)
                }
                switch fi23 {
                case 2147483647:
                    for fi30 := 0; fi30 < 6; fi30++ {
                        fmt.Println(v6)
                        _ = fi30
                    }
                case -2147483648:
                    _ = len(v5)
                case 42:
                    v6 = 1
                default:
                    if true {
                        v16 += (v16 + (v6 + 2))
                    }
                }
                _ = fi23
            }
            _ = v22
        }
    case "foo bar":
        v2 = (v2 | -875)
    default:
        v16 += (v6 + (^(v16)))
    }
    v4 -= 42
    v4 = v4
    switch ("hello" + ("hello" + (("\n\t" + (("" + "hello") + ("" + "a"))) + "a"))) {
    case "\n\t":
        var v31 float32 = float32(3.14159)
        _ = v31
    case "hello":
        v2 = 248
    default:
        v1 = (^(v1))
    }
    v32 := GenS3{F0: 32767, F1: v1}
    if len(v3) > 0 {
        v3[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
