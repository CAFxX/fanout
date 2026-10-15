package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int16, p1 string) uint32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        switch fi6 {
        case 1:
            for fi7 := 0; fi7 < 1; fi7++ {
                v8 := GenS1{F0: "a", F1: 2, F2: 65535}
                v9 := []int{4294967295}
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            p0 = (p0 - (p0 ^ 708))
        }
        if (1.5 <= (0.0 * 3.14159)) {
            var v10 float64 = (0.0 + 1e10)
            var v11 int = (-(fi6))
            fmt.Println(v11, p1, v10)
            _ = v10
            _ = v11
        }
        var v12 uint8 = 0
        _ = fi6
        _ = v12
    }
    v13 := GenS1{F0: p1, F1: 18446744073709551615, F2: 2147483648}
    var v14 bool = (!(0 < 2147483647))
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 4294967295
}
func genFn1(p0 uint, p1 uint8, p2 int64) uint64 {
    v15 := GenS1{F0: "", F1: 2, F2: genFn0(48, ("" + "a"))}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return uint64(p1)
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]bool)
    var v3 uint8 = 0
    var v4 uint8 = v3
    var v5 uint = 255
    var v16 float32 = 0.0
    var v17 float64 = float64(v16)
    switch ("\n\t" + "\n\t") {
    case "hello":
        v4 = uint8(v5)
    case "\n\t":
        v16 = v16
    case "a":
        v5 += 42
    }
    v18 := GenS1{F0: "\n\t", F1: 4294967295, F2: 812}
    if false {
        switch 32767 {
        case -732:
            delete(v2, 42)
        default:
            fmt.Println(v1, v1)
        }
    }
    var v19 uint64 = 4294967295
    v20 := GenS2{F0: -2147483648}
    var v21 uint = 18446744073709551615
    v22 := GenS2{F0: 2147483647}
    v5 = (v21 - (^(v21)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
