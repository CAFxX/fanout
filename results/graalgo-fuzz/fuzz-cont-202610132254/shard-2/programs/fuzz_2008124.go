package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0() uint8 {
    var v5 bool = false
    v5 = (("a" > "hello") || (3.14159 != 3.14159))
    var v6 uint16 = 530
    var v7 bool = v5
    _ = v5
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]string)
    v3 := []int{(v1 * v1), v1, (^(v1))}
    var v4 int16 = -1
    v8 := v2[0]
    var v9 int64 = 42
    switch 0 {
    case 32767:
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int8 = 127
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    case 467:
        v13 := GenS1{F0: -574, F1: (v4 + v4)}
        _ = v13
    case 426:
        var v14 int = -9223372036854775808
        _ = v14
    default:
        _ = len(v3)
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int = v1
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        v2[1] = (v8 + "foo bar")
        _ = fi17
    }
    var v18 uint32 = 42
    v19 := GenS1{F0: v9, F1: (v4 - v4)}
    v2[2] = (((v8 + v8) + "\n\t") + ("foo bar" + ("foo bar" + "foo bar")))
    var v20 float32 = (3.14159 / -2.25)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
