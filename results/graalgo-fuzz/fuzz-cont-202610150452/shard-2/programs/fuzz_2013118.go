package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint8 = 0
    var v2 uint64 = 260
    var v3 uint32 = 868
    v4 := make(map[string]uint32)
    v5 := make(map[int]uint32)
    var v6 string = "hello"
    var v7 int16 = 42
    if (1 <= (v3 / ((v3) | 1))) {
        if (!true) {
            if true {
                switch 42 {
                case -1:
                    switch -436 {
                    case -2147483648:
                        var v8 bool = true
                        _ = v8
                    }
                case -508:
                    for fi9 := 0; fi9 < 13; fi9++ {
                        var v10 int8 = int8(v3)
                        var v11 uint64 = (v2 | v2)
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case -9223372036854775808:
                    var v12 uint = 4294967295
                    _ = v12
                }
                v13 := GenS1{F0: -128, F1: -9223372036854775808, F2: v3}
                for fi14 := 0; fi14 < 13; fi14++ {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        v16 := []int{}
                        var v17 uint = 54
                        if false {
                            var v18 int64 = 9223372036854775807
                            _ = v18
                        }
                        v17 = 0
                        v6 = v6
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                _ = v13
            } else {
                v2 = (v2 + 42)
                v19 := GenS1{F0: 127, F1: int64(v3), F2: 2147483647}
                v20 := GenS1{F0: -1, F1: -428, F2: (v3 | v3)}
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 int8 = -1
                var v23 string = "a"
                var v24 uint8 = (v1 | (v1 << ((0) & 63)))
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            v6 = (("\n\t" + v6) + v6)
        }
    }
    v5[127] = v3
    var v25 uint64 = (v2 * v2)
    for fi26 := 0; fi26 < 3; fi26++ {
        v6 = ((("\n\t" + "a") + ((v6 + (((v6 + (v6 + v6)) + v6) + v6)) + "hello")) + (v6 + v6))
        _ = fi26
    }
    v3 -= (v3 >> ((4294967295) & 63))
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS1{F0: -1, F1: 9223372036854775807, F2: v3}
        _ = fi27
        _ = v28
    }
    if ((!(-9223372036854775808 == 42)) || (0.0 > 0.0)) {
        v3 = (v3 ^ (^(v3)))
    } else {
        if (-1 != 23) {
            for fi29 := 0; fi29 < 8; fi29++ {
                if (1.5 == -2.25) {
                    fmt.Println(v7, fi29, v1)
                    var v30 bool = (false || true)
                    _ = v30
                }
                if true {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        v25 = (v25 % (((v2 % (((v2 + (^(v25)))) | 1))) | 1))
                        var v32 int32 = 0
                        _ = fi31
                        _ = v32
                    }
                    v33 := GenS1{F0: -1, F1: -1, F2: (^(v3))}
                    _ = v33
                }
                v6 = (v6 + v6)
                _ = fi29
            }
        } else {
            var v34 float64 = (0.1 - (0.0 * 1.5))
            _ = v34
        }
        var v35 uint = 18446744073709551615
        v1 += (^(v1))
        _ = v35
    }
    v36 := GenS2{F0: 792}
    v37 := []int{32767, -1}
    fmt.Println(v2)
    if (!true) {
        v38 := GenS1{F0: 0, F1: 1, F2: 42}
        _ = v38
    } else {
        var v39 int8 = -37
        _ = v39
    }
    fmt.Println(v3)
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
