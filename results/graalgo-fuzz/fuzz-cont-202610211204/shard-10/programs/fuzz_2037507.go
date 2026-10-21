package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []uint32{0, 615, 34, 2147483648}
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    var v4 int16 = 0
    var v5 int = 127
    v6 := []string{(("a" + "foo bar") + "hello")}
    var v7 int32 = -2147483648
    v8 := v2[4294967295]
    var v9 float64 = 3.14159
    v1 = append(v1, v8)
    if (!true) {
        v6 = append(v6, "")
        v5 += (v5 | -9223372036854775808)
    }
    var v10 int = (v5 | 127)
    v8 = v8
    if ((-127 > 1) && ((v8 / ((v8) | 1)) <= (v8 ^ 633))) {
        v11 := GenS1{F0: int64(v5)}
        var v12 int32 = -2147483648
        _ = v11
        _ = v12
    }
    if (!((false && true) && ((3.14159 / 3.14159) > (0.1 - 3.14159)))) {
        var v13 int32 = v7
        v8 = 65535
        _ = v13
    }
    _ = len(v2)
    v6 = append(v6, "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
