package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint8, p1 int32) uint8 {
    if (!false) {
        var v4 float64 = ((1e10 / ((3.14159 + (3.14159 - (((((1.5 / 0.1) * float64(float32((1e10 * -2.25)))) * (0.1 / 3.14159)) / ((3.14159 * 1.5) - 1e10)) + (0.1 / (3.14159 / 0.1))))) / 1.0)) - 1e10)
        _ = v4
    }
    p1 = p1
    var v5 int8 = 110
    for fi6 := 0; fi6 < 18; fi6++ {
        if false {
            if ((fi6 & 453) > 9223372036854775807) {
                v7 := []int{(fi6 / (((-(fi6))) | 1)), (fi6 << ((1) & 63)), (fi6 & fi6), (fi6 >> ((1) & 63))}
                v5 += -15
                _ = v7
            }
            v8 := GenS1{F0: 1}
            _ = v8
        }
        p1 += p1
        var v9 bool = true
        _ = fi6
        _ = v9
    }
    p1 = (p1 * p1)
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1(p0 int8, p1 string) uint32 {
    var v10 uint8 = 255
    var v11 int = -348
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 992
}

func main() {
    v1 := []string{"a", "a", ("hello" + "hello")}
    var v2 int32 = 864
    v3 := []uint32{65535, 65535}
    v12 := GenS1{F0: 1}
    if false {
        if (v2 == v2) {
            v13 := GenS1{F0: 0}
            _ = v13
        }
    }
    if false {
        fmt.Println(v2)
        switch -1 {
        case 128:
            var v14 int32 = v2
            _ = v14
        case -105:
            fmt.Println(v2)
        default:
            switch "a" {
            case "a":
                for fi15 := 0; fi15 < 17; fi15++ {
                    var v16 uint32 = 0
                    _ = fi15
                    _ = v16
                }
            case "":
                v2 -= 42
            case "foo bar":
                var v17 int64 = -107
                _ = v17
            default:
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 int64 = -414
                    var v20 float32 = 1e10
                    v1 = append(v1, "\n\t")
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
        }
    }
    var v21 uint16 = 0
    if true {
        v2 = v2
        if (65535 != 4294967295) {
            if (42 >= 1) {
                v22 := GenS2{F0: 78, F1: 127, F2: (v21 >> ((2) & 63))}
                var v23 uint64 = 42
                var v24 int64 = -9223372036854775808
                _ = v22
                _ = v23
                _ = v24
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint16 = (v21 * v21)
            _ = fi25
            _ = v26
        }
    }
    v27 := GenS2{F0: 113, F1: -100, F2: (v21 * v21)}
    switch "" {
    case "hello":
        switch -9223372036854775808 {
        case -2147483648:
            v28 := GenS3{F0: false}
            _ = v28
        case -111:
            if true {
                v21 -= (v21 % ((v21) | 1))
            } else {
                for fi29 := 0; fi29 < 7; fi29++ {
                    switch "foo bar" {
                    case "hello":
                        if (v2 <= -719) {
                            var v30 uint32 = 0
                            var v31 uint16 = (v21 / ((v21) | 1))
                            if len(v1) > 0 {
                                v1[0] = ((("a" + "") + ("" + "foo bar")) + (("hello" + "hello") + "hello"))
                            }
                            _ = v30
                            _ = v31
                        }
                    }
                    v2 = (v2 % ((1) | 1))
                    _ = fi29
                }
                var v32 int32 = (v2 & (-(v2)))
                _ = v32
            }
        default:
            v21 -= (v21 - v21)
        }
    case "":
        var v33 int64 = -9223372036854775808
        _ = v33
    case "a":
        v2 = 42
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v27
    _ = v3
}
