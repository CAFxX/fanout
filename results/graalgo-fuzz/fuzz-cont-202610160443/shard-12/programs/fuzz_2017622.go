package main

import (
    "fmt"
)

func genFn0() string {
    v6 := []int{}
    var v7 uint64 = 18446744073709551615
    var v8 string = "foo bar"
    _ = v6
    _ = v7
    _ = v8
    return (v8 + v8)
}

func main() {
    v1 := []int{-874}
    v2 := []uint32{1, 1, 2}
    var v3 uint32 = 4294967295
    var v4 uint8 = 200
    v5 := make(map[uint32]string)
    if ((1.5 - -2.25) == 0.0) {
        var v9 int64 = -9223372036854775808
        _ = v9
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            fmt.Println(v4, v3)
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 int = fi11
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        v4 += 127
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 float64 = (3.14159 - 1e10)
        _ = fi13
        _ = v14
    }
    v4 -= v4
    v4 = 241
    if (("\n\t" + "hello") > "\n\t") {
        var v15 uint64 = 301
        v15 = (v15 * 0)
        _ = v15
    } else {
        v3 = v3
    }
    if true {
        fmt.Println(v4, v3)
    } else {
        v4 = (v4 + (v4 << ((0) & 63)))
        if true {
            var v16 int = 1
            _ = v16
        }
        fmt.Println(v4, v4)
    }
    v4 -= (v4 << ((65535) & 63))
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int64 = 42
        _ = fi17
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
