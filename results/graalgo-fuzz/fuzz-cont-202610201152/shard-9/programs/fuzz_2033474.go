package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 string = "a"
    var v2 bool = true
    var v3 uint16 = 1
    var v4 float32 = 1e10
    v5 := []uint32{70, 65535, 4294967295}
    v6 := make(map[uint32]uint32)
    v3 += 908
    v2 = (479 <= -32768)
    fmt.Println(v1)
    if true {
        v3 += 65535
        v7 := GenS1{F0: 1, F1: 110, F2: -11}
        _ = v7
    }
    var v8 int = 665
    var v9 int32 = -2147483648
    fmt.Println(v2, v8)
    var v10 float64 = 1.5
    var v11 uint32 = uint32(v3)
    v4 += v4
    v6[v11] = 1
    var v12 float32 = ((0.1 + float32(v10)) / float32(v10))
    v13 := GenS2{F0: 198}
    switch (v1 + v1) {
    case "":
        _ = len(v6)
    case "\n\t":
        switch (v8 & v8) {
        case -955:
            fmt.Println(v3, v12)
        case 32768:
            _ = len(v6)
        default:
            v6[(^(v11))] = v11
        }
    case "a":
        if ((100 <= 127) && (18446744073709551615 >= 658)) {
            v14 := GenS2{F0: 255}
            _ = v14
        } else {
            v5 = append(v5, v11)
            v3 = v3
            var v15 float64 = (v10 + v10)
            _ = v15
        }
    default:
        var v16 int = (v8 | (v8 / ((-2147483648) | 1)))
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS2{F0: 42}
        v11 = 875
        _ = fi17
        _ = v18
    }
    if (v4 != (v12 / ((v12 - 3.14159) - -2.25))) {
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 uint8 = 127
            _ = fi19
            _ = v20
        }
    } else {
        v21 := GenS1{F0: 42, F1: 127, F2: 1}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
