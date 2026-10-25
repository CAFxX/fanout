package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 uint64; F2 bool }
func genFn0(p0 int64, p1 int, p2 uint8) uint8 {
    if true {
        p0 = p0
        p0 = -594
    }
    p0 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint8 = 19
    var v2 int8 = 127
    var v3 string = (("a" + ("a" + ((((("a" + "hello") + ("\n\t" + "hello")) + "foo bar") + "foo bar") + ("foo bar" + ((("a" + "a") + "") + ((("" + "\n\t") + ((("hello" + "foo bar") + ("a" + "hello")) + "\n\t")) + ("a" + ("a" + "a")))))))) + "foo bar")
    v4 := make(map[int]string)
    var v5 int32 = -2147483648
    if true {
        var v6 uint32 = 42
        v7 := []int{}
        if (18446744073709551615 >= 4294967295) {
            if true {
                v6 += v6
                if ((3.14159 + (((-2.25 * (1.5 - 0.0)) * ((0.1 * 3.14159) - (float64((1.5 + 1.5)) - 1.5))) * -2.25)) >= ((float64(-2.25) * ((1e10 + (float64(3.14159) * 0.0)) + 3.14159)) / 0.1)) {
                    switch 546 {
                    case 9223372036854775807:
                        v8 := GenS1{F0: int32(v2)}
                        _ = v8
                    case 128:
                        v2 = v2
                    default:
                        var v9 int8 = (v2 ^ int8(v6))
                        _ = v9
                    }
                    v10 := GenS1{F0: (v5 + (v5 & v5))}
                    v5 = 497
                    _ = v10
                }
            } else {
                v11 := GenS1{F0: -2147483648}
                v5 = v5
                _ = v11
            }
            if (3.14159 < 0.1) {
                switch ("foo bar" + ("hello" + "\n\t")) {
                case "\n\t":
                    var v12 uint = 977
                    _ = v12
                case "a":
                    delete(v4, -352)
                case "foo bar":
                    for fi13 := 0; fi13 < 5; fi13++ {
                        v4[128] = v3
                        _ = fi13
                    }
                }
                var v14 int16 = 0
                _ = v14
            } else {
                fmt.Println(v6)
                v6 -= 354
                v4[4294967295] = v3
            }
        } else {
            for fi15 := 0; fi15 < 6; fi15++ {
                fmt.Println(v3, v5)
                fmt.Println(v5)
                _ = fi15
            }
            fmt.Println(v2)
            v6 += v6
        }
        _ = v6
        _ = v7
    } else {
        var v16 float64 = (float64((0.1 / (1e10 / 1.0))) / 1.0)
        v17 := []int{127, 4294967295, 32767}
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := []int{(fi19 >> ((18446744073709551615) & 63)), -617}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    fmt.Println(v3, v5)
    var v21 int = -261
    for fi22 := 0; fi22 < 3; fi22++ {
        v5 += 1
        v4[-500] = "\n\t"
        v3 = (v3 + "")
        _ = fi22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        switch v21 {
        case 0:
            if true {
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(fi23, v2, v1)
                    _ = fi24
                }
            }
        case 1:
            v25 := GenS1{F0: (-(v5))}
            _ = v25
        case 189:
            v26 := []int{}
            _ = v26
        default:
            v27 := GenS2{F0: 2, F1: (^(v2))}
            _ = v27
        }
        var v28 int = (v21 << ((508) & 63))
        var v29 bool = true
        _ = fi23
        _ = v28
        _ = v29
    }
    if true {
        var v30 float64 = -2.25
        _ = v30
    } else {
        if (32767 < 42) {
            v31 := GenS3{F0: "a", F1: uint64(v1), F2: ((v1 + (v1 - v1)) >= (v1 | genFn0(-605, (-(v21)), (v1 % ((v1) | 1)))))}
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 uint16 = 1
        v34 := GenS3{F0: "foo bar", F1: 42, F2: true}
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS2{F0: 0, F1: v2}
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
