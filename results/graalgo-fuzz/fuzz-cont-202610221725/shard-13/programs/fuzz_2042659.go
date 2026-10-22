package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint64 }
func genFn0() uint8 {
    var v4 string = "a"
    var v5 uint = 953
    v4 = "foo bar"
    var v6 uint = 255
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int8) int64 {
    p0 = p0
    fmt.Println(p0)
    p0 = (p0 ^ (p0 & -89))
    _ = p0
    return 42
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]string)
    var v3 uint8 = (v1 * v1)
    for fi7 := 0; fi7 < 10; fi7++ {
        if (!(42 <= -1)) {
            v8 := []int{}
            switch (fi7 >> ((42) & 63)) {
            case 916:
                for fi9 := 0; fi9 < 19; fi9++ {
                    v10 := GenS1{F0: 340, F1: 2, F2: 964}
                    fmt.Println(fi9, v1)
                    var v11 int = 128
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case 0:
                var v12 int16 = -1
                _ = v12
            case 1:
                if true {
                    if false {
                        var v13 string = ""
                        v3 = v3
                        v14 := v2[4294967295]
                        _ = v13
                        _ = v14
                    } else {
                        v15 := GenS1{F0: 42, F1: 18446744073709551615, F2: 42}
                        var v16 float64 = (-2.25 + float64(-2.25))
                        _ = v15
                        _ = v16
                    }
                }
            }
            switch 4294967295 {
            case 9223372036854775807:
                v17 := GenS1{F0: 0, F1: 517, F2: 2}
                _ = v17
            case 0:
                v3 += v3
            }
            _ = v8
        }
        if (-9223372036854775808 > 0) {
            v18 := GenS1{F0: 255, F1: 0, F2: 18446744073709551615}
            fmt.Println(v3)
            _ = v18
        } else {
            _ = len(v2)
        }
        _ = fi7
    }
    v1 += (v3 ^ genFn0())
    v19 := []int{-114, -2147483648}
    v1 -= (v3 * genFn0())
    fmt.Println(v1, v1)
    delete(v2, 2147483647)
    var v20 uint = 485
    fmt.Println(v20)
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 uint64 = 18446744073709551615
        v23 := GenS1{F0: 0, F1: (v22 - (v22 | 1)), F2: (^(v22))}
        if true {
            var v24 uint64 = v22
            _ = v24
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
