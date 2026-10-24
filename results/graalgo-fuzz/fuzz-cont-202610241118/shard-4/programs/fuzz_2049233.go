package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 bool; F2 uint }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 string, p1 uint, p2 int) int64 {
    p2 = (p2 | (p2 * p2))
    v6 := GenS1{F0: -2147483648, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1(p0 int16, p1 uint64, p2 int64) string {
    v7 := GenS2{F0: p2}
    v8 := GenS2{F0: (p2 + p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 int16 = -1
    var v2 int32 = 0
    v3 := []uint32{}
    var v4 int32 = -2147483648
    v5 := make(map[uint32]uint32)
    fmt.Println(v2, v1)
    var v9 uint16 = 65535
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(v1)
        var v11 bool = ("" == "hello")
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int8 = -89
        v9 -= v9
        v1 -= (v1 << ((18446744073709551615) & 63))
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 352, F1: v2}
    v15 := GenS3{F0: -108, F1: false, F2: 1}
    for fi16 := 0; fi16 < 8; fi16++ {
        if (1 >= v1) {
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 uint8 = 2
                switch ("hello" + "foo bar") {
                case "hello":
                    fmt.Println(fi17, v18, fi16)
                }
                var v19 uint64 = 2
                if (0.0 >= 0.1) {
                    switch "foo bar" {
                    case "foo bar":
                        v20 := GenS3{F0: 0, F1: (false && false), F2: 935}
                        _ = v20
                    case "":
                        var v21 bool = false
                        _ = v21
                    case "hello":
                        v3 = append(v3, 2)
                    default:
                        var v22 string = ("hello" + "a")
                        _ = v22
                    }
                    delete(v5, 382)
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                    case "hello":
                        v23 := GenS4{F0: v1, F1: -1}
                        v24 := GenS5{F0: v19}
                        v18 = (^(v18))
                        _ = v23
                        _ = v24
                    case "foo bar":
                        v4 = v2
                    default:
                        v25 := v5[65535]
                        _ = v25
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            for fi26 := 0; fi26 < 7; fi26++ {
                if (-48 < 120) {
                    v27 := GenS1{F0: v4, F1: (v4 | 898)}
                    switch "" {
                    case "":
                        v1 += (v1 >> ((4294967295) & 63))
                    case "a":
                        if (!true) {
                            var v28 string = ""
                            v2 -= 2147483647
                            v2 -= v2
                            _ = v28
                        } else {
                            var v29 string = "foo bar"
                            fmt.Println(v29, v9)
                            fmt.Println(v29)
                            _ = v29
                        }
                    case "foo bar":
                        delete(v5, 2147483647)
                        fmt.Println(v9, v2, v9)
                    }
                    v30 := v5[uint32(v9)]
                    _ = v27
                    _ = v30
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v31 float64 = 0.0
                var v32 uint64 = 18446744073709551615
                fmt.Println(fi16)
                v2 = 1
                _ = fi26
                _ = v31
                _ = v32
            }
        }
        _ = fi16
    }
    if (false && true) {
        var v33 bool = false
        for fi34 := 0; fi34 < 2; fi34++ {
            fmt.Println(v9)
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 uint64 = 119
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
        _ = v33
    } else {
        fmt.Println(v9)
    }
    if false {
        if (!(0 >= 2)) {
            v37 := v5[707]
            fmt.Println(v37, v1)
            var v38 int8 = 127
            _ = v37
            _ = v38
        } else {
            if len(v3) > 0 {
                v3[0] = 1
            }
            if (v1 < (v1 * v1)) {
                var v39 int64 = -1
                _ = v39
            }
        }
        v5[702] = 65535
        fmt.Println(v4, v4)
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
