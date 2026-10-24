package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int32) int8 {
    var v5 int8 = -128
    v6 := []int{615, 42, -45}
    _ = p0
    _ = v5
    _ = v6
    return v5
}

func main() {
    v1 := []int{4294967295, 4294967295, -18, 128}
    var v2 int8 = -1
    var v3 uint32 = 767
    v4 := []int{127, 739, 0}
    for fi7 := 0; fi7 < 20; fi7++ {
        v4 = append(v4, 4294967295)
        if false {
            v3 -= 42
            v3 = 4294967295
        } else {
            var v8 string = "\n\t"
            v9 := GenS1{F0: false}
            var v10 uint32 = (^(v3))
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    var v11 int8 = (v2 | -1)
    v12 := GenS1{F0: false}
    fmt.Println(v2, v3)
    if (42 == 852) {
        v11 = v11
        switch 9223372036854775807 {
        case 42:
            for fi13 := 0; fi13 < 16; fi13++ {
                v2 += (v11 & v2)
                _ = fi13
            }
        }
        v3 = (v3 ^ (v3 >> ((65535) & 63)))
    }
    v11 -= -128
    if (0.0 < 1.5) {
        var v14 int16 = 42
        if false {
            v14 -= v14
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 int16 = (v14 & -1)
                if false {
                    v17 := GenS2{F0: uint64(v3)}
                    if (2 >= 1) {
                        var v18 float64 = (3.14159 / float64(-2.25))
                        v19 := GenS1{F0: (genFn0(-2147483648) == v2)}
                        _ = v18
                        _ = v19
                    }
                    _ = v17
                } else {
                    for fi20 := 0; fi20 < 18; fi20++ {
                        v21 := GenS3{F0: 18446744073709551615}
                        _ = fi20
                        _ = v21
                    }
                }
                v16 = 1000
                _ = fi15
                _ = v16
            }
            v22 := GenS3{F0: 279}
            _ = v22
        }
        _ = v14
    }
    fmt.Println(v3, v11, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
