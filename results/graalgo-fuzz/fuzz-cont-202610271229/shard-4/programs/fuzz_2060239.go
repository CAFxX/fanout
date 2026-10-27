package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }

func main() {
    var v1 float64 = 1.5
    var v2 uint32 = 0
    v3 := []uint32{v2, 2147483647, (v2 / ((v2) | 1)), (v2 & (v2 + (^(v2))))}
    fmt.Println(v1, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 := GenS1{F0: 65535, F1: v2, F2: -9223372036854775808}
    v2 = v2
    v1 += v1
    for fi5 := 0; fi5 < 1; fi5++ {
        if (!true) {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 int8 = -14
                v1 += (v1 + 0.0)
                if len(v3) > 0 {
                    v3[0] = v2
                }
                _ = fi6
                _ = v7
            }
            var v8 int64 = 9223372036854775807
            _ = v8
        }
        v9 := GenS1{F0: 417, F1: 42, F2: (fi5 | (fi5 - fi5))}
        _ = fi5
        _ = v9
    }
    switch 127 {
    case 32768:
        var v10 uint = 2
        _ = v10
    case 467:
        if len(v3) > 0 {
            v3[0] = 65535
        }
    case 42:
        v1 = (3.14159 + (v1 - (1.5 / -2.25)))
    }
    var v11 int32 = 2147483647
    switch 2147483647 {
    case 4294967295:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -43:
        v3 = append(v3, v2)
    }
    v11 = (v11 + (v11 & (v11 >> ((0) & 63))))
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(fi12, v1, v1)
        _ = fi12
    }
    v2 = 2147483647
    var v13 uint8 = 127
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!(437 == 0)) {
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 uint64 = 2
            var v16 float32 = ((0.0 / (3.14159 - (-2.25 + float32(v1)))) + -2.25)
            _ = fi14
            _ = v15
            _ = v16
        }
        if ((float64(((1.5 + 0.1) - 1.5)) / ((float64(0.1) - v1) - v1)) > (v1 - v1)) {
            _ = len(v3)
            fmt.Println(v11, v2, v2)
            v3 = append(v3, 4294967295)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        switch -658 {
        case -628:
            var v17 float64 = v1
            _ = v17
        case 127:
            var v18 uint32 = 42
            _ = v18
        case 2147483647:
            fmt.Println(v1, v1)
        }
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 float64 = -2.25
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi19
        _ = v20
    }
    var v21 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
