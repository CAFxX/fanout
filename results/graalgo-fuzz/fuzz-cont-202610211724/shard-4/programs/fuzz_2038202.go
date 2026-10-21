package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string }

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 1
    var v3 uint8 = (v2 | (v2 * (^(v2))))
    var v4 int = 127
    v1 = (v1 + v1)
    if (992 < 407) {
        var v5 bool = true
        _ = v5
    }
    switch (v1 + v1) {
    case "a":
        v3 = v2
    case "\n\t":
        fmt.Println(v4, v3)
    }
    var v6 float64 = float64(0.0)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int32 = -145
        v9 := GenS1{F0: 87, F1: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int32 = 42
    v6 = v6
    v4 = v4
    switch (v4 * 9223372036854775807) {
    case 1:
        v2 = 42
    default:
        v1 = "foo bar"
    }
    var v11 float32 = float32(v6)
    switch v4 {
    case 42:
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                switch (v1 + "a") {
                case "a":
                    var v14 float64 = (v6 * 0.0)
                    _ = v14
                }
                var v15 string = v1
                _ = fi13
                _ = v15
            }
            v6 = float64(v11)
            v3 -= 181
            _ = fi12
        }
    default:
        if (1 <= v3) {
            if false {
                v16 := GenS2{F0: "\n\t"}
                if (v3 != v3) {
                    for fi17 := 0; fi17 < 17; fi17++ {
                        v18 := GenS1{F0: 255, F1: 0}
                        v2 -= v3
                        _ = fi17
                        _ = v18
                    }
                } else {
                    var v19 int16 = 1000
                    _ = v19
                }
                if ((^(v3)) < (v2 + v3)) {
                    var v20 int16 = 1000
                    _ = v20
                }
                _ = v16
            }
        } else {
            v2 += (v2 * (v3 & 255))
        }
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        fmt.Println(v10, fi21)
        _ = fi21
    }
    v3 += (v3 - (^(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
