package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }
func genFn0(p0 bool, p1 int32, p2 bool) int {
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{fi5, (fi5 ^ (fi5 / ((fi5) | 1)))}
        _ = fi5
        _ = v6
    }
    var v7 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}
func genFn1() uint {
    var v8 uint64 = 4294967295
    v9 := GenS1{F0: 1, F1: 65535, F2: 128}
    v8 -= 42
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := []int{128, 128, 9223372036854775807, 128}
    var v4 int16 = 0
    v10 := GenS1{F0: -2147483648, F1: 4294967295, F2: -58}
    if true {
        var v11 bool = (false || false)
        _ = v11
    } else {
        v4 -= (v4 & (v4 << ((genFn1()) & 63)))
    }
    var v12 int32 = 119
    if (1000 > 452) {
        v1[2] = (!(false || ((-2.25 - float32(1.5)) > float32((0.0 * 1e10)))))
    }
    v13 := GenS1{F0: 1, F1: 1, F2: 9223372036854775807}
    fmt.Println(v4, v12)
    if ((!(((1 <= 18446744073709551615) && false) && true)) || true) {
        if false {
            switch -9223372036854775808 {
            case -722:
                for fi14 := 0; fi14 < 12; fi14++ {
                    v15 := GenS1{F0: v12, F1: 1, F2: (fi14 & (^(fi14)))}
                    var v16 int16 = (v4 / (((v4 % (((v4 % ((v4) | 1))) | 1))) | 1))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case 1:
                var v17 uint16 = 0
                _ = v17
            default:
                if (!(87 <= 127)) {
                    fmt.Println(v4)
                    _ = len(v1)
                }
            }
        } else {
            var v18 string = "hello"
            _ = v18
        }
    }
    var v19 int8 = -1
    _ = len(v2)
    fmt.Println(v4, v19)
    if false {
        v4 = (v4 & (v4 % ((v4) | 1)))
    }
    fmt.Println(v19, v12, v12)
    fmt.Println(v19)
    v4 += v4
    v12 += 0
    _ = len(v3)
    v12 += (v12 | 363)
    var v20 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
