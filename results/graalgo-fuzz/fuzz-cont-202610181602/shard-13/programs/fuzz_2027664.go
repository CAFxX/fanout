package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint64, p1 int) uint {
    p1 = (p1 / ((9223372036854775807) | 1))
    var v6 int64 = -9223372036854775808
    v7 := []int{(p1 * (p1 & -821)), (^(p1))}
    p1 -= (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int32) int64 {
    fmt.Println(p0)
    v8 := []int{127}
    _ = p0
    _ = v8
    return 83
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{1, 4294967295, 42}
    var v3 int = 42
    v4 := []int{v3, 469, (-(v3)), (v3 << ((18446744073709551615) & 63))}
    var v5 uint32 = 614
    switch "\n\t" {
    case "hello":
        v9 := GenS1{F0: -1}
        _ = v9
    default:
        var v10 uint16 = 42
        _ = v10
    }
    if (-1 < -128) {
        if false {
            delete(v1, v5)
        }
        v11 := GenS1{F0: -1}
        _ = v11
    } else {
        var v12 float32 = 1e10
        var v13 float64 = (float64(0.0) * -2.25)
        _ = v12
        _ = v13
    }
    v3 += -2147483648
    v3 = (v3 << ((0) & 63))
    v5 = 492
    if (-712 != 42) {
        fmt.Println(v3, v3)
    }
    v3 -= (v3 & 32768)
    if (false && ((!((255 >= uint16(v5)) || true)) && false)) {
        var v14 uint16 = 353
        _ = v14
    }
    v5 += v5
    switch (v3 << ((65535) & 63)) {
    case 199:
        for fi15 := 0; fi15 < 9; fi15++ {
            switch "foo bar" {
            case "\n\t":
                v2 = append(v2, (v5 ^ (v5 | 900)))
            case "a":
                for fi16 := 0; fi16 < 2; fi16++ {
                    v4 = append(v4, 128)
                    _ = fi16
                }
            case "foo bar":
                v3 = (v3 - fi15)
            default:
                switch 127 {
                case 4294967295:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        var v18 int = -2147483648
                        _ = fi17
                        _ = v18
                    }
                case 172:
                    v3 += 9223372036854775807
                default:
                    var v19 uint32 = v5
                    _ = v19
                }
            }
            _ = fi15
        }
    case -1:
        var v20 int32 = 2147483647
        _ = v20
    default:
        if true {
            if true {
                v5 += v5
            } else {
                var v21 uint32 = v5
                _ = v21
            }
        } else {
            v3 += v3
            var v22 bool = (0.1 >= 3.14159)
            _ = v22
        }
    }
    delete(v1, 791)
    var v23 float64 = 1e10
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 int8 = -52
        for fi26 := 0; fi26 < 4; fi26++ {
            switch (fi24 ^ fi26) {
            case 585:
                switch fi26 {
                case -929:
                    v5 -= 0
                case 4294967295:
                    v27 := v1[v5]
                    _ = v27
                default:
                    switch fi26 {
                    case 9223372036854775807:
                        var v28 uint = 255
                        if (false && false) {
                            if ((v28 & 330) == 42) {
                                v28 = 2
                            } else {
                                var v29 uint32 = 960
                                _ = v29
                            }
                            v3 = v3
                        } else {
                            v25 = v25
                            var v30 int64 = 42
                            var v31 int32 = 1
                            _ = v30
                            _ = v31
                        }
                        _ = v28
                    case -307:
                        var v32 int32 = -2147483648
                        _ = v32
                    case 168:
                        v3 += (^(v3))
                    }
                }
            case 4294967295:
                var v33 uint = 715
                _ = v33
            case -509:
                v5 -= (^(v5))
            }
            _ = fi26
        }
        v3 += (fi24 >> ((255) & 63))
        _ = fi24
        _ = v25
    }
    v34 := GenS1{F0: -16}
    v35 := GenS1{F0: 478}
    var v36 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
