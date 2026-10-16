package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int64 = 42
    var v2 int64 = v1
    var v3 uint = 0
    v4 := make(map[string]uint32)
    var v5 uint16 = 493
    var v6 uint32 = uint32(v5)
    v3 = (v3 & v3)
    delete(v4, ("hello" + (("foo bar" + "hello") + "a")))
    if false {
        for fi7 := 0; fi7 < 14; fi7++ {
            if (v6 != (^(v6))) {
                var v8 uint32 = (v6 % (((v6 + (v6 >> (((v3 + v3)) & 63)))) | 1))
                _ = v8
            }
            fmt.Println(v1)
            v9 := []int{9223372036854775807, (fi7 + fi7), 9223372036854775807}
            _ = fi7
            _ = v9
        }
    }
    v10 := GenS1{F0: v5, F1: 69}
    if (true || (v6 <= uint32(v3))) {
        v1 -= 42
    }
    var v11 string = ("\n\t" + ((("a" + "foo bar") + "a") + "foo bar"))
    var v12 float64 = ((0.0 * 3.14159) * 0.1)
    var v13 int8 = -63
    v14 := GenS2{F0: (v6 >> ((v3) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
