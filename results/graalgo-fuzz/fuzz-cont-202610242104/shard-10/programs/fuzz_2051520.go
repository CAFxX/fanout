package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint16 = 1
    var v2 uint8 = 255
    var v3 bool = false
    var v4 uint = 2
    var v5 int32 = 414
    v6 := GenS1{F0: 42}
    v7 := []int{128, -288, int(v1), 2147483647}
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS1{F0: 1}
            v3 = v3
            v11 := GenS1{F0: 0}
            v3 = (0 != 2)
            v12 := GenS1{F0: 1000}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int8 = 34
            if len(v7) > 0 {
                v7[0] = (fi14 + 482)
            }
            var v16 bool = v3
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int8 = 100
        if (2 > 2) {
            v7 = append(v7, 9223372036854775807)
            fmt.Println(v3)
            var v18 uint8 = (v2 - v2)
            _ = v18
        }
        _ = fi13
        _ = v17
    }
    var v19 uint8 = 128
    v2 += (v19 << (((v4 | 475)) & 63))
    var v20 bool = v3
    var v21 uint16 = v1
    fmt.Println(v5, v4, v3)
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            v19 -= 242
            var v24 float64 = 1.5
            _ = fi23
            _ = v24
        }
        var v25 string = "\n\t"
        _ = fi22
        _ = v25
    }
    if v20 {
        v7 = append(v7, -972)
        var v26 float64 = 0.1
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
