package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int16, p1 uint) int64 {
    var v6 int64 = 42
    v7 := GenS1{F0: -1, F1: "", F2: 251}
    v8 := GenS1{F0: p0, F1: "a", F2: -1}
    switch -2147483648 {
    case 32768:
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := []int{9223372036854775807, (-(fi10)), (fi9 << (((p1 ^ 18446744073709551615)) & 63))}
                _ = fi10
                _ = v11
            }
            switch (fi9 & fi9) {
            case -1:
                p0 = 41
            case -815:
                var v12 int8 = -1
                _ = v12
            case 128:
                var v13 uint8 = 0
                _ = v13
            }
            _ = fi9
        }
    case 42:
        v6 += (v6 * v6)
    case 697:
        v14 := GenS2{F0: 0, F1: 511}
        _ = v14
    }
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (-(v6))
}

func main() {
    v1 := []string{((("a" + "foo bar") + ("foo bar" + "foo bar")) + (("a" + "a") + "\n\t"))}
    v2 := []string{"a", "", "foo bar"}
    var v3 int = 42
    var v4 float64 = (3.14159 * (float64(3.14159) + -2.25))
    var v5 int = 32767
    v5 += -392
    var v15 int = (v3 | -2147483648)
    fmt.Println(v15, v5, v4)
    v5 += (v15 * (v15 / ((v5) | 1)))
    var v16 float32 = 0.1
    v15 += (^(v3))
    var v17 uint32 = 2147483648
    v1 = append(v1, "foo bar")
    switch "foo bar" {
    case "a":
        if true {
            v4 = (v4 + v4)
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if false {
        var v18 int64 = 9223372036854775807
        v19 := GenS2{F0: 32767, F1: v17}
        v20 := GenS2{F0: 322, F1: v17}
        _ = v18
        _ = v19
        _ = v20
    }
    v3 += (v3 - 128)
    fmt.Println(v5, v16, v17)
    var v21 bool = ((false || (!false)) || (254 < 127))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
