package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int; F2 uint16 }
type GenS5 struct { F0 int; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 int32; F2 int }
type GenS8 struct { F0 string; F1 uint16; F2 uint32 }
func genFn0(p0 int16) int32 {
    switch "hello" {
    case "a":
        p0 += p0
    case "hello":
        for fi6 := 0; fi6 < 3; fi6++ {
            p0 = (^(p0))
            _ = fi6
        }
    case "foo bar":
        var v7 int16 = (p0 | (p0 * p0))
        _ = v7
    default:
        if true {
            fmt.Println(p0)
        }
    }
    var v8 uint8 = 128
    var v9 int16 = p0
    v10 := []int{9223372036854775807, 42, 2147483647}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 42
}
func genFn1() uint64 {
    switch 2147483647 {
    case 389:
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(fi11)
            _ = fi11
        }
    case 42:
        var v12 bool = false
        _ = v12
    case -257:
        if ("a" != "a") {
            if (true || false) {
                v13 := GenS1{F0: 0}
                var v14 int64 = -149
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := GenS2{F0: (214 == 0), F1: 416}
                var v17 int64 = -687
                v18 := []int{(fi15 + -2147483648), (fi15 << ((0) & 63)), (-(fi15)), -665}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        } else {
            v19 := []int{-560, 385, 2147483647}
            if len(v19) > 0 {
                v19[0] = 32767
            }
            _ = v19
        }
    }
    v20 := GenS1{F0: 664}
    if ("\n\t" == "hello") {
        var v21 float64 = (1.5 * -2.25)
        var v22 uint16 = 932
        switch 42 {
        case 32767:
            v23 := GenS3{F0: 0}
            _ = v23
        case 128:
            fmt.Println(v22)
        case 127:
            if ((!(!(0 == -9223372036854775808))) || ((118 > -127) && ((1 < v22) || (!true)))) {
                v24 := GenS2{F0: true, F1: 2147483647}
                v25 := GenS2{F0: (0 >= 42), F1: genFn0(-32768)}
                v26 := []int{42}
                _ = v24
                _ = v25
                _ = v26
            }
        }
        _ = v21
        _ = v22
    }
    if (!true) {
        v27 := GenS2{F0: ((false || true) && true), F1: 225}
        _ = v27
    }
    _ = v20
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]string)
    var v3 int32 = 2147483647
    var v4 uint8 = 2
    var v5 uint64 = 687
    var v28 int64 = -991
    for fi29 := 0; fi29 < 9; fi29++ {
        fmt.Println(v4)
        _ = fi29
    }
    v3 = (-(v3))
    fmt.Println(v4, v1)
    for fi30 := 0; fi30 < 2; fi30++ {
        var v31 uint16 = 32767
        if true {
            switch ("hello" + "") {
            case "":
                var v32 string = ("foo bar" + "hello")
                _ = v32
            default:
                if (((v1 & (v1 * (v31 << ((2) & 63)))) > v1) || false) {
                    v33 := GenS4{F0: 2, F1: (fi30 << ((42) & 63)), F2: (v1 - (v31 | v31))}
                    for fi34 := 0; fi34 < 1; fi34++ {
                        v35 := GenS5{F0: fi30, F1: -128, F2: (v5 & (v5 >> ((2) & 63)))}
                        v2["\n\t"] = "\n\t"
                        _ = fi34
                        _ = v35
                    }
                    var v36 int64 = -9223372036854775808
                    _ = v33
                    _ = v36
                }
            }
            if (false || true) {
                for fi37 := 0; fi37 < 6; fi37++ {
                    v31 = (v1 ^ v31)
                    _ = fi37
                }
            } else {
                switch 2147483647 {
                case -125:
                    fmt.Println(v31)
                }
            }
            v38 := GenS5{F0: 32767, F1: 127, F2: (v5 >> ((1) & 63))}
            _ = v38
        } else {
            fmt.Println(v31, v3)
            fmt.Println(v1, v4)
            v3 += (^(v3))
        }
        _ = fi30
        _ = v31
    }
    var v39 uint16 = v1
    var v40 int16 = 1
    var v41 uint16 = (v39 | v39)
    var v42 uint16 = 65535
    var v43 uint16 = 2
    v28 -= -9223372036854775808
    var v44 float64 = 0.0
    if false {
        fmt.Println(v41, v3)
    } else {
        if true {
            v4 += v4
            if (v4 == (v4 % ((v4) | 1))) {
                for fi45 := 0; fi45 < 18; fi45++ {
                    v46 := v2["hello"]
                    if false {
                        var v47 int = fi45
                        var v48 int = fi45
                        _ = v47
                        _ = v48
                    }
                    _ = fi45
                    _ = v46
                }
            } else {
                v43 = v42
                v49 := GenS6{F0: 998}
                _ = v49
            }
            v50 := GenS1{F0: v5}
            _ = v50
        }
        switch ("" + "\n\t") {
        case "hello":
            v51 := []int{-473, 127, 879}
            _ = v51
        case "":
            for fi52 := 0; fi52 < 15; fi52++ {
                var v53 int = fi52
                _ = fi52
                _ = v53
            }
        case "\n\t":
            v54 := []int{int(v3), 4294967295}
            _ = v54
        default:
            for fi55 := 0; fi55 < 14; fi55++ {
                fmt.Println(v1)
                _ = fi55
            }
        }
    }
    v44 = v44
    switch ("foo bar" + ("" + ("" + "foo bar"))) {
    case "foo bar":
        v56 := GenS7{F0: ((v3 / ((v3) | 1)) >= (v3 % ((42) | 1))), F1: -437, F2: -804}
        _ = v56
    case "\n\t":
        var v57 uint16 = 65535
        _ = v57
    default:
        v58 := GenS8{F0: "a", F1: (v42 - v43), F2: 710}
        _ = v58
    }
    fmt.Println(v4, v3, v1)
    if ((v40 & 1000) == (v40 / ((-32768) | 1))) {
        v4 = v4
        v5 += (v5 + 610)
        for fi59 := 0; fi59 < 6; fi59++ {
            v1 += 2
            if false {
                v28 = 1
            } else {
                var v60 uint8 = (v4 + 242)
                _ = v60
            }
            _ = fi59
        }
    }
    var v61 int32 = (v3 * (v3 * (v3 << ((0) & 63))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v61
}
