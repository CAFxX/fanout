package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 bool; F2 uint16 }
type GenS5 struct { F0 int; F1 string }
func genFn0(p0 bool) string {
    var v7 uint = 4294967295
    if p0 {
        var v8 int = 9223372036854775807
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := []int{(v8 & fi9)}
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        p0 = ((9223372036854775807 != 1) || ((!p0) || (-86 <= 755)))
        fmt.Println(v7, p0)
        switch ("" + (((((("hello" + (((("" + ("" + "foo bar")) + ("\n\t" + "\n\t")) + "a") + "hello")) + ("\n\t" + "\n\t")) + ("foo bar" + "a")) + "") + "hello") + "\n\t")) {
        case "foo bar":
            fmt.Println(p0, p0)
        case "a":
            var v11 int16 = -690
            _ = v11
        default:
            fmt.Println(p0)
        }
    }
    var v12 float64 = -2.25
    _ = p0
    _ = v12
    _ = v7
    return ("a" + "\n\t")
}
func genFn1(p0 bool, p1 uint16, p2 uint16) string {
    var v13 int8 = 115
    var v14 string = ""
    if ((!true) && p0) {
        v15 := GenS1{F0: 656, F1: 1}
        v16 := []int{}
        fmt.Println(p2, v14)
        _ = v15
        _ = v16
    } else {
        if (32767 == (p2 - 0)) {
            v14 = v14
            var v17 uint32 = 354
            v18 := GenS1{F0: 1, F1: p2}
            _ = v17
            _ = v18
        } else {
            var v19 int64 = 0
            fmt.Println(p1, p2)
            _ = v19
        }
        var v20 int8 = -21
        v21 := []int{0}
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: -1, F1: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v22
    return ((v14 + v14) + v14)
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 18446744073709551615
    var v3 uint = v2
    v4 := make(map[int]string)
    v5 := []int{127, 128, 4294967295}
    var v6 int = 128
    var v23 string = ((("a" + ("a" + ((("hello" + "\n\t") + ("a" + ("a" + "hello"))) + "\n\t"))) + (("hello" + ("hello" + ("a" + ((("a" + ("hello" + "a")) + ("hello" + "foo bar")) + "a")))) + (("foo bar" + ("hello" + ((("\n\t" + ("a" + "foo bar")) + "") + "\n\t"))) + ("foo bar" + "hello")))) + "foo bar")
    v6 -= v6
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            if len(v5) > 0 {
                v5[0] = fi24
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v6 -= fi25
            _ = fi25
        }
        v26 := GenS1{F0: -105, F1: 42}
        for fi27 := 0; fi27 < 18; fi27++ {
            switch (fi27 + -9223372036854775808) {
            case -2147483648:
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v28 uint16 = 1
                    _ = v28
                }
            case 1:
                v29 := GenS1{F0: 0, F1: 42}
                _ = v29
            case 0:
                fmt.Println(v23, v23)
            default:
                v3 = (^(v3))
            }
            fmt.Println(v3)
            v2 = (v3 + v2)
            _ = fi27
        }
        _ = fi24
        _ = v26
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        for fi31 := 0; fi31 < 16; fi31++ {
            if (false && false) {
                v2 = (v2 | v3)
                switch (v23 + ((v23 + "foo bar") + v23)) {
                case "hello":
                    v3 -= (v2 + (^(v2)))
                case "\n\t":
                    if (255 != 255) {
                        v1 = v1
                        if (9223372036854775807 == -9223372036854775808) {
                            v3 += 65535
                            var v32 int = (v6 << (((v3 - v3)) & 63))
                            _ = v32
                        }
                        fmt.Println(v3, v23)
                        v3 = v3
                    }
                default:
                    v33 := GenS1{F0: 2147483647, F1: 20}
                    _ = v33
                }
                if false {
                    v34 := GenS1{F0: 2147483647, F1: uint16(v3)}
                    if true {
                        v3 -= v2
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v34
                } else {
                    var v35 int = fi31
                    if true {
                        v1 -= 1.5
                        switch (v6 ^ 331) {
                        case 339:
                            v36 := GenS2{F0: 1000, F1: (!true), F2: 42}
                            _ = v36
                        case -1:
                            v6 = 130
                        case 32768:
                            if (!false) {
                                v1 = ((v1 * float64(float32((0.1 + v1)))) + (0.1 / (((v1 / v1) - v1) * ((v1 + -2.25) - v1))))
                                if (false && false) {
                                    v6 += (fi31 + 0)
                                    _ = len(v5)
                                }
                                v3 = v2
                                _ = len(v5)
                            }
                        }
                        v37 := GenS3{F0: (^(fi30))}
                        _ = v37
                    } else {
                        _ = len(v4)
                        v23 = "\n\t"
                    }
                    v1 = v1
                    _ = v35
                }
            }
            if true {
                v6 -= fi31
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if len(v5) > 0 {
                    v5[0] = fi31
                }
            } else {
                fmt.Println(v3, v1, fi30)
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
                fmt.Println(fi30)
            }
            if false {
                delete(v4, (v6 * 259))
            } else {
                fmt.Println(v3, fi30, v2)
                fmt.Println(fi30)
            }
            _ = fi31
        }
        _ = fi30
    }
    v3 = 623
    switch v6 {
    case 1:
        for fi38 := 0; fi38 < 10; fi38++ {
            v39 := GenS4{F0: 40, F1: (111 != 42), F2: 2}
            for fi40 := 0; fi40 < 14; fi40++ {
                v41 := GenS5{F0: 691, F1: v23}
                var v42 int8 = 42
                switch v23 {
                case "hello":
                    switch ("" + "\n\t") {
                    case "\n\t":
                        var v43 float64 = float64(((0.1 + (-2.25 + 1.5)) / (3.14159 + 1e10)))
                        _ = v43
                    case "hello":
                        fmt.Println(v6)
                    default:
                        v44 := GenS4{F0: v42, F1: true, F2: 858}
                        _ = v44
                    }
                }
                _ = fi40
                _ = v41
                _ = v42
            }
            for fi45 := 0; fi45 < 11; fi45++ {
                var v46 float64 = (((v1 / (v1 - ((v1 + 3.14159) - v1))) + float64(-2.25)) + (-2.25 - 0.0))
                switch -9223372036854775808 {
                case 2147483647:
                    if false {
                        v5 = append(v5, v6)
                        var v47 uint32 = 42
                        _ = v47
                    }
                case -9223372036854775808:
                    _ = len(v5)
                case 0:
                    _ = len(v4)
                }
                v2 = v2
                _ = fi45
                _ = v46
            }
            _ = fi38
            _ = v39
        }
    case 42:
        _ = len(v5)
    }
    v3 = v2
    v3 = v3
    v48 := GenS5{F0: 9223372036854775807, F1: "\n\t"}
    v49 := GenS3{F0: v6}
    _ = len(v4)
    var v50 bool = true
    v51 := GenS3{F0: 184}
    switch v6 {
    case 2147483647:
        v50 = (!(((3.14159 == 3.14159) || v50) || v50))
    case -143:
        fmt.Println(v3)
    default:
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    }
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    v52 := v4[(v6 ^ v6)]
    v53 := GenS2{F0: -32768, F1: true, F2: 4294967295}
    var v54 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
