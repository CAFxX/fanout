package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 bool; F1 string; F2 int8 }
type GenS10 struct { F0 string; F1 uint; F2 int8 }
func genFn0() uint {
    var v5 uint64 = 42
    fmt.Println(v5)
    v5 += (v5 | (^(v5)))
    var v6 bool = (true || (((false && true) && false) && (-89 >= 91)))
    v7 := GenS1{F0: 1}
    _ = v5
    _ = v6
    _ = v7
    return 921
}
func genFn1(p0 int64) int64 {
    var v8 uint8 = 128
    var v9 uint = 255
    var v10 uint8 = v8
    v11 := []int{32768}
    if (!(0 > 0)) {
        if (((0.0 <= 1.5) || (true || (!("hello" >= "")))) && false) {
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: 1}
                v8 = v10
                _ = fi12
                _ = v13
            }
            if (false || false) {
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v9, v8)
                    _ = fi14
                }
                for fi15 := 0; fi15 < 2; fi15++ {
                    fmt.Println(v8, v9, p0)
                    fmt.Println(p0)
                    for fi16 := 0; fi16 < 2; fi16++ {
                        var v17 int32 = -2147483648
                        p0 += -308
                        if true {
                            if false {
                                v18 := GenS2{F0: 326, F1: ((1 >= v9) && true)}
                                if (!false) {
                                    var v19 uint16 = 2
                                    var v20 int32 = (v17 - v17)
                                    var v21 uint8 = uint8(v9)
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                } else {
                                    if ("" > "") {
                                        if false {
                                            v22 := GenS3{F0: p0, F1: 4294967295}
                                            _ = v22
                                        } else {
                                            fmt.Println(p0, v17, v8)
                                            p0 = -429
                                        }
                                        if false {
                                            if true {
                                                _ = len(v11)
                                                p0 = int64(v8)
                                            } else {
                                                var v23 uint = 18446744073709551615
                                                var v24 string = ("" + ("a" + ""))
                                                v25 := GenS4{F0: 98}
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                        }
                                        if false {
                                            var v26 int8 = -87
                                            if (0 <= 18446744073709551615) {
                                                v9 += genFn0()
                                                v8 -= 2
                                                v17 = v17
                                                if (!true) {
                                                    switch ("" + "a") {
                                                    case "\n\t":
                                                        switch ("hello" + ("hello" + ("foo bar" + ""))) {
                                                        case "foo bar":
                                                            v17 = (v17 * v17)
                                                        }
                                                    }
                                                }
                                                switch "\n\t" {
                                                case "":
                                                    var v27 int = fi15
                                                    _ = v27
                                                case "foo bar":
                                                    v28 := GenS5{F0: (v9 - 42), F1: (v9 & v9)}
                                                    _ = v28
                                                case "hello":
                                                    var v29 uint64 = 645
                                                    _ = v29
                                                }
                                            }
                                            v9 -= (v9 * (v9 >> ((12) & 63)))
                                            _ = v26
                                        } else {
                                            v17 -= (^(v17))
                                            switch (fi16 ^ -2147483648) {
                                            case 32768:
                                                if len(v11) > 0 {
                                                    v11[0] = fi15
                                                }
                                            default:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    switch "hello" {
                                                    case "hello":
                                                        var v30 int64 = (p0 / (((p0 << (((v9 << ((722) & 63))) & 63))) | 1))
                                                        _ = v30
                                                    case "foo bar":
                                                        var v31 float64 = 3.14159
                                                        _ = v31
                                                    }
                                                }
                                            }
                                            var v32 int = fi15
                                            _ = v32
                                        }
                                        _ = len(v11)
                                    }
                                    var v33 int16 = -32768
                                    _ = v33
                                }
                                _ = v18
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                    _ = fi15
                }
                v34 := GenS6{F0: 4294967295, F1: (^(v9))}
                _ = v34
            }
            if (1 > (v9 - (v9 | v9))) {
                switch "hello" {
                case "hello":
                    fmt.Println(v8, v8)
                }
            } else {
                v8 = v10
            }
        } else {
            switch 0 {
            case 42:
                if true {
                    v10 = (v8 & (v8 & (v10 ^ (v10 - (^(v10))))))
                }
            case 753:
                fmt.Println(v8, v9, p0)
            default:
                switch 42 {
                case 0:
                    var v35 float64 = 1.5
                    _ = v35
                }
            }
            var v36 int64 = p0
            fmt.Println(v10, v36)
            _ = v36
        }
        if true {
            var v37 int64 = p0
            v11 = append(v11, -370)
            v38 := GenS4{F0: 4294967295}
            _ = v37
            _ = v38
        } else {
            for fi39 := 0; fi39 < 17; fi39++ {
                var v40 uint8 = (^(v8))
                for fi41 := 0; fi41 < 10; fi41++ {
                    v42 := GenS3{F0: (^(p0)), F1: uint64(v9)}
                    v10 = v10
                    var v43 uint = 18446744073709551615
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = fi39
                _ = v40
            }
            v9 = 515
        }
    } else {
        v8 += v8
        if false {
            if (-439 != 1) {
                _ = len(v11)
                v10 -= 48
            } else {
                var v44 float32 = 0.1
                _ = v44
            }
        }
        if len(v11) > 0 {
            v11[0] = 1
        }
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]bool)
    v3 := []string{("foo bar" + ("a" + "")), "a", "hello", (("" + "foo bar") + ("\n\t" + (("foo bar" + ("\n\t" + "foo bar")) + ("hello" + "a"))))}
    var v4 int16 = -32768
    delete(v2, "a")
    v4 = (v4 | -547)
    for fi45 := 0; fi45 < 11; fi45++ {
        switch "foo bar" {
        case "":
            fmt.Println(v1)
        case "foo bar":
            v4 = v4
        case "hello":
            switch (^(fi45)) {
            case 42:
                var v46 int16 = v4
                _ = v46
            }
        default:
            v47 := GenS4{F0: 18446744073709551615}
            _ = v47
        }
        for fi48 := 0; fi48 < 10; fi48++ {
            if (true && (!true)) {
                if false {
                    v49 := GenS7{F0: 245}
                    if false {
                        v50 := GenS2{F0: 1, F1: ((3.14159 / float64((-2.25 / (1.5 / 1e10)))) != 0.1)}
                        _ = v50
                    } else {
                        var v51 int8 = -1
                        v51 -= v51
                        _ = v51
                    }
                    v4 = v4
                    v1 -= 255
                    var v52 uint = 414
                    _ = v49
                    _ = v52
                } else {
                    switch "a" {
                    case "":
                        v4 -= v4
                    case "a":
                        v4 += (v4 & v4)
                    case "hello":
                        v53 := GenS8{F0: 127, F1: "\n\t"}
                        v54 := GenS9{F0: true, F1: "hello", F2: 81}
                        var v55 uint8 = 0
                        var v56 int32 = 42
                        _ = v53
                        _ = v54
                        _ = v55
                        _ = v56
                    }
                }
                fmt.Println(v4, v1, v1)
            }
            _ = fi48
        }
        _ = fi45
    }
    v1 = v1
    v57 := GenS10{F0: (("a" + ("hello" + "a")) + "a"), F1: v1, F2: -128}
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    _ = len(v2)
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
}
