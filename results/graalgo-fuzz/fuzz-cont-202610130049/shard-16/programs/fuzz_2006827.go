package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int8; F2 int16 }
type GenS4 struct { F0 bool; F1 int; F2 bool }
func genFn0() uint {
    var v4 int64 = 45
    if (((false || (false || ((128 < 127) || ("" <= "hello")))) || (-955 != 369)) || (138 <= 2)) {
        var v5 int64 = -923
        _ = v5
    }
    _ = v4
    return 197
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 65535
    var v3 uint32 = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: 18446744073709551615, F1: 1, F2: 0}
        _ = fi6
        _ = v7
    }
    v3 = v3
    v2 = (^(v2))
    v8 := GenS2{F0: 164}
    v9 := GenS3{F0: (3.14159 >= float64(0.0)), F1: 0, F2: -152}
    v3 = (^(v3))
    var v10 float32 = 1.5
    for fi11 := 0; fi11 < 18; fi11++ {
        if true {
            switch "hello" {
            case "hello":
                if (true || true) {
                    var v12 uint = 0
                    _ = v12
                }
            default:
                v10 += (0.0 / v10)
            }
            fmt.Println(v3, v3)
            var v13 int8 = 42
            _ = v13
        }
        _ = fi11
    }
    v14 := GenS1{F0: 1, F1: 4294967295, F2: (v2 & v2)}
    fmt.Println(v10, v3)
    switch ("" + "a") {
    case "a":
        switch "a" {
        case "hello":
            var v15 uint = 4294967295
            _ = v15
        case "foo bar":
            v16 := GenS4{F0: false, F1: 9223372036854775807, F2: (true || (true && true))}
            _ = v16
        default:
            v10 += ((float32(float64(v10)) / 3.14159) / (v10 + float32(1.5)))
        }
    case "\n\t":
        v17 := []int{}
        _ = v17
    }
    delete(v1, 1)
    v18 := []int{1, -9223372036854775808, 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
