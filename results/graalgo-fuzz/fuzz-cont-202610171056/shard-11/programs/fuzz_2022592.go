package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool; F1 int64; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint; F2 string }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 int; F2 int32 }
type GenS9 struct { F0 int16; F1 string; F2 uint16 }
type GenS10 struct { F0 int64; F1 uint64; F2 uint }
func genFn0(p0 uint8) int32 {
    if ((0.1 + 3.14159) < 1.5) {
        var v4 float32 = (3.14159 * 3.14159)
        var v5 uint32 = 2147483648
        p0 = 127
        _ = v4
        _ = v5
    } else {
        v6 := GenS1{F0: 154, F1: 2147483648}
        if (true && true) {
            if false {
                v7 := GenS2{F0: true, F1: 9223372036854775807, F2: p0}
                p0 = (^(p0))
                v8 := []int{1, 32767, 127}
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: 2, F1: 2147483648}
            var v10 int32 = 229
            _ = v10
            _ = v9
        } else {
            v11 := []int{32767, 2147483647, -414, -9223372036854775808}
            _ = v11
        }
        if (1 < 1) {
            var v12 uint16 = 42
            _ = v12
        } else {
            p0 -= p0
            var v13 uint64 = 489
            _ = v13
        }
        _ = v6
    }
    p0 = p0
    v14 := GenS3{F0: 1, F1: 921, F2: ("\n\t" + ("foo bar" + "\n\t"))}
    _ = p0
    _ = v14
    return 2147483647
}

