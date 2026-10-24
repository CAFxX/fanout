package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = -2147483648
    v2 := []int{-805, 32768, -2147483648, 790}
    var v3 bool = false
    var v4 int = -748
    var v5 float32 = (-2.25 / 1e10)
    var v6 string = (("a" + "\n\t") + "")
    v7 := GenS1{F0: 128}
    for fi8 := 0; fi8 < 12; fi8++ {
        if true {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 uint16 = 0
                v11 := GenS1{F0: 255}
                v12 := GenS2{F0: 42, F1: v1, F2: 128}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            fmt.Println(v1, v5, v3)
        } else {
            switch v4 {
            case 2147483647:
                if len(v2) > 0 {
                    v2[0] = -9223372036854775808
                }
            default:
                var v13 uint = 4294967295
                _ = v13
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                if len(v2) > 0 {
                    v2[0] = (fi8 / ((fi8) | 1))
                }
                _ = fi14
            }
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 uint = 826
            v3 = (true || (v16 < v16))
            _ = fi15
            _ = v16
        }
        var v17 float64 = 0.1
        _ = fi8
        _ = v17
    }
    var v18 string = "foo bar"
    var v19 uint16 = 255
    _ = len(v2)
    v20 := GenS3{F0: 158}
    _ = len(v2)
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS3{F0: 1}
        _ = fi21
        _ = v22
    }
    if len(v2) > 0 {
        v2[0] = v4
    }
    v23 := GenS3{F0: 18446744073709551615}
    fmt.Println(v4)
    var v24 uint16 = v19
    v24 += 65535
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
