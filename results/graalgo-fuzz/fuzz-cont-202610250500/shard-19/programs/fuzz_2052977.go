package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 bool; F2 int64 }
func genFn0(p0 bool, p1 int32) uint8 {
    fmt.Println(p0)
    if (774 > 887) {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: 65535}
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 32767}
        fmt.Println(p0, p0)
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        p1 -= (p1 / ((-2147483648) | 1))
        _ = fi10
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 127
    v3 := make(map[string]int)
    var v4 uint = 18446744073709551615
    var v5 float32 = float32(-2.25)
    var v6 uint = 198
    var v11 int32 = 28
    v2 -= v2
    var v12 int32 = int32(v4)
    var v13 bool = false
    switch "" {
    case "\n\t":
        switch "hello" {
        case "foo bar":
            var v14 int16 = 1000
            _ = v14
        case "a":
            v11 += (v11 | v12)
        default:
            v15 := GenS1{F0: 32767}
            _ = v15
        }
    default:
        switch "\n\t" {
        case "a":
            switch "" {
            case "a":
                v16 := GenS1{F0: uint16(v6)}
                _ = v16
            case "hello":
                for fi17 := 0; fi17 < 1; fi17++ {
                    switch "\n\t" {
                    case "\n\t":
                        v18 := []int{}
                        _ = v18
                    default:
                        v19 := []int{}
                        _ = v19
                    }
                    v5 -= float32((((float64(v5) / 3.14159) + (0.0 + float64(((v5 + (v5 - (v5 + 3.14159))) * (float32((0.0 / 1.0)) - (((v5 / 1.0) + ((v5 * 1e10) / 1.0)) * (1e10 + 1.5))))))) - ((0.1 * (-2.25 / 1.5)) + (1e10 + 0.1))))
                    if (2147483647 != (v12 / ((680) | 1))) {
                        v12 = -2147483648
                    } else {
                        if (!(v13 && true)) {
                            v20 := GenS1{F0: 1}
                            if (879 < 42) {
                                var v21 bool = true
                                _ = v21
                            } else {
                                v22 := GenS1{F0: 1}
                                fmt.Println(v5, v4)
                                switch fi17 {
                                case 4294967295:
                                    switch "" {
                                    case "a":
                                        v4 = (v4 * v6)
                                    default:
                                        v13 = v13
                                    }
                                }
                                _ = v22
                            }
                            _ = v20
                        }
                        for fi23 := 0; fi23 < 4; fi23++ {
                            fmt.Println(v13, v11)
                            v13 = (v13 || v13)
                            _ = fi23
                        }
                        var v24 int32 = (v11 << ((2) & 63))
                        _ = v24
                    }
                    _ = fi17
                }
            }
        default:
            if (850 != 42) {
                var v25 int8 = -30
                _ = v25
            } else {
                var v26 uint8 = v2
                v26 = 1
                v27 := GenS2{F0: 18446744073709551615, F1: false, F2: 0}
                _ = v26
                _ = v27
            }
        }
    }
    var v28 uint64 = 4294967295
    var v29 int16 = 32767
    if v13 {
        v30 := []int{32768, -2147483648}
        _ = v30
    } else {
        v13 = ((false && true) && (42 > 544))
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 float32 = float32(1e10)
            _ = fi31
            _ = v32
        }
    }
    var v33 int32 = (v11 << (((v6 ^ (v6 & 243))) & 63))
    v29 -= v29
    v34 := v1["foo bar"]
    if v13 {
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 uint32 = 65535
            _ = fi35
            _ = v36
        }
    }
    v2 = genFn0((!v13), v33)
    v34 = (v34 << (((v6 | (^(v4)))) & 63))
    var v37 uint16 = 32767
    var v38 uint64 = (^(v28))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