func main() {
    var v1 string = "a"
    v2 := []int{}
    var v3 uint = 18446744073709551615
    if (true && ((v3 | v3) < 255)) {
        _ = len(v2)
        for fi15 := 0; fi15 < 20; fi15++ {
            if ((3.14159 / 1e10) >= 1.5) {
                fmt.Println(v1, v3, v1)
            } else {
                if true {
                    fmt.Println(v3, v1, v3)
                    v1 = v1
                } else {
                    var v16 uint64 = 2
                    _ = v16
                }
                v17 := GenS1{F0: 0, F1: 0}
                for fi18 := 0; fi18 < 9; fi18++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi18
                }
                _ = v17
            }
            var v19 float32 = (-2.25 + ((1.5 + -2.25) / 1.0))
            v19 = v19
            _ = fi15
            _ = v19
        }
    } else {
        v3 = 857
    }
    var v20 uint8 = 0
    if false {
        if (!(0.0 <= ((((0.1 / 1e10) * 1.5) / ((-2.25 - ((1.5 - ((((3.14159 - (0.1 - float64(0.1))) - 0.0) - 1e10) * float64((3.14159 - 0.1)))) + (0.0 + 3.14159))) + -2.25)) * 1.5))) {
            if false {
                var v21 int8 = 19
                switch "" {
                case "a":
                    var v22 uint64 = 18446744073709551615
                    _ = v22
                case "foo bar":
                    fmt.Println(v20)
                }
                for fi23 := 0; fi23 < 7; fi23++ {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        var v25 int8 = (v21 >> (((v3 + v3)) & 63))
                        var v26 uint32 = uint32(v3)
                        v21 = v25
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS2{F0: true, F1: -586, F2: 1}
                    _ = fi23
                    _ = v27
                }
                _ = v21
            }
        }
        v1 = v1
    }
    v2 = append(v2, 4294967295)
    var v28 int8 = 127
    var v29 bool = ((0.0 / 3.14159) == -2.25)
    for fi30 := 0; fi30 < 14; fi30++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v31 := GenS4{F0: 652, F1: genFn0((v20 << ((v3) & 63)))}
        fmt.Println(v28)
        _ = fi30
        _ = v31
    }
    var v32 float64 = 3.14159
    v33 := GenS1{F0: 2, F1: 4294967295}
    var v34 float32 = -2.25
    for fi35 := 0; fi35 < 4; fi35++ {
        switch (fi35 + fi35) {
        case 32768:
            if (true && (42 != 220)) {
                if ((false || ((v20 ^ v20) != v20)) && (!(!false))) {
                    if v29 {
                        if (32767 != -32768) {
                            fmt.Println(fi35)
                            switch fi35 {
                            case -1:
                                v3 = (v3 * 132)
                            }
                        }
                    }
                    if false {
                        v36 := GenS5{F0: v3, F1: 23}
                        for fi37 := 0; fi37 < 7; fi37++ {
                            v1 = (v1 + (v1 + ""))
                            _ = fi37
                        }
                        _ = v36
                    }
                }
                var v38 int = (fi35 << ((v3) & 63))
                _ = v38
            } else {
                v39 := GenS6{F0: 909}
                if false {
                    for fi40 := 0; fi40 < 17; fi40++ {
                        var v41 uint = 693
                        v2 = append(v2, fi40)
                        fmt.Println(v1, v3)
                        _ = fi40
                        _ = v41
                    }
                    v42 := GenS7{F0: (!(!(226 <= 42)))}
                    _ = v42
                }
                fmt.Println(v1)
                _ = v39
            }
        case 4294967295:
            switch v1 {
            case "a":
                if v29 {
                    var v43 float64 = 1e10
                    v28 = v28
                    switch 0 {
                    case 338:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = v43
                } else {
                    switch v1 {
                    case "foo bar":
                        v44 := GenS8{F0: (v3 * 42), F1: fi35, F2: 0}
                        _ = v44
                    case "hello":
                        fmt.Println(v20, v34)
                    default:
                        if len(v2) > 0 {
                            v2[0] = 773
                        }
                    }
                    if (3.14159 > (((0.1 - v34) * 0.0) * 1.5)) {
                        v3 -= v3
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v28, v28)
                        }
                        switch "a" {
                        case "a":
                            switch "foo bar" {
                            case "\n\t":
                                for fi45 := 0; fi45 < 13; fi45++ {
                                    switch ("hello" + v1) {
                                    case "foo bar":
                                        fmt.Println(v28)
                                    }
                                    _ = fi45
                                }
                            }
                        case "hello":
                            v3 += v3
                        default:
                            var v46 uint16 = 32767
                            _ = v46
                        }
                        if (true && (!(578 <= 2))) {
                            if v29 {
                                for fi47 := 0; fi47 < 10; fi47++ {
                                    var v48 uint16 = 42
                                    var v49 float64 = 1.5
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            }
                        } else {
                            if false {
                                for fi50 := 0; fi50 < 20; fi50++ {
                                    if (v29 || v29) {
                                        v51 := GenS7{F0: (2 == 178)}
                                        switch (^(fi35)) {
                                        case -722:
                                            v20 = 255
                                        case 32768:
                                            v52 := GenS4{F0: 1, F1: 0}
                                            _ = v52
                                        case 127:
                                            var v53 int = fi35
                                            _ = v53
                                        default:
                                            _ = len(v2)
                                        }
                                        _ = v51
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    switch fi50 {
                                    case -1:
                                        v32 = v32
                                    }
                                    var v54 int64 = 0
                                    var v55 string = v1
                                    switch (fi50 % ((fi50) | 1)) {
                                    case -19:
                                        switch v1 {
                                        case "\n\t":
                                            if (fi50 >= -712) {
                                                if len(v2) > 0 {
                                                    v2[0] = fi50
                                                }
                                                switch 2147483647 {
                                                case 151:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case 0:
                                                    switch v1 {
                                                    case "foo bar":
                                                        fmt.Println(fi35)
                                                    case "hello":
                                                        v34 += float32((v32 / -2.25))
                                                    }
                                                default:
                                                    v20 += v20
                                                }
                                                v32 -= (0.0 / v32)
                                                v56 := GenS6{F0: 894}
                                                _ = v56
                                            } else {
                                                var v57 uint16 = 42
                                                v1 = ""
                                                fmt.Println(v28, v55, v3)
                                                _ = v57
                                            }
                                        case "":
                                            var v58 uint32 = 2147483647
                                            _ = v58
                                        case "a":
                                            var v59 uint16 = 0
                                            _ = v59
                                        }
                                    case 947:
                                        fmt.Println(v34, v20, v32)
                                    }
                                    _ = fi50
                                    _ = v54
                                    _ = v55
                                }
                            }
                            var v60 int32 = genFn0(v20)
                            switch (^(fi35)) {
                            case 0:
                                v60 = genFn0(v20)
                            case -1:
                                v32 += (v32 * float64((((-2.25 + 1e10) + (v34 + 3.14159)) - (((v34 - float32((float64(v34) - (((((-2.25 - 1.5) - v32) * v32) * 3.14159) + -2.25)))) - 3.14159) - v34))))
                            case 32767:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                fmt.Println(v20)
                            }
                            _ = v60
                        }
                    }
                    var v61 int32 = -1
                    _ = v61
                }
            }
        case 676:
            v62 := GenS2{F0: true, F1: 872, F2: 128}
            _ = v62
        }
        var v63 int = fi35
        var v64 int32 = 505
        _ = fi35
        _ = v63
        _ = v64
    }
    v65 := GenS9{F0: -1, F1: v1, F2: 1}
    var v66 int16 = 185
    if (42 <= -1) {
        var v67 uint64 = 0
        var v68 uint64 = 960
        _ = v67
        _ = v68
    } else {
        var v69 float64 = -2.25
        _ = v69
    }
    if false {
        v3 = v3
        switch (v1 + (((((v1 + "\n\t") + (v1 + (v1 + "hello"))) + "foo bar") + v1) + "hello")) {
        case "a":
            var v70 uint32 = 1
            _ = v70
        case "foo bar":
            if v29 {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v71 int32 = 553
                fmt.Println(v28)
                _ = v71
            }
        }
        if (128 > 389) {
            v72 := GenS10{F0: 663, F1: 1, F2: 18446744073709551615}
            v34 -= 1.5
            var v73 uint64 = 907
            _ = v72
            _ = v73
        }
    } else {
        v74 := GenS10{F0: -3, F1: 0, F2: 0}
        _ = v74
    }
    var v75 uint16 = 554
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v66:", v66)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v65
    _ = v66
    _ = v75
}
