package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint }
func genFn0(p0 int64) int64 {
    if (true && true) {
        p0 += (p0 | 0)
    }
    p0 = p0
    p0 += p0
    _ = p0
    return (p0 | (p0 << ((1) & 63)))
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]uint32)
    var v3 uint32 = 42
    v4 := []int{-405, 9223372036854775807}
    v5 := []string{}
    v1 -= (v1 + (v1 ^ (v1 | uint64(v3))))
    v6 := v2[-338]
    v7 := GenS1{F0: (^(v1)), F1: ("" + ("a" + (("" + ("\n\t" + "\n\t")) + ("foo bar" + (("hello" + "foo bar") + "\n\t"))))), F2: 118}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int32 = 0
        if (42 == 42) {
            _ = len(v4)
            v3 = v3
            for fi10 := 0; fi10 < 5; fi10++ {
                switch (-(fi8)) {
                case -37:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case 9223372036854775807:
                    var v11 int = fi10
                    _ = v11
                case 127:
                    if true {
                        v12 := GenS2{F0: 0, F1: 3, F2: 42}
                        v6 -= 2147483647
                        _ = v12
                    } else {
                        var v13 int8 = int8(fi10)
                        var v14 int64 = 9223372036854775807
                        fmt.Println(fi8)
                        _ = v13
                        _ = v14
                    }
                default:
                    v9 -= v9
                }
                if (!(true && (-298 <= 42))) {
                    _ = len(v2)
                    fmt.Println(fi10)
                    v15 := GenS1{F0: v1, F1: ("a" + ("a" + "a")), F2: 128}
                    var v16 float32 = 0.1
                    _ = v15
                    _ = v16
                } else {
                    var v17 int = fi10
                    v9 -= (v9 | v9)
                    v6 += (v6 - (v3 ^ v3))
                    switch "" {
                    case "a":
                        fmt.Println(v17, v1, v9)
                    }
                    v9 = v9
                    _ = v17
                }
                _ = fi10
            }
        } else {
            var v18 uint = 65535
            v2[42] = uint32(v18)
            v9 = 1
            _ = v18
        }
        _ = fi8
        _ = v9
    }
    v3 = 2147483648
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS3{F0: 18446744073709551615}
        v1 = (v1 * (v1 / (((v1 % ((v1) | 1))) | 1)))
        _ = fi19
        _ = v20
    }
    if ((v6 * 948) <= (v3 % ((v3) | 1))) {
        var v21 uint8 = 1
        _ = v21
    }
    var v22 float64 = 0.1
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
