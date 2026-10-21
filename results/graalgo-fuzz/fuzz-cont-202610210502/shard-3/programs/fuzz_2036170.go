package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 string }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(fi5)
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 32768, F1: false}
    var v8 int32 = 1
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 int32 = 0
    var v2 int64 = 42
    var v3 uint = 217
    var v4 int64 = (v2 << (((^(v3))) & 63))
    v9 := []int{-2147483648, -2147483648}
    var v10 float64 = 0.1
    v1 = -880
    if true {
        var v11 uint32 = 2
        fmt.Println(v1, v10, v2)
        _ = v11
    }
    if len(v9) > 0 {
        v9[0] = 2147483647
    }
    var v12 int32 = v1
    var v13 uint16 = 32767
    v14 := GenS1{F0: int(v3), F1: ((-2.25 - 3.14159) == v10)}
    if false {
        var v15 bool = false
        switch "" {
        case "\n\t":
            v16 := GenS2{F0: 1, F1: (v12 - (v1 & (-(v1))))}
            _ = v16
        case "hello":
            v17 := GenS3{F0: v3, F1: 427}
            _ = v17
        case "foo bar":
            var v18 int = 42
            _ = v18
        }
        _ = v15
    } else {
        for fi19 := 0; fi19 < 14; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                v12 += (^(v12))
                switch fi20 {
                case 32767:
                    fmt.Println(v13)
                case -2147483648:
                    v10 -= v10
                default:
                    v4 = 1
                }
                v21 := GenS4{F0: (v3 / (((v3 * (v3 ^ v3))) | 1))}
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        fmt.Println(v2, v3, v10)
        for fi22 := 0; fi22 < 8; fi22++ {
            v23 := GenS1{F0: (-(fi22)), F1: (42 < 0)}
            fmt.Println(v3)
            v24 := GenS5{F0: (v13 << ((446) & 63)), F1: "foo bar"}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    if false {
        for fi25 := 0; fi25 < 14; fi25++ {
            _ = len(v9)
            var v26 float64 = 1.5
            for fi27 := 0; fi27 < 7; fi27++ {
                v12 += (-(v1))
                var v28 bool = false
                v13 = 1
                _ = fi27
                _ = v28
            }
            _ = fi25
            _ = v26
        }
        if (!((42 <= uint32(v13)) && (!(false && true)))) {
            fmt.Println(v12)
            var v29 float64 = 0.1
            var v30 bool = true
            _ = v29
            _ = v30
        }
        if (1 > 1) {
            v31 := GenS2{F0: 716, F1: v12}
            v10 -= (0.0 / float64(1e10))
            _ = v31
        } else {
            fmt.Println(v2, v10, v10)
            for fi32 := 0; fi32 < 10; fi32++ {
                fmt.Println(v3)
                var v33 uint16 = (v13 * uint16(v3))
                if (v33 != uint16(v3)) {
                    v4 -= v2
                } else {
                    v33 = (v33 ^ (v33 + 330))
                    v2 -= (v4 << (((v3 >> (((v3 - (v3 << ((1) & 63)))) & 63))) & 63))
                }
                _ = fi32
                _ = v33
            }
        }
    }
    var v34 bool = false
    fmt.Println(v12, v12)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    if len(v9) > 0 {
        v9[0] = 32768
    }
    if v34 {
        if false {
            if ((v4 ^ v4) < (v4 * v4)) {
                for fi35 := 0; fi35 < 12; fi35++ {
                    v34 = true
                    _ = fi35
                }
            } else {
                fmt.Println(v2, v1)
                v36 := GenS3{F0: v3, F1: 762}
                _ = v36
            }
            for fi37 := 0; fi37 < 14; fi37++ {
                v12 = (v1 % (((v1 % ((1) | 1))) | 1))
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = fi37
            }
            for fi38 := 0; fi38 < 12; fi38++ {
                var v39 uint16 = 1
                for fi40 := 0; fi40 < 8; fi40++ {
                    var v41 int16 = 42
                    var v42 int16 = v41
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = fi38
                _ = v39
            }
        }
    }
    var v43 int32 = (v12 + (v1 / (((^(v12))) | 1)))
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v9
}
