package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 int8 }
func genFn0(p0 uint, p1 uint32, p2 int) uint32 {
    v6 := []int{}
    v6 = append(v6, p2)
    p0 = 580
    p0 -= p0
    for fi7 := 0; fi7 < 7; fi7++ {
        v6 = append(v6, p2)
        if (!(1 != 349)) {
            p0 = 255
            p2 -= -768
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                _ = len(v6)
                fmt.Println(p2, fi7, fi7)
                switch "hello" {
                case "\n\t":
                    switch (fi7 & -2147483648) {
                    case 128:
                        if (!false) {
                            v9 := GenS1{F0: "", F1: (^(fi8))}
                            var v10 uint32 = 4294967295
                            p1 += v10
                            fmt.Println(v10)
                            var v11 int8 = 42
                            _ = v10
                            _ = v11
                            _ = v9
                        } else {
                            if (323 < 601) {
                                var v12 float64 = (((0.1 * 1.5) / (1.5 / 3.14159)) * ((-2.25 * 3.14159) * 3.14159))
                                switch ("\n\t" + "a") {
                                case "a":
                                    var v13 int8 = 127
                                    _ = v13
                                default:
                                    fmt.Println(fi7, fi7, fi8)
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if (42 <= 790) {
                                    var v14 float32 = -2.25
                                    var v15 float64 = 1.5
                                    var v16 uint8 = 2
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        var v17 uint32 = 65535
                                        _ = v17
                                    default:
                                        var v18 int64 = 1
                                        _ = v18
                                    }
                                }
                                if false {
                                    fmt.Println(fi7, v12)
                                    switch ("\n\t" + "") {
                                    case "":
                                        v19 := GenS1{F0: "foo bar", F1: -2147483648}
                                        _ = v19
                                    default:
                                        v20 := GenS1{F0: "a", F1: -788}
                                        _ = v20
                                    }
                                    var v21 int = 1
                                    _ = v21
                                } else {
                                    switch fi7 {
                                    case 42:
                                        var v22 int = 4294967295
                                        _ = v22
                                    case 9223372036854775807:
                                        v23 := GenS2{F0: 439, F1: -1, F2: 1000}
                                        _ = v23
                                    }
                                }
                                _ = v12
                            } else {
                                p2 += (fi7 << ((p0) & 63))
                                v24 := GenS3{F0: 65535, F1: 1}
                                _ = v24
                            }
                        }
                    case -538:
                        v25 := GenS1{F0: "foo bar", F1: (p2 ^ (fi7 * (fi8 * fi7)))}
                        _ = v25
                    default:
                        if (!true) {
                            v26 := GenS1{F0: ("\n\t" + "\n\t"), F1: fi7}
                            switch p2 {
                            case 210:
                                switch (fi8 - (fi8 ^ (fi8 | p2))) {
                                case -2147483648:
                                    var v27 uint8 = 128
                                    _ = v27
                                }
                            case 32768:
                                var v28 int32 = 42
                                _ = v28
                            default:
                                p2 -= (fi7 | 2147483647)
                            }
                            _ = v26
                        }
                    }
                case "foo bar":
                    v29 := GenS1{F0: "\n\t", F1: 32767}
                    _ = v29
                case "a":
                    if true {
                        switch "hello" {
                        case "\n\t":
                            if (false && (165 <= 127)) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v30 string = ("foo bar" + ("foo bar" + ("foo bar" + "a")))
                                _ = v30
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            p0 += p0
                        case "":
                            var v31 float32 = ((0.0 * 0.1) / 1.0)
                            _ = v31
                        default:
                            _ = len(v6)
                        }
                        var v32 uint16 = 458
                        if (42 >= 65535) {
                            p0 += 1
                            v33 := GenS4{F0: p1, F1: 42}
                            var v34 int = fi8
                            _ = len(v6)
                            _ = v33
                            _ = v34
                        }
                        _ = v32
                    }
                }
                _ = fi8
            }
            p1 += (p1 & (p1 + p1))
        }
        switch -117 {
        case 9223372036854775807:
            if len(v6) > 0 {
                v6[0] = p2
            }
        case 32767:
            fmt.Println(p1, fi7)
        case 2147483647:
            p2 += (fi7 % (((p2 | -801)) | 1))
        default:
            fmt.Println(p2, p2)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    v1 := []string{"foo bar", (("a" + "") + "hello"), ("a" + (("hello" + "") + "a"))}
    var v2 uint16 = 946
    var v3 float32 = (0.1 / float32((1e10 - ((1e10 / (1.5 * float64((0.1 - 1e10)))) - (3.14159 / float64((-2.25 + -2.25)))))))
    v4 := make(map[int]bool)
    var v5 int = 551
    var v35 uint64 = 2
    if (0.0 > v3) {
        if true {
            fmt.Println(v3, v2)
            var v36 uint32 = 4294967295
            _ = v36
        }
        if true {
            v37 := GenS1{F0: "\n\t", F1: 227}
            if false {
                v2 = (v2 * v2)
            } else {
                v5 += (^(v5))
                v35 += (v35 / ((1) | 1))
                var v38 int32 = -1
                _ = v38
            }
            for fi39 := 0; fi39 < 4; fi39++ {
                v40 := GenS5{F0: "\n\t"}
                _ = fi39
                _ = v40
            }
            _ = v37
        }
        var v41 int = (v5 | v5)
        _ = v41
    } else {
        v5 = v5
        delete(v4, v5)
        if (-666 >= 1000) {
            for fi42 := 0; fi42 < 20; fi42++ {
                v43 := GenS4{F0: 1, F1: -1}
                _ = fi42
                _ = v43
            }
            _ = len(v1)
            fmt.Println(v5)
        }
    }
    var v44 uint32 = 489
    var v45 float32 = (-2.25 - float32(1.5))
    fmt.Println(v35)
    if (255 != 42) {
        fmt.Println(v5, v45, v3)
    }
    fmt.Println(v35, v5)
    for fi46 := 0; fi46 < 13; fi46++ {
        switch (("a" + "") + "\n\t") {
        case "a":
            var v47 uint32 = uint32(v2)
            _ = v47
        case "":
            v4[(-(fi46))] = true
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi46
    }
    switch int(v2) {
    case 4294967295:
        v2 += v2
    case 2147483647:
        var v48 float64 = 1.5
        _ = v48
    case -1:
        switch "hello" {
        case "\n\t":
            for fi49 := 0; fi49 < 11; fi49++ {
                v50 := v4[(v5 & -253)]
                _ = fi49
                _ = v50
            }
        }
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        var v52 uint64 = 2
        _ = fi51
        _ = v52
    }
    if false {
        var v53 string = "hello"
        _ = v53
    }
    for fi54 := 0; fi54 < 8; fi54++ {
        fmt.Println(v3)
        fmt.Println(v2, v3, v44)
        var v55 int8 = 100
        _ = fi54
        _ = v55
    }
    switch "\n\t" {
    case "\n\t":
        v1 = append(v1, ("a" + (("hello" + "\n\t") + "")))
    case "hello":
        if true {
            v35 += (v35 | (v35 + v35))
            for fi56 := 0; fi56 < 2; fi56++ {
                for fi57 := 0; fi57 < 6; fi57++ {
                    var v58 bool = (2 > 127)
                    if v58 {
                        v1 = append(v1, "")
                        v45 -= 3.14159
                        v58 = ((("a" + "") + "a") != "foo bar")
                    }
                    _ = fi57
                    _ = v58
                }
                _ = fi56
            }
        } else {
            var v59 uint32 = genFn0(255, 0, -2147483648)
            _ = v59
        }
    default:
        v44 -= v44
    }
    v60 := GenS1{F0: ("\n\t" + "foo bar"), F1: v5}
    for fi61 := 0; fi61 < 19; fi61++ {
        for fi62 := 0; fi62 < 7; fi62++ {
            if true {
                fmt.Println(v35, v45, fi62)
                fmt.Println(v5, fi62)
                fmt.Println(fi62, v2)
            }
            if ("hello" != ((((("" + "a") + "") + "") + "") + "foo bar")) {
                switch 4294967295 {
                case 128:
                    v63 := GenS6{F0: 42, F1: 1}
                    _ = v63
                case 32768:
                    fmt.Println(v35, v45, v45)
                }
                var v64 float32 = v45
                fmt.Println(fi61)
                _ = v64
            }
            switch ((("a" + "\n\t") + "\n\t") + ("foo bar" + ("" + ("hello" + "")))) {
            case "":
                fmt.Println(v45, fi62, fi62)
            case "hello":
                v2 += v2
            }
            var v65 float32 = 1.5
            _ = fi62
            _ = v65
        }
        _ = len(v4)
        v2 += v2
        _ = fi61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v60
}
