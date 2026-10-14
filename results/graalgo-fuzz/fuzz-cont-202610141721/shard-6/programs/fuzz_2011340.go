package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int }
type GenS2 struct { F0 int64; F1 uint; F2 bool }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 int32) int8 {
    if ((float32(-2.25) * float32(-2.25)) <= (0.0 * 0.1)) {
        fmt.Println(p0)
        var v6 uint32 = 2147483647
        _ = v6
    } else {
        fmt.Println(p0)
        v7 := GenS1{F0: -1, F1: (1.5 >= 0.1), F2: -749}
        _ = v7
    }
    var v8 int64 = 42
    var v9 bool = (!(false || true))
    v8 = v8
    if v9 {
        var v10 float32 = -2.25
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := []int{(fi11 >> ((42) & 63)), fi11}
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        var v13 int16 = 1
        if (42 < 18446744073709551615) {
            var v14 string = ""
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 uint16 = 890
                fmt.Println(v9, v16)
                v17 := GenS1{F0: (v8 << ((444) & 63)), F1: false, F2: fi15}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        } else {
            p0 -= (-(p0))
        }
        _ = v13
    }
    _ = p0
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    v4 := []uint32{1}
    var v5 int = 961
    _ = len(v3)
    v4 = append(v4, 45)
    for fi18 := 0; fi18 < 3; fi18++ {
        v5 += (v5 + (fi18 * 86))
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 float64 = 3.14159
            v3[0] = 715
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v21 := GenS2{F0: 0, F1: 42, F2: true}
    for fi22 := 0; fi22 < 7; fi22++ {
        v23 := GenS2{F0: 0, F1: 0, F2: false}
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        v1 -= 65535
        _ = fi24
    }
    v25 := GenS3{F0: -128, F1: 657}
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v26 int = v5
    var v27 string = ("a" + ("foo bar" + ""))
    var v28 float32 = (-2.25 - 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
