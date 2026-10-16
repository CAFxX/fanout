package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint16; F1 int }

func main() {
    var v1 int32 = 724
    v2 := make(map[uint32]int)
    v3 := []int{}
    var v4 bool = (117 == 127)
    var v5 uint = 776
    var v6 float32 = ((-2.25 - 1.5) * 1.5)
    if v4 {
        for fi7 := 0; fi7 < 1; fi7++ {
            v5 -= (v5 >> ((710) & 63))
            _ = fi7
        }
        if true {
            v8 := GenS1{F0: -195, F1: -2147483648}
            if len(v3) > 0 {
                v3[0] = 97
            }
            var v9 int8 = 0
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: -9223372036854775808, F1: v1}
        _ = v10
    } else {
        _ = len(v3)
        v6 = ((float32(3.14159) * (3.14159 - v6)) * v6)
        var v11 bool = ((^(v5)) <= v5)
        _ = v11
    }
    var v12 uint64 = 0
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS1{F0: -9223372036854775808, F1: 0}
        _ = fi13
        _ = v14
    }
    var v15 int8 = -70
    v1 -= -2147483648
    v16 := GenS1{F0: -9223372036854775808, F1: -2147483648}
    for fi17 := 0; fi17 < 9; fi17++ {
        switch 0 {
        case 127:
            switch 127 {
            case 32768:
                for fi18 := 0; fi18 < 4; fi18++ {
                    v12 -= v12
                    _ = fi18
                }
            case 32767:
                _ = len(v3)
            case -713:
                for fi19 := 0; fi19 < 16; fi19++ {
                    if v4 {
                        switch 0 {
                        case -800:
                            v20 := GenS2{F0: 42, F1: fi17}
                            if ((v4 && (-136 > 0)) && (0 >= 32767)) {
                                v21 := v2[2147483647]
                                if true {
                                    v12 += (v12 >> ((633) & 63))
                                    v2[4294967295] = (fi19 * 128)
                                } else {
                                    var v22 uint64 = v12
                                    _ = v22
                                }
                                _ = v21
                            }
                            _ = v20
                        case 2147483647:
                            v4 = v4
                        case -924:
                            switch "foo bar" {
                            case "\n\t":
                                var v23 int64 = 9223372036854775807
                                _ = v23
                            }
                        default:
                            switch ("foo bar" + "") {
                            case "":
                                switch "\n\t" {
                                case "":
                                    fmt.Println(v1, v4)
                                default:
                                    fmt.Println(fi17, fi19)
                                }
                            case "a":
                                var v24 int64 = 9223372036854775807
                                _ = v24
                            default:
                                v1 -= v1
                            }
                        }
                        switch fi19 {
                        case 42:
                            if (221 != 1000) {
                                v25 := GenS1{F0: 44, F1: -1}
                                _ = v25
                            }
                        default:
                            _ = len(v2)
                            v26 := GenS2{F0: 2, F1: (fi19 * fi19)}
                            _ = v26
                        }
                    } else {
                        if (true && v4) {
                            if v4 {
                                var v27 float32 = v6
                                var v28 uint64 = (v12 - v12)
                                v28 = (v12 * 570)
                                switch "foo bar" {
                                case "foo bar":
                                    if (474 != 464) {
                                        var v29 uint64 = (v12 ^ v12)
                                        _ = v29
                                    }
                                    fmt.Println(fi19, v12, v27)
                                default:
                                    v5 += 65535
                                }
                                v30 := GenS2{F0: 1, F1: 883}
                                _ = v27
                                _ = v28
                                _ = v30
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(fi17, v5, v6)
                                    var v31 float32 = float32((0.0 / (3.14159 / float64((float32(1.5) * v6)))))
                                    _ = v31
                                default:
                                    var v32 uint = v5
                                    _ = v32
                                }
                            }
                            var v33 bool = true
                            switch ("foo bar" + ("" + "")) {
                            case "a":
                                if (1 <= 36) {
                                    switch (fi19 ^ -2147483648) {
                                    case 694:
                                        v4 = ((!v33) || false)
                                    }
                                    v6 -= v6
                                }
                            case "":
                                fmt.Println(v4, fi17, v33)
                            default:
                                var v34 uint = 2
                                _ = v34
                            }
                            _ = v33
                        }
                        if ((v6 - (1e10 * -2.25)) != 0.0) {
                            if v4 {
                                v12 = (v12 * v12)
                            } else {
                                v12 -= v12
                                v35 := GenS2{F0: 32767, F1: -535}
                                v36 := GenS1{F0: 42, F1: (v1 / (((v1 << ((v5) & 63))) | 1))}
                                _ = v35
                                _ = v36
                            }
                        }
                        v6 += -2.25
                    }
                    if v4 {
                        var v37 bool = (1 < 0)
                        fmt.Println(v37, v4)
                        _ = v37
                    } else {
                        var v38 int = fi19
                        _ = len(v2)
                        fmt.Println(v12, fi19, v4)
                        _ = v38
                    }
                    _ = fi19
                }
            }
        }
        var v39 float64 = 1.5
        fmt.Println(v1, v12)
        _ = fi17
        _ = v39
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v40 := GenS1{F0: 0, F1: v1}
    var v41 string = ("foo bar" + "\n\t")
    v12 = v12
    v5 = v5
    v42 := GenS2{F0: 1, F1: -2147483648}
    var v43 uint8 = 2
    v44 := GenS1{F0: 69, F1: v1}
    fmt.Println(v6, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
