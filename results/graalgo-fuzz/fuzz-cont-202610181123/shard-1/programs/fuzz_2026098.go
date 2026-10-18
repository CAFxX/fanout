package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int16 = -5
    var v2 uint = 255
    var v3 int32 = 42
    var v4 uint = v2
    var v5 float64 = (0.1 - (-2.25 - (1.5 + 0.0)))
    v3 = (-(v3))
    var v6 uint = v2
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = v1
        var v9 int16 = -122
        _ = fi7
        _ = v8
        _ = v9
    }
    if (v2 == v6) {
        for fi10 := 0; fi10 < 1; fi10++ {
            if false {
                if (v3 < 42) {
                    v11 := GenS1{F0: -1, F1: 127, F2: fi10}
                    var v12 uint64 = 18446744073709551615
                    _ = v11
                    _ = v12
                }
            } else {
                v4 += 0
            }
            _ = fi10
        }
        var v13 string = ((("foo bar" + ("a" + "")) + "") + ("hello" + "foo bar"))
        _ = v13
    }
    v4 = v6
    v3 = (v3 / ((int32(v4)) | 1))
    if (0 != 918) {
        v5 = -2.25
        v14 := GenS2{F0: 869}
        _ = v14
    }
    v3 -= (v3 ^ v3)
    var v15 int64 = -353
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
