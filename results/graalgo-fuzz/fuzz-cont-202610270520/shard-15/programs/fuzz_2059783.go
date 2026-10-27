package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int64 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 uint16 }
func genFn0(p0 int64, p1 int8) int8 {
    v5 := GenS1{F0: 255, F1: (^(p0)), F2: p0}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v5
    return (p1 << ((1) & 63))
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    v3 := []uint32{}
    v4 := []string{}
    v2 = (!(true && (787 < 1)))
    switch ("a" + "") {
    case "a":
        for fi6 := 0; fi6 < 15; fi6++ {
            for fi7 := 0; fi7 < 3; fi7++ {
                v2 = (("foo bar" <= "a") || v2)
                if v2 {
                    var v8 string = ("foo bar" + (("foo bar" + "a") + ("a" + ("" + (("foo bar" + "hello") + "foo bar")))))
                    _ = v8
                } else {
                    var v9 int64 = -1
                    v10 := GenS1{F0: 1, F1: -9223372036854775808, F2: (v9 << ((463) & 63))}
                    v11 := GenS2{F0: 145, F1: false}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                v12 := GenS3{F0: "a"}
                _ = fi7
                _ = v12
            }
            switch (fi6 >> ((1) & 63)) {
            case 42:
                v13 := GenS4{F0: 0}
                _ = v13
            }
            _ = fi6
        }
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        switch ("" + "") {
        case "foo bar":
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 int = fi14
                _ = fi15
                _ = v16
            }
        default:
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 string = "foo bar"
                if (4294967295 == 2147483647) {
                    fmt.Println(v2)
                } else {
                    if v2 {
                        v19 := GenS2{F0: 23, F1: v2}
                        _ = v19
                    }
                    if (2 != 370) {
                        fmt.Println(v18)
                    }
                    switch v18 {
                    case "":
                        v18 = v18
                    case "foo bar":
                        if (65535 <= 4294967295) {
                            fmt.Println(v2, fi17)
                            v20 := GenS2{F0: 42, F1: v2}
                            _ = v20
                        } else {
                            var v21 int32 = 385
                            var v22 uint16 = 2
                            _ = v21
                            _ = v22
                        }
                    case "\n\t":
                        v23 := GenS5{F0: -1}
                        v24 := GenS6{F0: (v18 + (("\n\t" + ((("hello" + v18) + ("foo bar" + "hello")) + v18)) + (v18 + ""))), F1: 2}
                        _ = v23
                        _ = v24
                    default:
                        var v25 int = fi17
                        _ = v25
                    }
                }
                _ = fi17
                _ = v18
            }
        }
        var v26 int16 = 32767
        _ = len(v3)
        _ = fi14
        _ = v26
    }
    var v27 uint32 = 42
    switch ("foo bar" + "hello") {
    case "foo bar":
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 int8 = genFn0(1, 0)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi28
            _ = v29
        }
    case "\n\t":
        v30 := GenS7{F0: 32768, F1: 1}
        _ = v30
    default:
        v31 := GenS3{F0: "foo bar"}
        _ = v31
    }
    if true {
        var v32 string = "a"
        var v33 uint = 65535
        var v34 int16 = 42
        _ = v32
        _ = v33
        _ = v34
    } else {
        for fi35 := 0; fi35 < 18; fi35++ {
            var v36 int8 = -1
            _ = fi35
            _ = v36
        }
        var v37 float64 = -2.25
        _ = v37
    }
    if ("" > "foo bar") {
        var v38 int8 = 1
        _ = v38
    } else {
        if ((^(v27)) == 1) {
            if ("" < "a") {
                v4 = append(v4, "\n\t")
                for fi39 := 0; fi39 < 11; fi39++ {
                    _ = len(v3)
                    var v40 float64 = float64(0.0)
                    _ = fi39
                    _ = v40
                }
                if (0 < 0) {
                    for fi41 := 0; fi41 < 19; fi41++ {
                        v1["\n\t"] = "\n\t"
                        switch "" {
                        case "\n\t":
                            var v42 uint8 = 2
                            _ = v42
                        case "":
                            v43 := GenS4{F0: 4294967295}
                            _ = v43
                        }
                        _ = fi41
                    }
                } else {
                    for fi44 := 0; fi44 < 18; fi44++ {
                        var v45 uint64 = 514
                        _ = fi44
                        _ = v45
                    }
                    var v46 int64 = -215
                    _ = v46
                }
            } else {
                switch "a" {
                case "foo bar":
                    if (439 > 1000) {
                        if ((false || (1 == 4294967295)) && v2) {
                            v47 := GenS8{F0: 6}
                            var v48 uint16 = 2
                            _ = v47
                            _ = v48
                        }
                        v27 = (v27 / (((v27 * (v27 ^ (v27 ^ 929)))) | 1))
                        v2 = (!((((-2.25 / -2.25) / 1e10) - (1e10 - 1e10)) != (((0.0 * 3.14159) + 1.5) / (float32((-2.25 + 0.0)) + 1e10))))
                    }
                default:
                    v27 += (v27 << ((4294967295) & 63))
                }
                v49 := v1[("a" + "foo bar")]
                if ("a" >= v49) {
                    v3 = append(v3, v27)
                }
                _ = v49
            }
            var v50 int16 = 285
            if (!(true || v2)) {
                if (!(uint64(v27) >= 0)) {
                    if true {
                        fmt.Println(v50, v2)
                        v50 -= v50
                        v51 := GenS6{F0: ("" + "\n\t"), F1: 0}
                        _ = v51
                    }
                }
                for fi52 := 0; fi52 < 4; fi52++ {
                    if v2 {
                        v50 = int16(v27)
                    } else {
                        _ = len(v3)
                        var v53 int64 = -235
                        v54 := GenS9{F0: fi52, F1: 67}
                        _ = v53
                        _ = v54
                    }
                    v2 = v2
                    if false {
                        fmt.Println(v50, v27, fi52)
                    } else {
                        v27 = 2
                    }
                    _ = fi52
                }
                if (-1 <= 379) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v55 int16 = (v50 + -1)
                    _ = v55
                }
            }
            _ = v50
        }
        v56 := GenS10{F0: 905}
        switch -1 {
        case 0:
            v2 = ((0.1 - float32(float64(-2.25))) == 0.1)
        case 32767:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 191:
            v4 = append(v4, "foo bar")
        }
        _ = v56
    }
    v27 += v27
    var v57 uint64 = 42
    switch "foo bar" {
    case "":
        v2 = v2
    case "hello":
        v58 := v1["foo bar"]
        _ = v58
    default:
        if (!true) {
            for fi59 := 0; fi59 < 19; fi59++ {
                switch "hello" {
                case "":
                    switch -1 {
                    case 197:
                        for fi60 := 0; fi60 < 2; fi60++ {
                            switch (-(fi60)) {
                            case -249:
                                var v61 uint = 269
                                _ = v61
                            case -2147483648:
                                var v62 int8 = 0
                                _ = v62
                            case 128:
                                if true {
                                    v57 += (v57 << ((2) & 63))
                                    v2 = ((1e10 + 1.5) >= 1e10)
                                    fmt.Println(fi59, v2)
                                }
                            }
                            _ = fi60
                        }
                    }
                default:
                    v63 := GenS7{F0: fi59, F1: (^(v27))}
                    _ = v63
                }
                v3 = append(v3, (v27 | v27))
                _ = fi59
            }
        }
    }
    for fi64 := 0; fi64 < 8; fi64++ {
        var v65 int16 = 264
        for fi66 := 0; fi66 < 16; fi66++ {
            var v67 string = ("foo bar" + "foo bar")
            switch (fi64 & (fi66 * 42)) {
            case -1:
                var v68 int8 = -1
                _ = v68
            case -89:
                v69 := GenS2{F0: 42, F1: (true || false)}
                _ = v69
            }
            if false {
                if len(v3) > 0 {
                    v3[0] = 879
                }
                var v70 int16 = v65
                var v71 int8 = genFn0(int64(v57), 127)
                _ = v70
                _ = v71
            }
            var v72 int16 = (v65 % ((v65) | 1))
            var v73 uint32 = 145
            _ = fi66
            _ = v67
            _ = v72
            _ = v73
        }
        _ = fi64
        _ = v65
    }
    v27 -= uint32(v57)
    fmt.Println(v27)
    v74 := GenS1{F0: uint8(v27), F1: 0, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v57
    _ = v74
}
