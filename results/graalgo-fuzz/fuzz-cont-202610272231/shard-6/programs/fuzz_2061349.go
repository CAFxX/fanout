package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint8) uint32 {
    v5 := []int{-1, -1, -9223372036854775808, 670}
    for fi6 := 0; fi6 < 20; fi6++ {
        _ = len(v5)
        p0 -= 128
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    var v8 float32 = 1.5
    _ = p0
    _ = v5
    _ = v8
    return 4294967295
}

func main() {
    v1 := []string{("hello" + "a"), "a", ((("a" + "foo bar") + "foo bar") + ((("a" + "") + "hello") + "a"))}
    var v2 string = "\n\t"
    var v3 int8 = 100
    var v4 int64 = 925
    var v9 uint = 4294967295
    fmt.Println(v3, v4)
    fmt.Println(v4, v3)
    var v10 float32 = 3.14159
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(v3)
        v9 = (v9 >> ((v9) & 63))
        _ = fi11
    }
    v3 = 0
    for fi12 := 0; fi12 < 15; fi12++ {
        switch (fi12 % ((fi12) | 1)) {
        case 1:
            v13 := GenS1{F0: 128}
            _ = v13
        case 128:
            v14 := GenS2{F0: v3}
            _ = v14
        case 4294967295:
            v3 -= v3
        }
        _ = fi12
    }
    if (!true) {
        var v15 int64 = (v4 ^ (v4 & -984))
        _ = v15
    }
    v16 := GenS2{F0: 127}
    fmt.Println(v2, v3)
    v10 += float32(1e10)
    v4 = v4
    var v17 int32 = 279
    v18 := GenS2{F0: (v3 << ((v9) & 63))}
    var v19 bool = (v3 < (^(v3)))
    var v20 string = v2
    v21 := GenS1{F0: 1}
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 float64 = 0.1
        switch "" {
        case "a":
            var v24 string = v2
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
