package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int, p1 bool, p2 int8) uint8 {
    v6 := []int{128, p0}
    var v7 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 string = ("foo bar" + ("a" + ("hello" + ((((("foo bar" + ("a" + "hello")) + "") + "") + "") + ""))))
    var v2 uint8 = 209
    var v3 int32 = 1
    var v4 float64 = 0.0
    v5 := make(map[int]int)
    switch 229 {
    case 2147483647:
        v2 += 2
    case 32767:
        var v8 int = 2147483647
        _ = v8
    }
    var v9 bool = (760 >= 608)
    v2 += (v2 * v2)
    v3 += v3
    switch 412 {
    case 4294967295:
        if v9 {
            var v10 int32 = 1
            _ = v10
        } else {
            v11 := GenS1{F0: -2147483648}
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: (v3 / ((2147483647) | 1))}
                fmt.Println(v4, v9)
                var v14 int16 = 892
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
    case 42:
        v2 -= v2
    case -9223372036854775808:
        var v15 int8 = 0
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v9 = (-770 >= (^(fi16)))
        _ = fi16
    }
    v1 = ((v1 + v1) + "foo bar")
    var v17 float64 = ((0.0 / v4) - v4)
    var v18 float32 = -2.25
    v19 := GenS1{F0: -754}
    fmt.Println(v3, v4)
    var v20 uint64 = 2
    var v21 int16 = -32768
    switch v1 {
    case "\n\t":
        var v22 int = -469
        _ = v22
    default:
        v17 = 1.5
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v3 = v3
        v4 -= 1.5
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
