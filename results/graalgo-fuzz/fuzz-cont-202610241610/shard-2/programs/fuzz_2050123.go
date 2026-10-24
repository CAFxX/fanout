package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 int }
type GenS3 struct { F0 uint; F1 uint }

func main() {
    var v1 uint8 = 23
    var v2 bool = true
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 255
    var v5 int64 = 0
    var v6 float32 = 1e10
    v1 += v1
    var v7 uint16 = 2
    fmt.Println(v7)
    v5 += (v5 ^ int64(v1))
    v8 := GenS1{F0: 1}
    switch ("a" + "\n\t") {
    case "foo bar":
        switch "hello" {
        case "\n\t":
            fmt.Println(v4, v6, v1)
        case "":
            var v9 int32 = 681
            _ = v9
        case "a":
            if true {
                for fi10 := 0; fi10 < 3; fi10++ {
                    fmt.Println(v7, v2, fi10)
                    _ = fi10
                }
            }
        }
    }
    var v11 bool = v2
    if v2 {
        if true {
            v12 := GenS1{F0: -2147483648}
            _ = v12
        } else {
            switch "hello" {
            case "foo bar":
                v11 = false
            case "":
                delete(v3, 65535)
            case "\n\t":
                fmt.Println(v7, v1, v1)
            default:
                if (v2 && ((v11 && true) && v2)) {
                    v11 = (100 <= 128)
                    var v13 float64 = 1e10
                    v13 = ((((v13 + v13) + (float64(v6) - (v13 * -2.25))) / v13) / 3.14159)
                    _ = v13
                }
            }
        }
        var v14 float32 = 0.1
        _ = v14
    } else {
        if v11 {
            v15 := GenS1{F0: 1}
            _ = v15
        } else {
            v7 = 1
            v16 := GenS2{F0: 200, F1: 0, F2: -1}
            v2 = v11
            _ = v16
        }
        switch -1 {
        case 127:
            v3[4294967295] = 2147483647
        case 1:
            var v17 float64 = -2.25
            _ = v17
        default:
            var v18 string = "hello"
            _ = v18
        }
        v19 := GenS2{F0: 730, F1: 1, F2: 0}
        _ = v19
    }
    fmt.Println(v1, v7, v4)
    switch ("a" + (("a" + "\n\t") + "a")) {
    case "\n\t":
        v6 = v6
    case "foo bar":
        var v20 int64 = (v5 & 84)
        _ = v20
    }
    v21 := []int{4294967295, 32768, 32768}
    if ((("hello" + ("a" + "")) + "a") >= "") {
        v2 = (!(42 < 9223372036854775807))
        v11 = (v7 <= (v4 + 42))
    } else {
        switch 587 {
        case 802:
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 string = "\n\t"
                switch "hello" {
                case "hello":
                    _ = len(v3)
                case "\n\t":
                    var v24 uint16 = v7
                    _ = v24
                case "":
                    v6 -= (v6 / float32(0.1))
                default:
                    v25 := GenS2{F0: 1, F1: 4294967295, F2: fi22}
                    _ = v25
                }
                switch v23 {
                case "foo bar":
                    v1 = v1
                case "hello":
                    for fi26 := 0; fi26 < 15; fi26++ {
                        if ((42 <= 4294967295) && ((!v11) || v11)) {
                            switch v23 {
                            case "a":
                                delete(v3, 2)
                            case "\n\t":
                                var v27 uint64 = 275
                                _ = v27
                            case "hello":
                                fmt.Println(v4)
                            default:
                                if v2 {
                                    if v11 {
                                        v6 += 0.1
                                    } else {
                                        if (!((((!v2) || false) && false) && (!(1.5 < 0.1)))) {
                                            var v28 float64 = 3.14159
                                            _ = v28
                                        }
                                    }
                                    if v2 {
                                        if v2 {
                                            v29 := GenS2{F0: 605, F1: uint64(v1), F2: (fi22 >> ((272) & 63))}
                                            _ = v29
                                        }
                                    }
                                }
                            }
                            var v30 float32 = ((-2.25 * float32(float64(v6))) + v6)
                            switch (fi22 / (((fi26 * fi22)) | 1)) {
                            case -1:
                                v7 -= v4
                                var v31 uint16 = 434
                                v6 += v6
                                _ = v31
                            }
                            _ = v30
                        } else {
                            var v32 uint16 = v7
                            _ = v32
                        }
                        v1 = v1
                        _ = fi26
                    }
                }
                _ = fi22
                _ = v23
            }
        case 32767:
            var v33 float32 = v6
            _ = v33
        default:
            if (!v2) {
                if false {
                    var v34 string = "hello"
                    if v2 {
                        if len(v21) > 0 {
                            v21[0] = 128
                        }
                    }
                    _ = v34
                } else {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                }
                switch "foo bar" {
                case "hello":
                    if ((v7 ^ v4) != (^(v7))) {
                        v21 = append(v21, 9223372036854775807)
                        for fi35 := 0; fi35 < 20; fi35++ {
                            switch (((("foo bar" + "hello") + ("" + ("hello" + ""))) + "foo bar") + "hello") {
                            case "foo bar":
                                var v36 uint8 = v1
                                _ = v36
                            case "":
                                v2 = (-2.25 != 1.5)
                            case "hello":
                                fmt.Println(v2, fi35, v1)
                            default:
                                v2 = (v11 || false)
                            }
                            v37 := GenS3{F0: 4294967295, F1: 970}
                            _ = fi35
                            _ = v37
                        }
                    } else {
                        if v11 {
                            v1 = v1
                        }
                        if (v5 < v5) {
                            v11 = (v2 || v2)
                            if (false || (!v2)) {
                                var v38 float32 = ((float32(-2.25) - ((float32(-2.25) / 3.14159) - (v6 / v6))) - 0.1)
                                _ = v38
                            } else {
                                var v39 int64 = (v5 + v5)
                                _ = v39
                            }
                            var v40 uint32 = 4294967295
                            _ = v40
                        } else {
                            v6 = (1e10 * v6)
                            switch "foo bar" {
                            case "":
                                var v41 float32 = (1e10 + 1.5)
                                _ = v41
                            default:
                                if v2 {
                                    var v42 uint8 = (^(v1))
                                    v5 += v5
                                    _ = v42
                                }
                            }
                        }
                    }
                case "foo bar":
                    var v43 int64 = (v5 % ((9223372036854775807) | 1))
                    _ = v43
                case "a":
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                default:
                    var v44 uint32 = 65535
                    _ = v44
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
