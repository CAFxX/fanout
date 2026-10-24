package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 bool }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 string; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int64 = 9223372036854775807
    var v3 float64 = (float64((-2.25 - ((-2.25 / 1.0) / (1.5 / (-2.25 * ((1.5 + 0.0) / (1e10 / float32(1e10)))))))) * float64(1e10))
    var v4 uint = 65535
    var v5 bool = (0 >= 100)
    for fi6 := 0; fi6 < 10; fi6++ {
        v4 += v4
        v7 := []int{32767}
        v8 := GenS1{F0: 42, F1: -9223372036854775808, F2: (!false)}
        _ = fi6
        _ = v7
        _ = v8
    }
    v2 -= v2
    if (!(v5 && (v5 || (true || (!v5))))) {
        v9 := []int{667, 127, 32767}
        var v10 int8 = -66
        _ = v10
        _ = v9
    }
    v11 := []int{128}
    if (v2 > (-(v2))) {
        v12 := GenS2{F0: 2147483647, F1: (v4 - v4)}
        _ = len(v11)
        _ = v12
    }
    _ = len(v11)
    if ((true || ((v3 - (v3 - v3)) <= (v3 * v3))) && false) {
        v13 := GenS3{F0: (("foo bar" + "a") + "a"), F1: v1, F2: 65535}
        _ = v13
    }
    v14 := GenS4{F0: 32767, F1: 42, F2: (v1 / ((0) | 1))}
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 float32 = 0.0
        _ = len(v11)
        v17 := GenS1{F0: -32768, F1: 510, F2: (true && v5)}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            v3 = v3
            var v20 int64 = v2
            _ = fi19
            _ = v20
        }
        if v5 {
            v4 = (v4 - (v4 - v4))
            v21 := GenS5{F0: (-2.25 != (0.1 + 1.5)), F1: 2}
            var v22 uint32 = 2147483648
            _ = v21
            _ = v22
        } else {
            var v23 int64 = 0
            _ = v23
        }
        _ = fi18
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 int16 = -63
        var v26 uint = (v4 + v4)
        _ = fi24
        _ = v25
        _ = v26
    }
    v1 += v1
    v2 = 504
    v11 = append(v11, 160)
    fmt.Println(v4)
    v1 -= v1
    if false {
        if v5 {
            v2 -= 922
        }
        v1 = v1
    } else {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if (v2 != -1) {
            _ = len(v11)
            v2 += v2
        } else {
            fmt.Println(v1, v1, v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
