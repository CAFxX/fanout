package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint }

func main() {
    var v1 uint = 2
    var v2 bool = ("\n\t" != "\n\t")
    var v3 uint64 = 2
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{127, fi5, (fi5 << ((v1) & 63)), fi5}
        var v7 int16 = 32767
        _ = fi5
        _ = v6
        _ = v7
    }
    if (true || (255 > (v1 + 0))) {
        switch "foo bar" {
        case "\n\t":
            v2 = (2147483647 >= 516)
        case "a":
            fmt.Println(v3, v1)
        case "foo bar":
            if v2 {
                if (2147483647 <= 109) {
                    fmt.Println(v2, v3, v1)
                }
                for fi8 := 0; fi8 < 16; fi8++ {
                    if (223 != 186) {
                        for fi9 := 0; fi9 < 9; fi9++ {
                            fmt.Println(v1, fi8)
                            var v10 uint32 = 2147483647
                            _ = fi9
                            _ = v10
                        }
                    } else {
                        if v2 {
                            v3 -= (v3 & v3)
                        } else {
                            if v2 {
                                delete(v4, 42)
                                var v11 float32 = 1.5
                                _ = v11
                            } else {
                                v1 = 2
                                var v12 uint8 = uint8(v3)
                                _ = v12
                            }
                            if (-1 >= 0) {
                                v1 += 255
                                var v13 int = fi8
                                _ = v13
                            }
                        }
                        v14 := GenS1{F0: true}
                        var v15 int16 = 0
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS1{F0: (702 >= 40)}
                    if true {
                        var v17 int8 = -1
                        _ = v17
                    } else {
                        var v18 int16 = 105
                        _ = v18
                    }
                    _ = fi8
                    _ = v16
                }
            }
        default:
            switch (("hello" + "hello") + ("" + "a")) {
            case "hello":
                var v19 uint16 = 0
                _ = v19
            case "\n\t":
                for fi20 := 0; fi20 < 2; fi20++ {
                    fmt.Println(v3, fi20, v1)
                    v21 := GenS1{F0: v2}
                    v2 = true
                    _ = fi20
                    _ = v21
                }
            default:
                var v22 uint = 368
                _ = v22
            }
        }
    } else {
        v23 := []int{822}
        v3 -= (v3 * (v3 % ((v3) | 1)))
        _ = v23
    }
    v3 = v3
    for fi24 := 0; fi24 < 1; fi24++ {
        v25 := []int{fi24, (fi24 / ((42) | 1)), 777, (-(fi24))}
        _ = fi24
        _ = v25
    }
    if (!(3.14159 < float64(1e10))) {
        if (32767 > 1000) {
            switch 912 {
            case -880:
                v26 := GenS2{F0: 42, F1: v1}
                _ = v26
            default:
                for fi27 := 0; fi27 < 19; fi27++ {
                    var v28 uint16 = 32767
                    _ = fi27
                    _ = v28
                }
            }
        } else {
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 uint16 = 65535
                fmt.Println(v2, fi29, v2)
                _ = fi29
                _ = v30
            }
            v31 := []int{}
            var v32 uint16 = 800
            _ = v31
            _ = v32
        }
    } else {
        if (0 < -9223372036854775808) {
            fmt.Println(v3, v1, v2)
        } else {
            if true {
                var v33 uint8 = 42
                var v34 int64 = 9223372036854775807
                _ = v33
                _ = v34
            } else {
                var v35 uint64 = (v3 << ((v1) & 63))
                fmt.Println(v2)
                _ = v35
            }
        }
        if (843 > 1) {
            v36 := GenS3{F0: 2147483647}
            _ = v36
        } else {
            var v37 int32 = 1
            v1 = (^(v1))
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        if true {
            var v39 string = "foo bar"
            for fi40 := 0; fi40 < 16; fi40++ {
                switch v39 {
                case "a":
                    v39 = v39
                }
                v4[72] = v39
                v1 += (v1 ^ (^(v1)))
                _ = fi40
            }
            _ = v39
        } else {
            v41 := []int{fi38, 2147483647, (fi38 ^ fi38)}
            _ = v41
        }
        v2 = v2
        switch "foo bar" {
        case "foo bar":
            switch fi38 {
            case -2147483648:
                for fi42 := 0; fi42 < 5; fi42++ {
                    delete(v4, 2147483647)
                    _ = fi42
                }
            case 944:
                delete(v4, 134)
            case 46:
                v2 = ((!((v3 / ((0) | 1)) >= 18446744073709551615)) && v2)
            default:
                v3 -= 536
            }
        case "a":
            var v43 string = (("" + "foo bar") + "hello")
            _ = v43
        case "":
            for fi44 := 0; fi44 < 14; fi44++ {
                v45 := GenS4{F0: (^(v3)), F1: -27, F2: (^(v1))}
                _ = fi44
                _ = v45
            }
        default:
            v46 := []int{fi38, fi38, (fi38 / (((fi38 << ((373) & 63))) | 1))}
            _ = v46
        }
        _ = fi38
    }
    v47 := GenS2{F0: -2147483648, F1: 903}
    fmt.Println(v2)
    v48 := GenS4{F0: (^(v3)), F1: -1, F2: v1}
    var v49 uint = 18446744073709551615
    var v50 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
