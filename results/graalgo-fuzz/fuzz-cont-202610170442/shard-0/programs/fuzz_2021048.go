package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 string; F1 uint8; F2 bool }
func genFn0(p0 uint) uint64 {
    switch 151 {
    case -880:
        if true {
            v6 := []int{32768, 127, 2147483647}
            p0 -= (p0 % (((p0 >> (((p0 | p0)) & 63))) | 1))
            _ = v6
        }
    case 774:
        v7 := []int{-9223372036854775808}
        _ = v7
    default:
        v8 := []int{2147483647, 32768, 32768, 377}
        _ = v8
    }
    v9 := GenS1{F0: "hello", F1: true, F2: p0}
    var v10 int32 = -787
    v11 := []int{-587, 32767, -862}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    v3 := []uint32{2147483647, 0, 17}
    v4 := make(map[int]string)
    v5 := make(map[int]string)
    if (282 <= 1000) {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(fi12)
        var v13 uint = 4294967295
        _ = fi12
        _ = v13
    }
    if v2 {
        var v14 int8 = 0
        v15 := GenS1{F0: "a", F1: true, F2: 14}
        _ = v14
        _ = v15
    } else {
        v16 := GenS2{F0: 0, F1: 2}
        v2 = true
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        fmt.Println(fi17, fi17)
        v2 = v2
        var v18 int16 = 590
        _ = fi17
        _ = v18
    }
    var v19 bool = v2
    var v20 uint16 = 1
    v19 = true
    var v21 string = "\n\t"
    var v22 int8 = -1
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            if v2 {
                var v25 uint8 = 255
                _ = v25
            }
            if (false || v2) {
                v26 := GenS2{F0: 4294967295, F1: (v20 & (^(v20)))}
                fmt.Println(v21)
                _ = v26
            }
            if true {
                switch -1 {
                case -1:
                    var v27 bool = (!(255 < 250))
                    if ((fi24 - -9223372036854775808) <= fi23) {
                        var v28 string = "hello"
                        v29 := GenS2{F0: 263, F1: 32767}
                        switch "hello" {
                        case "hello":
                            v22 -= 8
                        default:
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                        }
                        _ = v28
                        _ = v29
                    }
                    v22 += v22
                    _ = v27
                case 320:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 128:
                    v20 = (v20 % (((v20 << ((18446744073709551615) & 63))) | 1))
                default:
                    var v30 int32 = 0
                    _ = v30
                }
                var v31 int16 = int16(fi24)
                var v32 float64 = (3.14159 + 0.0)
                _ = v31
                _ = v32
            }
            _ = fi24
        }
        v33 := GenS3{F0: "a", F1: 255, F2: (1.5 < 1e10)}
        v22 += v22
        _ = fi23
        _ = v33
    }
    v34 := GenS1{F0: v21, F1: v19, F2: 241}
    var v35 uint64 = 653
    for fi36 := 0; fi36 < 13; fi36++ {
        var v37 int32 = -2147483648
        for fi38 := 0; fi38 < 16; fi38++ {
            switch (-(fi36)) {
            case 42:
                switch v21 {
                case "\n\t":
                    v21 = ("\n\t" + ((v21 + v21) + ("" + v21)))
                }
            }
            v21 = "a"
            _ = fi38
        }
        var v39 int = fi36
        _ = fi36
        _ = v37
        _ = v39
    }
    v21 = "a"
    var v40 uint16 = 566
    var v41 int8 = (v22 ^ (v22 >> ((uint(v40)) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
