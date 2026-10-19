package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int }
func genFn0(p0 bool, p1 int64, p2 string) int8 {
    p2 = p2
    if (688 == 18446744073709551615) {
        switch 9223372036854775807 {
        case 127:
            var v7 float32 = 0.0
            _ = v7
        default:
            if p0 {
                var v8 int8 = 39
                switch 9223372036854775807 {
                case -2147483648:
                    p2 = p2
                default:
                    v8 += -1
                }
                _ = v8
            }
        }
        v9 := GenS1{F0: p1, F1: 18446744073709551615, F2: -9223372036854775808}
        _ = v9
    }
    p2 = ((p2 + ("a" + p2)) + "a")
    p0 = p0
    p0 = true
    _ = p0
    _ = p1
    _ = p2
    return 22
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "a"
    var v3 int16 = 32767
    var v4 int = -2147483648
    var v5 int16 = (v3 / ((v3) | 1))
    var v6 uint8 = 255
    var v10 uint8 = 117
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := []int{(v4 << ((0) & 63)), v4}
        _ = fi11
        _ = v12
    }
    v6 = 229
    var v13 bool = (127 > 100)
    v4 -= v4
    var v14 float64 = 1.5
    v13 = v13
    v15 := []int{v4}
    v5 = (v5 | (v5 % ((v3) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
