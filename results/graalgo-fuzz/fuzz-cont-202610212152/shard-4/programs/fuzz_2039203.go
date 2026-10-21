package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint8) int64 {
    v5 := []int{174, -1, 42, -1}
    p0 = p0
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 1}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return -9223372036854775808
}
func genFn1(p0 int) uint64 {
    p0 = p0
    v8 := GenS2{F0: 18446744073709551615, F1: 4294967295, F2: 255}
    _ = p0
    _ = v8
    return 761
}

func main() {
    var v1 int32 = 499
    var v2 int = 128
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    _ = len(v3)
    if true {
        var v9 uint = 970
        v9 = (v9 % (((v9 % ((v9) | 1))) | 1))
        _ = v9
    } else {
        v10 := GenS3{F0: 1}
        if false {
            switch (("\n\t" + ("foo bar" + "foo bar")) + (("hello" + "foo bar") + "\n\t")) {
            case "\n\t":
                v11 := GenS2{F0: 18446744073709551615, F1: 587, F2: 2}
                _ = v11
            case "":
                v12 := []int{}
                _ = v12
            case "hello":
                var v13 uint32 = 2147483648
                _ = v13
            default:
                var v14 float32 = 1e10
                _ = v14
            }
        } else {
            var v15 int16 = 689
            _ = v15
        }
        _ = v10
    }
    fmt.Println(v1, v1)
    var v16 uint64 = 0
    fmt.Println(v2, v2)
    fmt.Println(v16, v1)
    var v17 bool = (-32768 >= 1)
    var v18 float64 = (0.0 * 1.5)
    var v19 int64 = 1
    var v20 int64 = genFn0(1)
    var v21 int32 = v1
    v20 = v20
    v1 += v1
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int16 = -436
        _ = fi22
        _ = v23
    }
    v24 := GenS4{F0: (v16 ^ v16)}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
}
