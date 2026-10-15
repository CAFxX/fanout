package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
type GenS3 struct { F0 int32; F1 uint16; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint; F1 uint32 }
type GenS9 struct { F0 int; F1 int16; F2 int }
type GenS10 struct { F0 uint32; F1 uint64; F2 uint32 }
func genFn0() int {
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint = 191
        v8 := GenS1{F0: -126, F1: true, F2: -1}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 1
    v10 := []int{4294967295, 1, 128, 4294967295}
    switch 127 {
    case -2147483648:
        v11 := GenS2{F0: 0, F1: 32767, F2: 116}
        _ = v11
    case -9223372036854775808:
        var v12 uint64 = (v9 / ((v9) | 1))
        _ = v12
    default:
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS3{F0: -1, F1: 124, F2: 0}
            v9 = v9
            fmt.Println(v9, fi13)
            _ = fi13
            _ = v14
        }
    }
    if (true && ("a" <= "a")) {
        if false {
            var v15 int64 = 592
            _ = v15
        }
        v16 := GenS3{F0: -291, F1: 2, F2: 59}
        _ = v16
    } else {
        if len(v10) > 0 {
            v10[0] = 128
        }
    }
    _ = v10
    _ = v9
    return 127
}
func genFn1(p0 int8, p1 int64) bool {
    for fi17 := 0; fi17 < 9; fi17++ {
        p0 = (p0 << ((975) & 63))
        p0 = (p0 + -72)
        _ = fi17
    }
    p1 = 42
    p1 = p1
    var v18 int32 = -2147483648
    var v19 uint16 = 32767
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return true
}

