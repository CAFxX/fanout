package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint64, p1 int, p2 int8) uint32 {
    var v6 uint64 = p0
    v7 := []int{(p1 >> ((668) & 63)), (p1 + p1)}
    var v8 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1() int8 {
    if true {
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    var v10 string = "hello"
    _ = v10
    return 1
}

func main() {
    var v1 float64 = (0.1 - 0.1)
    v2 := make(map[int]uint32)
    var v3 int64 = 1
    var v4 int16 = 1000
    var v5 int64 = 42
    v11 := GenS1{F0: 1, F1: 42}
    v5 = v3
    if ("foo bar" == "") {
        v3 = v3
        v5 = (v3 << ((192) & 63))
    }
    if (818 > 42) {
        v5 += (v5 - (-(v5)))
    }
    v12 := GenS1{F0: genFn1(), F1: 116}
    if true {
        var v13 int32 = 0
        v1 -= 1e10
        _ = v13
    }
    fmt.Println(v5, v3, v1)
    if false {
        v14 := GenS1{F0: -128, F1: 26}
        if true {
            v15 := GenS2{F0: false, F1: 42}
            for fi16 := 0; fi16 < 15; fi16++ {
                v3 += (v3 + 1)
                _ = fi16
            }
            _ = v15
        } else {
            var v17 uint16 = 65535
            var v18 bool = ((v3 + (v3 - (v3 >> ((2) & 63)))) != (v5 | v3))
            if true {
                v17 += v17
            } else {
                switch "hello" {
                case "a":
                    fmt.Println(v3, v5)
                default:
                    v1 += 0.1
                }
                if v18 {
                    v19 := []int{946, 56}
                    _ = v19
                } else {
                    var v20 int64 = (v5 % ((324) | 1))
                    v21 := GenS1{F0: -65, F1: 230}
                    _ = v20
                    _ = v21
                }
                var v22 int = -9223372036854775808
                _ = v22
            }
            _ = v17
            _ = v18
        }
        _ = v14
    }
    if ((-(v3)) <= v3) {
        var v23 int64 = v3
        _ = v23
    } else {
        var v24 uint8 = 128
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
