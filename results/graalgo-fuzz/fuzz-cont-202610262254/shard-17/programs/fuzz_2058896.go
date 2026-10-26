package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 43
    v3 := make(map[uint32]string)
    var v4 int64 = 9223372036854775807
    var v5 uint64 = 0
    var v6 uint8 = 2
    switch 0 {
    case 32767:
        if (2147483647 == 127) {
            v7 := GenS1{F0: -1}
            _ = v7
        }
    case 0:
        v5 -= uint64(v2)
    case 4294967295:
        v5 += v5
    }
    v5 += (v5 & (v5 * v5))
    v8 := []int{127, -417, 0, 9223372036854775807}
    _ = len(v8)
    v6 = v6
    fmt.Println(v6, v6, v6)
    switch ("\n\t" + "a") {
    case "a":
        var v9 int = 32768
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int64 = (v4 * -9223372036854775808)
            if len(v8) > 0 {
                v8[0] = fi10
            }
            switch (fi10 * 411) {
            case 32767:
                v8 = append(v8, (-(fi10)))
            case 127:
                v12 := GenS1{F0: 0}
                _ = v12
            case -1:
                var v13 int64 = v11
                _ = v13
            default:
                delete(v1, "\n\t")
            }
            _ = fi10
            _ = v11
        }
    default:
        var v14 int = 4294967295
        _ = v14
    }
    if ((!false) || false) {
        v6 = v6
        v15 := GenS2{F0: v5}
        fmt.Println(v4, v6)
        _ = v15
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v2, v2)
    }
    delete(v3, 4294967295)
    v5 += v5
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
