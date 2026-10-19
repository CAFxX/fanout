package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint16) uint16 {
    var v5 uint32 = 42
    var v6 int8 = 0
    for fi7 := 0; fi7 < 10; fi7++ {
        if false {
            var v8 uint64 = 42
            v8 = v8
            for fi9 := 0; fi9 < 12; fi9++ {
                switch "foo bar" {
                case "hello":
                    v6 += v6
                }
                v10 := GenS1{F0: (fi7 / ((32768) | 1)), F1: "hello"}
                v5 = 1
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        _ = fi7
    }
    v11 := []int{-848}
    _ = p0
    _ = v11
    _ = v5
    _ = v6
    return p0
}
func genFn1() int {
    var v12 int8 = 0
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(v12, fi13)
        _ = fi13
    }
    _ = v12
    return 127
}

func main() {
    v1 := []int{32768, -358}
    var v2 uint16 = 255
    var v3 int32 = -769
    var v4 float32 = 0.1
    v1 = append(v1, -1)
    if false {
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(fi14)
            v3 += v3
            _ = fi14
        }
        if true {
            v3 += v3
            for fi15 := 0; fi15 < 1; fi15++ {
                if true {
                    v16 := GenS2{F0: 0, F1: "a", F2: ""}
                    if (((fi15 & (-(fi15))) < -405) && false) {
                        for fi17 := 0; fi17 < 12; fi17++ {
                            v3 = v3
                            _ = fi17
                        }
                        var v18 int64 = 42
                        for fi19 := 0; fi19 < 16; fi19++ {
                            var v20 string = ((("" + ("" + "a")) + "foo bar") + ("a" + "a"))
                            v3 = (v3 - v3)
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    }
                    _ = v16
                }
                v21 := GenS2{F0: 507, F1: ("foo bar" + "a"), F2: "a"}
                v2 += v2
                _ = fi15
                _ = v21
            }
            var v22 int8 = 42
            _ = v22
        } else {
            fmt.Println(v2, v2, v3)
        }
    } else {
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int16 = 1
            _ = fi23
            _ = v24
        }
        var v25 int64 = -94
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 uint8 = 42
        var v28 uint32 = 0
        v29 := GenS3{F0: -1, F1: genFn1()}
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println(v3, v4, v4)
    if (!true) {
        v4 -= v4
    } else {
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := GenS1{F0: (^(fi30)), F1: ("\n\t" + "a")}
            _ = len(v1)
            var v32 string = ""
            _ = fi30
            _ = v31
            _ = v32
        }
        switch ("foo bar" + "") {
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "":
            var v33 string = "hello"
            _ = v33
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = 91
            }
        default:
            if len(v1) > 0 {
                v1[0] = 32767
            }
        }
    }
    fmt.Println(v2, v4)
    v3 = (^(v3))
    if true {
        v4 += v4
    }
    var v34 string = ("a" + ("\n\t" + ("foo bar" + (("foo bar" + "") + "foo bar"))))
    var v35 uint64 = 42
    switch v34 {
    case "foo bar":
        for fi36 := 0; fi36 < 7; fi36++ {
            if false {
                v2 = (v2 * (^(v2)))
            }
            if (((v34 + ((v34 + v34) + "hello")) > (("\n\t" + v34) + ((v34 + (v34 + ("\n\t" + ("hello" + "foo bar")))) + v34))) || false) {
                var v37 uint32 = 2147483648
                var v38 float64 = -2.25
                v39 := GenS4{F0: fi36, F1: (v35 + (v35 & 878))}
                _ = v37
                _ = v38
                _ = v39
            } else {
                switch (fi36 | -149) {
                case 42:
                    for fi40 := 0; fi40 < 7; fi40++ {
                        if true {
                            switch genFn1() {
                            case 2147483647:
                                var v41 uint64 = (^(v35))
                                _ = v41
                            case 9223372036854775807:
                                var v42 uint64 = v35
                                _ = v42
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = fi36
                        }
                        v1 = append(v1, (fi36 * (fi36 * fi36)))
                        var v43 string = v34
                        _ = fi40
                        _ = v43
                    }
                case -826:
                    if (!false) {
                        if (true && false) {
                            var v44 uint8 = 128
                            _ = v44
                        }
                    }
                case 4294967295:
                    fmt.Println(v35, v35, fi36)
                }
                var v45 float32 = ((0.1 / v4) + (((((1e10 - v4) * v4) - v4) - float32(((3.14159 * 3.14159) + 0.1))) * v4))
                _ = v45
            }
            if (true && true) {
                fmt.Println(v4)
                v4 = 0.0
            }
            _ = fi36
        }
    case "\n\t":
        for fi46 := 0; fi46 < 10; fi46++ {
            if true {
                if false {
                    fmt.Println(v2)
                    var v47 bool = (true || false)
                    if (fi46 != fi46) {
                        v48 := GenS5{F0: genFn1(), F1: (fi46 + (^(fi46)))}
                        v4 = v4
                        for fi49 := 0; fi49 < 1; fi49++ {
                            switch fi46 {
                            case -2147483648:
                                switch "\n\t" {
                                case "":
                                    v35 = (v35 << ((uint(v2)) & 63))
                                default:
                                    v50 := GenS6{F0: -35}
                                    _ = v50
                                }
                            case 0:
                                fmt.Println(v4, v47, v4)
                            case 447:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v2 = (^(v2))
                            }
                            _ = fi49
                        }
                        _ = v48
                    } else {
                        fmt.Println(fi46, v35, fi46)
                        var v51 bool = (v35 < v35)
                        if v51 {
                            switch v34 {
                            case "hello":
                                v3 = v3
                            default:
                                if (!(0 == 42)) {
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v35, v51, v51)
                                    case "hello":
                                        var v52 float32 = v4
                                        _ = v52
                                    case "\n\t":
                                        if (0 <= -9223372036854775808) {
                                            var v53 string = (v34 + v34)
                                            for fi54 := 0; fi54 < 9; fi54++ {
                                                var v55 uint64 = v35
                                                v56 := GenS7{F0: 0, F1: "hello"}
                                                if v51 {
                                                    v47 = (2 >= 128)
                                                    _ = len(v1)
                                                    var v57 bool = (128 >= 127)
                                                    _ = v57
                                                }
                                                var v58 int = fi54
                                                var v59 int = (v58 * v58)
                                                var v60 bool = (-150 != 517)
                                                _ = fi54
                                                _ = v55
                                                _ = v56
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            }
                                            for fi61 := 0; fi61 < 13; fi61++ {
                                                if len(v1) > 0 {
                                                    v1[0] = (fi61 * 0)
                                                }
                                                _ = fi61
                                            }
                                            _ = v53
                                        }
                                    }
                                }
                            }
                            fmt.Println(v35)
                            v4 += -2.25
                        }
                        _ = v51
                    }
                    _ = v47
                } else {
                    v35 = (v35 % (((v35 << ((2) & 63))) | 1))
                    v1 = append(v1, fi46)
                    var v62 int = (fi46 ^ (fi46 % (((fi46 / ((fi46) | 1))) | 1)))
                    _ = v62
                }
                fmt.Println(v3, fi46, v4)
            } else {
                fmt.Println(fi46)
            }
            if true {
                var v63 int = fi46
                _ = len(v1)
                var v64 int64 = 0
                _ = v63
                _ = v64
            } else {
                v1 = append(v1, (-(fi46)))
                v4 = (1.5 / (v4 + 3.14159))
                fmt.Println(v35)
            }
            _ = fi46
        }
    case "hello":
        if (42 < -1) {
            v34 = v34
            if false {
                for fi65 := 0; fi65 < 6; fi65++ {
                    for fi66 := 0; fi66 < 9; fi66++ {
                        if len(v1) > 0 {
                            v1[0] = (fi66 - fi65)
                        }
                        v4 = v4
                        if ((("hello" + v34) + "") >= v34) {
                            v2 = (v2 * (^(v2)))
                            var v67 uint16 = genFn0(v2)
                            v34 = v34
                            v1 = append(v1, genFn1())
                            switch "a" {
                            case "foo bar":
                                if len(v1) > 0 {
                                    v1[0] = 127
                                }
                            case "\n\t":
                                if (!((!true) || true)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                fmt.Println(v34)
                            default:
                                v68 := GenS8{F0: 121, F1: 100}
                                _ = v68
                            }
                            _ = v67
                        } else {
                            var v69 float32 = v4
                            _ = v69
                        }
                        _ = fi66
                    }
                    _ = fi65
                }
                var v70 int16 = 42
                v2 += genFn0(v2)
                _ = v70
            }
        }
    }
    var v71 int64 = 0
    var v72 float32 = 0.0
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if len(v1) > 0 {
        v1[0] = genFn1()
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v71
    _ = v72
}
