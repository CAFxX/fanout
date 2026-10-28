package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint64 = 280
    v2 := []int{-9223372036854775808, 188}
    var v3 int32 = 2147483647
    v4 := []string{(("\n\t" + (("hello" + "") + "a")) + ""), "a", "a", "a"}
    var v5 int = 2147483647
    var v6 string = "\n\t"
    v5 += v5
    if false {
        var v7 float64 = 1e10
        v8 := GenS1{F0: 817, F1: -478}
        var v9 uint32 = 2147483647
        _ = v7
        _ = v8
        _ = v9
    }
    _ = len(v2)
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: 543, F1: 459}
        var v12 float64 = (-2.25 - (3.14159 + 1e10))
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = 2147483647
    v13 := GenS1{F0: 2, F1: 32767}
    fmt.Println(v5)
    v5 -= -1
    if ((661 == 2147483647) && true) {
        switch v5 {
        case 791:
            v14 := GenS1{F0: 32767, F1: -1}
            _ = v14
        default:
            v4 = append(v4, (v6 + (v6 + v6)))
        }
        var v15 int8 = 1
        _ = len(v4)
        _ = v15
    }
    if true {
        if (100 >= 42) {
            switch (v5 << ((4294967295) & 63)) {
            case 309:
                var v16 uint = 4294967295
                _ = v16
            case -659:
                for fi17 := 0; fi17 < 9; fi17++ {
                    v4 = append(v4, v6)
                    switch fi17 {
                    case 127:
                        var v18 int = fi17
                        _ = v18
                    case 1:
                        if true {
                            v19 := GenS2{F0: -238, F1: (-(fi17))}
                            var v20 uint64 = v1
                            v20 = 511
                            _ = v19
                            _ = v20
                        }
                    default:
                        var v21 uint16 = 380
                        _ = v21
                    }
                    _ = fi17
                }
            default:
                for fi22 := 0; fi22 < 15; fi22++ {
                    var v23 float32 = (3.14159 - (0.0 * float32(float64((0.1 + float32((0.1 / float64(((((float32((1e10 / -2.25)) / 1e10) + 0.1) + (0.0 - (0.1 - ((float32(3.14159) + 1.5) - 0.1)))) / 1e10)))))))))
                    v24 := GenS2{F0: -9223372036854775808, F1: (fi22 * (v5 | (fi22 - 2147483647)))}
                    var v25 bool = (94 != 32767)
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
        } else {
            fmt.Println(v5, v1)
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    v26 := GenS3{F0: 371}
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