func main() {
    var v1 bool = (false || (((0.0 / (-2.25 / (1.5 / (1.5 + (((1.5 + 0.0) * (1e10 * (float64(0.0) / 3.14159))) / (3.14159 / 1.5)))))) / 0.1) < -2.25))
    v2 := []uint32{2147483647, 65535}
    var v3 uint16 = 0
    v4 := make(map[string]int)
    var v5 uint32 = 4294967295
    v20 := GenS2{F0: -2147483648, F1: 255, F2: -1}
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS2{F0: 1, F1: v3, F2: -128}
        _ = fi21
        _ = v22
    }
    if (("\n\t" + "a") > "a") {
        var v23 int = 42
        _ = v23
    } else {
        v24 := GenS3{F0: 1, F1: (^(v3)), F2: -59}
        _ = v24
    }
    var v25 uint16 = 0
    _ = len(v2)
    v25 = v3
    v2 = append(v2, v5)
    v26 := GenS2{F0: -1, F1: 255, F2: -28}
    if false {
        for fi27 := 0; fi27 < 10; fi27++ {
            switch ("" + "\n\t") {
            case "hello":
                _ = len(v2)
            case "\n\t":
                for fi28 := 0; fi28 < 4; fi28++ {
                    if v1 {
                        if v1 {
                            fmt.Println(fi28, fi27, fi28)
                        } else {
                            switch ("a" + "a") {
                            case "\n\t":
                                v3 = 966
                            case "foo bar":
                                if (42 > 42) {
                                    if (!(1.5 <= (1.5 * (1.5 - 3.14159)))) {
                                        v1 = ((1e10 >= 1e10) && v1)
                                    }
                                    v4[("" + "hello")] = fi28
                                    v25 += (v3 % ((v3) | 1))
                                } else {
                                    v2 = append(v2, v5)
                                }
                            default:
                                _ = len(v2)
                            }
                            switch int(v3) {
                            case 1:
                                switch (fi27 - (fi27 | (fi27 >> ((2) & 63)))) {
                                case 127:
                                    var v29 uint8 = 128
                                    _ = v29
                                case 1:
                                    if v1 {
                                        var v30 uint64 = 42
                                        v31 := GenS1{F0: int8(v3), F1: v1, F2: 1}
                                        _ = v30
                                        _ = v31
                                    } else {
                                        var v32 uint8 = 181
                                        _ = v32
                                    }
                                default:
                                    switch fi27 {
                                    case 0:
                                        v33 := GenS4{F0: 0}
                                        _ = v33
                                    case 42:
                                        fmt.Println(fi28, fi27, fi27)
                                    }
                                }
                            }
                            _ = len(v2)
                        }
                        delete(v4, "")
                        var v34 string = "hello"
                        _ = v34
                    }
                    v35 := GenS2{F0: -1, F1: (v3 * (v3 - (v3 << ((42) & 63)))), F2: 127}
                    var v36 string = "hello"
                    var v37 float64 = -2.25
                    var v38 int32 = -2147483648
                    _ = fi28
                    _ = v35
                    _ = v36
                    _ = v37
                    _ = v38
                }
            default:
                if (0.0 >= 1.5) {
                    v39 := GenS5{F0: -1}
                    _ = v39
                }
            }
            if v1 {
                v1 = true
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v25 = v3
                switch 9223372036854775807 {
                case 42:
                    if v1 {
                        v5 -= 2147483647
                    } else {
                        for fi40 := 0; fi40 < 18; fi40++ {
                            fmt.Println(fi27, fi40)
                            v3 -= 2
                            switch ("" + (("\n\t" + "hello") + ((("hello" + "foo bar") + ((("foo bar" + "foo bar") + "a") + ("foo bar" + ""))) + (("\n\t" + ("" + "foo bar")) + "a")))) {
                            case "":
                                _ = len(v4)
                            case "foo bar":
                                var v41 int = (-(fi27))
                                _ = v41
                            case "hello":
                                if (false && v1) {
                                    fmt.Println(v1, fi27)
                                    v25 += (v3 & (v25 + 32767))
                                    v42 := GenS3{F0: -634, F1: (v3 >> ((255) & 63)), F2: 42}
                                    var v43 int = genFn0()
                                    if (!true) {
                                        var v44 int16 = -1
                                        _ = v44
                                    } else {
                                        v43 = 32767
                                        v25 += 32767
                                    }
                                    _ = v42
                                    _ = v43
                                } else {
                                    switch (fi27 / (((^(fi27))) | 1)) {
                                    case -1:
                                        _ = len(v2)
                                    case -9223372036854775808:
                                        v45 := GenS5{F0: 1}
                                        _ = v45
                                    case -901:
                                        var v46 uint = 18446744073709551615
                                        _ = v46
                                    }
                                    switch ("" + "hello") {
                                    case "a":
                                        _ = len(v4)
                                        v5 -= v5
                                        var v47 uint32 = v5
                                        fmt.Println(fi27, v47, v47)
                                        _ = v47
                                    default:
                                        switch (fi27 >> ((4294967295) & 63)) {
                                        case 288:
                                            v5 += 156
                                        case -877:
                                            var v48 uint32 = 0
                                            _ = v48
                                        case 9223372036854775807:
                                            v3 += 2
                                        }
                                    }
                                }
                            default:
                                fmt.Println(v5, fi40, v3)
                                v49 := v4["a"]
                                _ = v49
                            }
                            var v50 int32 = 2147483647
                            _ = fi40
                            _ = v50
                        }
                        v51 := GenS6{F0: 0, F1: fi27}
                        _ = v51
                    }
                case 108:
                    v52 := GenS2{F0: 1, F1: 0, F2: 1}
                    _ = v52
                }
            }
            v53 := GenS7{F0: 128}
            _ = fi27
            _ = v53
        }
        v5 += (^(v5))
    }
    v4["\n\t"] = -1
    v54 := GenS8{F0: 506, F1: v5}
    for fi55 := 0; fi55 < 17; fi55++ {
        for fi56 := 0; fi56 < 2; fi56++ {
            v57 := GenS5{F0: -97}
            fmt.Println(v5, v3, fi55)
            _ = fi56
            _ = v57
        }
        switch fi55 {
        case -1:
            if (int16(v3) == 598) {
                v3 = (^(v3))
            } else {
                v58 := GenS9{F0: fi55, F1: -1, F2: 4294967295}
                switch "" {
                case "":
                    v4[(((("" + "hello") + "hello") + "foo bar") + ("\n\t" + (("a" + "") + "a")))] = fi55
                default:
                    for fi59 := 0; fi59 < 1; fi59++ {
                        fmt.Println(v25, v25)
                        fmt.Println(fi59, v3, fi59)
                        var v60 uint16 = (v25 + (^(v3)))
                        if v1 {
                            fmt.Println(v25)
                            if (1e10 > 3.14159) {
                                v61 := GenS1{F0: -93, F1: (!false), F2: 1}
                                fmt.Println(fi59)
                                _ = v61
                            }
                            v62 := v4[""]
                            _ = v62
                        } else {
                            fmt.Println(v5)
                            if (0 >= -527) {
                                var v63 int64 = -1
                                v1 = v1
                                delete(v4, ((("a" + ("\n\t" + "")) + "foo bar") + "a"))
                                _ = v63
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = v5
                        }
                        if false {
                            fmt.Println(fi55, fi59)
                            if (0.1 < 3.14159) {
                                _ = len(v2)
                                v4["a"] = (fi59 ^ (fi59 / ((-1) | 1)))
                                v3 -= 25
                            }
                            var v64 bool = v1
                            _ = v64
                        }
                        switch (^(fi59)) {
                        case -9223372036854775808:
                            fmt.Println(v60, fi55)
                        case 2147483647:
                            if (2 <= v5) {
                                v65 := GenS8{F0: 15, F1: (v5 & 281)}
                                _ = v65
                            } else {
                                fmt.Println(v60, fi55, v1)
                            }
                        case -1:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            if len(v2) > 0 {
                                v2[0] = v5
                            }
                        }
                        v66 := GenS10{F0: (v5 | (^(v5))), F1: 2, F2: v5}
                        v25 -= 1
                        _ = fi59
                        _ = v60
                        _ = v66
                    }
                }
                if true {
                    var v67 uint16 = 42
                    _ = v67
                } else {
                    for fi68 := 0; fi68 < 12; fi68++ {
                        if ((float32((3.14159 + 3.14159)) / float32(0.1)) != 0.0) {
                            fmt.Println(fi68, fi55, fi55)
                            v69 := v4["\n\t"]
                            _ = v69
                        }
                        v1 = (v1 || (0.1 == -2.25))
                        _ = fi68
                    }
                }
                _ = v58
            }
        case -436:
            fmt.Println(v25)
        }
        for fi70 := 0; fi70 < 20; fi70++ {
            var v71 uint64 = 2
            _ = fi70
            _ = v71
        }
        _ = fi55
    }
    v1 = (!v1)
    if v1 {
        _ = len(v2)
        delete(v4, ("hello" + ((("\n\t" + ("foo bar" + "\n\t")) + ("" + "hello")) + "hello")))
    } else {
        v1 = (v1 && (1 < 42))
        fmt.Println(v5)
        v2 = append(v2, (v5 & 326))
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v54
}
