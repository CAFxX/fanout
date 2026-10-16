package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint64, p1 string) uint32 {
    fmt.Println(p0)
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 float32 = (1e10 / 1.5)
    var v2 int = -589
    var v3 int64 = -9223372036854775808
    var v4 uint16 = 210
    var v5 float64 = 3.14159
    var v6 int16 = 1
    v7 := GenS1{F0: 2, F1: v3}
    v8 := GenS1{F0: 212, F1: v3}
    v5 = v5
    var v9 float32 = (0.1 / v1)
    switch "\n\t" {
    case "a":
        v9 = -2.25
    default:
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint8 = 38
            v12 := GenS1{F0: v11, F1: 9223372036854775807}
            if (false && (!(1e10 != float64(0.0)))) {
                fmt.Println(v5)
            } else {
                var v13 int16 = (^(v6))
                v5 = v5
                _ = v13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v14 int = 4294967295
    fmt.Println(v9)
    var v15 float32 = 0.0
    if true {
        v9 += 0.1
    }
    if false {
        var v16 int32 = 156
        switch "" {
        case "\n\t":
            fmt.Println(v6, v5, v14)
        case "":
            v17 := []int{127, 32767}
            _ = v17
        case "foo bar":
            if (!(609 > 0)) {
                v4 = 31
                v18 := GenS1{F0: 2, F1: (v3 - (v3 % ((-1) | 1)))}
                switch v2 {
                case 42:
                    if (932 <= 255) {
                        v2 = 109
                    } else {
                        if ((v16 / ((1) | 1)) > v16) {
                            var v19 uint64 = 65
                            var v20 int = v14
                            _ = v19
                            _ = v20
                        }
                    }
                }
                _ = v18
            } else {
                switch -9223372036854775808 {
                case -639:
                    var v21 uint = 65535
                    _ = v21
                }
                v4 = v4
            }
        }
        var v22 int = v2
        _ = v16
        _ = v22
    }
    fmt.Println(v1, v15, v15)
    v3 = (-(v3))
    if true {
        fmt.Println(v14, v15, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
