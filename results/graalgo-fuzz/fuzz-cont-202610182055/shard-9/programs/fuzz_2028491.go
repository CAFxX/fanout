package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 bool; F2 int8 }
func genFn0(p0 uint64, p1 int32, p2 uint16) uint8 {
    var v5 int16 = -32768
    v6 := GenS1{F0: ((("\n\t" + ((("hello" + "") + ("foo bar" + "\n\t")) + "\n\t")) + "foo bar") + "\n\t")}
    switch 32768 {
    case 0:
        v7 := GenS1{F0: "a"}
        _ = v7
    default:
        v8 := GenS1{F0: "a"}
        _ = v8
    }
    var v9 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v9
    return 44
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"a", ((("" + ("hello" + ("a" + ("foo bar" + ((("hello" + "a") + ("foo bar" + ("\n\t" + "foo bar"))) + ""))))) + "\n\t") + (("\n\t" + (("\n\t" + "\n\t") + (((("foo bar" + (("foo bar" + "\n\t") + ("a" + "hello"))) + ((("\n\t" + "foo bar") + "") + ((("hello" + ("a" + "a")) + "hello") + ""))) + "foo bar") + (("foo bar" + "foo bar") + "a")))) + "hello"))}
    v3 := make(map[uint32]string)
    var v4 float32 = 1e10
    var v10 float32 = (((v4 * 0.1) / 3.14159) + (3.14159 / (v4 * v4)))
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            fmt.Println(fi11, v10, fi11)
            _ = fi12
        }
        v13 := GenS1{F0: "hello"}
        if true {
            if true {
                fmt.Println(v4, fi11)
            } else {
                v14 := GenS1{F0: ""}
                v10 = (v10 / float32(((((0.0 + (1e10 * (1.5 * 1.5))) * (-2.25 / 0.1)) * float64(float32((-2.25 / (0.1 * (1.5 - 0.1)))))) / float64(v10))))
                _ = v14
            }
            v15 := GenS2{F0: 2, F1: 0}
            v16 := GenS2{F0: 1, F1: -379}
            _ = v15
            _ = v16
        }
        _ = fi11
        _ = v13
    }
    v17 := GenS3{F0: -128, F1: 100}
    switch ("\n\t" + "\n\t") {
    case "":
        v18 := GenS1{F0: ""}
        _ = v18
    case "\n\t":
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 string = "foo bar"
            _ = fi19
            _ = v20
        }
    default:
        v4 = v4
    }
    switch 1 {
    case 32767:
        for fi21 := 0; fi21 < 10; fi21++ {
            v4 -= (v10 - 1.5)
            _ = fi21
        }
    case 354:
        fmt.Println(v4, v4)
    case -903:
        delete(v1, 9223372036854775807)
    }
    switch 42 {
    case 32768:
        for fi22 := 0; fi22 < 18; fi22++ {
            if true {
                for fi23 := 0; fi23 < 3; fi23++ {
                    if (146 <= 2) {
                        if true {
                            var v24 uint = 255
                            v10 += v10
                            _ = v24
                        } else {
                            v10 = float32(0.0)
                            fmt.Println(v10)
                            var v25 int8 = 82
                            _ = v25
                        }
                        v26 := GenS3{F0: -128, F1: 42}
                        _ = v26
                    }
                    _ = fi23
                }
                var v27 int8 = -1
                switch "hello" {
                case "\n\t":
                    for fi28 := 0; fi28 < 14; fi28++ {
                        _ = len(v1)
                        _ = fi28
                    }
                }
                _ = v27
            } else {
                switch "hello" {
                case "hello":
                    switch (fi22 / (((fi22 | (fi22 % ((fi22) | 1)))) | 1)) {
                    case -373:
                        switch "\n\t" {
                        case "a":
                            var v29 uint = 1
                            _ = v29
                        }
                    case 593:
                        v30 := GenS2{F0: 32767, F1: -1}
                        _ = v30
                    default:
                        for fi31 := 0; fi31 < 17; fi31++ {
                            if false {
                                switch (fi22 << ((1) & 63)) {
                                case 475:
                                    if len(v2) > 0 {
                                        v2[0] = (((("a" + ("\n\t" + "")) + "") + "a") + "foo bar")
                                    }
                                    if true {
                                        if (2147483647 < 65535) {
                                            var v32 uint8 = 128
                                            var v33 int = 128
                                            v32 = v32
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v32
                                            _ = v33
                                        }
                                    } else {
                                        v4 += (float32((1.5 * ((((0.1 + 3.14159) - 0.1) / 1.0) + (1.5 / 1.0)))) / 1e10)
                                        fmt.Println(fi22, v4, v10)
                                        if true {
                                            v34 := v1[fi22]
                                            v35 := GenS1{F0: v34}
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                case -125:
                                    v4 -= ((((v4 / v10) / float32(1.5)) + float32((1e10 + 0.1))) - 0.1)
                                default:
                                    fmt.Println(v4, fi22, v10)
                                }
                            } else {
                                fmt.Println(fi31, v4)
                                var v36 uint32 = 50
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v36
                            }
                            v10 += 0.0
                            if (0.1 > 0.1) {
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    _ = len(v1)
                                case "":
                                    var v37 int8 = 42
                                    v2 = append(v2, "foo bar")
                                    _ = v37
                                default:
                                    var v38 uint32 = 1
                                    if (801 <= 4294967295) {
                                        fmt.Println(v10)
                                    } else {
                                        fmt.Println(fi22, v4)
                                        if ((!(false || false)) || (!true)) {
                                            v38 = 2147483647
                                            v38 = (v38 >> ((487) & 63))
                                            var v39 float64 = (1e10 / 1e10)
                                            _ = v39
                                        }
                                        if (!(32767 < 1000)) {
                                            v38 -= v38
                                        }
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v38
                                }
                                if false {
                                    if (49 >= 128) {
                                        fmt.Println(v10, fi31, fi31)
                                        _ = len(v2)
                                        v40 := GenS4{F0: -328}
                                        _ = v40
                                    }
                                    var v41 uint32 = 2
                                    _ = v41
                                } else {
                                    v42 := GenS4{F0: 0}
                                    v4 += v10
                                    _ = v42
                                }
                                switch (-(fi22)) {
                                case -853:
                                    fmt.Println(v10, v10)
                                    if len(v2) > 0 {
                                        v2[0] = ("" + ((("hello" + (((("\n\t" + ("foo bar" + "hello")) + ("hello" + "\n\t")) + "") + (("hello" + ("" + "hello")) + ("" + "foo bar")))) + "hello") + "a"))
                                    }
                                case 346:
                                    v4 += v4
                                case -9223372036854775808:
                                    switch (^(fi22)) {
                                    case 32768:
                                        v43 := v3[4294967295]
                                        _ = v43
                                    case -9223372036854775808:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case 128:
                                        v4 -= float32((1.5 / (0.0 - 3.14159)))
                                    }
                                default:
                                    fmt.Println(fi22, fi22, fi22)
                                }
                            } else {
                                var v44 string = ("" + ("\n\t" + "\n\t"))
                                var v45 uint32 = 42
                                v45 = (v45 | 1)
                                _ = v44
                                _ = v45
                            }
                            _ = fi31
                        }
                    }
                case "\n\t":
                    delete(v3, 0)
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                fmt.Println(fi22, fi22, v4)
                var v46 uint16 = 65535
                _ = v46
            }
            switch "a" {
            case "foo bar":
                v47 := GenS4{F0: -282}
                _ = v47
            case "":
                v10 = ((1e10 * v4) * ((3.14159 + 3.14159) - -2.25))
            case "hello":
                _ = len(v3)
            default:
                v2 = append(v2, "\n\t")
            }
            v10 = 3.14159
            _ = fi22
        }
    }
    if false {
        v10 = v10
        var v48 int = 42
        _ = v48
    } else {
        fmt.Println(v10)
        _ = len(v2)
    }
    fmt.Println(v10, v4)
    v4 += float32(0.1)
    switch "\n\t" {
    case "":
        switch "foo bar" {
        case "":
            fmt.Println(v4, v10)
        case "foo bar":
            for fi49 := 0; fi49 < 1; fi49++ {
                var v50 int8 = -1
                _ = fi49
                _ = v50
            }
        case "hello":
            var v51 uint = 2
            _ = v51
        }
    default:
        var v52 int = -102
        _ = v52
    }
    for fi53 := 0; fi53 < 3; fi53++ {
        v54 := GenS5{F0: -32768}
        for fi55 := 0; fi55 < 2; fi55++ {
            if false {
                fmt.Println(v4, fi55, v4)
                switch 301 {
                case 32768:
                    if true {
                        v56 := GenS6{F0: -1, F1: true, F2: -128}
                        _ = v56
                    }
                case 32767:
                    v2 = append(v2, "a")
                default:
                    var v57 int64 = 42
                    _ = v57
                }
                switch 32767 {
                case 127:
                    v58 := GenS1{F0: "\n\t"}
                    _ = v58
                case -1:
                    var v59 int8 = 57
                    _ = v59
                case -267:
                    fmt.Println(v4, v10, fi55)
                }
            }
            _ = fi55
        }
        v4 = 1e10
        _ = fi53
        _ = v54
    }
    v4 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
