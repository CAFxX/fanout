package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint16 }
func genFn0() bool {
    var v5 bool = (42 != 956)
    v5 = ((242 < 1) && v5)
    fmt.Println(v5)
    _ = v5
    return (32768 != 42)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 0
    v3 := make(map[int]bool)
    var v4 uint = 813
    if true {
        v6 := []int{}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := []int{(fi7 % (((v2 + v2)) | 1))}
            switch fi7 {
            case 32768:
                fmt.Println(v4, fi7)
            case -2147483648:
                var v9 float64 = float64(0.1)
                _ = v9
            default:
                _ = len(v3)
            }
            var v10 float64 = (0.0 / float64(1.5))
            _ = fi7
            _ = v10
            _ = v8
        }
    }
    v11 := GenS1{F0: -74, F1: 0, F2: 255}
    switch (v2 >> (((v4 ^ v4)) & 63)) {
    case -1:
        v4 = v4
    case -9223372036854775808:
        var v12 uint32 = 1
        _ = v12
    }
    var v13 string = ""
    var v14 uint32 = 65535
    _ = len(v1)
    v14 -= v14
    v15 := GenS1{F0: 9223372036854775807, F1: 284, F2: 851}
    var v16 int32 = 1
    fmt.Println(v16, v14, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
