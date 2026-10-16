package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint32; F1 uint16; F2 int }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int; F2 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int; F1 int16 }
type GenS10 struct { F0 string; F1 uint64 }
type GenS11 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0(p0 bool, p1 uint64, p2 uint32) uint16 {
    v7 := []int{42, 12}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: (("a" + ("foo bar" + ("a" + "hello"))) + ""), F1: 128}
        _ = fi8
        _ = v9
    }
    var v10 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 207
}

func main() {
    v1 := []uint32{42, 2147483647}
    var v2 uint16 = 0
    v3 := []string{"a", "foo bar"}
    var v4 int = -1
    v5 := []string{"a", "\n\t", "a"}
    var v6 string = "a"
    v11 := GenS1{F0: "foo bar", F1: 2}
    v6 = v6
    if false {
        switch (v6 + v6) {
        case "":
            if false {
                if false {
                    v6 = ("" + v6)
                    if (213 == 1) {
                        for fi12 := 0; fi12 < 14; fi12++ {
                            var v13 string = "\n\t"
                            v2 += 2
                            _ = fi12
                            _ = v13
                        }
                    } else {
                        var v14 float32 = (-2.25 * 0.0)
                        for fi15 := 0; fi15 < 12; fi15++ {
                            var v16 float32 = ((v14 / (v14 + 0.0)) - v14)
                            v6 = v6
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = len(v3)
                }
            } else {
                if len(v5) > 0 {
                    v5[0] = ""
                }
            }
        case "\n\t":
            fmt.Println(v6, v2, v2)
        }
        if ((!false) || false) {
            for fi17 := 0; fi17 < 13; fi17++ {
                if (false && (!false)) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 float64 = 0.0
                        var v20 bool = (4294967295 == 0)
                        var v21 int16 = 1
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    if true {
                        var v22 uint = 760
                        _ = v22
                    } else {
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        }
                    }
                } else {
                    var v23 int = v4
                    var v24 int64 = 1
                    v4 -= (-(v23))
                    _ = v23
                    _ = v24
                }
                if false {
                    v2 = (v2 + (v2 | v2))
                }
                _ = fi17
            }
            var v25 float32 = -2.25
            v2 = (v2 + 0)
            _ = v25
        }
        fmt.Println(v4, v4)
    } else {
        if (!true) {
            if false {
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 int64 = -9223372036854775808
                    _ = fi26
                    _ = v27
                }
                v2 -= v2
            } else {
                if false {
                    v6 = ((("hello" + "a") + v6) + "hello")
                } else {
                    v28 := GenS2{F0: true, F1: (v4 % (((v4 + 1)) | 1))}
                    _ = v28
                }
            }
            if false {
                switch v6 {
                case "\n\t":
                    v29 := GenS1{F0: v6, F1: 2}
                    _ = v29
                case "a":
                    var v30 uint32 = 4294967295
                    _ = v30
                case "hello":
                    var v31 int16 = 358
                    _ = v31
                }
            }
            v32 := GenS2{F0: (-1 < 1), F1: -1}
            _ = v32
        } else {
            v33 := GenS3{F0: 699, F1: v2, F2: (v4 | v4)}
            v1 = append(v1, 2147483648)
            _ = v33
        }
        for fi34 := 0; fi34 < 8; fi34++ {
            if false {
                switch (fi34 * fi34) {
                case 127:
                    if false {
                        v6 = ((v6 + (v6 + v6)) + (v6 + ""))
                    }
                default:
                    switch (v4 + (fi34 - (v4 | (-(v4))))) {
                    case -855:
                        for fi35 := 0; fi35 < 14; fi35++ {
                            v4 = (fi34 - -2147483648)
                            v6 = "hello"
                            _ = fi35
                        }
                    case 32767:
                        v36 := GenS4{F0: (fi34 / (((v4 >> ((0) & 63))) | 1)), F1: v2}
                        _ = v36
                    case 9223372036854775807:
                        v37 := GenS5{F0: 400, F1: (v4 ^ (v4 ^ (v4 >> ((1) & 63)))), F2: -1}
                        _ = v37
                    }
                }
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                v2 += 0
            }
            var v38 int64 = 893
            switch v6 {
            case "\n\t":
                if true {
                    v38 += (v38 / ((v38) | 1))
                }
            case "a":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi34
            _ = v38
        }
        for fi39 := 0; fi39 < 7; fi39++ {
            v2 = genFn0(true, 42, 2147483647)
            v2 += v2
            v40 := GenS3{F0: 2, F1: (v2 * (v2 - (^(v2)))), F2: v4}
            _ = fi39
            _ = v40
        }
    }
    v1 = append(v1, 2147483648)
    switch (v4 - v4) {
    case -9223372036854775808:
        switch v4 {
        case 9223372036854775807:
            v2 -= genFn0(((v4 & v4) != v4), 675, 1)
        case 127:
            v41 := GenS6{F0: 128}
            _ = v41
        case 128:
            v42 := GenS7{F0: 18446744073709551615, F1: v4}
            _ = v42
        }
    case -967:
        var v43 uint64 = 18446744073709551615
        _ = v43
    }
    if (0 == 2147483647) {
        if len(v3) > 0 {
            v3[0] = v6
        }
        for fi44 := 0; fi44 < 15; fi44++ {
            switch v6 {
            case "foo bar":
                v1 = append(v1, 2147483648)
            case "\n\t":
                if true {
                    fmt.Println(fi44, v4)
                }
            case "hello":
                v4 = (fi44 >> ((4294967295) & 63))
            default:
                _ = len(v5)
            }
            _ = fi44
        }
        switch (v6 + ("\n\t" + "foo bar")) {
        case "hello":
            v45 := GenS8{F0: 4294967295}
            _ = v45
        case "foo bar":
            v46 := GenS9{F0: (-(v4)), F1: 745}
            _ = v46
        case "":
            v4 = (v4 + v4)
        }
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        if (!(-2147483648 == 1)) {
            if len(v5) > 0 {
                v5[0] = (v6 + v6)
            }
            var v48 float32 = 0.0
            _ = v48
        }
        _ = fi47
    }
    v49 := GenS10{F0: v6, F1: 2}
    v50 := GenS3{F0: 2147483648, F1: genFn0((255 != 740), 624, 160), F2: v4}
    for fi51 := 0; fi51 < 4; fi51++ {
        if ((0 != 42) || false) {
            var v52 int16 = 0
            var v53 uint16 = v2
            v54 := GenS11{F0: 255, F1: 2147483647, F2: -812}
            _ = v52
            _ = v53
            _ = v54
        }
        _ = len(v1)
        _ = fi51
    }
    if true {
        fmt.Println(v4, v6, v2)
        var v55 int16 = 0
        fmt.Println(v55)
        _ = v55
    }
    var v56 float32 = (float32(0.0) * 0.0)
    if (true || (true && false)) {
        v2 = v2
        var v57 uint64 = 979
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v57
    }
    v56 += float32(float64(0.0))
    var v58 bool = (35 > 127)
    v56 = ((((-2.25 * v56) / ((((3.14159 - (1.5 / 3.14159)) - v56) + 0.0) / (v56 * v56))) + ((1e10 + 1.5) * v56)) + v56)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v56
    _ = v58
    _ = v6
}
