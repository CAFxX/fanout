package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int64 }
func genFn0(p0 int64, p1 int64, p2 uint8) int8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        if ((0 > 277) && true) {
            p2 = (p2 ^ 155)
        }
        _ = fi7
    }
    fmt.Println(p1, p1)
    var v8 uint32 = 752
    fmt.Println(p1, p2, v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 100
}
func genFn1(p0 uint64, p1 int16, p2 int16) uint32 {
    var v9 bool = true
    fmt.Println(p2, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 65535
}

func main() {
    var v1 float32 = 0.1
    var v2 float64 = 1e10
    v3 := []int{}
    var v4 int32 = -11
    v5 := make(map[int]int)
    v6 := make(map[uint32]bool)
    fmt.Println(v1)
    if (2 >= 18446744073709551615) {
        if (!false) {
            if (2 < 2) {
                switch "foo bar" {
                case "a":
                    v4 -= v4
                case "foo bar":
                    switch 2147483647 {
                    case 0:
                        for fi10 := 0; fi10 < 17; fi10++ {
                            _ = len(v6)
                            _ = fi10
                        }
                    case 4294967295:
                        _ = len(v5)
                    case 128:
                        v11 := GenS1{F0: 255, F1: genFn1(2, 635, -823), F2: 32767}
                        _ = v11
                    }
                }
                var v12 int16 = -1
                var v13 int16 = v12
                _ = v12
                _ = v13
            }
            var v14 uint16 = 32767
            _ = v14
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            fmt.Println(v1)
            v1 -= (v1 * float32(1e10))
            for fi16 := 0; fi16 < 9; fi16++ {
                if (false || false) {
                    var v17 uint16 = 42
                    v17 += v17
                    if len(v3) > 0 {
                        v3[0] = (fi15 ^ 32767)
                    }
                    _ = v17
                } else {
                    var v18 uint32 = genFn1(295, 1, -32768)
                    if false {
                        v19 := GenS2{F0: 0, F1: (v18 * 65535), F2: 18446744073709551615}
                        var v20 uint16 = 736
                        if len(v3) > 0 {
                            v3[0] = fi16
                        }
                        _ = v19
                        _ = v20
                    }
                    _ = v18
                }
                delete(v6, genFn1(18446744073709551615, 42, 42))
                _ = fi16
            }
            _ = fi15
        }
    }
    v21 := GenS3{F0: 32768}
    v2 = float64((v1 - v1))
    v2 -= 0.0
    var v22 int = 406
    var v23 uint32 = genFn1(2, 29, 1)
    v1 = v1
    if ((v4 - -2147483648) > v4) {
        v24 := GenS4{F0: (-(v22)), F1: 42}
        _ = v24
    } else {
        if false {
            if true {
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := GenS4{F0: fi25, F1: -138}
                    _ = fi25
                    _ = v26
                }
                switch "\n\t" {
                case "hello":
                    v2 += (float64((v1 - 3.14159)) - v2)
                case "foo bar":
                    var v27 uint16 = 545
                    _ = v27
                }
            } else {
                var v28 float64 = v2
                _ = len(v6)
                _ = v28
            }
            if false {
                v29 := GenS2{F0: -1, F1: genFn1(0, 42, -187), F2: 2}
                for fi30 := 0; fi30 < 6; fi30++ {
                    v1 -= float32(float64(v1))
                    for fi31 := 0; fi31 < 14; fi31++ {
                        v1 += ((v1 / v1) + ((1e10 - 3.14159) - 0.0))
                        v4 -= (v4 + 0)
                        var v32 int64 = -1
                        _ = fi31
                        _ = v32
                    }
                    _ = fi30
                }
                _ = v29
            }
            switch "\n\t" {
            case "a":
                var v33 int32 = -2147483648
                _ = v33
            case "hello":
                v3 = append(v3, v22)
            }
        }
    }
    if (false && true) {
        v4 = (v4 - v4)
    }
    var v34 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
