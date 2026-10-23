package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint8; F2 string }
type GenS4 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 int16; F1 int; F2 bool }
type GenS8 struct { F0 bool; F1 int16; F2 int16 }
type GenS9 struct { F0 int16; F1 uint32; F2 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int; F1 int16; F2 uint16 }
func genFn0(p0 int, p1 uint) int {
    v4 := []int{-649, (p0 * (p0 + -89))}
    v5 := GenS1{F0: (p1 | 1), F1: ("a" + "a")}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p0 % ((p0) | 1))
}
func genFn1(p0 uint32, p1 uint32) int {
    p0 = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        p0 += p0
        var v7 int64 = 652
        _ = fi6
        _ = v7
    }
    var v8 int16 = -546
    switch "hello" {
    case "foo bar":
        v9 := GenS2{F0: 65535}
        _ = v9
    case "":
        v10 := GenS2{F0: 255}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 0
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]bool)
    var v3 int32 = (v1 << ((255) & 63))
    var v11 uint = 576
    v11 = 18446744073709551615
    if (!true) {
        v1 = 954
        v12 := GenS3{F0: ("" + "hello"), F1: 139, F2: ("\n\t" + "")}
        _ = v12
    }
    v13 := GenS4{F0: -1, F1: 1, F2: 0}
    v11 = (^(v11))
    v14 := v2["foo bar"]
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int8 = 0
        var v17 int32 = v1
        if false {
            if v14 {
                v11 += 2
                if (-2.25 > 1.5) {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v19 := []int{32768, genFn1(2147483647, 65535)}
                        v3 -= v3
                        switch "" {
                        case "\n\t":
                            var v20 uint32 = 65535
                            _ = v20
                        case "foo bar":
                            v19 = append(v19, (fi18 % ((1) | 1)))
                        }
                        _ = fi18
                        _ = v19
                    }
                    v16 = v16
                    var v21 float32 = (1.5 + (((1e10 + 0.0) + (((((0.1 * float32((0.0 + 0.1))) + float32(0.1)) * 1e10) + (1.5 - (3.14159 / 1e10))) - float32((0.1 * 3.14159)))) / float32(-2.25)))
                    _ = v21
                }
            } else {
                v22 := []int{}
                _ = v22
            }
            switch (fi15 * 2147483647) {
            case 128:
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 float32 = 0.1
                    v25 := GenS5{F0: (-(fi15))}
                    if v14 {
                        if v14 {
                            switch -118 {
                            case 127:
                                var v26 bool = (!true)
                                if (127 == 2) {
                                    var v27 float64 = float64(v24)
                                    v28 := GenS6{F0: -9223372036854775808, F1: 506}
                                    _ = v27
                                    _ = v28
                                }
                                _ = v26
                            case -9223372036854775808:
                                v29 := GenS7{F0: 1000, F1: fi23, F2: v14}
                                _ = v29
                            case -118:
                                v3 = v3
                            }
                            var v30 uint32 = 0
                            var v31 uint64 = 695
                            _ = v30
                            _ = v31
                        } else {
                            var v32 uint64 = 1
                            var v33 uint32 = 0
                            switch (("\n\t" + "hello") + "") {
                            case "\n\t":
                                var v34 bool = ((^(v33)) < v33)
                                _ = v34
                            }
                            v33 += 42
                            if (2 > 128) {
                                v35 := GenS3{F0: "foo bar", F1: 128, F2: "a"}
                                _ = len(v2)
                                _ = len(v2)
                                _ = v35
                            } else {
                                v36 := GenS2{F0: 1}
                                if (v14 || (2147483647 < (v33 + v33))) {
                                    var v37 int16 = -32768
                                    var v38 uint64 = v32
                                    _ = v37
                                    _ = v38
                                } else {
                                    v16 += (^(v16))
                                    switch ("\n\t" + "\n\t") {
                                    case "":
                                        var v39 uint8 = 255
                                        _ = v39
                                    }
                                    v40 := []int{(fi23 << ((v11) & 63))}
                                    var v41 float32 = v24
                                    switch fi15 {
                                    case 9223372036854775807:
                                        switch "foo bar" {
                                        case "foo bar":
                                            fmt.Println(v33)
                                        }
                                    case -1:
                                        v42 := GenS6{F0: 675, F1: 42}
                                        _ = v42
                                    case 4294967295:
                                        if (!(!false)) {
                                            fmt.Println(fi23, v41, v33)
                                        }
                                    }
                                    var v43 uint16 = 1
                                    _ = v40
                                    _ = v41
                                    _ = v43
                                }
                                var v44 uint = (v11 % ((v11) | 1))
                                _ = v36
                                _ = v44
                            }
                            _ = v32
                            _ = v33
                        }
                        if ((((-2.25 * (v24 - 1.5)) / v24) * 1.5) == ((3.14159 * (0.1 / v24)) * ((-2.25 - ((v24 + v24) / 3.14159)) / v24))) {
                            var v45 uint64 = 2
                            _ = v45
                        } else {
                            v46 := GenS8{F0: true, F1: -32768, F2: 230}
                            v47 := GenS9{F0: 1, F1: 760, F2: false}
                            _ = v46
                            _ = v47
                        }
                        if true {
                            var v48 int64 = -9223372036854775808
                            switch 95 {
                            case 579:
                                var v49 float64 = ((-2.25 - 1e10) * (3.14159 / (0.1 - 1.5)))
                                _ = v49
                            default:
                                v50 := GenS6{F0: v48, F1: 776}
                                _ = v50
                            }
                            _ = v48
                        }
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            v11 -= v11
        } else {
            for fi51 := 0; fi51 < 5; fi51++ {
                var v52 uint16 = 1
                v53 := GenS8{F0: v14, F1: 0, F2: 32767}
                switch -153 {
                case 210:
                    switch "foo bar" {
                    case "hello":
                        v54 := GenS2{F0: (v52 + v52)}
                        var v55 float32 = (-2.25 / (1e10 / -2.25))
                        _ = v54
                        _ = v55
                    case "foo bar":
                        if v14 {
                            fmt.Println(v3, v11, v1)
                            v16 = v16
                            v52 -= 2
                        } else {
                            v3 += (v17 * (-(v3)))
                        }
                    }
                case -247:
                    switch "a" {
                    case "a":
                        switch (fi51 | genFn0((fi15 & genFn1(1, 2147483647)), (v11 - v11))) {
                        case 176:
                            var v56 string = "foo bar"
                            _ = v56
                        }
                    case "":
                        var v57 float32 = 1.5
                        _ = v57
                    case "foo bar":
                        v58 := GenS10{F0: 826}
                        if ((!(uint32(v52) < 2)) || v14) {
                            if (genFn1(434, 1) != fi51) {
                                var v59 uint16 = 42
                                var v60 int16 = -32768
                                _ = v59
                                _ = v60
                            }
                        }
                        v16 = v16
                        _ = v58
                    }
                case 128:
                    if v14 {
                        var v61 int = fi15
                        var v62 int32 = (^(v3))
                        _ = v61
                        _ = v62
                    }
                    fmt.Println(v1, fi51)
                default:
                    switch (("hello" + "\n\t") + (("foo bar" + ("" + "")) + "a")) {
                    case "\n\t":
                        v63 := []int{1, (-(fi51)), fi51, fi15}
                        _ = v63
                    case "a":
                        fmt.Println(fi15, v11)
                    }
                }
                _ = fi51
                _ = v52
                _ = v53
            }
            v64 := []int{(fi15 + (fi15 & (-(fi15)))), fi15}
            _ = v64
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v65 := GenS8{F0: v14, F1: 1000, F2: -1}
    switch "\n\t" {
    case "":
        for fi66 := 0; fi66 < 11; fi66++ {
            if ((221 <= 255) || ((42 > 42) || v14)) {
                switch ("\n\t" + "a") {
                case "foo bar":
                    var v67 int8 = -1
                    _ = v67
                case "\n\t":
                    v68 := GenS11{F0: fi66, F1: -32768, F2: 1}
                    _ = v68
                case "a":
                    v3 += (v3 ^ (v1 << (((v11 << ((v11) & 63))) & 63)))
                }
            } else {
                fmt.Println(v1)
            }
            _ = fi66
        }
    default:
        v69 := v2[("\n\t" + ("foo bar" + ("" + "hello")))]
        _ = v69
    }
    if (!v14) {
        switch "foo bar" {
        case "":
            delete(v2, "")
        }
    } else {
        v1 += v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v65
}
