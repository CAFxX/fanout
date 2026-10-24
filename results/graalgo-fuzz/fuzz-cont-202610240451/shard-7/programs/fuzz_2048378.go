package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int8 }

func main() {
    var v1 bool = false
    var v2 float64 = float64(((0.1 - (0.1 - (((3.14159 + 3.14159) / 1.5) / -2.25))) * -2.25))
    var v3 int = 584
    var v4 float32 = (float32((1.5 + float64((((1e10 - 0.1) + 1.5) / 1.0)))) / (-2.25 / 1.0))
    var v5 uint = 42
    if v1 {
        var v6 uint64 = 0
        v7 := GenS1{F0: v1, F1: 255}
        if (253 >= 42) {
            switch 9223372036854775807 {
            case 36:
                v8 := GenS2{F0: 127, F1: 198, F2: 42}
                _ = v8
            case 699:
                var v9 int16 = 967
                _ = v9
            case 0:
                switch "a" {
                case "hello":
                    v10 := GenS1{F0: false, F1: 128}
                    _ = v10
                case "foo bar":
                    var v11 int64 = int64(v5)
                    _ = v11
                case "a":
                    v12 := []int{(v3 - v3)}
                    _ = v12
                default:
                    v6 -= (v6 * (v6 - (v6 - (v6 << ((v5) & 63)))))
                }
            }
            v13 := []int{}
            _ = v13
        } else {
            var v14 uint8 = 128
            switch "\n\t" {
            case "\n\t":
                v15 := []int{v3}
                _ = v15
            case "a":
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := GenS1{F0: ((v6 + (v6 % ((18446744073709551615) | 1))) > 18446744073709551615), F1: (v14 % ((152) | 1))}
                    switch (^(v3)) {
                    case 2147483647:
                        if v1 {
                            for fi18 := 0; fi18 < 4; fi18++ {
                                fmt.Println(v4)
                                _ = fi18
                            }
                            v2 += ((v2 * float64(-2.25)) - (v2 + v2))
                            v4 = ((3.14159 / v4) * 0.1)
                        } else {
                            v19 := GenS2{F0: 172, F1: 2, F2: -46}
                            _ = v19
                        }
                    default:
                        v20 := []int{-507, 128, fi16}
                        _ = v20
                    }
                    _ = fi16
                    _ = v17
                }
            case "foo bar":
                v3 = 42
            }
            _ = v14
        }
        _ = v6
        _ = v7
    }
    v21 := []int{(v3 + (v3 << ((v5) & 63))), -9223372036854775808, v3, 557}
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 uint8 = 0
        v23 = v23
        _ = fi22
        _ = v23
    }
    if (!(-1 == 42)) {
        switch "foo bar" {
        case "":
            v24 := GenS1{F0: (-2147483648 >= -1), F1: 0}
            _ = v24
        }
    } else {
        switch v3 {
        case 2147483647:
            v3 -= (-(v3))
        case 4294967295:
            if v1 {
                for fi25 := 0; fi25 < 6; fi25++ {
                    switch "" {
                    case "":
                        v5 = (v5 / ((v5) | 1))
                    case "hello":
                        var v26 int = v3
                        _ = v26
                    default:
                        v2 += 3.14159
                    }
                    var v27 float32 = float32((v2 + ((v2 + (v2 * (1e10 * v2))) - 1e10)))
                    _ = fi25
                    _ = v27
                }
                for fi28 := 0; fi28 < 1; fi28++ {
                    _ = len(v21)
                    for fi29 := 0; fi29 < 17; fi29++ {
                        if (!(-150 == int16(v5))) {
                            fmt.Println(fi29)
                            var v30 uint64 = 483
                            _ = v30
                        }
                        if false {
                            var v31 uint16 = 32767
                            fmt.Println(v5, fi28, v5)
                            _ = v31
                        } else {
                            v1 = (!(!(228 < -1)))
                            var v32 float64 = v2
                            _ = v32
                        }
                        _ = fi29
                    }
                    v2 += v2
                    _ = fi28
                }
                v3 += v3
            } else {
                fmt.Println(v3)
                v21 = append(v21, (v3 % ((v3) | 1)))
            }
        default:
            v1 = (!v1)
        }
        fmt.Println(v4, v1, v2)
        var v33 int64 = 0
        _ = v33
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        fmt.Println(v4, v3)
        _ = fi34
    }
    v21 = append(v21, 373)
    var v35 uint32 = 84
    fmt.Println(v35, v35)
    v35 = (v35 << ((v5) & 63))
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v5 = 1
    v35 = 925
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
