package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int64 = 0
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 string = ("a" + ("\n\t" + "hello"))
    var v5 uint32 = 2
    v6 := v3["hello"]
    switch 4294967295 {
    case 0:
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := GenS1{F0: (("foo bar" + "") + v4), F1: 247, F2: 1}
                v10 := GenS1{F0: "hello", F1: 238, F2: fi7}
                switch v4 {
                case "foo bar":
                    v11 := GenS2{F0: 0}
                    fmt.Println(fi8, v6)
                    switch fi8 {
                    case -2147483648:
                        v12 := v2[(fi8 & 0)]
                        _ = v12
                    case -1:
                        if v6 {
                            if false {
                                var v13 uint8 = 2
                                v1 = -9223372036854775808
                                var v14 uint64 = 134
                                if true {
                                    v5 -= (v5 | (v5 - v5))
                                    switch -9223372036854775808 {
                                    case -271:
                                        v4 = ("\n\t" + "a")
                                    case -1:
                                        if (2 <= 791) {
                                            fmt.Println(v14, v14)
                                        }
                                    default:
                                        v15 := GenS3{F0: v14}
                                        _ = v15
                                    }
                                    switch -685 {
                                    case 1:
                                        v6 = (v14 >= (v14 ^ (^(v14))))
                                    case 305:
                                        v16 := GenS3{F0: v14}
                                        _ = v16
                                    case 128:
                                        v17 := []int{(-(fi7)), 127}
                                        v6 = v6
                                        var v18 uint = 42
                                        if v6 {
                                            var v19 float64 = 1.5
                                            v20 := GenS4{F0: (v13 << ((602) & 63)), F1: v1}
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v3[("\n\t" + (v4 + (v4 + (((v4 + v4) + v4) + v4))))] = (4294967295 <= (v5 & 0))
                                            _ = len(v17)
                                            fmt.Println(fi8, fi8)
                                        }
                                        fmt.Println(v5)
                                        fmt.Println(v13)
                                        var v21 bool = ((v13 | v13) != v13)
                                        var v22 string = v4
                                        _ = v17
                                        _ = v18
                                        _ = v21
                                        _ = v22
                                    }
                                }
                                _ = v13
                                _ = v14
                            }
                        }
                    default:
                        v23 := []int{(fi8 | (^(fi8))), (fi7 + (fi7 & (fi7 % (((^(fi7))) | 1)))), 32767, 396}
                        _ = v23
                    }
                    _ = v11
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    }
    if ((-1 == 88) || (1 <= 18446744073709551615)) {
        v24 := v2[814]
        _ = v24
    }
    fmt.Println(v1, v1, v6)
    v1 += -853
    var v25 uint = 4294967295
    v1 -= (v1 << ((uint(v5)) & 63))
    v1 += (v1 / ((v1) | 1))
    v25 = 1
    if (!(v25 != 2)) {
        v1 -= v1
        v6 = v6
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        v5 += 163
        fmt.Println(v5, v1)
        for fi27 := 0; fi27 < 18; fi27++ {
            v28 := []int{}
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v25 = (v25 + 1)
        _ = fi29
    }
    v25 -= (^(v25))
    switch (v4 + (v4 + v4)) {
    case "foo bar":
        v30 := GenS3{F0: 983}
        _ = v30
    default:
        v31 := GenS3{F0: 42}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
