package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int; F2 int16 }
func genFn0(p0 int32, p1 uint16, p2 uint32) int {
    v5 := []int{865, 42, 9223372036854775807, 9223372036854775807}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 int16 = 59
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -672
}

func main() {
    var v1 uint = 0
    var v2 uint8 = 42
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]string)
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 2}
        v1 = v1
        _ = fi7
        _ = v8
    }
    var v9 float64 = -2.25
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint = 65535
        _ = fi10
        _ = v11
    }
    var v12 float32 = (0.0 - float32((-2.25 - (-2.25 / v9))))
    if false {
        switch ("a" + "\n\t") {
        case "a":
            var v13 int64 = -9223372036854775808
            _ = v13
        case "\n\t":
            fmt.Println(v1, v2)
        }
    }
    fmt.Println(v1)
    v14 := GenS1{F0: 2}
    v3 = v3
    var v15 int = 670
    v16 := GenS2{F0: -1000, F1: v15, F2: 680}
    var v17 float64 = v9
    if false {
        switch v15 {
        case 1:
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 int8 = -1
                v20 := GenS2{F0: -524, F1: 1, F2: -1}
                var v21 int64 = -898
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        case 32767:
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS2{F0: 0, F1: (fi22 ^ 4294967295), F2: 1}
                _ = fi22
                _ = v23
            }
        default:
            v24 := []int{}
            _ = v24
        }
    } else {
        if true {
            var v25 uint = 255
            _ = v25
        } else {
            for fi26 := 0; fi26 < 14; fi26++ {
                v27 := GenS2{F0: -2147483648, F1: genFn0(0, 708, 2), F2: 32767}
                _ = fi26
                _ = v27
            }
        }
    }
    v3 = v3
    v1 = v1
    v9 += 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
