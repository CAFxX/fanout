package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint64 }
func genFn0(p0 int64, p1 int, p2 uint8) int32 {
    var v5 string = "hello"
    var v6 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 int32, p1 int64) int16 {
    var v7 int8 = 42
    switch 32767 {
    case 221:
        var v8 float64 = (0.1 + ((1e10 + 1.5) / (1.5 + 1.5)))
        _ = v8
    }
    v9 := GenS1{F0: -9223372036854775808, F1: 491}
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 1
}

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{65535, 2147483648}
    var v3 int32 = 182
    var v4 int16 = 1
    v10 := GenS1{F0: 770, F1: v4}
    switch "hello" {
    case "hello":
        var v11 uint8 = 255
        _ = v11
    case "a":
        switch 0 {
        case 42:
            if true {
                var v12 int32 = -2147483648
                _ = v12
            } else {
                v13 := GenS2{F0: 65535, F1: 255}
                var v14 int32 = (^(v3))
                _ = v13
                _ = v14
            }
        case -1:
            for fi15 := 0; fi15 < 10; fi15++ {
                switch (("a" + ("foo bar" + "hello")) + "foo bar") {
                case "foo bar":
                    for fi16 := 0; fi16 < 18; fi16++ {
                        v3 += v3
                        var v17 int = fi15
                        if false {
                            var v18 float32 = v1
                            fmt.Println(v3)
                            _ = v18
                        }
                        v19 := GenS3{F0: 42, F1: fi15}
                        _ = fi16
                        _ = v17
                        _ = v19
                    }
                case "a":
                    var v20 uint8 = 42
                    _ = v20
                }
                _ = fi15
            }
        default:
            var v21 int = 0
            _ = v21
        }
    }
    var v22 string = ""
    v23 := GenS3{F0: 23, F1: 32767}
    v4 = 34
    if ((1.5 / (1e10 + 0.0)) > (1.5 / 1.0)) {
        v1 = v1
        switch 9223372036854775807 {
        case 745:
            v2 = append(v2, 4294967295)
        default:
            var v24 uint = 2
            _ = v24
        }
        if (1 != 42) {
            if (true || (false && (!true))) {
                v4 -= (v4 * (v4 << ((992) & 63)))
            } else {
                v25 := GenS3{F0: -1, F1: -2147483648}
                switch 901 {
                case 128:
                    switch v22 {
                    case "\n\t":
                        v1 = v1
                    case "foo bar":
                        switch 32768 {
                        case 0:
                            var v26 uint8 = 255
                            _ = v26
                        case 9223372036854775807:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            if (2 <= 604) {
                                var v27 float32 = v1
                                _ = v27
                            }
                        }
                    case "a":
                        var v28 float32 = 1.5
                        _ = v28
                    default:
                        v29 := GenS4{F0: 2147483647, F1: 0}
                        _ = v29
                    }
                }
                _ = v25
            }
            for fi30 := 0; fi30 < 18; fi30++ {
                _ = len(v2)
                var v31 bool = (42 == 9223372036854775807)
                _ = fi30
                _ = v31
            }
        }
    }
    v22 = "a"
    switch -1 {
    case -727:
        if true {
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            v22 = v22
            v22 = v22
        } else {
            v1 += v1
            v22 = "hello"
        }
    case 32767:
        v32 := GenS5{F0: (v3 >= (-(v3))), F1: 2}
        _ = v32
    case -2147483648:
        for fi33 := 0; fi33 < 6; fi33++ {
            v1 += v1
            for fi34 := 0; fi34 < 16; fi34++ {
                v4 -= v4
                var v35 int32 = genFn0(9223372036854775807, 9223372036854775807, 0)
                v4 += v4
                _ = fi34
                _ = v35
            }
            _ = fi33
        }
    }
    var v36 int64 = 9223372036854775807
    var v37 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
