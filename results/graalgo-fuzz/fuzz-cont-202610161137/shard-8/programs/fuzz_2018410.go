package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint8 = 116
    var v2 int64 = 1
    var v3 float64 = 1.5
    if (((2 == 65535) || (0 > 0)) || (false && false)) {
        v4 := []int{-2147483648, 1, 32767, -9223372036854775808}
        _ = v4
    }
    var v5 uint32 = 2147483647
    var v6 int64 = v2
    var v7 float64 = -2.25
    fmt.Println(v7)
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 = v7
        v1 = 128
        _ = fi8
    }
    var v9 int16 = 1
    var v10 uint64 = 4294967295
    switch "\n\t" {
    case "\n\t":
        if (-101 > (v9 + v9)) {
            v11 := GenS1{F0: v2}
            _ = v11
        } else {
            if true {
                if (!((v5 + (v5 - 2147483648)) != v5)) {
                    v12 := []int{1, 0, -9223372036854775808}
                    _ = v12
                } else {
                    var v13 uint8 = (^(v1))
                    v14 := GenS2{F0: 0}
                    _ = v13
                    _ = v14
                }
            } else {
                v15 := GenS2{F0: int32(v2)}
                _ = v15
            }
            fmt.Println(v5)
        }
    case "hello":
        fmt.Println(v2)
    default:
        v3 = float64(float32((float64(0.1) / (v7 - float64((-2.25 + -2.25))))))
    }
    fmt.Println(v9)
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(v5, v3, v7)
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int = 127
            var v19 float64 = v3
            var v20 float32 = 3.14159
            if (42 < 442) {
                var v21 uint = uint(v10)
                _ = v21
            } else {
                v6 += v6
                switch fi16 {
                case 179:
                    if ((((v19 / -2.25) + (float64(((0.1 - v20) - v20)) / (v3 * float64(3.14159)))) >= float64(float32(0.0))) || (2147483647 > 1)) {
                        fmt.Println(fi17, v6, v2)
                        var v22 int16 = 1
                        var v23 int32 = 0
                        _ = v22
                        _ = v23
                    }
                case -80:
                    v24 := GenS2{F0: int32(v10)}
                    _ = v24
                }
                v1 = v1
            }
            if true {
                v2 = (v6 / ((-1) | 1))
            }
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            v1 = (v1 + v1)
            _ = fi25
        }
        _ = fi16
    }
    v7 = ((v3 * v7) - 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
