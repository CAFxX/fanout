package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 uint32; F2 uint64 }
func genFn0(p0 string, p1 uint) int8 {
    p0 = p0
    p0 = p0
    switch 1 {
    case 1:
        switch p0 {
        case "":
            var v5 uint = p1
            _ = v5
        case "a":
            var v6 int64 = -9223372036854775808
            _ = v6
        case "\n\t":
            if ((!((("a" + ("" + (((p0 + p0) + "hello") + ((p0 + "") + p0)))) + (p0 + p0)) <= p0)) && true) {
                for fi7 := 0; fi7 < 7; fi7++ {
                    fmt.Println(p1, p1, fi7)
                    _ = fi7
                }
            }
        default:
            p0 = "\n\t"
        }
    case 32767:
        var v8 bool = true
        _ = v8
    case 614:
        v9 := []int{-50, 9223372036854775807}
        _ = v9
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]uint32)
    var v3 uint8 = 214
    v4 := make(map[uint32]int)
    var v10 int16 = -715
    v11 := []int{1, 32768}
    for fi12 := 0; fi12 < 9; fi12++ {
        fmt.Println(v10, v1, v3)
        _ = fi12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v3 -= (v3 / ((uint8(v1)) | 1))
        var v14 float32 = 0.1
        switch "" {
        case "":
            switch (fi13 - (-(fi13))) {
            case 32768:
                for fi15 := 0; fi15 < 19; fi15++ {
                    v14 = 3.14159
                    _ = fi15
                }
            case -674:
                for fi16 := 0; fi16 < 17; fi16++ {
                    v1 = (v1 | v1)
                    if true {
                        v17 := GenS1{F0: 100, F1: 1}
                        _ = v17
                    } else {
                        fmt.Println(v3, v1)
                        v18 := GenS1{F0: 42, F1: 160}
                        _ = v18
                    }
                    fmt.Println(v14, fi16)
                    v19 := GenS2{F0: v10}
                    _ = fi16
                    _ = v19
                }
            default:
                switch "foo bar" {
                case "foo bar":
                    v20 := GenS3{F0: (v3 & (^(v3))), F1: 147}
                    _ = v20
                case "\n\t":
                    var v21 string = "\n\t"
                    _ = v21
                case "a":
                    var v22 uint = v1
                    _ = v22
                default:
                    fmt.Println(v14, v1, fi13)
                }
            }
        case "\n\t":
            v23 := GenS3{F0: v3, F1: uint64(v3)}
            _ = v23
        default:
            if false {
                if true {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v3 -= (v3 / (((^(v3))) | 1))
                        fmt.Println(v3, fi24, v1)
                        switch (fi13 ^ 2147483647) {
                        case -803:
                            v25 := GenS3{F0: 2, F1: 523}
                            _ = v25
                        }
                        _ = fi24
                    }
                    var v26 int8 = genFn0("", v1)
                    _ = v26
                } else {
                    fmt.Println(v1, v10, v14)
                    v1 += v1
                }
                if len(v11) > 0 {
                    v11[0] = fi13
                }
            } else {
                delete(v4, 42)
                var v27 int32 = -1
                _ = v27
            }
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v10)
    v10 = v10
    var v28 int64 = 1
    for fi29 := 0; fi29 < 12; fi29++ {
        for fi30 := 0; fi30 < 15; fi30++ {
            v3 += v3
            v10 = (v10 + (v10 % ((1) | 1)))
            var v31 uint16 = 42
            if (false && true) {
                v1 = v1
                v11 = append(v11, fi29)
                fmt.Println(fi30)
                var v32 uint16 = v31
                v32 = (v32 << (((^(v1))) & 63))
                fmt.Println(v31, v3, v32)
                if len(v11) > 0 {
                    v11[0] = int(v10)
                }
                v33 := GenS4{F0: (v1 ^ v1)}
                _ = v32
                _ = v33
            } else {
                var v34 int16 = (^(v10))
                v10 = v34
                var v35 uint16 = 42
                _ = v34
                _ = v35
            }
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi30
            _ = v31
        }
        v10 = int16(v28)
        fmt.Println(fi29)
        _ = fi29
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        var v37 float64 = 0.0
        if (-128 == 100) {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
        for fi38 := 0; fi38 < 12; fi38++ {
            if (fi38 <= int(v10)) {
                switch "\n\t" {
                case "hello":
                    delete(v2, "a")
                }
            } else {
                v3 = (v3 - v3)
                v10 -= v10
            }
            _ = fi38
        }
        _ = fi36
        _ = v37
    }
    if (4294967295 >= 0) {
        v39 := GenS3{F0: (v3 >> (((v1 & 1)) & 63)), F1: 4294967295}
        _ = v39
    } else {
        v40 := GenS2{F0: v10}
        for fi41 := 0; fi41 < 17; fi41++ {
            v10 += -405
            var v42 uint64 = 585
            _ = fi41
            _ = v42
        }
        _ = v40
    }
    v3 = 127
    if false {
        fmt.Println(v10)
        v2["\n\t"] = 65535
        v43 := GenS2{F0: (v10 << (((v1 % ((65535) | 1))) & 63))}
        _ = v43
    } else {
        var v44 float64 = 1.5
        v10 -= -1
        v44 = (1.5 * (3.14159 * 1.5))
        _ = v44
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        if (!(1 <= 65535)) {
            for fi46 := 0; fi46 < 9; fi46++ {
                _ = len(v2)
                _ = fi46
            }
            if (!false) {
                v10 = v10
            }
        } else {
            if false {
                var v47 int32 = -67
                v28 -= (-(v28))
                if true {
                    var v48 uint32 = 65535
                    if (1 < genFn0("foo bar", (v1 & (v1 % (((v1 / (((v1 + v1)) | 1))) | 1))))) {
                        var v49 uint16 = 707
                        _ = v49
                    } else {
                        for fi50 := 0; fi50 < 10; fi50++ {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            v51 := GenS3{F0: 1, F1: uint64(v48)}
                            if false {
                                v1 = uint(v3)
                                var v52 int32 = -302
                                _ = v52
                            }
                            _ = fi50
                            _ = v51
                        }
                        var v53 float64 = 1e10
                        _ = v53
                    }
                    fmt.Println(v1, fi45)
                    _ = v48
                }
                _ = v47
            }
            v10 -= (^(v10))
        }
        for fi54 := 0; fi54 < 11; fi54++ {
            var v55 uint16 = 255
            _ = fi54
            _ = v55
        }
        v3 = (v3 * 0)
        _ = fi45
    }
    v28 = v28
    if true {
        fmt.Println(v3)
        v3 = (v3 & 1)
    }
    v56 := GenS5{F0: 1, F1: 4294967295, F2: 421}
    v3 += v3
    for fi57 := 0; fi57 < 1; fi57++ {
        var v58 float32 = (3.14159 + (1e10 - 1e10))
        for fi59 := 0; fi59 < 13; fi59++ {
            var v60 uint16 = uint16(v3)
            _ = fi59
            _ = v60
        }
        _ = fi57
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v56
}
