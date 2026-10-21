package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int; F2 int }
type GenS4 struct { F0 int32; F1 int16 }
func genFn0(p0 string, p1 uint32) uint {
    var v4 string = "foo bar"
    var v5 uint32 = (p1 + p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2
}

func main() {
    var v1 float32 = 1e10
    var v2 int8 = 0
    var v3 uint16 = 0
    var v6 int32 = -2147483648
    switch 32768 {
    case 187:
        v6 += (v6 & 96)
    case 2147483647:
        v7 := GenS1{F0: 42, F1: ("\n\t" + "hello")}
        _ = v7
    case -635:
        switch 128 {
        case 32768:
            v2 = -95
        case 9223372036854775807:
            for fi8 := 0; fi8 < 2; fi8++ {
                v1 = (-2.25 / v1)
                _ = fi8
            }
        case -1:
            v6 += (v6 + v6)
        }
    default:
        v6 += (v6 << ((0) & 63))
    }
    var v9 bool = true
    v1 -= float32((3.14159 / (-2.25 - float64((1e10 - (v1 / v1))))))
    if ("a" != "\n\t") {
        if ((^(v6)) >= v6) {
            var v10 uint8 = 42
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 float64 = ((1.5 + -2.25) - 3.14159)
                v1 -= 0.0
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            if v9 {
                var v13 uint64 = 558
                if (!v9) {
                    var v14 int64 = 670
                    if false {
                        var v15 int32 = (v6 | v6)
                        var v16 bool = v9
                        _ = v15
                        _ = v16
                    } else {
                        var v17 bool = ("\n\t" >= (("hello" + "a") + "hello"))
                        _ = v17
                    }
                    _ = v14
                }
                _ = v13
            } else {
                if v9 {
                    if (895 > 0) {
                        switch -725 {
                        case 0:
                            v6 = (v6 / ((v6) | 1))
                        case -74:
                            fmt.Println(v6, v1, v1)
                        default:
                            var v18 float32 = (float32(float64(((v1 - ((((3.14159 * 1e10) / (1e10 + (v1 / (v1 + v1)))) - (v1 + v1)) + v1)) - 0.0))) + v1)
                            _ = v18
                        }
                    }
                    v1 += ((v1 * float32(0.0)) * v1)
                    if ("a" == "hello") {
                        v19 := []int{9223372036854775807, 2147483647}
                        _ = v19
                    }
                } else {
                    var v20 int = -1
                    v21 := GenS2{F0: 670}
                    v9 = true
                    _ = v20
                    _ = v21
                }
                var v22 uint8 = 2
                _ = v22
            }
        }
    } else {
        v23 := []int{373, 915}
        if false {
            fmt.Println(v2, v9)
            v24 := GenS3{F0: 100, F1: 4294967295, F2: -529}
            if len(v23) > 0 {
                v23[0] = 127
            }
            _ = v24
        } else {
            v25 := GenS4{F0: (v6 | (v6 % (((^(v6))) | 1))), F1: 951}
            _ = v25
        }
        _ = v23
    }
    fmt.Println(v3)
    var v26 uint32 = 2
    if true {
        var v27 int = 414
        var v28 int8 = (-(v2))
        _ = v27
        _ = v28
    } else {
        v29 := GenS2{F0: 933}
        if true {
            fmt.Println(v2, v26)
        }
        v30 := GenS1{F0: (v3 + (v3 | (v3 % ((65535) | 1)))), F1: "\n\t"}
        _ = v29
        _ = v30
    }
    v1 -= 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v6
    _ = v9
}
