package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0(p0 uint32, p1 int8, p2 uint32) int8 {
    p1 = 1
    if true {
        v6 := GenS1{F0: 349, F1: false, F2: 2}
        _ = v6
    } else {
        p0 += p2
    }
    v7 := GenS1{F0: -9223372036854775808, F1: (p1 <= 0), F2: (p2 / ((p0) | 1))}
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(p2, p1, fi8)
        _ = fi8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(p0, fi9)
        v10 := []int{(fi9 << ((714) & 63)), -1, fi9}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    var v1 int32 = -2147483648
    var v2 float64 = 1.5
    var v3 float32 = 3.14159
    var v4 bool = true
    v5 := make(map[int]string)
    v11 := GenS2{F0: 1, F1: 42}
    if false {
        if ((!(0 < -896)) && v4) {
            var v12 uint64 = 2
            v13 := GenS2{F0: 32767, F1: 42}
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 int32 = -2147483648
                v16 := []int{32767, fi14}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v12
            _ = v13
        }
    }
    switch -2147483648 {
    case 32767:
        if ((!v4) || ((2147483647 <= 9223372036854775807) || v4)) {
            fmt.Println(v1, v2, v4)
            var v17 int16 = 830
            _ = len(v5)
            _ = v17
        }
    }
    switch "hello" {
    case "\n\t":
        switch 4294967295 {
        case 9223372036854775807:
            v18 := GenS2{F0: 1000, F1: 42}
            _ = v18
        }
    case "foo bar":
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 string = "foo bar"
            delete(v5, (fi19 >> ((255) & 63)))
            var v21 int8 = 100
            _ = fi19
            _ = v20
            _ = v21
        }
    default:
        v22 := GenS3{F0: 2, F1: (2 > 54)}
        _ = v22
    }
    v23 := GenS4{F0: "a", F1: 2147483648}
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS1{F0: fi24, F1: true, F2: 329}
        v26 := GenS2{F0: 1, F1: 2}
        var v27 uint16 = 182
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 uint16 = 65535
    v29 := GenS4{F0: "a", F1: 2147483648}
    v3 += v3
    v3 -= (v3 * (float32(float64(v3)) + v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
