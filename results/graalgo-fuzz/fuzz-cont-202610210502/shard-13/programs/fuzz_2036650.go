package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    var v4 int32 = 0
    v5 := make(map[uint32]bool)
    v4 = v4
    v4 = v4
    switch "" {
    case "":
        if (((2 != 65535) || true) && (993 > 2)) {
            var v6 float64 = 0.0
            v5[2147483648] = (false && ((!(4294967295 != 42)) && (4294967295 < 459)))
            var v7 int = -1
            _ = v6
            _ = v7
        }
    }
    switch -2147483648 {
    case -551:
        _ = len(v2)
    default:
        var v8 int32 = (^(v4))
        _ = v8
    }
    var v9 uint32 = 2
    v9 -= 0
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 float64 = (0.1 * ((((((0.0 + (-2.25 * -2.25)) + 0.0) / 1.0) / (1.5 * 0.1)) / 3.14159) - 3.14159))
        if (!(!true)) {
            switch "hello" {
            case "hello":
                v12 := []int{fi10, fi10}
                _ = v12
            }
        }
        _ = fi10
        _ = v11
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 int16 = -32768
        for fi15 := 0; fi15 < 9; fi15++ {
            switch fi15 {
            case -637:
                var v16 int = fi13
                _ = v16
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    v9 = (v9 & v9)
    v9 += (v9 / ((4294967295) | 1))
    v9 += 4294967295
    var v17 int64 = 1
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS1{F0: (fi18 & 4294967295)}
        v17 = 14
        _ = fi18
        _ = v19
    }
    fmt.Println(v4)
    v4 = v4
    v20 := GenS1{F0: 32767}
    var v21 bool = false
    var v22 int16 = -746
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
