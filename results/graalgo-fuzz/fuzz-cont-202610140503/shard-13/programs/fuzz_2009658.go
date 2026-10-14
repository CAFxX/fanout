package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint64 = 18446744073709551615
    var v3 float64 = 1.5
    v4 := make(map[uint32]bool)
    var v5 int64 = v1
    _ = len(v4)
    var v6 float64 = v3
    v1 += (v5 % (((v1 * -9223372036854775808)) | 1))
    v3 += v3
    v3 -= v3
    switch "" {
    case "a":
        var v7 uint = 2
        _ = v7
    }
    v4[478] = false
    v1 = (^(v5))
    v1 -= (v1 ^ 42)
    v1 = (v5 * 1)
    fmt.Println(v3, v5, v6)
    fmt.Println(v5, v2)
    v6 = ((1e10 / v3) * float64(0.1))
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: v2}
        if (v6 > 1.5) {
            for fi10 := 0; fi10 < 18; fi10++ {
                v3 = v6
                var v11 int64 = v5
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 11; fi12++ {
                v3 = v6
                v3 = v6
                var v13 float32 = -2.25
                _ = fi12
                _ = v13
            }
        }
        v2 -= v2
        _ = fi8
        _ = v9
    }
    switch 1 {
    case -898:
        var v14 uint16 = 127
        _ = v14
    case 452:
        fmt.Println(v5, v2)
    default:
        v15 := GenS1{F0: (v2 | v2)}
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        if false {
            var v17 uint64 = v2
            v18 := GenS1{F0: (^(v2))}
            _ = v17
            _ = v18
        } else {
            v6 -= v6
            fmt.Println(v1)
            v1 = v5
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            if (true || false) {
                delete(v4, 0)
                if false {
                    switch ("foo bar" + "a") {
                    case "":
                        v20 := v4[2]
                        _ = v20
                    case "hello":
                        v21 := GenS2{F0: (v1 << ((uint(v2)) & 63)), F1: 490}
                        _ = v21
                    case "a":
                        v3 = ((v3 * float64(-2.25)) / 1.0)
                    default:
                        v22 := []int{fi19}
                        _ = v22
                    }
                    v6 -= ((v6 / float64(((float32(v6) / 1.5) * float32(-2.25)))) + v6)
                } else {
                    var v23 float32 = 0.0
                    _ = v23
                }
                fmt.Println(v6, v1)
            }
            if false {
                delete(v4, 42)
                v24 := []int{}
                var v25 string = "\n\t"
                _ = v24
                _ = v25
            } else {
                v1 += (v1 ^ v5)
            }
            _ = fi19
        }
        v6 = (3.14159 - v6)
        _ = fi16
    }
    fmt.Println(v3, v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
