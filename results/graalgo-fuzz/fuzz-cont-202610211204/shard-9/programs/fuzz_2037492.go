package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    var v1 uint64 = 42
    var v2 int32 = 738
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v5 := GenS1{F0: v2, F1: -92}
    for fi6 := 0; fi6 < 10; fi6++ {
        if false {
            if false {
                if false {
                    v2 = v2
                    v7 := []int{(fi6 * fi6), (fi6 >> ((4294967295) & 63)), fi6}
                    _ = v7
                }
            }
        }
        _ = fi6
    }
    v1 -= v1
    switch 32768 {
    case 127:
        var v8 uint = 0
        _ = v8
    case -2147483648:
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                if (1 != 1) {
                    v10 := GenS1{F0: (-(v2)), F1: (fi9 % (((fi9 + fi9)) | 1))}
                    _ = v10
                } else {
                    var v11 uint8 = 237
                    fmt.Println(fi9, v11, v2)
                    var v12 bool = true
                    _ = v11
                    _ = v12
                }
                var v13 float64 = 1.5
                _ = fi9
                _ = v13
            }
            v1 = 435
        } else {
            var v14 int64 = 1
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v2 = 986
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS1{F0: v2, F1: 4294967295}
            v18 := []int{-9223372036854775808}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    v3[65535] = 2147483648
    var v19 uint16 = 532
    v1 += 470
    v19 += (v19 * v19)
    var v20 uint = 295
    var v21 int64 = 0
    fmt.Println(v19, v19)
    var v22 uint32 = 786
    v20 = 48
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
