package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int64, p1 string) int {
    if false {
        fmt.Println(p1, p1)
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := []int{-1, fi6}
            v8 := GenS1{F0: 65535}
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    var v9 int64 = p0
    for fi10 := 0; fi10 < 13; fi10++ {
        p1 = p1
        v9 += v9
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 335
    var v3 int16 = 0
    v4 := make(map[int]int)
    var v5 uint64 = 1
    v2 += v2
    v1[("" + (("" + "\n\t") + "foo bar"))] = 72
    v5 += v5
    if (-22 >= 1) {
        if (0.1 <= (3.14159 / ((1e10 / (float32(0.0) + (float32(1.5) / (-2.25 * 1.5)))) + 1.5))) {
            fmt.Println(v2, v5)
        }
        var v11 uint16 = 0
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int = (fi12 - (fi12 % (((fi12 >> ((0) & 63))) | 1)))
            v14 := GenS1{F0: 42}
            v3 += v3
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    var v15 string = ("hello" + "hello")
    v2 = -9223372036854775808
    v15 = ("\n\t" + v15)
    for fi16 := 0; fi16 < 6; fi16++ {
        v5 = (v5 % ((v5) | 1))
        fmt.Println(v15, fi16, v3)
        var v17 int8 = 70
        _ = fi16
        _ = v17
    }
    switch 383 {
    case 32768:
        if false {
            v18 := GenS1{F0: 708}
            _ = v18
        }
    }
    var v19 uint32 = 139
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
