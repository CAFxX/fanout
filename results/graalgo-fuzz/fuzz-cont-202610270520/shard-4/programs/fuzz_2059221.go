package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool; F1 string; F2 uint32 }
func genFn0(p0 bool, p1 int16) uint {
    fmt.Println(p1, p1)
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: -1, F1: 2}
        p0 = p0
        switch (fi5 - (fi5 * fi5)) {
        case -1:
            for fi7 := 0; fi7 < 20; fi7++ {
                fmt.Println(fi5)
                _ = fi7
            }
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(p1)
    p1 += p1
    _ = p0
    _ = p1
    return 886
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 255
    var v3 uint64 = 42
    v4 := []int{}
    v2 = v2
    for fi8 := 0; fi8 < 4; fi8++ {
        if false {
            var v9 float32 = 0.0
            _ = v9
        }
        var v10 uint8 = 2
        v1 = append(v1, 517)
        _ = fi8
        _ = v10
    }
    if ((684 <= 1) || false) {
        var v11 uint = 732
        v11 = 695
        if false {
            v3 -= (v3 >> (((v11 & v11)) & 63))
        }
        _ = v11
    }
    if (0 == 2) {
        if len(v4) > 0 {
            v4[0] = 359
        }
        var v12 float32 = 0.0
        if len(v4) > 0 {
            v4[0] = 127
        }
        _ = v12
    } else {
        fmt.Println(v3)
        v13 := GenS1{F0: -1, F1: 32767}
        _ = v13
    }
    var v14 int64 = 555
    if (((("a" + ("" + "")) + "") == "hello") && false) {
        fmt.Println(v14, v14, v3)
        _ = len(v4)
    } else {
        var v15 uint = 42
        if true {
            fmt.Println(v2, v2)
        } else {
            v15 += 0
        }
        _ = v15
    }
    fmt.Println(v2, v14)
    v3 = v3
    v16 := GenS2{F0: (float32(3.14159) == (-2.25 / float32((0.1 + (0.1 + ((float64(((0.0 - (float32(0.0) / float32((((0.0 / -2.25) * 0.1) + 3.14159)))) - 1e10)) * 1e10) - 0.0)))))), F1: "foo bar", F2: 565}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
