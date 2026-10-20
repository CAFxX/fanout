package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int; F1 int32; F2 int8 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int; F1 uint16; F2 int16 }
func genFn0(p0 int16, p1 int8) bool {
    var v6 float32 = (0.1 / 1.5)
    v7 := []int{-46}
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            p0 += p0
            _ = fi9
        }
        switch (fi8 & -250) {
        case 1:
            var v10 uint32 = 834
            _ = v10
        }
        _ = fi8
    }
    var v11 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 uint8, p1 uint8) string {
    v12 := GenS1{F0: 2, F1: 0, F2: 572}
    v13 := []int{}
    if true {
        v14 := GenS2{F0: "a", F1: -47}
        var v15 bool = false
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return "hello"
}

func main() {
    v1 := []uint32{51, 65535, 1, 964}
    var v2 string = ("" + "")
    v3 := make(map[uint32]int)
    var v4 bool = true
    var v5 int32 = -399
    fmt.Println(v2, v5, v5)
    _ = len(v3)
    v16 := GenS3{F0: 2, F1: 1}
    switch v2 {
    case "":
        v4 = v4
    }
    var v17 int32 = v5
    var v18 int16 = 0
    v4 = (1 == 42)
    v17 -= -1
    v5 = -2147483648
    switch -1 {
    case -1:
        switch "\n\t" {
        case "":
            if (!(42 <= 2)) {
                for fi19 := 0; fi19 < 8; fi19++ {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 int = fi19
                        v2 = (v2 + (v2 + ("a" + v2)))
                        switch -5 {
                        case -196:
                            var v22 int32 = (^(v17))
                            _ = v22
                        case 0:
                            if false {
                                v17 = -422
                                switch (fi20 & fi20) {
                                case 1:
                                    var v23 float64 = 1.5
                                    _ = v23
                                case 32768:
                                    fmt.Println(v4, fi19, fi19)
                                }
                            }
                        case 32768:
                            v24 := GenS4{F0: (v21 << ((23) & 63)), F1: (v17 / (((^(v5))) | 1)), F2: -128}
                            v5 += -268
                            _ = v24
                        default:
                            var v25 int8 = 1
                            _ = v25
                        }
                        _ = fi20
                        _ = v21
                    }
                    _ = fi19
                }
                fmt.Println(v5)
                if false {
                    _ = len(v3)
                    var v26 int16 = (v18 ^ v18)
                    v27 := v3[2147483647]
                    _ = v26
                    _ = v27
                }
            }
        case "a":
            v28 := GenS4{F0: 127, F1: -2147483648, F2: 127}
            _ = v28
        }
    case 2147483647:
        fmt.Println(v17, v17)
    default:
        if (1 == -501) {
            v29 := GenS1{F0: 294, F1: 9223372036854775807, F2: 811}
            _ = v29
        }
    }
    switch v2 {
    case "":
        for fi30 := 0; fi30 < 13; fi30++ {
            v4 = (v4 || v4)
            v18 += (v18 | (v18 & (v18 | -32768)))
            v31 := GenS1{F0: 42, F1: 9223372036854775807, F2: 42}
            _ = fi30
            _ = v31
        }
    default:
        for fi32 := 0; fi32 < 15; fi32++ {
            for fi33 := 0; fi33 < 8; fi33++ {
                if (!v4) {
                    fmt.Println(fi33, v17)
                } else {
                    var v34 int32 = -2147483648
                    _ = v34
                }
                v35 := GenS5{F0: fi33, F1: -9223372036854775808}
                v36 := GenS6{F0: v17, F1: fi32}
                if v4 {
                    var v37 int64 = 0
                    var v38 bool = ("foo bar" == "a")
                    _ = v37
                    _ = v38
                }
                var v39 uint64 = 1
                _ = fi33
                _ = v35
                _ = v36
                _ = v39
            }
            if v4 {
                for fi40 := 0; fi40 < 7; fi40++ {
                    v41 := v3[809]
                    _ = fi40
                    _ = v41
                }
                v1 = append(v1, 65535)
            } else {
                var v42 int32 = (v17 & (v17 & (v17 % ((v17) | 1))))
                var v43 int = 983
                fmt.Println(v18)
                _ = v42
                _ = v43
            }
            _ = fi32
        }
    }
    v44 := GenS7{F0: -540, F1: 107, F2: (v18 ^ 32767)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
