package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint32 }
func genFn0(p0 uint64) string {
    var v5 uint = 576
    switch "" {
    case "foo bar":
        if (32767 >= 2) {
            var v6 int64 = -120
            var v7 float32 = (float32(0.1) + ((float32((3.14159 - ((float64(float32(0.1)) + ((1e10 * 0.0) * 1e10)) / (1e10 - 1.5)))) + 0.1) - (0.0 / ((0.0 - ((0.1 + 3.14159) / 1e10)) + (1e10 + float32(1e10))))))
            _ = v6
            _ = v7
        } else {
            var v8 bool = ((358 >= 0) && true)
            _ = v8
        }
    case "":
        var v9 uint8 = 42
        _ = v9
    default:
        v10 := []int{1, 60, -1, -1}
        _ = v10
    }
    if false {
        v11 := GenS1{F0: 42, F1: (v5 ^ v5), F2: p0}
        fmt.Println(v5, v5)
        _ = v11
    } else {
        var v12 int8 = -1
        v13 := []int{2147483647}
        _ = v12
        _ = v13
    }
    switch 128 {
    case 4294967295:
        fmt.Println(p0)
    case 127:
        for fi14 := 0; fi14 < 12; fi14++ {
            if (100 > 127) {
                var v15 uint32 = 2147483648
                v5 -= (v5 - (v5 % (((v5 - (v5 ^ v5))) | 1)))
                _ = v15
            }
            _ = fi14
        }
    case 42:
        if (1 <= 65535) {
            v5 = v5
            p0 += (p0 * (p0 >> ((v5) & 63)))
        } else {
            v5 += 65535
            fmt.Println(p0, p0)
        }
    default:
        v5 -= 255
    }
    var v16 uint64 = p0
    _ = p0
    _ = v16
    _ = v5
    return "\n\t"
}

func main() {
    var v1 uint = 135
    var v2 int = 1
    var v3 int32 = 649
    var v4 int = 32767
    v17 := []int{(-(v4)), v2, (v4 + -2147483648), v2}
    v1 = 1
    switch (("hello" + "foo bar") + "\n\t") {
    case "hello":
        fmt.Println(v4, v3, v3)
    }
    var v18 int8 = -14
    v19 := GenS2{F0: 2, F1: uint16(v1), F2: 42}
    var v20 bool = false
    if false {
        var v21 int8 = v18
        fmt.Println(v4)
        for fi22 := 0; fi22 < 1; fi22++ {
            if true {
                var v23 uint16 = 65535
                _ = v23
            }
            _ = fi22
        }
        _ = v21
    } else {
        v18 -= 0
    }
    var v24 uint32 = 2147483647
    v25 := GenS3{F0: 1, F1: (1 <= -507), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
