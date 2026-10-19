package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 string; F1 int64; F2 int32 }
func genFn0(p0 int, p1 uint32) int32 {
    var v6 uint = 639
    v7 := GenS1{F0: -1, F1: 0, F2: (!(0.0 <= float64((3.14159 * 0.0))))}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 string = "foo bar"
    v2 := []string{}
    v3 := []uint32{468}
    v4 := []string{}
    var v5 int16 = -653
    switch (v1 + "\n\t") {
    case "foo bar":
        v8 := GenS1{F0: 63, F1: 33, F2: false}
        _ = v8
    case "\n\t":
        _ = len(v3)
    }
    v9 := GenS1{F0: -1, F1: 493, F2: false}
    if len(v4) > 0 {
        v4[0] = "a"
    }
    v1 = "\n\t"
    if (4294967295 == 805) {
        var v10 int8 = 61
        _ = v10
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS1{F0: 127, F1: 238, F2: false}
            v5 = v5
            _ = fi11
            _ = v12
        }
    }
    if (true && (!((v5 - 411) <= v5))) {
        _ = len(v3)
    } else {
        var v13 int = -9223372036854775808
        _ = v13
    }
    var v14 int8 = -1
    switch 127 {
    case 32767:
        v15 := GenS1{F0: v14, F1: 31, F2: false}
        _ = v15
    }
    v16 := GenS2{F0: v1, F1: 42, F2: genFn0(32767, 388)}
    switch v1 {
    case "a":
        fmt.Println(v5, v1)
    }
    v1 = v1
    fmt.Println(v14, v14, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
