package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int32; F2 string }
type GenS4 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 bool }

func main() {
    v1 := []uint32{188, 1}
    v2 := make(map[int]int)
    var v3 int32 = 0
    v4 := GenS1{F0: 2147483647, F1: 15, F2: 293}
    if (!false) {
        fmt.Println(v3)
        switch "\n\t" {
        case "\n\t":
            v5 := GenS1{F0: v3, F1: 100, F2: 255}
            _ = v5
        default:
            v3 -= v3
        }
        v6 := v2[-1]
        _ = v6
    } else {
        v7 := v2[0]
        _ = v7
    }
    switch "" {
    case "":
        fmt.Println(v3)
    case "a":
        if true {
            var v8 uint16 = 65535
            _ = v8
        } else {
            v3 += 0
            var v9 int32 = (v3 / ((194) | 1))
            fmt.Println(v9, v3)
            _ = v9
        }
    default:
        if ((!true) || (!false)) {
            if true {
                v10 := GenS1{F0: v3, F1: 1, F2: 65535}
                v3 = (v3 % (((v3 | v3)) | 1))
                if (0.0 <= (((-2.25 * 3.14159) + 1.5) - 0.1)) {
                    v3 = (-(v3))
                } else {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        var v12 float64 = ((0.0 * float64(0.0)) + ((3.14159 - float64((3.14159 * 0.1))) - 0.0))
                        _ = fi11
                        _ = v12
                    }
                    v3 -= (v3 % (((v3 << ((4294967295) & 63))) | 1))
                    var v13 uint8 = 2
                    _ = v13
                }
                _ = v10
            } else {
                for fi14 := 0; fi14 < 8; fi14++ {
                    v15 := GenS2{F0: 1}
                    v3 = v3
                    v3 += (v3 & v3)
                    _ = fi14
                    _ = v15
                }
                var v16 float64 = 1e10
                _ = v16
            }
            if true {
                switch 0 {
                case 204:
                    var v17 int = -172
                    _ = v17
                case 127:
                    v3 = v3
                }
                switch -1 {
                case 128:
                    for fi18 := 0; fi18 < 8; fi18++ {
                        var v19 uint = 0
                        v19 = (v19 | (v19 % ((v19) | 1)))
                        _ = fi18
                        _ = v19
                    }
                case 32767:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v20 int32 = (v3 | 887)
                    _ = v20
                }
            } else {
                fmt.Println(v3)
            }
        } else {
            var v21 int16 = 1
            var v22 int16 = 1
            v23 := GenS1{F0: v3, F1: 1, F2: 2}
            _ = v21
            _ = v22
            _ = v23
        }
    }
    fmt.Println(v3)
    for fi24 := 0; fi24 < 3; fi24++ {
        v3 = v3
        var v25 int32 = v3
        var v26 int8 = 68
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint = 255
    v28 := GenS2{F0: 1}
    for fi29 := 0; fi29 < 8; fi29++ {
        if (45 < 1) {
            for fi30 := 0; fi30 < 10; fi30++ {
                delete(v2, fi30)
                var v31 uint = 2
                _ = fi30
                _ = v31
            }
        } else {
            var v32 int32 = int32(fi29)
            for fi33 := 0; fi33 < 19; fi33++ {
                var v34 uint32 = 153
                var v35 bool = false
                v36 := GenS3{F0: v35, F1: v3, F2: ("foo bar" + "a")}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = v32
        }
        if ((((1e10 + 0.0) * 0.1) + 3.14159) >= 0.0) {
            for fi37 := 0; fi37 < 20; fi37++ {
                v38 := GenS4{F0: 138, F1: 127, F2: -128}
                switch -54 {
                case 32767:
                    v2[fi37] = (fi37 & (fi29 >> ((v27) & 63)))
                }
                v39 := GenS3{F0: (780 <= 2147483648), F1: -2147483648, F2: ("" + ((("hello" + "hello") + ("foo bar" + (("foo bar" + ("\n\t" + "foo bar")) + ("\n\t" + (("hello" + ("\n\t" + ("" + ""))) + "foo bar"))))) + "foo bar"))}
                _ = fi37
                _ = v38
                _ = v39
            }
        } else {
            v40 := GenS5{F0: (v3 | (v3 + (v3 % (((v3 << ((v27) & 63))) | 1))))}
            fmt.Println(fi29)
            v3 -= v3
            _ = v40
        }
        fmt.Println(fi29)
        _ = fi29
    }
    v41 := GenS6{F0: 42, F1: true}
    v3 -= (v3 % ((-1) | 1))
    var v42 int = -1
    v43 := GenS1{F0: (v3 ^ (v3 % ((2147483647) | 1))), F1: 42, F2: 0}
    switch "foo bar" {
    case "":
        fmt.Println(v3)
    case "hello":
        if (18446744073709551615 > 498) {
            fmt.Println(v42)
            var v44 int32 = (v3 - v3)
            for fi45 := 0; fi45 < 10; fi45++ {
                var v46 int = v42
                v42 += v46
                v1 = append(v1, 42)
                _ = fi45
                _ = v46
            }
            _ = v44
        } else {
            if (3.14159 < 0.0) {
                fmt.Println(v27, v27)
                v1 = append(v1, 866)
            } else {
                v27 += v27
            }
        }
    default:
        v47 := GenS5{F0: v3}
        _ = v47
    }
    if false {
        var v48 int8 = 56
        var v49 uint = (^(v27))
        for fi50 := 0; fi50 < 12; fi50++ {
            var v51 int64 = 42
            _ = fi50
            _ = v51
        }
        _ = v48
        _ = v49
    } else {
        var v52 uint8 = 255
        _ = v52
    }
    var v53 float32 = float32((-2.25 - ((0.1 - 0.0) * (-2.25 * 1.5))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v53
}
