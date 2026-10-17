package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32, p1 int) uint {
    v5 := []int{(p1 % (((p1 | p1)) | 1)), (p1 & p1)}
    var v6 int64 = 0
    var v7 bool = ((p1 & (-(p1))) >= p1)
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(p0, p1, p1)
        if true {
            var v9 string = "hello"
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: 1000, F1: 1, F2: p1}
                v12 := GenS1{F0: -1, F1: 9223372036854775807, F2: p1}
                if len(v5) > 0 {
                    v5[0] = fi10
                }
                if (true && v7) {
                    var v13 uint = 63
                    _ = v13
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        } else {
            fmt.Println(v7, v6)
            switch fi8 {
            case 127:
                v6 = -737
            default:
                v6 = v6
            }
        }
        if false {
            if v7 {
                v6 -= v6
                if (!v7) {
                    var v14 uint64 = 810
                    _ = v14
                } else {
                    var v15 uint8 = 127
                    _ = v15
                }
                fmt.Println(p1, fi8, p0)
            }
            if len(v5) > 0 {
                v5[0] = (fi8 + 4294967295)
            }
            v16 := GenS1{F0: 42, F1: v6, F2: (-(fi8))}
            _ = v16
        } else {
            fmt.Println(p1, p0)
            v6 -= 9223372036854775807
            fmt.Println(v7)
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int32) int16 {
    if false {
        p0 -= (p0 * 315)
    }
    fmt.Println(p0)
    _ = p0
    return 32767
}

func main() {
    var v1 uint64 = 0
    var v2 bool = (("a" + "a") <= "\n\t")
    var v3 int64 = -971
    v4 := make(map[int]bool)
    var v17 bool = (v3 < (v3 * v3))
    fmt.Println(v2)
    for fi18 := 0; fi18 < 9; fi18++ {
        v3 = (v3 - (-(v3)))
        v19 := v4[(fi18 - (^(fi18)))]
        _ = fi18
        _ = v19
    }
    fmt.Println(v2, v2)
    switch 20 {
    case -1:
        for fi20 := 0; fi20 < 4; fi20++ {
            switch "a" {
            case "foo bar":
                delete(v4, (fi20 & 797))
            case "":
                _ = len(v4)
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                v1 -= 570
                var v22 uint8 = 127
                v23 := []int{fi21}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
    case -826:
        for fi24 := 0; fi24 < 1; fi24++ {
            v3 = -9223372036854775808
            for fi25 := 0; fi25 < 6; fi25++ {
                v17 = (!true)
                _ = fi25
            }
            _ = fi24
        }
    }
    var v26 int8 = 100
    var v27 float32 = float32(0.1)
    var v28 string = (("a" + ("foo bar" + ("" + ("a" + ("foo bar" + ("\n\t" + "a")))))) + "hello")
    if v2 {
        if v2 {
            if ((v3 & v3) <= v3) {
                v29 := []int{-476, -820, 2147483647}
                v26 = -1
                for fi30 := 0; fi30 < 12; fi30++ {
                    if (2 == 42) {
                        var v31 uint = 0
                        _ = v31
                    } else {
                        if (!(!(2 >= 42))) {
                            v4[(fi30 | (fi30 % (((fi30 << ((uint(v1)) & 63))) | 1)))] = ((127 != 0) || false)
                            switch fi30 {
                            case 554:
                                var v32 float64 = 1e10
                                _ = v32
                            case -2147483648:
                                if len(v29) > 0 {
                                    v29[0] = (^(fi30))
                                }
                            }
                            var v33 uint8 = 0
                            _ = v33
                        } else {
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                        }
                        delete(v4, (fi30 | -2147483648))
                    }
                    _ = fi30
                }
                _ = v29
            }
        } else {
            v34 := v4[977]
            _ = v34
        }
    }
    v28 = "\n\t"
    for fi35 := 0; fi35 < 13; fi35++ {
        fmt.Println(v26)
        for fi36 := 0; fi36 < 9; fi36++ {
            var v37 uint64 = 4294967295
            fmt.Println(v3)
            var v38 int32 = -1
            fmt.Println(v37)
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    v27 = (((v27 / (v27 + v27)) / float32((-2.25 + (-2.25 + ((0.0 * 0.1) + 0.0))))) * (float32((1e10 + (1e10 - 1e10))) * -2.25))
    switch (v28 + v28) {
    case "foo bar":
        for fi39 := 0; fi39 < 1; fi39++ {
            v40 := GenS1{F0: genFn1(-344), F1: v3, F2: fi39}
            _ = fi39
            _ = v40
        }
    case "":
        fmt.Println(v26, v28)
    }
    v41 := GenS2{F0: 4294967295}
    var v42 int16 = 491
    v3 = (v3 * (v3 | (v3 % ((v3) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
