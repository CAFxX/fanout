package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int) uint8 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint64 = 82
        _ = fi5
        _ = v6
    }
    _ = p0
    return 255
}

func main() {
    var v1 int64 = 358
    var v2 uint8 = 128
    var v3 int64 = (v1 ^ 590)
    var v4 int8 = -1
    v4 = v4
    var v7 int64 = (v1 * v3)
    v8 := []int{4294967295, 32767}
    var v9 int8 = v4
    v8 = append(v8, -105)
    fmt.Println(v4, v9)
    if len(v8) > 0 {
        v8[0] = 482
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float32 = 1.5
            _ = fi11
            _ = v12
        }
        v2 = (v2 - 1)
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi10
    }
    if false {
        if false {
            fmt.Println(v3, v3, v3)
        } else {
            v2 += 128
            var v13 uint64 = 0
            switch "a" {
            case "\n\t":
                v4 = -58
            }
            _ = v13
        }
    }
    v14 := GenS1{F0: v2}
    if (1 != 42) {
        v4 += (-(v9))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
