package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint }
func genFn0(p0 int8, p1 int16, p2 int64) uint8 {
    fmt.Println(p0)
    v5 := []int{127, 478}
    p1 -= (p1 / ((-453) | 1))
    var v6 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -2147483648
    var v3 uint64 = 493
    var v4 int8 = 42
    for fi7 := 0; fi7 < 7; fi7++ {
        v2 += v2
        _ = fi7
    }
    v8 := GenS1{F0: 1, F1: (v3 | 0)}
    for fi9 := 0; fi9 < 5; fi9++ {
        v4 = (v4 & v4)
        _ = fi9
    }
    v2 -= (v2 / ((v2) | 1))
    if (false || (255 > 0)) {
        v3 = 1
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: 32767, F1: (v3 ^ (v3 * 42))}
            v4 = (v4 / ((v4) | 1))
            var v12 uint32 = 477
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int64 = 854
        _ = v13
    } else {
        switch ("a" + ("foo bar" + (("hello" + "foo bar") + ("\n\t" + "foo bar")))) {
        case "hello":
            if (!true) {
                v14 := GenS2{F0: 299}
                switch -2147483648 {
                case 4294967295:
                    var v15 uint32 = 0
                    _ = v15
                case 32768:
                    fmt.Println(v2, v2)
                default:
                    for fi16 := 0; fi16 < 3; fi16++ {
                        v4 += 0
                        for fi17 := 0; fi17 < 10; fi17++ {
                            if true {
                                v18 := []int{(fi16 >> ((1) & 63)), 127}
                                var v19 uint32 = 2147483647
                                v19 += 0
                                _ = v18
                                _ = v19
                            }
                            _ = fi17
                        }
                        _ = fi16
                    }
                }
                _ = v14
            } else {
                switch 1 {
                case -275:
                    v3 = (v3 + (v3 % (((v3 - 220)) | 1)))
                case 128:
                    v4 = (v4 / (((v4 & v4)) | 1))
                case 4294967295:
                    if ((true || (!(-2147483648 < -2147483648))) && (genFn0(v4, -1, 0) < 20)) {
                        for fi20 := 0; fi20 < 19; fi20++ {
                            var v21 uint32 = 2
                            _ = fi20
                            _ = v21
                        }
                    }
                }
                var v22 uint = 563
                _ = v22
            }
        case "foo bar":
            var v23 uint64 = v3
            _ = v23
        default:
            v2 -= (-(v2))
        }
        var v24 bool = false
        fmt.Println(v3)
        _ = v24
    }
    v3 = v3
    v25 := GenS3{F0: 0}
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 float32 = -2.25
        v4 = -1
        v2 += v2
        _ = fi26
        _ = v27
    }
    var v28 bool = ((!("\n\t" == "a")) && true)
    if v28 {
        fmt.Println(v4, v4, v2)
        var v29 float32 = 0.0
        _ = v29
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        v31 := GenS4{F0: -26, F1: 296}
        _ = fi30
        _ = v31
    }
    var v32 uint32 = 507
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v8
}
