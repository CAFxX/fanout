package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    var v1 int8 = 0
    v2 := make(map[int]int)
    var v3 int16 = -624
    var v4 float64 = -2.25
    var v5 bool = (0 > 0)
    v6 := GenS1{F0: v3, F1: v3}
    v7 := GenS1{F0: v3, F1: v3}
    var v8 float32 = (1.5 + (3.14159 - 0.1))
    var v9 int32 = 1
    var v10 uint = 452
    v11 := GenS1{F0: v3, F1: v3}
    var v12 uint32 = 65535
    v12 = v12
    var v13 int16 = 32767
    fmt.Println(v9, v8, v8)
    v13 = v3
    for fi14 := 0; fi14 < 19; fi14++ {
        if (0.0 < 1e10) {
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := GenS1{F0: v13, F1: -32768}
                v17 := []int{fi14, (fi15 % ((fi14) | 1))}
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 20; fi18++ {
                if v5 {
                    v19 := []int{}
                    switch "" {
                    case "hello":
                        var v20 uint16 = 587
                        _ = v20
                    }
                    _ = v19
                } else {
                    if false {
                        fmt.Println(v5, v1, fi14)
                    } else {
                        if (v8 >= ((0.0 + float32(0.1)) + v8)) {
                            v5 = ((!v5) && v5)
                            if ((fi14 >> ((0) & 63)) >= fi18) {
                                v2[fi14] = (fi14 + 898)
                            } else {
                                v9 -= v9
                                delete(v2, -2147483648)
                            }
                            if true {
                                v21 := GenS2{F0: (v9 % ((v9) | 1)), F1: (v1 + (v1 & v1)), F2: 1}
                                switch (^(fi18)) {
                                case 990:
                                    v8 -= 1.5
                                case 128:
                                    fmt.Println(v4)
                                    var v22 int16 = (v3 ^ v13)
                                    _ = v22
                                default:
                                    var v23 bool = true
                                    _ = v23
                                }
                                v24 := []int{(^(fi14)), (fi18 + 128), (fi14 << ((v10) & 63)), fi18}
                                _ = v21
                                _ = v24
                            }
                            v4 += v4
                        }
                    }
                    if ((v1 << (((v10 ^ v10)) & 63)) <= (^(v1))) {
                        if (((v4 - v4) >= ((v4 + float64(1e10)) * (float64(0.0) / v4))) || (-9223372036854775808 > 0)) {
                            fmt.Println(fi18, v5)
                            fmt.Println(v3, v12, v4)
                            var v25 uint16 = 2
                            _ = v25
                        }
                        _ = len(v2)
                        if (2 == 42) {
                            v26 := GenS2{F0: (v9 + 42), F1: v1, F2: -9223372036854775808}
                            v5 = false
                            _ = v26
                        } else {
                            switch (fi18 + fi18) {
                            case 32767:
                                fmt.Println(v1)
                            }
                            var v27 uint8 = 127
                            if ((2147483647 != 631) && (1 == 18446744073709551615)) {
                                v5 = false
                                if (v10 >= v10) {
                                    var v28 int8 = v1
                                    _ = v28
                                }
                                var v29 int16 = (v3 << (((v10 << (((v10 ^ (v10 / ((v10) | 1)))) & 63))) & 63))
                                _ = v29
                            }
                            _ = v27
                        }
                        v30 := []int{fi14}
                        var v31 float32 = v8
                        _ = v30
                        _ = v31
                    } else {
                        v32 := v2[(^(fi14))]
                        if (!v5) {
                            fmt.Println(v10, v5, v5)
                        } else {
                            var v33 int16 = v13
                            v1 += v1
                            v34 := GenS1{F0: (v13 << (((v10 % ((v10) | 1))) & 63)), F1: (v13 ^ 1000)}
                            _ = v33
                            _ = v34
                        }
                        _ = v32
                    }
                    var v35 string = "a"
                    v36 := GenS3{F0: false, F1: (v3 << ((439) & 63))}
                    _ = v35
                    _ = v36
                }
                _ = fi18
            }
        }
        _ = fi14
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        v38 := []int{fi37, (fi37 * fi37)}
        _ = fi37
        _ = v38
    }
    delete(v2, 409)
    for fi39 := 0; fi39 < 10; fi39++ {
        v40 := []int{}
        v8 = -2.25
        v2[(fi39 | 32768)] = (fi39 / (((fi39 & 203)) | 1))
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
