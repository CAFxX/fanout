package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
type GenS2 struct { F0 string; F1 int16; F2 int16 }
type GenS3 struct { F0 int64; F1 bool; F2 string }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int8; F2 uint8 }

func main() {
    var v1 int = -733
    var v2 string = "foo bar"
    var v3 bool = (1 >= 0)
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 int32 = 0
        v6 := GenS1{F0: ((v1 + (v1 ^ (v1 ^ (v1 % ((32767) | 1))))) == v1), F1: v2, F2: -2147483648}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: (60 > 127), F1: (v2 + v2), F2: -272}
    var v8 float64 = -2.25
    v9 := GenS1{F0: (!((-(v1)) == v1)), F1: v2, F2: 837}
    var v10 int = 42
    v11 := []int{9223372036854775807, v1, 32768, v10}
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            fmt.Println(v1, v8)
            _ = fi13
        }
        v1 -= (v10 >> ((0) & 63))
        _ = fi12
    }
    var v14 int32 = -1
    v8 = (v8 + (v8 * v8))
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 int = 42
            var v18 uint32 = 2147483647
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 uint64 = 42
            switch v1 {
            case 32768:
                if false {
                    if v3 {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        v1 -= -897
                        v10 -= (v10 << ((18446744073709551615) & 63))
                    }
                    switch v10 {
                    case 9223372036854775807:
                        _ = len(v11)
                    default:
                        var v21 int = fi15
                        _ = v21
                    }
                }
            case -560:
                fmt.Println(fi19, fi19)
            case -9223372036854775808:
                if (0.1 != 1e10) {
                    v8 = ((float64((0.0 / ((float32(v8) + 1e10) / 1e10))) * v8) / (((3.14159 - v8) / (v8 * float64(0.0))) / 1.5))
                    switch "hello" {
                    case "foo bar":
                        v22 := GenS2{F0: v2, F1: 32767, F2: 872}
                        _ = v22
                    default:
                        v23 := GenS3{F0: 9223372036854775807, F1: true, F2: ""}
                        _ = v23
                    }
                    v24 := GenS1{F0: (v8 == v8), F1: "", F2: v14}
                    _ = v24
                }
            default:
                if v3 {
                    var v25 int16 = 42
                    fmt.Println(v20, fi19)
                    v8 += v8
                    v26 := GenS2{F0: v2, F1: (^(v25)), F2: (v25 + (v25 << ((18446744073709551615) & 63)))}
                    v11 = append(v11, v10)
                    _ = v25
                    _ = v26
                } else {
                    var v27 string = ""
                    v28 := GenS2{F0: v2, F1: 42, F2: 32767}
                    fmt.Println(v10)
                    fmt.Println(v3, v1)
                    if (true || v3) {
                        var v29 uint8 = 37
                        v30 := GenS4{F0: v29, F1: (v20 - 487)}
                        _ = v29
                        _ = v30
                    } else {
                        var v31 float64 = v8
                        _ = v31
                    }
                    _ = v27
                    _ = v28
                }
            }
            _ = fi19
            _ = v20
        }
        _ = fi15
    }
    switch v1 {
    case 42:
        v1 -= v1
    case 127:
        if (!((537 <= 65535) || v3)) {
            var v32 int32 = (v14 << ((255) & 63))
            _ = v32
        }
    case 2147483647:
        var v33 int = v10
        _ = v33
    default:
        var v34 uint32 = 1
        _ = v34
    }
    var v35 string = (("\n\t" + (v2 + v2)) + ("" + "a"))
    v2 = v35
    if v3 {
        for fi36 := 0; fi36 < 1; fi36++ {
            v14 = (v14 / (((-(v14))) | 1))
            var v37 uint32 = 2147483648
            v38 := GenS5{F0: 42, F1: 1, F2: 128}
            _ = fi36
            _ = v37
            _ = v38
        }
        fmt.Println(v8, v8)
    } else {
        v3 = ((0.1 / float32(((v8 - (v8 * (((1.5 + v8) + 1.5) + v8))) + (v8 + v8)))) != 3.14159)
        for fi39 := 0; fi39 < 15; fi39++ {
            v14 = (v14 << ((595) & 63))
            _ = fi39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v7
    _ = v8
    _ = v9
}
