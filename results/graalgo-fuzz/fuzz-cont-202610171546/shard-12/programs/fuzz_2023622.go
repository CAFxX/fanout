package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint8; F1 int32; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -674
    var v3 uint16 = 0
    var v4 float32 = 3.14159
    switch 9223372036854775807 {
    case -27:
        switch -387 {
        case 127:
            v2 -= (v2 + v2)
        }
    case 32767:
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int32 = 781
            switch ("foo bar" + (("foo bar" + (("hello" + "a") + (("hello" + ((("a" + ("foo bar" + (((("hello" + "hello") + "foo bar") + "foo bar") + ("a" + "\n\t")))) + "foo bar") + "a")) + ""))) + ("hello" + "\n\t"))) {
            case "foo bar":
                v3 = (v3 / (((v3 & (^(v3)))) | 1))
            case "a":
                v3 += v3
            case "\n\t":
                for fi7 := 0; fi7 < 7; fi7++ {
                    v3 = (^(v3))
                    v8 := v1[792]
                    var v9 float32 = float32(-2.25)
                    if true {
                        var v10 uint16 = v3
                        _ = v10
                    }
                    if true {
                        v9 -= v4
                    } else {
                        if (18446744073709551615 != 18446744073709551615) {
                            v11 := []int{}
                            var v12 uint = 369
                            _ = v11
                            _ = v12
                        } else {
                            var v13 int32 = 42
                            if (!true) {
                                v14 := []int{}
                                _ = v14
                            } else {
                                var v15 uint = 255
                                v16 := GenS1{F0: (((("hello" + ("a" + ("\n\t" + "hello"))) + "a") + "") + ""), F1: false, F2: v6}
                                switch "hello" {
                                case "\n\t":
                                    v9 = v9
                                case "hello":
                                    fmt.Println(v6)
                                case "foo bar":
                                    v8 = v8
                                default:
                                    var v17 uint64 = 18446744073709551615
                                    _ = v17
                                }
                                _ = v15
                                _ = v16
                            }
                            _ = v13
                        }
                    }
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
            _ = fi5
            _ = v6
        }
    case -1:
        var v18 int = -653
        _ = v18
    default:
        var v19 int16 = 146
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS1{F0: "", F1: true, F2: (v2 | v2)}
        if false {
            fmt.Println(v3, v3)
            if true {
                for fi22 := 0; fi22 < 10; fi22++ {
                    switch 0 {
                    case -9223372036854775808:
                        v2 += (v2 * 42)
                    case 9223372036854775807:
                        v23 := v1[2147483648]
                        _ = v23
                    case -1:
                        delete(v1, 65535)
                        v24 := []int{0, fi20, 52, 273}
                        _ = v24
                    default:
                        v25 := GenS1{F0: "", F1: (((((v4 / 3.14159) * v4) <= v4) || (!true)) && true), F2: v2}
                        _ = v25
                    }
                    _ = fi22
                }
            } else {
                if true {
                    v26 := []int{0}
                    _ = v26
                } else {
                    var v27 uint8 = 1
                    _ = v27
                }
            }
        } else {
            if (((177 < 255) && true) && false) {
                var v28 int16 = -1
                v4 += 3.14159
                v29 := GenS1{F0: (("hello" + ("a" + "hello")) + (((("" + "") + "hello") + "") + "a")), F1: ("hello" <= ("" + ("hello" + "a"))), F2: (v2 % ((v2) | 1))}
                _ = v28
                _ = v29
            } else {
                v30 := GenS2{F0: fi20}
                _ = v30
            }
        }
        fmt.Println(v2)
        _ = fi20
        _ = v21
    }
    v1[4294967295] = 662
    v31 := []int{-2147483648, -470}
    var v32 uint32 = 2
    v33 := GenS3{F0: "\n\t", F1: 18446744073709551615}
    var v34 int16 = 0
    var v35 int = 204
    v32 = (v32 % (((v32 * 0)) | 1))
    delete(v1, (v32 >> ((0) & 63)))
    if (v34 <= v34) {
        v35 += (v35 << ((333) & 63))
        _ = len(v31)
        switch (-(v35)) {
        case -2147483648:
            var v36 string = (("\n\t" + "foo bar") + "a")
            _ = v36
        }
    }
    var v37 string = "foo bar"
    fmt.Println(v34, v2)
    var v38 uint64 = 18446744073709551615
    if (4294967295 < (v35 << ((1) & 63))) {
        v39 := GenS4{F0: 192, F1: v2, F2: v32}
        _ = v39
    }
    v40 := GenS1{F0: "a", F1: (100 <= -1), F2: 0}
    v32 = 42
    v34 -= v34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v40
}
