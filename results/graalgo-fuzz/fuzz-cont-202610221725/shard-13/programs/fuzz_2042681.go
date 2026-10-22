package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int32; F1 int; F2 int8 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int; F1 int; F2 uint8 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    var v3 bool = false
    v4 := make(map[uint32]int)
    var v5 float64 = -2.25
    var v6 int64 = 9223372036854775807
    v4[4294967295] = -1
    v7 := []int{}
    fmt.Println(v3, v6)
    _ = len(v7)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int32 = -55
        for fi10 := 0; fi10 < 11; fi10++ {
            v7 = append(v7, -456)
            if (1 <= 0) {
                _ = len(v7)
                v11 := GenS1{F0: ((v3 && false) || (!(!false))), F1: 4294967295, F2: v3}
                _ = v11
            } else {
                if v3 {
                    v1 = v3
                    var v12 bool = (((v9 + v9) != -647) && (!(v1 && (32767 < -397))))
                    _ = v12
                } else {
                    var v13 uint64 = 18446744073709551615
                    v6 -= (v6 & v6)
                    switch 128 {
                    case 32767:
                        switch (^(fi10)) {
                        case 9223372036854775807:
                            if v1 {
                                var v14 uint64 = v13
                                var v15 uint16 = 655
                                switch (fi10 - -9223372036854775808) {
                                case 494:
                                    v13 = v13
                                default:
                                    switch (fi8 * 128) {
                                    case 42:
                                        var v16 uint32 = 65535
                                        _ = v16
                                    }
                                }
                                _ = v14
                                _ = v15
                            } else {
                                if (!true) {
                                    var v17 int = (fi10 >> ((4294967295) & 63))
                                    v1 = false
                                    var v18 string = ("" + ((("" + "foo bar") + ("a" + ("" + "foo bar"))) + "\n\t"))
                                    _ = v17
                                    _ = v18
                                }
                                if (!(!v3)) {
                                    v6 = -9223372036854775808
                                }
                                switch "a" {
                                case "foo bar":
                                    if ((3.14159 >= 1.5) && ((1000 > 244) && v3)) {
                                        v13 += (v13 + 2)
                                        switch (-(fi8)) {
                                        case 9223372036854775807:
                                            var v19 float32 = (1e10 - 0.1)
                                            _ = v19
                                        default:
                                            switch ("" + "hello") {
                                            case "":
                                                v1 = true
                                            case "a":
                                                if false {
                                                    var v20 float64 = (3.14159 + v5)
                                                    v5 = v5
                                                    switch (fi8 + (fi8 * -1)) {
                                                    case 0:
                                                        if ((v9 | -10) != (v9 - int32(fi8))) {
                                                            v5 = 0.0
                                                        } else {
                                                            var v21 int64 = -1
                                                            if (-847 != 1000) {
                                                                var v22 int32 = v9
                                                                if len(v7) > 0 {
                                                                    v7[0] = (fi10 << ((1) & 63))
                                                                }
                                                                v6 += (-(v6))
                                                                _ = v22
                                                            }
                                                            var v23 uint = 952
                                                            _ = v21
                                                            _ = v23
                                                        }
                                                    case 127:
                                                        v24 := GenS2{F0: fi10, F1: (fi8 % ((fi10) | 1))}
                                                        _ = v24
                                                    default:
                                                        var v25 uint64 = 4294967295
                                                        _ = v25
                                                    }
                                                    _ = v20
                                                } else {
                                                    if true {
                                                        v26 := GenS1{F0: false, F1: 1, F2: (-1 <= v9)}
                                                        var v27 uint8 = 0
                                                        var v28 float64 = (v5 + 3.14159)
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                }
                                            case "foo bar":
                                                if ((("" + "") + ("\n\t" + "\n\t")) != "a") {
                                                    if v3 {
                                                        var v29 uint32 = 0
                                                        _ = v29
                                                    } else {
                                                        fmt.Println(v1, v1)
                                                        v13 -= 718
                                                    }
                                                }
                                            }
                                        }
                                        if v1 {
                                            v7 = append(v7, fi10)
                                        } else {
                                            if v1 {
                                                v30 := GenS3{F0: (v9 << ((219) & 63)), F1: fi8, F2: 127}
                                                v13 -= v13
                                                v6 = (^(v6))
                                                _ = v30
                                            } else {
                                                var v31 uint16 = 0
                                                var v32 string = "a"
                                                _ = v31
                                                _ = v32
                                            }
                                            fmt.Println(v13)
                                            switch ("a" + "a") {
                                            case "hello":
                                                v1 = (v3 || true)
                                            case "":
                                                var v33 int16 = 1
                                                _ = v33
                                            case "foo bar":
                                                fmt.Println(fi10, v1, v3)
                                            }
                                        }
                                    }
                                case "":
                                    v6 = (v6 + -9223372036854775808)
                                }
                            }
                        }
                    case -9223372036854775808:
                        v13 += v13
                    default:
                        if v1 {
                            if v1 {
                                var v34 uint8 = 255
                                _ = v34
                            }
                        }
                    }
                    _ = v13
                }
            }
            var v35 int32 = (^(v9))
            v36 := GenS1{F0: (32767 < -554), F1: 65535, F2: false}
            v9 += v35
            _ = fi10
            _ = v35
            _ = v36
        }
        _ = fi8
        _ = v9
    }
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    v37 := GenS1{F0: false, F1: 100, F2: true}
    v38 := GenS4{F0: -128, F1: v6}
    var v39 int8 = 100
    for fi40 := 0; fi40 < 14; fi40++ {
        v7 = append(v7, (fi40 & (-(fi40))))
        for fi41 := 0; fi41 < 3; fi41++ {
            v6 = v6
            fmt.Println(v3, v5)
            v42 := v2[472]
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    v5 = float64(float32((v5 + v5)))
    if false {
        switch "hello" {
        case "a":
            fmt.Println(v1, v5)
        case "\n\t":
            switch (("foo bar" + "foo bar") + (("\n\t" + "a") + "a")) {
            case "hello":
                var v43 float32 = (float32(0.0) + (3.14159 / 3.14159))
                _ = v43
            default:
                var v44 uint = 4294967295
                _ = v44
            }
        case "foo bar":
            var v45 string = "\n\t"
            _ = v45
        }
        switch -829 {
        case 1:
            v3 = true
        case -2147483648:
            if v3 {
                v46 := GenS4{F0: v39, F1: 1}
                _ = v46
            }
        default:
            switch 128 {
            case 0:
                v47 := GenS5{F0: -745, F1: 9223372036854775807, F2: 2}
                _ = v47
            case 1:
                v48 := v2[0]
                _ = v48
            }
        }
    }
    switch 9223372036854775807 {
    case -2147483648:
        v5 -= v5
    case 13:
        var v49 int64 = -1
        _ = v49
    default:
        v50 := GenS1{F0: v3, F1: 353, F2: v3}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
