package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 bool }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint8; F1 int16; F2 int8 }

func main() {
    var v1 bool = (-596 != -1)
    var v2 bool = (1 != 0)
    var v3 bool = false
    v4 := []string{"foo bar", "hello"}
    v5 := make(map[uint32]bool)
    var v6 uint = 2
    v5[1] = (918 < 65535)
    v7 := GenS1{F0: 127, F1: true, F2: (-1 < -461)}
    v8 := GenS2{F0: 130, F1: 2147483647}
    var v9 float32 = (float32(1.5) - (float32(3.14159) * (3.14159 * 0.1)))
    for fi10 := 0; fi10 < 8; fi10++ {
        if (!v3) {
            var v11 int8 = -113
            var v12 string = "hello"
            _ = v11
            _ = v12
        }
        v6 = (^(v6))
        _ = fi10
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v9 = 1e10
        if v1 {
            _ = len(v4)
            var v14 int8 = 41
            _ = v14
        }
        _ = fi13
    }
    if (float32(-2.25) <= 1.5) {
        v4 = append(v4, "a")
        var v15 int64 = -9223372036854775808
        switch 32768 {
        case 127:
            var v16 int64 = (v15 & v15)
            _ = v16
        }
        _ = v15
    } else {
        fmt.Println(v6, v6)
        fmt.Println(v1, v1)
        v5[245] = (1 != 2)
    }
    var v17 uint8 = 127
    if false {
        switch 637 {
        case 9223372036854775807:
            if (100 > 1) {
                v2 = false
                if false {
                    var v18 uint32 = 2147483648
                    _ = v18
                } else {
                    var v19 int64 = 165
                    switch 32767 {
                    case 4294967295:
                        if v1 {
                            fmt.Println(v19)
                        } else {
                            var v20 float32 = ((((float32((1.5 * (1.5 * 0.1))) / 1.0) - float32(1e10)) - v9) * 1.5)
                            v21 := GenS3{F0: 4294967295}
                            switch "a" {
                            case "a":
                                for fi22 := 0; fi22 < 12; fi22++ {
                                    switch fi22 {
                                    case 1:
                                        var v23 float32 = v9
                                        _ = v23
                                    case 676:
                                        var v24 uint = (v6 >> ((4294967295) & 63))
                                        _ = v24
                                    case 127:
                                        if v2 {
                                            if true {
                                                v25 := GenS1{F0: 42, F1: v1, F2: v3}
                                                _ = v25
                                            }
                                            v17 = v17
                                            for fi26 := 0; fi26 < 10; fi26++ {
                                                fmt.Println(fi26, v9)
                                                v5[uint32(v17)] = ((!(0 <= -1)) || false)
                                                _ = fi26
                                            }
                                        } else {
                                            for fi27 := 0; fi27 < 9; fi27++ {
                                                fmt.Println(v19)
                                                _ = fi27
                                            }
                                            switch -166 {
                                            case -627:
                                                for fi28 := 0; fi28 < 10; fi28++ {
                                                    fmt.Println(fi22)
                                                    v2 = v3
                                                    _ = fi28
                                                }
                                            }
                                        }
                                    }
                                    _ = fi22
                                }
                            }
                            _ = v20
                            _ = v21
                        }
                    case 128:
                        v29 := GenS4{F0: 2, F1: v6}
                        _ = v29
                    case -1:
                        v30 := GenS4{F0: 32767, F1: 845}
                        _ = v30
                    }
                    v3 = (v6 >= 238)
                    _ = v19
                }
                fmt.Println(v2)
            } else {
                if (((0 >= 92) || false) && (3 == 100)) {
                    if false {
                        for fi31 := 0; fi31 < 19; fi31++ {
                            v2 = (false && ((fi31 <= (^(fi31))) || ((^(v17)) <= v17)))
                            var v32 float64 = (0.0 * (-2.25 - 1.5))
                            for fi33 := 0; fi33 < 6; fi33++ {
                                v34 := GenS5{F0: (v17 % ((2) | 1)), F1: 1000, F2: -54}
                                v9 = v9
                                _ = fi33
                                _ = v34
                            }
                            _ = fi31
                            _ = v32
                        }
                        v2 = true
                    } else {
                        fmt.Println(v1)
                        v3 = ((v17 >> (((v6 << ((65535) & 63))) & 63)) >= v17)
                    }
                } else {
                    v35 := GenS4{F0: 0, F1: v6}
                    v2 = (v2 && true)
                    _ = v35
                }
                var v36 uint64 = 0
                v37 := v5[uint32(v36)]
                _ = v36
                _ = v37
            }
        case -2147483648:
            for fi38 := 0; fi38 < 12; fi38++ {
                var v39 int = 9223372036854775807
                for fi40 := 0; fi40 < 17; fi40++ {
                    delete(v5, 608)
                    var v41 int16 = -1
                    _ = fi40
                    _ = v41
                }
                fmt.Println(v9)
                _ = fi38
                _ = v39
            }
        }
    }
    v42 := v5[65535]
    v6 -= (v6 * v6)
    v43 := GenS5{F0: v17, F1: 1000, F2: -1}
    switch ("foo bar" + "hello") {
    case "":
        var v44 uint = (v6 / ((v6) | 1))
        _ = v44
    case "\n\t":
        for fi45 := 0; fi45 < 2; fi45++ {
            var v46 float64 = (3.14159 + 0.0)
            _ = len(v5)
            _ = fi45
            _ = v46
        }
    default:
        var v47 int64 = 9223372036854775807
        _ = v47
    }
    v42 = (!(uint64(v17) > 18446744073709551615))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
