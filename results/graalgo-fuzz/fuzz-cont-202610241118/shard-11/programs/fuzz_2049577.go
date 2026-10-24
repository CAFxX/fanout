package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int8; F1 int64; F2 string }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint, p1 int16, p2 bool) int {
    for fi6 := 0; fi6 < 9; fi6++ {
        p0 -= 733
        var v7 uint = (^(p0))
        var v8 uint8 = 16
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 = (p1 - 0)
    v9 := GenS1{F0: 0, F1: 1}
    v10 := []int{42, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 128
}

func main() {
    v1 := []int{9223372036854775807, -571, 32767}
    v2 := make(map[string]string)
    var v3 uint = 1
    v4 := make(map[uint32]string)
    v5 := make(map[int]int)
    switch "" {
    case "a":
        var v11 int8 = 0
        _ = v11
    case "":
        fmt.Println(v3)
    case "\n\t":
        v12 := GenS1{F0: 65535, F1: 2147483647}
        _ = v12
    }
    if (false && true) {
        switch -1 {
        case 1:
            for fi13 := 0; fi13 < 12; fi13++ {
                switch (("\n\t" + "a") + ("" + "hello")) {
                case "":
                    var v14 uint32 = 65535
                    _ = v14
                case "foo bar":
                    if (1 <= 127) {
                        for fi15 := 0; fi15 < 6; fi15++ {
                            if true {
                                v3 += v3
                                if false {
                                    v3 -= 255
                                }
                            } else {
                                var v16 int8 = 37
                                v3 -= 255
                                fmt.Println(fi15, fi13)
                                v3 -= 2
                                fmt.Println(v16, fi15)
                                _ = v16
                            }
                            fmt.Println(v3, v3, fi15)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi15
                        }
                        var v17 float64 = ((3.14159 * 3.14159) + 0.0)
                        _ = v17
                    }
                case "hello":
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v19 := GenS2{F0: -32768, F1: ("a" + ("\n\t" + "foo bar"))}
                        if true {
                            if false {
                                v3 = (v3 >> ((v3) & 63))
                                if true {
                                    var v20 bool = (!false)
                                    var v21 int32 = 0
                                    _ = v20
                                    _ = v21
                                } else {
                                    switch fi18 {
                                    case 42:
                                        v22 := GenS3{F0: true}
                                        _ = v22
                                    case 4294967295:
                                        var v23 uint = (v3 - v3)
                                        _ = v23
                                    default:
                                        v3 = (v3 % (((^(v3))) | 1))
                                    }
                                }
                            }
                            v24 := GenS1{F0: 759, F1: -2147483648}
                            _ = v24
                        }
                        _ = fi18
                        _ = v19
                    }
                default:
                    for fi25 := 0; fi25 < 8; fi25++ {
                        switch ("" + "hello") {
                        case "a":
                            v3 = (v3 % ((v3) | 1))
                        case "\n\t":
                            var v26 uint16 = 266
                            v27 := GenS3{F0: ("\n\t" != "")}
                            var v28 uint8 = 48
                            _ = v26
                            _ = v27
                            _ = v28
                        case "hello":
                            switch fi25 {
                            case -2147483648:
                                v3 = v3
                            case 32767:
                                v3 = v3
                            case 4294967295:
                                v3 -= v3
                                v3 -= v3
                                var v29 uint8 = uint8(v3)
                                _ = v29
                            default:
                                if true {
                                    var v30 float64 = float64(0.0)
                                    var v31 uint64 = 18446744073709551615
                                    v32 := GenS3{F0: false}
                                    v31 = (v31 << (((v3 << (((v3 ^ (^(v3)))) & 63))) & 63))
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                } else {
                                    v3 += (v3 ^ v3)
                                    v33 := v5[fi13]
                                    _ = v33
                                }
                            }
                        default:
                            v34 := GenS4{F0: 514, F1: 1}
                            _ = v34
                        }
                        _ = fi25
                    }
                }
                _ = fi13
            }
        default:
            v3 += (v3 >> ((v3) & 63))
        }
        v2[("foo bar" + "hello")] = (("" + ("a" + (("" + ("a" + "")) + "a"))) + "\n\t")
        switch 0 {
        case 585:
            fmt.Println(v3)
        default:
            var v35 int64 = 1
            _ = v35
        }
    }
    fmt.Println(v3)
    var v36 int16 = 1
    var v37 uint8 = 255
    var v38 float64 = 3.14159
    v36 -= v36
    switch "" {
    case "a":
        _ = len(v1)
    case "foo bar":
        v1 = append(v1, 32767)
    case "hello":
        v2[("hello" + "hello")] = "\n\t"
    }
    if (true && (true && ((v3 << (((^(v3))) & 63)) < 42))) {
        var v39 uint8 = 1
        _ = v39
    } else {
        v40 := GenS5{F0: -100, F1: -1, F2: "foo bar"}
        _ = len(v2)
        _ = v40
    }
    v37 += (^(v37))
    v36 += 32767
    var v41 uint = 18446744073709551615
    switch "\n\t" {
    case "foo bar":
        _ = len(v1)
    case "hello":
        for fi42 := 0; fi42 < 4; fi42++ {
            v43 := GenS6{F0: genFn0((^(v41)), (v36 | v36), false), F1: v3}
            if true {
                if false {
                    switch "foo bar" {
                    case "hello":
                        var v44 uint16 = 2
                        _ = v44
                    case "":
                        var v45 float32 = 1e10
                        _ = v45
                    case "a":
                        v36 = (v36 & 42)
                    default:
                        v46 := GenS7{F0: (^(v36))}
                        _ = v46
                    }
                } else {
                    delete(v2, "\n\t")
                }
            }
            fmt.Println(v41, v37, v36)
            _ = fi42
            _ = v43
        }
    case "":
        v47 := v5[128]
        _ = v47
    default:
        switch 9223372036854775807 {
        case -893:
            var v48 float64 = 1.5
            _ = v48
        case 0:
            for fi49 := 0; fi49 < 16; fi49++ {
                v37 -= (v37 - (v37 + 255))
                v50 := GenS5{F0: 122, F1: 9223372036854775807, F2: ("hello" + "")}
                fmt.Println(v38, fi49, v37)
                _ = fi49
                _ = v50
            }
        }
    }
    for fi51 := 0; fi51 < 14; fi51++ {
        var v52 int8 = 0
        if (751 < 374) {
            switch (fi51 / ((fi51) | 1)) {
            case 698:
                var v53 uint16 = 740
                _ = v53
            }
            v38 += float64(float32(float64(3.14159)))
        }
        var v54 uint = 62
        _ = fi51
        _ = v52
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}
