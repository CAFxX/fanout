package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint16 }
func genFn0(p0 uint8) uint16 {
    fmt.Println(p0)
    switch 32767 {
    case 67:
        var v5 int16 = -182
        _ = v5
    }
    _ = p0
    return 0
}
func genFn1(p0 uint16) uint {
    switch "hello" {
    case "\n\t":
        var v6 uint = 42
        _ = v6
    case "a":
        p0 = (p0 / (((p0 * genFn0(127))) | 1))
    case "foo bar":
        fmt.Println(p0)
    default:
        v7 := GenS1{F0: 0, F1: ("a" + (("\n\t" + ("" + "")) + ("foo bar" + "\n\t")))}
        _ = v7
    }
    var v8 float32 = -2.25
    if (-128 > 127) {
        var v9 uint32 = 589
        if false {
            var v10 uint64 = 2
            v11 := []int{32768, 790}
            v12 := GenS1{F0: v9, F1: "foo bar"}
            _ = v10
            _ = v11
            _ = v12
        } else {
            v8 = float32(0.0)
        }
        _ = v9
    }
    _ = p0
    _ = v8
    return 951
}

func main() {
    v1 := []string{"hello", "\n\t", "\n\t"}
    var v2 uint32 = 607
    var v3 uint = 4294967295
    var v4 int16 = 1
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 = (v4 | -1)
        _ = fi13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 float64 = 1e10
        v1 = append(v1, "a")
        _ = fi14
        _ = v15
    }
    if true {
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int16 = v4
        v1 = append(v1, "")
        var v18 int = (fi16 + fi16)
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int8 = 127
    v1 = append(v1, "hello")
    for fi20 := 0; fi20 < 5; fi20++ {
        switch 65 {
        case -331:
            var v21 uint16 = 940
            _ = v21
        case -9223372036854775808:
            switch (fi20 << ((0) & 63)) {
            case -579:
                for fi22 := 0; fi22 < 1; fi22++ {
                    if true {
                        var v23 uint8 = 193
                        switch (("a" + "foo bar") + (("\n\t" + "foo bar") + ("a" + "foo bar"))) {
                        case "foo bar":
                            v19 = (v19 * v19)
                        default:
                            v24 := GenS1{F0: v2, F1: "foo bar"}
                            _ = v24
                        }
                        var v25 uint32 = (v2 << ((1) & 63))
                        _ = v23
                        _ = v25
                    }
                    switch ("\n\t" + ("foo bar" + "a")) {
                    case "":
                        switch 128 {
                        case -2147483648:
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                        case -997:
                            fmt.Println(v19)
                        default:
                            v4 -= v4
                        }
                    case "\n\t":
                        if (-1 > -813) {
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                    case "foo bar":
                        fmt.Println(v4, v19, v4)
                    default:
                        v26 := GenS1{F0: v2, F1: ("foo bar" + "foo bar")}
                        _ = v26
                    }
                    switch fi20 {
                    case -517:
                        var v27 uint64 = 196
                        v27 = (v27 >> (((^(v3))) & 63))
                        var v28 int64 = 0
                        _ = v27
                        _ = v28
                    case -9223372036854775808:
                        fmt.Println(fi22)
                    case 128:
                        if ((v2 + (v2 | 65535)) < v2) {
                            switch 192 {
                            case -388:
                                var v29 int8 = v19
                                _ = v29
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                    default:
                        v30 := GenS1{F0: v2, F1: (("\n\t" + "foo bar") + "\n\t")}
                        _ = v30
                    }
                    _ = fi22
                }
            case 4294967295:
                for fi31 := 0; fi31 < 17; fi31++ {
                    v2 = (v2 / ((607) | 1))
                    var v32 int32 = 1
                    if true {
                        v33 := GenS1{F0: (v2 | v2), F1: ("hello" + "\n\t")}
                        v34 := GenS1{F0: (^(v2)), F1: "foo bar"}
                        v35 := GenS2{F0: ((v3 + genFn1(612)) <= genFn1(65535))}
                        _ = v33
                        _ = v34
                        _ = v35
                    } else {
                        if (((3.14159 + ((float64(1e10) / ((0.1 - ((float64(3.14159) - 1e10) * -2.25)) / ((0.0 + float64(3.14159)) * (((0.0 * ((3.14159 + 0.0) + (0.1 / -2.25))) * 3.14159) + ((1.5 - float64(-2.25)) + 0.0))))) - 1e10)) - 0.1) >= 0.0) {
                            v36 := GenS2{F0: (783 < 176)}
                            var v37 int32 = (^(v32))
                            _ = v36
                            _ = v37
                        } else {
                            _ = len(v1)
                        }
                        v19 += (v19 & v19)
                    }
                    _ = fi31
                    _ = v32
                }
            default:
                if (!false) {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        v19 -= (-(v19))
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            fmt.Println(v2)
                        }
                        fmt.Println(v19)
                        _ = fi38
                    }
                }
            }
        case 341:
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        default:
            if false {
                switch "a" {
                case "\n\t":
                    for fi39 := 0; fi39 < 15; fi39++ {
                        v40 := GenS2{F0: false}
                        _ = fi39
                        _ = v40
                    }
                case "":
                    if false {
                        v41 := GenS3{F0: "\n\t"}
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        _ = v41
                    }
                }
                v2 = v2
            }
        }
        if (!(25 < 1)) {
            var v42 int8 = (v19 & 42)
            v2 = 65535
            _ = v42
        }
        for fi43 := 0; fi43 < 6; fi43++ {
            switch fi20 {
            case 9223372036854775807:
                var v44 uint8 = 42
                v45 := GenS4{F0: genFn0(v44), F1: uint16(v44), F2: 32767}
                if (fi20 >= (fi20 | 9223372036854775807)) {
                    v44 = (v44 | (v44 / (((v44 - (v44 * v44))) | 1)))
                    v4 = v4
                    var v46 uint64 = 18446744073709551615
                    _ = v46
                }
                _ = v44
                _ = v45
            }
            v47 := GenS4{F0: 2, F1: 42, F2: 1}
            _ = fi43
            _ = v47
        }
        _ = fi20
    }
    v2 -= (^(v2))
    v2 = (v2 % ((1) | 1))
    fmt.Println(v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
