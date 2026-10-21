package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
func genFn0(p0 int, p1 int64) int16 {
    var v6 uint32 = 0
    v7 := []int{9223372036854775807, -1, 0}
    if ((!true) || ("a" >= "")) {
        v8 := GenS1{F0: (^(p1)), F1: 42}
        _ = v8
    } else {
        fmt.Println(v6)
        v6 = (v6 >> ((2) & 63))
        var v9 int = 128
        _ = v9
    }
    if false {
        if false {
            v10 := GenS1{F0: (p1 ^ p1), F1: 127}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                if ("" <= "\n\t") {
                    p0 += (p0 - (^(fi11)))
                    for fi12 := 0; fi12 < 9; fi12++ {
                        switch fi12 {
                        case -977:
                            if len(v7) > 0 {
                                v7[0] = 2147483647
                            }
                        default:
                            var v13 string = "hello"
                            _ = v13
                        }
                        if (!(-32768 > 845)) {
                            v7 = append(v7, -1)
                            fmt.Println(v6, v6)
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            v14 := GenS2{F0: 167, F1: 127}
                            _ = v14
                        }
                        v6 -= (^(v6))
                        _ = fi12
                    }
                    if len(v7) > 0 {
                        v7[0] = (fi11 | p0)
                    }
                }
                var v15 int = 0
                v6 = v6
                _ = fi11
                _ = v15
            }
            var v16 float32 = -2.25
            _ = v16
        }
        v17 := GenS1{F0: p1, F1: 1}
        if false {
            for fi18 := 0; fi18 < 4; fi18++ {
                p0 += 127
                v6 -= (v6 >> ((374) & 63))
                _ = fi18
            }
            p0 -= (p0 << ((1) & 63))
        } else {
            var v19 int = (p0 | p0)
            _ = v19
        }
        _ = v17
    } else {
        if len(v7) > 0 {
            v7[0] = -1
        }
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 int64) bool {
    fmt.Println(p0)
    p0 += p0
    p0 -= 42
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 float32 = 0.0
    var v2 int8 = -1
    v3 := []string{("foo bar" + ((("hello" + "hello") + ("\n\t" + "\n\t")) + "hello"))}
    var v4 uint16 = 1
    var v5 uint32 = 493
    for fi20 := 0; fi20 < 6; fi20++ {
        fmt.Println(v5, v4, v1)
        v2 += (v2 * 95)
        var v21 int16 = genFn0(fi20, 9223372036854775807)
        _ = fi20
        _ = v21
    }
    v4 = (^(v4))
    v2 += v2
    if true {
        var v22 uint64 = 103
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := GenS3{F0: 2147483647, F1: 902, F2: (v4 / ((v4) | 1))}
        switch "" {
        case "\n\t":
            v2 += 127
        }
        for fi25 := 0; fi25 < 18; fi25++ {
            if true {
                var v26 uint32 = 42
                switch (fi25 % (((fi23 / ((fi25) | 1))) | 1)) {
                case 4294967295:
                    var v27 int = fi25
                    _ = v27
                }
                fmt.Println(v2, v1)
                v28 := GenS2{F0: uint8(v4), F1: 0}
                fmt.Println(fi23)
                _ = v26
                _ = v28
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    if true {
        v1 = (float32(3.14159) / ((float32((1e10 * (-2.25 + -2.25))) * ((v1 * 0.1) + 1.5)) - float32(((-2.25 - -2.25) + ((0.1 / 0.1) / 1.0)))))
        var v29 int64 = -1
        _ = v29
    } else {
        var v30 int64 = 0
        _ = v30
    }
    var v31 float64 = float64(((((v1 + (v1 / v1)) * (v1 + 0.1)) - ((v1 + (v1 * 0.1)) / ((v1 / v1) + v1))) + v1))
    v2 = (v2 * 59)
    v2 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
