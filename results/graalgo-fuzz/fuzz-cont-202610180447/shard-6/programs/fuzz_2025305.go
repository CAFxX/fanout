package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64 }
func genFn0() bool {
    var v6 uint16 = 0
    v7 := []int{2147483647, 32768, -361}
    var v8 int16 = -1
    _ = v6
    _ = v7
    _ = v8
    return (1 != 19)
}

func main() {
    var v1 string = "hello"
    var v2 int32 = -1
    var v3 float32 = (1e10 * 1e10)
    var v4 bool = false
    var v5 uint32 = 1
    var v9 float32 = v3
    fmt.Println(v2, v4)
    switch -9223372036854775808 {
    case 32768:
        for fi10 := 0; fi10 < 7; fi10++ {
            v1 = v1
            switch v1 {
            case "hello":
                if (-128 <= 85) {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        v4 = true
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 3; fi12++ {
                        var v13 uint16 = 1
                        fmt.Println(v4)
                        v1 = v1
                        _ = fi12
                        _ = v13
                    }
                    var v14 uint = 42
                    _ = v14
                }
            case "":
                v15 := []int{(fi10 | 128), (fi10 % (((fi10 >> ((42) & 63))) | 1)), (^(fi10)), 182}
                _ = v15
            default:
                v16 := GenS1{F0: ("hello" + v1), F1: (v5 >> ((4294967295) & 63)), F2: 127}
                _ = v16
            }
            _ = fi10
        }
    case -9223372036854775808:
        v17 := GenS1{F0: (v1 + "foo bar"), F1: 2147483647, F2: 255}
        _ = v17
    default:
        v2 -= v2
    }
    var v18 uint64 = 18446744073709551615
    var v19 int8 = -61
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 float64 = float64(float32((1.5 + 0.0)))
            _ = fi21
            _ = v22
        }
        v23 := GenS1{F0: ((v1 + v1) + "a"), F1: v5, F2: 207}
        _ = fi20
        _ = v23
    }
    fmt.Println(v1)
    v18 -= v18
    v3 = 1e10
    v24 := GenS2{F0: -340}
    v25 := GenS2{F0: -697}
    var v26 float32 = float32(1e10)
    v5 += v5
    v27 := GenS2{F0: 0}
    v28 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
