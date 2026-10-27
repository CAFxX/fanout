package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 uint64, p1 uint8) uint8 {
    var v7 float32 = 1e10
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 float64 = (0.1 + 1.5)
        for fi10 := 0; fi10 < 18; fi10++ {
            p0 = (p0 + (p0 << ((42) & 63)))
            if (false && true) {
                v11 := GenS1{F0: int(p1), F1: false, F2: 100}
                _ = v11
            }
            v7 -= ((1.5 / float32(1.5)) + (float32(((3.14159 / v9) - -2.25)) + 3.14159))
            _ = fi10
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint16 = 1
            _ = fi12
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    if false {
        p1 += 1
        if false {
            v14 := []int{4294967295}
            if (!false) {
                switch 32768 {
                case 9223372036854775807:
                    var v15 int64 = 9223372036854775807
                    _ = v15
                case 0:
                    fmt.Println(v7, p0)
                case 338:
                    p0 = (p0 * (^(p0)))
                default:
                    switch 128 {
                    case -9223372036854775808:
                        p1 -= (p1 / ((1) | 1))
                    default:
                        for fi16 := 0; fi16 < 14; fi16++ {
                            fmt.Println(p0)
                            v7 += 3.14159
                            v17 := GenS1{F0: int(p1), F1: true, F2: 79}
                            _ = fi16
                            _ = v17
                        }
                    }
                }
                for fi18 := 0; fi18 < 10; fi18++ {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        fmt.Println(fi19)
                        _ = fi19
                    }
                    v20 := GenS1{F0: fi18, F1: (("hello" == "\n\t") || (!(p0 <= 725))), F2: 127}
                    _ = fi18
                    _ = v20
                }
                if (("\n\t" + "foo bar") == "hello") {
                    var v21 uint64 = 107
                    _ = v21
                }
            } else {
                v7 = -2.25
                for fi22 := 0; fi22 < 8; fi22++ {
                    fmt.Println(p1, p0, v7)
                    if len(v14) > 0 {
                        v14[0] = (fi22 / (((^(fi22))) | 1))
                    }
                    _ = fi22
                }
                if true {
                    var v23 uint = 65535
                    var v24 string = (("hello" + ("hello" + "")) + (("\n\t" + "\n\t") + (("foo bar" + ("\n\t" + "a")) + "")))
                    _ = v23
                    _ = v24
                } else {
                    p0 += p0
                    var v25 float32 = 0.0
                    _ = v25
                }
            }
            _ = v14
        }
    } else {
        v26 := GenS1{F0: 1, F1: false, F2: -18}
        _ = v26
    }
    p0 = (p0 << ((4294967295) & 63))
    v7 = (v7 - (0.1 + (3.14159 * 1e10)))
    _ = p0
    _ = p1
    _ = v7
    return p1
}
func genFn1() int64 {
    var v27 uint16 = 367
    if false {
        var v28 uint64 = 1
        v28 = 570
        v29 := GenS2{F0: 65535, F1: 0}
        _ = v28
        _ = v29
    } else {
        v27 = v27
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := GenS3{F0: ("\n\t" + "a")}
            _ = fi30
            _ = v31
        }
    }
    switch "" {
    case "foo bar":
        v32 := []int{798, -803, 32767, 32768}
        _ = v32
    case "\n\t":
        v27 += 0
    case "hello":
        v27 += v27
    default:
        var v33 uint16 = (v27 % (((v27 - v27)) | 1))
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        var v35 int8 = -1
        v36 := []int{(-(fi34)), fi34}
        _ = fi34
        _ = v35
        _ = v36
    }
    v27 += (v27 ^ (v27 << ((4294967295) & 63)))
    _ = v27
    return 0
}

func main() {
    var v1 bool = (2 <= 0)
    var v2 uint = 2
    v3 := []uint32{1, 964}
    v4 := make(map[uint32]bool)
    var v5 int64 = 0
    var v6 uint8 = 0
    v6 = 47
    var v37 int = 268
    switch ("hello" + "foo bar") {
    case "foo bar":
        v1 = v1
    default:
        var v38 uint32 = 65535
        _ = v38
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 float32 = 0.1
        v41 := GenS4{F0: "a", F1: (v1 || true)}
        _ = fi39
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 15; fi42++ {
        var v43 int8 = 0
        v43 = -1
        _ = fi42
        _ = v43
    }
    fmt.Println(v5, v6)
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        delete(v4, 870)
    case "a":
        v1 = (!(-1 != (^(v5))))
    case "\n\t":
        _ = len(v4)
    default:
        for fi44 := 0; fi44 < 9; fi44++ {
            fmt.Println(v37, v2)
            _ = fi44
        }
    }
    if v1 {
        v37 = int(v5)
    } else {
        v1 = (!v1)
        for fi45 := 0; fi45 < 15; fi45++ {
            if (-97 != 0) {
                var v46 int32 = -2147483648
                v47 := GenS4{F0: "\n\t", F1: ((1.5 - (float32((-2.25 + 1e10)) - 3.14159)) != 1e10)}
                switch "a" {
                case "foo bar":
                    var v48 int32 = v46
                    _ = v48
                case "\n\t":
                    v5 -= v5
                default:
                    var v49 uint64 = 781
                    _ = v49
                }
                _ = v46
                _ = v47
            } else {
                if (v2 > (v2 * v2)) {
                    var v50 int64 = 130
                    _ = v50
                }
                v5 += genFn1()
                var v51 uint32 = 42
                _ = v51
            }
            var v52 int16 = -32768
            _ = fi45
            _ = v52
        }
        v53 := v4[2147483647]
        _ = v53
    }
    var v54 uint8 = v6
    if (!false) {
        if v1 {
            v6 = v54
        }
    }
    var v55 int16 = 42
    var v56 bool = v1
    v1 = v56
    v5 -= genFn1()
    var v57 float64 = 1.5
    if v1 {
        var v58 uint16 = 255
        _ = v58
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}
