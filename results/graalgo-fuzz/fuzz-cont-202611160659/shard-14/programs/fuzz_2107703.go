package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint; F2 int8 }
func genFn0(p0 uint16, p1 uint8) int {
    switch 980 {
    case 1:
        var v6 int64 = 217
        _ = v6
    default:
        fmt.Println(p1)
    }
    v7 := []int{548, 2147483647, 2147483647, -190}
    _ = p0
    _ = p1
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 int64 = 0
    var v4 uint64 = 1
    v5 := make(map[string]int)
    var v8 uint64 = 1
    v9 := GenS1{F0: -491, F1: 32767}
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(fi10)
        switch 2147483647 {
        case 1:
            v4 += v8
        default:
            v3 -= (v3 << ((65535) & 63))
        }
        if (!((711 != 42) && (!(!(float64(1e10) != float64(1.5)))))) {
            switch -890 {
            case -1:
                v11 := GenS2{F0: 127}
                _ = v11
            }
        } else {
            var v12 uint16 = 1
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        fmt.Println(v4, fi13)
        _ = fi13
    }
    if ((((-(v3)) <= v3) && (!false)) || true) {
        v14 := []int{32768, 32767}
        v2[42] = 355
        if len(v14) > 0 {
            v14[0] = 128
        }
        _ = v14
    } else {
        v3 -= v3
        var v15 bool = (!true)
        switch 0 {
        case 32767:
            v4 -= v4
        case 32768:
            switch genFn0(795, 195) {
            case 32767:
                v16 := []int{genFn0(113, 6), 127}
                _ = v16
            case 4294967295:
                if (v15 && ((v4 | (v8 & (v8 >> ((255) & 63)))) == 646)) {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        if v15 {
                            v15 = true
                            switch genFn0(0, 114) {
                            case 128:
                                if false {
                                    for fi18 := 0; fi18 < 16; fi18++ {
                                        v3 = 42
                                        var v19 float32 = (3.14159 + float32(0.0))
                                        v20 := GenS2{F0: 113}
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                } else {
                                    var v21 int = 584
                                    _ = v21
                                }
                            default:
                                fmt.Println(fi17)
                            }
                        } else {
                            v15 = v15
                        }
                        _ = fi17
                    }
                    for fi22 := 0; fi22 < 13; fi22++ {
                        if v15 {
                            if true {
                                var v23 int8 = -106
                                fmt.Println(v15)
                                var v24 uint8 = 1
                                _ = v23
                                _ = v24
                            }
                        } else {
                            v25 := []int{421, (-(fi22)), int(v8)}
                            _ = v25
                        }
                        v8 = v4
                        for fi26 := 0; fi26 < 13; fi26++ {
                            v15 = (uint16(v8) >= 365)
                            _ = fi26
                        }
                        _ = fi22
                    }
                } else {
                    v27 := GenS2{F0: 1}
                    var v28 uint64 = 18446744073709551615
                    v29 := GenS1{F0: 1, F1: 354}
                    _ = v27
                    _ = v28
                    _ = v29
                }
            case 32768:
                v15 = (1 > 1)
            default:
                v8 = (^(v8))
            }
        default:
            var v30 uint8 = 228
            _ = v30
        }
        _ = v15
    }
    var v31 uint32 = 1
    v32 := GenS3{F0: v8, F1: uint(v31), F2: -128}
    var v33 uint8 = uint8(v31)
    for fi34 := 0; fi34 < 1; fi34++ {
        fmt.Println(fi34)
        _ = fi34
    }
    v3 -= v3
    v35 := v5["foo bar"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
