package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 int8, p1 bool) int {
    v5 := []int{}
    p0 += p0
    var v6 float32 = 1.5
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    var v1 uint8 = 128
    var v2 int = 32767
    var v3 uint64 = 1
    var v4 int32 = 2147483647
    var v7 int8 = -64
    fmt.Println(v3, v7, v2)
    v7 -= v7
    var v8 int8 = v7
    if (49 >= (v8 >> ((4294967295) & 63))) {
        var v9 float32 = ((1.5 + ((-2.25 + float32((0.1 / 1e10))) - (1.5 + 1.5))) - 1.5)
        switch 2147483647 {
        case 127:
            for fi10 := 0; fi10 < 13; fi10++ {
                v3 -= v3
                v11 := []int{42, genFn0(75, false)}
                for fi12 := 0; fi12 < 6; fi12++ {
                    v7 -= v7
                    _ = fi12
                }
                _ = fi10
                _ = v11
            }
        case 0:
            v13 := GenS1{F0: 942, F1: true}
            _ = v13
        case 4294967295:
            v14 := GenS1{F0: 42, F1: (996 != 1)}
            _ = v14
        default:
            fmt.Println(v9, v8)
        }
        fmt.Println(v8, v3, v1)
        _ = v9
    } else {
        v15 := []int{(v2 - (^(v2))), v2}
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = v15
    }
    v4 = (v4 % ((v4) | 1))
    var v16 uint64 = (^(v3))
    if (-2.25 == 1.5) {
        var v17 uint8 = (v1 * (v1 >> ((608) & 63)))
        _ = v17
    } else {
        v18 := GenS2{F0: 341, F1: (v7 - (v7 << ((0) & 63)))}
        var v19 string = "\n\t"
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
