package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 int16 = 1000
    var v5 float64 = 1e10
    if (false || true) {
        delete(v3, 4294967295)
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 int32 = 2147483647
            if (4294967295 < 926) {
                v5 = float64(-2.25)
                v4 = 1000
            }
            v7 = 1
            _ = fi6
            _ = v7
        }
    } else {
        v8 := v2[-359]
        v1 -= 578
        _ = v8
    }
    if true {
        v5 -= v5
        v5 = 3.14159
    }
    switch (("\n\t" + ("hello" + "a")) + (("" + ("a" + "\n\t")) + "\n\t")) {
    case "a":
        var v9 float32 = (3.14159 - 3.14159)
        _ = v9
    case "":
        var v10 int64 = 0
        _ = v10
    default:
        v11 := GenS1{F0: 42}
        _ = v11
    }
    v5 = 1e10
    fmt.Println(v1, v1)
    v5 += 1e10
    v12 := []int{}
    _ = len(v12)
    for fi13 := 0; fi13 < 3; fi13++ {
        if false {
            if true {
                v14 := GenS1{F0: -300}
                _ = v14
            }
        } else {
            v5 += (v5 * -2.25)
        }
        _ = fi13
    }
    if (1e10 > float64(1.5)) {
        for fi15 := 0; fi15 < 8; fi15++ {
            fmt.Println(fi15, fi15, v5)
            _ = fi15
        }
        var v16 uint32 = (v1 + (v1 ^ v1))
        _ = v16
    } else {
        _ = len(v2)
    }
    var v17 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
