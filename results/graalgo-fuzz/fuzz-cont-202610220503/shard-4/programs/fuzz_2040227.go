package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint32; F1 string; F2 uint8 }
type GenS6 struct { F0 int32; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint, p2 string) int8 {
    var v7 uint64 = 2
    v8 := GenS1{F0: 0}
    p2 = "foo bar"
    v9 := []int{9223372036854775807, -1}
    switch 0 {
    case -1000:
        v10 := GenS1{F0: 2147483648}
        _ = v10
    case 161:
        if false {
            p1 = 255
        }
    case -2147483648:
        var v11 int = -9223372036854775808
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -53
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]int)
    var v3 float64 = 0.1
    var v4 bool = (1e10 == 1e10)
    var v5 uint32 = 2147483647
    var v6 uint16 = uint16(v1)
    var v12 uint8 = v1
    v13 := []int{2147483647, 127}
    if ((127 == 128) && true) {
        if v4 {
            if v4 {
                var v14 int = int(v6)
                fmt.Println(v5, v4, v12)
                v15 := GenS1{F0: v5}
                _ = v14
                _ = v15
            } else {
                var v16 bool = v4
                _ = v16
            }
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 int64 = -1
                var v19 float32 = ((-2.25 / 1e10) / -2.25)
                v20 := GenS1{F0: 0}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                for fi22 := 0; fi22 < 17; fi22++ {
                    v23 := GenS2{F0: (v5 / (((v5 & (v5 - 2147483647))) | 1))}
                    _ = fi22
                    _ = v23
                }
                switch "a" {
                case "":
                    var v24 float32 = 0.0
                    _ = v24
                }
                v25 := GenS3{F0: genFn0(42, 18446744073709551615, "foo bar"), F1: (^(v1))}
                _ = fi21
                _ = v25
            }
        } else {
            switch "hello" {
            case "foo bar":
                v4 = false
            case "":
                switch 1 {
                case 32768:
                    var v26 int = -276
                    _ = v26
                case 4294967295:
                    v2[(v5 ^ v5)] = 2147483647
                case 42:
                    v4 = (!(1 >= 2))
                }
            default:
                if (1 > 208) {
                    var v27 int16 = int16(v12)
                    for fi28 := 0; fi28 < 12; fi28++ {
                        v6 = (v6 - (^(v6)))
                        _ = fi28
                    }
                    _ = v27
                }
            }
        }
    } else {
        var v29 float64 = float64(1.5)
        _ = v29
    }
    switch -1 {
    case 128:
        switch (("\n\t" + ("a" + "\n\t")) + "a") {
        case "hello":
            v5 = v5
        case "foo bar":
            v30 := GenS3{F0: 127, F1: (v1 >> ((18446744073709551615) & 63))}
            _ = v30
        default:
            if (v4 || v4) {
                v1 += v1
                v6 += v6
            } else {
                var v31 int32 = 860
                switch 42 {
                case 0:
                    if (((-2.25 * float32((float64(float32((((v3 - (-2.25 - (-2.25 - v3))) - v3) / (3.14159 / -2.25)))) * (((-2.25 + (((1e10 + v3) * v3) * v3)) + float64((0.1 - 1e10))) / 0.1)))) * (0.0 * 0.0)) == -2.25) {
                        fmt.Println(v4)
                        var v32 uint32 = v5
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            v31 += (^(v31))
                            switch (fi33 * fi33) {
                            case 994:
                                v6 = 255
                            }
                            var v34 int8 = -128
                            _ = fi33
                            _ = v34
                        }
                        _ = len(v13)
                        for fi35 := 0; fi35 < 12; fi35++ {
                            if ("hello" > ("" + ("\n\t" + "foo bar"))) {
                                var v36 string = "foo bar"
                                var v37 uint64 = 42
                                _ = v36
                                _ = v37
                            } else {
                                v31 += v31
                                fmt.Println(v31, v3, fi35)
                                v38 := GenS1{F0: v5}
                                _ = v38
                            }
                            _ = fi35
                        }
                    }
                case 32768:
                    v3 = v3
                case 429:
                    switch "" {
                    case "hello":
                        fmt.Println(v31, v31, v31)
                    case "foo bar":
                        v39 := GenS4{F0: 42, F1: genFn0(18446744073709551615, 255, "")}
                        _ = v39
                    }
                default:
                    fmt.Println(v6, v1, v5)
                }
                _ = v31
            }
        }
    case 2147483647:
        var v40 int16 = 457
        _ = v40
    case 4294967295:
        for fi41 := 0; fi41 < 3; fi41++ {
            fmt.Println(v4, v12, v4)
            _ = fi41
        }
    default:
        fmt.Println(v5)
    }
    if (-901 > 2147483647) {
        v12 += v1
    } else {
        v2[v5] = -9223372036854775808
        switch 42 {
        case -248:
            if v4 {
                v4 = ((false || (!v4)) || ((v5 << ((42) & 63)) != v5))
                fmt.Println(v4, v12, v3)
                for fi42 := 0; fi42 < 1; fi42++ {
                    if (v4 || (-128 < 127)) {
                        if (!(25 == v5)) {
                            var v43 bool = (v4 && v4)
                            if (!v43) {
                                v5 = v5
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                if (!(!((!v4) && true))) {
                                    for fi44 := 0; fi44 < 2; fi44++ {
                                        switch "" {
                                        case "hello":
                                            if (-530 >= -417) {
                                                v2[v5] = 4294967295
                                                if ((!true) && ("hello" == "hello")) {
                                                    var v45 int = (fi42 * fi42)
                                                    _ = v45
                                                }
                                            } else {
                                                var v46 uint64 = 18446744073709551615
                                                _ = v46
                                            }
                                        case "a":
                                            var v47 int = 1
                                            _ = v47
                                        case "foo bar":
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                        }
                                        _ = fi44
                                    }
                                    v48 := GenS3{F0: -14, F1: v12}
                                    _ = v48
                                }
                            }
                            var v49 int32 = -227
                            _ = v43
                            _ = v49
                        } else {
                            v50 := GenS5{F0: (v5 ^ 2147483647), F1: (("foo bar" + "a") + ("foo bar" + "hello")), F2: (v1 + uint8(v5))}
                            _ = v50
                        }
                        fmt.Println(v4)
                    }
                    v6 += (v6 * v6)
                    _ = fi42
                }
            }
        default:
            if (0.0 == 0.0) {
                fmt.Println(v4, v5, v5)
                var v51 uint16 = v6
                _ = v51
            }
        }
    }
    for fi52 := 0; fi52 < 14; fi52++ {
        v12 = uint8(v5)
        for fi53 := 0; fi53 < 12; fi53++ {
            fmt.Println(v5, v3, v4)
            if (128 <= v12) {
                var v54 float64 = v3
                _ = v54
            }
            var v55 uint64 = 2
            _ = fi53
            _ = v55
        }
        _ = fi52
    }
    var v56 float32 = 1.5
    if len(v13) > 0 {
        v13[0] = -984
    }
    for fi57 := 0; fi57 < 20; fi57++ {
        v56 = v56
        if v4 {
            var v58 string = "foo bar"
            switch (-(fi57)) {
            case -2147483648:
                fmt.Println(v1)
            case 32767:
                var v59 int64 = -1
                _ = v59
            default:
                v60 := GenS6{F0: 42, F1: 317, F2: v12}
                _ = v60
            }
            _ = v58
        } else {
            for fi61 := 0; fi61 < 9; fi61++ {
                fmt.Println(v1)
                var v62 int32 = -2147483648
                var v63 uint16 = (v6 | uint16(v1))
                _ = fi61
                _ = v62
                _ = v63
            }
        }
        _ = fi57
    }
    var v64 float64 = v3
    var v65 uint = 414
    var v66 uint64 = 23
    if v4 {
        v1 = (v12 / (((^(v1))) | 1))
        if (v4 || false) {
            var v67 uint32 = 2147483648
            v56 -= 0.1
            v68 := GenS7{F0: (^(v66))}
            _ = v67
            _ = v68
        }
    } else {
        v6 -= v6
        for fi69 := 0; fi69 < 20; fi69++ {
            var v70 uint32 = (^(v5))
            fmt.Println(v65, v64)
            if (42 <= -128) {
                fmt.Println(v66, v4, v4)
                fmt.Println(v12, v12, v12)
            } else {
                v12 = 74
            }
            _ = fi69
            _ = v70
        }
        v71 := v2[v5]
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v64
    _ = v65
    _ = v66
}
