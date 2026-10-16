package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    var v4 uint16 = 42
    v5 := GenS1{F0: v1, F1: 510, F2: (v4 / (((v4 + (v4 << ((692) & 63)))) | 1))}
    fmt.Println(v4, v3)
    if true {
        if (!(false || true)) {
            for fi6 := 0; fi6 < 19; fi6++ {
                switch "\n\t" {
                case "a":
                    v3 = v3
                case "\n\t":
                    v1 += v1
                default:
                    v7 := []int{}
                    _ = v7
                }
                v8 := GenS1{F0: v1, F1: 65535, F2: v4}
                fmt.Println(v3, v3, v1)
                _ = fi6
                _ = v8
            }
        }
        var v9 int8 = 42
        if (1 >= 65535) {
            if (-2.25 > 0.0) {
                v10 := GenS1{F0: (v1 >> ((18446744073709551615) & 63)), F1: 0, F2: v4}
                v3 = v3
                v11 := GenS2{F0: -610}
                _ = v10
                _ = v11
            } else {
                switch -2147483648 {
                case 32768:
                    _ = len(v2)
                case -9223372036854775808:
                    v12 := GenS2{F0: -9223372036854775808}
                    _ = v12
                case 128:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v1 = v1
                        _ = fi13
                    }
                default:
                    v9 = v9
                }
                if ((v1 / ((v1) | 1)) <= 850) {
                    v3 = v3
                    fmt.Println(v9, v9, v4)
                    v4 -= (v4 + (v4 % (((v4 % (((v4 % ((v4) | 1))) | 1))) | 1)))
                }
                v4 += 218
            }
        }
        _ = v9
    } else {
        if false {
            v1 = (^(v1))
        }
    }
    fmt.Println(v4, v4)
    v14 := []int{242}
    v2[v3] = true
    var v15 uint64 = 2
    var v16 bool = true
    v3 = (v3 + v3)
    v17 := GenS1{F0: (v1 >> ((65535) & 63)), F1: 1, F2: v4}
    var v18 float32 = ((1.5 + -2.25) / 3.14159)
    v19 := GenS1{F0: v1, F1: 65535, F2: (^(v4))}
    v18 = 1e10
    v1 = 432
    var v20 int16 = 32767
    var v21 uint16 = (v4 << ((18446744073709551615) & 63))
    v22 := GenS2{F0: -736}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
