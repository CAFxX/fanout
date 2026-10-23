package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    var v3 uint64 = 42
    var v4 string = ""
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: fi5}
        var v7 float64 = (0.1 / 3.14159)
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint32 = 430
    switch v4 {
    case "\n\t":
        v8 -= (v8 << ((1) & 63))
    case "hello":
        fmt.Println(v8, v4)
    case "a":
        var v9 int = 32767
        _ = v9
    default:
        switch v4 {
        case "foo bar":
            var v10 int8 = 0
            _ = v10
        default:
            v4 = "\n\t"
        }
    }
    var v11 uint32 = v8
    fmt.Println(v8, v3)
    var v12 uint = 65535
    var v13 string = (v4 + (v4 + ("" + ((v4 + v4) + "foo bar"))))
    v14 := GenS1{F0: 0}
    if (uint32(v3) >= 2147483647) {
        var v15 uint16 = 32767
        _ = v15
    } else {
        var v16 uint16 = 287
        v17 := []int{4294967295, 32768, -510, 736}
        _ = v16
        _ = v17
    }
    fmt.Println(v13, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
