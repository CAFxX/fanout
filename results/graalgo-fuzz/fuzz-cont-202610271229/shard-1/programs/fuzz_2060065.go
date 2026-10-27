package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint8) uint16 {
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{(fi7 * 1), fi7, (fi7 ^ fi7), -861}
        _ = fi7
        _ = v8
    }
    _ = p0
    return 65535
}

func main() {
    v1 := []uint32{98}
    var v2 uint32 = 42
    var v3 int64 = 42
    v4 := make(map[string]int)
    var v5 float32 = 1e10
    for fi9 := 0; fi9 < 7; fi9++ {
        switch fi9 {
        case -9223372036854775808:
            v5 = 3.14159
        default:
            v10 := GenS1{F0: 134, F1: 1, F2: fi9}
            _ = v10
        }
        fmt.Println(v2, fi9)
        if true {
            if ((!true) && true) {
                _ = len(v1)
            }
            switch "hello" {
            case "":
                delete(v4, "foo bar")
            default:
                switch 32768 {
                case 32768:
                    switch (("" + "") + "\n\t") {
                    case "hello":
                        fmt.Println(v2, fi9, fi9)
                    }
                default:
                    switch "a" {
                    case "foo bar":
                        switch "\n\t" {
                        case "hello":
                            switch (fi9 >> ((4294967295) & 63)) {
                            case 32768:
                                fmt.Println(v2, v2)
                            case 128:
                                var v11 bool = (false || false)
                                _ = v11
                            case 517:
                                if len(v1) > 0 {
                                    v1[0] = v2
                                }
                            default:
                                switch "\n\t" {
                                case "":
                                    for fi12 := 0; fi12 < 3; fi12++ {
                                        var v13 uint = 227
                                        _ = fi12
                                        _ = v13
                                    }
                                case "\n\t":
                                    fmt.Println(v3, v3, v2)
                                default:
                                    for fi14 := 0; fi14 < 17; fi14++ {
                                        v3 = v3
                                        _ = fi14
                                    }
                                }
                            }
                        case "":
                            var v15 uint8 = 128
                            _ = v15
                        }
                    }
                }
            }
            if false {
                var v16 string = ((("\n\t" + "\n\t") + ("a" + ("hello" + ("\n\t" + ("\n\t" + "\n\t"))))) + "")
                v3 = v3
                v16 = (v16 + v16)
                _ = v16
            } else {
                fmt.Println(v5, v2, v5)
                switch fi9 {
                case 127:
                    v2 += (v2 & (^(v2)))
                case 2147483647:
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v18 := GenS1{F0: 42, F1: 42, F2: -631}
                        v19 := GenS1{F0: 0, F1: 42, F2: (fi17 >> ((0) & 63))}
                        v2 = (v2 - (v2 * (v2 | 412)))
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                v1 = append(v1, v2)
            }
        } else {
            var v20 int64 = 42
            var v21 uint32 = (^(v2))
            _ = v20
            _ = v21
        }
        _ = fi9
    }
    var v22 float64 = (float64(v5) - 1.5)
    v2 += (v2 ^ (v2 ^ 0))
    switch -9223372036854775808 {
    case 32767:
        var v23 float32 = v5
        _ = v23
    case -2147483648:
        if (!true) {
            var v24 int8 = 0
            _ = v24
        } else {
            v25 := GenS2{F0: (true || true), F1: 42}
            var v26 int8 = 127
            _ = v25
            _ = v26
        }
    case 127:
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 uint32 = v2
            _ = fi27
            _ = v28
        }
    default:
        var v29 float32 = ((v5 * 0.1) - (0.0 * (v5 + v5)))
        _ = v29
    }
    v5 = v5
    if len(v1) > 0 {
        v1[0] = (v2 % ((4294967295) | 1))
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v5 -= v5
        _ = fi30
    }
    fmt.Println(v3)
    var v31 float32 = v5
    var v32 uint16 = 402
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
