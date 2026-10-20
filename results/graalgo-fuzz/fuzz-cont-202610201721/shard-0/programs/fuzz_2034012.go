package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 bool }
func genFn0(p0 bool, p1 uint16, p2 int8) uint64 {
    p0 = true
    p2 = (p2 * (p2 - (-(p2))))
    if p0 {
        if (((0.1 + 3.14159) + 3.14159) < ((-2.25 / 3.14159) - 1e10)) {
            if (!false) {
                fmt.Println(p2, p1, p1)
                p0 = p0
            } else {
                var v6 uint8 = 28
                fmt.Println(p2, v6)
                v7 := []int{9223372036854775807}
                _ = v6
                _ = v7
            }
            for fi8 := 0; fi8 < 9; fi8++ {
                for fi9 := 0; fi9 < 20; fi9++ {
                    if (p1 < 0) {
                        v10 := GenS1{F0: "\n\t", F1: 4294967295, F2: p0}
                        v11 := []int{(fi8 / ((706) | 1)), (-(fi9))}
                        if (4294967295 >= 1) {
                            fmt.Println(p2, p0)
                        }
                        _ = v10
                        _ = v11
                    }
                    _ = fi9
                }
                if p0 {
                    p0 = false
                    p1 -= (p1 * (p1 << ((0) & 63)))
                } else {
                    v12 := GenS1{F0: ("foo bar" + "hello"), F1: 1, F2: (-1 == -1)}
                    var v13 float64 = 1e10
                    var v14 uint64 = 42
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 uint16 = 367
                _ = fi8
                _ = v15
            }
        }
        p2 = 1
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint8 = 42
            v18 := []int{fi16, fi16, int(p2)}
            _ = fi16
            _ = v17
            _ = v18
        }
        switch "\n\t" {
        case "foo bar":
            var v19 string = "hello"
            _ = v19
        }
        p0 = (int64(p1) >= -1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 string = "\n\t"
    v3 := make(map[uint32]int)
    var v4 int64 = 42
    var v5 int16 = 0
    v4 = v4
    v20 := GenS1{F0: v2, F1: 42, F2: true}
    if (true || ((!false) && ((^(v4)) != v4))) {
        v21 := GenS1{F0: "foo bar", F1: 4294967295, F2: (884 > 2147483647)}
        _ = v21
    }
    fmt.Println(v2)
    v5 -= -32768
    v22 := []int{-598, 127, 2147483647}
    var v23 float64 = -2.25
    v5 = v5
    delete(v3, 4294967295)
    var v24 float32 = 1.5
    if true {
        var v25 uint = v1
        fmt.Println(v5)
        _ = v25
    }
    var v26 string = ("foo bar" + "hello")
    v2 = (v2 + v26)
    v27 := GenS1{F0: v26, F1: 428, F2: true}
    v24 -= 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
