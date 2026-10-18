package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8, p1 int32, p2 uint32) uint64 {
    v6 := GenS1{F0: "", F1: 939, F2: 127}
    p1 = 0
    if ((-(p0)) <= p0) {
        var v7 uint32 = 2
        _ = v7
    }
    var v8 uint64 = 4294967295
    var v9 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 int = 42
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    fmt.Println(v3)
    fmt.Println(v3)
    var v10 int64 = 368
    v3 += (v3 << ((1) & 63))
    if false {
        fmt.Println(v10, v3)
        v11 := GenS1{F0: "foo bar", F1: 1000, F2: v3}
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 string = "foo bar"
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 int = v3
            var v16 int16 = 1000
            var v17 int64 = v10
            var v18 uint64 = 18446744073709551615
            if (true && true) {
                var v19 uint64 = 42
                fmt.Println(v16, v15)
                if (3.14159 <= -2.25) {
                    fmt.Println(v19, fi14, v15)
                    switch (fi14 * (fi14 * (-(v15)))) {
                    case -821:
                        var v20 int32 = int32(v19)
                        _ = v20
                    case 32767:
                        var v21 int = v15
                        _ = v21
                    case 32768:
                        _ = len(v4)
                    default:
                        v22 := v5["a"]
                        _ = v22
                    }
                    _ = len(v1)
                    v13 = v13
                }
                _ = v19
            } else {
                fmt.Println(fi14, v18, fi14)
                v2[v13] = (!(1e10 > -2.25))
            }
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi12
        _ = v13
    }
    v10 -= v10
    v3 += (v3 * (-(v3)))
    var v23 bool = true
    var v24 uint64 = 18446744073709551615
    v25 := []int{(v3 - 2147483647), (^(v3))}
    v3 = 42
    var v26 int32 = 0
    for fi27 := 0; fi27 < 6; fi27++ {
        v28 := GenS2{F0: 0}
        if (!v23) {
            v10 = (v10 >> ((404) & 63))
            v3 -= (fi27 - 2147483647)
            v3 = fi27
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            v10 = (v10 + (v10 & v10))
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    if v23 {
        v25 = append(v25, (^(v3)))
        var v30 string = "a"
        var v31 uint = 2
        _ = v30
        _ = v31
    } else {
        var v32 uint32 = 4294967295
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
