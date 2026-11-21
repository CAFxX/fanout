package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 string, p1 uint32) uint32 {
    v4 := []int{}
    v5 := GenS1{F0: 2}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 string = p0
        v8 := GenS1{F0: 221}
        if len(v4) > 0 {
            v4[0] = (-(fi6))
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 = 65535
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (^(p1))
}

func main() {
    var v1 uint = 290
    var v2 uint32 = 534
    var v3 int64 = 9223372036854775807
    v9 := []int{194, 71}
    var v10 int8 = 127
    var v11 bool = false
    fmt.Println(v3, v11)
    var v12 int32 = -1
    v11 = (786 < 4294967295)
    switch -2147483648 {
    case 0:
        v10 = v10
    default:
        v10 += 100
    }
    v13 := GenS1{F0: 127}
    v3 += -9223372036854775808
    var v14 float64 = (float64(3.14159) * 3.14159)
    v14 += 0.0
    switch "" {
    case "\n\t":
        switch -1 {
        case -1:
            for fi15 := 0; fi15 < 15; fi15++ {
                v2 -= (v2 >> ((425) & 63))
                v12 = (v12 - int32(v10))
                _ = fi15
            }
        case 0:
            v16 := GenS1{F0: 128}
            _ = v16
        }
    }
    switch ((("hello" + "foo bar") + ("hello" + "hello")) + ("\n\t" + "a")) {
    case "\n\t":
        v17 := GenS1{F0: 0}
        _ = v17
    default:
        v3 = v3
    }
    if v11 {
        v2 = (v2 & (v2 >> (((^(v1))) & 63)))
        var v18 int64 = v3
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
