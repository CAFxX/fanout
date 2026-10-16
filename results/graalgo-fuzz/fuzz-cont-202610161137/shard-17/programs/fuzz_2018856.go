package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []int{32768, 0, -2147483648}
    v2 := []int{128, -85, -4, 154}
    var v3 uint8 = 198
    var v4 int8 = 42
    var v5 int32 = 2147483647
    v6 := make(map[int]string)
    _ = len(v2)
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 float64 = 3.14159
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            v4 += (v4 >> ((817) & 63))
            _ = fi10
        }
        v4 -= int8(v3)
        _ = fi7
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if false {
            var v12 uint16 = 1
            for fi13 := 0; fi13 < 9; fi13++ {
                v2 = append(v2, fi13)
                _ = fi13
            }
            _ = v12
        } else {
            delete(v6, 42)
            var v14 bool = true
            v3 += (v3 >> ((0) & 63))
            _ = v14
        }
        _ = fi11
    }
    fmt.Println(v3, v4)
    for fi15 := 0; fi15 < 2; fi15++ {
        v3 += (v3 ^ 2)
        v4 = 100
        _ = fi15
    }
    if (18446744073709551615 < 795) {
        v16 := v6[42]
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v16
    }
    v4 -= -128
    var v17 string = "foo bar"
    var v18 int16 = int16(v4)
    fmt.Println(v17)
    var v19 int8 = 1
    v19 = v4
    v20 := GenS1{F0: 127}
    var v21 uint16 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
