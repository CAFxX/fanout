package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint, p1 int64) uint16 {
    v5 := GenS1{F0: 42}
    fmt.Println(p1)
    p1 -= p1
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint16 = 0
        var v8 float32 = 1e10
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 32768
    v3 := make(map[string]string)
    var v4 uint32 = 442
    v4 -= v4
    v2 = v2
    v9 := []int{-637, v2, 399}
    var v10 int16 = 42
    v2 -= int(v4)
    _ = len(v1)
    for fi11 := 0; fi11 < 15; fi11++ {
        switch "foo bar" {
        case "a":
            for fi12 := 0; fi12 < 20; fi12++ {
                if false {
                    var v13 uint64 = 1
                    v14 := GenS1{F0: (^(v4))}
                    if (true || false) {
                        var v15 int = fi12
                        switch v2 {
                        case 91:
                            _ = len(v1)
                        case 9223372036854775807:
                            var v16 int16 = (v10 % ((v10) | 1))
                            _ = v16
                        case 2147483647:
                            fmt.Println(fi11, v13, v2)
                        default:
                            v17 := GenS1{F0: 2147483648}
                            _ = v17
                        }
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = v15
                    }
                    var v18 uint8 = 0
                    _ = v13
                    _ = v14
                    _ = v18
                } else {
                    v19 := GenS1{F0: v4}
                    v2 += 552
                    _ = v19
                }
                var v20 int = 9223372036854775807
                var v21 bool = (false && false)
                v20 += fi11
                v20 -= (fi12 / ((-452) | 1))
                _ = fi12
                _ = v20
                _ = v21
            }
        case "":
            if (true || false) {
                for fi22 := 0; fi22 < 18; fi22++ {
                    v1[int(v10)] = false
                    if (100 > 28) {
                        if (false && ((false && false) && (v4 >= v4))) {
                            if (0 < 42) {
                                var v23 uint32 = (v4 - 4294967295)
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                _ = v23
                            } else {
                                v10 += v10
                            }
                            v4 += 2147483648
                            switch "\n\t" {
                            case "\n\t":
                                var v24 bool = false
                                _ = v24
                            default:
                                v4 = (v4 & v4)
                            }
                            if (1 == 1) {
                                if ((fi22 < (-(v2))) && false) {
                                    var v25 int32 = -248
                                    v10 += (v10 >> ((255) & 63))
                                    var v26 uint16 = 42
                                    _ = v25
                                    _ = v26
                                }
                            } else {
                                if len(v9) > 0 {
                                    v9[0] = -2147483648
                                }
                            }
                            v27 := GenS1{F0: (^(v4))}
                            v28 := GenS2{F0: (-(v10)), F1: 42}
                            v29 := GenS2{F0: -246, F1: (v4 * v4)}
                            fmt.Println(fi22, fi11)
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            v2 += (v2 / (((-(fi22))) | 1))
                            v4 = (^(v4))
                            v2 -= (v2 - v2)
                        }
                        if (2 >= 255) {
                            var v30 int32 = -2147483648
                            var v31 uint64 = 42
                            _ = v30
                            _ = v31
                        } else {
                            fmt.Println(v2)
                            var v32 string = ("a" + ("hello" + "\n\t"))
                            _ = v32
                        }
                        v33 := GenS1{F0: v4}
                        _ = v33
                    }
                    if (false && ((^(v2)) != 32767)) {
                        _ = len(v3)
                        var v34 float64 = float64((0.0 + 0.0))
                        _ = v34
                    } else {
                        v4 -= (^(v4))
                        v35 := GenS2{F0: (^(v10)), F1: (v4 % ((v4) | 1))}
                        v10 -= (v10 * 42)
                        _ = v35
                    }
                    _ = fi22
                }
                v10 = v10
            }
        default:
            for fi36 := 0; fi36 < 17; fi36++ {
                v10 = 32767
                v37 := GenS2{F0: v10, F1: 2147483647}
                _ = len(v3)
                _ = fi36
                _ = v37
            }
        }
        _ = fi11
    }
    var v38 int16 = (v10 % ((v10) | 1))
    v38 += 32767
    v39 := GenS2{F0: v38, F1: (v4 | 42)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v9
}
