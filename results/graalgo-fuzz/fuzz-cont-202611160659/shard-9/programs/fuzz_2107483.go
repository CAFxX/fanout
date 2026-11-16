package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0(p0 uint, p1 bool, p2 uint32) uint8 {
    var v7 int = 127
    var v8 uint = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 bool) string {
    var v9 int32 = 1
    p0 = (2 >= 0)
    _ = p0
    _ = v9
    return "foo bar"
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (1e10 / (1e10 * 1e10))
    var v3 uint16 = 162
    v4 := []string{"\n\t"}
    var v5 uint64 = 1
    var v6 string = ("foo bar" + ((("hello" + "") + "foo bar") + "hello"))
    var v10 string = v6
    for fi11 := 0; fi11 < 3; fi11++ {
        v2 = 0.1
        v3 = 254
        _ = fi11
    }
    v2 -= 3.14159
    if (127 > -328) {
        for fi12 := 0; fi12 < 9; fi12++ {
            v1[fi12] = fi12
            _ = fi12
        }
    }
    var v13 float32 = (0.1 * 0.1)
    var v14 int16 = 32767
    v14 += 1
    for fi15 := 0; fi15 < 1; fi15++ {
        v10 = ("\n\t" + v10)
        var v16 bool = (true || (127 <= int8(fi15)))
        _ = fi15
        _ = v16
    }
    v4 = append(v4, ((v10 + v6) + v6))
    v5 += (v5 % ((v5) | 1))
    var v17 float64 = (-2.25 / v2)
    var v18 int8 = -110
    if false {
        v19 := GenS1{F0: (!true), F1: 42}
        _ = v19
    }
    fmt.Println(v13, v10)
    var v20 float32 = v13
    if (!(!false)) {
        fmt.Println(v18, v20)
    } else {
        var v21 int32 = 1
        v17 += v17
        if (-9223372036854775808 == 128) {
            v2 = v2
        } else {
            v22 := GenS1{F0: true, F1: 42}
            var v23 bool = (-963 == (v14 % ((v14) | 1)))
            _ = v22
            _ = v23
        }
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
