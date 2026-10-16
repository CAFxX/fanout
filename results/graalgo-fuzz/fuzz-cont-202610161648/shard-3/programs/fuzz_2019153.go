package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16) uint8 {
    switch "hello" {
    case "a":
        v6 := GenS1{F0: ("a" + "\n\t")}
        _ = v6
    case "":
        v7 := GenS1{F0: "a"}
        _ = v7
    default:
        v8 := GenS2{F0: 2147483647}
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 float64 = ((0.0 + 1.5) - (float64((0.1 / (1e10 - ((3.14159 / 1e10) / ((-2.25 / 1e10) * 1.5))))) / float64(-2.25)))
        var v11 int16 = 42
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    return 9
}

func main() {
    var v1 int16 = 32767
    var v2 int = 919
    var v3 uint32 = 1
    var v4 int32 = 540
    var v5 uint32 = 65535
    v3 += v5
    fmt.Println(v1, v3, v5)
    if (9223372036854775807 <= 9223372036854775807) {
        v5 += v5
        v12 := []int{-9223372036854775808}
        _ = v12
    }
    var v13 int = (v2 / (((v2 % ((v2) | 1))) | 1))
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 float32 = float32(((((3.14159 / 3.14159) + 0.0) * 3.14159) / -2.25))
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: ""}
    var v17 uint = 29
    v1 = v1
    if (2 == genFn0(0)) {
        v18 := []int{v2, 9223372036854775807, 4294967295, (v13 >> ((0) & 63))}
        if false {
            var v19 bool = (!false)
            _ = v19
        }
        _ = v18
    }
    fmt.Println(v17)
    var v20 string = ""
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 string = v20
        var v23 int16 = 601
        if ((v1 >> (((^(v17))) & 63)) == (v23 * (^(v1)))) {
            var v24 uint64 = 4294967295
            _ = v24
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    v2 += -9223372036854775808
    for fi25 := 0; fi25 < 3; fi25++ {
        if false {
            if false {
                if ((v1 / ((870) | 1)) <= (v1 / ((v1) | 1))) {
                    if (((1e10 <= 1.5) || (v1 != (-(v1)))) || true) {
                        for fi26 := 0; fi26 < 15; fi26++ {
                            var v27 int = fi26
                            switch (v13 + v27) {
                            case 2147483647:
                                var v28 int = (fi25 | int(v17))
                                fmt.Println(fi25)
                                _ = v28
                            default:
                                fmt.Println(v4, v2)
                            }
                            v1 += v1
                            _ = fi26
                            _ = v27
                        }
                        for fi29 := 0; fi29 < 16; fi29++ {
                            var v30 float32 = (0.0 / 1.5)
                            v13 = (-(fi25))
                            _ = fi29
                            _ = v30
                        }
                    } else {
                        v31 := []int{v2, (v2 * 2147483647), v13}
                        if (((v17 << ((v17) & 63)) != 18446744073709551615) || ((false || true) && false)) {
                            v32 := GenS1{F0: v20}
                            if len(v31) > 0 {
                                v31[0] = (v13 * fi25)
                            }
                            for fi33 := 0; fi33 < 17; fi33++ {
                                var v34 bool = false
                                if v34 {
                                    var v35 float32 = float32(1.5)
                                    if ((!v34) || (0.0 != (1e10 - 0.1))) {
                                        if ((502 > uint32(v17)) || (575 < uint64(v17))) {
                                            v36 := GenS2{F0: 0}
                                            _ = v36
                                        }
                                        var v37 float32 = v35
                                        if ((v17 != (^(v17))) || v34) {
                                            var v38 int = v2
                                            _ = v38
                                        } else {
                                            v31 = append(v31, fi33)
                                            var v39 uint8 = 128
                                            _ = v39
                                        }
                                        _ = v37
                                    }
                                    _ = v35
                                }
                                _ = fi33
                                _ = v34
                            }
                            _ = v32
                        }
                        _ = v31
                    }
                    v2 = fi25
                    v2 -= fi25
                }
                fmt.Println(fi25)
            }
            v4 -= -2147483648
        }
        fmt.Println(fi25)
        for fi40 := 0; fi40 < 8; fi40++ {
            fmt.Println(fi25)
            _ = fi40
        }
        _ = fi25
    }
    v41 := []int{9223372036854775807, -9223372036854775808, (v2 | v2), v2}
    var v42 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
