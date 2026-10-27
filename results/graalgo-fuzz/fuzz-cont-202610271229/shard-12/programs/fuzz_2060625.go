package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 int, p1 uint64, p2 int64) string {
    var v5 int = (p0 + p0)
    var v6 uint16 = 0
    switch ("foo bar" + "a") {
    case "a":
        v7 := GenS1{F0: 127, F1: "hello"}
        _ = v7
    case "":
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 float64 = -2.25
            p0 += (p0 % (((-(fi8))) | 1))
            _ = fi8
            _ = v9
        }
    case "foo bar":
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int64 = p2
            v12 := []int{}
            var v13 int16 = 32767
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    p0 = v5
    v6 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ("a" + "foo bar")
}
func genFn1() uint {
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 uint = 83
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS1{F0: 42, F1: "\n\t"}
        _ = fi16
        _ = v17
    }
    v18 := []int{747, -2147483648, -753}
    _ = v18
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{822, 638, 127, 32767}
    var v3 int64 = 1
    var v4 float32 = 1e10
    var v19 int64 = (v3 / ((920) | 1))
    var v20 float64 = 1.5
    fmt.Println(v4, v20)
    if (true || true) {
        v3 = (^(v3))
        for fi21 := 0; fi21 < 9; fi21++ {
            fmt.Println(v4, v3, v19)
            v19 += (v19 ^ (v3 * 0))
            _ = fi21
        }
        v3 -= (v19 + (v3 & -9223372036854775808))
    }
    v4 += v4
    if false {
        v19 = -1
        v19 -= v19
    } else {
        fmt.Println(v20)
        v19 += (v3 & (v19 + 42))
        v22 := v1[-1]
        _ = v22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        switch (fi23 + 32768) {
        case 789:
            var v24 int = 32768
            _ = v24
        case -100:
            var v25 int8 = 127
            _ = v25
        }
        fmt.Println(v20)
        _ = fi23
    }
    switch -765 {
    case 9223372036854775807:
        var v26 int32 = -447
        _ = v26
    case 131:
        v4 -= (1e10 / v4)
    }
    v4 += v4
    v19 = (v3 << ((genFn1()) & 63))
    v19 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
