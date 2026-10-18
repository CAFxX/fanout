package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint; F2 int16 }

func main() {
    var v1 string = ("\n\t" + "\n\t")
    v2 := []string{}
    var v3 int32 = 0
    var v4 uint16 = 255
    fmt.Println(v3)
    v3 += (^(v3))
    var v5 uint = 0
    v6 := GenS1{F0: v3, F1: 65535}
    var v7 int = 53
    _ = len(v2)
    var v8 float64 = float64(float32(1e10))
    v9 := GenS1{F0: 2147483647, F1: v4}
    var v10 int16 = -32768
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = v1
    }
    v11 := GenS1{F0: v3, F1: v4}
    if false {
        v8 += float64(3.14159)
        switch "" {
        case "hello":
            v7 -= (v7 ^ v7)
        }
    } else {
        if true {
            if len(v2) > 0 {
                v2[0] = ((v1 + (v1 + v1)) + v1)
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                for fi13 := 0; fi13 < 10; fi13++ {
                    v4 = v4
                    if ((-2.25 * float32((float64(0.1) - v8))) == (float32(1.5) - (-2.25 / ((0.0 / float32(v8)) - float32(1e10))))) {
                        var v14 bool = true
                        _ = v14
                    } else {
                        switch v1 {
                        case "\n\t":
                            var v15 int8 = -128
                            _ = v15
                        case "foo bar":
                            var v16 float32 = -2.25
                            _ = v16
                        default:
                            v8 -= 1.5
                        }
                    }
                    _ = fi13
                }
                _ = fi12
            }
            v4 += v4
        }
        var v17 int32 = v3
        v2 = append(v2, "a")
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 bool = false
        v20 := GenS2{F0: 127, F1: v5, F2: v10}
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: 42, F1: v5, F2: (v10 << ((v5) & 63))}
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
