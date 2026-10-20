package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 bool; F2 int }
func genFn0() int32 {
    var v4 int64 = -850
    v4 -= v4
    v4 = v4
    var v5 int64 = v4
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1() int {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float32 = 0.0
        switch "hello" {
        case "\n\t":
            if true {
                v7 += float32(1e10)
                v7 = 1.5
            }
        }
        var v8 string = ("" + "\n\t")
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 bool = (0 != 0)
        var v11 int16 = 1000
        _ = fi9
        _ = v10
        _ = v11
    }
    if false {
        var v12 float64 = 1.5
        _ = v12
    } else {
        var v13 int16 = -32768
        v14 := GenS1{F0: (v13 - (v13 * 65)), F1: 889}
        _ = v13
        _ = v14
    }
    return 32767
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]int)
    v3 := []int{9223372036854775807, 32768, -642}
    v1 = 18446744073709551615
    for fi15 := 0; fi15 < 15; fi15++ {
        if false {
            switch (^(fi15)) {
            case 101:
                if len(v3) > 0 {
                    v3[0] = (-(fi15))
                }
            default:
                if len(v3) > 0 {
                    v3[0] = (fi15 % (((fi15 & -2147483648)) | 1))
                }
            }
            for fi16 := 0; fi16 < 14; fi16++ {
                v1 += (v1 * 700)
                v1 = 42
                _ = fi16
            }
            v1 = v1
        } else {
            fmt.Println(v1, v1)
        }
        v1 = 0
        for fi17 := 0; fi17 < 3; fi17++ {
            _ = len(v3)
            _ = fi17
        }
        _ = fi15
    }
    _ = len(v3)
    for fi18 := 0; fi18 < 9; fi18++ {
        switch "\n\t" {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi18
    }
    var v19 int8 = -16
    var v20 int32 = 0
    var v21 int32 = -1
    var v22 float32 = -2.25
    v23 := GenS1{F0: 0, F1: 4294967295}
    if ((42 > 1) && (((false || (-1 != -32768)) && false) && ((-279 == -580) && (v22 == (float32(1e10) - (v22 + v22)))))) {
        fmt.Println(v1)
        if false {
            if (float64(0.0) != (0.1 / (0.1 + 0.0))) {
                v22 += v22
            } else {
                v24 := GenS1{F0: -32768, F1: 0}
                _ = len(v3)
                _ = v24
            }
            v22 = v22
            fmt.Println(v19, v22, v21)
        } else {
            v2[858] = -9223372036854775808
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 string = (("foo bar" + ("foo bar" + "hello")) + "")
                var v27 int64 = -9223372036854775808
                _ = fi25
                _ = v26
                _ = v27
            }
            if ((4294967295 != 2147483647) || (127 <= v19)) {
                v28 := GenS1{F0: 32767, F1: 277}
                v19 = (v19 - 34)
                _ = v28
            }
        }
        fmt.Println(v19, v19)
    } else {
        fmt.Println(v22, v20, v21)
    }
    v20 += -1
    for fi29 := 0; fi29 < 4; fi29++ {
        v30 := GenS1{F0: -813, F1: 103}
        _ = fi29
        _ = v30
    }
    switch 32767 {
    case -2147483648:
        switch (("a" + ("" + ("hello" + ("\n\t" + ("hello" + "hello"))))) + "hello") {
        case "\n\t":
            _ = len(v2)
        case "a":
            switch ("foo bar" + "hello") {
            case "\n\t":
                v1 -= v1
            case "":
                if false {
                    v31 := GenS1{F0: -1, F1: 988}
                    _ = v31
                } else {
                    v22 = 0.1
                    v21 = (v20 | genFn0())
                }
            default:
                var v32 float32 = (v22 / v22)
                _ = v32
            }
        case "":
            fmt.Println(v22, v22, v20)
        default:
            v19 -= v19
        }
    case -491:
        v22 = float32(0.1)
    default:
        for fi33 := 0; fi33 < 10; fi33++ {
            if (-32768 < -1) {
                fmt.Println(v20)
                v34 := v2[32767]
                _ = v34
            }
            switch (fi33 ^ 42) {
            case 315:
                for fi35 := 0; fi35 < 5; fi35++ {
                    fmt.Println(fi35, fi35, v22)
                    _ = fi35
                }
            case 314:
                var v36 float32 = (0.0 - v22)
                _ = v36
            case 735:
                for fi37 := 0; fi37 < 17; fi37++ {
                    _ = len(v3)
                    if false {
                        v38 := GenS1{F0: 1, F1: 1}
                        var v39 int64 = 334
                        _ = v38
                        _ = v39
                    } else {
                        var v40 int16 = 532
                        _ = v40
                    }
                    fmt.Println(v1, v19, fi33)
                    _ = fi37
                }
            }
            fmt.Println(v22)
            _ = fi33
        }
    }
    fmt.Println(v22, v22)
    if false {
        v41 := GenS1{F0: 1, F1: 507}
        if len(v3) > 0 {
            v3[0] = -992
        }
        _ = v41
    }
    v42 := GenS2{F0: 255, F1: false, F2: 0}
    v20 += v21
    for fi43 := 0; fi43 < 4; fi43++ {
        if false {
            v22 += v22
        } else {
            v22 = (float32(0.0) / -2.25)
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v42
}
