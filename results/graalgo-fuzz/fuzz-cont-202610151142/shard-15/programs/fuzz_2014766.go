package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS5 struct { F0 string; F1 uint16; F2 uint32 }
func genFn0() uint32 {
    switch "\n\t" {
    case "\n\t":
        for fi4 := 0; fi4 < 6; fi4++ {
            var v5 int16 = 1
            var v6 float64 = float64((((((1.5 * 1.5) / 1.0) + (0.1 / (-2.25 - 1.5))) / 3.14159) - (3.14159 + 1.5)))
            var v7 string = ("" + "a")
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    case "hello":
        var v8 bool = ((!true) || (512 < 2147483647))
        _ = v8
    }
    var v9 int16 = -204
    switch 1 {
    case 4294967295:
        var v10 int8 = 0
        _ = v10
    default:
        v11 := GenS1{F0: 1, F1: 1, F2: -1}
        _ = v11
    }
    _ = v9
    return 2147483648
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    v3[1] = (-292 > -1)
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS2{F0: ("" + ""), F1: 2}
        v1 -= -2.25
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 2
    v15 := GenS3{F0: 100, F1: -32768, F2: 2}
    if false {
        switch "a" {
        case "foo bar":
            var v16 uint64 = 701
            _ = v16
        }
    } else {
        v14 = (^(v14))
    }
    v1 = v1
    var v17 float64 = 3.14159
    v18 := GenS4{F0: 1, F1: int64(v14), F2: 1}
    if (2147483648 <= 2) {
        for fi19 := 0; fi19 < 5; fi19++ {
            v17 = float64(float32((0.1 + float64(0.0))))
            var v20 int32 = 1
            var v21 float32 = ((1e10 * (1.5 * 1e10)) + 1e10)
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        if false {
            var v22 float32 = -2.25
            v23 := GenS3{F0: -128, F1: -981, F2: 10}
            _ = v22
            _ = v23
        }
    }
    fmt.Println(v1)
    switch 32767 {
    case -2147483648:
        v24 := GenS5{F0: "foo bar", F1: 65535, F2: 2147483647}
        _ = v24
    }
    var v25 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
}
