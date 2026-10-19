package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int32) uint {
    fmt.Println(p0)
    switch -859 {
    case -390:
        for fi4 := 0; fi4 < 16; fi4++ {
            p0 = p0
            _ = fi4
        }
    case 42:
        p0 += p0
    case 32767:
        p0 = p0
    default:
        v5 := []int{-2147483648, 4294967295}
        _ = v5
    }
    _ = p0
    return 0
}

func main() {
    var v1 int16 = -166
    v2 := make(map[string]uint32)
    var v3 float64 = -2.25
    for fi6 := 0; fi6 < 4; fi6++ {
        v1 = 1000
        if true {
            var v7 uint8 = 174
            _ = v7
        } else {
            if false {
                for fi8 := 0; fi8 < 17; fi8++ {
                    switch -2147483648 {
                    case 9223372036854775807:
                        var v9 int64 = 1
                        _ = v9
                    default:
                        var v10 int32 = -1
                        _ = v10
                    }
                    v1 = 0
                    v2["hello"] = 275
                    _ = fi8
                }
            } else {
                v11 := []int{-11, fi6, (fi6 % ((fi6) | 1)), fi6}
                v12 := GenS1{F0: 2, F1: true}
                v1 += (v1 | 42)
                _ = v11
                _ = v12
            }
            if (0 == 255) {
                var v13 float64 = v3
                for fi14 := 0; fi14 < 13; fi14++ {
                    switch (fi14 ^ 32767) {
                    case 127:
                        if true {
                            v1 += v1
                            v15 := GenS1{F0: 2, F1: ((v3 / -2.25) > (v3 - 1.5))}
                            var v16 string = "foo bar"
                            _ = v15
                            _ = v16
                        } else {
                            v13 += v13
                        }
                    case -842:
                        var v17 float64 = float64(((1e10 + (3.14159 / float32((v3 - float64(0.0))))) - (0.1 - 1.5)))
                        _ = v17
                    case 32767:
                        var v18 float64 = 1e10
                        _ = v18
                    }
                    _ = fi14
                }
                _ = v13
            }
        }
        var v19 bool = false
        _ = fi6
        _ = v19
    }
    v3 -= 0.1
    for fi20 := 0; fi20 < 15; fi20++ {
        v1 += (v1 ^ v1)
        if false {
            v1 = v1
        } else {
            v3 += v3
        }
        _ = fi20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 string = "foo bar"
        delete(v2, (v22 + v22))
        _ = fi21
        _ = v22
    }
    v23 := []int{127, -1, 2147483647}
    v1 = v1
    for fi24 := 0; fi24 < 15; fi24++ {
        if false {
            fmt.Println(v3)
            if true {
                if false {
                    for fi25 := 0; fi25 < 14; fi25++ {
                        v26 := GenS2{F0: 908, F1: 32767}
                        _ = fi25
                        _ = v26
                    }
                    v3 = v3
                } else {
                    v1 -= (v1 + v1)
                    fmt.Println(fi24, fi24)
                }
            }
        }
        v2[((((("foo bar" + "foo bar") + "\n\t") + "foo bar") + "") + "a")] = 2147483648
        v3 += 0.0
        _ = fi24
    }
    v27 := GenS2{F0: 0, F1: 449}
    fmt.Println(v3, v3)
    var v28 int8 = 88
    var v29 uint32 = 0
    v30 := GenS1{F0: 42, F1: (789 <= 42)}
    var v31 int16 = (v1 >> ((2) & 63))
    var v32 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
