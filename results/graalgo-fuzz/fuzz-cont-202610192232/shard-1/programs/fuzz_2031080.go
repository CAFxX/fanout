package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0() uint {
    v6 := GenS1{F0: 4294967295, F1: 894}
    var v7 int16 = 0
    _ = v6
    _ = v7
    return 61
}

func main() {
    var v1 int16 = 42
    var v2 uint64 = 2
    var v3 uint8 = 42
    v4 := make(map[int]string)
    var v5 uint64 = v2
    v2 += 18446744073709551615
    v5 += v5
    switch "foo bar" {
    case "hello":
        var v8 uint16 = 65535
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 1; fi9++ {
            v5 = 42
            v10 := GenS2{F0: fi9, F1: 490}
            if (!true) {
                v5 -= (v2 | v5)
                v2 = 18446744073709551615
            } else {
                for fi11 := 0; fi11 < 6; fi11++ {
                    fmt.Println(fi11, v1)
                    _ = fi11
                }
            }
            _ = fi9
            _ = v10
        }
    case "":
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 float32 = (3.14159 + 1e10)
            if true {
                if (((true || false) && false) && false) {
                    v1 = (^(v1))
                    fmt.Println(v2, v13)
                }
                _ = len(v4)
                v13 -= v13
            } else {
                var v14 int16 = (v1 ^ (^(v1)))
                v13 += ((float32((0.0 * (3.14159 + 1.5))) - ((0.1 + 3.14159) + 0.1)) - 1.5)
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    }
    v5 -= (v2 & v5)
    v5 -= (^(v5))
    v1 = 1000
    v5 -= (v5 - (v5 * v2))
    if ((2 <= v3) && (!(42 == 544))) {
        var v15 int8 = -1
        var v16 int = 958
        _ = v15
        _ = v16
    }
    v17 := []int{127, -1, 0}
    fmt.Println(v5, v5, v2)
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 bool = true
        _ = fi18
        _ = v19
    }
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS2{F0: 2147483647, F1: v2}
            v2 = v2
            if (1e10 == 3.14159) {
                var v22 uint64 = v2
                _ = v22
            } else {
                var v23 uint16 = 65535
                _ = v23
            }
            _ = fi20
            _ = v21
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v4[fi24] = ""
            v25 := GenS2{F0: 42, F1: (^(v2))}
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
