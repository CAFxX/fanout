package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS5 struct { F0 string; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 uint32; F1 uint; F2 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint }
func genFn0(p0 int8, p1 int8, p2 uint) int16 {
    v7 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807}
    if false {
        v8 := GenS2{F0: (132 <= 255)}
        if false {
            v9 := GenS1{F0: -331, F1: 819}
            v10 := GenS3{F0: 42}
            p1 = p0
            _ = v10
            _ = v9
        } else {
            p2 = p2
            if (!(603 == (p2 - (p2 << ((p2) & 63))))) {
                if false {
                    switch 4294967295 {
                    case 128:
                        switch "\n\t" {
                        case "":
                            switch "a" {
                            case "a":
                                p1 -= p1
                            case "\n\t":
                                var v11 uint32 = 65535
                                _ = v11
                            }
                        default:
                            p1 -= (p0 | -7)
                        }
                    case 630:
                        var v12 bool = ((!true) && (42 != -1))
                        _ = v12
                    }
                    var v13 uint32 = 0
                    _ = v13
                } else {
                    v14 := GenS4{F0: 850, F1: 65535, F2: 0}
                    switch "" {
                    case "hello":
                        if true {
                            for fi15 := 0; fi15 < 16; fi15++ {
                                v16 := []int{}
                                for fi17 := 0; fi17 < 5; fi17++ {
                                    if false {
                                        v18 := GenS1{F0: (fi17 + fi17), F1: -156}
                                        fmt.Println(p1)
                                        _ = v18
                                    } else {
                                        if len(v16) > 0 {
                                            v16[0] = fi17
                                        }
                                        p2 = (^(p2))
                                    }
                                    v19 := GenS5{F0: "hello", F1: 2, F2: 255}
                                    var v20 int = (^(fi15))
                                    _ = fi17
                                    _ = v19
                                    _ = v20
                                }
                                _ = fi15
                                _ = v16
                            }
                            var v21 int16 = 1000
                            _ = v21
                        }
                    default:
                        v22 := []int{128, -589}
                        _ = v22
                    }
                    if (!true) {
                        p0 = p0
                        p0 = 0
                        var v23 uint = (p2 % ((p2) | 1))
                        _ = v23
                    }
                    _ = v14
                }
                var v24 int64 = 1
                p2 = p2
                _ = v24
            } else {
                for fi25 := 0; fi25 < 10; fi25++ {
                    switch -1 {
                    case 127:
                        fmt.Println(p0)
                    case 32767:
                        fmt.Println(p2, fi25)
                    }
                    for fi26 := 0; fi26 < 6; fi26++ {
                        if true {
                            switch ("a" + "") {
                            case "hello":
                                v27 := []int{}
                                _ = v27
                            case "\n\t":
                                var v28 int16 = 42
                                _ = v28
                            case "a":
                                p1 -= p0
                                var v29 uint16 = 42
                                var v30 uint8 = 2
                                v31 := []int{fi26, -117, fi26}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        } else {
                            var v32 float32 = float32(0.0)
                            switch (^(fi25)) {
                            case 127:
                                v33 := []int{(fi26 % ((fi25) | 1)), fi26, (^(fi26)), fi26}
                                _ = v33
                            case 42:
                                fmt.Println(fi25, p0, p2)
                            case -509:
                                v32 = (v32 - ((v32 / 0.1) / (v32 + 0.0)))
                                var v34 uint32 = 65535
                                fmt.Println(fi25)
                                _ = v34
                            default:
                                p0 = p0
                            }
                            if (!false) {
                                var v35 int16 = 32767
                                switch ("" + ((((("hello" + "") + "a") + "foo bar") + (("a" + ("a" + (("hello" + ("a" + "a")) + ""))) + "hello")) + ("foo bar" + "hello"))) {
                                case "a":
                                    v36 := GenS6{F0: 4294967295, F1: -9223372036854775808}
                                    _ = v36
                                }
                                p2 += p2
                                var v37 bool = (!((3.14159 * 1.5) < v32))
                                _ = v35
                                _ = v37
                            } else {
                                p0 += p1
                            }
                            _ = v32
                        }
                        _ = fi26
                    }
                    fmt.Println(p2, p2)
                    _ = fi25
                }
                var v38 float32 = -2.25
                _ = v38
            }
            v39 := GenS3{F0: 42}
            _ = v39
        }
        v40 := GenS1{F0: -961, F1: 9223372036854775807}
        _ = v40
        _ = v8
    }
    fmt.Println(p1, p1)
    if false {
        p0 = 1
        switch (("a" + "hello") + ("\n\t" + "foo bar")) {
        case "hello":
            for fi41 := 0; fi41 < 17; fi41++ {
                p0 = 42
                v42 := GenS7{F0: 2147483648, F1: p2, F2: "\n\t"}
                v43 := GenS4{F0: -32768, F1: 934, F2: 65535}
                _ = fi41
                _ = v42
                _ = v43
            }
        }
    }
    var v44 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    _ = v7
    return -1
}
func genFn1() uint16 {
    switch "\n\t" {
    case "a":
        var v45 int16 = 761
        _ = v45
    }
    var v46 float32 = -2.25
    fmt.Println(v46)
    v46 = v46
    _ = v46
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -98
    v3 := []int{153}
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    v6 := make(map[int]bool)
    var v47 float64 = 1.5
    switch (("a" + ("hello" + "foo bar")) + "hello") {
    case "hello":
        v48 := GenS8{F0: -2147483648}
        _ = v48
    default:
        var v49 uint = 18446744073709551615
        _ = v49
    }
    switch "a" {
    case "hello":
        if (!true) {
            v50 := GenS5{F0: ("hello" + "hello"), F1: 801, F2: 0}
            _ = v50
        } else {
            v2 -= (v2 - v2)
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v47, v47)
            }
        }
    case "\n\t":
        var v51 bool = (587 > 2)
        _ = v51
    case "":
        var v52 uint64 = 0
        _ = v52
    }
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    switch "\n\t" {
    case "":
        v53 := GenS4{F0: genFn0((v2 - (^(v2))), (v2 - (v2 + (v2 >> ((65535) & 63)))), 4294967295), F1: 4294967295, F2: genFn1()}
        _ = v53
    }
    if false {
        switch ((((("" + "\n\t") + "\n\t") + "foo bar") + (("foo bar" + "hello") + "a")) + "foo bar") {
        case "foo bar":
            v54 := GenS8{F0: 1}
            _ = v54
        case "hello":
            v47 += (v47 + (v47 - v47))
        }
        var v55 int8 = v2
        _ = v55
    } else {
        fmt.Println(v47)
        for fi56 := 0; fi56 < 8; fi56++ {
            var v57 int32 = 401
            _ = fi56
            _ = v57
        }
    }
    for fi58 := 0; fi58 < 6; fi58++ {
        if (2 <= 255) {
            fmt.Println(fi58, fi58)
            v59 := GenS5{F0: "foo bar", F1: 1, F2: 0}
            v3 = append(v3, (fi58 | fi58))
            _ = v59
        }
        fmt.Println(fi58, fi58)
        _ = fi58
    }
    if false {
        var v60 uint32 = 0
        if (0 >= 1) {
            var v61 bool = false
            if false {
                switch 4294967295 {
                case 32768:
                    var v62 string = "a"
                    _ = v62
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if v61 {
                    if (v61 || (("hello" <= "\n\t") || ("\n\t" <= ((("foo bar" + "foo bar") + ("hello" + "hello")) + "hello")))) {
                        switch -763 {
                        case 1:
                            fmt.Println(v47, v47)
                        default:
                            if v61 {
                                var v63 int16 = -64
                                _ = v63
                            }
                        }
                    } else {
                        var v64 int8 = -128
                        _ = v64
                    }
                    var v65 bool = true
                    var v66 float32 = 1.5
                    _ = v65
                    _ = v66
                }
            } else {
                v2 += (v2 ^ (-(v2)))
                v47 = float64(1e10)
            }
            switch 9223372036854775807 {
            case 9223372036854775807:
                v61 = v61
            case 0:
                var v67 int8 = v2
                _ = v67
            }
            _ = v61
        }
        _ = v60
    }
    var v68 int64 = 560
    v69 := GenS9{F0: 65535}
    v70 := GenS5{F0: "foo bar", F1: 874, F2: 255}
    v2 += v2
    var v71 string = (("foo bar" + "a") + "\n\t")
    v5[-9223372036854775808] = 2147483648
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v68:", v68)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
