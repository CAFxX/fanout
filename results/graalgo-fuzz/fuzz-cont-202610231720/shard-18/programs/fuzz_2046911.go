package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint32) int8 {
    fmt.Println(p0)
    v4 := []int{32768, -959}
    v4 = append(v4, -1)
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: p1}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 100
}

func main() {
    var v1 string = (("" + "a") + ("hello" + "a"))
    var v2 uint32 = 438
    v3 := []int{32767, -727, 9223372036854775807}
    v1 = v1
    v3 = append(v3, 1)
    for fi7 := 0; fi7 < 10; fi7++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi7
    }
    v8 := GenS1{F0: (v2 % ((2147483647) | 1))}
    for fi9 := 0; fi9 < 19; fi9++ {
        v2 -= v2
        _ = fi9
    }
    v3 = append(v3, -9223372036854775808)
    v10 := GenS1{F0: 348}
    if (1.5 <= float32(-2.25)) {
        v2 -= v2
    } else {
        if (373 >= 2147483647) {
            var v11 int8 = genFn0(2, v2)
            var v12 uint = 0
            _ = v11
            _ = v12
        }
    }
    var v13 int8 = -113
    v14 := GenS1{F0: 42}
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    v13 = (^(v13))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
}
