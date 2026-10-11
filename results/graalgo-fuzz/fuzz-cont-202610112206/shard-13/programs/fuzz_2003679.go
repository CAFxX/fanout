package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint }
type GenS2 struct { F0 int32; F1 int }
func genFn0() int {
    var v5 int64 = -1
    var v6 uint16 = 761
    var v7 uint64 = 51
    var v8 float32 = -2.25
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1() uint16 {
    var v9 int64 = -9223372036854775808
    var v10 float64 = float64(3.14159)
    v11 := []int{}
    v9 += (v9 - (v9 / ((-1) | 1)))
    _ = v10
    _ = v11
    _ = v9
    return 65535
}

func main() {
    var v1 string = ""
    var v2 uint16 = 65535
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    for fi12 := 0; fi12 < 11; fi12++ {
        v1 = ("\n\t" + v1)
        fmt.Println(fi12, v1, fi12)
        _ = fi12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS1{F0: -1, F1: -41, F2: 1}
        v2 += v2
        _ = fi13
        _ = v14
    }
    v15 := v3[42]
    var v16 bool = false
    if (-1 < 600) {
        v16 = v16
    } else {
        if v16 {
            if true {
                var v17 float32 = float32(3.14159)
                if v16 {
                    if true {
                        v15 = v15
                    }
                    _ = len(v4)
                    v18 := []int{1, 42, 127}
                    _ = v18
                } else {
                    var v19 uint = 18446744073709551615
                    if ((v1 + v1) <= "foo bar") {
                        var v20 int8 = -1
                        _ = v20
                    } else {
                        if ((("\n\t" + v1) + ("hello" + v1)) == v1) {
                            v1 = "a"
                        }
                        v19 -= v19
                    }
                    v21 := []int{genFn0(), -9223372036854775808, 127}
                    _ = v19
                    _ = v21
                }
                _ = v17
            }
            var v22 uint8 = 190
            v1 = v1
            _ = v22
        }
        switch 9223372036854775807 {
        case -49:
            v16 = v16
        default:
            var v23 int8 = 1
            _ = v23
        }
        v24 := GenS2{F0: -1, F1: 9223372036854775807}
        _ = v24
    }
    v2 = v2
    fmt.Println(v16, v16)
    delete(v3, 65)
    v1 = "\n\t"
    if (!v16) {
        var v25 uint = 4294967295
        _ = v25
    }
    v16 = false
    v16 = (uint8(v2) == 128)
    v26 := GenS1{F0: -1, F1: 0, F2: 2}
    if (100 > -106) {
        var v27 int16 = 0
        v2 = 2
        var v28 int16 = v27
        _ = v27
        _ = v28
    }
    var v29 uint32 = v15
    var v30 uint = 69
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
