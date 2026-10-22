package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int64 }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 int, p1 string, p2 int32) int64 {
    v5 := GenS1{F0: false, F1: (p2 < p2), F2: -979}
    v6 := GenS1{F0: (int(p2) != p0), F1: false, F2: 9223372036854775807}
    var v7 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint32, p1 int64, p2 int8) int64 {
    for fi8 := 0; fi8 < 14; fi8++ {
        p0 = 2147483648
        var v9 uint32 = 0
        fmt.Println(p2, p2, p2)
        _ = fi8
        _ = v9
    }
    var v10 float64 = ((3.14159 / 1e10) / (3.14159 - 0.0))
    var v11 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return genFn0(-2147483648, v11, 0)
}

func main() {
    var v1 uint64 = 955
    v2 := []int{667}
    v3 := []string{"foo bar", "", "\n\t"}
    var v4 uint32 = 1
    v4 = v4
    v12 := GenS2{F0: 257, F1: 0}
    for fi13 := 0; fi13 < 3; fi13++ {
        switch (fi13 + 9223372036854775807) {
        case 150:
            if ((v1 * (v1 % ((v1) | 1))) >= (v1 / (((v1 - (v1 - (v1 * (v1 | 18446744073709551615))))) | 1))) {
                switch fi13 {
                case 9223372036854775807:
                    var v14 int = fi13
                    _ = v14
                }
                fmt.Println(v4, v1, v1)
            } else {
                v1 -= 1
                v4 = (v4 / ((v4) | 1))
            }
        case 42:
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 int8 = 0
                var v17 int16 = -628
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint64 = 1
            var v20 string = ""
            var v21 float32 = (1e10 - 0.0)
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        if true {
            v2 = append(v2, 1)
        }
        _ = fi13
    }
    v4 = 0
    var v22 int32 = -1
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = 171
    }
    fmt.Println(v22, v4, v1)
    if ((3.14159 - 1e10) != (0.1 - 1.5)) {
        for fi23 := 0; fi23 < 9; fi23++ {
            v22 += -2147483648
            _ = fi23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
