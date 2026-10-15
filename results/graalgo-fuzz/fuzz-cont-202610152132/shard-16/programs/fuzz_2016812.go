package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0(p0 bool, p1 uint64) int64 {
    for fi5 := 0; fi5 < 18; fi5++ {
        if (p0 || (("hello" + "\n\t") != "")) {
            fmt.Println(p0)
        } else {
            p0 = (p0 || (p0 || false))
            var v6 uint = 93
            v6 = (v6 ^ 0)
            _ = v6
        }
        p0 = (9223372036854775807 >= 9223372036854775807)
        _ = fi5
    }
    p1 = p1
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = (0.1 + 1e10)
    var v3 bool = false
    var v4 int64 = 9223372036854775807
    fmt.Println(v3, v4, v3)
    var v7 float32 = v2
    fmt.Println(v2, v3)
    v8 := GenS1{F0: -766, F1: 0, F2: 128}
    fmt.Println(v7, v2)
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 float64 = (3.14159 * 3.14159)
        var v11 float32 = v7
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v7, v3, v4)
    if (864 <= 0) {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int64 = 238
            switch "foo bar" {
            case "\n\t":
                var v14 uint8 = 42
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        v15 := GenS1{F0: 4294967295, F1: 171, F2: 1}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS1{F0: (fi16 << ((42) & 63)), F1: 128, F2: 128}
            var v18 bool = v3
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v19 := []int{45, -1, -341}
    v1[4294967295] = ("a" + ("hello" + "a"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
