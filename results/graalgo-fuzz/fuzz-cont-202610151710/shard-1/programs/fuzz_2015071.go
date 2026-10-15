package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8; F1 uint64; F2 string }
type GenS4 struct { F0 int64; F1 uint16 }

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    v5 := []int{42, -1}
    v1 = (v1 & (^(v1)))
    switch 128 {
    case 32768:
        v6 := GenS1{F0: (504 == 42), F1: 63}
        _ = v6
    case -11:
        switch "hello" {
        case "hello":
            if (!(9223372036854775807 >= 4294967295)) {
                if ((false && true) || false) {
                    if (!(1 < -9223372036854775808)) {
                        var v7 int64 = -9223372036854775808
                        _ = v7
                    }
                    for fi8 := 0; fi8 < 7; fi8++ {
                        if ((1.5 >= (-2.25 - 3.14159)) || (127 < -128)) {
                            for fi9 := 0; fi9 < 19; fi9++ {
                                var v10 uint = 18446744073709551615
                                if false {
                                    var v11 uint16 = 42
                                    v11 = v11
                                    _ = v11
                                } else {
                                    var v12 uint8 = 185
                                    if true {
                                        v10 += (^(v10))
                                    }
                                    _ = v12
                                }
                                v13 := GenS2{F0: "foo bar", F1: (v10 & 18446744073709551615)}
                                _ = fi9
                                _ = v10
                                _ = v13
                            }
                        }
                        for fi14 := 0; fi14 < 11; fi14++ {
                            var v15 uint64 = 1
                            var v16 float64 = 0.0
                            v17 := GenS2{F0: "\n\t", F1: 558}
                            _ = fi14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        v1 += 130
                        _ = fi8
                    }
                } else {
                    fmt.Println(v1)
                    if ("foo bar" != "foo bar") {
                        _ = len(v5)
                    } else {
                        for fi18 := 0; fi18 < 6; fi18++ {
                            var v19 uint32 = 354
                            _ = fi18
                            _ = v19
                        }
                        v3[-323] = 2
                    }
                }
                var v20 int = 656
                var v21 int8 = 0
                _ = v20
                _ = v21
            } else {
                v1 = (v1 % ((v1) | 1))
                v1 = (v1 | (v1 >> ((255) & 63)))
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case "\n\t":
            _ = len(v3)
        default:
            if (("" + "a") < ("a" + "a")) {
                switch "foo bar" {
                case "hello":
                    var v22 uint32 = 0
                    _ = v22
                case "a":
                    if (!(!true)) {
                        for fi23 := 0; fi23 < 11; fi23++ {
                            var v24 float64 = 3.14159
                            switch "foo bar" {
                            case "hello":
                                if false {
                                    v5 = append(v5, (fi23 + -9223372036854775808))
                                } else {
                                    var v25 float32 = float32(0.0)
                                    fmt.Println(v1)
                                    _ = v25
                                }
                            }
                            v1 -= 81
                            _ = fi23
                            _ = v24
                        }
                    }
                case "":
                    v1 += (^(v1))
                }
                _ = len(v4)
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch "\n\t" {
                case "hello":
                    var v26 bool = true
                    _ = v26
                case "foo bar":
                    v27 := GenS3{F0: (v1 / ((v1) | 1)), F1: 1, F2: "a"}
                    _ = v27
                case "":
                    var v28 uint16 = 1
                    _ = v28
                default:
                    delete(v2, "")
                }
            }
        }
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        if false {
            var v29 int = 4294967295
            var v30 float32 = 3.14159
            var v31 uint = 2
            _ = v29
            _ = v30
            _ = v31
        } else {
            var v32 uint8 = (^(v1))
            _ = v32
        }
    }
    fmt.Println(v1)
    v33 := GenS4{F0: -9223372036854775808, F1: 65535}
    var v34 float32 = 0.1
    v34 = v34
    var v35 float64 = 0.1
    fmt.Println(v35)
    v1 = (^(v1))
    v1 = (v1 << ((385) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
