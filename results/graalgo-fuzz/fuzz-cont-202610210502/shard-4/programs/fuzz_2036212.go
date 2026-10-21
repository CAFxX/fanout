package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint32 }
func genFn0(p0 int64, p1 int16) int16 {
    switch -9223372036854775808 {
    case 858:
        var v6 int16 = -1
        _ = v6
    default:
        if ((!(!true)) || false) {
            var v7 uint32 = 2147483647
            _ = v7
        } else {
            p1 = p1
        }
    }
    if false {
        for fi8 := 0; fi8 < 19; fi8++ {
            p0 = (p0 << ((18446744073709551615) & 63))
            v9 := GenS1{F0: 180, F1: (^(p1)), F2: 2147483648}
            _ = fi8
            _ = v9
        }
        v10 := []int{404, 42, 32768, 1}
        fmt.Println(p0, p0)
        _ = v10
    } else {
        var v11 float64 = 0.1
        _ = v11
    }
    switch -2147483648 {
    case 0:
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := []int{fi12, 127, (fi12 * 9223372036854775807), 0}
            _ = fi12
            _ = v13
        }
    }
    var v14 uint8 = 127
    _ = p0
    _ = p1
    _ = v14
    return 32767
}
func genFn1(p0 uint16, p1 int) string {
    p0 += p0
    if (false && false) {
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 uint32 = 1
            _ = fi15
            _ = v16
        }
    } else {
        fmt.Println(p1)
    }
    var v17 uint = 641
    var v18 string = "\n\t"
    v19 := []int{387}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v19
    return v18
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 1.5
    v3 := []uint32{2147483648, 2147483647}
    var v4 int8 = 100
    v5 := make(map[int]bool)
    v1 = v1
    v20 := v5[32767]
    var v21 bool = (0 == 4294967295)
    v3 = append(v3, 2147483647)
    v21 = v20
    if ((!(("" + "hello") != "\n\t")) || (v20 && v20)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v1, v20, v2)
    } else {
        v2 = v2
    }
    if true {
        delete(v5, -938)
    }
    if (172 >= 1) {
        v22 := GenS1{F0: 723, F1: (v1 & genFn0(-358, genFn0(249, v1))), F2: 2147483647}
        for fi23 := 0; fi23 < 4; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                v2 += v2
                _ = fi24
            }
            var v25 int64 = -329
            _ = fi23
            _ = v25
        }
        _ = v22
    }
    v26 := GenS1{F0: 42, F1: (v1 / (((v1 - genFn0(-1, v1))) | 1)), F2: 1}
    var v27 int16 = (^(v1))
    for fi28 := 0; fi28 < 15; fi28++ {
        for fi29 := 0; fi29 < 15; fi29++ {
            v27 += v1
            var v30 uint = 250
            _ = fi29
            _ = v30
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi28
    }
    var v31 int16 = v1
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
