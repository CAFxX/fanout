package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int16 = -400
    var v2 int = 32768
    var v3 string = "foo bar"
    var v4 int32 = int32(v1)
    v5 := make(map[string]uint32)
    v6 := []int{-468}
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: 127}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v9 uint32 = 65535
    var v10 uint8 = 128
    v11 := GenS1{F0: 69}
    var v12 string = v3
    switch "\n\t" {
    case "a":
        v1 -= v1
    case "":
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint = 18446744073709551615
            fmt.Println(v4)
            var v15 int = v2
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v2 = (v2 >> ((0) & 63))
    var v16 int16 = (v1 << ((467) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
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
