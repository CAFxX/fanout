package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int32 }
func genFn0() uint {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 int32 = -2147483648
        _ = fi4
        _ = v5
    }
    var v6 int16 = 1
    var v7 uint32 = 2147483648
    _ = v6
    _ = v7
    return 972
}

func main() {
    var v1 bool = true
    var v2 int32 = -574
    v3 := make(map[uint32]int)
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint32 = 2
        v1 = ((v1 && false) && v1)
        v10 := v3[(^(v9))]
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint = 512
    if (1.5 < 1e10) {
        fmt.Println(v11, v11)
        v11 -= (v11 >> (((v11 / ((v11) | 1))) & 63))
        if (!true) {
            var v12 uint16 = 609
            if (((!v1) && (0 < 127)) && (1 > 619)) {
                if v1 {
                    v2 += v2
                    if (!(-32768 <= -706)) {
                        if true {
                            var v13 float64 = (3.14159 * float64(3.14159))
                            _ = v13
                        } else {
                            v14 := []int{310}
                            _ = v14
                        }
                    }
                    v12 -= v12
                } else {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        switch (-(fi15)) {
                        case -9223372036854775808:
                            for fi16 := 0; fi16 < 14; fi16++ {
                                var v17 uint16 = (v12 / ((v12) | 1))
                                v18 := GenS1{F0: 1, F1: ("foo bar" + ("a" + ("a" + ("a" + "\n\t")))), F2: v2}
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        case 9223372036854775807:
                            v2 += (v2 ^ (-(v2)))
                        }
                        v12 -= v12
                        _ = fi15
                    }
                }
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := GenS1{F0: 141, F1: "hello", F2: (v2 * v2)}
                    _ = fi19
                    _ = v20
                }
                var v21 uint = v11
                _ = v21
            } else {
                if true {
                    v12 -= v12
                } else {
                    v22 := []int{}
                    v2 += 2147483647
                    _ = v22
                }
                _ = len(v3)
                fmt.Println(v11)
            }
            _ = v12
        } else {
            v23 := GenS1{F0: 42, F1: (("" + "a") + (("" + ("\n\t" + "hello")) + "a")), F2: (v2 & 426)}
            v24 := []int{532}
            _ = v23
            _ = v24
        }
    } else {
        v1 = v1
        var v25 uint16 = 42
        for fi26 := 0; fi26 < 16; fi26++ {
            v1 = (((v2 <= (v2 / ((v2) | 1))) || v1) || (-76 >= -9223372036854775808))
            var v27 uint32 = 42
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    v2 += 2147483647
    v28 := []int{-574, 127, 25}
    v3[584] = 128
    _ = len(v28)
    var v29 uint32 = 0
    switch -129 {
    case -962:
        v30 := GenS1{F0: 1, F1: "\n\t", F2: (v2 % (((v2 << ((uint(v29)) & 63))) | 1))}
        _ = v30
    default:
        v3[(^(v29))] = 128
    }
    var v31 int32 = (v2 - v2)
    var v32 int = 127
    _ = len(v28)
    v2 -= (-(v2))
    v11 = v11
    v33 := GenS1{F0: 2, F1: ("foo bar" + "a"), F2: -1}
    switch (v32 / ((v32) | 1)) {
    case 9223372036854775807:
        v34 := v3[v29]
        _ = v34
    default:
        v2 += 42
    }
    v3[v29] = (v32 & 32767)
    var v35 bool = ((!true) && v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
}
