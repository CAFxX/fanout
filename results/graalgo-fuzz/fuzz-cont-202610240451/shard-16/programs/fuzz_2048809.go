package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 string }
type GenS2 struct { F0 bool }
func genFn0(p0 int8, p1 int8) int8 {
    var v5 int64 = -1
    if (42 >= 32767) {
        v5 = (v5 | 0)
        v5 += 42
        p0 = (p0 * (p0 << ((18446744073709551615) & 63)))
    } else {
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: 0, F1: 555, F2: "foo bar"}
            _ = fi6
            _ = v7
        }
        if true {
            var v8 uint64 = 2
            var v9 int16 = int16(p1)
            _ = v8
            _ = v9
        } else {
            fmt.Println(v5)
            v10 := []int{}
            var v11 int8 = p0
            _ = v10
            _ = v11
        }
        if true {
            v12 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: ("foo bar" + "foo bar")}
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 int, p1 uint, p2 uint8) int {
    p2 = (p2 & 139)
    if (0.0 == ((((1e10 / (float64(0.0) - 0.1)) / (3.14159 + 3.14159)) / -2.25) - float64(-2.25))) {
        if false {
            var v13 int16 = 42
            _ = v13
        } else {
            v14 := []int{p0, 870}
            _ = v14
        }
        p2 -= (p2 | p2)
        p2 = 2
    } else {
        fmt.Println(p2)
    }
    if false {
        fmt.Println(p0, p0)
    }
    var v15 float64 = -2.25
    p0 = -816
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return p0
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 255
    v3 := make(map[int]bool)
    var v4 uint16 = 2
    v1 = "a"
    switch v1 {
    case "\n\t":
        v1 = "hello"
    }
    if true {
        var v16 uint32 = 4294967295
        var v17 string = ((v1 + (v1 + "\n\t")) + ("hello" + v1))
        _ = v16
        _ = v17
    }
    v3[127] = (!false)
    v2 += v2
    var v18 int8 = 0
    var v19 int16 = -1
    v20 := GenS2{F0: (65535 == 42)}
    fmt.Println(v2, v2)
    fmt.Println(v1, v18, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
