package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 18446744073709551615
    var v3 float64 = 1.5
    var v4 int8 = 0
    v5 := make(map[string]string)
    v6 := make(map[uint32]uint32)
    switch ("\n\t" + "hello") {
    case "\n\t":
        v7 := GenS1{F0: -909, F1: 42, F2: 255}
        _ = v7
    case "":
        v2 -= 4294967295
    case "hello":
        v1 += 1
    default:
        v8 := GenS1{F0: -1, F1: 661, F2: 25}
        _ = v8
    }
    if false {
        var v9 int32 = 920
        fmt.Println(v2, v1)
        for fi10 := 0; fi10 < 10; fi10++ {
            v3 = v3
            var v11 int = fi10
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        var v12 int64 = 9223372036854775807
        for fi13 := 0; fi13 < 12; fi13++ {
            v12 -= v12
            v14 := []int{fi13, (fi13 << ((60) & 63)), fi13, 42}
            v3 -= v3
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            v4 += 127
            fmt.Println(v1)
            _ = fi15
        }
        _ = v12
    }
    var v16 float64 = 0.1
    if true {
        v17 := []int{32768, 977, 128}
        v18 := GenS1{F0: -85, F1: v2, F2: 127}
        _ = v17
        _ = v18
    }
    var v19 int64 = 9223372036854775807
    v3 += (v3 * float64(1.5))
    switch 32767 {
    case 128:
        var v20 int16 = 0
        _ = v20
    case 0:
        v21 := []int{68, 42, -2147483648, 128}
        _ = v21
    default:
        var v22 int16 = 1
        _ = v22
    }
    if false {
        switch 759 {
        case 1:
            fmt.Println(v19, v1)
        default:
            v1 = (v1 ^ 65535)
        }
        _ = len(v5)
    }
    var v23 int64 = 42
    switch 4294967295 {
    case 42:
        var v24 float32 = 0.0
        _ = v24
    default:
        v19 = (v23 >> ((18446744073709551615) & 63))
    }
    var v25 uint = 219
    v25 += (v25 | 2)
    v26 := GenS2{F0: 207}
    v1 += 2147483647
    var v27 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
