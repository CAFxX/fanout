package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int; F2 int }
type GenS3 struct { F0 bool; F1 int8; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{280, 78}
    var v3 int = -191
    var v4 uint16 = 2
    var v5 uint = 783
    var v6 int = (v3 + (v3 * (v3 * v3)))
    var v7 uint64 = 825
    switch "" {
    case "foo bar":
        var v8 uint8 = 255
        _ = v8
    case "":
        if (32768 > v3) {
            _ = len(v1)
        }
    case "\n\t":
        v9 := GenS1{F0: (0 >= (v4 & 222))}
        _ = v9
    }
    switch "hello" {
    case "\n\t":
        v10 := GenS1{F0: ((v5 * uint(v7)) == (v5 << ((v5) & 63)))}
        _ = v10
    case "":
        v5 += v5
    case "hello":
        fmt.Println(v4, v6)
    }
    var v11 float64 = (1e10 / (3.14159 - 0.1))
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(v5, v11)
            _ = fi13
        }
        switch v3 {
        case 42:
            v11 -= -2.25
        case 2147483647:
            switch "" {
            case "\n\t":
                var v14 int64 = -960
                _ = v14
            case "a":
                if (9223372036854775807 <= 9223372036854775807) {
                    v11 += ((v11 / (v11 - 3.14159)) / float64(1.5))
                } else {
                    var v15 int8 = 105
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    var v16 uint64 = uint64(v5)
                    _ = v15
                    _ = v16
                }
            default:
                v17 := GenS1{F0: false}
                _ = v17
            }
        default:
            if false {
                if (true || false) {
                    var v18 bool = (false || ((^(v5)) >= (v5 >> ((uint(v4)) & 63))))
                    _ = v18
                }
                _ = len(v2)
                delete(v1, "")
            } else {
                var v19 int16 = -436
                v6 += (v6 ^ fi12)
                switch "a" {
                case "foo bar":
                    var v20 bool = ((v4 & (v4 % (((v4 ^ v4)) | 1))) < (v4 & 925))
                    _ = v20
                case "hello":
                    fmt.Println(v3)
                case "a":
                    for fi21 := 0; fi21 < 3; fi21++ {
                        var v22 uint32 = 363
                        v3 = (fi21 >> (((v5 & 0)) & 63))
                        var v23 uint64 = (v7 / ((1) | 1))
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
                _ = v19
            }
        }
        _ = fi12
    }
    var v24 int64 = -572
    v25 := GenS2{F0: v4, F1: (v6 * -109), F2: (v3 + (v3 & v3))}
    fmt.Println(v7)
    var v26 float64 = 3.14159
    v1["a"] = uint32(v5)
    if true {
        v26 = (v26 * v26)
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        fmt.Println(v3)
    }
    var v27 int = 127
    if (true || (true || (4294967295 <= v7))) {
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := v1["a"]
            v30 := GenS2{F0: 2, F1: (v27 - v6), F2: v6}
            for fi31 := 0; fi31 < 9; fi31++ {
                switch -9223372036854775808 {
                case -351:
                    v6 += (v6 * (v27 & v6))
                }
                v32 := v1[(("foo bar" + "") + "\n\t")]
                v1["hello"] = (v29 | (^(v29)))
                _ = fi31
                _ = v32
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        v24 = v24
        v33 := GenS2{F0: 0, F1: v6, F2: (v3 ^ 127)}
        _ = v33
    } else {
        for fi34 := 0; fi34 < 8; fi34++ {
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            v35 := GenS1{F0: false}
            _ = fi34
            _ = v35
        }
        for fi36 := 0; fi36 < 9; fi36++ {
            for fi37 := 0; fi37 < 18; fi37++ {
                _ = len(v2)
                v24 = int64(v3)
                switch "" {
                case "hello":
                    v38 := GenS1{F0: (!false)}
                    v39 := GenS3{F0: (42 == 0), F1: 0, F2: "\n\t"}
                    fmt.Println(v26, v5, v7)
                    _ = v38
                    _ = v39
                case "\n\t":
                    if (600 == (v7 * v7)) {
                        var v40 uint = (v5 % (((v5 / (((v5 ^ v5)) | 1))) | 1))
                        fmt.Println(v40)
                        _ = v40
                    }
                case "":
                    v4 = v4
                }
                if (1 > 65535) {
                    if false {
                        var v41 string = ("foo bar" + "foo bar")
                        _ = v41
                    } else {
                        if false {
                            fmt.Println(v24, v7, fi36)
                            fmt.Println(v27, v4, fi36)
                            v6 = (^(fi36))
                            var v42 uint = v5
                            switch "\n\t" {
                            case "foo bar":
                                v26 -= 1.5
                            case "":
                                var v43 int64 = v24
                                _ = v43
                            case "\n\t":
                                var v44 uint32 = 2
                                _ = v44
                            default:
                                v42 -= uint(v7)
                            }
                            v42 -= v5
                            _ = v42
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch v3 {
                    case 672:
                        if false {
                            var v45 float32 = (float32(v26) * -2.25)
                            switch -28 {
                            case 811:
                                fmt.Println(v5)
                            case 0:
                                fmt.Println(v27, fi37, v5)
                            }
                            switch 127 {
                            case -9223372036854775808:
                                v2 = append(v2, uint32(v7))
                            default:
                                v46 := GenS1{F0: false}
                                _ = v46
                            }
                            _ = v45
                        } else {
                            v47 := GenS1{F0: false}
                            v27 += v3
                            var v48 uint8 = 59
                            _ = v47
                            _ = v48
                        }
                    case -9223372036854775808:
                        switch (-(v3)) {
                        case -182:
                            v49 := GenS4{F0: -42}
                            _ = v49
                        default:
                            v11 = -2.25
                        }
                    default:
                        fmt.Println(v26, fi36, v4)
                    }
                } else {
                    switch fi36 {
                    case 32768:
                        var v50 float32 = (1.5 + 0.0)
                        _ = v50
                    }
                    switch (v3 - 4294967295) {
                    case 1:
                        if ((fi36 > (v3 >> ((v5) & 63))) && false) {
                            v7 -= (^(v7))
                        } else {
                            if false {
                                v51 := GenS3{F0: (false || false), F1: 42, F2: ""}
                                _ = len(v1)
                                _ = v51
                            } else {
                                fmt.Println(v7)
                                v52 := GenS5{F0: 18446744073709551615}
                                _ = v52
                            }
                            var v53 bool = false
                            _ = v53
                        }
                    case -9223372036854775808:
                        v54 := GenS3{F0: (!false), F1: 127, F2: "foo bar"}
                        _ = v54
                    case 960:
                        var v55 int32 = 724
                        _ = v55
                    }
                }
                if ("a" <= "hello") {
                    var v56 uint64 = 0
                    var v57 float64 = -2.25
                    v56 = (v56 % ((v56) | 1))
                    _ = v56
                    _ = v57
                }
                v58 := GenS6{F0: (v5 / ((v5) | 1)), F1: 127}
                _ = fi37
                _ = v58
            }
            _ = fi36
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v59 int64 = 0
    v5 -= (v5 / (((v5 / ((18446744073709551615) | 1))) | 1))
    for fi60 := 0; fi60 < 9; fi60++ {
        v4 += v4
        _ = fi60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v7
}
