package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int = 0
    var v2 bool = true
    v3 := []int{}
    v4 := make(map[int]bool)
    v5 := []int{(-(v1)), (v1 << ((65535) & 63)), (v1 * v1), v1}
    v1 += (v1 % ((v1) | 1))
    v1 -= (v1 << ((293) & 63))
    fmt.Println(v1, v2)
    v1 -= (v1 & -320)
    var v6 float64 = -2.25
    if v2 {
        switch 1 {
        case 465:
            v4[v1] = false
        case 32767:
            if true {
                for fi7 := 0; fi7 < 10; fi7++ {
                    var v8 uint64 = 763
                    if (((-2.25 * 0.0) == ((((((((0.0 / 1.5) - 1e10) / 3.14159) / ((1.5 / 1.5) / 1.5)) - ((float32(1.5) + 1e10) + (3.14159 * 0.1))) - 1e10) - -2.25) / 1.5)) && (!(-1 != 0))) {
                        v9 := GenS1{F0: 4294967295, F1: 42, F2: 255}
                        var v10 uint32 = 65535
                        v10 = v10
                        _ = v10
                        _ = v9
                    } else {
                        var v11 float32 = 0.0
                        switch 32767 {
                        case 115:
                            v2 = (148 < 2)
                        }
                        v2 = (62 < 0)
                        _ = v11
                    }
                    _ = fi7
                    _ = v8
                }
            }
        case 2147483647:
            switch "\n\t" {
            case "a":
                v6 = (v6 * 0.1)
            case "hello":
                for fi12 := 0; fi12 < 17; fi12++ {
                    for fi13 := 0; fi13 < 6; fi13++ {
                        fmt.Println(v2, v1)
                        v14 := GenS2{F0: 42, F1: 65535}
                        v2 = (!true)
                        _ = fi13
                        _ = v14
                    }
                    if (-9223372036854775808 >= 9223372036854775807) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v15 float32 = 1e10
                        for fi16 := 0; fi16 < 1; fi16++ {
                            v17 := v4[(-(fi12))]
                            _ = fi16
                            _ = v17
                        }
                        _ = v15
                    }
                    _ = fi12
                }
            }
        default:
            v18 := GenS3{F0: (!v2)}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint8 = 76
        var v21 bool = false
        _ = fi19
        _ = v20
        _ = v21
    }
    if v2 {
        fmt.Println(v2)
        switch "" {
        case "hello":
            if (!(65535 > 4294967295)) {
                var v22 uint32 = 42
                _ = v22
            } else {
                switch ("a" + "hello") {
                case "a":
                    v23 := GenS1{F0: (v1 >> ((138) & 63)), F1: 1, F2: 65535}
                    _ = v23
                case "hello":
                    v24 := GenS4{F0: 0, F1: -9223372036854775808}
                    _ = v24
                case "\n\t":
                    v25 := GenS2{F0: -9223372036854775808, F1: 0}
                    _ = v25
                }
                v26 := GenS1{F0: 4294967295, F1: 65535, F2: 255}
                _ = v26
            }
        default:
            v27 := GenS1{F0: (v1 * v1), F1: 477, F2: 255}
            _ = v27
        }
        var v28 float32 = 3.14159
        _ = v28
    }
    fmt.Println(v1, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
