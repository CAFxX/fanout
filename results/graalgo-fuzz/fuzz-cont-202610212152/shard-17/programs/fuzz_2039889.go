package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8 }
func genFn0() uint16 {
    var v4 uint16 = 0
    v5 := []int{-559}
    _ = v4
    _ = v5
    return (v4 + (v4 + 42))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 1
    var v3 bool = false
    if (v3 && v3) {
        delete(v1, 320)
        v6 := GenS1{F0: 1, F1: 80}
        _ = v6
    } else {
        v2 = (v2 | 0)
        if (!(65535 == 65535)) {
            fmt.Println(v2, v2)
            v7 := GenS1{F0: 42, F1: 127}
            v3 = (-75 < 1)
            _ = v7
        }
        fmt.Println(v2, v2)
    }
    fmt.Println(v3, v2)
    if false {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int16 = -954
            _ = fi8
            _ = v9
        }
    } else {
        var v10 bool = (v3 && v3)
        _ = v10
    }
    if false {
        v3 = (2 == 65535)
    } else {
        v11 := GenS1{F0: 4294967295, F1: 105}
        _ = v11
    }
    var v12 uint32 = 4294967295
    var v13 int8 = 42
    v12 = v12
    var v14 int64 = -9223372036854775808
    if (255 > 58) {
        for fi15 := 0; fi15 < 3; fi15++ {
            switch "foo bar" {
            case "":
                v2 += (^(v2))
            default:
                if (int(v12) == fi15) {
                    v16 := v1[v12]
                    fmt.Println(v2)
                    _ = v16
                } else {
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v17 := GenS1{F0: uint64(v12), F1: 255}
                        _ = v17
                    case "foo bar":
                        v2 += uint16(v12)
                    default:
                        if false {
                            for fi18 := 0; fi18 < 3; fi18++ {
                                if ((127 >= 210) || (2147483647 == 42)) {
                                    v12 -= 65535
                                    var v19 int16 = int16(v13)
                                    fmt.Println(v19, v19)
                                    var v20 int16 = v19
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 uint32 = v12
                                    _ = v21
                                }
                                _ = fi18
                            }
                        }
                    }
                }
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                if v3 {
                    v2 = v2
                    switch "a" {
                    case "foo bar":
                        v14 -= (^(v14))
                    case "":
                        var v23 uint8 = 128
                        _ = v23
                    case "a":
                        var v24 int = (fi15 / ((fi22) | 1))
                        _ = v24
                    default:
                        v25 := []int{(fi22 >> ((18446744073709551615) & 63)), (fi22 >> ((18446744073709551615) & 63)), (fi22 - fi22), 128}
                        _ = v25
                    }
                    if v3 {
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v14, v13, v3)
                        default:
                            var v26 float64 = 1e10
                            _ = v26
                        }
                    } else {
                        v2 += genFn0()
                        if (v3 || true) {
                            delete(v1, 65535)
                            var v27 bool = (!false)
                            v3 = false
                            _ = v27
                        }
                    }
                } else {
                    var v28 int64 = v14
                    _ = v28
                }
                var v29 int64 = v14
                if (!(uint64(v12) >= 18446744073709551615)) {
                    var v30 int = fi15
                    if v3 {
                        var v31 int16 = 1000
                        _ = v31
                    }
                    v12 += (v12 - v12)
                    _ = v30
                } else {
                    v13 = 42
                }
                delete(v1, 65535)
                _ = fi22
                _ = v29
            }
            v2 = 255
            _ = fi15
        }
    } else {
        var v32 int = 9223372036854775807
        var v33 int64 = (v14 >> ((1) & 63))
        _ = v32
        _ = v33
    }
    var v34 float32 = 1e10
    if (v3 || false) {
        for fi35 := 0; fi35 < 8; fi35++ {
            switch (fi35 | 263) {
            case -814:
                var v36 bool = false
                _ = v36
            case -662:
                v2 = 202
            case 127:
                v34 = (v34 - (float32(0.1) * v34))
            default:
                v37 := GenS2{F0: (v13 | (v13 | (v13 * v13)))}
                _ = v37
            }
            if ((^(v2)) == (v2 | 255)) {
                v38 := GenS2{F0: int8(v2)}
                _ = v38
            } else {
                switch "\n\t" {
                case "":
                    var v39 uint64 = 42
                    _ = v39
                default:
                    var v40 uint = 0
                    _ = v40
                }
                v3 = (!(!v3))
            }
            v13 += v13
            _ = fi35
        }
    }
    var v41 float64 = (1e10 * (1.5 * 1e10))
    var v42 uint8 = 243
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v41
    _ = v42
}
