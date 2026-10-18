package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint, p1 string, p2 uint8) uint32 {
    var v4 int8 = 100
    var v5 uint = 2
    p2 -= p2
    var v6 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 348
}
func genFn1(p0 int, p1 uint, p2 bool) uint {
    p1 += 491
    v7 := GenS1{F0: true}
    switch p0 {
    case 9223372036854775807:
        fmt.Println(p1, p0)
    case -1:
        var v8 string = "foo bar"
        _ = v8
    case 42:
        var v9 float64 = 0.0
        _ = v9
    default:
        switch ("hello" + "") {
        case "":
            switch "hello" {
            case "\n\t":
                v10 := []int{p0}
                _ = v10
            case "a":
                var v11 uint64 = 0
                _ = v11
            }
        }
    }
    v12 := []int{0, p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return 0
}

func main() {
    v1 := []string{"hello", (("a" + "") + ""), "a", ""}
    var v2 bool = (!(2147483647 >= 42))
    var v3 float64 = 1e10
    switch (((("a" + "\n\t") + "\n\t") + "hello") + (("a" + "\n\t") + ("foo bar" + "a"))) {
    case "foo bar":
        _ = len(v1)
    case "a":
        v13 := GenS1{F0: false}
        _ = v13
    }
    var v14 int8 = -128
    fmt.Println(v3, v3)
    v2 = v2
    if ("" == ("\n\t" + (("" + "") + "a"))) {
        v1 = append(v1, "a")
    } else {
        if (true || (105 != -9223372036854775808)) {
            v14 = (v14 * v14)
            v2 = false
            var v15 uint8 = 1
            _ = v15
        }
        v3 = 3.14159
    }
    switch 42 {
    case -174:
        for fi16 := 0; fi16 < 12; fi16++ {
            switch "foo bar" {
            case "":
                switch (fi16 & (fi16 / ((fi16) | 1))) {
                case 9223372036854775807:
                    var v17 string = "a"
                    _ = v17
                case -9223372036854775808:
                    if (((float64((0.1 * 1e10)) - float64(0.0)) - ((0.0 - float64((-2.25 + 1e10))) - 0.1)) != v3) {
                        v18 := GenS2{F0: 815, F1: "\n\t"}
                        for fi19 := 0; fi19 < 16; fi19++ {
                            switch "a" {
                            case "\n\t":
                                v2 = ((-9223372036854775808 == -242) && true)
                            case "a":
                                if (!(!true)) {
                                    var v20 int8 = (^(v14))
                                    if (((18446744073709551615 >= 2) && v2) && v2) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v14 = v20
                                        case "":
                                            v3 = float64(1e10)
                                        }
                                        v21 := GenS3{F0: 34, F1: genFn0(476, "a", 42), F2: 18446744073709551615}
                                        var v22 int16 = 0
                                        _ = len(v1)
                                        _ = v21
                                        _ = v22
                                    }
                                    _ = v20
                                } else {
                                    switch "a" {
                                    case "\n\t":
                                        switch (fi19 - (fi16 % ((-9223372036854775808) | 1))) {
                                        case 793:
                                            v3 -= 1.5
                                        }
                                        fmt.Println(v2)
                                    }
                                }
                            case "":
                                v23 := GenS2{F0: 482, F1: ""}
                                _ = v23
                            default:
                                if true {
                                    fmt.Println(v3)
                                } else {
                                    v14 -= (v14 << ((42) & 63))
                                }
                            }
                            switch (("foo bar" + "") + "hello") {
                            case "foo bar":
                                v14 = (v14 & v14)
                            case "\n\t":
                                v14 = 42
                            }
                            switch "a" {
                            case "foo bar":
                                switch (fi19 % (((fi16 + fi16)) | 1)) {
                                case 4294967295:
                                    fmt.Println(fi16, v2, v14)
                                }
                            case "":
                                if v2 {
                                    if (1 < -1) {
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                        var v24 float64 = float64((3.14159 / (3.14159 * 0.1)))
                                        _ = v24
                                    }
                                } else {
                                    v14 = (-(v14))
                                    if true {
                                        var v25 uint32 = 2147483648
                                        _ = len(v1)
                                        _ = v25
                                    }
                                    fmt.Println(v14, fi19)
                                }
                            case "a":
                                v2 = v2
                            default:
                                v14 = 100
                            }
                            _ = fi19
                        }
                        if (0 < 1) {
                            fmt.Println(v2)
                            var v26 uint8 = 1
                            if true {
                                fmt.Println(v26, v26)
                            } else {
                                var v27 int = fi16
                                v3 = v3
                                _ = v27
                            }
                            _ = v26
                        }
                        _ = v18
                    } else {
                        if v2 {
                            for fi28 := 0; fi28 < 8; fi28++ {
                                if (264 > -178) {
                                    var v29 int64 = 0
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v29
                                } else {
                                    v30 := GenS1{F0: false}
                                    v14 = 100
                                    if v2 {
                                        v31 := GenS3{F0: 127, F1: 1, F2: genFn1(fi28, 0, v2)}
                                        if (!v2) {
                                            switch (fi16 - fi28) {
                                            case -1:
                                                v32 := GenS4{F0: (v14 ^ (v14 - (-(v14))))}
                                                _ = v32
                                            case 127:
                                                v33 := GenS4{F0: v14}
                                                _ = v33
                                            default:
                                                v34 := GenS3{F0: 142, F1: 4294967295, F2: 501}
                                                _ = v34
                                            }
                                            v35 := GenS5{F0: int16(fi28)}
                                            v3 = (0.1 / v3)
                                            _ = v35
                                        } else {
                                            var v36 uint = 2
                                            _ = v36
                                        }
                                        fmt.Println(fi28, v14, v2)
                                        _ = v31
                                    }
                                    _ = v30
                                }
                                switch -9223372036854775808 {
                                case 533:
                                    if v2 {
                                        v1 = append(v1, "hello")
                                        var v37 int32 = 2147483647
                                        var v38 float64 = ((((v3 + -2.25) + (1.5 / 0.1)) / 0.1) * ((float64(0.1) - (v3 / (float64(((0.0 - (float32(0.1) - (0.0 / 1.0))) * (0.1 * 0.0))) * v3))) / v3))
                                        var v39 int32 = (v37 << ((genFn1(fi16, 1, ("foo bar" == "hello"))) & 63))
                                        v39 += 1
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                case 2147483647:
                                    if len(v1) > 0 {
                                        v1[0] = ((("" + "foo bar") + "a") + ("hello" + ("\n\t" + (("hello" + "\n\t") + "a"))))
                                    }
                                case 335:
                                    var v40 int64 = 812
                                    _ = v40
                                }
                                switch "" {
                                case "a":
                                    var v41 bool = (2 > 18446744073709551615)
                                    _ = v41
                                case "\n\t":
                                    v42 := GenS4{F0: (v14 / (((v14 ^ (v14 % (((v14 * (v14 % (((v14 & -109)) | 1)))) | 1)))) | 1))}
                                    _ = v42
                                case "":
                                    var v43 float32 = 1e10
                                    _ = v43
                                }
                                switch (fi28 ^ (fi16 % (((fi16 - (fi28 & 95))) | 1))) {
                                case 32768:
                                    var v44 uint8 = 243
                                    _ = v44
                                case -110:
                                    fmt.Println(fi28)
                                }
                                switch fi16 {
                                case 42:
                                    var v45 int32 = 1
                                    _ = v45
                                case 4294967295:
                                    v46 := GenS4{F0: 42}
                                    _ = v46
                                case -9223372036854775808:
                                    var v47 int32 = -92
                                    _ = v47
                                default:
                                    v2 = false
                                }
                                _ = fi28
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                default:
                    var v48 int = (fi16 << ((819) & 63))
                    _ = v48
                }
            }
            switch ("\n\t" + "") {
            case "\n\t":
                v14 = v14
            case "a":
                v49 := GenS2{F0: 689, F1: "hello"}
                _ = v49
            }
            _ = fi16
        }
    }
    if v2 {
        var v50 float32 = 0.1
        v51 := GenS4{F0: v14}
        _ = v50
        _ = v51
    }
    fmt.Println(v2, v3, v3)
    v3 += v3
    fmt.Println(v3)
    v1 = append(v1, "")
    v2 = (!v2)
    var v52 uint = 228
    v52 += 556
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v52
}
