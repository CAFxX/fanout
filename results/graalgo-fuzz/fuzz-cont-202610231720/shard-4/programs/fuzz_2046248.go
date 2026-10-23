package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0() bool {
    var v7 int = 42
    v7 += v7
    _ = v7
    return false
}
func genFn1(p0 int64, p1 int) string {
    v8 := []int{(p1 >> ((382) & 63)), (p1 / (((p1 >> ((42) & 63))) | 1)), 2147483647}
    var v9 uint64 = 1
    p0 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 float32 = 3.14159
    var v4 uint16 = 872
    v5 := make(map[uint32]int)
    var v6 int32 = -879
    v10 := GenS1{F0: 42, F1: 42, F2: -128}
    var v11 uint8 = 147
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS2{F0: 159, F1: -1, F2: (v11 % (((v11 ^ v11)) | 1))}
        for fi14 := 0; fi14 < 6; fi14++ {
            v6 = v6
            v15 := []int{fi12, -2147483648, fi12, (fi12 & (fi14 >> ((4294967295) & 63)))}
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    if (0.0 <= float64(0.1)) {
        v4 = v4
        v4 -= (v4 << ((4294967295) & 63))
        v11 += (^(v11))
    }
    v4 = (v4 << ((255) & 63))
    v16 := []int{9223372036854775807, -839}
    var v17 string = ""
    switch ("foo bar" + v17) {
    case "":
        v3 -= v3
    default:
        var v18 uint8 = v11
        _ = v18
    }
    _ = len(v16)
    v3 += 3.14159
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            v4 = v4
            _ = fi20
        }
        var v21 float64 = 0.1
        _ = fi19
        _ = v21
    }
    v22 := v2[1]
    v3 = ((3.14159 / (float32(((3.14159 - (0.0 / 0.1)) + (0.0 - (-2.25 * (1e10 - 1e10))))) / 1.5)) * float32((1.5 * -2.25)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
