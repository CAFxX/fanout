package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int32; F1 int8; F2 int }
type GenS7 struct { F0 int64 }

func main() {
    var v1 int32 = -1
    v2 := make(map[string]int)
    v3 := []int{-724, 9223372036854775807, 0, 310}
    v4 := GenS1{F0: 1, F1: -294, F2: 255}
    _ = len(v3)
    v1 = 42
    var v5 uint64 = 133
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5, v5)
    if (!(!true)) {
        var v6 uint8 = 2
        v1 -= v1
        _ = v6
    }
    switch 32767 {
    case 32768:
        switch (((("a" + "") + "hello") + "foo bar") + "") {
        case "foo bar":
            v7 := GenS2{F0: -1, F1: 0}
            _ = v7
        case "hello":
            var v8 uint8 = 2
            _ = v8
        case "\n\t":
            v2[("a" + "")] = 2147483647
        default:
            if (1 > 42) {
                for fi9 := 0; fi9 < 12; fi9++ {
                    v10 := GenS1{F0: 127, F1: fi9, F2: 42}
                    if (!true) {
                        fmt.Println(v1)
                        v5 = (^(v5))
                    } else {
                        for fi11 := 0; fi11 < 16; fi11++ {
                            var v12 int8 = 100
                            _ = fi11
                            _ = v12
                        }
                        var v13 float32 = 1e10
                        v14 := GenS3{F0: 1, F1: 775}
                        _ = v13
                        _ = v14
                    }
                    v15 := GenS2{F0: -1, F1: 1000}
                    _ = fi9
                    _ = v10
                    _ = v15
                }
                v2[("" + (("foo bar" + ("a" + ("\n\t" + ("\n\t" + "")))) + "foo bar"))] = 128
                fmt.Println(v5)
            }
        }
    }
    v5 = (v5 - (v5 & v5))
    var v16 uint8 = 42
    if false {
        if false {
            v16 = v16
            if false {
                var v17 int16 = 42
                var v18 uint64 = 42
                var v19 uint32 = 4294967295
                _ = v17
                _ = v18
                _ = v19
            } else {
                var v20 int32 = (v1 ^ v1)
                var v21 int16 = 42
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 10; fi22++ {
                _ = len(v2)
                var v23 string = "hello"
                _ = fi22
                _ = v23
            }
        } else {
            var v24 int32 = 42
            var v25 uint8 = v16
            _ = v24
            _ = v25
        }
        delete(v2, "\n\t")
        fmt.Println(v1)
    } else {
        for fi26 := 0; fi26 < 7; fi26++ {
            v16 = v16
            _ = fi26
        }
        if (65535 < uint32(v16)) {
            v27 := GenS4{F0: v5}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 3; fi28++ {
                if (v5 == v5) {
                    switch ("" + ("\n\t" + (((("hello" + "\n\t") + ("hello" + "\n\t")) + "foo bar") + ("\n\t" + "foo bar")))) {
                    case "hello":
                        for fi29 := 0; fi29 < 11; fi29++ {
                            if true {
                                if (true || true) {
                                    switch "hello" {
                                    case "hello":
                                        var v30 string = ("\n\t" + (("foo bar" + "a") + "foo bar"))
                                        _ = v30
                                    case "foo bar":
                                        v31 := GenS3{F0: 920, F1: (v5 | (^(v5)))}
                                        _ = v31
                                    case "a":
                                        var v32 float64 = (-2.25 + (1.5 / -2.25))
                                        _ = v32
                                    default:
                                        v33 := GenS1{F0: 128, F1: 555, F2: v16}
                                        _ = v33
                                    }
                                }
                                v16 = 216
                                v34 := GenS1{F0: v16, F1: fi28, F2: (v16 * (v16 ^ (v16 % ((v16) | 1))))}
                                _ = v34
                            } else {
                                fmt.Println(v16)
                                v35 := GenS5{F0: (v5 << ((389) & 63)), F1: ("\n\t" + "a")}
                                var v36 uint64 = v5
                                v3 = append(v3, fi29)
                                _ = v35
                                _ = v36
                            }
                            _ = fi29
                        }
                    case "foo bar":
                        for fi37 := 0; fi37 < 11; fi37++ {
                            fmt.Println(fi28, v1)
                            if true {
                                v38 := GenS2{F0: 100, F1: 32767}
                                v3 = append(v3, (fi28 + fi37))
                                var v39 string = ""
                                _ = v38
                                _ = v39
                            } else {
                                if ((false && ((v1 | v1) == 105)) || ((42 >= 1000) && true)) {
                                    v1 = -726
                                    fmt.Println(fi28, v5, v5)
                                    var v40 int8 = 1
                                    _ = v40
                                } else {
                                    v16 += v16
                                    var v41 int16 = 919
                                    v42 := GenS6{F0: v1, F1: -1, F2: fi28}
                                    _ = v41
                                    _ = v42
                                }
                                if true {
                                    var v43 uint64 = (v5 & (v5 ^ (v5 | (^(v5)))))
                                    v1 -= (^(v1))
                                    v3 = append(v3, fi37)
                                    _ = v43
                                } else {
                                    if true {
                                        if (false || ((^(v1)) > (v1 ^ v1))) {
                                            v44 := GenS7{F0: 1}
                                            v45 := GenS5{F0: 18446744073709551615, F1: "\n\t"}
                                            v16 = (v16 + 255)
                                            _ = v44
                                            _ = v45
                                        }
                                    } else {
                                        if (!((fi28 << ((766) & 63)) > fi28)) {
                                            v5 -= v5
                                            fmt.Println(v5, v16, v5)
                                            fmt.Println(fi37, fi28, fi37)
                                        } else {
                                            v46 := GenS5{F0: (v5 >> ((0) & 63)), F1: "\n\t"}
                                            if true {
                                                fmt.Println(v1, fi37, fi37)
                                            } else {
                                                switch "hello" {
                                                case "foo bar":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "a":
                                                    v47 := v2["foo bar"]
                                                    _ = v47
                                                }
                                                if (true && false) {
                                                    if true {
                                                        fmt.Println(v16, fi37)
                                                    }
                                                }
                                                v2["foo bar"] = fi28
                                                fmt.Println(v5)
                                            }
                                            _ = v46
                                        }
                                        v48 := GenS5{F0: v5, F1: ("" + "foo bar")}
                                        v1 += -2147483648
                                        _ = v48
                                    }
                                    switch "" {
                                    case "hello":
                                        v49 := GenS3{F0: -1, F1: v5}
                                        _ = v49
                                    case "foo bar":
                                        if false {
                                            var v50 int64 = 0
                                            if true {
                                                v16 += 153
                                            } else {
                                                delete(v2, "hello")
                                            }
                                            v16 = v16
                                            _ = v50
                                        }
                                    default:
                                        var v51 int64 = 9223372036854775807
                                        _ = v51
                                    }
                                }
                            }
                            fmt.Println(v1, v1)
                            _ = fi37
                        }
                    case "":
                        _ = len(v3)
                    }
                    if true {
                        var v52 uint16 = 921
                        v16 += 251
                        _ = v52
                    }
                }
                v5 = (v5 % ((v5) | 1))
                _ = fi28
            }
        }
    }
    var v53 uint32 = 2147483648
    for fi54 := 0; fi54 < 5; fi54++ {
        fmt.Println(fi54, v16, v16)
        _ = fi54
    }
    switch "" {
    case "hello":
        v16 = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
