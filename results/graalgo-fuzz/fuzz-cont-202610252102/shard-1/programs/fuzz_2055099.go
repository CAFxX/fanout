package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint32 = 42
    var v2 uint32 = (v1 * (^(v1)))
    var v3 int16 = 42
    v4 := make(map[uint32]string)
    var v5 float32 = (float32(((float64(0.0) - float64((0.1 + 1e10))) + (1e10 + -2.25))) * 1.5)
    var v6 int64 = 105
    switch int(v6) {
    case 128:
        v7 := GenS1{F0: 475}
        _ = v7
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            v6 = (v6 | v6)
            switch (("foo bar" + "foo bar") + "a") {
            case "hello":
                v5 = v5
            }
            if false {
                if (((v3 - (-(v3))) < (v3 / ((42) | 1))) || true) {
                    var v9 uint16 = 255
                    _ = v9
                } else {
                    fmt.Println(v1, v5)
                    v6 = (^(v6))
                }
                switch "hello" {
                case "foo bar":
                    var v10 int = fi8
                    _ = v10
                case "\n\t":
                    fmt.Println(fi8, v1)
                case "":
                    switch 535 {
                    case 0:
                        var v11 uint16 = 1
                        _ = v11
                    case 32768:
                        v6 -= (v6 & -573)
                    }
                default:
                    for fi12 := 0; fi12 < 10; fi12++ {
                        var v13 uint8 = 2
                        _ = fi12
                        _ = v13
                    }
                }
            }
            _ = fi8
        }
    }
    switch 2147483647 {
    case 1:
        var v14 int = 127
        _ = v14
    case 9223372036854775807:
        fmt.Println(v5)
    case 704:
        switch -1 {
        case 127:
            for fi15 := 0; fi15 < 14; fi15++ {
                delete(v4, (v2 + 436))
                var v16 float32 = (float32((1e10 / 1.0)) + (((v5 + v5) * v5) / 3.14159))
                _ = fi15
                _ = v16
            }
        case 806:
            var v17 int32 = 0
            _ = v17
        }
    }
    var v18 uint32 = v2
    v19 := GenS1{F0: 0}
    var v20 int32 = 0
    var v21 bool = true
    v22 := GenS1{F0: 255}
    fmt.Println(v18)
    var v23 uint8 = 127
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
