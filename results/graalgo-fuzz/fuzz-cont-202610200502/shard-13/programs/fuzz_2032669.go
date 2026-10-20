package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 uint; F1 uint16; F2 int64 }
func genFn0() int {
    if true {
        v6 := GenS1{F0: 42, F1: -9223372036854775808, F2: false}
        _ = v6
    }
    var v7 uint = 745
    var v8 string = "foo bar"
    v9 := []int{128, -834}
    _ = v7
    _ = v8
    _ = v9
    return 205
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 2147483647
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 18446744073709551615
    var v5 uint16 = 32767
    fmt.Println(v5, v2)
    v10 := []int{(^(v2))}
    v10 = append(v10, (^(v2)))
    if false {
        fmt.Println(v2, v4, v2)
        _ = len(v10)
    }
    fmt.Println(v2)
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := v3[2147483647]
        for fi13 := 0; fi13 < 2; fi13++ {
            if true {
                v14 := GenS2{F0: 42, F1: v5, F2: -1}
                var v15 int8 = 127
                _ = v14
                _ = v15
            } else {
                v4 += v4
            }
            if (false || (!((371 >= 353) && true))) {
                v12 += 4294967295
            }
            fmt.Println(v4, v2, fi11)
            _ = fi13
        }
        if (!(((0.0 - (0.1 + float32(float64(((3.14159 / float32(0.1)) / -2.25))))) - 1e10) > ((3.14159 * (-2.25 - 0.0)) + (0.0 / 1.5)))) {
            var v16 int = v2
            var v17 uint16 = v5
            _ = v16
            _ = v17
        } else {
            switch ("a" + "\n\t") {
            case "\n\t":
                v12 += v12
            case "hello":
                var v18 uint64 = v4
                _ = v18
            case "a":
                if false {
                    v3[394] = v12
                    if ((1 <= 32767) || (!(1 > -992))) {
                        var v19 int = (fi11 & (-(v2)))
                        _ = v19
                    }
                }
            }
        }
        _ = fi11
        _ = v12
    }
    var v20 int = genFn0()
    if (false && (0 >= 14)) {
        v20 -= v2
        v2 = (v2 + 32767)
        var v21 uint32 = 0
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
