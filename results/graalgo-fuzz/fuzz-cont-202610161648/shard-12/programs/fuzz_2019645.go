package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint16 {
    var v4 int16 = 1
    var v5 uint16 = 42
    v6 := []int{9223372036854775807}
    var v7 string = "a"
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 770
    v3 := []uint32{}
    v2 -= (^(v2))
    var v8 int16 = 0
    var v9 uint16 = v2
    v10 := GenS1{F0: 864}
    var v11 uint64 = 841
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    if ((0.0 * 1e10) > float32(v1)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            v8 -= v8
            var v13 int = (fi12 | 127)
            if (genFn0() == v9) {
                var v14 int8 = 87
                var v15 int = -2147483648
                _ = v14
                _ = v15
            } else {
                v13 += -697
            }
            _ = fi12
            _ = v13
        }
        v1 = v1
        v3 = append(v3, 42)
    }
    var v16 int32 = 42
    switch "\n\t" {
    case "":
        if len(v3) > 0 {
            v3[0] = uint32(v2)
        }
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
    case "a":
        var v17 int16 = v8
        _ = v17
    default:
        var v18 bool = false
        _ = v18
    }
    v11 += 291
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
