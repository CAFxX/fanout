package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 bool; F2 int16 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int16 = 1
    v2 := make(map[int]uint32)
    var v3 uint8 = 42
    var v4 float64 = 1.5
    var v5 int32 = 183
    v6 := []int{-603}
    if (-505 != 42) {
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(fi7, v1, v3)
            switch "\n\t" {
            case "\n\t":
                v8 := GenS1{F0: 943}
                _ = v8
            case "hello":
                v4 = v4
            default:
                var v9 uint32 = 2147483647
                _ = v9
            }
            _ = fi7
        }
    }
    if (!true) {
        if true {
            if true {
                var v10 int = -9223372036854775808
                if true {
                    if false {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        switch ("hello" + "a") {
                        case "hello":
                            v10 -= v10
                        case "a":
                            for fi11 := 0; fi11 < 7; fi11++ {
                                v5 = v5
                                for fi12 := 0; fi12 < 9; fi12++ {
                                    var v13 uint = 2
                                    v4 = (3.14159 - float64(0.0))
                                    _ = fi12
                                    _ = v13
                                }
                                if true {
                                    if (!true) {
                                        v4 -= (v4 - v4)
                                    }
                                    var v14 float64 = 1.5
                                    _ = v14
                                }
                                _ = fi11
                            }
                        default:
                            v10 -= v10
                        }
                        v2[v10] = 42
                    }
                }
                _ = v10
            }
            fmt.Println(v4, v3)
            var v15 int = 4294967295
            _ = v15
        }
        if false {
            v3 -= 1
            fmt.Println(v1, v3)
            v3 += v3
        }
        fmt.Println(v4, v4)
    }
    v3 -= (^(v3))
    _ = len(v2)
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(v1, v1)
        var v17 int32 = v5
        _ = fi16
        _ = v17
    }
    if ((v1 < 0) || true) {
        var v18 int = 4294967295
        _ = v18
    } else {
        fmt.Println(v5)
        if (v5 == (v5 + (v5 & v5))) {
            v4 -= -2.25
        } else {
            v4 -= float64(1e10)
            if (!true) {
                v1 = (-(v1))
            } else {
                var v19 int64 = -1
                _ = v19
            }
        }
    }
    v20 := GenS2{F0: 2, F1: (v3 >= 255), F2: v1}
    v4 = ((0.1 / 1.0) - ((((((float64(0.0) + (v4 / v4)) / 1e10) + 3.14159) * (((float64(0.0) / 1.0) / 0.1) * (v4 - (0.0 * (1e10 - (v4 - 1.5)))))) + ((v4 / float64((1e10 + 3.14159))) + v4)) * v4))
    if len(v6) > 0 {
        v6[0] = 1
    }
    var v21 string = "\n\t"
    delete(v2, -199)
    v4 = (float64(float32(0.0)) / float64(((0.1 * 1e10) / 1e10)))
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 uint64 = 2
        var v24 uint8 = 156
        _ = fi22
        _ = v23
        _ = v24
    }
    if (false || ((32767 == v1) || false)) {
        v3 -= v3
        fmt.Println(v3)
        if false {
            v6 = append(v6, -9223372036854775808)
            _ = len(v6)
        } else {
            v25 := GenS2{F0: 2, F1: (!(2147483648 != 2147483648)), F2: v1}
            if ((v4 + (1e10 * (0.0 - v4))) == float64((float32((v4 / v4)) * ((1e10 + float32((1.5 - v4))) / 0.1)))) {
                v1 += (v1 & v1)
                for fi26 := 0; fi26 < 5; fi26++ {
                    v21 = v21
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v21 = v21
                        if (true || (v4 >= (v4 / v4))) {
                            v4 = (0.1 + v4)
                        } else {
                            var v28 float64 = v4
                            _ = v28
                        }
                        _ = fi27
                    }
                    v6 = append(v6, fi26)
                    _ = fi26
                }
                switch "" {
                case "\n\t":
                    v3 += (^(v3))
                }
            }
            v29 := v2[-1]
            _ = v25
            _ = v29
        }
    } else {
        if false {
            if false {
                var v30 uint32 = 617
                _ = v30
            } else {
                fmt.Println(v3, v1, v21)
                v31 := GenS3{F0: 65535, F1: v3}
                var v32 string = v21
                _ = v31
                _ = v32
            }
            _ = len(v2)
        } else {
            switch 32767 {
            case 127:
                v33 := GenS2{F0: 0, F1: false, F2: (v1 % ((-883) | 1))}
                _ = v33
            case 2147483647:
                v4 += v4
            default:
                if true {
                    var v34 float32 = (1e10 / 0.1)
                    _ = v34
                }
            }
        }
        v35 := GenS4{F0: 32767}
        var v36 float64 = (1e10 - v4)
        _ = v35
        _ = v36
    }
    v5 = v5
    var v37 int = -9223372036854775808
    var v38 int = (v37 - v37)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
