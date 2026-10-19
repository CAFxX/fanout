package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int8; F1 string; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }

func main() {
    var v1 uint = 622
    var v2 int64 = 42
    var v3 int8 = 100
    var v4 int = 1
    v2 += v2
    fmt.Println(v1, v2)
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{}
        var v7 string = (("foo bar" + "\n\t") + "hello")
        var v8 bool = (-237 > 42)
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if (255 >= 0) {
        v4 += 128
        v9 := GenS1{F0: 2}
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint16 = 0
        for fi12 := 0; fi12 < 7; fi12++ {
            v11 = 0
            v13 := GenS1{F0: 2147483648}
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    if true {
        v14 := []int{(v4 / ((-1) | 1)), 4294967295, (v4 + v4)}
        var v15 int = (v4 ^ v4)
        if false {
            var v16 int16 = 5
            if (42 != (v16 << ((v1) & 63))) {
                v15 -= int(v3)
            } else {
                if true {
                    var v17 string = "foo bar"
                    var v18 int16 = (v16 ^ int16(v1))
                    _ = v17
                    _ = v18
                } else {
                    if false {
                        if (true || false) {
                            var v19 uint = v1
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 14; fi20++ {
                                v15 += (fi20 << ((v1) & 63))
                                if (!(2 < 1)) {
                                    if false {
                                        var v21 string = ""
                                        fmt.Println(v3, v4)
                                        fmt.Println(v2)
                                        _ = v21
                                    }
                                }
                                if (!true) {
                                    v22 := GenS1{F0: 65535}
                                    _ = v22
                                }
                                _ = fi20
                            }
                        }
                        switch "foo bar" {
                        case "a":
                            if len(v14) > 0 {
                                v14[0] = v15
                            }
                        case "\n\t":
                            if false {
                                if false {
                                    switch "\n\t" {
                                    case "hello":
                                        if false {
                                            switch "\n\t" {
                                            case "hello":
                                                var v23 uint = 42
                                                _ = v23
                                            default:
                                                for fi24 := 0; fi24 < 15; fi24++ {
                                                    switch (("a" + "hello") + "foo bar") {
                                                    case "foo bar":
                                                        var v25 int16 = int16(v2)
                                                        _ = v25
                                                    default:
                                                        if false {
                                                            if true {
                                                                if false {
                                                                    var v26 string = "hello"
                                                                    _ = v26
                                                                } else {
                                                                    if (((!false) && false) && true) {
                                                                        switch 2147483647 {
                                                                        case -2147483648:
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                v15 += (v15 & (v15 + 0))
                                                                            case "\n\t":
                                                                                v3 = v3
                                                                            case "hello":
                                                                                v27 := GenS2{F0: 61, F1: 65535}
                                                                                _ = v27
                                                                            }
                                                                        case 9223372036854775807:
                                                                            fmt.Println(v15)
                                                                        default:
                                                                            fmt.Println(v1)
                                                                        }
                                                                    } else {
                                                                        for fi28 := 0; fi28 < 9; fi28++ {
                                                                            if len(v14) >= 2 {
                                                                                v14 = v14[1:len(v14)]
                                                                            }
                                                                            v14 = append(v14, v4)
                                                                            var v29 uint16 = 147
                                                                            _ = fi28
                                                                            _ = v29
                                                                        }
                                                                    }
                                                                    if len(v14) > 0 {
                                                                        v14[0] = -348
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v15 += v4
                                                    _ = fi24
                                                }
                                            }
                                            switch (((("" + "a") + "hello") + "") + "") {
                                            case "\n\t":
                                                for fi30 := 0; fi30 < 19; fi30++ {
                                                    if ((!(false && (float64(3.14159) > 1.5))) && (((v1 % ((v1) | 1)) == 2) && true)) {
                                                        v16 = (v16 * (v16 >> (((^(v1))) & 63)))
                                                    }
                                                    _ = fi30
                                                }
                                            default:
                                                v31 := GenS3{F0: -1, F1: ((("" + "\n\t") + ("" + ("hello" + (("a" + "a") + "foo bar")))) + "hello"), F2: 42}
                                                _ = v31
                                            }
                                            v32 := GenS4{F0: 18446744073709551615}
                                            _ = v32
                                        }
                                    }
                                } else {
                                    v15 = v15
                                    v33 := GenS1{F0: 89}
                                    _ = v33
                                }
                                if (65535 == v1) {
                                    if false {
                                        v15 += v4
                                        var v34 uint8 = 42
                                        v35 := GenS1{F0: 65535}
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v4 -= 4294967295
                                        var v36 uint64 = 1
                                        var v37 int64 = (v2 | (v2 & 924))
                                        _ = v36
                                        _ = v37
                                    }
                                    v4 += 127
                                    if (v15 > (v15 % (((v15 & v15)) | 1))) {
                                        for fi38 := 0; fi38 < 3; fi38++ {
                                            fmt.Println(v2, v16)
                                            _ = fi38
                                        }
                                        v14 = append(v14, (v15 / ((v15) | 1)))
                                        v2 = (v2 | 9223372036854775807)
                                    }
                                }
                                var v39 int16 = (v16 | (v16 / (((v16 / ((int16(v3)) | 1))) | 1)))
                                _ = v39
                            }
                        case "":
                            var v40 uint64 = 18446744073709551615
                            _ = v40
                        }
                    }
                    v1 = 65535
                }
                v15 += v15
            }
            _ = v16
        }
        _ = v14
        _ = v15
    }
    v41 := []int{}
    if len(v41) >= 2 {
        v41 = v41[1:len(v41)]
    }
    for fi42 := 0; fi42 < 16; fi42++ {
        fmt.Println(v4, v3)
        v3 = -68
        _ = fi42
    }
    fmt.Println(v1)
    if len(v41) > 0 {
        v41[0] = (v4 % ((v4) | 1))
    }
    var v43 int8 = (v3 - (^(v3)))
    var v44 uint8 = 0
    if (0.1 > 1e10) {
        fmt.Println(v1, v44, v2)
    }
    for fi45 := 0; fi45 < 6; fi45++ {
        v46 := GenS4{F0: 2}
        v47 := GenS3{F0: v3, F1: ("hello" + ("\n\t" + "")), F2: v44}
        _ = fi45
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 18; fi48++ {
        for fi49 := 0; fi49 < 13; fi49++ {
            var v50 float64 = 3.14159
            v50 += 0.0
            _ = fi49
            _ = v50
        }
        switch v4 {
        case 1:
            if (false || false) {
                v44 = 219
                var v51 int8 = v43
                _ = v51
            }
        case -2147483648:
            v44 = 0
        case -9223372036854775808:
            for fi52 := 0; fi52 < 1; fi52++ {
                v53 := GenS4{F0: 663}
                v54 := GenS5{F0: ((true && ((v2 - (v2 - v2)) == v2)) && false)}
                v1 -= 298
                _ = fi52
                _ = v53
                _ = v54
            }
        default:
            if len(v41) > 0 {
                v41[0] = 62
            }
        }
        if len(v41) > 0 {
            v41[0] = 32767
        }
        _ = fi48
    }
    v2 -= v2
    if true {
        var v55 int16 = 42
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v44
}
