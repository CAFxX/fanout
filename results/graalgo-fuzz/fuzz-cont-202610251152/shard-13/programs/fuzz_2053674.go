package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint }

func main() {
    v1 := []string{}
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    var v5 float32 = 1.5
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 10; fi7++ {
        v5 += ((1.5 + v5) * v5)
        fmt.Println(v5, v6)
        _ = fi7
    }
    var v8 bool = ((!(true || true)) || (130 >= 24))
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 bool = v8
            var v12 int16 = 32767
            v12 += v12
            switch ("hello" + "foo bar") {
            case "\n\t":
                v5 += 0.1
            }
            v12 += -1
            _ = fi10
            _ = v11
            _ = v12
        }
        v8 = true
        v5 -= 0.1
        _ = fi9
    }
    if v8 {
        if (-1 >= 42) {
            v5 = float32(0.0)
            var v13 int32 = 179
            _ = v13
        } else {
            var v14 int16 = 871
            switch 4294967295 {
            case 9223372036854775807:
                if v8 {
                    v5 += float32((0.1 * 3.14159))
                    v6 = v6
                }
            case 555:
                v14 = (v14 / ((v14) | 1))
            }
            switch "a" {
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "foo bar":
                v15 := GenS1{F0: -1, F1: 175}
                _ = v15
            default:
                v5 -= v5
            }
            _ = v14
        }
        v6 -= (v6 >> ((1) & 63))
        v16 := v2["foo bar"]
        _ = v16
    }
    var v17 int = 362
    delete(v4, 4294967295)
    switch "a" {
    case "\n\t":
        v18 := GenS1{F0: -128, F1: 1}
        _ = v18
    }
    var v19 float32 = (v5 * -2.25)
    switch "a" {
    case "foo bar":
        fmt.Println(v8, v5)
    case "hello":
        for fi20 := 0; fi20 < 7; fi20++ {
            v19 = float32((-2.25 * (float64(float32((0.0 - (3.14159 * -2.25)))) - 0.1)))
            v21 := GenS1{F0: -128, F1: 18446744073709551615}
            _ = fi20
            _ = v21
        }
    case "a":
        fmt.Println(v19, v19, v6)
    }
    v8 = true
    switch 9223372036854775807 {
    case 4294967295:
        var v22 int16 = -1
        _ = v22
    case -156:
        var v23 uint = 72
        _ = v23
    case 32768:
        fmt.Println(v17)
    }
    v24 := GenS1{F0: 42, F1: 4294967295}
    switch "a" {
    case "foo bar":
        var v25 uint32 = 2
        _ = v25
    case "hello":
        v8 = v8
    default:
        v6 = v6
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        v6 = (v6 / ((1) | 1))
        if (v17 > (fi26 << ((739) & 63))) {
            switch "a" {
            case "a":
                var v27 int32 = 950
                _ = v27
            case "hello":
                if v8 {
                    if ((0.0 + (0.0 / (-2.25 - 3.14159))) > ((1.5 / 3.14159) - float64((v19 - 1.5)))) {
                        v8 = v8
                    }
                    if true {
                        for fi28 := 0; fi28 < 7; fi28++ {
                            v6 -= (^(v6))
                            _ = fi28
                        }
                        if false {
                            switch (-(v17)) {
                            case 128:
                                var v29 float32 = float32((3.14159 / -2.25))
                                _ = v29
                            case 1:
                                var v30 int64 = 42
                                _ = v30
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            switch (fi26 - v17) {
                            case 128:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case 197:
                                _ = len(v2)
                            default:
                                var v31 bool = ((fi26 + fi26) == 42)
                                _ = v31
                            }
                            switch "" {
                            case "\n\t":
                                v1 = append(v1, "hello")
                            case "foo bar":
                                var v32 uint = 599
                                _ = v32
                            case "":
                                fmt.Println(fi26)
                            default:
                                v5 = 0.1
                            }
                        }
                    } else {
                        var v33 int = (v17 | -9223372036854775808)
                        var v34 uint8 = (v6 << ((1) & 63))
                        _ = v33
                        _ = v34
                    }
                }
            case "foo bar":
                fmt.Println(v5)
            default:
                for fi35 := 0; fi35 < 17; fi35++ {
                    if true {
                        var v36 uint16 = 2
                        if (!(!v8)) {
                            if true {
                                switch ("a" + ("hello" + "a")) {
                                case "foo bar":
                                    v5 -= float32(0.1)
                                }
                                fmt.Println(v8, v6, v5)
                                v19 = 1.5
                            }
                            v37 := GenS1{F0: -122, F1: 228}
                            _ = v37
                        }
                        _ = v36
                    }
                    fmt.Println(fi26, v17, v5)
                    var v38 int16 = 922
                    _ = fi35
                    _ = v38
                }
            }
            fmt.Println(v5, v5, v5)
            v1 = append(v1, "hello")
        } else {
            var v39 bool = (!(((2 <= 0) || (255 > 742)) || (141 != 2)))
            v40 := GenS2{F0: (("" + ("foo bar" + "")) + ("\n\t" + ("" + ""))), F1: 32767}
            _ = v39
            _ = v40
        }
        var v41 int8 = 1
        _ = fi26
        _ = v41
    }
    if false {
        var v42 float64 = 3.14159
        for fi43 := 0; fi43 < 10; fi43++ {
            for fi44 := 0; fi44 < 20; fi44++ {
                if true {
                    v19 = ((1e10 + ((0.1 + (v5 / 3.14159)) - (((v5 + (v19 * -2.25)) + (v19 * 1.5)) - (1e10 * 0.1)))) * (v5 * ((((v19 + float32(v42)) + float32((v42 - (0.0 + (v42 / ((1e10 * float64((v19 + 1.5))) - -2.25)))))) + float32((v42 / 1e10))) + 0.0)))
                    switch v17 {
                    case 127:
                        _ = len(v1)
                    case -815:
                        if (true || (!(408 > 812))) {
                            fmt.Println(v5, v6, v42)
                            _ = len(v4)
                        }
                        if v8 {
                            fmt.Println(v8)
                            v19 += -2.25
                            if true {
                                fmt.Println(v5, v5)
                                fmt.Println(v19, v17)
                                v45 := GenS3{F0: 9223372036854775807, F1: 255, F2: 42}
                                _ = v45
                            } else {
                                var v46 int16 = -1
                                v47 := GenS3{F0: -640, F1: 218, F2: 219}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v46
                                _ = v47
                            }
                            if true {
                                fmt.Println(v8, fi43, fi44)
                                var v48 uint32 = 53
                                _ = v48
                            }
                        }
                    }
                } else {
                    v49 := GenS4{F0: 0, F1: v6, F2: 18446744073709551615}
                    if (((9223372036854775807 > 0) || false) || v8) {
                        var v50 float32 = ((1.5 / (v19 - (-2.25 / v5))) + 3.14159)
                        switch -455 {
                        case 567:
                            var v51 uint64 = 0
                            _ = v51
                        }
                        _ = v50
                    } else {
                        v52 := GenS2{F0: ("foo bar" + ("hello" + "foo bar")), F1: 32767}
                        _ = v52
                    }
                    _ = v49
                }
                if ((("a" + "hello") + ("hello" + ((("a" + "\n\t") + ("a" + "\n\t")) + "hello"))) > "\n\t") {
                    v17 = 1
                    if false {
                        v19 = (v19 / v5)
                        v8 = (42 > 740)
                    } else {
                        var v53 bool = true
                        switch 149 {
                        case -578:
                            _ = len(v3)
                        default:
                            if len(v1) > 0 {
                                v1[0] = ("" + (("hello" + ((("\n\t" + "a") + ("hello" + "foo bar")) + "a")) + "foo bar"))
                            }
                            if v8 {
                                v42 += ((v42 / ((((3.14159 + (float64((((v5 + v19) / 1.0) + v19)) / v42)) + (3.14159 / v42)) + (v42 + -2.25)) / 3.14159)) * v42)
                            }
                        }
                        _ = v53
                    }
                    var v54 uint32 = 65535
                    _ = v54
                }
                _ = fi44
            }
            switch (v17 * (-(v17))) {
            case -2147483648:
                var v55 int = (v17 * (v17 & (v17 & 128)))
                _ = v55
            case 32768:
                fmt.Println(fi43, v42, fi43)
            case -9223372036854775808:
                var v56 uint64 = 42
                _ = v56
            }
            _ = fi43
        }
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
