package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 uint64 }

func main() {
    var v1 uint16 = 32767
    v2 := []string{}
    v3 := make(map[int]int)
    var v4 uint8 = 0
    v4 = 102
    switch 1 {
    case 0:
        v4 = 108
    default:
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := GenS1{F0: int32(v4)}
            v2 = append(v2, "a")
            _ = fi5
            _ = v6
        }
    }
    var v7 uint8 = v4
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v7, v4)
        v4 = (v7 % (((v4 << ((18446744073709551615) & 63))) | 1))
        _ = fi8
    }
    if true {
        v9 := GenS2{F0: 42}
        var v10 uint16 = v1
        v11 := GenS3{F0: 4294967295, F1: v1, F2: "foo bar"}
        _ = v10
        _ = v11
        _ = v9
    } else {
        v3[32767] = 9223372036854775807
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int64 = 1
            if true {
                var v15 int64 = -9223372036854775808
                switch fi12 {
                case 32768:
                    var v16 int32 = -6
                    _ = v16
                default:
                    delete(v3, fi12)
                }
                _ = v15
            } else {
                delete(v3, (fi12 - fi12))
                var v17 float32 = 0.1
                if false {
                    switch fi12 {
                    case 92:
                        delete(v3, fi13)
                    case 0:
                        var v18 string = "a"
                        switch ((((v18 + "a") + "a") + "a") + (v18 + v18)) {
                        case "":
                            if true {
                                delete(v3, fi13)
                            } else {
                                fmt.Println(v17)
                                v2 = append(v2, v18)
                                switch (fi12 << ((uint(v1)) & 63)) {
                                case 9223372036854775807:
                                    fmt.Println(v7, v7, fi12)
                                case -2147483648:
                                    v19 := GenS4{F0: 179}
                                    _ = v19
                                case -9223372036854775808:
                                    switch ("a" + (v18 + "hello")) {
                                    case "hello":
                                        v20 := GenS3{F0: 65535, F1: v1, F2: (v18 + "hello")}
                                        _ = v20
                                    case "\n\t":
                                        if false {
                                            _ = len(v3)
                                            v3[fi13] = (-(fi13))
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            v3[fi13] = fi12
                                            v7 = v4
                                            v1 -= 2
                                            v21 := GenS1{F0: -1}
                                            _ = v21
                                        }
                                    case "foo bar":
                                        var v22 uint = 0
                                        _ = v22
                                    default:
                                        fmt.Println(v4, fi13)
                                    }
                                default:
                                    var v23 int64 = v14
                                    _ = v23
                                }
                            }
                        case "foo bar":
                            switch fi12 {
                            case -9223372036854775808:
                                v17 -= 1e10
                            case -2147483648:
                                var v24 float64 = float64(0.1)
                                _ = v24
                            default:
                                if ((v14 >= v14) && false) {
                                    var v25 uint16 = 42
                                    v25 += (v25 << ((2) & 63))
                                    v2 = append(v2, v18)
                                    _ = v25
                                } else {
                                    v1 -= v1
                                    v4 += uint8(v1)
                                }
                            }
                            v4 -= (v4 - v7)
                        default:
                            fmt.Println(v1, fi12, v17)
                        }
                        delete(v3, fi13)
                        fmt.Println(v4, v14)
                        switch (fi13 % (((fi12 * fi13)) | 1)) {
                        case 651:
                            var v26 uint32 = 2
                            _ = v26
                        case 1:
                            v27 := GenS4{F0: 828}
                            _ = v27
                        default:
                            v17 = (((v17 * v17) * ((v17 / (((v17 - v17) / v17) / 3.14159)) * v17)) - ((v17 * v17) - 3.14159))
                        }
                        _ = v18
                    case -540:
                        fmt.Println(v14)
                    }
                    switch "\n\t" {
                    case "foo bar":
                        v4 = (^(v4))
                    case "":
                        var v28 uint8 = v4
                        _ = v28
                    case "\n\t":
                        switch fi13 {
                        case 127:
                            v3[(-(fi13))] = (fi13 | 2147483647)
                        case 9223372036854775807:
                            v29 := GenS5{F0: 0}
                            _ = v29
                        case 42:
                            v30 := GenS1{F0: -162}
                            _ = v30
                        }
                    }
                } else {
                    _ = len(v2)
                    v17 = ((1.5 / (float32(1e10) * (-2.25 / 1.5))) - v17)
                }
                _ = v17
            }
            v14 = v14
            var v31 int32 = 42
            switch "hello" {
            case "foo bar":
                v4 = v4
            case "a":
                var v32 int32 = v31
                _ = v32
            case "":
                var v33 string = ""
                _ = v33
            default:
                _ = len(v2)
            }
            _ = fi13
            _ = v14
            _ = v31
        }
        _ = fi12
    }
    if (((0.1 / 1.5) - 0.1) == -2.25) {
        var v34 uint = 0
        _ = v34
    }
    if true {
        if (!(!(!false))) {
            var v35 uint64 = 388
            var v36 float64 = 0.1
            v37 := GenS6{F0: 1, F1: v35}
            _ = v35
            _ = v36
            _ = v37
        }
        v7 = v4
    } else {
        v7 = (v4 ^ 128)
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        if len(v2) > 0 {
            v2[0] = (("hello" + "a") + "a")
        }
        _ = fi38
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        for fi40 := 0; fi40 < 16; fi40++ {
            var v41 int8 = 0
            v7 += (v7 >> ((18446744073709551615) & 63))
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    fmt.Println(v4, v4, v4)
    fmt.Println(v4, v7)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
