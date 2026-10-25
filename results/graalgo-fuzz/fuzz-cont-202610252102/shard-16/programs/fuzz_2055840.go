package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int8; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int16; F2 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 uint64 }

func main() {
    var v1 int8 = -128
    var v2 uint64 = 2
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v5 uint = 302
    v6 := []string{("foo bar" + ("a" + "hello")), "", "", "\n\t"}
    var v7 uint32 = 0
    var v8 int16 = 104
    v9 := GenS1{F0: (v5 << ((v5) & 63))}
    switch ("foo bar" + "a") {
    case "\n\t":
        for fi10 := 0; fi10 < 8; fi10++ {
            if (255 < 128) {
                _ = len(v6)
            }
            var v11 uint = v5
            _ = fi10
            _ = v11
        }
    case "foo bar":
        delete(v4, 1)
    default:
        v1 -= v1
    }
    var v12 uint32 = (v7 << (((v5 & 427)) & 63))
    v2 = (^(v2))
    v12 = 2147483647
    var v13 float64 = (0.0 * (1e10 / float64(((0.0 + 3.14159) / (1.5 / 1.0)))))
    v14 := v4[42]
    fmt.Println(v5)
    fmt.Println(v5, v7)
    if false {
        if (v13 > ((((-2.25 * ((v13 * (-2.25 + (float64(1e10) / float64(0.1)))) / v13)) / v13) / ((v13 / v13) * float64(1e10))) - (0.0 - v13))) {
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 20; fi16++ {
                    var v17 uint32 = 2147483647
                    _ = fi16
                    _ = v17
                }
                var v18 int16 = -1
                _ = fi15
                _ = v18
            }
            v19 := GenS1{F0: 694}
            fmt.Println(v7, v5, v2)
            _ = v19
        } else {
            delete(v4, -1)
            v20 := GenS2{F0: 1}
            _ = v20
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            if (-2147483648 >= fi21) {
                v22 := v3[(v14 % (((v7 << (((v5 % ((42) | 1))) & 63))) | 1))]
                _ = len(v6)
                _ = v22
            } else {
                switch (fi21 % ((int(v1)) | 1)) {
                case 127:
                    switch (fi21 % (((-(fi21))) | 1)) {
                    case 2147483647:
                        v23 := GenS2{F0: 2}
                        _ = v23
                    case -2147483648:
                        for fi24 := 0; fi24 < 13; fi24++ {
                            v13 += (float64(3.14159) + v13)
                            if false {
                                switch 397 {
                                case 32767:
                                    var v25 int64 = 42
                                    _ = v25
                                case -311:
                                    v5 -= (v5 + v5)
                                case -9223372036854775808:
                                    v26 := GenS3{F0: -2147483648, F1: (-(v1)), F2: 32767}
                                    _ = v26
                                default:
                                    if false {
                                        switch (-(fi24)) {
                                        case 32768:
                                            if true {
                                                v27 := GenS4{F0: true}
                                                switch 501 {
                                                case 2147483647:
                                                    v28 := GenS4{F0: false}
                                                    _ = v28
                                                case 32767:
                                                    if true {
                                                        if ("foo bar" == "\n\t") {
                                                            if true {
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v29 int16 = (v8 >> (((v5 >> (((^(v5))) & 63))) & 63))
                                                                    _ = v29
                                                                case "a":
                                                                    var v30 uint8 = 128
                                                                    _ = v30
                                                                case "foo bar":
                                                                    v3[(^(v12))] = (true && false)
                                                                }
                                                            }
                                                            fmt.Println(v2, v7, v7)
                                                            fmt.Println(v1)
                                                            fmt.Println(fi24, fi24, fi24)
                                                        }
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            v12 -= (v7 | (v12 + v14))
                                                        }
                                                        var v31 uint = (v5 + (v5 & 42))
                                                        _ = v31
                                                    }
                                                case 1:
                                                    var v32 int64 = 870
                                                    _ = v32
                                                }
                                                _ = v27
                                            }
                                        case 9223372036854775807:
                                            if ((!(42 >= 0)) || (-844 > -352)) {
                                                if true {
                                                    var v33 int16 = v8
                                                    if ((!true) && true) {
                                                        var v34 int64 = 9223372036854775807
                                                        _ = v34
                                                    }
                                                    fmt.Println(v33, v14, v5)
                                                    _ = v33
                                                } else {
                                                    var v35 int64 = -169
                                                    _ = v35
                                                }
                                                if ((v8 / (((v8 & int16(v2))) | 1)) < v8) {
                                                    v12 = (v12 + 269)
                                                    var v36 uint64 = (v2 * 12)
                                                    _ = v36
                                                }
                                            }
                                        case -3:
                                            if false {
                                                fmt.Println(v7, v12)
                                            } else {
                                                v37 := GenS5{F0: 42}
                                                _ = v37
                                            }
                                        default:
                                            if false {
                                                v3[v7] = ("foo bar" >= "foo bar")
                                                fmt.Println(v5, v7)
                                            }
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = (("\n\t" + ((("hello" + ("foo bar" + "hello")) + "hello") + ("hello" + "a"))) + ("hello" + "\n\t"))
                                        }
                                        var v38 float64 = (v13 - v13)
                                        _ = v38
                                    } else {
                                        delete(v4, fi21)
                                        var v39 float32 = 1e10
                                        fmt.Println(v8)
                                        _ = v39
                                    }
                                }
                                if true {
                                    fmt.Println(v2, fi24)
                                }
                                v5 -= (v5 & 65535)
                            }
                            _ = fi24
                        }
                    default:
                        var v40 int8 = 111
                        _ = v40
                    }
                case 162:
                    v8 = v8
                case 32768:
                    for fi41 := 0; fi41 < 1; fi41++ {
                        if true {
                            if true {
                                fmt.Println(v14, v2, fi21)
                                _ = len(v6)
                                fmt.Println(v12)
                            } else {
                                if false {
                                    switch ("foo bar" + "") {
                                    case "hello":
                                        var v42 uint64 = v2
                                        _ = v42
                                    case "foo bar":
                                        var v43 uint = (v5 - 4294967295)
                                        _ = v43
                                    }
                                } else {
                                    v5 = (v5 / ((v5) | 1))
                                    v14 -= 2147483648
                                }
                                if true {
                                    var v44 uint64 = (v2 - v2)
                                    _ = len(v3)
                                    if (!(v14 != (v12 >> ((uint(v2)) & 63)))) {
                                        if true {
                                            v1 += v1
                                            v45 := GenS1{F0: (v5 / ((v5) | 1))}
                                            _ = v45
                                        }
                                        if (v2 <= v2) {
                                            v2 += 43
                                        }
                                        v6 = append(v6, "hello")
                                    }
                                    _ = v44
                                }
                                var v46 bool = true
                                switch (fi41 | 9223372036854775807) {
                                case -1:
                                    v47 := GenS6{F0: ("hello" + ""), F1: (v8 + v8), F2: -178}
                                    _ = v47
                                case -857:
                                    if len(v6) > 0 {
                                        v6[0] = "a"
                                    }
                                case 32768:
                                    switch ("a" + "hello") {
                                    case "":
                                        fmt.Println(v1, v8, v7)
                                    case "\n\t":
                                        v5 = (v5 % ((uint(v7)) | 1))
                                    default:
                                        var v48 int8 = v1
                                        _ = v48
                                    }
                                }
                                switch "a" {
                                case "hello":
                                    v49 := GenS7{F0: uint(v2)}
                                    _ = v49
                                case "foo bar":
                                    if ((v2 & (v2 - (v2 + 18446744073709551615))) > uint64(v5)) {
                                        fmt.Println(v7)
                                        v12 = 0
                                        var v50 int = fi41
                                        _ = v50
                                    } else {
                                        var v51 float32 = (-2.25 * (((0.1 / (3.14159 / 3.14159)) * 1e10) + 1.5))
                                        _ = v51
                                    }
                                case "":
                                    var v52 uint32 = (v14 ^ (^(v7)))
                                    _ = v52
                                default:
                                    v53 := GenS8{F0: -2147483648, F1: 127}
                                    _ = v53
                                }
                                _ = v46
                            }
                            v14 += v12
                        } else {
                            var v54 float64 = float64(1e10)
                            var v55 uint8 = 21
                            _ = v54
                            _ = v55
                        }
                        v12 = v7
                        var v56 int32 = -2147483648
                        _ = fi41
                        _ = v56
                    }
                default:
                    for fi57 := 0; fi57 < 15; fi57++ {
                        v58 := GenS4{F0: ((((v13 + 0.0) + (v13 * ((-2.25 + (v13 - v13)) * (v13 - v13)))) - 0.0) <= (v13 * 1.5))}
                        fmt.Println(fi21, v12, v5)
                        _ = fi57
                        _ = v58
                    }
                }
                var v59 uint64 = (v2 >> (((^(v5))) & 63))
                var v60 uint32 = (v7 | v12)
                _ = v59
                _ = v60
            }
            v61 := GenS9{F0: v2}
            var v62 uint = 1
            _ = fi21
            _ = v61
            _ = v62
        }
    } else {
        var v63 bool = true
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
