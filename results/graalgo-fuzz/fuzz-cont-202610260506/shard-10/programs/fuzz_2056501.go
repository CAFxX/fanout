package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 uint8; F2 int16 }

func main() {
    v1 := []string{}
    var v2 uint64 = 857
    v3 := []int{-1, 550, -9223372036854775808}
    v4 := make(map[int]uint32)
    var v5 int16 = -32768
    v6 := make(map[int]uint32)
    var v7 int8 = 42
    var v8 string = ""
    switch -9223372036854775808 {
    case -986:
        _ = len(v6)
    case 127:
        fmt.Println(v5, v2)
    case -2147483648:
        switch v8 {
        case "":
            for fi9 := 0; fi9 < 15; fi9++ {
                v5 -= v5
                v10 := GenS1{F0: v7, F1: 42}
                _ = fi9
                _ = v10
            }
        case "hello":
            var v11 int = int(v2)
            _ = v11
        case "a":
            v7 = -125
        default:
            var v12 uint16 = 2
            _ = v12
        }
    }
    var v13 string = "hello"
    var v14 uint64 = (v2 << ((42) & 63))
    fmt.Println(v5, v14)
    fmt.Println(v2)
    if false {
        for fi15 := 0; fi15 < 6; fi15++ {
            switch v13 {
            case "hello":
                var v16 bool = (!true)
                _ = v16
            }
            _ = fi15
        }
        if (!(552 <= 471)) {
            v5 = (v5 / (((v5 >> ((0) & 63))) | 1))
        } else {
            var v17 float64 = 0.0
            _ = v17
        }
    } else {
        var v18 uint8 = 42
        var v19 string = v13
        _ = v18
        _ = v19
    }
    if true {
        var v20 uint = 255
        v21 := GenS1{F0: (^(v7)), F1: 418}
        _ = v20
        _ = v21
    } else {
        fmt.Println(v2, v7, v13)
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                if false {
                    var v24 uint32 = 2147483648
                    _ = v24
                }
                _ = fi23
            }
            _ = fi22
        }
    }
    v25 := GenS2{F0: (^(v5)), F1: v14}
    if true {
        v6[-274] = 622
        for fi26 := 0; fi26 < 4; fi26++ {
            if ((v7 - (v7 % ((v7) | 1))) > -99) {
                var v27 uint8 = 128
                _ = v27
            }
            _ = fi26
        }
    } else {
        var v28 uint32 = 2147483648
        var v29 bool = (!(1e10 < ((1.5 / 1e10) * ((1.5 - 3.14159) / ((1.5 / (((float32((0.1 + float64(0.1))) / -2.25) + (((((1.5 * float32((0.1 / -2.25))) / 1.5) + (3.14159 * (3.14159 / 1.5))) + float32(1.5)) / float32((((-2.25 * (-2.25 + (0.0 + -2.25))) - 1.5) + 1.5)))) * float32((-2.25 + float64(0.0))))) * -2.25)))))
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        switch "" {
        case "hello":
            var v31 uint8 = 127
            _ = v31
        }
        fmt.Println(v14, v8, v8)
        switch "hello" {
        case "foo bar":
            var v32 string = (v13 + v8)
            _ = v32
        }
        _ = fi30
    }
    v14 = (v14 ^ (v14 + v2))
    v33 := GenS2{F0: (v5 << ((42) & 63)), F1: (^(v2))}
    fmt.Println(v13)
    for fi34 := 0; fi34 < 7; fi34++ {
        v35 := GenS1{F0: 127, F1: 0}
        _ = len(v4)
        if (!false) {
            if false {
                if len(v1) > 0 {
                    v1[0] = ((v13 + "") + v13)
                }
                switch ("" + ((v8 + ("a" + v13)) + v8)) {
                case "foo bar":
                    var v36 uint64 = (v14 * (v2 / ((1) | 1)))
                    _ = v36
                default:
                    if (!true) {
                        switch fi34 {
                        case -830:
                            v37 := v6[(fi34 + fi34)]
                            _ = v37
                        case 1:
                            if len(v1) > 0 {
                                v1[0] = v8
                            }
                        }
                    } else {
                        var v38 bool = true
                        var v39 int = fi34
                        switch v39 {
                        case 4294967295:
                            var v40 uint8 = 127
                            _ = v40
                        case -2147483648:
                            switch "a" {
                            case "foo bar":
                                v41 := GenS2{F0: v5, F1: 18446744073709551615}
                                _ = v41
                            case "a":
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    v7 = (v7 << ((255) & 63))
                                    var v43 bool = v38
                                    var v44 uint16 = 1
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                }
                            }
                        default:
                            fmt.Println(v39)
                        }
                        _ = v38
                        _ = v39
                    }
                }
                switch (fi34 << ((413) & 63)) {
                case 127:
                    for fi45 := 0; fi45 < 17; fi45++ {
                        var v46 uint8 = 151
                        _ = fi45
                        _ = v46
                    }
                default:
                    var v47 int32 = 590
                    _ = v47
                }
            } else {
                v3 = append(v3, fi34)
            }
            if (false && true) {
                v48 := GenS3{F0: fi34, F1: 0}
                fmt.Println(v2, v8)
                _ = v48
            } else {
                var v49 float32 = 0.1
                if true {
                    var v50 float32 = v49
                    switch (v8 + ((v8 + "") + "a")) {
                    case "hello":
                        var v51 string = "foo bar"
                        _ = v51
                    default:
                        if len(v1) > 0 {
                            v1[0] = ((v8 + "") + (v8 + "hello"))
                        }
                    }
                    for fi52 := 0; fi52 < 3; fi52++ {
                        if true {
                            if (false && true) {
                                switch 831 {
                                case -2147483648:
                                    var v53 uint8 = 42
                                    _ = v53
                                case 252:
                                    fmt.Println(v5, v5)
                                default:
                                    if false {
                                        v54 := GenS3{F0: -2147483648, F1: 0}
                                        v55 := GenS4{F0: 0}
                                        switch (fi34 * (fi34 << ((456) & 63))) {
                                        case 782:
                                            fmt.Println(v7, v8, v13)
                                        case 4:
                                            var v56 uint64 = v14
                                            _ = v56
                                        default:
                                            delete(v4, (fi52 / (((fi34 << ((65535) & 63))) | 1)))
                                            var v57 int16 = (v5 - (^(v5)))
                                            _ = v57
                                        }
                                        _ = v54
                                        _ = v55
                                    }
                                }
                                v58 := GenS5{F0: 32767}
                                switch 123 {
                                case 42:
                                    switch ((v8 + ((v13 + v8) + (v13 + "hello"))) + ("hello" + v8)) {
                                    case "foo bar":
                                        var v59 bool = (!true)
                                        _ = v59
                                    }
                                case 0:
                                    fmt.Println(v7)
                                }
                                _ = v58
                            }
                        }
                        v50 -= 0.1
                        v60 := GenS6{F0: (v7 / ((55) | 1)), F1: uint8(v14), F2: v5}
                        _ = fi52
                        _ = v60
                    }
                    _ = v50
                } else {
                    var v61 bool = true
                    v7 -= v7
                    v14 += v14
                    _ = v61
                }
                _ = v49
            }
        } else {
            for fi62 := 0; fi62 < 6; fi62++ {
                v63 := GenS3{F0: fi62, F1: 128}
                var v64 float64 = ((1.5 + 3.14159) * (float64(1e10) + (((0.0 / -2.25) * ((1e10 - -2.25) / 3.14159)) + (1.5 + 3.14159))))
                fmt.Println(v13)
                _ = fi62
                _ = v63
                _ = v64
            }
            v2 += (v14 & 2)
            v8 = (v8 + v8)
        }
        _ = fi34
        _ = v35
    }
    v65 := GenS4{F0: uint16(v2)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v7
    _ = v8
}
