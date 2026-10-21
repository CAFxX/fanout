package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 string }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0() bool {
    v5 := []int{42}
    switch 0 {
    case -300:
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 float32 = 0.1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi6
            _ = v7
        }
    case 32768:
        var v8 uint32 = 2147483647
        _ = v8
    default:
        var v9 uint16 = 367
        _ = v9
    }
    var v10 int64 = 0
    if false {
        v10 = (v10 + -675)
        v11 := GenS1{F0: 521, F1: 0, F2: ("" + "hello")}
        _ = v11
    }
    var v12 int16 = 742
    _ = v10
    _ = v12
    _ = v5
    return true
}

func main() {
    var v1 uint8 = 1
    var v2 float32 = 1e10
    var v3 int64 = 1
    var v4 float64 = (-2.25 / float64(((float32(1e10) + v2) + v2)))
    v3 = (v3 - (^(v3)))
    v13 := []int{-9223372036854775808, -9223372036854775808}
    _ = len(v13)
    if (2 == 789) {
        var v14 float32 = 1e10
        switch -633 {
        case 32768:
            v2 -= v14
        case 4294967295:
            v15 := GenS2{F0: 4294967295, F1: 32767}
            _ = v15
        case 128:
            for fi16 := 0; fi16 < 13; fi16++ {
                for fi17 := 0; fi17 < 5; fi17++ {
                    v13 = append(v13, fi17)
                    _ = fi17
                }
                var v18 uint64 = 1
                if len(v13) > 0 {
                    v13[0] = 32767
                }
                _ = fi16
                _ = v18
            }
        default:
            var v19 uint8 = 27
            _ = v19
        }
        _ = v14
    }
    fmt.Println(v1)
    var v20 uint = 255
    fmt.Println(v20, v4)
    _ = len(v13)
    for fi21 := 0; fi21 < 17; fi21++ {
        if true {
            var v22 float64 = 0.0
            _ = v22
        }
        v1 = (v1 | uint8(v20))
        _ = fi21
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(v20, v20)
        for fi24 := 0; fi24 < 6; fi24++ {
            if ((^(v3)) <= v3) {
                fmt.Println(v3, fi24, v4)
                v2 = (float32(v4) * (-2.25 - v2))
                fmt.Println(fi23, v1)
            }
            _ = fi24
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 int8 = 1
            _ = fi25
            _ = v26
        }
        _ = fi23
    }
    var v27 int64 = -887
    switch 366 {
    case 959:
        var v28 int = -2147483648
        _ = v28
    case -9223372036854775808:
        var v29 uint64 = uint64(v1)
        _ = v29
    default:
        if (v1 == v1) {
            var v30 int = -290
            _ = v30
        } else {
            if false {
                for fi31 := 0; fi31 < 13; fi31++ {
                    v2 -= 1e10
                    if (-128 != 100) {
                        if true {
                            var v32 int32 = 0
                            var v33 uint = (^(v20))
                            _ = v32
                            _ = v33
                        }
                        if true {
                            _ = len(v13)
                            fmt.Println(v3, v20)
                        }
                    }
                    v13 = append(v13, fi31)
                    _ = fi31
                }
            }
        }
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v1 += v1
        for fi35 := 0; fi35 < 19; fi35++ {
            v1 = 198
            _ = fi35
        }
        v36 := GenS2{F0: 4294967295, F1: -32768}
        _ = fi34
        _ = v36
    }
    if true {
        var v37 int8 = 0
        _ = v37
    } else {
        v38 := GenS1{F0: v3, F1: 1, F2: "\n\t"}
        v39 := GenS1{F0: (v3 / (((v27 & v27)) | 1)), F1: 42, F2: ((("foo bar" + ("hello" + "")) + "") + "a")}
        for fi40 := 0; fi40 < 13; fi40++ {
            switch (fi40 % ((32767) | 1)) {
            case -1:
                var v41 uint64 = 42
                _ = v41
            case -239:
                var v42 int64 = 609
                _ = v42
            case -9223372036854775808:
                var v43 uint = 42
                _ = v43
            }
            _ = fi40
        }
        _ = v38
        _ = v39
    }
    switch "hello" {
    case "":
        for fi44 := 0; fi44 < 8; fi44++ {
            v1 += 189
            _ = fi44
        }
    default:
        var v45 string = "\n\t"
        _ = v45
    }
    v46 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615, F2: "\n\t"}
    v47 := GenS1{F0: (v3 << ((v20) & 63)), F1: 1, F2: ("" + "a")}
    v27 = (v3 / ((v27) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v46
    _ = v47
}
