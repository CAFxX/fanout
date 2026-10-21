package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int; F1 bool; F2 uint64 }
func genFn0(p0 uint) uint32 {
    switch 128 {
    case -1:
        fmt.Println(p0)
    }
    p0 = (p0 ^ p0)
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{}
        var v8 uint16 = 2
        var v9 string = "foo bar"
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    return 844
}

func main() {
    var v1 int64 = -117
    var v2 int = 193
    v3 := make(map[string]int)
    var v4 int8 = 100
    var v5 float32 = 0.0
    switch v2 {
    case 1:
        v10 := GenS1{F0: 1, F1: int32(v4)}
        _ = v10
    case -2147483648:
        fmt.Println(v4)
    }
    var v11 uint64 = 0
    var v12 int16 = 1000
    switch ("a" + ("" + "\n\t")) {
    case "a":
        v13 := GenS1{F0: 872, F1: 1}
        _ = v13
    case "foo bar":
        if false {
            v5 = (v5 * ((v5 / (((((v5 + v5) + 1.5) * 3.14159) - (1e10 * ((v5 - -2.25) * (1.5 * v5)))) * v5)) - -2.25))
            v3[("hello" + ("foo bar" + ""))] = v2
            var v14 uint16 = 363
            _ = v14
        } else {
            var v15 string = "\n\t"
            v16 := GenS2{F0: 42, F1: false, F2: v11}
            var v17 uint64 = (v11 - (v11 & v11))
            _ = v15
            _ = v16
            _ = v17
        }
    default:
        v18 := []int{v2, int(v12)}
        _ = v18
    }
    switch "a" {
    case "hello":
        var v19 bool = false
        _ = v19
    default:
        v20 := GenS2{F0: (v2 - (v2 % (((v2 - -529)) | 1))), F1: false, F2: (v11 + 0)}
        _ = v20
    }
    v12 = -784
    v2 -= (v2 % ((0) | 1))
    fmt.Println(v1, v11, v12)
    v1 -= v1
    fmt.Println(v1, v1, v1)
    var v21 uint64 = v11
    v12 += int16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
