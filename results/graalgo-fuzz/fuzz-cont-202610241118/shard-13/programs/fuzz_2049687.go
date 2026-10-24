package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 string; F1 string; F2 int }
type GenS4 struct { F0 uint; F1 int16; F2 string }
func genFn0() int16 {
    v4 := []int{}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (true || true) {
        v5 := GenS1{F0: 32767, F1: 32767}
        var v6 int32 = 0
        _ = v5
        _ = v6
    } else {
        var v7 int = 32768
        _ = v7
    }
    _ = v4
    return -32768
}
func genFn1(p0 uint64, p1 uint8, p2 bool) int64 {
    p1 = p1
    if (-388 == -9223372036854775808) {
        v8 := []int{-1, -9223372036854775808, 9223372036854775807, -9223372036854775808}
        _ = v8
    }
    switch "foo bar" {
    case "foo bar":
        var v9 uint16 = 65535
        _ = v9
    case "\n\t":
        p0 += 2
    case "a":
        var v10 float64 = 0.0
        _ = v10
    }
    v11 := []int{128, 4294967295, 2147483647}
    if ((^(p0)) > (p0 - 18446744073709551615)) {
        p2 = true
        var v12 int = 4294967295
        _ = v12
    } else {
        v13 := GenS2{F0: uint32(p0), F1: 139}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 9223372036854775807
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 229
    var v3 float32 = -2.25
    v1 = v1
    v14 := []int{-468, -32, 4294967295, -9223372036854775808}
    var v15 uint64 = 42
    v16 := GenS3{F0: ("hello" + "hello"), F1: ("" + "a"), F2: 32768}
    if (uint(v15) != 4294967295) {
        var v17 int8 = -1
        _ = v17
    }
    if (true && (!false)) {
        var v18 int32 = (v1 / ((0) | 1))
        _ = v18
    } else {
        if false {
            switch "" {
            case "a":
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            case "hello":
                v15 += 42
            case "foo bar":
                var v19 string = "\n\t"
                _ = v19
            default:
                if false {
                    var v20 float64 = 0.0
                    v1 = 0
                    _ = v20
                }
            }
        }
    }
    v21 := GenS4{F0: 65535, F1: 42, F2: ""}
    switch 9223372036854775807 {
    case 42:
        if true {
            if (false && (!true)) {
                for fi22 := 0; fi22 < 9; fi22++ {
                    switch fi22 {
                    case -923:
                        var v23 uint32 = 2
                        _ = v23
                    case 9223372036854775807:
                        if ((v15 >> ((1) & 63)) == 312) {
                            v24 := GenS4{F0: 0, F1: genFn0(), F2: "a"}
                            v25 := GenS2{F0: 645, F1: uint32(v2)}
                            switch 32767 {
                            case 4294967295:
                                if len(v14) > 0 {
                                    v14[0] = fi22
                                }
                            case -371:
                                switch fi22 {
                                case -1:
                                    fmt.Println(v1, v15)
                                }
                            }
                            _ = v24
                            _ = v25
                        } else {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            var v26 uint16 = 224
                            _ = v26
                        }
                    default:
                        if false {
                            var v27 int8 = 0
                            _ = v27
                        }
                    }
                    _ = fi22
                }
            } else {
                var v28 int = 32767
                _ = v28
            }
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = len(v14)
        }
    case -683:
        v15 += 4294967295
    }
    var v29 uint16 = 42
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 int64 = int64(fi30)
        v29 = v29
        switch "foo bar" {
        case "hello":
            var v32 int = (fi30 * (^(fi30)))
            _ = v32
        case "\n\t":
            if len(v14) > 0 {
                v14[0] = (fi30 << ((4294967295) & 63))
            }
        case "":
            if len(v14) > 0 {
                v14[0] = fi30
            }
        }
        _ = fi30
        _ = v31
    }
    var v33 int32 = (v1 << ((966) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v33
}
