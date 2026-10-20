package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint; F1 int; F2 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 int8 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 127
    v3 := []string{""}
    var v4 uint16 = 42
    var v5 int = 980
    var v6 uint = 0
    v5 = (v5 - v5)
    if true {
        if false {
            var v7 int64 = -9223372036854775808
            for fi8 := 0; fi8 < 3; fi8++ {
                v3 = append(v3, (("a" + "foo bar") + (("foo bar" + ("a" + "foo bar")) + "hello")))
                _ = fi8
            }
            v9 := GenS1{F0: v4, F1: 1000}
            _ = v7
            _ = v9
        }
    }
    var v10 int32 = 2147483647
    if true {
        if true {
            v5 -= v5
            v11 := GenS2{F0: -834, F1: (v4 - 255)}
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint8 = 1
                _ = fi13
                _ = v14
            }
            var v15 float32 = ((0.0 / float32((-2.25 / 1e10))) * 3.14159)
            _ = fi12
            _ = v15
        }
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v4, v2)
    case "":
        fmt.Println(v5, v6, v2)
    case "\n\t":
        _ = len(v3)
    }
    var v16 int32 = v10
    switch "foo bar" {
    case "a":
        v16 += v16
    }
    if false {
        if ((v6 >> ((v6) & 63)) != (v6 >> ((v6) & 63))) {
            switch ("a" + ("\n\t" + "\n\t")) {
            case "\n\t":
                v17 := GenS1{F0: 0, F1: 1000}
                _ = v17
            default:
                var v18 string = "foo bar"
                _ = v18
            }
            fmt.Println(v4, v16, v6)
            v2 = v2
        }
        if (-1 >= 42) {
            v10 += v10
            switch "a" {
            case "foo bar":
                if (!((v16 >> ((v6) & 63)) >= (-(v16)))) {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        var v20 int32 = v10
                        var v21 int8 = 42
                        v3 = append(v3, "hello")
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    switch v5 {
                    case 127:
                        fmt.Println(v2, v2)
                    case 4294967295:
                        if false {
                            var v22 int64 = 580
                            v3 = append(v3, "foo bar")
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 14; fi23++ {
                                v4 = 32767
                                _ = fi23
                            }
                            if (v6 <= 0) {
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                v1 -= v1
                                v24 := GenS3{F0: v6, F1: (v5 | v5), F2: -9223372036854775808}
                                _ = v24
                            } else {
                                for fi25 := 0; fi25 < 20; fi25++ {
                                    for fi26 := 0; fi26 < 12; fi26++ {
                                        v6 = v6
                                        _ = fi26
                                    }
                                    v27 := GenS4{F0: 2}
                                    for fi28 := 0; fi28 < 1; fi28++ {
                                        var v29 float64 = float64(-2.25)
                                        if len(v3) > 0 {
                                            v3[0] = (("foo bar" + "") + "a")
                                        }
                                        var v30 bool = true
                                        _ = fi28
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = fi25
                                    _ = v27
                                }
                                for fi31 := 0; fi31 < 15; fi31++ {
                                    if false {
                                        var v32 uint64 = 0
                                        _ = v32
                                    } else {
                                        switch "a" {
                                        case "":
                                            v33 := GenS1{F0: (v4 * v4), F1: 0}
                                            _ = v33
                                        default:
                                            v10 += (v10 & 2147483647)
                                        }
                                        v34 := GenS5{F0: 0}
                                        if len(v3) > 0 {
                                            v3[0] = "hello"
                                        }
                                        _ = v34
                                    }
                                    for fi35 := 0; fi35 < 6; fi35++ {
                                        _ = len(v3)
                                        var v36 int = -512
                                        fmt.Println(v16, v6, fi35)
                                        _ = fi35
                                        _ = v36
                                    }
                                    _ = fi31
                                }
                                fmt.Println(v16, v2, v4)
                            }
                            v4 = (^(v4))
                        }
                    }
                }
            case "a":
                var v37 int = (v5 << (((v6 & (v6 - (v6 ^ 781)))) & 63))
                _ = v37
            default:
                for fi38 := 0; fi38 < 19; fi38++ {
                    var v39 uint16 = (v4 >> (((v6 | 55)) & 63))
                    v39 = v4
                    v1 = float64(float32((v1 / 3.14159)))
                    _ = fi38
                    _ = v39
                }
            }
        } else {
            if len(v3) > 0 {
                v3[0] = ""
            }
        }
        if (false || ((v6 << ((2) & 63)) != 1)) {
            for fi40 := 0; fi40 < 2; fi40++ {
                for fi41 := 0; fi41 < 4; fi41++ {
                    switch int(v2) {
                    case -1:
                        if (70 == 100) {
                            v42 := GenS2{F0: 0, F1: (v4 % ((2) | 1))}
                            switch "" {
                            case "\n\t":
                                fmt.Println(v1, v2, fi40)
                            }
                            v4 = v4
                            _ = v42
                        } else {
                            if true {
                                v1 = 1.5
                            }
                        }
                    case 0:
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            v10 = v10
                            v1 -= v1
                            v43 := GenS6{F0: fi40}
                            _ = v43
                        }
                    case 127:
                        fmt.Println(v16, v2, v10)
                    default:
                        v5 -= (^(fi40))
                    }
                    _ = fi41
                }
                fmt.Println(v5, v10, v16)
                _ = fi40
            }
            var v44 uint8 = 127
            _ = v44
        } else {
            v3 = append(v3, ("" + "\n\t"))
        }
    }
    var v45 float32 = float32(float64(float32(((0.1 * 0.1) - 1e10))))
    var v46 int16 = 32767
    for fi47 := 0; fi47 < 2; fi47++ {
        v1 = v1
        _ = fi47
    }
    var v48 uint32 = 2147483648
    var v49 uint16 = (v4 - (v4 + 255))
    v50 := GenS7{F0: v4, F1: -325}
    var v51 int = v5
    if true {
        if true {
            v49 = v49
            for fi52 := 0; fi52 < 19; fi52++ {
                var v53 int16 = 1000
                _ = fi52
                _ = v53
            }
        } else {
            var v54 bool = true
            v4 = v4
            switch (v51 << ((18446744073709551615) & 63)) {
            case -51:
                var v55 uint16 = (v49 & uint16(v2))
                _ = v55
            default:
                if v54 {
                    var v56 uint32 = v48
                    if (!(v48 == v56)) {
                        fmt.Println(v54)
                        v46 = (v46 + (v46 + (v46 << ((v6) & 63))))
                    }
                    _ = v56
                } else {
                    if ((v45 + v45) < float32(1e10)) {
                        switch ("\n\t" + ((("hello" + "") + (((("\n\t" + "") + "hello") + (("\n\t" + "hello") + "")) + "\n\t")) + "\n\t")) {
                        case "\n\t":
                            switch (v5 * v5) {
                            case -1:
                                var v57 uint8 = 255
                                _ = v57
                            }
                        default:
                            if (float32(v1) < ((0.1 + v45) + ((float32(((v1 - 0.1) * (float64((v45 - v45)) * (v1 - v1)))) + ((v45 + -2.25) - (((v45 * 0.1) - -2.25) + (((v45 * (((((v45 * (v45 + (v45 * 1e10))) - (v45 / (v45 / v45))) / 1e10) / 1e10) / (float32(v1) + float32(((float64(float32(1.5)) + float64((0.0 / v45))) - (v1 / float64(v45))))))) * (((0.1 * ((0.0 + v45) * v45)) * v45) - 3.14159)) / v45)))) - 0.0))) {
                                var v58 int16 = -1
                                v48 += (v48 * v48)
                                for fi59 := 0; fi59 < 12; fi59++ {
                                    for fi60 := 0; fi60 < 14; fi60++ {
                                        fmt.Println(v46, v45)
                                        var v61 uint64 = 18446744073709551615
                                        v6 += v6
                                        var v62 uint64 = v61
                                        _ = fi60
                                        _ = v61
                                        _ = v62
                                    }
                                    if v54 {
                                        if false {
                                            switch -1 {
                                            case -9223372036854775808:
                                                v51 = v5
                                            }
                                        } else {
                                            for fi63 := 0; fi63 < 15; fi63++ {
                                                if (v2 <= (v2 - (v2 | (v2 / (((^(v2))) | 1))))) {
                                                    if ((v58 ^ v46) > 42) {
                                                        if (v54 && ((^(v16)) >= 0)) {
                                                            _ = len(v3)
                                                        } else {
                                                            if (!v54) {
                                                                var v64 float32 = 0.1
                                                                v65 := GenS8{F0: (v49 << (((^(v6))) & 63)), F1: 2147483647}
                                                                var v66 uint64 = 2
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            v51 = (fi59 - (v5 + 9223372036854775807))
                                                            var v67 uint16 = (^(v49))
                                                            _ = v67
                                                        }
                                                        fmt.Println(v48)
                                                    }
                                                }
                                                v1 = float64(v45)
                                                _ = fi63
                                            }
                                            for fi68 := 0; fi68 < 14; fi68++ {
                                                var v69 uint32 = (v48 ^ (v48 / ((20) | 1)))
                                                switch "hello" {
                                                case "a":
                                                    v70 := GenS6{F0: (v5 % (((^(v5))) | 1))}
                                                    _ = v70
                                                }
                                                _ = fi68
                                                _ = v69
                                            }
                                            v16 = (-(v10))
                                        }
                                    } else {
                                        v3 = append(v3, "")
                                    }
                                    if v54 {
                                        if ((v10 % ((1) | 1)) < (v16 | 602)) {
                                            fmt.Println(v45, v49, v16)
                                            var v71 uint64 = 2
                                            _ = v71
                                        }
                                        switch (-(fi59)) {
                                        case 42:
                                            var v72 uint32 = 2147483647
                                            _ = v72
                                        }
                                    }
                                    _ = fi59
                                }
                                _ = v58
                            }
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                    }
                    v73 := GenS9{F0: 127}
                    _ = v73
                }
            }
            _ = v54
        }
    }
    var v74 float32 = v45
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v74
}
