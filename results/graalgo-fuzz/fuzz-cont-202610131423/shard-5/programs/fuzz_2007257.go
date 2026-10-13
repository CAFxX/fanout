package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 int64) uint8 {
    fmt.Println(p0)
    p0 += (p0 + (p0 ^ (p0 % (((p0 << ((p0) & 63))) | 1))))
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = -32768
    var v3 float64 = -2.25
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 int64 = 42
        v6 := []int{}
        if (!false) {
            v7 := GenS1{F0: (!(1 == 1)), F1: 2147483648, F2: 0}
            _ = v7
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    var v8 int32 = -1
    v9 := GenS1{F0: true, F1: 367, F2: 42}
    v10 := GenS2{F0: 154, F1: int8(v8), F2: v2}
    v11 := []int{128, -469}
    var v12 float64 = v3
    var v13 uint64 = 2
    v3 += (3.14159 + 1e10)
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int = 32767
        var v16 int64 = -1
        for fi17 := 0; fi17 < 14; fi17++ {
            if false {
                var v18 float64 = (1.5 - v3)
                switch (("foo bar" + "foo bar") + ("a" + "hello")) {
                case "a":
                    _ = len(v11)
                case "hello":
                    switch fi14 {
                    case 2147483647:
                        switch (fi17 / (((-(fi14))) | 1)) {
                        case -2147483648:
                            switch (fi17 + (v15 - fi14)) {
                            case 362:
                                fmt.Println(v13)
                            }
                        default:
                            v2 = v2
                        }
                    case -2147483648:
                        fmt.Println(v2, v2, v8)
                    case -226:
                        fmt.Println(fi14)
                    default:
                        if (((true || false) || false) || (!(!(!(!false))))) {
                            if ((fi14 | (-(fi17))) <= (fi14 - 176)) {
                                fmt.Println(fi14, v18)
                                if (-128 >= 42) {
                                    var v19 int8 = 1
                                    switch ("a" + "hello") {
                                    case "foo bar":
                                        v3 = v18
                                    }
                                    if false {
                                        var v20 int = (v15 << ((42) & 63))
                                        _ = v20
                                    } else {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v21 int8 = v19
                                            _ = v21
                                        case "a":
                                            v3 = v18
                                        default:
                                            switch 4294967295 {
                                            case -732:
                                                v3 = ((v12 / 1.0) - v18)
                                            default:
                                                var v22 int64 = (v16 | v16)
                                                _ = v22
                                            }
                                        }
                                        v23 := GenS2{F0: 255, F1: 100, F2: v2}
                                        _ = v23
                                    }
                                    _ = v19
                                } else {
                                    var v24 uint8 = 129
                                    if true {
                                        var v25 int8 = 127
                                        fmt.Println(v8, v8)
                                        v26 := GenS1{F0: (!(((v2 + v2) >= v2) || true)), F1: 2147483647, F2: 100}
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v24
                                }
                                var v27 uint = 4294967295
                                _ = v27
                            } else {
                                if (!true) {
                                    if true {
                                        fmt.Println(v13)
                                        v11 = append(v11, (fi14 & fi14))
                                        var v28 uint16 = 1
                                        _ = v28
                                    }
                                    var v29 int8 = -27
                                    var v30 uint = 848
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 int16 = (v2 + (v2 ^ v2))
                                    fmt.Println(v16, v3, fi17)
                                    v32 := GenS2{F0: 0, F1: 42, F2: 1000}
                                    _ = v31
                                    _ = v32
                                }
                            }
                        } else {
                            if (42 >= 1) {
                                switch v15 {
                                case -557:
                                    if len(v11) > 0 {
                                        v11[0] = v15
                                    }
                                case -296:
                                    if len(v11) > 0 {
                                        v11[0] = int(v2)
                                    }
                                default:
                                    if true {
                                        v8 -= 42
                                        switch (fi14 + (v15 % (((^(v15))) | 1))) {
                                        case 4294967295:
                                            var v33 uint = 42
                                            _ = v33
                                        default:
                                            var v34 uint32 = 667
                                            _ = v34
                                        }
                                        switch ("" + "foo bar") {
                                        case "":
                                            var v35 int = fi14
                                            _ = v35
                                        case "a":
                                            switch -2147483648 {
                                            case 4294967295:
                                                v36 := GenS1{F0: (!(763 >= (v13 + 42))), F1: 4294967295, F2: 100}
                                                _ = v36
                                            case -914:
                                                switch "foo bar" {
                                                case "a":
                                                    var v37 int = (-(v15))
                                                    _ = v37
                                                case "":
                                                    switch v15 {
                                                    case 9223372036854775807:
                                                        v38 := GenS2{F0: 255, F1: -54, F2: (v2 | (^(v2)))}
                                                        _ = v38
                                                    case 4294967295:
                                                        v2 += v2
                                                    default:
                                                        var v39 int64 = -156
                                                        _ = v39
                                                    }
                                                case "foo bar":
                                                    v40 := GenS2{F0: 65535, F1: -46, F2: int16(fi17)}
                                                    _ = v40
                                                }
                                                switch 0 {
                                                case -1:
                                                    fmt.Println(fi17)
                                                case -9223372036854775808:
                                                    var v41 uint8 = 125
                                                    _ = v41
                                                }
                                                switch (("foo bar" + "") + ("\n\t" + "foo bar")) {
                                                case "":
                                                    var v42 uint16 = 1
                                                    _ = v42
                                                case "\n\t":
                                                    v43 := GenS2{F0: 914, F1: -126, F2: 42}
                                                    _ = v43
                                                }
                                                fmt.Println(v12, v16)
                                            }
                                        case "\n\t":
                                            if false {
                                                v2 = v2
                                                var v44 bool = false
                                                _ = v44
                                            } else {
                                                v15 = 0
                                                var v45 int16 = (v2 & 32767)
                                                _ = v45
                                            }
                                        }
                                    } else {
                                        if ((v13 ^ (v13 / (((^(v13))) | 1))) < 42) {
                                            v13 -= v13
                                            v13 = (v13 ^ (v13 >> ((4294967295) & 63)))
                                            switch (-(fi17)) {
                                            case 212:
                                                v13 = v13
                                            case 32767:
                                                if false {
                                                    if true {
                                                        v46 := GenS3{F0: 255}
                                                        v12 -= float64(((((float32(v3) / (float32(-2.25) * -2.25)) + 1.5) / (1e10 + -2.25)) - 1.5))
                                                        _ = v46
                                                    }
                                                    v2 = v2
                                                }
                                            }
                                        } else {
                                            switch int(v13) {
                                            case 645:
                                                var v47 uint64 = (v13 | (v13 / (((v13 + (v13 % ((v13) | 1)))) | 1)))
                                                _ = v47
                                            default:
                                                var v48 uint8 = 42
                                                _ = v48
                                            }
                                            var v49 uint8 = 0
                                            _ = v49
                                        }
                                    }
                                }
                            }
                            v8 = v8
                        }
                    }
                }
                v50 := GenS1{F0: (-107 == 100), F1: 839, F2: 1}
                _ = v18
                _ = v50
            } else {
                v51 := GenS1{F0: (!false), F1: 2147483647, F2: 42}
                _ = v51
            }
            if true {
                v3 += v12
            }
            _ = fi17
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
