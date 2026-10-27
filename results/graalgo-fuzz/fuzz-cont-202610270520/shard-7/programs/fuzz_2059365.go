package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint64; F1 int32; F2 int }
func genFn0(p0 string, p1 uint) uint {
    p0 = p0
    p1 = p1
    var v6 bool = true
    _ = p0
    _ = p1
    _ = v6
    return (^(p1))
}
func genFn1(p0 int8, p1 bool, p2 uint16) string {
    if (370 <= uint32(p2)) {
        var v7 int64 = -714
        _ = v7
    }
    p1 = p1
    if p1 {
        v8 := []int{2147483647}
        var v9 float32 = 1e10
        v10 := GenS1{F0: 42}
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    v1 := []string{"\n\t", ("\n\t" + "hello"), "\n\t", "hello"}
    v2 := make(map[string]uint32)
    var v3 float64 = float64(0.0)
    var v4 int32 = -1
    v5 := []string{"a", ("\n\t" + ("a" + ("\n\t" + "hello"))), "foo bar", ("\n\t" + "")}
    switch 32767 {
    case -2147483648:
        var v11 bool = (2 >= 65535)
        _ = v11
    default:
        v4 = v4
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        v2[("foo bar" + "\n\t")] = 65535
        v13 := GenS2{F0: (^(v4)), F1: 1}
        var v14 uint16 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint32 = 65535
    for fi16 := 0; fi16 < 7; fi16++ {
        v4 += (v4 % ((v4) | 1))
        var v17 uint = 255
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v15, v4)
    v4 = v4
    v3 += v3
    for fi18 := 0; fi18 < 13; fi18++ {
        v3 += (v3 + ((v3 * float64((1e10 / 3.14159))) + ((float64(0.1) / (float64(1.5) * -2.25)) + v3)))
        v2["hello"] = v15
        _ = fi18
    }
    v19 := GenS2{F0: (v4 * (-(v4))), F1: (v15 >> ((0) & 63))}
    v20 := GenS2{F0: 989, F1: v15}
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 uint8 = 0
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v1 = append(v1, ("hello" + "\n\t"))
        var v24 int8 = -128
        var v25 float32 = (float32(0.1) / 1.0)
        _ = fi23
        _ = v24
        _ = v25
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 int16 = 32767
        _ = fi26
        _ = v27
    }
    switch "a" {
    case "foo bar":
        if true {
            if (18446744073709551615 <= 899) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        } else {
            v28 := GenS3{F0: (v4 % ((v4) | 1)), F1: 0, F2: -32768}
            v29 := GenS4{F0: 395, F1: v4, F2: 42}
            _ = v28
            _ = v29
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
