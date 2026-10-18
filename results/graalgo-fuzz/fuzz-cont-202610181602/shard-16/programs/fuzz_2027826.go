package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int64, p1 string, p2 uint64) bool {
    var v6 float64 = 1.5
    v7 := GenS1{F0: 255, F1: (p1 + p1), F2: p2}
    switch -9223372036854775808 {
    case 187:
        p2 += 563
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]int)
    var v3 int32 = -977
    var v4 uint = 18446744073709551615
    var v5 float32 = (1e10 + 0.0)
    fmt.Println(v5)
    if false {
        var v8 bool = true
        v3 = (v3 | (v3 | v3))
        fmt.Println(v3, v1, v4)
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        switch ("foo bar" + ("foo bar" + ("a" + ("a" + "")))) {
        case "\n\t":
            for fi10 := 0; fi10 < 19; fi10++ {
                v1 += (v1 - (^(v1)))
                _ = fi10
            }
        }
        fmt.Println(fi9, v5, v3)
        v5 = 3.14159
        _ = fi9
    }
    fmt.Println(v4, v5)
    var v11 uint8 = 42
    fmt.Println(v1)
    v12 := GenS1{F0: 971, F1: "hello", F2: uint64(v1)}
    var v13 uint8 = (v11 % ((127) | 1))
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(v5, v5, v13)
        var v15 uint16 = 42
        v16 := []int{fi14}
        _ = fi14
        _ = v15
        _ = v16
    }
    if false {
        v5 = (float32(-2.25) * ((v5 * (v5 - v5)) / (3.14159 - -2.25)))
        v5 += ((0.0 + v5) - v5)
        v5 -= v5
    } else {
        v4 += 65535
    }
    var v17 uint8 = 148
    if ((v1 | 2) >= v1) {
        v1 = v1
        if false {
            if false {
                switch "" {
                case "\n\t":
                    if false {
                        var v18 int16 = 0
                        _ = v18
                    } else {
                        switch 2147483647 {
                        case -380:
                            var v19 uint8 = v17
                            _ = v19
                        case 0:
                            fmt.Println(v11, v17, v4)
                        default:
                            var v20 bool = false
                            _ = v20
                        }
                        v11 = (v17 % (((v13 << (((v4 << ((641) & 63))) & 63))) | 1))
                        v21 := GenS1{F0: 1, F1: "\n\t", F2: 4294967295}
                        _ = v21
                    }
                case "":
                    v2[4294967295] = -2147483648
                case "hello":
                    v22 := GenS1{F0: 121, F1: (("" + "a") + "a"), F2: 858}
                    _ = v22
                default:
                    v2[-350] = 32767
                }
                var v23 int32 = v3
                v24 := v2[0]
                _ = v23
                _ = v24
            }
            switch 128 {
            case -2147483648:
                var v25 bool = false
                _ = v25
            case 0:
                fmt.Println(v17, v4, v1)
            default:
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 uint8 = (v13 / ((143) | 1))
                    var v28 uint8 = 42
                    for fi29 := 0; fi29 < 11; fi29++ {
                        if (true && (v3 >= (-(v3)))) {
                            switch "hello" {
                            case "\n\t":
                                v2[fi26] = 42
                                fmt.Println(v3, v27, fi26)
                                if ((true || (!false)) || true) {
                                    v30 := []int{fi29, fi26, fi29, 128}
                                    _ = v30
                                }
                            }
                            v11 = v11
                            v31 := v2[(fi29 + (fi29 >> ((v4) & 63)))]
                            _ = v31
                        }
                        var v32 uint8 = (v17 * 246)
                        if (-1 < -1) {
                            v5 -= 1e10
                            v33 := []int{(^(fi29))}
                            var v34 int = fi26
                            _ = v33
                            _ = v34
                        } else {
                            var v35 float64 = 0.0
                            v36 := []int{(fi29 / (((fi29 - int(v3))) | 1))}
                            _ = v35
                            _ = v36
                        }
                        _ = fi29
                        _ = v32
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            switch "foo bar" {
            case "":
                switch "hello" {
                case "a":
                    switch -1 {
                    case 128:
                        v4 += v4
                    case -1:
                        v4 = (v4 | (v4 * 0))
                    }
                default:
                    v37 := GenS2{F0: 108}
                    _ = v37
                }
            default:
                v38 := []int{32768, 0, -1, 32767}
                _ = v38
            }
        }
    } else {
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 int = fi39
            v4 -= (v4 | 251)
            _ = fi39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
