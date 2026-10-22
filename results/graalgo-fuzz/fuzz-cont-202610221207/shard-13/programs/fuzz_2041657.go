package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 255
    v3 := make(map[string]uint32)
    var v4 float64 = 3.14159
    var v5 int16 = 42
    var v6 int8 = 0
    v7 := GenS1{F0: -9223372036854775808}
    switch 32768 {
    case 222:
        v1["a"] = false
    case 32768:
        var v8 int16 = v5
        _ = v8
    case -2147483648:
        v5 = v5
    }
    v9 := []int{1, 128, -209, 0}
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(v6)
        _ = fi10
    }
    switch ("a" + "") {
    case "a":
        var v11 int32 = 1
        _ = v11
    }
    var v12 int8 = 0
    if false {
        var v13 float32 = -2.25
        if false {
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := GenS1{F0: -777}
                v2 = 32767
                v16 := GenS1{F0: 42}
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            v3["\n\t"] = 1
        }
        _ = v13
    } else {
        var v17 float32 = (3.14159 / -2.25)
        v18 := v3[("foo bar" + "a")]
        fmt.Println(v18)
        _ = v17
        _ = v18
    }
    v12 = 100
    switch 0 {
    case -1:
        v6 = (v12 - (v6 * (v12 % ((v12) | 1))))
    default:
        fmt.Println(v12, v4)
    }
    v2 -= (v2 ^ 65535)
    v19 := GenS1{F0: -9223372036854775808}
    v20 := GenS1{F0: -9223372036854775808}
    var v21 int32 = -474
    var v22 int8 = (v6 & v6)
    v23 := GenS2{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
