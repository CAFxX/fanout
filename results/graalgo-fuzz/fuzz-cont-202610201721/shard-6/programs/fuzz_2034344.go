package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 bool, p1 uint) uint {
    p0 = (-48 >= 101)
    p0 = (0.1 >= 0.1)
    v7 := GenS1{F0: 4294967295, F1: p0}
    switch -2147483648 {
    case 42:
        p1 -= p1
    case -9223372036854775808:
        v8 := []int{42, 4294967295, -1, 127}
        _ = v8
    case 127:
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS2{F0: 0}
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS3{F0: 913, F1: 0}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return (p1 << ((4294967295) & 63))
}

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 595
    v3 := make(map[uint32]bool)
    var v4 uint8 = 128
    var v5 bool = (float64(v1) >= 3.14159)
    v6 := make(map[string]string)
    if v5 {
        switch "a" {
        case "":
            v1 = ((-2.25 + v1) / float32(1.5))
        }
    }
    v12 := GenS4{F0: true, F1: (v4 | (^(v4)))}
    var v13 int = 1
    var v14 uint32 = 1
    v15 := GenS1{F0: 945, F1: ((v2 & (v2 * (v2 / ((726) | 1)))) > v2)}
    v1 = v1
    for fi16 := 0; fi16 < 9; fi16++ {
        delete(v6, "\n\t")
        _ = fi16
    }
    switch int(v14) {
    case 999:
        var v17 int16 = -573
        _ = v17
    case -2147483648:
        v13 = (v13 + 127)
    default:
        switch (v13 & v13) {
        case 127:
            v1 = (v1 + ((float32(0.1) + (1.5 * v1)) / (float32(1.5) / ((v1 - 0.0) * (3.14159 - v1)))))
        case -398:
            if true {
                v6[("" + ("hello" + "hello"))] = "foo bar"
                switch -293 {
                case 800:
                    fmt.Println(v2, v1, v1)
                default:
                    v18 := GenS1{F0: v14, F1: (1 == 127)}
                    _ = v18
                }
            } else {
                v4 -= v4
            }
        default:
            var v19 uint = 18446744073709551615
            _ = v19
        }
    }
    v3[v14] = v5
    switch v13 {
    case 0:
        v1 = (((v1 * float32((1.5 * 1.5))) + (((1.5 * (v1 / 0.1)) - float32(-2.25)) + v1)) + (v1 - v1))
    }
    v13 += -1
    v14 = 213
    fmt.Println(v14, v13)
    for fi20 := 0; fi20 < 14; fi20++ {
        v21 := GenS3{F0: (^(v2)), F1: 1}
        fmt.Println(v2, v13, v14)
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: 4294967295}
    var v23 string = "a"
    v6["hello"] = ("a" + v23)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
