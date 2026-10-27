package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }

func main() {
    var v1 bool = (255 <= 614)
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v4 uint16 = 255
    _ = len(v3)
    v1 = (913 < 18446744073709551615)
    if (1 >= 42) {
        delete(v2, 0)
    } else {
        var v5 int16 = -1
        fmt.Println(v4, v1)
        _ = v5
    }
    v1 = false
    if v1 {
        v1 = (0 != 4294967295)
    }
    switch "a" {
    case "\n\t":
        for fi6 := 0; fi6 < 3; fi6++ {
            v4 = v4
            fmt.Println(v1)
            _ = fi6
        }
    case "a":
        var v7 uint16 = v4
        _ = v7
    default:
        switch ("" + ((("\n\t" + "a") + "") + ("foo bar" + "a"))) {
        case "a":
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 float32 = (((((0.1 / float32(0.1)) + -2.25) + 1e10) - 1.5) - 0.0)
                if true {
                    fmt.Println(fi8)
                    var v10 int8 = 1
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
        case "\n\t":
            switch 2147483647 {
            case 2147483647:
                _ = len(v3)
            case -2147483648:
                var v11 uint8 = 1
                _ = v11
            case 9223372036854775807:
                if v1 {
                    if (0.1 == -2.25) {
                        v1 = (1 >= 0)
                        v12 := GenS1{F0: 1, F1: 0, F2: (("foo bar" + "") + "\n\t")}
                        _ = v12
                    }
                    v13 := GenS1{F0: -92, F1: 663, F2: (((("foo bar" + "a") + "hello") + (("a" + "\n\t") + "a")) + "foo bar")}
                    if v1 {
                        v14 := GenS2{F0: 0, F1: -2147483648, F2: 1}
                        _ = len(v3)
                        v4 = 808
                        _ = v14
                    } else {
                        v4 += 0
                        switch 1 {
                        case -9223372036854775808:
                            if v1 {
                                if (388 >= 2) {
                                    if (127 > 1) {
                                        v1 = false
                                        switch ("" + ("hello" + "foo bar")) {
                                        case "":
                                            v15 := GenS2{F0: 0, F1: 558, F2: 0}
                                            _ = v15
                                        case "foo bar":
                                            for fi16 := 0; fi16 < 13; fi16++ {
                                                var v17 int8 = -122
                                                var v18 uint64 = 980
                                                if v1 {
                                                    for fi19 := 0; fi19 < 2; fi19++ {
                                                        v20 := GenS2{F0: 7, F1: -2147483648, F2: 42}
                                                        _ = len(v2)
                                                        _ = fi19
                                                        _ = v20
                                                    }
                                                    v18 = 2
                                                    v4 -= (v4 << ((524) & 63))
                                                }
                                                _ = fi16
                                                _ = v17
                                                _ = v18
                                            }
                                        case "\n\t":
                                            var v21 int32 = 42
                                            _ = v21
                                        default:
                                            for fi22 := 0; fi22 < 6; fi22++ {
                                                fmt.Println(fi22)
                                                _ = fi22
                                            }
                                        }
                                    } else {
                                        switch 802 {
                                        case 737:
                                            fmt.Println(v1, v4)
                                        case 9223372036854775807:
                                            fmt.Println(v4)
                                        case 128:
                                            var v23 int16 = 1000
                                            _ = v23
                                        }
                                        v24 := v2[679]
                                        var v25 int16 = 1
                                        _ = v24
                                        _ = v25
                                    }
                                } else {
                                    var v26 uint32 = 2
                                    _ = len(v3)
                                    _ = v26
                                }
                            } else {
                                if true {
                                    switch 2147483647 {
                                    case 1:
                                        if (-2147483648 > 32767) {
                                            var v27 int8 = -47
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v28 int8 = 42
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v29 := GenS3{F0: 32767}
                                            fmt.Println(v4)
                                            _ = v29
                                        }
                                    default:
                                        var v30 int32 = 2147483647
                                        _ = v30
                                    }
                                    fmt.Println(v1, v4)
                                } else {
                                    v4 += 375
                                    if ((42 <= -32768) || (2 >= 42)) {
                                        var v31 int8 = 42
                                        if (!(213 < 1)) {
                                            v32 := GenS4{F0: -32768, F1: -9223372036854775808}
                                            _ = v32
                                        }
                                        _ = v31
                                    }
                                    v4 = v4
                                }
                            }
                        case 32767:
                            for fi33 := 0; fi33 < 4; fi33++ {
                                v1 = v1
                                v34 := GenS5{F0: 0}
                                v35 := GenS6{F0: 9223372036854775807}
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        default:
                            v3 = append(v3, "foo bar")
                        }
                    }
                    _ = v13
                } else {
                    if true {
                        var v36 float64 = ((0.0 - float64((1e10 * 1.5))) * 3.14159)
                        switch "foo bar" {
                        case "foo bar":
                            v1 = v1
                        case "":
                            fmt.Println(v4, v4)
                        case "a":
                            v36 = v36
                        }
                        for fi37 := 0; fi37 < 5; fi37++ {
                            if (v1 && false) {
                                v4 = v4
                            }
                            v38 := GenS2{F0: 35, F1: int32(fi37), F2: 560}
                            _ = fi37
                            _ = v38
                        }
                        _ = v36
                    }
                }
            default:
                var v39 int64 = 1
                _ = v39
            }
        }
    }
    v1 = v1
    v4 = v4
    if v1 {
        v1 = v1
        v1 = (95 >= 100)
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        var v41 uint32 = 969
        _ = fi40
        _ = v41
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v4 = v4
    var v42 uint16 = 255
    if (723 <= 1) {
        v1 = (!false)
        var v43 int = -115
        _ = v43
    }
    v42 = 2
    v44 := v2[9223372036854775807]
    if (("foo bar" + "hello") == ((("" + "a") + "a") + (("" + ("a" + ("a" + ""))) + ((("" + "") + (("\n\t" + "\n\t") + ("a" + "a"))) + "")))) {
        switch -897 {
        case 996:
            var v45 int32 = 65
            _ = v45
        case -652:
            var v46 string = (("hello" + "hello") + "")
            _ = v46
        default:
            v47 := GenS6{F0: 0}
            _ = v47
        }
        var v48 float64 = (3.14159 / (1e10 / ((3.14159 - ((3.14159 / 0.1) - -2.25)) + (0.0 + float64(-2.25)))))
        var v49 int32 = 2147483647
        _ = v48
        _ = v49
    } else {
        var v50 bool = true
        _ = v50
    }
    fmt.Println(v4, v44, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v44
}
