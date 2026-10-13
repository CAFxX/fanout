package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int, p1 int64, p2 uint16) uint32 {
    var v6 int64 = (-(p1))
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{}
        p0 += p0
        var v9 bool = (0 < 42)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 636
}

func main() {
    var v1 uint8 = 42
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    v4 := []int{170, -1, -372}
    var v5 int64 = 42
    v10 := GenS1{F0: -14}
    var v11 int = -2147483648
    for fi12 := 0; fi12 < 11; fi12++ {
        v11 += (-(fi12))
        if (827 > 18446744073709551615) {
            v11 = fi12
            var v13 int32 = -2147483648
            _ = v13
        }
        _ = fi12
    }
    if false {
        v11 -= (v11 ^ (v11 << ((42) & 63)))
        if true {
            v14 := v3[356]
            _ = len(v4)
            _ = v14
        } else {
            _ = len(v2)
            switch 898 {
            case 1:
                for fi15 := 0; fi15 < 20; fi15++ {
                    if (916 != 1) {
                        var v16 uint32 = genFn0((v11 - 1), v5, 65535)
                        var v17 int = (^(v11))
                        delete(v3, fi15)
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(fi15, v1)
                    }
                    _ = fi15
                }
            }
            var v18 float64 = (-2.25 * 1.5)
            _ = v18
        }
    }
    v1 = (v1 % ((v1) | 1))
    delete(v3, -9223372036854775808)
    if true {
        fmt.Println(v11, v11, v11)
    }
    var v19 float32 = ((3.14159 * 1.5) + 3.14159)
    v5 = v5
    v1 = v1
    for fi20 := 0; fi20 < 7; fi20++ {
        v1 += (v1 % (((^(v1))) | 1))
        _ = fi20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        if true {
            v4 = append(v4, 42)
        }
        var v22 int64 = v5
        v3[32768] = 682
        _ = fi21
        _ = v22
    }
    fmt.Println(v19, v1)
    var v23 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
