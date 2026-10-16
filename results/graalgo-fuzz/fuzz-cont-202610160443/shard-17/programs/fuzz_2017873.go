package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint64, p1 uint16, p2 uint16) uint {
    p0 = 4294967295
    var v7 uint = 345
    if false {
        var v8 uint64 = 876
        if false {
            switch 127 {
            case -1:
                v9 := []int{-436, 2147483647, 9223372036854775807, 42}
                _ = v9
            case -9223372036854775808:
                v7 -= v7
            default:
                v10 := []int{-9223372036854775808, int(p0)}
                _ = v10
            }
            v8 += 974
            var v11 int64 = 0
            _ = v11
        }
        v12 := []int{4294967295, -1, 127, 2147483647}
        _ = v12
        _ = v8
    } else {
        p2 = 255
        if true {
            fmt.Println(v7, p2, p2)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 813
}

func main() {
    var v1 int = 945
    var v2 uint = 42
    var v3 int8 = 0
    var v4 string = ("foo bar" + "")
    var v5 int = (v1 >> (((v2 % ((1) | 1))) & 63))
    v6 := make(map[string]uint32)
    var v13 uint16 = 532
    v4 = v4
    var v14 float32 = 1.5
    var v15 int8 = (v3 << (((v2 & 402)) & 63))
    fmt.Println(v5, v3, v3)
    v4 = (v4 + "a")
    if false {
        var v16 uint8 = 0
        v4 = "hello"
        _ = v16
    }
    var v17 bool = (!false)
    delete(v6, v4)
    switch "hello" {
    case "foo bar":
        v4 = v4
    case "":
        if ((v15 / ((v3) | 1)) >= -1) {
            var v18 string = (v4 + v4)
            var v19 uint8 = 40
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                v17 = false
                var v21 int = (v5 - (v5 + (^(v1))))
                _ = fi20
                _ = v21
            }
            switch (v1 - (v1 % (((v1 >> (((v2 ^ (v2 << (((v2 * 229)) & 63)))) & 63))) | 1))) {
            case 745:
                switch v5 {
                case 128:
                    v22 := []int{(v1 / (((v1 * (^(v5)))) | 1))}
                    _ = v22
                case 2147483647:
                    if (((((0.0 - (float32(0.1) * ((1e10 - v14) + 1e10))) + v14) > (v14 / 3.14159)) && ((v13 << (((v2 ^ (v2 * (v2 - 249)))) & 63)) < v13)) && (-544 <= -1)) {
                        v23 := []int{(v1 * 11), (-(v1)), v1}
                        _ = v23
                    }
                default:
                    if v17 {
                        v13 -= v13
                        for fi24 := 0; fi24 < 11; fi24++ {
                            if (42 <= 29) {
                                v25 := GenS1{F0: 5}
                                _ = v25
                            } else {
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    var v27 uint32 = 531
                                    var v28 uint = (^(v2))
                                    v13 -= 2
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                            }
                            _ = fi24
                        }
                    } else {
                        for fi29 := 0; fi29 < 12; fi29++ {
                            for fi30 := 0; fi30 < 1; fi30++ {
                                var v31 float32 = 0.1
                                fmt.Println(fi30, v3, v15)
                                v32 := []int{(v1 | v1), v5, (^(v5))}
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            for fi33 := 0; fi33 < 17; fi33++ {
                                v1 = (fi33 * (-(fi33)))
                                _ = fi33
                            }
                            fmt.Println(v4, v15)
                            _ = fi29
                        }
                        for fi34 := 0; fi34 < 6; fi34++ {
                            v35 := GenS2{F0: (-(v15))}
                            v36 := []int{(v5 % ((v1) | 1)), fi34}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                }
            default:
                v17 = (754 != -1)
            }
            _ = len(v6)
        }
    default:
        for fi37 := 0; fi37 < 9; fi37++ {
            if (!v17) {
                var v38 uint16 = 32767
                _ = v38
            }
            for fi39 := 0; fi39 < 1; fi39++ {
                fmt.Println(v15)
                _ = fi39
            }
            _ = fi37
        }
    }
    v40 := GenS2{F0: (^(v3))}
    v41 := GenS3{F0: uint8(v2), F1: genFn0(0, 32767, 0)}
    v42 := []int{v1, (v5 % ((v5) | 1)), (v1 & (^(v1))), (v1 ^ -807)}
    fmt.Println(v13, v15)
    switch (("foo bar" + "\n\t") + (v4 + "hello")) {
    case "\n\t":
        v42 = append(v42, (v1 ^ (-(v1))))
    default:
        v43 := GenS4{F0: v15, F1: 29}
        _ = v43
    }
    v2 = (v2 / (((v2 * v2)) | 1))
    v44 := GenS1{F0: 334}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v44
    _ = v5
    _ = v6
}
