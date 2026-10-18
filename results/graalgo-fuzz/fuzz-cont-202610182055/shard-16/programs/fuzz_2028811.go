package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() int {
    var v6 int64 = 9223372036854775807
    var v7 uint8 = 42
    if (!(false && true)) {
        var v8 int16 = 841
        var v9 int8 = 127
        switch "a" {
        case "hello":
            var v10 int8 = v9
            _ = v10
        }
        _ = v8
        _ = v9
    } else {
        switch "\n\t" {
        case "hello":
            var v11 uint = 65535
            _ = v11
        case "a":
            var v12 uint = 479
            _ = v12
        }
    }
    v13 := []int{32767, 473, 951}
    v7 += 255
    _ = v13
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 bool, p1 uint16) int8 {
    p1 = (p1 & (p1 * (p1 | p1)))
    p1 = p1
    var v14 uint = 632
    _ = p0
    _ = p1
    _ = v14
    return 1
}

func main() {
    v1 := []string{"hello", "a"}
    var v2 uint16 = 32767
    var v3 int8 = 99
    v4 := []uint32{839}
    v5 := []string{}
    v3 -= (v3 << ((0) & 63))
    v2 = (^(v2))
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (259 != 4294967295) {
        v3 = v3
        v2 = (v2 ^ v2)
    }
    var v15 uint8 = 127
    if ((true || true) || false) {
        v16 := GenS1{F0: 42}
        _ = v16
    }
    var v17 float32 = float32((1.5 / ((0.0 - 1.5) * 1e10)))
    v15 -= (^(v15))
    v2 += 665
    v18 := GenS1{F0: 0}
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 int8 = (v3 / (((^(v3))) | 1))
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 934}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 string = "\n\t"
    var v23 uint64 = 148
    v24 := GenS1{F0: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
