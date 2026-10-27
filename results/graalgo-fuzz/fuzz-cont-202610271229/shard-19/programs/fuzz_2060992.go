package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0(p0 int32, p1 uint32) int32 {
    v4 := []int{-657, 4294967295}
    fmt.Println(p0, p1)
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    _ = p0
    _ = p1
    _ = v4
    return p0
}
func genFn1(p0 uint16) uint {
    var v5 bool = (0.0 > float32(float64(((1e10 / (1e10 / 1.0)) + 0.1))))
    var v6 int8 = 100
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := []int{fi7, (fi8 << ((4294967295) & 63))}
            fmt.Println(v6, fi7)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if v5 {
        var v10 int32 = 313
        fmt.Println(v10)
        _ = v10
    } else {
        v5 = ((1 < 42) || (9223372036854775807 <= -2147483648))
    }
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint = 42
    var v2 uint64 = 42
    var v3 int = 228
    v11 := []int{}
    var v12 int16 = 32767
    if false {
        var v13 int8 = int8(v2)
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v3 = (fi14 - -2147483648)
        _ = fi14
    }
    _ = len(v11)
    var v15 int64 = 0
    if len(v11) > 0 {
        v11[0] = 127
    }
    if true {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if false {
            switch (-(v3)) {
            case 9223372036854775807:
                switch "\n\t" {
                case "\n\t":
                    var v16 float32 = 3.14159
                    _ = v16
                case "foo bar":
                    for fi17 := 0; fi17 < 15; fi17++ {
                        if (false || true) {
                            var v18 string = "a"
                            _ = v18
                        }
                        var v19 int8 = -128
                        v3 = (fi17 << ((1) & 63))
                        _ = fi17
                        _ = v19
                    }
                case "a":
                    v3 += 2147483647
                }
            default:
                if ((!false) || (0.1 > 0.0)) {
                    v12 = v12
                }
            }
            if (!(v12 <= -32768)) {
                var v20 int64 = (v15 % (((v15 | v15)) | 1))
                if true {
                    if len(v11) > 0 {
                        v11[0] = (-(v3))
                    }
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    fmt.Println(v20)
                }
                var v21 float32 = float32(float64(1e10))
                _ = v20
                _ = v21
            }
        } else {
            v15 -= (v15 * v15)
        }
    } else {
        switch 127 {
        case 128:
            var v22 uint = v1
            _ = v22
        case 42:
            var v23 uint = genFn1(1)
            _ = v23
        case 1:
            var v24 float32 = 0.0
            _ = v24
        default:
            switch v3 {
            case 1:
                switch "" {
                case "a":
                    _ = len(v11)
                case "\n\t":
                    var v25 uint16 = uint16(v1)
                    _ = v25
                case "hello":
                    var v26 uint8 = 131
                    _ = v26
                default:
                    fmt.Println(v3)
                }
            case -241:
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 uint32 = 2147483647
                    _ = fi27
                    _ = v28
                }
            }
        }
        v11 = append(v11, (v3 ^ 9223372036854775807))
    }
    if (v12 > (v12 | (v12 - (v12 / (((v12 % ((v12) | 1))) | 1))))) {
        v29 := GenS1{F0: v1, F1: 1}
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS2{F0: -32768, F1: -864}
        var v32 uint8 = 0
        _ = fi30
        _ = v31
        _ = v32
    }
    if true {
        switch (v3 ^ (v3 ^ v3)) {
        case -1:
            for fi33 := 0; fi33 < 16; fi33++ {
                var v34 float64 = 1e10
                var v35 int16 = (^(v12))
                v1 = 42
                _ = fi33
                _ = v34
                _ = v35
            }
        default:
            fmt.Println(v2, v1)
        }
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 int16 = (v12 * 0)
            _ = fi36
            _ = v37
        }
        if (!true) {
            var v38 uint32 = 4294967295
            for fi39 := 0; fi39 < 11; fi39++ {
                for fi40 := 0; fi40 < 17; fi40++ {
                    _ = len(v11)
                    _ = fi40
                }
                _ = fi39
            }
            v2 = v2
            _ = v38
        } else {
            var v41 int = 127
            _ = v41
        }
    } else {
        v11 = append(v11, -140)
        for fi42 := 0; fi42 < 15; fi42++ {
            fmt.Println(v15, v3, v15)
            for fi43 := 0; fi43 < 7; fi43++ {
                v12 -= v12
                _ = fi43
            }
            v15 += (v15 + v15)
            _ = fi42
        }
        var v44 int8 = -1
        _ = v44
    }
    if true {
        v45 := GenS2{F0: (v12 << (((v1 << (((v1 + v1)) & 63))) & 63)), F1: (v15 << ((genFn1(270)) & 63))}
        if len(v11) > 0 {
            v11[0] = (^(v3))
        }
        _ = v45
    } else {
        fmt.Println(v15, v15)
        v11 = append(v11, v3)
    }
    var v46 float64 = 1e10
    v47 := GenS1{F0: (^(v1)), F1: v12}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println(v1)
    for fi48 := 0; fi48 < 1; fi48++ {
        var v49 bool = false
        var v50 int = fi48
        if false {
            for fi51 := 0; fi51 < 5; fi51++ {
                fmt.Println(v49)
                _ = fi51
            }
            if v49 {
                switch "a" {
                case "a":
                    for fi52 := 0; fi52 < 2; fi52++ {
                        var v53 int64 = v15
                        _ = fi52
                        _ = v53
                    }
                }
                fmt.Println(v1)
            }
        } else {
            switch ("hello" + (("\n\t" + ("foo bar" + "\n\t")) + "hello")) {
            case "hello":
                var v54 int32 = 0
                _ = v54
            case "":
                for fi55 := 0; fi55 < 19; fi55++ {
                    v12 = 5
                    _ = len(v11)
                    if ((v46 / v46) == (v46 * (float64(float32(v46)) * v46))) {
                        var v56 uint64 = v2
                        switch fi48 {
                        case -9223372036854775808:
                            if ((v2 * (v56 >> (((v1 >> ((v1) & 63))) & 63))) >= v2) {
                                if (-1 == 0) {
                                    if len(v11) > 0 {
                                        v11[0] = fi48
                                    }
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    fmt.Println(v3, v49)
                                }
                            }
                        }
                        _ = v56
                    } else {
                        v57 := GenS2{F0: v12, F1: -9223372036854775808}
                        var v58 int16 = v12
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        fmt.Println(fi55, v46, v49)
                        var v59 uint8 = 240
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                    var v60 uint32 = 4294967295
                    _ = fi55
                    _ = v60
                }
            case "foo bar":
                var v61 int = 0
                _ = v61
            default:
                v62 := GenS1{F0: (^(v1)), F1: (v12 >> ((genFn1(32767)) & 63))}
                _ = v62
            }
        }
        _ = fi48
        _ = v49
        _ = v50
    }
    var v63 uint = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v63
}
