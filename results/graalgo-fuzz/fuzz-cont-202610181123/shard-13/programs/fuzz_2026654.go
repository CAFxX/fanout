package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0(p0 int32, p1 bool, p2 int8) int32 {
    v6 := GenS1{F0: 990}
    p1 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -2147483648
    var v3 uint16 = 255
    v4 := make(map[string]string)
    var v5 int32 = v2
    var v7 uint = 42
    for fi8 := 0; fi8 < 15; fi8++ {
        v3 -= 646
        if false {
            _ = len(v4)
            if false {
                v7 = (v7 << ((42) & 63))
                var v9 int32 = (v5 << ((v7) & 63))
                var v10 string = "a"
                _ = v10
                _ = v9
            } else {
                var v11 float32 = (-2.25 - 1.5)
                _ = v11
            }
        }
        v7 -= 787
        _ = fi8
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        if true {
            var v13 float64 = 3.14159
            v1[42] = 0
            v14 := []int{(fi12 >> ((18446744073709551615) & 63))}
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS2{F0: 2}
                fmt.Println(v7, v7, v2)
                fmt.Println(v5)
                _ = fi15
                _ = v16
            }
            var v17 bool = true
            _ = v17
        }
        v18 := GenS3{F0: 9223372036854775807, F1: 127, F2: fi12}
        _ = fi12
        _ = v18
    }
    v7 = v7
    switch (("\n\t" + "") + "\n\t") {
    case "a":
        var v19 uint8 = 31
        _ = v19
    default:
        v3 += 32767
    }
    fmt.Println(v7)
    v5 += (v5 * v2)
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint32 = 4294967295
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS2{F0: 128}
        v24 := GenS4{F0: (v5 ^ (v5 ^ 0)), F1: 14}
        v5 = (v2 | v5)
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v26 := GenS1{F0: (v3 / ((2) | 1))}
        var v27 uint16 = 42
        _ = fi25
        _ = v26
        _ = v27
    }
    v3 += v3
    v3 += 5
    var v28 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
