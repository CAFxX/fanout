package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int = 127
    var v2 int = v1
    var v3 int32 = 1
    var v4 uint64 = 4294967295
    var v5 int32 = -189
    if (!true) {
        var v6 uint16 = 869
        _ = v6
    } else {
        switch -2147483648 {
        case 0:
            v5 += (v3 << ((868) & 63))
        case -1:
            if false {
                v7 := []int{127}
                _ = v7
            } else {
                v8 := GenS1{F0: 767, F1: (v2 | v1), F2: 1}
                v3 -= v5
                _ = v8
            }
        case 9223372036854775807:
            if true {
                if (2 >= -90) {
                    var v9 int64 = 0
                    _ = v9
                } else {
                    v10 := []int{(v1 % ((-1) | 1)), (-(v1))}
                    v10 = append(v10, v2)
                    if (false && (true || (false && false))) {
                        var v11 uint16 = 0
                        _ = v11
                    }
                    _ = v10
                }
                var v12 uint = 42
                var v13 uint64 = v4
                _ = v12
                _ = v13
            } else {
                v14 := []int{}
                if ((v4 * (^(v4))) >= v4) {
                    v15 := GenS2{F0: 0}
                    var v16 uint64 = (v4 ^ v4)
                    fmt.Println(v16, v2, v3)
                    _ = v15
                    _ = v16
                }
                _ = v14
            }
        default:
            v17 := []int{777}
            _ = v17
        }
    }
    v18 := []int{(v2 * (v2 | (-(v1))))}
    v18 = append(v18, (v2 ^ -9223372036854775808))
    if true {
        var v19 uint8 = uint8(v4)
        _ = v19
    }
    if (((-1 > -203) || (9223372036854775807 != -1)) || (!false)) {
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 string = "hello"
            var v22 int64 = 9223372036854775807
            _ = fi20
            _ = v21
            _ = v22
        }
        v1 = v1
        var v23 uint = 338
        _ = v23
    }
    if true {
        switch "hello" {
        case "\n\t":
            v1 += (^(v2))
        case "a":
            if true {
                var v24 float32 = 0.0
                v4 = v4
                var v25 uint32 = 2
                _ = v24
                _ = v25
            }
        case "hello":
            _ = len(v18)
        }
        v1 = v1
    }
    var v26 uint8 = 43
    if len(v18) > 0 {
        v18[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
