package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int }
func genFn0(p0 uint64) int16 {
    var v5 float32 = (0.1 * -2.25)
    var v6 int16 = 904
    _ = p0
    _ = v5
    _ = v6
    return v6
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    var v4 int = 607
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 0, F1: false, F2: (-(v4))}
        v1 = ""
        v9 := GenS1{F0: 4294967295, F1: false, F2: fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 302, F1: true, F2: v4}
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint32 = 123
        _ = fi11
        _ = v12
    }
    _ = len(v2)
    v1 = ("" + "a")
    if false {
        v2[v1] = 869
    }
    v1 = v1
    var v13 float64 = (1.5 / (1.5 * 3.14159))
    if false {
        v1 = ""
        for fi14 := 0; fi14 < 13; fi14++ {
            v4 = fi14
            var v15 int16 = -1
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 int = (v4 + -2147483648)
            var v18 float64 = v13
            v17 = -1
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        fmt.Println(v13, v1)
    }
    v19 := GenS1{F0: 4294967295, F1: false, F2: v4}
    var v20 uint64 = 98
    v21 := GenS1{F0: (v20 & v20), F1: ((42 < 65535) && false), F2: v4}
    v20 -= v20
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint = 65535
        _ = fi22
        _ = v23
    }
    fmt.Println(v13, v1, v1)
    v20 = (^(v20))
    v20 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
