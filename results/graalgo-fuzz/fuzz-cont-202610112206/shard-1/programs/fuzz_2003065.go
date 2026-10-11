package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 bool; F1 int }
func genFn0() uint64 {
    v4 := []int{42, -2147483648, -1}
    v5 := GenS1{F0: 184, F1: 148}
    var v6 int8 = 0
    _ = v4
    _ = v5
    _ = v6
    return 509
}
func genFn1(p0 uint16, p1 int16, p2 uint16) uint32 {
    if true {
        for fi7 := 0; fi7 < 11; fi7++ {
            switch "a" {
            case "":
                var v8 uint8 = 127
                _ = v8
            default:
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 int32 = 0
                    var v11 int16 = (p1 * p1)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            v12 := GenS1{F0: 766, F1: 2}
            _ = fi7
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int64 = 0
        var v15 uint8 = 0
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]string)
    var v3 bool = true
    var v16 int16 = 147
    v3 = (127 <= 42)
    var v17 int = 32767
    v18 := GenS2{F0: v3, F1: 1}
    v1 = (^(v1))
    var v19 int = v17
    switch ("" + "foo bar") {
    case "foo bar":
        v20 := GenS2{F0: v3, F1: 42}
        _ = v20
    case "":
        v21 := v2[624]
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 float64 = 1e10
        var v24 uint16 = 42
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v1 = (v1 >> ((18446744073709551615) & 63))
        for fi26 := 0; fi26 < 2; fi26++ {
            switch "hello" {
            case "a":
                if v3 {
                    v1 -= v1
                    if (0 == 42) {
                        var v27 int64 = 1
                        v1 += genFn0()
                        v16 += 42
                        _ = v27
                    } else {
                        v1 = 18446744073709551615
                        v17 += fi26
                        var v28 int32 = 350
                        _ = v28
                    }
                    v19 += fi25
                } else {
                    if v3 {
                        if ((^(v17)) > (v19 | (^(fi26)))) {
                            if (128 >= 152) {
                                fmt.Println(v3, v19, v16)
                                v16 += (v16 & -371)
                                var v29 uint16 = 1
                                _ = v29
                            }
                        } else {
                            if v3 {
                                var v30 int32 = 2147483647
                                var v31 int = (v17 ^ 629)
                                switch -346 {
                                case 127:
                                    var v32 float32 = float32(0.1)
                                    _ = v32
                                default:
                                    var v33 bool = (-1 < -112)
                                    _ = v33
                                }
                                _ = v30
                                _ = v31
                            }
                            fmt.Println(v19, fi26, v19)
                        }
                        var v34 int16 = (v16 & v16)
                        v34 -= (v34 & v16)
                        var v35 uint = 1
                        var v36 int16 = -8
                        switch (v17 >> (((^(v35))) & 63)) {
                        case -2147483648:
                            var v37 int = 584
                            _ = v37
                        default:
                            fmt.Println(v1, v16, v17)
                        }
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                }
            }
            var v38 int16 = 1
            var v39 string = "foo bar"
            _ = fi26
            _ = v38
            _ = v39
        }
        _ = fi25
    }
    v40 := GenS2{F0: ((v1 | (v1 + genFn0())) <= genFn0()), F1: (v17 << ((18446744073709551615) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v40
}
