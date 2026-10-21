package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int32 }
func genFn0(p0 string) int {
    v5 := []int{-720, 42, -9223372036854775808, 9223372036854775807}
    switch "a" {
    case "foo bar":
        var v6 int16 = -1
        _ = v6
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = 32767
        }
    default:
        fmt.Println(p0)
    }
    var v7 int64 = -1
    var v8 int64 = 173
    p0 = (p0 + ("hello" + "\n\t"))
    _ = p0
    _ = v5
    _ = v7
    _ = v8
    return 1
}
func genFn1() int32 {
    var v9 string = "foo bar"
    switch 631 {
    case 42:
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 int16 = 1
                v9 = "hello"
                var v13 uint32 = 639
                _ = fi11
                _ = v12
                _ = v13
            }
            if (183 == 127) {
                var v14 uint32 = 65535
                for fi15 := 0; fi15 < 17; fi15++ {
                    v16 := GenS1{F0: (("a" + v9) + ("" + v9))}
                    _ = fi15
                    _ = v16
                }
                v14 = (v14 / ((v14) | 1))
                _ = v14
            } else {
                v9 = "hello"
                var v17 uint64 = 18446744073709551615
                _ = v17
            }
            _ = fi10
        }
    case 32767:
        var v18 int8 = 127
        _ = v18
    default:
        switch "" {
        case "\n\t":
            switch -894 {
            case 127:
                var v19 string = v9
                _ = v19
            case 579:
                v9 = v9
            }
        default:
            if (1 >= 0) {
                fmt.Println(v9)
                v20 := []int{genFn0("a"), -1}
                _ = v20
            }
        }
    }
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -8
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    _ = len(v4)
    v2 = v2
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int64 = 9223372036854775807
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        if ((v2 << ((42) & 63)) != v2) {
            var v24 int = fi23
            _ = v24
        }
        v25 := GenS2{F0: ((1e10 * -2.25) <= -2.25), F1: genFn1()}
        v26 := GenS2{F0: false, F1: -855}
        _ = fi23
        _ = v25
        _ = v26
    }
    if (0 != -128) {
        v27 := v3[2147483647]
        _ = v27
    }
    v2 = (v2 - v2)
    var v28 float32 = (1.5 + 1.5)
    v28 += ((0.1 - v28) - 0.1)
    switch -448 {
    case 9223372036854775807:
        v28 = v28
    default:
        v28 = (((v28 + ((v28 / (v28 * v28)) + v28)) * (3.14159 * ((v28 + v28) * 0.1))) * v28)
    }
    switch 42 {
    case -174:
        var v29 uint8 = 128
        _ = v29
    default:
        fmt.Println(v28)
    }
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
