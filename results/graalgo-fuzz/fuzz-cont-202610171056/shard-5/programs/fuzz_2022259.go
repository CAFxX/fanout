package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []string{"foo bar", ("a" + (("" + ("" + "\n\t")) + "a")), "hello"}
    var v2 int64 = 1
    var v3 int32 = 294
    var v4 int16 = 1000
    var v5 uint32 = 2
    var v6 float64 = -2.25
    if (32767 >= -210) {
        v2 -= (v2 << ((4294967295) & 63))
    }
    v6 = (1e10 + v6)
    var v7 uint64 = 384
    var v8 uint32 = v5
    v9 := GenS1{F0: v2}
    v10 := GenS2{F0: true}
    var v11 uint8 = uint8(v5)
    var v12 int64 = v2
    v13 := GenS2{F0: false}
    var v14 uint = 838
    v1 = append(v1, "foo bar")
    var v15 uint8 = (v11 | (v11 & (v11 & uint8(v14))))
    v8 = (^(v5))
    var v16 float64 = float64(((0.1 + 1.5) - 3.14159))
    if true {
        var v17 uint64 = (v7 + uint64(v15))
        if (v3 >= 2147483647) {
            var v18 uint64 = (v17 >> ((4294967295) & 63))
            _ = v18
        }
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
