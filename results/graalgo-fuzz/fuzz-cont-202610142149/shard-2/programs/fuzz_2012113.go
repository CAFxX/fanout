package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint8 }
func genFn0(p0 int64) int {
    v5 := []int{-490, 128}
    v6 := GenS1{F0: "hello", F1: 255, F2: p0}
    _ = p0
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 uint = 87
    var v2 int64 = -1
    v3 := make(map[uint32]uint32)
    var v4 int8 = -1
    fmt.Println(v2)
    if false {
        v7 := GenS2{F0: 4294967295}
        fmt.Println(v2)
        _ = v7
    }
    var v8 int8 = v4
    v3[2147483647] = 874
    if true {
        if false {
            v4 += v4
            v8 += (v4 % (((v8 % (((v8 & (v8 + v4))) | 1))) | 1))
        } else {
            var v9 float64 = 0.0
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 uint8 = 59
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        v12 := GenS3{F0: 80, F1: -32768, F2: 127}
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        if true {
            var v14 int16 = 1
            _ = v14
        }
        v8 = 30
        _ = fi13
    }
    v8 -= v8
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            v4 -= v4
            _ = fi16
        }
        v17 := []int{(fi15 ^ 128)}
        switch "\n\t" {
        case "":
            v8 = v8
        default:
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        }
        _ = fi15
        _ = v17
    }
    v18 := GenS3{F0: 65535, F1: 1, F2: 42}
    v19 := GenS3{F0: 1, F1: 1000, F2: 255}
    var v20 uint16 = 0
    v2 -= -9223372036854775808
    v8 = v4
    v21 := []int{592, 42}
    var v22 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
}
