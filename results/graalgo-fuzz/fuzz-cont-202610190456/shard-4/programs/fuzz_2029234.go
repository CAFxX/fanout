package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 int; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint16 }
func genFn0() uint64 {
    if true {
        var v7 uint32 = 2147483647
        v7 = (v7 << ((83) & 63))
        _ = v7
    }
    var v8 int64 = -1
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: false, F1: fi9, F2: "foo bar"}
        v8 = (v8 ^ (v8 | -953))
        v8 = (v8 + (v8 % ((1) | 1)))
        _ = fi9
        _ = v10
    }
    _ = v8
    return 42
}

func main() {
    var v1 uint32 = 718
    v2 := make(map[int]uint32)
    var v3 string = (("hello" + "\n\t") + "\n\t")
    var v4 string = (v3 + v3)
    var v5 uint16 = 1
    var v6 int = 42
    v11 := GenS2{F0: -850, F1: 1, F2: 9223372036854775807}
    for fi12 := 0; fi12 < 4; fi12++ {
        if true {
            v13 := GenS2{F0: v6, F1: genFn0(), F2: -1}
            v3 = "foo bar"
            _ = v13
        }
        switch v6 {
        case 0:
            var v14 float32 = 0.1
            _ = v14
        case 1:
            fmt.Println(v5, v5, v5)
        case -453:
            v1 += v1
        default:
            v4 = (v4 + v3)
        }
        v15 := GenS1{F0: (!(1.5 >= ((1.5 * (((-2.25 + 0.0) * -2.25) / (1e10 - (0.1 / (1.5 + 0.0))))) - float32(1e10)))), F1: (v6 >> ((uint(v5)) & 63)), F2: (v4 + v4)}
        _ = fi12
        _ = v15
    }
    var v16 uint = 18446744073709551615
    var v17 string = v4
    v16 = v16
    var v18 int8 = 42
    v19 := []int{(^(v6)), (v6 * (v6 & (v6 ^ 2147483647)))}
    fmt.Println(v6, v4)
    var v20 int8 = v18
    v21 := GenS2{F0: 4294967295, F1: 1, F2: -9223372036854775808}
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 int32 = 173
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: 152, F1: 42}
    var v25 uint16 = (^(v5))
    v26 := GenS2{F0: -1, F1: 2, F2: 9223372036854775807}
    if (((!(892 != -9223372036854775808)) && true) || ((v6 + v6) >= 128)) {
        if false {
            v6 -= 9223372036854775807
            v4 = v3
            for fi27 := 0; fi27 < 15; fi27++ {
                if (!true) {
                    var v28 int64 = 1
                    fmt.Println(v16, v4)
                    _ = v28
                }
                _ = fi27
            }
        } else {
            v20 += -128
            v18 = (v18 >> ((v16) & 63))
            var v29 string = ""
            _ = v29
        }
    }
    v25 += (v25 & v25)
    v5 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
