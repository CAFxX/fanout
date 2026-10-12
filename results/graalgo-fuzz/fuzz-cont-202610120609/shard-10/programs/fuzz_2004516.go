package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint }
func genFn0(p0 string, p1 string, p2 uint16) uint64 {
    p1 = p0
    var v6 uint32 = uint32(p2)
    switch 42 {
    case 42:
        fmt.Println(v6, v6, p2)
    default:
        if true {
            v7 := GenS1{F0: 42, F1: v6}
            for fi8 := 0; fi8 < 7; fi8++ {
                fmt.Println(p2, fi8, p0)
                _ = fi8
            }
            _ = v7
        }
    }
    if (!(851 >= 167)) {
        v9 := []int{32767, -262, 1, -9223372036854775808}
        var v10 bool = (true || false)
        v11 := GenS1{F0: 0, F1: 2147483647}
        _ = v10
        _ = v11
        _ = v9
    } else {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 uint32 = v6
            if ((float32((1e10 * 3.14159)) == 1.5) && true) {
                p1 = ""
                p2 += uint16(v6)
                v14 := []int{(fi12 * fi12), fi12, fi12}
                _ = v14
            } else {
                fmt.Println(fi12)
                v13 += v6
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 int64 = 161
                    _ = fi15
                    _ = v16
                }
            }
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 157
}
func genFn1(p0 bool, p1 bool, p2 uint) int32 {
    p0 = p0
    var v17 uint = p2
    fmt.Println(v17)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 1
}

func main() {
    v1 := []string{"a", "hello", "a"}
    v2 := make(map[int]uint32)
    var v3 int64 = -9223372036854775808
    var v4 int64 = (v3 << ((26) & 63))
    var v5 uint64 = 18446744073709551615
    delete(v2, 830)
    var v18 int32 = -870
    var v19 bool = (1.5 > (-2.25 - (0.0 * 3.14159)))
    fmt.Println(v19)
    if (1 == -9223372036854775808) {
        var v20 int16 = -1
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 int64 = v4
            v2[127] = 871
            _ = fi21
            _ = v22
        }
        switch (("hello" + "foo bar") + "a") {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v20
    }
    var v23 int32 = 19
    v24 := GenS2{F0: (true && (!false)), F1: 0, F2: 0}
    v25 := v2[128]
    v2[-157] = (v25 / ((2147483648) | 1))
    var v26 int32 = (v18 - genFn1(false, true, 42))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
