package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int8; F1 uint8; F2 uint8 }

func main() {
    var v1 int = 42
    var v2 float32 = -2.25
    var v3 uint32 = 1
    var v4 uint32 = (v3 & v3)
    v5 := make(map[uint32]bool)
    var v6 int8 = -128
    v7 := GenS1{F0: false, F1: 30, F2: true}
    var v8 int32 = 0
    v9 := v5[968]
    var v10 bool = (v9 && v9)
    if v9 {
        for fi11 := 0; fi11 < 17; fi11++ {
            if (157 <= 1) {
                var v12 float32 = v2
                switch v1 {
                case -15:
                    v9 = v9
                case 4294967295:
                    v13 := GenS1{F0: v9, F1: 0, F2: (false || v9)}
                    _ = v13
                }
                _ = v12
            } else {
                v14 := GenS2{F0: fi11}
                for fi15 := 0; fi15 < 6; fi15++ {
                    v16 := GenS3{F0: 42, F1: 42, F2: (^(v3))}
                    v9 = (-1 == -9223372036854775808)
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            _ = fi11
        }
        var v17 int = (v1 & (v1 << ((163) & 63)))
        _ = v17
    }
    v3 = v3
    var v18 uint8 = 0
    for fi19 := 0; fi19 < 1; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            fmt.Println(fi19, v4)
            var v21 int32 = v8
            fmt.Println(v21)
            v22 := GenS2{F0: fi19}
            switch fi19 {
            case -806:
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v21 = int32(v6)
                case 127:
                    if ((v4 >> ((4294967295) & 63)) != v3) {
                        v23 := []int{32767}
                        fmt.Println(v10, v4)
                        if v10 {
                            v24 := GenS2{F0: (fi20 / ((151) | 1))}
                            _ = v24
                        }
                        var v25 float32 = (3.14159 - v2)
                        _ = v23
                        _ = v25
                    } else {
                        v2 -= -2.25
                        v26 := GenS1{F0: (true && (v9 && ((v1 % ((fi19) | 1)) < (fi19 - (v1 | fi19))))), F1: (v18 * (v18 - (v18 / ((v18) | 1)))), F2: false}
                        v4 -= v4
                        _ = v26
                    }
                case 173:
                    v4 -= v3
                default:
                    switch "\n\t" {
                    case "a":
                        v9 = (-128 < v6)
                    case "hello":
                        fmt.Println(v21, fi19, v3)
                    default:
                        var v27 uint8 = v18
                        _ = v27
                    }
                }
            case 2147483647:
                v5[(v4 - (v3 + v3))] = false
            case 940:
                v6 = v6
                v18 += v18
            default:
                fmt.Println(v8, v3)
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    v10 = v9
    var v28 int8 = (v6 | v6)
    fmt.Println(v6)
    v29 := GenS4{F0: 32767, F1: true}
    switch (v1 % ((v1) | 1)) {
    case -1:
        v30 := GenS5{F0: v28, F1: v18, F2: (^(v18))}
        _ = v30
    case 290:
        if v10 {
            for fi31 := 0; fi31 < 7; fi31++ {
                var v32 uint32 = (v3 % ((v3) | 1))
                if (18446744073709551615 == 1) {
                    v33 := []int{(^(v1))}
                    _ = v33
                } else {
                    v34 := v5[v3]
                    _ = v34
                }
                _ = fi31
                _ = v32
            }
        } else {
            fmt.Println(v9, v6)
            v35 := GenS1{F0: ((42 == -128) && v10), F1: v18, F2: true}
            v36 := GenS1{F0: (((v1 | (v1 << ((18446744073709551615) & 63))) < (v1 | v1)) && (true && (1 == 42))), F1: v18, F2: false}
            _ = v35
            _ = v36
        }
    }
    v9 = (375 <= -1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
