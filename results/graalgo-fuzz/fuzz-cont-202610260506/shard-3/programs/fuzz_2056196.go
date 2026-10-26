package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint32, p1 bool) int64 {
    v4 := []int{}
    fmt.Println(p1)
    var v5 int64 = -1
    if (2147483647 > 42) {
        p1 = p1
        p1 = p1
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            p0 += p0
        case "foo bar":
            if p1 {
                fmt.Println(v5)
            }
        default:
            _ = len(v4)
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -347
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]uint32)
    var v3 uint32 = 35
    var v6 uint32 = v3
    var v7 int16 = -671
    var v8 int16 = -1
    v9 := GenS1{F0: 2}
    v6 = v3
    var v10 string = ("" + ("foo bar" + ""))
    v11 := GenS1{F0: 178}
    var v12 float64 = (1.5 * 0.0)
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS1{F0: 128}
            switch (fi13 - (fi13 - -2147483648)) {
            case 9223372036854775807:
                v10 = "hello"
            default:
                v6 += 2
            }
            _ = fi14
            _ = v15
        }
        v12 += v12
        v2[(v10 + "\n\t")] = uint32(v1)
        _ = fi13
    }
    if true {
        var v16 int32 = 2147483647
        var v17 int64 = -1
        fmt.Println(v16, v10, v17)
        _ = v16
        _ = v17
    } else {
        v18 := GenS1{F0: 255}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
