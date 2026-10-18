package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() uint32 {
    var v7 uint64 = 42
    v7 -= (v7 + v7)
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 string) bool {
    fmt.Println(p1, p0)
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: 128}
        fmt.Println(p0, p0)
        var v10 uint64 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{892, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v11
    return (2 != 42)
}

func main() {
    var v1 int16 = 32767
    var v2 float64 = 1e10
    var v3 uint8 = 255
    v4 := make(map[string]uint32)
    v5 := []uint32{}
    v6 := make(map[uint32]string)
    switch -1 {
    case -9223372036854775808:
        if (2 >= v3) {
            var v12 int = 32768
            _ = v12
        }
    }
    v2 -= v2
    delete(v6, 35)
    if (("\n\t" == "\n\t") || (127 < 100)) {
        if (true && false) {
            v2 += v2
            var v13 float64 = v2
            _ = v13
        }
    } else {
        fmt.Println(v3, v3, v3)
    }
    fmt.Println(v1, v1, v1)
    _ = len(v5)
    v2 -= (v2 - ((v2 + (v2 / v2)) + v2))
    v14 := v6[uint32(v3)]
    for fi15 := 0; fi15 < 11; fi15++ {
        v2 = (v2 + float64((0.1 / (1e10 + (1e10 * 1e10)))))
        for fi16 := 0; fi16 < 13; fi16++ {
            v2 = (float64(-2.25) - (v2 - (((((1.5 + -2.25) + 1.5) / -2.25) * 0.0) + v2)))
            var v17 int64 = -9223372036854775808
            v17 += (v17 | (v17 | (v17 & (v17 ^ (v17 % (((v17 | (^(v17)))) | 1))))))
            _ = len(v5)
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    fmt.Println(v1, v14)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
