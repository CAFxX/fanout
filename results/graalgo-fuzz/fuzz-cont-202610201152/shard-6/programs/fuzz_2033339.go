package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint }

func main() {
    v1 := []int{464, 9223372036854775807, 28}
    var v2 float32 = 0.1
    var v3 float64 = 1e10
    v4 := []uint32{}
    for fi5 := 0; fi5 < 1; fi5++ {
        v3 += ((float64(v2) + v3) * (v3 / v3))
        _ = fi5
    }
    v3 += 1.5
    fmt.Println(v3, v2)
    v6 := GenS1{F0: (!false)}
    var v7 int32 = -1
    var v8 uint8 = 50
    v8 -= v8
    fmt.Println(v8, v2, v3)
    var v9 bool = false
    if (4294967295 > -9223372036854775808) {
        v4 = append(v4, 0)
        fmt.Println(v9, v3, v9)
    }
    v10 := GenS1{F0: (!((v7 | v7) > 2147483647))}
    v8 += (v8 * v8)
    for fi11 := 0; fi11 < 17; fi11++ {
        if len(v4) > 0 {
            v4[0] = 42
        }
        _ = fi11
    }
    var v12 uint8 = v8
    _ = len(v1)
    if len(v4) > 0 {
        v4[0] = 1
    }
    var v13 uint32 = 2
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: (true && (v12 == (v8 % ((214) | 1))))}
            var v17 uint32 = (v13 * (v13 & v13))
            v12 += v8
            if len(v4) > 0 {
                v4[0] = v17
            }
            switch fi15 {
            case -54:
                v18 := GenS2{F0: int16(v13)}
                _ = v18
            case -2147483648:
                v7 += 453
                if v9 {
                    v19 := GenS3{F0: 1, F1: 1000, F2: 42}
                    if v9 {
                        v3 += v3
                    }
                    _ = v19
                } else {
                    if len(v1) > 0 {
                        v1[0] = fi15
                    }
                    var v20 uint8 = 42
                    if ("hello" < "a") {
                        fmt.Println(v17)
                        if v9 {
                            if (3.14159 > (v3 + v3)) {
                                if (18446744073709551615 < 4294967295) {
                                    fmt.Println(v2, v20, fi14)
                                    v12 = uint8(v13)
                                    switch (fi14 & fi15) {
                                    case 2147483647:
                                        v17 = 213
                                    default:
                                        v8 -= 127
                                    }
                                    if false {
                                        v12 = v8
                                    } else {
                                        v3 = ((v3 / v3) * 1e10)
                                        _ = len(v4)
                                        var v21 int8 = -38
                                        _ = v21
                                    }
                                }
                                v17 = v17
                                var v22 int64 = -9223372036854775808
                                _ = v22
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            _ = len(v4)
                        }
                    } else {
                        switch "\n\t" {
                        case "hello":
                            if len(v4) > 0 {
                                v4[0] = v17
                            }
                        case "a":
                            v8 = v12
                        }
                        _ = len(v1)
                        if (!v9) {
                            if true {
                                v23 := GenS2{F0: -473}
                                if len(v4) > 0 {
                                    v4[0] = v13
                                }
                                var v24 uint = 4294967295
                                _ = v23
                                _ = v24
                            }
                        }
                    }
                    _ = v20
                }
            default:
                var v25 uint = 2
                _ = v25
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
