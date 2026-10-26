package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int; F2 int32 }
type GenS3 struct { F0 uint16; F1 string; F2 uint32 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16 }
func genFn0(p0 string) uint32 {
    p0 = p0
    p0 = ((("\n\t" + "a") + "\n\t") + (("" + "hello") + p0))
    _ = p0
    return 42
}
func genFn1(p0 int, p1 uint8) int {
    fmt.Println(p1, p1)
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int16 = 0
        v6 -= v6
        v7 := GenS1{F0: 467}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "" {
    case "":
        p0 += p0
    case "hello":
        var v8 float32 = (0.1 / 3.14159)
        _ = v8
    case "foo bar":
        v9 := GenS2{F0: (p1 - p1), F1: (p0 + 127), F2: 991}
        _ = v9
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]uint32)
    v3 := []int{}
    v4 := make(map[int]int)
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: 35}
        v1 = 1
        _ = fi10
        _ = v11
    }
    fmt.Println(v1)
    switch "\n\t" {
    case "a":
        var v12 float32 = 1e10
        _ = v12
    case "foo bar":
        switch "\n\t" {
        case "":
            v13 := GenS1{F0: 2147483647}
            _ = v13
        case "hello":
            v14 := GenS1{F0: 0}
            _ = v14
        }
    }
    var v15 int = -824
    var v16 int = (^(v15))
    if false {
        v16 += 0
        v1 = (v1 % ((int8(v15)) | 1))
        if true {
            v17 := GenS3{F0: 42, F1: "foo bar", F2: genFn0("")}
            _ = v17
        }
    }
    var v18 uint8 = 0
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 uint8 = 110
            v22 := v4[42]
            fmt.Println(v22, fi19)
            _ = fi20
            _ = v21
            _ = v22
        }
        v18 -= v18
        _ = fi19
    }
    v23 := v4[genFn1(v16, 1)]
    for fi24 := 0; fi24 < 19; fi24++ {
        v15 += v23
        switch "" {
        case "\n\t":
            v25 := GenS1{F0: 42}
            _ = v25
        case "a":
            var v26 int32 = 489
            _ = v26
        }
        delete(v2, genFn0("a"))
        _ = fi24
    }
    v1 += 127
    for fi27 := 0; fi27 < 19; fi27++ {
        fmt.Println(v1, v18, fi27)
        if false {
            if true {
                switch (("hello" + "hello") + "hello") {
                case "\n\t":
                    if ((v23 - genFn1(32767, (v18 + v18))) != genFn1(1, v18)) {
                        v15 = -1
                        v28 := GenS3{F0: 42, F1: "foo bar", F2: 2147483648}
                        if (("\n\t" >= ("\n\t" + ("\n\t" + ""))) && false) {
                            v18 -= 109
                            if false {
                                var v29 int16 = -244
                                _ = v29
                            } else {
                                v15 += v15
                                v30 := GenS1{F0: 1}
                                _ = v30
                            }
                        }
                        _ = v28
                    }
                case "a":
                    v31 := GenS2{F0: v18, F1: (v16 ^ v16), F2: -2147483648}
                    _ = v31
                case "foo bar":
                    var v32 uint16 = 329
                    _ = v32
                }
            } else {
                if true {
                    switch "\n\t" {
                    case "":
                        if (42 <= 4294967295) {
                            for fi33 := 0; fi33 < 20; fi33++ {
                                var v34 int64 = 952
                                switch "a" {
                                case "foo bar":
                                    if (42 != 1) {
                                        var v35 uint16 = 255
                                        if (true || true) {
                                            v36 := GenS2{F0: v18, F1: (v23 << ((65535) & 63)), F2: 862}
                                            var v37 bool = true
                                            _ = v36
                                            _ = v37
                                        }
                                        v2[783] = 4294967295
                                        _ = v35
                                    }
                                }
                                if (2147483647 >= 2147483647) {
                                    v18 = v18
                                    if true {
                                        var v38 int16 = 0
                                        v39 := v2[2]
                                        _ = v38
                                        _ = v39
                                    }
                                    v34 = 493
                                    if (("" + ((("foo bar" + "a") + "a") + "foo bar")) <= "hello") {
                                        v34 -= v34
                                        v16 += fi27
                                        if (3.14159 == 3.14159) {
                                            var v40 string = "a"
                                            if (!(!((v34 * v34) <= (v34 - (v34 << ((0) & 63)))))) {
                                                v41 := GenS2{F0: v18, F1: genFn1(genFn1((v23 - (v16 * (v23 + fi33))), 128), v18), F2: 42}
                                                _ = v41
                                            }
                                            v42 := GenS3{F0: 643, F1: (v40 + (("foo bar" + v40) + v40)), F2: genFn0("hello")}
                                            fmt.Println(v34, v40, fi27)
                                            v15 = genFn1(fi33, v18)
                                            _ = v40
                                            _ = v42
                                        }
                                    }
                                    if (42 < 42) {
                                        var v43 float64 = 3.14159
                                        if true {
                                            v44 := GenS2{F0: v18, F1: (^(fi27)), F2: 42}
                                            var v45 float64 = v43
                                            var v46 uint = 18446744073709551615
                                            v23 = 946
                                            var v47 string = "hello"
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                        _ = v43
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = -599
                                        }
                                    }
                                    switch ("a" + ("hello" + ("hello" + "foo bar"))) {
                                    case "":
                                        v18 = v18
                                    case "hello":
                                        v18 = (v18 << ((65535) & 63))
                                    default:
                                        v23 = v16
                                        v48 := GenS4{F0: uint16(v18), F1: v34}
                                        _ = v48
                                    }
                                    switch (-(v15)) {
                                    case 921:
                                        v1 = (^(v1))
                                    case -715:
                                        v1 += (v1 * v1)
                                    case -1:
                                        var v49 uint8 = v18
                                        _ = v49
                                    default:
                                        if (!true) {
                                            v15 -= (v16 << ((4294967295) & 63))
                                            var v50 uint8 = (v18 + (v18 - v18))
                                            _ = v50
                                        } else {
                                            switch ("a" + "hello") {
                                            case "":
                                                var v51 int16 = 42
                                                _ = v51
                                            case "hello":
                                                v16 += 4294967295
                                            case "\n\t":
                                                if (!false) {
                                                    v3 = append(v3, 128)
                                                }
                                            }
                                            if (!(false || (!(!false)))) {
                                                v34 -= 42
                                                if (!(v34 > v34)) {
                                                    switch v15 {
                                                    case -602:
                                                        _ = len(v4)
                                                    case -295:
                                                        if false {
                                                            v16 = (^(v16))
                                                        }
                                                    }
                                                }
                                                var v52 uint8 = v18
                                                _ = v52
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v53 uint64 = 48
                                                v54 := v2[1]
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                    }
                                }
                                _ = fi33
                                _ = v34
                            }
                            if ((-657 == -1) && true) {
                                v55 := GenS4{F0: 255, F1: 1}
                                _ = v55
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = v16
                                }
                                var v56 uint = 128
                                _ = v56
                            }
                            v57 := GenS3{F0: 454, F1: "\n\t", F2: 1}
                            _ = v57
                        }
                    }
                    fmt.Println(v18)
                }
                fmt.Println(fi27, v16, fi27)
                if true {
                    var v58 int64 = 42
                    for fi59 := 0; fi59 < 7; fi59++ {
                        v60 := v2[1]
                        switch (fi27 % ((0) | 1)) {
                        case -9223372036854775808:
                            fmt.Println(v16, v23)
                        }
                        v1 = -1
                        v2[(v60 - (^(v60)))] = v60
                        v61 := GenS5{F0: true}
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                    v62 := v2[42]
                    _ = v58
                    _ = v62
                }
            }
            for fi63 := 0; fi63 < 8; fi63++ {
                switch "foo bar" {
                case "hello":
                    var v64 int8 = (v1 + 100)
                    _ = v64
                case "foo bar":
                    switch (fi27 + -979) {
                    case 223:
                        var v65 string = "\n\t"
                        v18 = (v18 % ((v18) | 1))
                        var v66 uint8 = v18
                        switch v65 {
                        case "a":
                            var v67 uint32 = 42
                            _ = v67
                        }
                        fmt.Println(v23, v15, fi27)
                        v1 -= v1
                        _ = v65
                        _ = v66
                    case 32768:
                        v68 := v4[v15]
                        _ = v68
                    default:
                        var v69 uint32 = 2147483647
                        var v70 uint8 = v18
                        v71 := GenS6{F0: 1}
                        _ = v69
                        _ = v70
                        _ = v71
                    }
                }
                v18 = v18
                _ = fi63
            }
            v15 = (^(fi27))
        }
        switch v23 {
        case -531:
            for fi72 := 0; fi72 < 17; fi72++ {
                var v73 int16 = 42
                var v74 int = -869
                if true {
                    var v75 float32 = 0.1
                    _ = v75
                } else {
                    v76 := GenS2{F0: 2, F1: (v15 - (v23 << ((uint(v18)) & 63))), F2: -835}
                    _ = v76
                }
                _ = fi72
                _ = v73
                _ = v74
            }
        case 0:
            v77 := GenS4{F0: 1, F1: -1}
            _ = v77
        case 4294967295:
            switch 735 {
            case 127:
                var v78 int16 = 1000
                _ = v78
            case 1:
                var v79 bool = true
                _ = v79
            }
        default:
            v15 = (fi27 << ((1) & 63))
        }
        _ = fi27
    }
    if (false || false) {
        if (v1 != v1) {
            v18 = v18
            fmt.Println(v23, v18, v23)
            v80 := GenS7{F0: 432}
            _ = v80
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
