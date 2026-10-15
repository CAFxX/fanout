package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int64 }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int16; F2 int }
type GenS6 struct { F0 int64 }
func genFn0(p0 bool, p1 uint64) uint {
    p1 = (p1 ^ (p1 % ((p1) | 1)))
    if (p0 && (!true)) {
        fmt.Println(p1, p1)
    } else {
        fmt.Println(p0, p0)
        p1 += p1
    }
    var v6 uint32 = 696
    v7 := GenS1{F0: (v6 * 120), F1: (p0 && ((0 != 18446744073709551615) || (65535 > 964))), F2: -34}
    v8 := GenS2{F0: 1, F1: -524, F2: 867}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := []int{}
    v2 := []string{"foo bar", "hello", ((("foo bar" + "") + (("foo bar" + (("a" + "a") + "a")) + "")) + "")}
    var v3 bool = false
    var v4 uint = 42
    v5 := make(map[string]int)
    v9 := GenS3{F0: 867}
    delete(v5, ("" + ""))
    if (v3 && false) {
        v4 = (v4 + 1)
        fmt.Println(v4, v4)
        v4 -= 2
    }
    v10 := GenS4{F0: "foo bar"}
    var v11 int = -536
    var v12 bool = (9223372036854775807 <= 9223372036854775807)
    v11 = -597
    var v13 uint16 = 1
    var v14 uint8 = 128
    if v3 {
        if (-195 > -9223372036854775808) {
            if (v4 < (v4 << ((1) & 63))) {
                v15 := GenS5{F0: "a", F1: 32767, F2: v11}
                switch -212 {
                case 1:
                    v3 = (((-2.25 + 1.5) >= (1e10 * 3.14159)) || ((v14 / (((^(v14))) | 1)) == 128))
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = v15
            } else {
                v3 = v12
                fmt.Println(v12)
            }
            v2 = append(v2, "hello")
        } else {
            var v16 int64 = -9223372036854775808
            _ = v16
        }
        var v17 float32 = float32((0.0 * 1e10))
        _ = v17
    }
    var v18 string = "\n\t"
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 int16 = 767
        v21 := GenS6{F0: 9223372036854775807}
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 bool = (((v4 >> (((v4 + v4)) & 63)) < (^(v4))) || (-1 == int32(v11)))
    _ = len(v5)
    var v23 uint64 = 523
    delete(v5, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
