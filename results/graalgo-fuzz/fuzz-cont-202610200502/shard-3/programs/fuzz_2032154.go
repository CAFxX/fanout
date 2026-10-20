package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint }
func genFn0(p0 int, p1 uint, p2 uint8) int {
    v4 := []int{(^(p0)), -448, p0}
    var v5 string = ("foo bar" + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p0 % ((-2147483648) | 1))
}

func main() {
    var v1 float32 = 0.1
    var v2 int16 = -32768
    v3 := make(map[uint32]uint32)
    if true {
        v6 := GenS1{F0: 128, F1: 128}
        _ = v6
    } else {
        delete(v3, 2147483648)
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := v3[65535]
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(v1, v1)
    var v9 uint16 = 255
    v10 := GenS1{F0: 181, F1: 1}
    var v11 float32 = -2.25
    v1 += (0.1 - v1)
    fmt.Println(v11, v9, v2)
    v12 := GenS2{F0: 686, F1: 1, F2: 100}
    var v13 uint64 = 1
    fmt.Println(v1, v2)
    switch 127 {
    case 32768:
        for fi14 := 0; fi14 < 8; fi14++ {
            delete(v3, 0)
            _ = fi14
        }
    case 42:
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS3{F0: v9, F1: 255, F2: 0}
            v17 := GenS3{F0: (v9 - (v9 / ((v9) | 1))), F1: 1, F2: -2147483648}
            for fi18 := 0; fi18 < 14; fi18++ {
                switch "foo bar" {
                case "hello":
                    v1 += -2.25
                case "a":
                    var v19 uint = 18446744073709551615
                    _ = v19
                case "\n\t":
                    v9 = 718
                default:
                    v13 -= 2
                }
                v13 -= v13
                v13 = (^(v13))
                _ = fi18
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    case 32767:
        v20 := []int{genFn0(genFn0(4294967295, 18446744073709551615, 127), 255, 1), -9223372036854775808, 2147483647}
        _ = v20
    }
    fmt.Println(v9, v2)
    var v21 int64 = 1
    var v22 uint64 = (^(v13))
    var v23 uint64 = v22
    v11 += 3.14159
    v9 += (v9 % (((v9 >> ((766) & 63))) | 1))
    switch "hello" {
    case "hello":
        for fi24 := 0; fi24 < 15; fi24++ {
            v23 += v13
            v25 := GenS4{F0: 191, F1: 582}
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v9
}
