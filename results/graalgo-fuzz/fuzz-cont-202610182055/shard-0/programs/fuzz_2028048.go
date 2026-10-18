package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint }
func genFn0(p0 uint64, p1 int64) string {
    v6 := []int{128}
    var v7 int32 = 42
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(p0)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    var v1 bool = true
    v2 := []int{}
    var v3 int = 665
    v4 := make(map[string]int)
    var v5 int64 = 786
    if (324 < 2) {
        var v9 uint8 = 1
        _ = v9
    }
    v1 = (("foo bar" + "hello") <= "")
    if (2147483647 != 2) {
        v10 := GenS1{F0: 162, F1: 0, F2: 18446744073709551615}
        v1 = true
        _ = v10
    }
    if v1 {
        v1 = ((v5 + (v5 << ((42) & 63))) > (v5 - (v5 >> ((0) & 63))))
    } else {
        fmt.Println(v5, v1)
        v5 -= v5
    }
    var v11 string = (("foo bar" + ("a" + "foo bar")) + "a")
    v1 = v1
    v3 -= (-(v3))
    var v12 uint8 = 159
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int32 = 2147483647
        _ = fi13
        _ = v14
    }
    if len(v2) > 0 {
        v2[0] = (v3 ^ (v3 - v3))
    }
    v12 += v12
    v15 := GenS1{F0: 615, F1: 42, F2: 86}
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v12, v11, v5)
        v5 = (v5 - 472)
        _ = fi16
    }
    v3 += 127
    if ((v1 || true) || (42 >= 0)) {
        switch v3 {
        case 42:
            switch v3 {
            case 21:
                v3 = (v3 - (v3 | (v3 & (v3 << ((307) & 63)))))
            }
        case -317:
            var v17 string = v11
            _ = v17
        default:
            var v18 uint8 = 141
            _ = v18
        }
        fmt.Println(v11, v5, v3)
    } else {
        v3 = 1
    }
    var v19 float32 = float32(0.1)
    var v20 float32 = 1e10
    if ((2147483647 == 4294967295) || false) {
        v11 = (v11 + ("" + (((v11 + v11) + "") + v11)))
        var v21 int64 = (-(v5))
        for fi22 := 0; fi22 < 8; fi22++ {
            if len(v2) > 0 {
                v2[0] = v3
            }
            _ = fi22
        }
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
