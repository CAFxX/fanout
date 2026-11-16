package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }

func main() {
    v1 := []string{"\n\t", "a"}
    var v2 uint32 = 92
    v3 := []int{2147483647, -189, 2147483647}
    var v4 int16 = 1
    v5 := make(map[uint32]bool)
    switch 4294967295 {
    case 424:
        v4 += (v4 % ((-14) | 1))
    default:
        var v6 uint16 = 255
        _ = v6
    }
    if true {
        switch -1 {
        case 32767:
            fmt.Println(v4, v4)
        default:
            for fi7 := 0; fi7 < 5; fi7++ {
                v2 -= (v2 & v2)
                fmt.Println(fi7, v2, v2)
                var v8 int8 = 42
                _ = fi7
                _ = v8
            }
        }
        _ = len(v1)
    } else {
        v9 := GenS1{F0: (-801 < 127), F1: 1, F2: (v4 % (((-(v4))) | 1))}
        v2 = (v2 << ((255) & 63))
        if true {
            v4 += (v4 * 476)
        } else {
            v4 += v4
            if true {
                v4 += v4
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 uint16 = 742
                    _ = fi10
                    _ = v11
                }
                v4 = v4
            }
        }
        _ = v9
    }
    if false {
        var v12 int32 = -2147483648
        _ = v12
    } else {
        var v13 uint32 = v2
        fmt.Println(v4, v2, v13)
        _ = v13
    }
    var v14 int = 127
    switch (v14 % ((4294967295) | 1)) {
    case 849:
        switch "hello" {
        case "":
            v2 += v2
        case "a":
            for fi15 := 0; fi15 < 10; fi15++ {
                if len(v1) > 0 {
                    v1[0] = ("" + "")
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    switch "" {
                    case "\n\t":
                        v2 = (v2 ^ (v2 * v2))
                    case "hello":
                        var v17 string = "\n\t"
                        _ = v17
                    }
                    switch (-(fi15)) {
                    case 66:
                        if false {
                            if (!true) {
                                v18 := GenS1{F0: false, F1: 202, F2: (^(v4))}
                                var v19 float64 = 0.0
                                _ = v18
                                _ = v19
                            } else {
                                var v20 uint8 = 4
                                _ = v20
                            }
                            v21 := GenS1{F0: (!(true || (1.5 != -2.25))), F1: 65535, F2: v4}
                            v22 := v5[v2]
                            switch "hello" {
                            case "foo bar":
                                var v23 string = ""
                                _ = v23
                            case "hello":
                                v24 := GenS1{F0: (true || true), F1: 1, F2: (v4 * (v4 * v4))}
                                _ = v24
                            case "a":
                                fmt.Println(v2, v2, v4)
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    _ = len(v5)
                                }
                            }
                            v14 -= v14
                            v22 = (0.1 > 3.14159)
                            var v25 int32 = -2147483648
                            _ = v21
                            _ = v22
                            _ = v25
                        } else {
                            switch "foo bar" {
                            case "hello":
                                v4 += 1000
                            case "foo bar":
                                if ((true && true) && ((^(fi16)) > (v14 - fi16))) {
                                    fmt.Println(v14, v4)
                                    v4 = v4
                                    var v26 uint8 = 128
                                    _ = v26
                                } else {
                                    fmt.Println(fi15)
                                }
                            }
                            var v27 string = "a"
                            v28 := v5[v2]
                            _ = v27
                            _ = v28
                        }
                    default:
                        var v29 int16 = -32768
                        _ = v29
                    }
                    var v30 int = (v14 / ((fi15) | 1))
                    if ((false || true) && (!((true && ((true || (255 > 42)) || (false && false))) || true))) {
                        fmt.Println(v4, fi15, v2)
                        _ = len(v5)
                    }
                    if len(v3) > 0 {
                        v3[0] = (^(v14))
                    }
                    _ = fi16
                    _ = v30
                }
                if false {
                    var v31 int8 = 127
                    var v32 float64 = (((3.14159 * (float64(-2.25) / 1.0)) / 1.0) - 3.14159)
                    _ = v31
                    _ = v32
                }
                _ = fi15
            }
        case "foo bar":
            var v33 int16 = 42
            _ = v33
        }
    case -737:
        var v34 int64 = -1
        _ = v34
    }
    var v35 uint64 = 851
    v2 = (^(v2))
    for fi36 := 0; fi36 < 20; fi36++ {
        v35 += (^(v35))
        v37 := GenS1{F0: (false && false), F1: 18446744073709551615, F2: 1}
        _ = fi36
        _ = v37
    }
    v38 := GenS1{F0: true, F1: 331, F2: (v4 << ((255) & 63))}
    v2 = 65535
    v3 = append(v3, 128)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v5
}
