package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 int, p1 uint32, p2 int64) uint {
    p1 = 386
    p0 += p0
    switch p0 {
    case 217:
        v5 := []int{(p0 >> ((4294967295) & 63)), (p0 >> ((915) & 63)), p0}
        _ = v5
    default:
        p2 -= p2
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint = 321
        p1 -= p1
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}
func genFn1() uint32 {
    var v8 uint16 = 2
    switch ("hello" + "foo bar") {
    case "hello":
        v9 := []int{32768, -9223372036854775808}
        _ = v9
    }
    v10 := []int{903, 42}
    v11 := GenS1{F0: 32767, F1: 255, F2: 0}
    _ = v10
    _ = v11
    _ = v8
    return 2147483648
}

func main() {
    var v1 int64 = 1
    var v2 int8 = 1
    var v3 float32 = float32(1.5)
    var v4 int16 = 1000
    v12 := GenS1{F0: 2147483647, F1: 4294967295, F2: v1}
    if false {
        if true {
            var v13 uint16 = 65535
            v3 += 1e10
            for fi14 := 0; fi14 < 1; fi14++ {
                v1 += (v1 + (v1 % (((v1 % ((1) | 1))) | 1)))
                _ = fi14
            }
            _ = v13
        }
        v15 := []int{}
        if (false && (!(false && ((true && false) && false)))) {
            v4 = -115
            if true {
                for fi16 := 0; fi16 < 12; fi16++ {
                    if len(v15) > 0 {
                        v15[0] = (fi16 / (((fi16 & (fi16 * fi16))) | 1))
                    }
                    fmt.Println(v3, v1, v1)
                    _ = fi16
                }
                v4 = v4
            }
            var v17 int64 = (v1 + (v1 % ((v1) | 1)))
            _ = v17
        }
        _ = v15
    } else {
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := []int{32767, (fi18 & (fi18 | (fi18 + (fi18 ^ fi18)))), (fi18 * (fi18 - (^(fi18))))}
            fmt.Println(v1, fi18, v2)
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 uint8 = 2
        if (!false) {
            if false {
                v3 = 0.1
                if false {
                    var v22 int16 = v4
                    var v23 int16 = v4
                    if (v1 < v1) {
                        fmt.Println(v4, v1)
                        v24 := []int{fi20, (fi20 ^ (fi20 & fi20)), (fi20 - (fi20 - fi20)), 412}
                        for fi25 := 0; fi25 < 6; fi25++ {
                            v1 -= (-(v1))
                            _ = fi25
                        }
                        _ = v24
                    } else {
                        v21 += v21
                    }
                    _ = v22
                    _ = v23
                }
            }
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v2, v3, v1)
    if ((float64(v3) * 0.1) < (((-2.25 * 0.0) * (0.1 + -2.25)) / 1e10)) {
        v4 = (v4 % ((1000) | 1))
        var v26 int8 = (^(v2))
        _ = v26
    } else {
        v2 = (v2 % ((-128) | 1))
        fmt.Println(v4, v4, v2)
    }
    v27 := GenS2{F0: -688, F1: v4}
    switch "a" {
    case "foo bar":
        v28 := []int{1, int(v2)}
        _ = v28
    default:
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := []int{fi29, -1, fi29, 0}
            _ = fi29
            _ = v30
        }
    }
    var v31 int16 = v4
    switch 32768 {
    case 2147483647:
        fmt.Println(v3, v1)
    default:
        if true {
            var v32 float64 = -2.25
            _ = v32
        } else {
            var v33 bool = (636 != -108)
            _ = v33
        }
    }
    var v34 uint = 255
    for fi35 := 0; fi35 < 18; fi35++ {
        fmt.Println(v2)
        var v36 int64 = v1
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v34
    _ = v4
}
