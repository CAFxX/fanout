package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 string, p1 int64, p2 int64) uint16 {
    p2 = p2
    p2 = -1
    v5 := GenS1{F0: 1000}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    var v1 bool = (!(-725 <= 127))
    var v2 int16 = 1
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    v1 = v1
    v1 = v1
    v2 += 0
    if (994 > 4294967295) {
        var v6 uint16 = 764
        _ = v6
    }
    fmt.Println(v2)
    v1 = (true && true)
    fmt.Println(v2)
    var v7 int8 = 1
    var v8 int32 = 918
    fmt.Println(v1, v8, v7)
    var v9 uint64 = 85
    v10 := GenS1{F0: 1}
    v11 := GenS2{F0: "a"}
    var v12 bool = (2147483647 > -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
