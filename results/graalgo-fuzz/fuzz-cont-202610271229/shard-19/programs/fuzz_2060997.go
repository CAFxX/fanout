package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    var v1 bool = false
    var v2 uint = 18446744073709551615
    v3 := make(map[int]uint32)
    var v4 uint64 = 0
    var v5 int64 = 155
    v4 -= v4
    v5 = v5
    if (1 < int(v5)) {
        switch 0 {
        case -9223372036854775808:
            var v6 float32 = float32(-2.25)
            _ = v6
        default:
            v2 += (v2 % ((uint(v4)) | 1))
        }
    }
    var v7 float32 = 1.5
    v8 := []int{}
    v9 := GenS1{F0: -9223372036854775808, F1: (^(v4))}
    switch 4294967295 {
    case -1:
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS1{F0: 1, F1: 408}
            v12 := GenS2{F0: 255, F1: (("a" + "hello") + (("hello" + "foo bar") + ""))}
            if ((true && v1) && v1) {
                var v13 float32 = v7
                v1 = ((v7 * (((v7 / v13) + v13) * v7)) == (v13 * (v7 * v13)))
                fmt.Println(v1, v5)
                _ = v13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    case 32768:
        var v14 uint64 = v4
        _ = v14
    case 42:
        for fi15 := 0; fi15 < 18; fi15++ {
            v1 = (!(("foo bar" + ("foo bar" + "a")) < "\n\t"))
            if true {
                var v16 int = fi15
                switch 127 {
                case 2147483647:
                    v17 := GenS1{F0: 1, F1: (v4 * v4)}
                    _ = v17
                case -9223372036854775808:
                    v18 := GenS1{F0: v5, F1: 962}
                    _ = v18
                case 128:
                    if v1 {
                        v19 := GenS1{F0: v5, F1: (v4 / (((v4 ^ 894)) | 1))}
                        _ = v19
                    } else {
                        fmt.Println(v16)
                    }
                default:
                    for fi20 := 0; fi20 < 1; fi20++ {
                        var v21 int = (^(fi15))
                        _ = fi20
                        _ = v21
                    }
                }
                if ((v1 || (!(3.14159 > (v7 - float32(0.0))))) || true) {
                    var v22 int64 = 42
                    v1 = (!v1)
                    v22 += (-(v5))
                    _ = v22
                }
                _ = v16
            }
            _ = fi15
        }
    }
    v23 := v3[int(v4)]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
