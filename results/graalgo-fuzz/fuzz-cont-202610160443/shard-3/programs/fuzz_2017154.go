package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 127
    v3 := []int{}
    var v4 int8 = 42
    fmt.Println(v2, v4)
    fmt.Println(v2, v4)
    v5 := GenS1{F0: 65535}
    if true {
        switch 1 {
        case -301:
            for fi6 := 0; fi6 < 13; fi6++ {
                fmt.Println(v2, v4)
                _ = fi6
            }
        case -735:
            var v7 uint16 = 32767
            _ = v7
        case 0:
            v4 = 117
        default:
            v8 := GenS1{F0: 2}
            _ = v8
        }
        var v9 float32 = (1e10 - (float32(float64(0.0)) - 3.14159))
        _ = v9
    } else {
        for fi10 := 0; fi10 < 17; fi10++ {
            v4 += (v4 | (v4 * v4))
            _ = fi10
        }
    }
    var v11 int64 = -61
    _ = len(v3)
    v12 := GenS2{F0: 255, F1: 9223372036854775807}
    v11 -= -1
    if true {
        v2 += v2
        v13 := GenS1{F0: 322}
        switch "\n\t" {
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v13
    }
    if (0 <= 127) {
        var v14 int64 = v11
        switch (v2 >> ((454) & 63)) {
        case -27:
            for fi15 := 0; fi15 < 12; fi15++ {
                if (2 > 810) {
                    fmt.Println(fi15)
                    v16 := GenS1{F0: 279}
                    v17 := GenS2{F0: 85, F1: (v14 - v11)}
                    _ = v16
                    _ = v17
                }
                _ = fi15
            }
        case 9223372036854775807:
            if (121 > 125) {
                v4 = v4
                v18 := v1["a"]
                _ = v18
            } else {
                switch 42 {
                case -1:
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := GenS1{F0: 0}
                        fmt.Println(v11)
                        _ = fi19
                        _ = v20
                    }
                case -694:
                    v11 = (v11 * (v11 * v11))
                }
            }
        default:
            var v21 uint = 247
            _ = v21
        }
        _ = v14
    }
    var v22 int64 = (v11 * (v11 & -1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
