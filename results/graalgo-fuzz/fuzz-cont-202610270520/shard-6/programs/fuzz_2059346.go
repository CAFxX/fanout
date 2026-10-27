package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint8, p1 uint64, p2 uint) bool {
    var v7 int64 = 1
    var v8 float32 = (1.5 * float32(1e10))
    var v9 bool = (475 <= 1)
    for fi10 := 0; fi10 < 3; fi10++ {
        v7 = -9223372036854775808
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(p2, v9)
            if ((!(true || false)) && (696 > 0)) {
                var v12 uint = (p2 / ((p2) | 1))
                v13 := []int{-700, fi11, 676}
                if true {
                    if v9 {
                        var v14 int64 = 752
                        var v15 int8 = 127
                        var v16 uint32 = 1
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        v8 += (0.1 + (v8 - (v8 / 1e10)))
                    }
                } else {
                    v7 -= v7
                }
                _ = v12
                _ = v13
            }
            v17 := GenS1{F0: 776, F1: 1}
            _ = fi11
            _ = v17
        }
        switch int(p1) {
        case 1:
            fmt.Println(v8, p0)
        }
        _ = fi10
    }
    fmt.Println(v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 uint = 0
    var v2 uint32 = 2
    v3 := []uint32{(^(v2)), 4294967295}
    var v4 int32 = 811
    var v5 uint64 = 883
    var v6 uint = v1
    var v18 uint16 = 65535
    if true {
        v2 = 2147483648
    }
    var v19 bool = false
    v20 := GenS1{F0: (v2 >> ((790) & 63)), F1: (v18 >> ((v6) & 63))}
    var v21 int32 = (v4 - v4)
    v21 += v4
    v22 := GenS1{F0: 529, F1: (v18 ^ 623)}
    v3 = append(v3, v2)
    var v23 uint64 = (v5 << ((4294967295) & 63))
    var v24 float64 = (3.14159 * ((-2.25 - float64((-2.25 - 0.1))) * ((0.1 - 1.5) + (-2.25 * 0.0))))
    v1 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
