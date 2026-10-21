package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 bool }

func main() {
    v1 := []int{923, -58, 1, 2147483647}
    v2 := make(map[string]int)
    var v3 int64 = -9223372036854775808
    v4 := []int{}
    v5 := []int{-655}
    var v6 float32 = 1e10
    var v7 int8 = 1
    if true {
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                if false {
                    var v10 bool = true
                    v11 := GenS1{F0: 0}
                    var v12 uint32 = 278
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    var v13 float64 = 1e10
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v14 := GenS2{F0: 42, F1: (v7 ^ v7), F2: (v3 ^ int64(v7))}
                        _ = v14
                    } else {
                        v7 += (v7 >> ((2) & 63))
                        var v15 int8 = 109
                        _ = v15
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v13
                }
                v6 = (((v6 / v6) * (1e10 / v6)) + v6)
                _ = fi9
            }
            v7 += (v7 + 44)
            _ = fi8
        }
        var v16 int = 2147483647
        _ = v16
    } else {
        v17 := GenS2{F0: v7, F1: v7, F2: v3}
        v18 := GenS3{F0: true, F1: 32767, F2: 42}
        for fi19 := 0; fi19 < 15; fi19++ {
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(v6, v3, v6)
                var v21 uint64 = 4294967295
                v7 = -40
                _ = fi20
                _ = v21
            }
            v6 = float32(0.1)
            _ = fi19
        }
        _ = v17
        _ = v18
    }
    if (!false) {
        v22 := GenS1{F0: v7}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 2; fi23++ {
            for fi24 := 0; fi24 < 4; fi24++ {
                v3 = (v3 << ((65535) & 63))
                v25 := GenS4{F0: 42, F1: 18446744073709551615, F2: ((6 > 175) || false)}
                _ = fi24
                _ = v25
            }
            if false {
                var v26 int16 = int16(v3)
                if (1 >= v3) {
                    v27 := GenS1{F0: -1}
                    v6 = 3.14159
                    v26 = v26
                    _ = v27
                }
                _ = v26
            }
            delete(v2, ("" + ""))
            _ = fi23
        }
        v28 := GenS3{F0: (!(-9223372036854775808 == 216)), F1: 32767, F2: 190}
        _ = v28
    }
    fmt.Println(v6)
    var v29 int16 = -242
    v6 += float32(0.0)
    if false {
        if ("a" <= (("hello" + "a") + "")) {
            v4 = append(v4, -9223372036854775808)
        }
        v30 := GenS2{F0: (v7 + (v7 - (v7 * v7))), F1: v7, F2: (v3 | (v3 / (((v3 | (v3 ^ (v3 & v3)))) | 1)))}
        v31 := v2["hello"]
        _ = v30
        _ = v31
    }
    switch 1 {
    case 127:
        v6 += ((0.1 - v6) - (((1.5 * 3.14159) + 0.0) / v6))
    case 0:
        v32 := GenS4{F0: v3, F1: 18446744073709551615, F2: (v3 > 0)}
        _ = v32
    case 1:
        v5 = append(v5, 42)
    }
    var v33 string = "a"
    var v34 int8 = (v7 << ((4294967295) & 63))
    if ((65535 > 0) || true) {
        var v35 uint64 = 4294967295
        _ = v35
    }
    var v36 uint16 = 42
    v37 := GenS5{F0: (v6 <= (v6 - -2.25))}
    v38 := GenS3{F0: (65535 == 65535), F1: v36, F2: 136}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
