package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint8; F2 int16 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint; F2 uint }

func main() {
    v1 := []int{4294967295, 9223372036854775807}
    var v2 uint8 = 179
    v3 := []uint32{}
    var v4 int32 = -638
    if true {
        var v5 int16 = -32768
        v4 = int32(v5)
        v5 += (^(v5))
        _ = v5
    }
    if (32767 >= 2) {
        var v6 uint = 42
        _ = v6
    }
    var v7 uint8 = v2
    fmt.Println(v2, v7, v4)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int64 = 9223372036854775807
        _ = fi8
        _ = v9
    }
    if (!false) {
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                switch 1 {
                case -9223372036854775808:
                    switch "\n\t" {
                    case "":
                        var v12 uint64 = 462
                        _ = v12
                    default:
                        v7 = 0
                    }
                case 445:
                    v13 := GenS1{F0: (v4 + (v4 - (v4 % ((v4) | 1)))), F1: -9223372036854775808}
                    _ = v13
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v14 := GenS2{F0: fi11}
                v4 = (v4 + v4)
                _ = fi11
                _ = v14
            }
            _ = fi10
        }
    }
    fmt.Println(v2)
    var v15 uint32 = 147
    switch 128 {
    case -2147483648:
        var v16 uint32 = (^(v15))
        _ = v16
    case -841:
        if true {
            if len(v1) > 0 {
                v1[0] = -154
            }
            v7 = v7
            var v17 float64 = (1.5 + -2.25)
            _ = v17
        } else {
            var v18 bool = false
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
            v15 = (v15 / ((v15) | 1))
            _ = v18
        }
    case 32767:
        _ = len(v3)
    }
    v19 := GenS3{F0: v4, F1: (v2 + v7), F2: 204}
    v20 := GenS4{F0: ((v15 / ((v15) | 1)) < (v15 + (v15 - 4294967295))), F1: v7}
    if true {
        v3 = append(v3, 762)
        if false {
            if (!false) {
                v21 := GenS4{F0: (18446744073709551615 < 255), F1: (v2 + v2)}
                _ = v21
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                v4 += 42
                var v23 int = (fi22 & (-(fi22)))
                _ = fi22
                _ = v23
            }
        }
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS3{F0: v4, F1: v7, F2: -1}
        v15 -= v15
        var v26 int8 = -128
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 20; fi28++ {
            fmt.Println(v15)
            v29 := GenS5{F0: v7}
            var v30 int8 = 1
            _ = fi28
            _ = v29
            _ = v30
        }
        if (int32(v15) <= 781) {
            if len(v3) > 0 {
                v3[0] = 0
            }
            fmt.Println(v2)
            switch fi27 {
            case 32767:
                v4 -= (v4 >> ((65535) & 63))
            }
        }
        var v31 int64 = -9223372036854775808
        _ = fi27
        _ = v31
    }
    var v32 uint64 = 2
    v33 := GenS6{F0: 42, F1: 2, F2: 526}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v7
}
