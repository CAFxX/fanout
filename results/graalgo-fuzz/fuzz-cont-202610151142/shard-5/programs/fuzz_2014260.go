package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int, p1 int32) uint64 {
    p1 -= p1
    v7 := []int{p0, p0}
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    var v4 int8 = -19
    var v5 int32 = 1
    var v6 string = "foo bar"
    var v8 bool = true
    var v9 uint32 = 2147483648
    v10 := GenS1{F0: (v5 << ((1) & 63))}
    switch (v6 + "foo bar") {
    case "hello":
        for fi11 := 0; fi11 < 13; fi11++ {
            v6 = v6
            var v12 string = (v6 + (v6 + ""))
            _ = fi11
            _ = v12
        }
    }
    if (1 > 1) {
        v9 = 4294967295
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int64 = 367
        switch (^(fi13)) {
        case 0:
            v14 -= (v14 * 9223372036854775807)
        case -2147483648:
            fmt.Println(v14)
        default:
            var v15 string = (v6 + "")
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v16 int16 = 650
    v8 = false
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
