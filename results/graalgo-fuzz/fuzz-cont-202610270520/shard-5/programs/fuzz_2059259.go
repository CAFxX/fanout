package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint32, p1 uint16) uint64 {
    var v5 bool = true
    fmt.Println(p0, v5, p0)
    v5 = ((v5 || v5) && false)
    fmt.Println(p0, p1, p1)
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1() uint32 {
    var v6 bool = false
    var v7 string = "\n\t"
    fmt.Println(v7)
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 float64 = 1e10
    v2 := []uint32{0}
    var v3 uint8 = 102
    var v4 int8 = 104
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 float64 = (float64((float32(0.0) / (((1.5 - 0.1) - (0.1 * 3.14159)) / 1.0))) + (v1 / -2.25))
        _ = fi8
        _ = v9
    }
    fmt.Println(v4, v1, v4)
    var v10 float64 = v1
    var v11 uint = 1
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    var v12 uint64 = 1
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            v11 = v11
            v2 = append(v2, 65535)
            v1 += v10
            _ = fi14
        }
        v4 = (v4 & int8(fi13))
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(v1, v12, fi15)
            v4 -= (v4 | (v4 - v4))
            var v16 uint8 = uint8(v11)
            _ = fi15
            _ = v16
        }
        _ = fi13
    }
    v17 := GenS1{F0: "a"}
    switch "foo bar" {
    case "hello":
        switch -1 {
        case 32767:
            fmt.Println(v12, v3)
        }
    case "":
        v18 := GenS1{F0: "\n\t"}
        _ = v18
    }
    var v19 string = "hello"
    switch (v19 + v19) {
    case "":
        fmt.Println(v3, v1)
    case "foo bar":
        if (384 < 42) {
            var v20 int = -933
            v12 = 18446744073709551615
            var v21 uint64 = uint64(v11)
            _ = v20
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 12; fi22++ {
            v19 = ("\n\t" + (v19 + v19))
            for fi23 := 0; fi23 < 2; fi23++ {
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                _ = fi23
            }
            var v24 uint16 = 986
            _ = fi22
            _ = v24
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v11 = (v11 << ((v11) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
