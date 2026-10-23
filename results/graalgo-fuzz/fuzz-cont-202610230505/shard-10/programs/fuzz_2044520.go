package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 bool; F1 bool }
func genFn0() uint8 {
    var v5 uint32 = 2147483648
    v5 = (v5 % ((65535) | 1))
    var v6 uint8 = 0
    for fi7 := 0; fi7 < 13; fi7++ {
        if (!(true || false)) {
            v8 := GenS1{F0: "\n\t", F1: 809, F2: 100}
            _ = v8
        } else {
            if ((v6 | 128) != v6) {
                var v9 uint8 = (^(v6))
                _ = v9
            } else {
                v10 := []int{(fi7 + fi7)}
                v11 := GenS2{F0: ("a" + ("" + "foo bar"))}
                if false {
                    v12 := GenS3{F0: 679, F1: v5}
                    fmt.Println(v6)
                    if (1 > 42) {
                        if (-35 == 42) {
                            v5 += v5
                        } else {
                            for fi13 := 0; fi13 < 3; fi13++ {
                                v14 := GenS3{F0: 0, F1: 747}
                                v6 -= v6
                                switch 128 {
                                case 128:
                                    fmt.Println(fi13, v6, v6)
                                case -610:
                                    v6 = 91
                                }
                                if true {
                                    fmt.Println(v6)
                                }
                                _ = fi13
                                _ = v14
                            }
                            switch "" {
                            case "foo bar":
                                for fi15 := 0; fi15 < 6; fi15++ {
                                    v6 += v6
                                    v6 -= (v6 / ((237) | 1))
                                    _ = fi15
                                }
                            case "a":
                                fmt.Println(v6)
                            default:
                                for fi16 := 0; fi16 < 13; fi16++ {
                                    var v17 bool = ((127 >= -128) && (-1 != 0))
                                    _ = fi16
                                    _ = v17
                                }
                            }
                        }
                        var v18 int16 = 1000
                        _ = v18
                    } else {
                        var v19 int64 = 0
                        _ = v19
                    }
                    _ = v12
                } else {
                    v20 := GenS2{F0: ("hello" + "foo bar")}
                    _ = v20
                }
                _ = v10
                _ = v11
            }
            fmt.Println(v5, v6, v5)
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return v6
}

func main() {
    v1 := []uint32{65535, 1}
    var v2 float32 = (((1.5 + (-2.25 + (3.14159 - float32(3.14159)))) * -2.25) / 1e10)
    var v3 float32 = v2
    v4 := []int{804}
    var v21 int16 = 42
    switch "foo bar" {
    case "a":
        v3 -= v2
    default:
        v3 -= v3
    }
    if false {
        var v22 int64 = 42
        _ = v22
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch -1 {
    case 9223372036854775807:
        var v23 uint = 65535
        _ = v23
    case 2147483647:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -1:
        v21 -= (v21 / ((v21) | 1))
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        if (!((^(v21)) > -32768)) {
            if false {
                for fi25 := 0; fi25 < 8; fi25++ {
                    v26 := GenS4{F0: 1}
                    switch "" {
                    case "\n\t":
                        v1 = append(v1, 65535)
                    case "foo bar":
                        v21 += -1
                    case "hello":
                        if (1 > 924) {
                            v2 = (0.0 + v2)
                        } else {
                            var v27 uint8 = 0
                            if false {
                                v2 = v3
                                v28 := GenS4{F0: 946}
                                v29 := GenS2{F0: "a"}
                                _ = v28
                                _ = v29
                            }
                            var v30 uint32 = 2147483648
                            _ = v27
                            _ = v30
                        }
                    }
                    if true {
                        v4 = append(v4, 368)
                        var v31 uint = 255
                        if ((false && (true && (3.14159 <= 0.0))) && true) {
                            fmt.Println(fi25, fi25)
                            fmt.Println(v21, fi24, fi25)
                            if len(v4) > 0 {
                                v4[0] = (fi25 << (((^(v31))) & 63))
                            }
                        } else {
                            _ = len(v4)
                            fmt.Println(fi24, fi24, v21)
                            v4 = append(v4, 2147483647)
                            v32 := GenS3{F0: 1, F1: 4294967295}
                            switch ("a" + "\n\t") {
                            case "hello":
                                switch fi25 {
                                case -805:
                                    v33 := GenS5{F0: 86, F1: 315}
                                    _ = v33
                                case 1:
                                    v3 = v3
                                case -2147483648:
                                    v31 = (v31 & 4294967295)
                                }
                            case "\n\t":
                                v2 = v2
                            default:
                                var v34 float64 = 0.1
                                if false {
                                    var v35 uint8 = 0
                                    if ((v21 + (v21 + 1)) >= (v21 ^ (v21 & v21))) {
                                        var v36 uint = 490
                                        if (0 >= 1) {
                                            fmt.Println(v31)
                                            if false {
                                                v37 := GenS6{F0: 208}
                                                switch "" {
                                                case "\n\t":
                                                    var v38 uint8 = genFn0()
                                                    _ = v38
                                                default:
                                                    v39 := GenS7{F0: "a", F1: "foo bar"}
                                                    _ = v39
                                                }
                                                _ = v37
                                            }
                                            var v40 uint8 = genFn0()
                                            _ = v40
                                        }
                                        _ = v36
                                    }
                                    _ = v35
                                }
                                fmt.Println(v31, v2)
                                _ = v34
                            }
                            _ = v32
                        }
                        _ = v31
                    } else {
                        _ = len(v1)
                    }
                    _ = fi25
                    _ = v26
                }
            } else {
                switch (fi24 / ((327) | 1)) {
                case 32768:
                    if false {
                        var v41 uint = 65535
                        switch "\n\t" {
                        case "foo bar":
                            v1 = append(v1, uint32(v41))
                        }
                        fmt.Println(v3, v21, v2)
                        _ = v41
                    } else {
                        if (109 < 127) {
                            var v42 int = fi24
                            v3 = ((v2 + v2) + 3.14159)
                            if true {
                                var v43 int16 = (^(v21))
                                _ = v43
                            }
                            _ = v42
                        } else {
                            v44 := GenS5{F0: 127, F1: 4294967295}
                            _ = v44
                        }
                    }
                case 0:
                    var v45 string = ("" + "foo bar")
                    _ = v45
                default:
                    if ((true && true) && (42 <= 2)) {
                        v46 := GenS8{F0: (-629 > 42), F1: false}
                        _ = v46
                    } else {
                        v47 := GenS7{F0: ("a" + ("a" + ("foo bar" + "hello"))), F1: (((("" + "foo bar") + ("hello" + "hello")) + ("hello" + ("hello" + (((("foo bar" + ("hello" + "a")) + ("" + "")) + ("\n\t" + "foo bar")) + "a")))) + "a")}
                        var v48 string = ("a" + "a")
                        _ = v47
                        _ = v48
                    }
                }
                v4 = append(v4, (fi24 >> ((65535) & 63)))
            }
            var v49 int64 = -9223372036854775808
            var v50 bool = false
            _ = v49
            _ = v50
        } else {
            v51 := GenS4{F0: 42}
            v21 += v21
            _ = v51
        }
        if true {
            var v52 string = ("a" + ((("\n\t" + ((("foo bar" + "hello") + "foo bar") + ("a" + "foo bar"))) + ("hello" + "\n\t")) + ""))
            var v53 int = (fi24 + (fi24 * -9223372036854775808))
            for fi54 := 0; fi54 < 20; fi54++ {
                if true {
                    if (fi54 >= (fi24 & (v53 >> ((4294967295) & 63)))) {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        v3 = float32((1e10 - float64(v3)))
                        switch v52 {
                        case "\n\t":
                            v53 = v53
                        default:
                            v21 -= (v21 * v21)
                        }
                    }
                    switch v52 {
                    case "":
                        if false {
                            var v55 float64 = (1e10 / 1.0)
                            _ = v55
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 916
                            }
                            v56 := GenS7{F0: "foo bar", F1: "foo bar"}
                            v21 += v21
                            _ = v56
                        }
                    case "hello":
                        fmt.Println(v53, v53, v21)
                    default:
                        var v57 uint = 18446744073709551615
                        _ = v57
                    }
                } else {
                    switch v52 {
                    case "foo bar":
                        var v58 uint = 811
                        _ = v58
                    case "":
                        var v59 bool = false
                        _ = v59
                    case "hello":
                        fmt.Println(v52, v3, v2)
                    }
                    var v60 uint = 4294967295
                    _ = v60
                }
                _ = fi54
            }
            _ = v52
            _ = v53
        }
        if len(v4) > 0 {
            v4[0] = fi24
        }
        _ = fi24
    }
    v21 -= v21
    switch "" {
    case "\n\t":
        var v61 int = -1
        _ = v61
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v62 int64 = 82
    switch "a" {
    case "hello":
        var v63 int16 = v21
        _ = v63
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v62
}
