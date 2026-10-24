package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 bool; F2 bool }
type GenS4 struct { F0 bool; F1 int16; F2 int32 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 string, p2 int32) uint64 {
    var v5 uint64 = 18446744073709551615
    var v6 int = p0
    v7 := []int{(^(v6)), -19, (p0 * (-(v6))), (p0 ^ -2147483648)}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint64 = v5
        if (float32(1e10) >= 0.1) {
            if true {
                v6 = fi8
                p1 = ("a" + ((("foo bar" + p1) + "\n\t") + p1))
                for fi10 := 0; fi10 < 8; fi10++ {
                    p0 += p0
                    _ = fi10
                }
            }
        } else {
            v11 := GenS1{F0: (p0 * 352), F1: int64(fi8)}
            p2 -= (p2 % ((p2) | 1))
            _ = v11
        }
        fmt.Println(p1, v9, fi8)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]uint32)
    v3 := []int{728, -988}
    v4 := make(map[int]string)
    if (true || true) {
        v12 := GenS2{F0: 29}
        _ = v12
    }
    v1 = (v1 + "foo bar")
    var v13 uint = 42
    var v14 uint16 = 255
    v1 = v1
    v14 = v14
    if true {
        switch 1 {
        case 32767:
            _ = len(v3)
        case -1:
            switch 32767 {
            case -916:
                switch 9223372036854775807 {
                case -447:
                    var v15 uint16 = (v14 * (v14 ^ 255))
                    _ = v15
                default:
                    v14 -= (v14 & 42)
                }
            case 9223372036854775807:
                if true {
                    v16 := GenS1{F0: 42, F1: -9223372036854775808}
                    _ = v16
                }
            }
        default:
            var v17 uint = uint(v14)
            _ = v17
        }
        delete(v2, 42)
        fmt.Println(v14, v14, v1)
    } else {
        if true {
            var v18 uint = 18446744073709551615
            var v19 float64 = (0.1 / 1e10)
            _ = len(v2)
            _ = v18
            _ = v19
        }
        if (1e10 <= ((1e10 / (((((1e10 + (1e10 * (-2.25 + 3.14159))) - (float64(1e10) / 1.0)) / 1.0) / -2.25) / 3.14159)) / 3.14159)) {
            v20 := GenS1{F0: -805, F1: 9223372036854775807}
            var v21 uint16 = v14
            switch -1 {
            case -651:
                v22 := GenS1{F0: -358, F1: -1}
                _ = v22
            case -9223372036854775808:
                v23 := GenS3{F0: 728, F1: (v14 > (v21 << ((v13) & 63))), F2: false}
                _ = v23
            }
            _ = v20
            _ = v21
        } else {
            fmt.Println(v1)
            for fi24 := 0; fi24 < 4; fi24++ {
                v13 += 394
                v25 := GenS4{F0: (false && (!false)), F1: 1, F2: 1}
                for fi26 := 0; fi26 < 2; fi26++ {
                    if (483 < 1) {
                        v27 := GenS4{F0: false, F1: 42, F2: -2147483648}
                        _ = v27
                    } else {
                        v1 = v1
                        fmt.Println(v1, fi24)
                        if false {
                            switch (-(fi24)) {
                            case 997:
                                fmt.Println(fi24)
                            case -1:
                                if (!((v13 % (((^(v13))) | 1)) < (^(v13)))) {
                                    v1 = ((v1 + "foo bar") + (((v1 + "\n\t") + v1) + v1))
                                    v28 := GenS2{F0: -127}
                                    _ = v28
                                }
                            }
                            v1 = v1
                            v1 = "\n\t"
                        } else {
                            var v29 uint8 = 255
                            _ = v29
                        }
                    }
                    _ = fi26
                }
                _ = fi24
                _ = v25
            }
            fmt.Println(v14)
        }
    }
    if false {
        v14 -= 65535
    } else {
        switch 9223372036854775807 {
        case -1:
            v30 := v2[65535]
            _ = v30
        }
        if false {
            v31 := GenS4{F0: (71 >= -90), F1: 0, F2: 42}
            _ = v31
        } else {
            if false {
                v32 := GenS5{F0: v14, F1: genFn0(42, v1, -752), F2: 26}
                _ = v32
            } else {
                v1 = "a"
                if (int16(v13) >= 42) {
                    v1 = v1
                }
            }
        }
    }
    v1 = "foo bar"
    var v33 uint8 = 249
    var v34 int16 = 0
    var v35 int = 127
    switch ("" + ((v1 + "foo bar") + v1)) {
    case "\n\t":
        v4[v35] = (v1 + ((v1 + "foo bar") + "foo bar"))
    case "hello":
        v36 := GenS3{F0: genFn0((v35 >> ((v13) & 63)), (((v1 + "hello") + v1) + v1), -311), F1: (!true), F2: (!true)}
        _ = v36
    case "a":
        for fi37 := 0; fi37 < 19; fi37++ {
            switch ("foo bar" + v1) {
            case "foo bar":
                if ((v14 == v14) && true) {
                    v14 += v14
                    switch (v35 << ((v13) & 63)) {
                    case -9223372036854775808:
                        v38 := GenS5{F0: (v14 << ((v13) & 63)), F1: 634, F2: (v33 + v33)}
                        _ = v38
                    default:
                        v39 := GenS6{F0: 42, F1: uint16(v13)}
                        _ = v39
                    }
                    v3 = append(v3, 4294967295)
                } else {
                    v35 = (fi37 * (v35 ^ 279))
                    v1 = "\n\t"
                }
            default:
                v40 := GenS1{F0: (^(fi37)), F1: 0}
                _ = v40
            }
            var v41 bool = false
            _ = fi37
            _ = v41
        }
    default:
        var v42 bool = ((^(v13)) >= 255)
        _ = v42
    }
    if len(v3) > 0 {
        v3[0] = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
