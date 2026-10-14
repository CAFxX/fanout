package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int }
func genFn0(p0 uint8, p1 uint16, p2 uint64) int8 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{(^(fi5))}
        _ = fi5
        _ = v6
    }
    var v7 uint = 854
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int = 128
    v3 := make(map[string]int)
    var v4 string = "hello"
    if false {
        v1 = (v1 | (v1 ^ 0))
    } else {
        if (-2.25 >= 0.1) {
            switch -243 {
            case 4294967295:
                v8 := []int{0}
                _ = v8
            }
            v4 = v4
            v9 := GenS1{F0: -1, F1: v1, F2: (v2 + 128)}
            _ = v9
        }
        switch "" {
        case "foo bar":
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: genFn0(42, 42, 1), F1: v1, F2: -9223372036854775808}
                var v12 uint = 4294967295
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            var v13 int = v2
            _ = v13
        }
    }
    v14 := []int{v2, (v2 << ((2) & 63)), v2, (v2 + v2)}
    var v15 uint32 = 386
    v14 = append(v14, (v2 | (v2 * -9223372036854775808)))
    var v16 float64 = 0.0
    v1 += (v1 & v1)
    if (true || (false && true)) {
        if len(v14) > 0 {
            v14[0] = v2
        }
        _ = len(v14)
    }
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
