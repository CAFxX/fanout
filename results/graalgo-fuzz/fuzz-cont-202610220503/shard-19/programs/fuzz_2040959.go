package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int64) uint64 {
    p0 += (p0 << ((1) & 63))
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{1, fi5, fi5}
        var v7 int64 = p0
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(fi8)
        p0 = -1
        v9 := []int{799, (fi8 >> ((4294967295) & 63)), -9223372036854775808, 4294967295}
        _ = fi8
        _ = v9
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 42
    var v3 uint32 = v2
    var v4 float32 = (float32(-2.25) - (float32(0.0) * 3.14159))
    v10 := []int{383, 9223372036854775807, -100}
    if false {
        for fi11 := 0; fi11 < 13; fi11++ {
            if true {
                for fi12 := 0; fi12 < 7; fi12++ {
                    if (1 < -2147483648) {
                        var v13 uint32 = v3
                        fmt.Println(fi11, v3)
                        v4 = 0.0
                        _ = v13
                    }
                    v1[-1] = 0
                    v2 = (v3 | (v3 << ((65535) & 63)))
                    _ = fi12
                }
            }
            _ = fi11
        }
        v3 += v2
    }
    v2 = 4294967295
    var v14 int64 = -1
    v2 = 42
    switch (("" + "hello") + "foo bar") {
    case "a":
        if (true && true) {
            var v15 int8 = -123
            var v16 uint32 = v2
            _ = v15
            _ = v16
        }
    case "\n\t":
        var v17 int64 = (v14 + (v14 ^ -1))
        _ = v17
    }
    v3 += (v3 | (v2 | 65535))
    _ = len(v10)
    fmt.Println(v14, v4)
    if (0 <= 127) {
        v1[-381] = 2147483647
        if ((!(-9223372036854775808 <= (v14 | v14))) && true) {
            v3 = (v2 ^ v3)
        } else {
            fmt.Println(v14)
        }
        v18 := GenS1{F0: 32767}
        _ = v18
    } else {
        v19 := GenS1{F0: 361}
        _ = v19
    }
    if len(v10) > 0 {
        v10[0] = -2147483648
    }
    if true {
        v3 += v3
    }
    v14 -= -9223372036854775808
    v20 := GenS1{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
