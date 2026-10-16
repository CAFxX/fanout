package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int8; F1 int; F2 bool }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 int16; F2 uint }
func genFn0(p0 int16, p1 bool) string {
    fmt.Println(p1)
    v6 := GenS1{F0: ("a" + ("\n\t" + ("" + "\n\t")))}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int8 = 1
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint32 = 65535
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            p1 = p1
            var v12 int16 = 763
            _ = fi11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    if true {
        v13 := GenS1{F0: ""}
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 int = fi14
            p0 = (p0 << ((255) & 63))
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        p0 -= 42
    }
    var v16 uint16 = 32767
    _ = p0
    _ = p1
    _ = v16
    _ = v6
    return "a"
}

func main() {
    var v1 uint64 = 774
    var v2 int8 = -124
    v3 := []uint32{}
    var v4 int64 = -772
    v5 := make(map[int]uint32)
    var v17 bool = ((((^(v1)) < v1) && (1 >= 38)) || (false || (false && (v1 > v1))))
    if v17 {
        if v17 {
            if true {
                var v18 uint16 = 395
                _ = v18
            }
            v2 -= 100
            _ = len(v3)
        }
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                if (fi20 <= 128) {
                    v21 := GenS1{F0: ("\n\t" + "foo bar")}
                    switch ("" + "foo bar") {
                    case "":
                        v22 := GenS2{F0: 255, F1: "foo bar"}
                        _ = v22
                    case "hello":
                        if true {
                            fmt.Println(fi20, fi19, fi20)
                            delete(v5, fi19)
                        }
                    default:
                        var v23 uint = 2
                        _ = v23
                    }
                    var v24 uint = 208
                    var v25 uint16 = 694
                    v17 = v17
                    _ = v21
                    _ = v24
                    _ = v25
                } else {
                    v4 = v4
                    v1 = (^(v1))
                    v2 = v2
                }
                v26 := GenS1{F0: ""}
                _ = fi20
                _ = v26
            }
            v3 = append(v3, 1)
            switch "foo bar" {
            case "hello":
                var v27 uint32 = 65535
                _ = v27
            }
            _ = fi19
        }
        v28 := GenS2{F0: 177, F1: ("a" + ("foo bar" + ""))}
        v29 := GenS3{F0: v2, F1: 2147483647, F2: v17}
        _ = v28
        _ = v29
    }
    var v30 float32 = -2.25
    for fi31 := 0; fi31 < 7; fi31++ {
        fmt.Println(v17, v30, v4)
        switch (fi31 & (fi31 | (fi31 << ((0) & 63)))) {
        case 9223372036854775807:
            if v17 {
                var v32 int32 = 1
                v33 := GenS2{F0: 204, F1: "foo bar"}
                v1 -= (^(v1))
                _ = v32
                _ = v33
            }
        default:
            v34 := GenS4{F0: v4, F1: 1}
            _ = v34
        }
        _ = fi31
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 uint32 = 42
        var v37 uint16 = 122
        _ = fi35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        fmt.Println(v2, fi38, fi38)
        _ = fi38
    }
    if v17 {
        for fi39 := 0; fi39 < 17; fi39++ {
            for fi40 := 0; fi40 < 17; fi40++ {
                v4 = (v4 ^ 711)
                v17 = (65535 >= 959)
                _ = fi40
            }
            _ = fi39
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            delete(v5, (^(fi41)))
            v30 = 3.14159
            _ = fi41
        }
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        if v17 {
            fmt.Println(v4, fi42)
            switch (("foo bar" + "") + "foo bar") {
            case "hello":
                var v43 int = fi42
                _ = v43
            case "foo bar":
                for fi44 := 0; fi44 < 10; fi44++ {
                    v4 = v4
                    _ = fi44
                }
            case "":
                if len(v3) > 0 {
                    v3[0] = 190
                }
            default:
                v2 -= (v2 / ((int8(v4)) | 1))
            }
        }
        switch 32768 {
        case 0:
            if ("" != "a") {
                var v45 uint = 727
                _ = len(v5)
                v46 := GenS5{F0: 2147483648}
                _ = v45
                _ = v46
            } else {
                if true {
                    v4 -= 9223372036854775807
                    fmt.Println(v30)
                    v30 = ((float32((1.5 / 3.14159)) + (((3.14159 / v30) * v30) * ((v30 / (v30 + v30)) + (0.0 + v30)))) - v30)
                }
                if ((float32(0.1) + 1.5) == (v30 / 1.0)) {
                    for fi47 := 0; fi47 < 17; fi47++ {
                        if ((!v17) || v17) {
                            var v48 bool = (702 == 65535)
                            _ = v48
                        } else {
                            v30 = v30
                            switch int(v4) {
                            case -796:
                                fmt.Println(v4, v4, v30)
                            case 2147483647:
                                switch (fi47 + (fi42 ^ fi47)) {
                                case 42:
                                    if false {
                                        v1 -= v1
                                        fmt.Println(fi42, v30)
                                        var v49 int8 = v2
                                        _ = v49
                                    }
                                case 770:
                                    var v50 uint64 = (v1 << ((42) & 63))
                                    v4 += -689
                                    fmt.Println(fi42, v1)
                                    _ = v50
                                default:
                                    v2 = v2
                                }
                            default:
                                var v51 uint16 = 65535
                                _ = v51
                            }
                            var v52 uint64 = (v1 / ((745) | 1))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v52
                        }
                        v53 := GenS6{F0: 144, F1: 0, F2: 42}
                        _ = fi47
                        _ = v53
                    }
                }
                var v54 int = fi42
                _ = v54
            }
        case 431:
            var v55 bool = false
            _ = v55
        case 9223372036854775807:
            _ = len(v3)
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
