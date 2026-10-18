package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int32; F2 bool }
func genFn0(p0 string, p1 uint32, p2 string) uint8 {
    v5 := GenS1{F0: ((p0 + (p0 + p2)) + p0)}
    var v6 string = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 147
}
func genFn1(p0 bool, p1 uint8, p2 int8) int32 {
    if (0 == 4294967295) {
        if true {
            v7 := []int{-2147483648, -860, 128, 9223372036854775807}
            _ = v7
        }
    }
    fmt.Println(p1, p2, p0)
    for fi8 := 0; fi8 < 9; fi8++ {
        p2 = (p2 * -1)
        p1 -= (p1 * 101)
        var v9 float64 = (float64(3.14159) + (0.0 + float64(0.1)))
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        if (2147483648 == 0) {
            fmt.Println(fi10)
            var v11 uint = 881
            _ = v11
        } else {
            for fi12 := 0; fi12 < 14; fi12++ {
                p1 = genFn0(("foo bar" + ("foo bar" + "hello")), 2, (((("hello" + "foo bar") + ("\n\t" + "")) + ("a" + (("" + "a") + ("foo bar" + ("hello" + "\n\t"))))) + (("" + "a") + "foo bar")))
                var v13 int = (fi10 ^ fi12)
                v14 := []int{(v13 / (((^(v13))) | 1)), -9223372036854775808}
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        p0 = (!p0)
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := []int{fi10, (^(fi10))}
            _ = fi15
            _ = v16
        }
        _ = fi10
    }
    v17 := GenS1{F0: ("a" + ("hello" + "foo bar"))}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 1
}

func main() {
    var v1 float32 = 1.5
    v2 := []int{32768, 4294967295}
    var v3 int8 = 127
    var v4 int8 = 42
    for fi18 := 0; fi18 < 10; fi18++ {
        switch 0 {
        case -749:
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 float64 = (1e10 * ((-2.25 + ((float64((1e10 - (v1 - 1e10))) / 1e10) / -2.25)) / ((1.5 + 1e10) - 1e10)))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    v21 := GenS1{F0: "foo bar"}
                    v22 := GenS2{F0: 962}
                    _ = v21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
        case 32767:
            v4 += v4
        }
        _ = fi18
    }
    switch "foo bar" {
    case "foo bar":
        var v23 uint16 = 255
        _ = v23
    default:
        v4 = (-(v3))
    }
    fmt.Println(v4)
    var v24 int16 = 42
    _ = len(v2)
    if false {
        v4 -= (v3 % (((v3 % ((100) | 1))) | 1))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 = (((-2.25 + (0.1 * 0.1)) + 0.1) * float32(0.0))
    } else {
        fmt.Println(v1)
        var v25 int = -40
        _ = v25
    }
    switch -2147483648 {
    case 9223372036854775807:
        v3 -= 127
    case -2147483648:
        v4 += (v4 | (v4 >> ((4294967295) & 63)))
    case -9223372036854775808:
        if false {
            v1 += 1e10
            if false {
                v26 := GenS2{F0: 2}
                _ = v26
            }
        } else {
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
            fmt.Println(v1, v1, v1)
            var v27 bool = (4294967295 != 0)
            _ = v27
        }
    default:
        v28 := GenS3{F0: 32767}
        _ = v28
    }
    v24 = (^(v24))
    if (true && (!false)) {
        switch 418 {
        case 127:
            for fi29 := 0; fi29 < 15; fi29++ {
                var v30 float64 = (-2.25 + 0.0)
                var v31 uint64 = 2
                v32 := GenS4{F0: ("\n\t" + "\n\t"), F1: 42, F2: (true && true)}
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        case -9223372036854775808:
            v1 = 1e10
        case 888:
            var v33 int16 = v24
            _ = v33
        default:
            v3 -= (v3 ^ 24)
        }
        fmt.Println(v3, v24)
    } else {
        v2 = append(v2, 127)
        v1 += ((1e10 * 3.14159) + v1)
        for fi34 := 0; fi34 < 20; fi34++ {
            v1 -= float32((0.0 + -2.25))
            v1 = (0.0 / ((v1 - -2.25) * v1))
            fmt.Println(v24, fi34)
            _ = fi34
        }
    }
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
