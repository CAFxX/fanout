package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8, p1 int32, p2 int16) uint64 {
    var v5 float64 = 3.14159
    var v6 int32 = (p1 ^ -2147483648)
    v7 := GenS1{F0: 127, F1: 395}
    v5 = ((((((((v5 / v5) - float64((3.14159 - 1e10))) / v5) / 1.0) - v5) + -2.25) - 3.14159) - 3.14159)
    switch 87 {
    case 0:
        var v8 bool = (594 <= -595)
        _ = v8
    case 236:
        p2 += p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]string)
    var v3 int8 = -128
    v4 := make(map[int]string)
    var v9 int32 = -1
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            v9 += (-(v9))
            _ = fi11
        }
        _ = fi10
    }
    v9 = v9
    v12 := GenS1{F0: 2, F1: 0}
    v13 := GenS1{F0: 0, F1: 872}
    for fi14 := 0; fi14 < 10; fi14++ {
        v3 = -40
        var v15 uint = 42
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v9 = v9
        _ = fi16
    }
    var v17 uint16 = 435
    var v18 bool = false
    switch "" {
    case "foo bar":
        v19 := GenS2{F0: (2147483647 == 4294967295)}
        _ = v19
    case "hello":
        switch 4294967295 {
        case -2147483648:
            for fi20 := 0; fi20 < 1; fi20++ {
                switch int(v1) {
                case 734:
                    var v21 int32 = 0
                    _ = v21
                case -512:
                    v1 = 42
                case 971:
                    var v22 float64 = (0.1 * 0.1)
                    _ = v22
                }
                var v23 int16 = -32768
                _ = fi20
                _ = v23
            }
        case 1:
            var v24 int64 = (v1 | -939)
            _ = v24
        case 9223372036854775807:
            var v25 string = "a"
            _ = v25
        default:
            var v26 uint16 = v17
            _ = v26
        }
    default:
        v1 -= -947
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := []int{int(v1), fi27}
        _ = fi27
        _ = v28
    }
    v1 = (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
