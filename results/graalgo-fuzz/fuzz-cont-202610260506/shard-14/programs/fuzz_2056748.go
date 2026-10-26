package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string }

func main() {
    var v1 bool = true
    v2 := []uint32{0, 0, 2147483647}
    var v3 float32 = 3.14159
    v4 := []string{"\n\t", "\n\t"}
    v5 := []uint32{42, 2147483647, 4294967295}
    var v6 int8 = 127
    fmt.Println(v6, v6)
    fmt.Println(v6, v1, v1)
    var v7 bool = v1
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: v1}
        var v10 float32 = v3
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 = v7
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    switch "" {
    case "foo bar":
        v6 -= -128
    case "a":
        fmt.Println(v7)
    default:
        switch "foo bar" {
        case "a":
            switch (((("\n\t" + "\n\t") + "\n\t") + ("foo bar" + "foo bar")) + "hello") {
            case "foo bar":
                v11 := GenS2{F0: 763, F1: 799, F2: (((-1 <= 42) || true) && v1)}
                _ = v11
            case "\n\t":
                if (!false) {
                    fmt.Println(v3)
                }
            case "a":
                v3 = -2.25
            default:
                v12 := GenS3{F0: 127}
                _ = v12
            }
        case "\n\t":
            if true {
                for fi13 := 0; fi13 < 15; fi13++ {
                    if v1 {
                        if (v7 && false) {
                            fmt.Println(v1)
                            if v7 {
                                var v14 uint8 = 127
                                _ = v14
                            }
                        } else {
                            var v15 uint64 = 1
                            v16 := GenS2{F0: 497, F1: 0, F2: ("hello" != "foo bar")}
                            switch (fi13 * fi13) {
                            case 954:
                                if (!(0.0 >= 1.5)) {
                                    var v17 string = "\n\t"
                                    var v18 float64 = 1.5
                                    if true {
                                        switch (("foo bar" + ("a" + v17)) + ((v17 + ("\n\t" + (v17 + (v17 + v17)))) + v17)) {
                                        case "a":
                                            if v7 {
                                                _ = len(v4)
                                                fmt.Println(v17, v18)
                                            } else {
                                                v3 -= v3
                                                v19 := GenS2{F0: 32767, F1: -1, F2: ((fi13 | fi13) >= (-(fi13)))}
                                                v6 -= -128
                                                _ = v19
                                            }
                                        }
                                        _ = len(v2)
                                        for fi20 := 0; fi20 < 2; fi20++ {
                                            v21 := GenS1{F0: true}
                                            if false {
                                                if true {
                                                    var v22 int = fi13
                                                    _ = v22
                                                }
                                                var v23 uint = 18446744073709551615
                                                fmt.Println(v3, v1)
                                                _ = v23
                                            } else {
                                                v17 = v17
                                                v24 := GenS3{F0: 128}
                                                _ = v24
                                            }
                                            _ = fi20
                                            _ = v21
                                        }
                                    }
                                    _ = v17
                                    _ = v18
                                } else {
                                    for fi25 := 0; fi25 < 14; fi25++ {
                                        fmt.Println(v1, v1)
                                        _ = fi25
                                    }
                                    var v26 uint16 = 32767
                                    _ = v26
                                }
                            default:
                                switch ((("a" + "foo bar") + "hello") + ("hello" + "a")) {
                                case "a":
                                    switch fi13 {
                                    case 32767:
                                        for fi27 := 0; fi27 < 13; fi27++ {
                                            var v28 string = ("foo bar" + "\n\t")
                                            fmt.Println(fi13)
                                            v29 := GenS2{F0: 1, F1: 42, F2: v1}
                                            _ = fi27
                                            _ = v28
                                            _ = v29
                                        }
                                    case 0:
                                        switch (("hello" + ((("\n\t" + "a") + ("a" + "foo bar")) + ("a" + ""))) + "hello") {
                                        case "hello":
                                            v15 = v15
                                        case "":
                                            v30 := GenS2{F0: 722, F1: 42, F2: ((v15 << ((65535) & 63)) <= v15)}
                                            _ = v30
                                        default:
                                            fmt.Println(v1)
                                        }
                                    case -2147483648:
                                        v6 = (v6 * (v6 / ((v6) | 1)))
                                    }
                                case "":
                                    var v31 int = fi13
                                    _ = v31
                                }
                            }
                            _ = v15
                            _ = v16
                        }
                    }
                    var v32 float32 = (v3 - ((0.1 + v3) / v3))
                    v32 = ((0.0 / -2.25) + ((v3 - 1.5) - v3))
                    _ = fi13
                    _ = v32
                }
            }
        case "hello":
            var v33 float32 = (((0.1 - v3) / (v3 + 1.5)) / v3)
            _ = v33
        default:
            if v7 {
                if true {
                    if v7 {
                        switch "hello" {
                        case "":
                            for fi34 := 0; fi34 < 6; fi34++ {
                                var v35 float64 = 3.14159
                                _ = fi34
                                _ = v35
                            }
                        default:
                            v36 := GenS4{F0: 1, F1: 18446744073709551615}
                            _ = v36
                        }
                        var v37 int = 128
                        _ = v37
                    } else {
                        v7 = v7
                        for fi38 := 0; fi38 < 1; fi38++ {
                            var v39 int = (fi38 + fi38)
                            v7 = ((0.0 * 1e10) > ((v3 * float32(0.0)) - v3))
                            _ = fi38
                            _ = v39
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                } else {
                    v3 = float32(((3.14159 + 0.1) - float64(0.0)))
                }
            }
        }
    }
    if (false || (false || (!(1 > 1)))) {
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 uint = 4294967295
            var v42 bool = v1
            _ = fi40
            _ = v41
            _ = v42
        }
        if false {
            if ((!false) || (!(930 == 0))) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v1 = (4294967295 <= 738)
            } else {
                if v1 {
                    fmt.Println(v3, v7, v3)
                }
            }
        } else {
            fmt.Println(v3)
        }
    }
    v43 := GenS2{F0: 0, F1: 42, F2: v1}
    v44 := GenS5{F0: 9223372036854775807}
    for fi45 := 0; fi45 < 7; fi45++ {
        v1 = ((v6 >> ((621) & 63)) == v6)
        var v46 uint = 1
        v47 := GenS2{F0: 65535, F1: 42, F2: false}
        _ = fi45
        _ = v46
        _ = v47
    }
    v3 -= 3.14159
    v48 := GenS5{F0: 42}
    var v49 int8 = -109
    v50 := GenS6{F0: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
