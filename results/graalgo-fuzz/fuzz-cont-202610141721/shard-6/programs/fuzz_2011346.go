package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint; F1 uint8; F2 int16 }
func genFn0() uint64 {
    v5 := []int{}
    v6 := GenS1{F0: -9223372036854775808, F1: 84}
    _ = v5
    _ = v6
    return 672
}
func genFn1(p0 int8, p1 int8, p2 uint16) int32 {
    v7 := []int{}
    var v8 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = -2.25
    v3 := make(map[int]string)
    var v4 int8 = 127
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: v4, F1: fi9}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: (-(v4)), F1: -2147483648}
    for fi12 := 0; fi12 < 20; fi12++ {
        if true {
            v13 := []int{}
            _ = v13
        }
        _ = fi12
    }
    var v14 uint8 = 127
    v15 := []int{}
    v4 += 100
    v4 = v4
    if (0 != 2) {
        if false {
            if true {
                v16 := GenS3{F0: ("a" + ("" + ("foo bar" + "a"))), F1: 431}
                _ = v16
            } else {
                v1[128] = 1
                v14 -= (^(v14))
                if (false && false) {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v14 = (v14 % ((v14) | 1))
                        v4 = v4
                        _ = fi17
                    }
                    v4 = (v4 >> ((869) & 63))
                } else {
                    var v18 uint32 = 2147483647
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                for fi20 := 0; fi20 < 5; fi20++ {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = fi20
                }
                v21 := v1[fi19]
                v22 := GenS4{F0: 2, F1: v14, F2: 1}
                _ = fi19
                _ = v21
                _ = v22
            }
        } else {
            var v23 int16 = 43
            _ = v23
        }
        for fi24 := 0; fi24 < 12; fi24++ {
            v1[fi24] = 2147483648
            for fi25 := 0; fi25 < 1; fi25++ {
                _ = len(v3)
                _ = fi25
            }
            _ = fi24
        }
        if (4294967295 == 187) {
            _ = len(v15)
            switch ("foo bar" + "\n\t") {
            case "a":
                v2 -= (1e10 + v2)
            case "foo bar":
                var v26 bool = (false || false)
                _ = v26
            case "\n\t":
                if len(v15) > 0 {
                    v15[0] = -2147483648
                }
            default:
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            }
            for fi27 := 0; fi27 < 5; fi27++ {
                v14 = (v14 + v14)
                if true {
                    fmt.Println(v4)
                    v15 = append(v15, fi27)
                    var v28 uint64 = 86
                    _ = v28
                }
                _ = fi27
            }
        }
    }
    switch ("foo bar" + "hello") {
    case "\n\t":
        v2 -= (v2 - (float32(float64(((((1.5 + 3.14159) / v2) * v2) / float32(float64((v2 + v2)))))) * v2))
    case "a":
        var v29 int8 = (v4 | (v4 >> ((2) & 63)))
        _ = v29
    default:
        v30 := v1[32767]
        _ = v30
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        v32 := GenS3{F0: ("a" + "a"), F1: -2147483648}
        fmt.Println(fi31)
        _ = fi31
        _ = v32
    }
    v15 = append(v15, 32767)
    for fi33 := 0; fi33 < 17; fi33++ {
        fmt.Println(v4, v2, v2)
        _ = fi33
    }
    var v34 int32 = 2147483647
    v15 = append(v15, int(v4))
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 bool = false
        delete(v1, fi35)
        _ = fi35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
