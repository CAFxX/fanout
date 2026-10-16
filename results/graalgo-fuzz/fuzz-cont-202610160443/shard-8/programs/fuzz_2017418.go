package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int8; F1 int16 }

func main() {
    var v1 uint8 = 112
    var v2 uint64 = 0
    var v3 float32 = 0.0
    v4 := []int{}
    v5 := []string{"", "\n\t", ("hello" + "\n\t"), "a"}
    v6 := []uint32{0}
    var v7 float64 = ((float64(((v3 - 0.0) + (float32(((-2.25 / 1.5) * ((0.0 * 1e10) * (-2.25 * 0.1)))) - 1e10))) + (3.14159 * (1.5 - (((1.5 - (3.14159 / 3.14159)) / 3.14159) * ((float64(-2.25) - 3.14159) / (0.0 - 1e10)))))) + 3.14159)
    switch "" {
    case "a":
        v2 += (v2 >> ((615) & 63))
    }
    v8 := GenS1{F0: -9223372036854775808}
    if false {
        if (!(true || (32767 < 0))) {
            if (!true) {
                fmt.Println(v2, v1)
                if false {
                    fmt.Println(v1, v1, v2)
                    fmt.Println(v2, v3, v3)
                    switch 127 {
                    case 128:
                        var v9 int = 287
                        _ = v9
                    case -61:
                        v7 += v7
                    default:
                        if (false || false) {
                            if false {
                                if (!(1 > 1)) {
                                    var v10 float64 = 0.0
                                    var v11 int = 32767
                                    switch ("a" + "foo bar") {
                                    case "a":
                                        fmt.Println(v3, v7)
                                    case "hello":
                                        if true {
                                            for fi12 := 0; fi12 < 20; fi12++ {
                                                for fi13 := 0; fi13 < 14; fi13++ {
                                                    v3 = 0.1
                                                    _ = fi13
                                                }
                                                _ = fi12
                                            }
                                            v11 = v11
                                        }
                                    default:
                                        switch 9223372036854775807 {
                                        case -9223372036854775808:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case 31:
                                            var v14 uint = 657
                                            _ = v14
                                        case 1:
                                            if len(v6) > 0 {
                                                v6[0] = 42
                                            }
                                        default:
                                            fmt.Println(v7)
                                        }
                                    }
                                    _ = v10
                                    _ = v11
                                }
                                var v15 bool = ("hello" == "foo bar")
                                switch ("a" + ("" + "")) {
                                case "":
                                    var v16 bool = ((v1 * (v1 ^ v1)) > (v1 % (((v1 % ((42) | 1))) | 1)))
                                    _ = v16
                                default:
                                    switch 475 {
                                    case 32768:
                                        switch 853 {
                                        case -746:
                                            if false {
                                                v2 += v2
                                            }
                                        case -1:
                                            if (!v15) {
                                                fmt.Println(v15, v3, v1)
                                            }
                                        case 32767:
                                            fmt.Println(v7)
                                        default:
                                            v7 += v7
                                        }
                                    case 4294967295:
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(v3)
                                        case "foo bar":
                                            v17 := GenS1{F0: 0}
                                            _ = v17
                                        case "\n\t":
                                            var v18 uint32 = 1
                                            _ = v18
                                        }
                                    }
                                }
                                _ = v15
                            }
                            v19 := GenS1{F0: -669}
                            _ = v19
                        } else {
                            var v20 int16 = 1000
                            if true {
                                v6 = append(v6, 481)
                                v2 -= v2
                            }
                            _ = v20
                        }
                    }
                }
            }
            if (-1 == 0) {
                v21 := GenS1{F0: 127}
                v22 := GenS2{F0: 844, F1: false}
                _ = v21
                _ = v22
            }
        } else {
            fmt.Println(v3)
            fmt.Println(v3)
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 14; fi24++ {
                if ((v1 + v1) > v1) {
                    v7 -= v7
                }
                var v25 uint = 4294967295
                v2 = (^(v2))
                _ = fi24
                _ = v25
            }
            v7 = ((0.1 - v7) - 0.1)
            v1 = (^(v1))
            _ = fi23
        }
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        if false {
            var v27 uint16 = 42
            var v28 uint16 = v27
            v27 -= v28
            _ = v27
            _ = v28
        } else {
            var v29 int64 = -9223372036854775808
            v30 := GenS3{F0: 0, F1: 0}
            if len(v5) > 0 {
                v5[0] = "hello"
            }
            _ = v29
            _ = v30
        }
        if true {
            var v31 uint64 = (v2 << ((42) & 63))
            v1 += (v1 / (((v1 * (v1 % ((v1) | 1)))) | 1))
            _ = len(v5)
            _ = v31
        } else {
            for fi32 := 0; fi32 < 4; fi32++ {
                v33 := GenS4{F0: -32768, F1: 0, F2: (v2 * (v2 - (v2 / ((1) | 1))))}
                _ = fi32
                _ = v33
            }
            _ = len(v6)
        }
        _ = fi26
    }
    switch -64 {
    case 42:
        if true {
            fmt.Println(v7, v7)
            if (((true || ((v2 != v2) && true)) && (32767 == -62)) && (!false)) {
                fmt.Println(v3, v7)
            }
        } else {
            var v34 int = -978
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 int64 = 232
                _ = fi35
                _ = v36
            }
            if true {
                v34 += (v34 / ((v34) | 1))
                v37 := GenS2{F0: 65535, F1: (((!false) && false) || ((v2 <= (^(v2))) && false))}
                _ = v37
            } else {
                for fi38 := 0; fi38 < 16; fi38++ {
                    v34 += (v34 / ((fi38) | 1))
                    v5 = append(v5, ("" + "\n\t"))
                    var v39 bool = (!true)
                    _ = fi38
                    _ = v39
                }
            }
            _ = v34
        }
    case 1:
        fmt.Println(v7, v2, v7)
    case -1:
        var v40 bool = (false || false)
        _ = v40
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v2 += 18446744073709551615
    var v41 int64 = 42
    v2 -= v2
    v1 += v1
    var v42 uint = 282
    var v43 int64 = v41
    v3 -= v3
    var v44 uint16 = 892
    v45 := GenS5{F0: 15, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
