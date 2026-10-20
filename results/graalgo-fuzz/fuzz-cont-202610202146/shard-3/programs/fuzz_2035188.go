package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint32; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint16; F1 uint32; F2 bool }
type GenS7 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int64; F1 int64; F2 uint64 }
func genFn0(p0 bool) uint64 {
    var v7 int16 = 42
    for fi8 := 0; fi8 < 20; fi8++ {
        if false {
            var v9 bool = p0
            _ = v9
        } else {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 int64 = -445
                v7 += 0
                _ = fi10
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: "hello", F1: (v7 | (v7 / (((v7 ^ -70)) | 1))), F2: 2147483648}
            v14 := GenS2{F0: -20}
            _ = fi12
            _ = v13
            _ = v14
        }
        p0 = true
        _ = fi8
    }
    var v15 float64 = 0.0
    _ = p0
    _ = v15
    _ = v7
    return 2
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]int)
    v3 := []int{}
    var v4 int8 = -128
    v5 := make(map[int]bool)
    var v6 uint16 = 267
    v1 = (v1 & (v1 >> ((uint(v6)) & 63)))
    if true {
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := GenS2{F0: 0}
            v4 += (v4 >> ((v1) & 63))
            for fi18 := 0; fi18 < 19; fi18++ {
                fmt.Println(fi16, v4)
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 uint = v1
            v4 -= (v4 + -1)
            _ = fi19
            _ = v20
        }
        var v21 int64 = -9223372036854775808
        _ = v21
    }
    fmt.Println(v1)
    var v22 int64 = 503
    v23 := GenS2{F0: v22}
    for fi24 := 0; fi24 < 7; fi24++ {
        switch 32767 {
        case 9223372036854775807:
            _ = len(v2)
        case 32768:
            v22 -= v22
        case 184:
            v22 = v22
        default:
            var v25 float32 = 1.5
            _ = v25
        }
        if (-1 != 0) {
            _ = len(v3)
        }
        var v26 int64 = v22
        _ = fi24
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        for fi28 := 0; fi28 < 13; fi28++ {
            v29 := GenS3{F0: v4, F1: 0, F2: v4}
            v22 -= (v22 + 0)
            _ = fi28
            _ = v29
        }
        switch ("foo bar" + (((("" + "hello") + "hello") + "hello") + "\n\t")) {
        case "\n\t":
            switch "\n\t" {
            case "hello":
                if len(v3) > 0 {
                    v3[0] = 169
                }
            case "\n\t":
                if true {
                    v30 := GenS1{F0: "a", F1: -368, F2: 4294967295}
                    v31 := GenS4{F0: (v4 << ((v1) & 63))}
                    _ = v30
                    _ = v31
                } else {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = int(v22)
                        }
                    }
                    for fi32 := 0; fi32 < 13; fi32++ {
                        v4 = v4
                        _ = fi32
                    }
                    if true {
                        fmt.Println(v4, v4)
                        _ = len(v3)
                    } else {
                        for fi33 := 0; fi33 < 8; fi33++ {
                            fmt.Println(v22)
                            delete(v5, (-(fi27)))
                            _ = fi33
                        }
                        v34 := GenS5{F0: 42, F1: true}
                        var v35 int = 1
                        _ = v34
                        _ = v35
                    }
                }
            default:
                v6 = (v6 << ((uint(v6)) & 63))
            }
        }
        v36 := GenS6{F0: 1, F1: 970, F2: ((float32((3.14159 * (0.0 * (-2.25 - (0.1 / 1e10))))) / 1.0) >= (float32(1e10) * 3.14159))}
        _ = fi27
        _ = v36
    }
    var v37 bool = false
    v1 -= (v1 / (((v1 - (^(v1)))) | 1))
    var v38 uint8 = 0
    for fi39 := 0; fi39 < 8; fi39++ {
        if len(v3) > 0 {
            v3[0] = (fi39 % ((fi39) | 1))
        }
        _ = len(v5)
        _ = fi39
    }
    if v37 {
        v4 -= int8(v6)
    } else {
        var v40 int16 = -1
        if v37 {
            if true {
                var v41 int64 = 1
                _ = v41
            }
            v3 = append(v3, 42)
            var v42 int16 = 1
            _ = v42
        }
        switch 128 {
        case 42:
            if v37 {
                fmt.Println(v37, v37, v22)
                var v43 float64 = 1e10
                _ = v43
            } else {
                if true {
                    for fi44 := 0; fi44 < 7; fi44++ {
                        switch -9223372036854775808 {
                        case 42:
                            fmt.Println(v37, v6, v22)
                        case 846:
                            v45 := GenS3{F0: (v4 & -116), F1: 4294967295, F2: (v4 & (^(v4)))}
                            _ = v45
                        case -1:
                            v5[(fi44 | (fi44 >> (((v1 & v1)) & 63)))] = v37
                        default:
                            v38 -= 0
                        }
                        _ = fi44
                    }
                    if (0.0 < 0.0) {
                        fmt.Println(v1, v6, v6)
                    }
                }
            }
        }
        _ = v40
    }
    if (!((0 >= -2147483648) && true)) {
        if v37 {
            for fi46 := 0; fi46 < 6; fi46++ {
                var v47 uint32 = 2147483648
                _ = fi46
                _ = v47
            }
        } else {
            v22 -= (v22 >> ((v1) & 63))
        }
        v38 -= (v38 - v38)
    }
    for fi48 := 0; fi48 < 13; fi48++ {
        var v49 int = (fi48 << (((v1 * v1)) & 63))
        _ = fi48
        _ = v49
    }
    if (((!(true || v37)) || v37) && (0.0 > 1e10)) {
        v50 := GenS7{F0: (435 <= 42), F1: 65535, F2: genFn0(false)}
        _ = v50
    } else {
        if ("hello" != (("foo bar" + (("hello" + "a") + "\n\t")) + (("foo bar" + "") + "a"))) {
            var v51 int64 = v22
            for fi52 := 0; fi52 < 1; fi52++ {
                var v53 int = 1
                var v54 int32 = -794
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = v51
        }
        switch "\n\t" {
        case "foo bar":
            v6 += 42
        case "\n\t":
            for fi55 := 0; fi55 < 8; fi55++ {
                fmt.Println(v4)
                v56 := GenS8{F0: (v22 | (v22 - -355)), F1: 1, F2: uint64(v1)}
                fmt.Println(v38, v4)
                _ = fi55
                _ = v56
            }
        case "":
            var v57 int64 = (v22 >> (((v1 % ((42) | 1))) & 63))
            _ = v57
        }
        switch -702 {
        case 4294967295:
            var v58 int32 = 2147483647
            _ = v58
        case 128:
            v4 = -1
        }
    }
    var v59 int8 = (v4 - v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
