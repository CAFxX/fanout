package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32, p1 int, p2 uint32) int16 {
    p1 += (p1 * p1)
    switch (p1 - (p1 % (((^(p1))) | 1))) {
    case 93:
        var v6 float32 = 1e10
        _ = v6
    case 9223372036854775807:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    return int16(p1)
}

func main() {
    var v1 uint8 = 1
    v2 := []int{-9223372036854775808, 128, 32768}
    var v3 float64 = 3.14159
    var v4 float64 = 0.1
    var v5 uint16 = 65535
    _ = len(v2)
    if true {
        fmt.Println(v4)
    }
    var v7 uint8 = (v1 ^ 2)
    fmt.Println(v3, v4, v3)
    var v8 uint64 = 1
    switch 2147483647 {
    case -860:
        v5 -= v5
    case -2147483648:
        if ("a" == ("a" + ("foo bar" + ("hello" + "foo bar")))) {
            var v9 int8 = 100
            _ = v9
        }
    case -119:
        if true {
            fmt.Println(v4, v5)
            var v10 int = int(v7)
            _ = v10
        } else {
            _ = len(v2)
            if (int16(v5) <= 0) {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v3 = (float64((((0.0 * -2.25) / float32(v3)) - ((float32((v3 * (v4 + 1.5))) * 1e10) - 3.14159))) / 3.14159)
                    _ = fi11
                }
                if false {
                    var v12 int8 = 42
                    _ = v12
                }
                for fi13 := 0; fi13 < 2; fi13++ {
                    _ = len(v2)
                    v8 = 86
                    v4 = v4
                    _ = fi13
                }
            } else {
                _ = len(v2)
                var v14 int32 = 0
                v15 := GenS1{F0: -32768, F1: int8(v8)}
                _ = v14
                _ = v15
            }
        }
    }
    var v16 uint16 = (^(v5))
    if true {
        v4 -= 3.14159
    }
    v5 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
