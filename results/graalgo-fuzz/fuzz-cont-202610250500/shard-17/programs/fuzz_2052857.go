package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint8; F1 string; F2 uint16 }
func genFn0(p0 uint8, p1 uint64) int64 {
    var v5 uint16 = 565
    var v6 string = "foo bar"
    fmt.Println(v6, p1, v6)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 int32 = 1
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 int32 = (v1 | 1)
    if true {
        switch 2147483647 {
        case 128:
            v7 := GenS1{F0: 21}
            _ = v7
        case 654:
            switch "" {
            case "hello":
                fmt.Println(v4, v4)
            case "a":
                v1 = v1
            }
        }
        if ((float64(float32(1.5)) * 3.14159) >= -2.25) {
            for fi8 := 0; fi8 < 11; fi8++ {
                fmt.Println(v4, v1, fi8)
                _ = fi8
            }
        }
    }
    v9 := v3[1]
    switch (("" + "") + "") {
    case "foo bar":
        var v10 int8 = 1
        _ = v10
    }
    v11 := []int{0, 1, -2147483648}
    _ = len(v2)
    var v12 int = -259
    var v13 float32 = 0.1
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint64 = 268
        if (v9 == v9) {
            v11 = append(v11, fi14)
            if true {
                switch "hello" {
                case "a":
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v1 -= v4
                        if (v13 != -2.25) {
                            v17 := GenS2{F0: "a", F1: 0}
                            v18 := GenS1{F0: 42}
                            _ = v17
                            _ = v18
                        } else {
                            v15 = (^(v15))
                            if true {
                                if false {
                                    if false {
                                        if ((408 != 1) && false) {
                                            _ = len(v11)
                                        } else {
                                            var v19 float64 = 0.0
                                            _ = v19
                                        }
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        v3[(-(v12))] = v9
                                    }
                                    var v20 float32 = v13
                                    _ = v20
                                }
                                var v21 float64 = (3.14159 - 0.1)
                                v22 := GenS2{F0: ("foo bar" + ""), F1: 0}
                                _ = v21
                                _ = v22
                            } else {
                                var v23 string = "a"
                                var v24 uint64 = 34
                                var v25 uint = 706
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            if (3.14159 != -2.25) {
                                v4 += 2147483647
                                var v26 int16 = -32768
                                _ = v26
                            }
                        }
                        v13 = 1.5
                        _ = fi16
                    }
                }
                if ((0.0 == v13) || false) {
                    if ((!(-1 > -9223372036854775808)) && true) {
                        var v27 int32 = -2147483648
                        _ = v27
                    }
                    v1 = (v1 % (((v4 * v1)) | 1))
                }
                switch "foo bar" {
                case "foo bar":
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                case "\n\t":
                    var v28 int = 322
                    _ = v28
                case "a":
                    if false {
                        switch "" {
                        case "a":
                            for fi29 := 0; fi29 < 20; fi29++ {
                                fmt.Println(v12, v4)
                                _ = fi29
                            }
                        case "hello":
                            var v30 float64 = ((((3.14159 + (-2.25 / 0.1)) / -2.25) * 1.5) + 1e10)
                            _ = v30
                        case "foo bar":
                            if true {
                                v13 += v13
                                v15 -= v15
                                v12 -= (-(fi14))
                            } else {
                                v31 := GenS2{F0: "foo bar", F1: 0}
                                _ = v31
                            }
                        }
                        v32 := GenS3{F0: 0, F1: ("hello" + ((("a" + "\n\t") + ("foo bar" + "\n\t")) + (("" + "a") + "foo bar"))), F2: 65535}
                        _ = v32
                    } else {
                        switch "a" {
                        case "a":
                            switch "foo bar" {
                            case "":
                                var v33 float64 = 1e10
                                _ = v33
                            }
                        case "foo bar":
                            switch ("" + "a") {
                            case "\n\t":
                                if true {
                                    var v34 int16 = 1
                                    if ((int32(v9) == (v1 - (^(v1)))) && (42 > 1)) {
                                        var v35 float32 = float32(3.14159)
                                        v34 = v34
                                        _ = v35
                                    }
                                    _ = v34
                                }
                            case "a":
                                v15 = (v15 + 42)
                            default:
                                var v36 uint32 = v9
                                _ = v36
                            }
                        case "":
                            v13 = ((3.14159 * (v13 / 1.0)) * 0.1)
                        default:
                            v37 := GenS1{F0: 93}
                            _ = v37
                        }
                    }
                }
            }
            var v38 float32 = 0.1
            _ = v38
        }
        _ = fi14
        _ = v15
    }
    v39 := GenS2{F0: ("hello" + (("\n\t" + "") + "\n\t")), F1: 32767}
    for fi40 := 0; fi40 < 19; fi40++ {
        for fi41 := 0; fi41 < 7; fi41++ {
            switch fi40 {
            case 4294967295:
                v4 = v4
            case -477:
                v1 += 621
                v42 := GenS2{F0: "foo bar", F1: 0}
                v43 := GenS1{F0: 255}
                _ = v42
                _ = v43
            default:
                if (1 == v1) {
                    if (525 < 0) {
                        v12 += (fi41 / ((v12) | 1))
                        var v44 bool = true
                        var v45 bool = v44
                        _ = v44
                        _ = v45
                    }
                    v9 = (v9 & (^(v9)))
                    var v46 bool = true
                    _ = v46
                }
            }
            fmt.Println(fi40)
            _ = fi41
        }
        v1 -= v4
        _ = fi40
    }
    var v47 int64 = -9223372036854775808
    fmt.Println(v12, v12)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v47
    _ = v9
}
