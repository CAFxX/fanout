package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint }
func genFn0() uint64 {
    var v7 bool = true
    v7 = (((((float32(1.5) - 1.5) * (((1.5 / 1.5) * 3.14159) + 0.1)) - (1e10 / 1e10)) * -2.25) <= 3.14159)
    var v8 int32 = -1
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 int32 = -2147483648
    v5 := []int{48, 42}
    var v6 bool = (true || (!(false || true)))
    v9 := v2[443]
    var v10 uint16 = 856
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            if v6 {
                var v13 uint = 65535
                v9 = (fi12 * (fi11 * v9))
                _ = v13
            }
            var v14 float64 = 0.0
            if v6 {
                if (v6 && (1 <= -1)) {
                    var v15 uint32 = 2147483648
                    v16 := GenS1{F0: v15, F1: "\n\t"}
                    v9 = fi11
                    v17 := GenS2{F0: (-(v9)), F1: 42, F2: -128}
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    v4 -= v4
                }
                var v18 string = ("a" + "hello")
                v19 := GenS3{F0: v6, F1: 116, F2: 18446744073709551615}
                _ = v18
                _ = v19
            } else {
                if (2147483647 <= 1) {
                    v14 = (v14 - float64((-2.25 + 0.0)))
                    var v20 int8 = 127
                    _ = v20
                }
                switch 2147483647 {
                case 127:
                    if len(v5) > 0 {
                        v5[0] = v9
                    }
                }
            }
            _ = fi12
            _ = v14
        }
        v2[2] = 127
        _ = len(v5)
        _ = fi11
    }
    if len(v5) > 0 {
        v5[0] = v9
    }
    var v21 int16 = 479
    fmt.Println(v9)
    fmt.Println(v9)
    v10 -= v10
    fmt.Println(v4, v9, v6)
    var v22 bool = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
