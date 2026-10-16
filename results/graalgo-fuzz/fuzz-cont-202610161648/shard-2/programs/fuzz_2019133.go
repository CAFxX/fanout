package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 string; F1 string; F2 int64 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 uint8, p1 string) bool {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int32 = -592
        if (65535 >= 2) {
            var v7 int32 = (v6 << ((18446744073709551615) & 63))
            v8 := GenS1{F0: -9223372036854775808, F1: -102}
            _ = v7
            _ = v8
        } else {
            if (!("\n\t" <= p1)) {
                if true {
                    switch (fi5 % (((fi5 * (fi5 ^ 42))) | 1)) {
                    case 42:
                        fmt.Println(v6, fi5, p0)
                    default:
                        v9 := GenS2{F0: p1, F1: ((("hello" + (p1 + p1)) + p1) + p1), F2: 1}
                        _ = v9
                    }
                }
                p0 = (p0 - (p0 | (p0 ^ 255)))
            }
        }
        _ = fi5
        _ = v6
    }
    p0 = (p0 % ((p0) | 1))
    var v10 int32 = 1
    _ = p0
    _ = p1
    _ = v10
    return (p1 <= (p1 + ""))
}

func main() {
    var v1 uint8 = 2
    v2 := []int{-2147483648, 0, 524, 2147483647}
    var v3 int32 = 42
    var v4 bool = (-32768 == int16(v3))
    v4 = (!v4)
    fmt.Println(v4, v4)
    if (!(v4 && (43 < 0))) {
        v3 -= (v3 / ((181) | 1))
    }
    var v11 float32 = (((1e10 * (1e10 - ((float32(float64((3.14159 - 0.1))) + float32(3.14159)) / (1e10 / 3.14159)))) + float32(0.0)) / 1.0)
    v2 = append(v2, -1)
    _ = len(v2)
    fmt.Println(v1, v3, v1)
    if (2 <= 510) {
        fmt.Println(v4, v3)
    } else {
        if len(v2) > 0 {
            v2[0] = -869
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        if (9223372036854775807 > 1) {
            var v13 bool = (true || v4)
            if (379 < 1) {
                if len(v2) > 0 {
                    v2[0] = (fi12 >> ((255) & 63))
                }
                fmt.Println(v1, v1, v3)
            } else {
                v14 := GenS2{F0: ("hello" + ""), F1: "foo bar", F2: 0}
                _ = v14
            }
            _ = v13
        } else {
            if (0.0 > 1.5) {
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := GenS3{F0: (v3 + (v3 % ((2147483647) | 1))), F1: v4}
                    _ = fi15
                    _ = v16
                }
            } else {
                v11 = ((v11 * v11) / 3.14159)
                var v17 float64 = 1e10
                var v18 float64 = float64((v11 + (v11 / v11)))
                _ = v17
                _ = v18
            }
        }
        _ = fi12
    }
    fmt.Println(v3, v1)
    v3 = v3
    switch 189 {
    case 128:
        fmt.Println(v3, v1, v1)
    case -9223372036854775808:
        fmt.Println(v3, v4)
    case 127:
        v3 += 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
