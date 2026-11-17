package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint64; F1 int32; F2 int16 }
type GenS5 struct { F0 uint8; F1 bool; F2 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int; F1 uint; F2 uint8 }
func genFn0() int {
    var v5 int = -782
    switch "foo bar" {
    case "foo bar":
        var v6 uint16 = 65535
        _ = v6
    default:
        var v7 bool = true
        _ = v7
    }
    var v8 int32 = 42
    switch (v5 * int(v8)) {
    case 32767:
        if ("" <= "foo bar") {
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(v5)
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := GenS1{F0: "\n\t"}
                    _ = fi10
                    _ = v11
                }
                v5 -= 32768
                _ = fi9
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint64 = 332
                _ = fi12
                _ = v13
            }
        }
    }
    _ = v5
    _ = v8
    return (v5 / ((-700) | 1))
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int16 = -658
    var v3 uint16 = 756
    var v4 float64 = (((3.14159 - (((((3.14159 / -2.25) * 1e10) * 3.14159) - float64(0.1)) - 0.0)) * 0.1) + (-2.25 + 0.0))
    for fi14 := 0; fi14 < 4; fi14++ {
        switch genFn0() {
        case 1:
            if (0 >= -361) {
                if (109 <= (^(v3))) {
                    v15 := GenS2{F0: 127, F1: v2}
                    if true {
                        var v16 uint32 = (v1 ^ (v1 >> ((65535) & 63)))
                        _ = v16
                    }
                    _ = v15
                }
                var v17 float64 = -2.25
                var v18 float32 = ((-2.25 - 1.5) * 0.1)
                _ = v17
                _ = v18
            }
        default:
            fmt.Println(v1, v4, v2)
        }
        fmt.Println(v4)
        v19 := []int{2147483647, (-(fi14)), (fi14 - fi14)}
        _ = fi14
        _ = v19
    }
    var v20 uint = 4294967295
    var v21 uint8 = 128
    v22 := []int{-1, -808, int(v1), 831}
    var v23 int8 = 100
    for fi24 := 0; fi24 < 5; fi24++ {
        fmt.Println(v3, fi24, v23)
        _ = fi24
    }
    var v25 int8 = (v23 & -1)
    v20 -= v20
    v21 = (v21 >> (((v20 >> ((v20) & 63))) & 63))
    if len(v22) > 0 {
        v22[0] = 2147483647
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        fmt.Println(v2, v25, v2)
        if (!(v23 == (-(v25)))) {
            v4 += -2.25
            var v27 int64 = 1
            var v28 string = ("" + "a")
            _ = v27
            _ = v28
        } else {
            if (-2.25 <= (0.0 / -2.25)) {
                switch (("hello" + "a") + "foo bar") {
                case "\n\t":
                    v29 := GenS3{F0: 998, F1: -215}
                    _ = v29
                }
                switch "foo bar" {
                case "foo bar":
                    for fi30 := 0; fi30 < 3; fi30++ {
                        v31 := GenS4{F0: 223, F1: 0, F2: (v2 / ((v2) | 1))}
                        if false {
                            v32 := GenS5{F0: (^(v21)), F1: false, F2: "foo bar"}
                            fmt.Println(v21, v2)
                            _ = v32
                        }
                        _ = fi30
                        _ = v31
                    }
                case "":
                    switch "a" {
                    case "":
                        for fi33 := 0; fi33 < 1; fi33++ {
                            var v34 int = int(v1)
                            _ = fi33
                            _ = v34
                        }
                    case "hello":
                        for fi35 := 0; fi35 < 19; fi35++ {
                            v36 := GenS3{F0: 529, F1: v2}
                            v37 := GenS2{F0: 0, F1: 1000}
                            v38 := GenS1{F0: ("hello" + "foo bar")}
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    case "a":
                        switch "" {
                        case "":
                            if (!true) {
                                var v39 int64 = 804
                                for fi40 := 0; fi40 < 11; fi40++ {
                                    var v41 string = "a"
                                    _ = fi40
                                    _ = v41
                                }
                                v22 = append(v22, fi26)
                                _ = v39
                            } else {
                                var v42 int16 = v2
                                v42 += (v2 | (^(v42)))
                                if true {
                                    var v43 bool = false
                                    v22 = append(v22, fi26)
                                    fmt.Println(v2)
                                    _ = v43
                                }
                                _ = v42
                            }
                        case "\n\t":
                            switch "a" {
                            case "foo bar":
                                v1 -= v1
                            case "":
                                fmt.Println(v2, v20)
                            case "hello":
                                if false {
                                    v44 := GenS3{F0: 4294967295, F1: -377}
                                    var v45 int8 = (-(v23))
                                    v46 := GenS6{F0: (v1 >> ((42) & 63))}
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    for fi47 := 0; fi47 < 10; fi47++ {
                                        var v48 uint32 = (^(v1))
                                        v49 := GenS2{F0: v21, F1: (v2 ^ (v2 & (v2 - v2)))}
                                        fmt.Println(fi26)
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                    var v50 int32 = 300
                                    v50 = v50
                                    _ = v50
                                }
                            }
                        }
                    }
                }
            }
            fmt.Println(v3, v25, v21)
            fmt.Println(v3)
        }
        _ = fi26
    }
    if false {
        fmt.Println(v23, v21, v20)
        v21 += v21
    } else {
        if false {
            v51 := GenS7{F0: genFn0(), F1: 33}
            fmt.Println(v23, v20)
            _ = v51
        } else {
            for fi52 := 0; fi52 < 14; fi52++ {
                v3 += (v3 << ((65535) & 63))
                _ = fi52
            }
            v22 = append(v22, 128)
            for fi53 := 0; fi53 < 14; fi53++ {
                if true {
                    for fi54 := 0; fi54 < 5; fi54++ {
                        if false {
                            var v55 int16 = (v2 - -32768)
                            v3 = (v3 * (v3 + (^(v3))))
                            var v56 uint64 = 1
                            var v57 uint = v20
                            _ = v55
                            _ = v56
                            _ = v57
                        } else {
                            v2 += -32768
                        }
                        fmt.Println(v2, v21, v20)
                        var v58 uint16 = v3
                        _ = fi54
                        _ = v58
                    }
                    if ((-2.25 / v4) != (3.14159 + ((((v4 - v4) * v4) + 1.5) / ((v4 * (v4 - 0.0)) - -2.25)))) {
                        if false {
                            fmt.Println(fi53)
                            if (!((v4 >= v4) || false)) {
                                var v59 int = fi53
                                _ = v59
                            }
                            if true {
                                var v60 bool = false
                                _ = v60
                            }
                        } else {
                            for fi61 := 0; fi61 < 1; fi61++ {
                                switch -663 {
                                case 42:
                                    var v62 uint64 = 2
                                    _ = v62
                                default:
                                    var v63 float32 = (0.0 / (-2.25 + (((1.5 + 0.0) * 1e10) * (3.14159 * 0.1))))
                                    _ = v63
                                }
                                v2 += v2
                                v22 = append(v22, fi61)
                                _ = fi61
                            }
                        }
                        var v64 int64 = 75
                        _ = len(v22)
                        _ = v64
                    } else {
                        if true {
                            if false {
                                for fi65 := 0; fi65 < 13; fi65++ {
                                    v66 := GenS3{F0: 4294967295, F1: (v2 - (v2 ^ 1))}
                                    v20 -= v20
                                    _ = fi65
                                    _ = v66
                                }
                                switch "foo bar" {
                                case "hello":
                                    v3 -= v3
                                case "a":
                                    v21 = 0
                                default:
                                    var v67 int16 = (v2 * -545)
                                    _ = v67
                                }
                            }
                        } else {
                            var v68 float64 = (1.5 - v4)
                            _ = v68
                        }
                    }
                    if false {
                        for fi69 := 0; fi69 < 4; fi69++ {
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            var v70 uint64 = 4294967295
                            v22 = append(v22, fi69)
                            _ = fi69
                            _ = v70
                        }
                        v71 := GenS8{F0: 1, F1: (v20 << ((v20) & 63)), F2: 128}
                        _ = v71
                    }
                }
                for fi72 := 0; fi72 < 8; fi72++ {
                    var v73 int16 = (^(v2))
                    _ = fi72
                    _ = v73
                }
                _ = len(v22)
                _ = fi53
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
