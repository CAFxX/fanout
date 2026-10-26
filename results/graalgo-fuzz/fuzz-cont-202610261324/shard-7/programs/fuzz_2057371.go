package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint64 }
func genFn0() int {
    if true {
        v5 := GenS1{F0: true, F1: 0}
        _ = v5
    } else {
        var v6 uint = 0
        v7 := []int{42, 0}
        var v8 uint8 = 255
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{436, -695, -9223372036854775808, 1}
    _ = v9
    return 32767
}
func genFn1(p0 uint32, p1 uint8) uint8 {
    v10 := []int{32768, genFn0()}
    switch "foo bar" {
    case "\n\t":
        v10 = append(v10, -9223372036854775808)
    case "a":
        v10 = append(v10, 0)
    default:
        fmt.Println(p0, p0)
    }
    var v11 float32 = (1.5 - 3.14159)
    var v12 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return (p1 * (p1 | (p1 - (p1 << ((4294967295) & 63)))))
}

func main() {
    v1 := []uint32{4294967295, 42, 2, 525}
    var v2 int32 = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    v2 += 1
    if (-2.25 >= 1e10) {
        if false {
            var v13 int8 = 100
            v14 := GenS2{F0: 42, F1: 126}
            var v15 int = 9223372036854775807
            _ = v13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: 4294967295, F1: -32768}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 1; fi17++ {
            v2 = (v2 ^ v2)
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 float64 = float64(0.1)
                _ = fi18
                _ = v19
            }
            delete(v3, 73)
            _ = fi17
        }
    }
    var v20 string = ""
    v21 := v3[2147483648]
    if v21 {
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS3{F0: 0}
            _ = fi22
            _ = v23
        }
    }
    v2 -= v2
    if v21 {
        var v24 uint32 = 2147483647
        _ = v24
    }
    v20 = ("foo bar" + "\n\t")
    v1 = append(v1, 1)
    v2 = v2
    v21 = false
    for fi25 := 0; fi25 < 4; fi25++ {
        if len(v1) > 0 {
            v1[0] = 2
        }
        _ = fi25
    }
    if (((3.14159 * 0.0) * 3.14159) >= 3.14159) {
        v2 = v2
    } else {
        v2 = (^(v2))
    }
    v21 = (0 <= 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
