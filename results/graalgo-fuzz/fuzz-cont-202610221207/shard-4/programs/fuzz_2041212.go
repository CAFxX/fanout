package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []int{1, 2147483647, 500}
    v3 := make(map[int]bool)
    var v4 uint16 = 65535
    v1 -= (v1 ^ v1)
    v4 -= 255
    _ = len(v3)
    if false {
        v4 -= v4
    } else {
        var v5 bool = ("hello" >= "a")
        v5 = (3.14159 < (0.1 + float64((float32(1e10) - 0.0))))
        _ = v5
    }
    var v6 uint16 = v4
    v6 = 32767
    var v7 uint8 = 127
    switch "hello" {
    case "hello":
        if true {
            var v8 int8 = 69
            var v9 int32 = 42
            _ = v8
            _ = v9
        } else {
            var v10 string = "foo bar"
            _ = v10
        }
    }
    v6 = 614
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := v3[fi11]
        _ = fi11
        _ = v12
    }
    v4 += v4
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 bool = (-128 > 1)
            var v16 int8 = 0
            switch ("foo bar" + "foo bar") {
            case "a":
                v16 = (v16 & v16)
            case "hello":
                if len(v2) > 0 {
                    v2[0] = fi14
                }
            default:
                var v17 int64 = -9223372036854775808
                _ = v17
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        var v18 float32 = float32((0.1 * (0.1 - -2.25)))
        _ = fi13
        _ = v18
    }
    v19 := GenS1{F0: 34}
    var v20 uint = 1
    v6 = (v6 | (v4 | (v6 % ((v4) | 1))))
    var v21 uint = v20
    var v22 int = -487
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
