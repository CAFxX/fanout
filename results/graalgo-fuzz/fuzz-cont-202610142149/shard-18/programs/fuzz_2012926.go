package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }

func main() {
    var v1 int32 = 42
    var v2 uint16 = 847
    var v3 float32 = 1.5
    var v4 bool = true
    var v5 float64 = 0.1
    var v6 uint32 = 2
    var v7 float64 = (float64(1e10) + v5)
    var v8 int16 = -1
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: 4294967295, F1: v1, F2: 0}
        var v11 uint32 = v6
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        switch "foo bar" {
        case "":
            v13 := []int{(fi12 - (fi12 % ((fi12) | 1))), (fi12 ^ 4294967295), fi12}
            _ = v13
        case "\n\t":
            var v14 bool = true
            _ = v14
        default:
            var v15 uint64 = 2
            _ = v15
        }
        _ = fi12
    }
    fmt.Println(v1, v2)
    var v16 float64 = v7
    var v17 bool = (9223372036854775807 < -1)
    v3 += 0.0
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 string = "foo bar"
        for fi20 := 0; fi20 < 1; fi20++ {
            if v4 {
                fmt.Println(v16, v7, v19)
            } else {
                if v4 {
                    v21 := GenS1{F0: 1, F1: v1, F2: 388}
                    if true {
                        v8 = v8
                    } else {
                        var v22 int32 = (v1 << ((4294967295) & 63))
                        _ = v22
                    }
                    v1 -= (v1 << ((uint(v2)) & 63))
                    _ = v21
                }
                v23 := []int{(^(fi20)), (fi20 - (fi20 & (fi20 + fi18)))}
                var v24 int8 = 0
                _ = v23
                _ = v24
            }
            var v25 uint64 = 4294967295
            _ = fi20
            _ = v25
        }
        if v4 {
            var v26 int16 = v8
            var v27 uint64 = 2
            if true {
                for fi28 := 0; fi28 < 12; fi28++ {
                    v19 = v19
                    var v29 uint8 = 0
                    _ = fi28
                    _ = v29
                }
                v5 += v5
                v8 += v26
            }
            _ = v26
            _ = v27
        }
        _ = fi18
        _ = v19
    }
    v17 = ((!(v1 != v1)) || (!(814 != 285)))
    for fi30 := 0; fi30 < 9; fi30++ {
        fmt.Println(v3)
        v31 := GenS1{F0: 681, F1: v1, F2: 18446744073709551615}
        _ = fi30
        _ = v31
    }
    switch "" {
    case "hello":
        v5 += v16
    default:
        v32 := []int{32768, 2147483647, -1, 1}
        _ = v32
    }
    v33 := []int{0, 42}
    var v34 int8 = -1
    fmt.Println(v7, v3)
    var v35 int32 = (v1 ^ (v1 | (v1 | int32(v8))))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
