package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
func genFn0(p0 int32) int64 {
    v6 := GenS1{F0: int(p0), F1: 867, F2: -137}
    p0 = p0
    _ = p0
    _ = v6
    return 923
}

func main() {
    var v1 uint = 695
    var v2 int64 = -698
    var v3 int32 = -1
    var v4 uint16 = 42
    var v5 bool = true
    v7 := GenS2{F0: 651}
    var v8 uint = (v1 / (((v1 & 18446744073709551615)) | 1))
    v9 := GenS2{F0: -1}
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS3{F0: (fi10 | fi10), F1: 32767}
        _ = fi10
        _ = v11
    }
    v12 := []int{32767, -1, 685, -802}
    var v13 uint64 = 426
    v13 += (v13 | v13)
    v1 = v1
    for fi14 := 0; fi14 < 5; fi14++ {
        if len(v12) > 0 {
            v12[0] = (fi14 >> (((v1 ^ 42)) & 63))
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            if len(v12) > 0 {
                v12[0] = (fi14 % (((-(fi15))) | 1))
            }
            if (v3 > (v3 << (((v1 << ((v8) & 63))) & 63))) {
                var v16 uint64 = v13
                if false {
                    if len(v12) > 0 {
                        v12[0] = 9223372036854775807
                    }
                } else {
                    switch fi15 {
                    case -573:
                        var v17 uint8 = 115
                        _ = v17
                    case 88:
                        v18 := GenS2{F0: 1000}
                        _ = v18
                    default:
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    }
                    var v19 bool = v5
                    var v20 float64 = (-2.25 / (0.1 + 1e10))
                    if v5 {
                        var v21 uint = 15
                        v3 = (^(v3))
                        if len(v12) > 0 {
                            v12[0] = 999
                        }
                        fmt.Println(v3, v16, v8)
                        _ = v21
                    }
                    _ = v19
                    _ = v20
                }
                if ((v4 / (((v4 / ((v4) | 1))) | 1)) == (^(v4))) {
                    var v22 uint32 = 0
                    _ = v22
                } else {
                    switch 9223372036854775807 {
                    case 0:
                        if (1e10 > -2.25) {
                            v3 -= int32(v4)
                        } else {
                            if len(v12) > 0 {
                                v12[0] = (fi15 ^ 0)
                            }
                            fmt.Println(v1, fi14)
                            var v23 int64 = v2
                            if true {
                                var v24 int32 = v3
                                _ = v24
                            } else {
                                v25 := GenS4{F0: v16, F1: fi15}
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                if len(v12) > 0 {
                                    v12[0] = (fi15 - 4294967295)
                                }
                                _ = v25
                            }
                            v26 := GenS3{F0: 479, F1: 1000}
                            v27 := GenS5{F0: -1}
                            v28 := GenS6{F0: (!(!(1 <= 2)))}
                            _ = v23
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    case -101:
                        v12 = append(v12, 0)
                    default:
                        var v29 int = (fi14 << (((^(v8))) & 63))
                        _ = v29
                    }
                    switch "\n\t" {
                    case "\n\t":
                        if len(v12) > 0 {
                            v12[0] = 42
                        }
                    }
                }
                fmt.Println(v4)
                _ = v16
            }
            var v30 int64 = (v2 % (((^(v2))) | 1))
            _ = fi15
            _ = v30
        }
        _ = fi14
    }
    if (v5 && v5) {
        switch (("hello" + "a") + "\n\t") {
        case "a":
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
        var v31 bool = (0.1 > 0.1)
        v4 -= (v4 | (^(v4)))
        _ = v31
    }
    fmt.Println(v2, v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
