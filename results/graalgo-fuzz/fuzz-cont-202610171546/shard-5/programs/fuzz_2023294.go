package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 string }
type GenS2 struct { F0 uint8; F1 int64; F2 int64 }
type GenS3 struct { F0 int; F1 uint; F2 uint16 }

func main() {
    var v1 uint16 = 526
    v2 := make(map[int]string)
    var v3 int = 2147483647
    var v4 int64 = -152
    v5 := []int{(v3 * (v3 % (((-(v3))) | 1))), v3, v3, v3}
    var v6 int64 = v4
    v7 := v2[(v3 | (v3 << ((371) & 63)))]
    switch v7 {
    case "":
        if ((false || (!true)) || ((((0.1 * 0.1) * 3.14159) * -2.25) >= ((1e10 - -2.25) * 0.0))) {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 int16 = 1000
                var v10 int32 = -450
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 float64 = float64(((0.0 * 1.5) * -2.25))
                    v13 := GenS1{F0: 1, F1: 2, F2: ((("foo bar" + v7) + v7) + v7)}
                    fmt.Println(v12)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            for fi14 := 0; fi14 < 18; fi14++ {
                if (!false) {
                    if false {
                        for fi15 := 0; fi15 < 17; fi15++ {
                            v6 -= v4
                            v3 += v3
                            _ = fi15
                        }
                        v16 := v2[v3]
                        for fi17 := 0; fi17 < 15; fi17++ {
                            v3 += v3
                            v1 = v1
                            _ = fi17
                        }
                        _ = v16
                    } else {
                        fmt.Println(v7, fi14, v7)
                        var v18 int16 = -32768
                        _ = len(v5)
                        _ = v18
                    }
                    var v19 uint8 = 44
                    _ = v19
                } else {
                    v6 = (^(v6))
                    if (!true) {
                        if (42 < 18446744073709551615) {
                            v1 = (v1 * (v1 + (v1 ^ 0)))
                            v5 = append(v5, fi14)
                            switch v7 {
                            case "hello":
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    if (!(!false)) {
                                        var v21 string = v7
                                        v22 := GenS1{F0: fi20, F1: 2147483647, F2: ("\n\t" + "")}
                                        _ = v21
                                        _ = v22
                                    }
                                    var v23 bool = (("a" + "a") < v7)
                                    if v23 {
                                        v24 := v2[fi14]
                                        fmt.Println(fi14, v24, v6)
                                        _ = v24
                                    } else {
                                        v3 = -1
                                        var v25 bool = true
                                        _ = v25
                                    }
                                    var v26 int8 = 1
                                    _ = fi20
                                    _ = v23
                                    _ = v26
                                }
                            case "a":
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    fmt.Println(v6, v6)
                                    fmt.Println(v6, v4)
                                    _ = fi27
                                }
                            case "\n\t":
                                switch v7 {
                                case "hello":
                                    fmt.Println(v3)
                                default:
                                    v6 = -9223372036854775808
                                }
                            default:
                                var v28 bool = false
                                _ = v28
                            }
                        }
                        var v29 bool = (true || (!true))
                        _ = v29
                    }
                }
                for fi30 := 0; fi30 < 11; fi30++ {
                    switch (fi30 & fi14) {
                    case 32768:
                        switch v7 {
                        case "hello":
                            fmt.Println(fi30, v7, v6)
                        case "\n\t":
                            _ = len(v5)
                        }
                    case 0:
                        var v31 uint = 2
                        _ = v31
                    case 942:
                        v7 = v7
                        _ = len(v2)
                        v32 := GenS1{F0: 127, F1: 0, F2: ((v7 + v7) + v7)}
                        _ = v32
                    default:
                        v33 := GenS2{F0: 165, F1: (v4 >> ((18446744073709551615) & 63)), F2: v6}
                        _ = v33
                    }
                    var v34 float32 = float32(3.14159)
                    fmt.Println(fi14, v34, v1)
                    _ = fi30
                    _ = v34
                }
                _ = fi14
            }
        }
    }
    v35 := GenS2{F0: 249, F1: (v4 - (^(v4))), F2: 0}
    v1 = v1
    v36 := GenS3{F0: (v3 | (v3 - (v3 - -9223372036854775808))), F1: 4294967295, F2: (v1 * v1)}
    var v37 uint32 = 42
    for fi38 := 0; fi38 < 3; fi38++ {
        for fi39 := 0; fi39 < 19; fi39++ {
            v37 = 42
            var v40 uint16 = v1
            _ = fi39
            _ = v40
        }
        v5 = append(v5, (-(fi38)))
        _ = fi38
    }
    v37 = v37
    v37 += (v37 / ((v37) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
