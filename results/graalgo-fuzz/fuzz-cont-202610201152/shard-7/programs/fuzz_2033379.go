package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint, p1 int32) int64 {
    p0 = p0
    fmt.Println(p0, p0)
    if ((426 > 32767) || true) {
        v5 := GenS1{F0: 1, F1: -32768, F2: 2147483648}
        if true {
            if ((3.14159 * 1.5) >= 1e10) {
                p0 = p0
            } else {
                v6 := []int{0, -453, -387, 32767}
                v7 := GenS2{F0: -32768}
                p1 += (p1 * p1)
                _ = v6
                _ = v7
            }
            fmt.Println(p1, p1)
            v8 := GenS1{F0: 229, F1: 32767, F2: 1}
            _ = v8
        } else {
            if true {
                p1 = -2147483648
            }
        }
        var v9 int16 = -237
        _ = v5
        _ = v9
    } else {
        switch 127 {
        case 2147483647:
            var v10 float64 = -2.25
            _ = v10
        default:
            p1 = p1
        }
        var v11 uint16 = 665
        _ = v11
    }
    v12 := GenS3{F0: "\n\t", F1: (^(p0))}
    _ = p0
    _ = p1
    _ = v12
    return -468
}
func genFn1(p0 uint32) int16 {
    p0 = (^(p0))
    v13 := GenS1{F0: 206, F1: -32768, F2: p0}
    var v14 bool = (0 == 2)
    _ = p0
    _ = v13
    _ = v14
    return -982
}

func main() {
    var v1 uint8 = 1
    var v2 uint = 42
    var v3 int = 42
    var v4 uint = (v2 - 255)
    fmt.Println(v2, v2, v2)
    if false {
        v15 := []int{}
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int32 = 714
            v17 -= -781
            v1 += 205
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        if false {
            v18 := []int{v3, (-(v3))}
            var v19 int = -309
            _ = v18
            _ = v19
        }
    }
    fmt.Println(v2, v1, v3)
    var v20 int64 = 0
    v21 := []int{v3, (^(v3))}
    v22 := GenS1{F0: 0, F1: 503, F2: 65535}
    switch ("hello" + (("" + "foo bar") + "a")) {
    case "hello":
        v20 -= (v20 - (v20 << (((^(v2))) & 63)))
    case "a":
        var v23 int = (v3 ^ (^(v3)))
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        if (!true) {
            v1 = v1
            if true {
                v4 -= (v2 % (((v4 | v4)) | 1))
            }
            v2 -= v2
        } else {
            if true {
                if true {
                    v2 = 255
                    var v25 int = (-(fi24))
                    var v26 float32 = ((0.0 + (-2.25 + -2.25)) / float32((1.5 / (((3.14159 / float64((((1e10 * 3.14159) * (3.14159 + 0.0)) * 3.14159))) / 1.0) / ((float64(float32((1e10 / 1e10))) - 3.14159) - 0.1)))))
                    _ = v25
                    _ = v26
                }
                for fi27 := 0; fi27 < 11; fi27++ {
                    v4 += (v4 + (v4 >> ((v4) & 63)))
                    if (511 == 2147483647) {
                        fmt.Println(v20, v2, v3)
                        if false {
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                        } else {
                            v28 := GenS4{F0: 349, F1: genFn0((v4 >> ((v2) & 63)), 0)}
                            v1 = 0
                            switch "\n\t" {
                            case "a":
                                var v29 int = (v3 % ((fi24) | 1))
                                _ = v29
                            default:
                                v20 = -9223372036854775808
                            }
                            _ = v28
                        }
                    }
                    _ = fi27
                }
            }
            var v30 int16 = genFn1(2147483648)
            v31 := GenS2{F0: v30}
            _ = v30
            _ = v31
        }
        v1 += v1
        _ = fi24
    }
    v1 = (v1 << ((v4) & 63))
    v32 := GenS5{F0: -1}
    switch ("\n\t" + (("hello" + "") + "a")) {
    case "a":
        v4 -= 577
    default:
        v21 = append(v21, 42)
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    var v33 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
