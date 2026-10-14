package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = float64((-2.25 - 1.5))
    var v3 uint32 = 636
    var v4 float32 = ((1e10 + 0.0) / ((float32(v2) - float32((((v2 - (v2 + (3.14159 - ((v2 * v2) / -2.25)))) * (v2 / float64(float32((-2.25 * 1e10))))) / 0.1))) / (1e10 + ((((1e10 * (0.1 + 0.1)) - 1e10) + 1.5) - 1e10))))
    v5 := GenS1{F0: 18446744073709551615, F1: ((v3 >> ((18446744073709551615) & 63)) == 0)}
    v1[v3] = "a"
    fmt.Println(v4, v3)
    fmt.Println(v2, v4)
    delete(v1, (v3 / (((v3 >> ((2) & 63))) | 1)))
    var v6 uint8 = 0
    var v7 uint8 = v6
    v3 = 2147483648
    switch 434 {
    case 4294967295:
        var v8 uint32 = v3
        _ = v8
    case -1:
        v9 := []int{1, 2147483647, 42, 128}
        _ = v9
    }
    v10 := v1[2147483648]
    var v11 int64 = 394
    switch 4294967295 {
    case -229:
        var v12 uint16 = 42
        _ = v12
    default:
        var v13 int64 = (v11 >> ((255) & 63))
        _ = v13
    }
    if (!(v3 >= 637)) {
        var v14 float32 = v4
        var v15 float64 = v2
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int = fi16
            v6 += 255
            _ = fi16
            _ = v17
        }
    }
    switch (v10 + v10) {
    case "":
        var v18 uint64 = 620
        _ = v18
    case "foo bar":
        var v19 uint = 845
        _ = v19
    default:
        var v20 int32 = -199
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
