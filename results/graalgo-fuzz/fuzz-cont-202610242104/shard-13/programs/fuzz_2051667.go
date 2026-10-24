package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
func genFn0(p0 string) int8 {
    v5 := []int{2147483647}
    p0 = (p0 + p0)
    for fi6 := 0; fi6 < 2; fi6++ {
        if len(v5) > 0 {
            v5[0] = (fi6 << ((18446744073709551615) & 63))
        }
        switch 0 {
        case 296:
            var v7 uint32 = 65535
            _ = v7
        }
        var v8 int64 = -579
        _ = fi6
        _ = v8
    }
    if (((2 != 0) && (!true)) || (-424 <= 32768)) {
        var v9 int = 9223372036854775807
        if (0.0 >= (-2.25 / (float64(-2.25) / -2.25))) {
            var v10 uint32 = 1
            var v11 int8 = 1
            _ = v10
            _ = v11
        }
        switch (p0 + (p0 + (p0 + (("foo bar" + "foo bar") + p0)))) {
        case "a":
            v12 := GenS1{F0: 42, F1: -1}
            _ = v12
        }
        _ = v9
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            for fi14 := 0; fi14 < 9; fi14++ {
                switch (fi13 & fi14) {
                case -495:
                    var v15 uint8 = 55
                    _ = v15
                case 32767:
                    var v16 uint = 429
                    _ = v16
                case 2147483647:
                    p0 = ("foo bar" + p0)
                    p0 = ("a" + p0)
                    v17 := GenS2{F0: 950, F1: (fi13 ^ (fi14 >> ((4294967295) & 63))), F2: 127}
                    _ = v17
                default:
                    var v18 int8 = 66
                    p0 = (p0 + p0)
                    fmt.Println(p0)
                    _ = v18
                }
                _ = fi14
            }
            var v19 bool = true
            _ = fi13
            _ = v19
        }
        fmt.Println(p0)
        var v20 uint = 18446744073709551615
        _ = v20
    }
    var v21 uint64 = 1
    _ = p0
    _ = v21
    _ = v5
    return 0
}
func genFn1(p0 int16, p1 int) uint8 {
    v22 := GenS2{F0: 42, F1: (p1 / ((p1) | 1)), F2: 0}
    switch (p1 - (^(p1))) {
    case 42:
        v23 := GenS1{F0: (p0 - p0), F1: 1}
        _ = v23
    case -9223372036854775808:
        fmt.Println(p1)
    case 32767:
        if false {
            for fi24 := 0; fi24 < 6; fi24++ {
                v25 := []int{(p1 ^ p1)}
                _ = fi24
                _ = v25
            }
            switch (((("foo bar" + "\n\t") + "a") + ("foo bar" + "hello")) + ("\n\t" + "")) {
            case "\n\t":
                if (false || ((127 != 100) && true)) {
                    v26 := []int{(p1 >> ((0) & 63)), p1, p1}
                    p1 -= (p1 | (p1 / ((p1) | 1)))
                    var v27 int16 = 1000
                    _ = v26
                    _ = v27
                } else {
                    v28 := []int{}
                    for fi29 := 0; fi29 < 14; fi29++ {
                        switch (p1 - (fi29 ^ p1)) {
                        case 2147483647:
                            var v30 uint8 = 127
                            _ = v30
                        case -414:
                            v31 := GenS1{F0: -972, F1: 127}
                            _ = v31
                        case 32768:
                            if (-2.25 < 0.0) {
                                v32 := GenS3{F0: "foo bar", F1: 1}
                                var v33 string = "\n\t"
                                v34 := GenS1{F0: int16(p1), F1: 42}
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                        default:
                            for fi35 := 0; fi35 < 8; fi35++ {
                                fmt.Println(fi35)
                                _ = fi35
                            }
                        }
                        _ = fi29
                    }
                    v36 := GenS3{F0: ("a" + ("foo bar" + "foo bar")), F1: 255}
                    _ = v28
                    _ = v36
                }
            }
            p0 = 0
        } else {
            switch int(p0) {
            case 492:
                v37 := GenS4{F0: 4294967295}
                _ = v37
            case 4294967295:
                v38 := []int{}
                _ = v38
            }
            fmt.Println(p1)
            p1 = (p1 % (((p1 % ((p1) | 1))) | 1))
        }
    }
    p0 = 1
    var v39 int16 = -32768
    _ = p0
    _ = p1
    _ = v22
    _ = v39
    return 1
}

func main() {
    var v1 int8 = -1
    var v2 int = 1
    var v3 int32 = -2147483648
    v4 := []int{v2}
    var v40 uint64 = 1
    for fi41 := 0; fi41 < 15; fi41++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi42 := 0; fi42 < 3; fi42++ {
            if (0.0 <= ((((-2.25 + 3.14159) * -2.25) + (0.0 / (float64(float32(1e10)) * 1e10))) - ((0.1 + -2.25) * 3.14159))) {
                fmt.Println(v3, fi42, fi42)
                var v43 int64 = 9223372036854775807
                _ = v43
            }
            _ = fi42
        }
        switch "foo bar" {
        case "":
            v44 := GenS5{F0: 65535}
            _ = v44
        case "a":
            for fi45 := 0; fi45 < 16; fi45++ {
                v40 -= 42
                v4 = append(v4, fi41)
                if false {
                    var v46 int16 = 592
                    v3 = -1
                    v47 := GenS4{F0: 0}
                    _ = len(v4)
                    _ = v46
                    _ = v47
                }
                _ = fi45
            }
        case "foo bar":
            var v48 uint = 878
            _ = v48
        }
        _ = fi41
    }
    for fi49 := 0; fi49 < 1; fi49++ {
        switch "\n\t" {
        case "a":
            _ = len(v4)
        case "foo bar":
            switch 127 {
            case -32:
                v3 -= -140
            default:
                v50 := GenS6{F0: (((!true) && true) && false)}
                _ = v50
            }
        case "":
            var v51 int16 = -117
            _ = v51
        default:
            v40 += (v40 << ((65535) & 63))
        }
        _ = fi49
    }
    for fi52 := 0; fi52 < 18; fi52++ {
        var v53 uint64 = (v40 & v40)
        if len(v4) > 0 {
            v4[0] = 32768
        }
        _ = len(v4)
        _ = fi52
        _ = v53
    }
    fmt.Println(v1, v2)
    v54 := GenS2{F0: 42, F1: (v2 / (((v2 ^ 127)) | 1)), F2: 132}
    for fi55 := 0; fi55 < 12; fi55++ {
        var v56 float64 = 0.1
        v56 = v56
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi55
        _ = v56
    }
    switch (-(v2)) {
    case 0:
        v40 = v40
    default:
        for fi57 := 0; fi57 < 8; fi57++ {
            v3 = 1
            switch (-(fi57)) {
            case -1:
                for fi58 := 0; fi58 < 17; fi58++ {
                    v59 := GenS4{F0: 2}
                    switch "a" {
                    case "foo bar":
                        v60 := GenS2{F0: 1000, F1: 128, F2: genFn1(1000, (v2 + fi58))}
                        _ = v60
                    case "hello":
                        var v61 float32 = 3.14159
                        _ = v61
                    default:
                        v62 := GenS4{F0: 2}
                        _ = v62
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi58
                    _ = v59
                }
            }
            _ = fi57
        }
    }
    var v63 uint8 = 128
    fmt.Println(v1, v40, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v54
    _ = v63
}
