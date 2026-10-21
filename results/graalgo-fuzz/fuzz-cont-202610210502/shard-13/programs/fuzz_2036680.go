package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 int64; F2 uint32 }

func main() {
    var v1 int32 = 42
    var v2 uint64 = 42
    var v3 int16 = 32767
    var v4 uint8 = 199
    var v5 float64 = (-2.25 + float64(1.5))
    if true {
        if false {
            v6 := GenS1{F0: 2147483647}
            v7 := []int{-2147483648, 32768, 2147483647, -1}
            _ = v6
            _ = v7
        }
        if false {
            if (("a" + "\n\t") <= "a") {
                var v8 int = 265
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := []int{-676, v8, (-(fi9))}
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                v3 = 0
                if (32768 >= 4294967295) {
                    v2 -= (v2 << ((774) & 63))
                    v5 = v5
                    v11 := GenS2{F0: v3, F1: (v2 + v2)}
                    _ = v11
                }
            }
            v12 := []int{-481, 9223372036854775807}
            _ = v12
        } else {
            v2 -= 4294967295
            fmt.Println(v4, v1)
            for fi13 := 0; fi13 < 17; fi13++ {
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 float32 = -2.25
                    _ = fi14
                    _ = v15
                }
                v16 := []int{fi13, 0, (fi13 + fi13), (fi13 + fi13)}
                _ = fi13
                _ = v16
            }
        }
    }
    v1 += (-(v1))
    v5 -= float64((1.5 + float32(((-2.25 / (((float64(1e10) * (v5 * (0.0 + v5))) + 3.14159) - v5)) * v5))))
    if false {
        switch "foo bar" {
        case "\n\t":
            v17 := GenS3{F0: 182, F1: 9223372036854775807, F2: 4294967295}
            _ = v17
        case "":
            if true {
                v18 := GenS2{F0: (v3 + 0), F1: v2}
                v19 := GenS1{F0: v1}
                v4 = 127
                _ = v18
                _ = v19
            }
        case "hello":
            if ((4294967295 <= uint32(v2)) || (0 > 900)) {
                if (!(-2.25 < ((((v5 * v5) / v5) - v5) / 1.0))) {
                    v5 = v5
                }
                v3 = (v3 - -32768)
                if false {
                    v5 -= float64(-2.25)
                    var v20 bool = (475 == 2)
                    _ = v20
                }
            }
        }
        var v21 uint32 = 954
        _ = v21
    }
    v1 = (v1 * v1)
    v22 := []int{0, 42, 0, -1}
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v3, v2, v3)
        _ = fi23
    }
    fmt.Println(v3, v3, v2)
    v24 := GenS3{F0: 248, F1: 419, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
