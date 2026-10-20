package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 int8; F2 int8 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 bool; F1 int16; F2 uint8 }
type GenS9 struct { F0 string; F1 int; F2 int32 }
type GenS10 struct { F0 uint64; F1 int64 }

func main() {
    v1 := []string{}
    var v2 uint = 42
    var v3 int8 = 42
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    if (42 >= 0) {
        v6 := GenS1{F0: 934}
        v7 := GenS1{F0: 255}
        v8 := GenS1{F0: 2}
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 bool = (false && (!true))
        _ = v9
    }
    var v10 float64 = (3.14159 - float64(3.14159))
    fmt.Println(v3)
    if true {
        v11 := GenS1{F0: 2}
        var v12 int32 = 2147483647
        delete(v4, 0)
        _ = v11
        _ = v12
    }
    v2 = v2
    switch "hello" {
    case "a":
        v13 := GenS1{F0: 255}
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 14; fi14++ {
            for fi15 := 0; fi15 < 1; fi15++ {
                if ((v3 - v3) > v3) {
                    fmt.Println(fi15, fi15)
                    if false {
                        var v16 int64 = 42
                        var v17 string = "hello"
                        if (193 > 255) {
                            var v18 string = "hello"
                            v18 = ("a" + (("foo bar" + ("" + "hello")) + (v18 + v18)))
                            var v19 uint8 = 2
                            v20 := GenS2{F0: int16(v16)}
                            if (true && true) {
                                var v21 uint32 = 4294967295
                                var v22 int64 = -1
                                var v23 bool = (255 <= 849)
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        switch (fi15 / ((fi15) | 1)) {
                        case -495:
                            v24 := GenS1{F0: 514}
                            _ = v24
                        default:
                            v25 := GenS1{F0: 2}
                            _ = v25
                        }
                        v16 = v16
                        v26 := v5[fi14]
                        _ = v16
                        _ = v17
                        _ = v26
                    }
                    v3 = v3
                }
                if (2147483648 <= 1) {
                    var v27 int = (fi14 * fi15)
                    v1 = append(v1, "")
                    _ = v27
                } else {
                    if (false || (139 == 2)) {
                        _ = len(v1)
                    }
                }
                var v28 float64 = (v10 - 0.0)
                _ = fi15
                _ = v28
            }
            if true {
                fmt.Println(v2, v10)
                v29 := GenS2{F0: -1}
                _ = v29
            }
            switch "" {
            case "foo bar":
                fmt.Println(fi14, fi14, v2)
            }
            _ = fi14
        }
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        fmt.Println(fi30, v10, fi30)
        if false {
            if (-32768 != 32767) {
                switch ("a" + ("foo bar" + (("\n\t" + "\n\t") + "\n\t"))) {
                case "":
                    if false {
                        if true {
                            v2 = 65535
                            switch (^(fi30)) {
                            case 32768:
                                if (!false) {
                                    _ = len(v1)
                                    v31 := GenS3{F0: v3}
                                    v3 = (v3 << ((v2) & 63))
                                    _ = v31
                                }
                            case 127:
                                var v32 uint = (^(v2))
                                _ = v32
                            case 128:
                                for fi33 := 0; fi33 < 5; fi33++ {
                                    v34 := GenS4{F0: v2, F1: 2, F2: -32768}
                                    _ = fi33
                                    _ = v34
                                }
                            }
                        }
                        var v35 int32 = 0
                        if (true && false) {
                            for fi36 := 0; fi36 < 16; fi36++ {
                                v37 := GenS5{F0: v35}
                                v38 := GenS2{F0: -839}
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                            if true {
                                v39 := GenS4{F0: 2, F1: 2, F2: int16(v3)}
                                if (!(((!(1 < 42)) || (1 < 255)) && true)) {
                                    v40 := GenS4{F0: v2, F1: 2, F2: 355}
                                    _ = v40
                                }
                                v3 += v3
                                _ = v39
                            }
                        }
                        _ = v35
                    } else {
                        if (true && false) {
                            for fi41 := 0; fi41 < 17; fi41++ {
                                switch ("foo bar" + "a") {
                                case "foo bar":
                                    var v42 int = 9223372036854775807
                                    _ = v42
                                case "\n\t":
                                    fmt.Println(v10, v3, fi30)
                                }
                                if false {
                                    switch ("" + (("a" + "a") + "")) {
                                    case "a":
                                        if true {
                                            v1 = append(v1, "a")
                                        }
                                    case "":
                                        var v43 int16 = 1000
                                        _ = v43
                                    case "\n\t":
                                        var v44 uint = (v2 / (((^(v2))) | 1))
                                        if (!(270 >= -1)) {
                                            v44 = 0
                                        }
                                        _ = v44
                                    default:
                                        v45 := v5[(-(fi41))]
                                        if (fi30 != 128) {
                                            v46 := GenS5{F0: -1}
                                            switch (fi30 + 2147483647) {
                                            case 127:
                                                var v47 uint64 = 259
                                                _ = v47
                                            case -2147483648:
                                                switch v45 {
                                                case "a":
                                                    v48 := GenS6{F0: int(v3), F1: (v3 + (v3 | 0)), F2: v3}
                                                    _ = v48
                                                }
                                            }
                                            var v49 int32 = 1
                                            _ = v46
                                            _ = v49
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(fi41)
                                            case "foo bar":
                                                v50 := GenS1{F0: 255}
                                                _ = v50
                                            case "\n\t":
                                                var v51 int8 = (v3 - v3)
                                                v52 := GenS5{F0: -232}
                                                var v53 bool = true
                                                fmt.Println(v10, fi41)
                                                v53 = (int(v3) != (fi30 % ((fi30) | 1)))
                                                v54 := GenS7{F0: 92, F1: v51}
                                                _ = len(v1)
                                                fmt.Println(fi41)
                                                switch v45 {
                                                case "\n\t":
                                                    v55 := GenS1{F0: 255}
                                                    _ = v55
                                                case "hello":
                                                    v56 := GenS7{F0: (v3 ^ -1), F1: (v51 % ((v3) | 1))}
                                                    _ = v56
                                                }
                                                if false {
                                                    fmt.Println(v10)
                                                }
                                                _ = v51
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        v3 += (v3 & (v3 + 79))
                                        _ = v45
                                    }
                                    if false {
                                        v57 := GenS8{F0: true, F1: 1000, F2: 255}
                                        switch "hello" {
                                        case "foo bar":
                                            v10 += (v10 + -2.25)
                                        }
                                        v3 = (v3 / ((v3) | 1))
                                        _ = v57
                                    }
                                    var v58 uint16 = 255
                                    _ = v58
                                }
                                _ = fi41
                            }
                            v59 := v4[(fi30 & fi30)]
                            _ = v59
                        }
                        switch "hello" {
                        case "hello":
                            if true {
                                v2 = (v2 + v2)
                            } else {
                                if (true || false) {
                                    for fi60 := 0; fi60 < 11; fi60++ {
                                        v2 = (^(v2))
                                        _ = fi60
                                    }
                                }
                            }
                        case "":
                            var v61 uint16 = 640
                            _ = v61
                        default:
                            var v62 uint16 = 255
                            _ = v62
                        }
                        var v63 int32 = 2147483647
                        _ = v63
                    }
                case "foo bar":
                    var v64 int8 = v3
                    _ = v64
                case "hello":
                    var v65 string = ("" + ("\n\t" + "\n\t"))
                    _ = v65
                }
                var v66 float64 = v10
                fmt.Println(v3, v3)
                _ = v66
            }
            v10 += v10
            v2 -= 257
        }
        _ = fi30
    }
    fmt.Println(v3)
    delete(v4, -164)
    fmt.Println(v3, v2)
    v67 := GenS4{F0: v2, F1: 2147483648, F2: 32767}
    v68 := GenS9{F0: "hello", F1: 32768, F2: -845}
    v69 := GenS10{F0: 1, F1: -9223372036854775808}
    var v70 int16 = 989
    if len(v1) > 0 {
        v1[0] = ""
    }
    if len(v1) > 0 {
        v1[0] = (("a" + "\n\t") + ("\n\t" + ("a" + "foo bar")))
    }
    v70 = v70
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
