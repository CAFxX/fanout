package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 0.1
    v3 := make(map[int]string)
    var v4 int = 39
    var v5 int = v4
    for fi6 := 0; fi6 < 5; fi6++ {
        if (((v2 - v2) * 1.5) >= (1e10 / (0.1 + float32((float64(v2) * 1.5))))) {
            var v7 bool = false
            var v8 int = 127
            fmt.Println(v4, fi6, v8)
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 0, F1: 1}
            _ = v9
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            v5 = fi6
            _ = fi10
        }
        switch "foo bar" {
        case "foo bar":
            v11 := v3[(^(v4))]
            _ = v11
        }
        _ = fi6
    }
    v12 := []int{(v4 | (v5 << ((234) & 63))), -1, 0}
    v13 := GenS1{F0: 0, F1: -940}
    _ = len(v3)
    v14 := GenS2{F0: 18446744073709551615, F1: 184}
    var v15 int = (v4 % (((-(v4))) | 1))
    if (866 == -796) {
        switch -2147483648 {
        case 42:
            _ = len(v1)
        case -37:
            v4 += v4
        }
        v16 := v3[v4]
        _ = v16
    }
    if len(v12) > 0 {
        v12[0] = (^(v4))
    }
    v4 = (v5 * v5)
    if (!true) {
        var v17 uint16 = 2
        v2 += 3.14159
        var v18 int16 = 0
        _ = v17
        _ = v18
    } else {
        if true {
            v19 := v3[v5]
            v4 -= v15
            _ = v19
        } else {
            if len(v12) > 0 {
                v12[0] = v15
            }
        }
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v2 = float32(0.1)
        var v21 uint64 = 42
        _ = fi20
        _ = v21
    }
    var v22 int = (^(v5))
    if (-2147483648 >= 0) {
        var v23 int64 = 42
        _ = v23
    } else {
        if ("foo bar" >= ("" + "hello")) {
            if ((((42 != 42) || ((true && false) && (true && false))) || true) || false) {
                v24 := GenS3{F0: "\n\t", F1: 875}
                switch 128 {
                case 127:
                    if true {
                        fmt.Println(v15)
                        fmt.Println(v15, v5)
                    } else {
                        switch (v5 ^ 1) {
                        case 552:
                            switch 127 {
                            case -898:
                                var v25 uint64 = 118
                                _ = v25
                            case 32768:
                                switch "foo bar" {
                                case "a":
                                    v26 := v1[""]
                                    _ = v26
                                case "foo bar":
                                    var v27 int64 = 9223372036854775807
                                    _ = v27
                                default:
                                    for fi28 := 0; fi28 < 1; fi28++ {
                                        var v29 int16 = 0
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                            case 0:
                                v30 := GenS2{F0: 701, F1: 181}
                                _ = v30
                            }
                        case 727:
                            for fi31 := 0; fi31 < 1; fi31++ {
                                if len(v12) > 0 {
                                    v12[0] = (v22 & (v15 << ((367) & 63)))
                                }
                                _ = fi31
                            }
                        }
                        for fi32 := 0; fi32 < 17; fi32++ {
                            v1[("" + "")] = "foo bar"
                            if (-617 == 1) {
                                var v33 float32 = v2
                                v12 = append(v12, -883)
                                for fi34 := 0; fi34 < 12; fi34++ {
                                    v35 := GenS4{F0: -388}
                                    _ = fi34
                                    _ = v35
                                }
                                _ = v33
                            }
                            v36 := GenS3{F0: ("hello" + ""), F1: 305}
                            _ = fi32
                            _ = v36
                        }
                    }
                }
                v4 = (v4 | (v5 | (v4 & v4)))
                _ = v24
            } else {
                var v37 string = "\n\t"
                _ = v37
            }
            fmt.Println(v5, v5)
        } else {
            _ = len(v12)
            v4 = (v22 - v15)
            var v38 int64 = -968
            _ = v38
        }
        v39 := GenS2{F0: 18446744073709551615, F1: 186}
        _ = len(v12)
        _ = v39
    }
    v40 := GenS3{F0: "hello", F1: 922}
    switch (v5 ^ v15) {
    case -1:
        if (!true) {
            var v41 float64 = -2.25
            for fi42 := 0; fi42 < 19; fi42++ {
                fmt.Println(v2)
                var v43 bool = false
                _ = fi42
                _ = v43
            }
            var v44 int = v5
            _ = v41
            _ = v44
        }
    default:
        v4 -= v15
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        for fi46 := 0; fi46 < 2; fi46++ {
            v47 := v3[(fi45 - (v4 << ((314) & 63)))]
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
