package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() uint {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 string = "foo bar"
        if true {
            var v9 uint16 = 32767
            v10 := []int{fi7}
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    v11 := []int{-9223372036854775808, 642}
    var v12 uint = 270
    v12 += 228
    _ = v11
    _ = v12
    return v12
}
func genFn1() uint32 {
    v13 := GenS1{F0: 42}
    var v14 float64 = -2.25
    v14 = (v14 / (v14 / v14))
    _ = v13
    _ = v14
    return 952
}

func main() {
    var v1 int16 = 1
    var v2 uint8 = 123
    v3 := []uint32{4294967295}
    var v4 uint8 = 42
    var v5 string = "foo bar"
    v6 := make(map[uint32]string)
    fmt.Println(v5)
    v4 = 200
    v2 = (^(v4))
    fmt.Println(v5)
    switch -1 {
    case 1:
        if (!false) {
            for fi15 := 0; fi15 < 1; fi15++ {
                _ = len(v3)
                _ = fi15
            }
        } else {
            v4 -= (v4 >> ((1) & 63))
            var v16 uint8 = v4
            v17 := GenS1{F0: 147}
            _ = v16
            _ = v17
        }
    case -364:
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 uint8 = v2
            v3 = append(v3, 4294967295)
            var v20 uint = 868
            _ = fi18
            _ = v19
            _ = v20
        }
    case -70:
        var v21 uint = 958
        _ = v21
    default:
        v1 -= int16(v2)
    }
    var v22 float64 = 0.1
    v23 := GenS1{F0: 9223372036854775807}
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
