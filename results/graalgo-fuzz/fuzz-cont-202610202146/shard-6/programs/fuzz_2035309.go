package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int8) uint8 {
    v6 := []int{0}
    var v7 uint32 = 42
    _ = p0
    _ = v6
    _ = v7
    return 127
}
func genFn1() string {
    v8 := GenS1{F0: -9223372036854775808}
    v9 := GenS1{F0: 32767}
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    v1 := []uint32{809, 559, 318, 0}
    var v2 uint8 = 255
    var v3 float64 = 3.14159
    var v4 int32 = 0
    v5 := []string{}
    v4 -= (^(v4))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v5)
    fmt.Println(v4, v2, v4)
    var v10 int32 = 42
    if len(v5) > 0 {
        v5[0] = ((("hello" + "hello") + "\n\t") + "a")
    }
    switch 1 {
    case 2147483647:
        v10 -= (v4 << ((18446744073709551615) & 63))
    case 0:
        var v11 uint = 4294967295
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int8 = 127
        var v14 uint = 4294967295
        v10 -= v4
        _ = fi12
        _ = v13
        _ = v14
    }
    v10 = v10
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(v4, v4)
            v17 := GenS1{F0: (fi16 / (((fi15 << ((161) & 63))) | 1))}
            v2 += (v2 << ((1) & 63))
            _ = fi16
            _ = v17
        }
        v10 += v10
        if (0.1 <= (0.0 + 1e10)) {
            var v18 uint64 = 42
            v2 += (v2 >> ((uint(v2)) & 63))
            v3 = v3
            _ = v18
        }
        _ = fi15
    }
    var v19 string = ("a" + "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
