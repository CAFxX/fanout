package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
func genFn0(p0 uint64, p1 int64, p2 int8) uint32 {
    var v7 uint8 = 42
    v8 := GenS1{F0: 2, F1: false, F2: v7}
    p2 += 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := []uint32{65535, 56}
    var v2 string = "a"
    var v3 int8 = 42
    var v4 uint64 = 1
    var v5 bool = (true || true)
    var v6 uint8 = 128
    fmt.Println(v2, v3)
    fmt.Println(v5, v3, v2)
    v6 = v6
    if (382 > 18446744073709551615) {
        v1 = append(v1, 65535)
        var v9 uint16 = 2
        _ = v9
    } else {
        if (!(267 <= -32768)) {
            var v10 uint32 = 113
            switch v2 {
            case "a":
                if len(v1) > 0 {
                    v1[0] = v10
                }
            }
            _ = v10
        }
    }
    var v11 uint16 = 65535
    for fi12 := 0; fi12 < 4; fi12++ {
        v4 -= (v4 & (v4 * (v4 + (v4 << ((707) & 63)))))
        _ = fi12
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v2 = v2
    v6 = 0
    v3 -= (^(v3))
    v4 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
