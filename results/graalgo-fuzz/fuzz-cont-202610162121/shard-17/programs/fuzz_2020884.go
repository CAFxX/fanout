package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int; F1 string; F2 string }
func genFn0(p0 int8, p1 int64) string {
    for fi5 := 0; fi5 < 13; fi5++ {
        if true {
            p1 -= (-(p1))
        } else {
            switch ("foo bar" + ((("a" + (("a" + "") + ("a" + ("foo bar" + (("\n\t" + "") + "a"))))) + "a") + ("" + ""))) {
            case "\n\t":
                fmt.Println(p0)
            default:
                fmt.Println(fi5, p0)
            }
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 int32 = 0
                fmt.Println(p0, p0)
                v8 := GenS1{F0: -32768, F1: -205, F2: 18446744073709551615}
                _ = fi6
                _ = v7
                _ = v8
            }
        }
        _ = fi5
    }
    var v9 uint32 = 778
    v9 = v9
    _ = p0
    _ = p1
    _ = v9
    return "foo bar"
}

func main() {
    var v1 float64 = ((((0.0 - 0.1) - (float64(1.5) * 1.5)) + 0.1) / 1.0)
    v2 := []int{0, 4294967295, 856, 9223372036854775807}
    v3 := []uint32{1}
    var v4 int = 817
    v4 = (v4 + -2147483648)
    var v10 bool = false
    switch (v4 << ((65535) & 63)) {
    case 32767:
        for fi11 := 0; fi11 < 7; fi11++ {
            _ = len(v3)
            v4 += 128
            var v12 int16 = 32767
            _ = fi11
            _ = v12
        }
    case 32768:
        v10 = (!(!v10))
    default:
        v13 := GenS1{F0: 1, F1: -1, F2: 181}
        _ = v13
    }
    var v14 int32 = -2147483648
    if (v10 || (2 < 255)) {
        var v15 int16 = -1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 float64 = (v1 / 1.5)
            _ = len(v2)
            v18 := GenS2{F0: 551, F1: -218, F2: 100}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    v19 := GenS3{F0: (v10 || v10), F1: 42}
    for fi20 := 0; fi20 < 2; fi20++ {
        v10 = true
        _ = fi20
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v21 := GenS4{F0: (v4 ^ (v4 << ((255) & 63))), F1: "hello", F2: "\n\t"}
    v3 = append(v3, 590)
    v10 = false
    _ = len(v3)
    v1 = (v1 + v1)
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 bool = v10
        _ = len(v3)
        v24 := GenS3{F0: (1 >= 10), F1: 42}
        _ = fi22
        _ = v23
        _ = v24
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
