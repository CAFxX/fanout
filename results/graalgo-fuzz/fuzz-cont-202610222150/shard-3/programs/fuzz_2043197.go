package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := []string{"foo bar", ("" + "\n\t")}
    var v3 bool = (2147483648 > 4294967295)
    var v4 uint64 = 2
    v5 := make(map[int]bool)
    v6 := GenS1{F0: -32768, F1: 42}
    for fi7 := 0; fi7 < 15; fi7++ {
        switch "foo bar" {
        case "":
            v3 = ("a" != "a")
        case "hello":
            v4 = (v4 % (((^(v4))) | 1))
        case "a":
            if v3 {
                if (1 <= 704) {
                    if v3 {
                        if (v4 <= 2) {
                            if v3 {
                                var v8 uint8 = 127
                                _ = v8
                            } else {
                                v9 := GenS2{F0: ((((((v4 >> ((0) & 63)) != (^(v4))) || (v4 != (v4 * v4))) || v3) || (("" <= "foo bar") || (127 > 133))) && ((!false) || (0 >= 375))), F1: 1}
                                _ = v9
                            }
                            fmt.Println(v3, v3)
                        }
                        v10 := GenS3{F0: 255, F1: 601, F2: 35}
                        _ = v10
                    }
                }
                if v3 {
                    var v11 uint16 = 65535
                    _ = v11
                } else {
                    if (1e10 != 3.14159) {
                        var v12 uint64 = 2
                        _ = v12
                    }
                }
                if true {
                    v3 = false
                    var v13 uint = 972
                    var v14 uint32 = 884
                    _ = v13
                    _ = v14
                } else {
                    var v15 int8 = 100
                    _ = v15
                }
            } else {
                fmt.Println(fi7)
                if ((v4 >> ((106) & 63)) != v4) {
                    fmt.Println(v3)
                } else {
                    if false {
                        var v16 bool = (!false)
                        v17 := GenS4{F0: false}
                        _ = v16
                        _ = v17
                    } else {
                        v18 := GenS4{F0: false}
                        if (!(!true)) {
                            v4 = v4
                        } else {
                            var v19 float32 = 1e10
                            _ = v19
                        }
                        _ = v18
                    }
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if (-1 > 0) {
                            if v3 {
                                var v21 float32 = (((3.14159 - 3.14159) - (0.1 / 1.5)) / 1.0)
                                switch ("hello" + "foo bar") {
                                case "a":
                                    v22 := GenS2{F0: (!v3), F1: 2}
                                    _ = v22
                                case "":
                                    if (!((v3 || (v3 || (!(1000 == 1)))) && v3)) {
                                        v3 = (((v3 || (v4 > v4)) || ("\n\t" <= "hello")) || (416 <= 2))
                                        v3 = (false && (1 >= 1))
                                        var v23 int16 = -667
                                        v3 = (65535 == 432)
                                        v4 -= v4
                                        _ = v23
                                    }
                                case "\n\t":
                                    v3 = (1 > -12)
                                }
                                var v24 int = 32767
                                _ = v21
                                _ = v24
                            } else {
                                var v25 uint16 = 0
                                fmt.Println(fi20, v4, fi7)
                                var v26 int = fi7
                                _ = v25
                                _ = v26
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        if true {
                            var v27 uint8 = 60
                            _ = v27
                        }
                        if false {
                            var v28 float64 = (1e10 * ((1.5 * -2.25) * ((0.1 * 1e10) / 1.0)))
                            if (fi20 < (fi20 * 127)) {
                                v29 := GenS3{F0: 633, F1: -508, F2: 4294967295}
                                _ = v29
                            }
                            v3 = (v3 || v3)
                            _ = v28
                        } else {
                            v30 := GenS4{F0: v3}
                            if (!true) {
                                v3 = (!((float64((1.5 / 1.0)) + 1.5) == 0.0))
                            }
                            var v31 int8 = 100
                            _ = v30
                            _ = v31
                        }
                        _ = fi20
                    }
                    switch "a" {
                    case "foo bar":
                        v3 = false
                    default:
                        v32 := GenS1{F0: 42, F1: 9223372036854775807}
                        _ = v32
                    }
                }
            }
        default:
            var v33 float32 = 1e10
            _ = v33
        }
        _ = fi7
    }
    v2 = append(v2, ("foo bar" + (("hello" + ("" + "a")) + "\n\t")))
    v34 := GenS5{F0: 1}
    var v35 bool = (270 <= 479)
    fmt.Println(v3, v4, v35)
    v36 := GenS2{F0: (!v35), F1: 537}
    var v37 uint8 = 2
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 bool = v35
        v40 := GenS1{F0: -232, F1: int64(fi38)}
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
