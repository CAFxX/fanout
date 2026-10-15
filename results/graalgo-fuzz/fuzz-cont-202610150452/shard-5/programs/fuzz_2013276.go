package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int8 }
func genFn0(p0 int8) int16 {
    var v4 uint32 = 2147483648
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 += p0
        for fi6 := 0; fi6 < 12; fi6++ {
            v4 -= v4
            var v7 float64 = (0.0 * 3.14159)
            var v8 int32 = -173
            _ = fi6
            _ = v7
            _ = v8
        }
        v4 -= v4
        _ = fi5
    }
    var v9 float64 = (3.14159 - -2.25)
    _ = p0
    _ = v4
    _ = v9
    return 1000
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = 0.0
    var v3 uint = 704
    var v10 float32 = (0.1 - 0.0)
    if (-1 < 127) {
        v11 := []int{128, 728, 42}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(v10, v10, v3)
            _ = fi12
        }
    }
    v13 := GenS1{F0: 42, F1: v3, F2: 1}
    for fi14 := 0; fi14 < 8; fi14++ {
        v1 = v1
        switch fi14 {
        case 0:
            v15 := []int{fi14, (fi14 * 128), (fi14 & fi14)}
            _ = v15
        case 1:
            fmt.Println(v1, v1, v3)
        default:
            var v16 uint16 = 65535
            _ = v16
        }
        _ = fi14
    }
    v10 = v10
    v17 := []int{127, -2147483648, -9223372036854775808, -599}
    if (127 >= -1) {
        var v18 int = -188
        _ = v18
    }
    v2 -= -2.25
    _ = len(v17)
    if (true && (!(0 == 1))) {
        v1 = v1
    }
    v1 -= v1
    fmt.Println(v2, v2, v10)
    var v19 uint64 = 42
    v20 := GenS1{F0: -32768, F1: 390, F2: 90}
    switch "hello" {
    case "hello":
        switch -2147483648 {
        case 633:
            if false {
                var v21 int32 = 0
                _ = v21
            }
        case -9223372036854775808:
            var v22 uint64 = (^(v1))
            _ = v22
        }
    case "\n\t":
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    default:
        if false {
            v19 -= v1
            var v23 int16 = 359
            var v24 uint = v3
            _ = v23
            _ = v24
        }
    }
    if (-2147483648 < -772) {
        v17 = append(v17, -535)
        v10 += v10
    }
    var v25 int32 = 0
    v26 := GenS1{F0: 1000, F1: v3, F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
}
