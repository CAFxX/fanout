package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 string; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int16 }
func genFn0(p0 bool, p1 uint16, p2 int16) string {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int = 1
        if (p0 || false) {
            v8 := []int{fi6, 32767, (v7 << ((120) & 63))}
            for fi9 := 0; fi9 < 17; fi9++ {
                switch ("a" + (("a" + ("" + "hello")) + ((("\n\t" + "\n\t") + (((("hello" + "") + ("" + ("foo bar" + ""))) + (("a" + "\n\t") + "foo bar")) + ("a" + ("foo bar" + "a")))) + ""))) {
                case "a":
                    var v10 uint = 19
                    _ = v10
                case "foo bar":
                    p0 = p0
                default:
                    if (-311 == 42) {
                        v11 := GenS1{F0: -71}
                        _ = v11
                    }
                }
                var v12 uint32 = 2
                var v13 float64 = (float64(1e10) - 0.0)
                fmt.Println(v12, p2)
                var v14 uint = 2
                if (1 > -2147483648) {
                    fmt.Println(p2, fi9, v12)
                    v15 := GenS2{F0: 1}
                    _ = v15
                } else {
                    v8 = append(v8, (v7 % ((0) | 1)))
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                fmt.Println(v7)
                p0 = p0
                _ = fi9
                _ = v12
                _ = v13
                _ = v14
            }
            p0 = true
            _ = v8
        } else {
            v7 -= (v7 / ((-859) | 1))
            v16 := GenS1{F0: 1}
            _ = v16
        }
        _ = fi6
        _ = v7
    }
    p1 = (p1 >> ((0) & 63))
    v17 := GenS2{F0: 42}
    var v18 bool = (42 < -2147483648)
    v19 := GenS3{F0: 9223372036854775807, F1: 252}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    return "a"
}

func main() {
    var v1 float64 = 0.1
    v2 := []string{"a", "hello"}
    v3 := []string{"hello", "\n\t"}
    v4 := make(map[uint32]string)
    var v5 int64 = 42
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 uint = 255
        var v22 int32 = -1
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS4{F0: -981, F1: 945}
            v1 += ((-2.25 - v1) - v1)
            _ = fi23
            _ = v24
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    v3 = append(v3, (("hello" + "hello") + "\n\t"))
    v25 := GenS5{F0: "foo bar", F1: 127, F2: -32768}
    v26 := GenS6{F0: 1, F1: ((2147483647 != 42) && (!false))}
    v27 := GenS3{F0: v5, F1: 0}
    v5 += (v5 >> ((534) & 63))
    _ = len(v3)
    switch "foo bar" {
    case "\n\t":
        var v28 int8 = 0
        _ = v28
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        switch "" {
        case "foo bar":
            for fi29 := 0; fi29 < 11; fi29++ {
                if len(v2) > 0 {
                    v2[0] = (("\n\t" + (("a" + "hello") + "\n\t")) + "")
                }
                _ = fi29
            }
        case "\n\t":
            _ = len(v3)
        default:
            for fi30 := 0; fi30 < 19; fi30++ {
                v1 = (v1 / float64((((1.5 - (1e10 - (0.1 - 0.1))) / ((float32(1e10) - ((1.5 - 1e10) / -2.25)) * float32((((v1 + v1) / v1) * (3.14159 + float64((3.14159 / (-2.25 + 1e10)))))))) - 1e10)))
                v31 := GenS7{F0: 42}
                fmt.Println(v1, fi30, v1)
                _ = fi30
                _ = v31
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
