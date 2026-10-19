package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64; F1 int8; F2 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    var v3 bool = ((true && true) && (true && (!(!false))))
    var v4 string = "\n\t"
    v5 := v1[-986]
    if v3 {
        fmt.Println(v3, v3)
        for fi6 := 0; fi6 < 9; fi6++ {
            v4 = v5
            fmt.Println(v3)
            _ = fi6
        }
    }
    fmt.Println(v4, v5, v5)
    if v3 {
        switch "foo bar" {
        case "hello":
            if (-1 > -32768) {
                switch "" {
                case "":
                    if true {
                        if ((42 < 100) && (!true)) {
                            for fi7 := 0; fi7 < 14; fi7++ {
                                fmt.Println(v2)
                                v3 = true
                                v8 := GenS1{F0: 18446744073709551615, F1: 494}
                                _ = fi7
                                _ = v8
                            }
                            var v9 int64 = -416
                            _ = v9
                        } else {
                            v10 := []int{}
                            _ = v10
                        }
                        if (683 > 78) {
                            for fi11 := 0; fi11 < 16; fi11++ {
                                var v12 uint16 = 855
                                if (v12 <= v12) {
                                    var v13 int64 = 235
                                    _ = v13
                                }
                                v2 = (v2 + 42)
                                _ = fi11
                                _ = v12
                            }
                        } else {
                            if v3 {
                                var v14 int16 = 869
                                _ = v14
                            }
                            if v3 {
                                if (0 < 0) {
                                    if ((-9223372036854775808 < -708) || (int64(v2) != -638)) {
                                        v3 = v3
                                        v1[9223372036854775807] = ""
                                    } else {
                                        v4 = v4
                                    }
                                    var v15 uint = 4294967295
                                    _ = v15
                                }
                            }
                            var v16 uint32 = 16
                            _ = v16
                        }
                        v17 := GenS2{F0: 575, F1: -128, F2: ("foo bar" + "\n\t")}
                        _ = v17
                    }
                case "foo bar":
                    for fi18 := 0; fi18 < 15; fi18++ {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            var v20 uint = 932
                            _ = fi19
                            _ = v20
                        }
                        _ = fi18
                    }
                }
                var v21 int64 = 781
                var v22 uint32 = 4294967295
                _ = v21
                _ = v22
            }
        case "":
            v23 := []int{}
            _ = v23
        case "foo bar":
            v24 := []int{}
            _ = v24
        }
    } else {
        if (578 > 180) {
            if v3 {
                fmt.Println(v5)
            } else {
                for fi25 := 0; fi25 < 18; fi25++ {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        switch "\n\t" {
                        case "\n\t":
                            var v27 uint = 2
                            _ = v27
                        case "":
                            fmt.Println(v3, v3, v5)
                        case "a":
                            v2 = v2
                        default:
                            var v28 float32 = (0.0 - 1e10)
                            _ = v28
                        }
                        _ = fi26
                    }
                    _ = fi25
                }
                var v29 int32 = -69
                v30 := []int{}
                _ = v29
                _ = v30
            }
        }
    }
    var v31 float32 = (-2.25 * 0.1)
    v32 := []int{-1, 86}
    fmt.Println(v2)
    var v33 int = -1
    var v34 string = "\n\t"
    var v35 int64 = -488
    v36 := GenS1{F0: 649, F1: -310}
    for fi37 := 0; fi37 < 4; fi37++ {
        v35 = v35
        var v38 float64 = 3.14159
        var v39 int = -127
        _ = fi37
        _ = v38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
