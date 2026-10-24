package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint16 = 255
    v2 := []int{128}
    var v3 int8 = 1
    var v4 uint32 = 881
    v2 = append(v2, 42)
    var v5 int32 = 1
    var v6 uint16 = (v1 & 2)
    switch int(v5) {
    case 42:
        fmt.Println(v4, v5, v3)
    }
    fmt.Println(v5)
    var v7 int = -2147483648
    if (!false) {
        if false {
            var v8 int32 = v5
            var v9 bool = ((!((-161 <= (v7 ^ v7)) || ((!true) || true))) || false)
            _ = v8
            _ = v9
        }
        var v10 int32 = v5
        _ = v10
    }
    if false {
        v2 = append(v2, v7)
    }
    var v11 uint32 = (v4 ^ (v4 & (v4 + 4294967295)))
    for fi12 := 0; fi12 < 3; fi12++ {
        _ = len(v2)
        v13 := GenS1{F0: -32768}
        _ = fi12
        _ = v13
    }
    var v14 uint16 = (v1 / ((32767) | 1))
    if len(v2) > 0 {
        v2[0] = v7
    }
    var v15 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
