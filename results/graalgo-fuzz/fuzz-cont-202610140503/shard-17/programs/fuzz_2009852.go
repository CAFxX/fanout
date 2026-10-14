package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int8 {
    var v6 float64 = (-2.25 * (0.1 / 1e10))
    var v7 uint16 = 546
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    var v4 float64 = v1
    var v5 uint64 = 42
    var v8 uint32 = 2147483648
    v9 := GenS1{F0: 893}
    for fi10 := 0; fi10 < 15; fi10++ {
        v3 = "\n\t"
        fmt.Println(v5, v3, v3)
        var v11 uint8 = 255
        _ = fi10
        _ = v11
    }
    var v12 bool = false
    var v13 string = v3
    if v12 {
        var v14 int32 = 1
        _ = v14
    } else {
        if (((3.14159 + (-2.25 - float32(1e10))) <= 1e10) || v12) {
            fmt.Println(v8, v5)
            var v15 int64 = 1
            _ = v15
        } else {
            if (2147483648 <= v8) {
                delete(v2, "a")
            } else {
                _ = len(v2)
            }
            var v16 int32 = 138
            _ = v16
        }
        v17 := GenS1{F0: -1}
        switch v3 {
        case "hello":
            switch 128 {
            case 4294967295:
                v12 = false
            default:
                var v18 int64 = 456
                _ = v18
            }
        case "":
            switch 0 {
            case -972:
                v5 -= (^(v5))
            }
        default:
            var v19 int32 = 344
            _ = v19
        }
        _ = v17
    }
    fmt.Println(v5)
    var v20 int32 = 1
    v5 += (v5 & v5)
    fmt.Println(v12)
    switch -9223372036854775808 {
    case -9223372036854775808:
        v12 = (!(!(-2.25 >= (-2.25 / 1.0))))
    case 2147483647:
        for fi21 := 0; fi21 < 8; fi21++ {
            v20 -= v20
            v8 = (v8 ^ (v8 | v8))
            v22 := []int{}
            _ = fi21
            _ = v22
        }
    case 32768:
        var v23 uint64 = v5
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
