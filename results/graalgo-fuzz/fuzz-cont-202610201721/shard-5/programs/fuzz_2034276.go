package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint16) uint8 {
    p0 -= p0
    v6 := GenS1{F0: 0, F1: 0}
    _ = p0
    _ = v6
    return 2
}
func genFn1() int32 {
    if (-128 >= 0) {
        if true {
            v7 := GenS2{F0: 660, F1: -32768}
            var v8 int32 = -705
            v9 := GenS1{F0: -243, F1: (v8 >> ((966) & 63))}
            _ = v7
            _ = v8
            _ = v9
        } else {
            v10 := []int{-2147483648, -136, -71, 128}
            _ = v10
        }
        if false {
            var v11 bool = (2 >= 2)
            _ = v11
        } else {
            var v12 uint16 = 0
            var v13 int16 = 42
            _ = v12
            _ = v13
        }
    }
    v14 := GenS1{F0: 32768, F1: 0}
    var v15 float64 = -2.25
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{v1, (v1 + 2), 2}
    var v4 int = 32768
    v5 := []uint32{v1, 2147483648, (v1 ^ uint32(v2)), (v1 + (^(v1)))}
    if true {
        for fi16 := 0; fi16 < 13; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                fmt.Println(v4, v4)
                _ = fi17
            }
            v2 += v2
            v2 -= (v2 % ((1) | 1))
            _ = fi16
        }
    } else {
        switch 0 {
        case 42:
            var v18 int64 = int64(v1)
            _ = v18
        case 0:
            _ = len(v5)
        default:
            fmt.Println(v2, v4, v1)
        }
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int = (^(fi19))
        _ = fi19
        _ = v20
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "" {
    case "foo bar":
        var v21 uint8 = 255
        _ = v21
    default:
        var v22 uint64 = (v2 % ((v2) | 1))
        _ = v22
    }
    var v23 int = v4
    if false {
        v24 := GenS3{F0: -9223372036854775808, F1: v2, F2: 517}
        fmt.Println(v4, v2, v2)
        if true {
            if false {
                v25 := GenS4{F0: 82, F1: 0}
                if (!false) {
                    var v26 float32 = 1.5
                    var v27 int8 = 100
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(v1, v4, v2)
                    fmt.Println(v4, v2, v4)
                    if false {
                        v4 += (-(v23))
                    }
                }
                var v28 float64 = (3.14159 + 1.5)
                _ = v25
                _ = v28
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 uint16 = 255
                var v31 uint64 = v2
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 float32 = 3.14159
            _ = v32
        } else {
            if (499 == (v1 & v1)) {
                var v33 int8 = -1
                _ = v33
            } else {
                switch v4 {
                case 149:
                    for fi34 := 0; fi34 < 17; fi34++ {
                        v5 = append(v5, (v1 & (^(v1))))
                        v2 = (v2 + (v2 % ((v2) | 1)))
                        if len(v5) > 0 {
                            v5[0] = v1
                        }
                        _ = fi34
                    }
                }
                var v35 string = ("foo bar" + "\n\t")
                fmt.Println(v1, v23)
                _ = v35
            }
        }
        _ = v24
    }
    if ("foo bar" == "\n\t") {
        v36 := GenS5{F0: 42, F1: 1000}
        if (!true) {
            v37 := GenS6{F0: genFn0(844)}
            fmt.Println(v23, v1)
            _ = v37
        }
        var v38 int16 = 42
        _ = v36
        _ = v38
    }
    v2 += (v2 ^ v2)
    v39 := GenS3{F0: -977, F1: (v2 / (((v2 % ((375) | 1))) | 1)), F2: (^(v2))}
    for fi40 := 0; fi40 < 15; fi40++ {
        if (((!false) || false) || true) {
            v41 := GenS2{F0: (v2 - 42), F1: -1}
            v4 += v23
            _ = v41
        } else {
            v42 := GenS1{F0: v23, F1: genFn1()}
            _ = v42
        }
        _ = fi40
    }
    v23 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
