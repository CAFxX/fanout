package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 uint8; F1 int; F2 int32 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint = 599
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: "\n\t", F1: "a"}
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int = 32767
        if (false || (true && true)) {
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(fi10)
                _ = fi12
            }
            v11 = (v11 & v11)
        } else {
            v11 += fi10
            v13 := []int{9223372036854775807, fi10, (fi10 % ((-459) | 1))}
            _ = v13
        }
        switch v11 {
        case 42:
            if (1.5 > 0.1) {
                v11 = fi10
            }
        case -9223372036854775808:
            var v14 float32 = (1.5 * 3.14159)
            _ = v14
        case 127:
            v15 := []int{(v11 | fi10), (fi10 / ((v11) | 1)), (v11 >> ((18446744073709551615) & 63)), fi10}
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v16 int16 = -824
    _ = v16
    _ = v9
    return 201
}

func main() {
    var v1 int8 = -128
    var v2 float32 = (1.5 - 1e10)
    var v3 int16 = 42
    var v4 uint64 = 42
    v5 := []uint32{2147483647, 44}
    v6 := make(map[string]int)
    v1 -= 0
    v17 := GenS2{F0: -1, F1: 2147483647, F2: "foo bar"}
    switch "foo bar" {
    case "foo bar":
        if true {
            v18 := GenS3{F0: 128, F1: 4294967295, F2: 578}
            v1 = (v1 + -1)
            _ = len(v6)
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v3 -= (v3 & (^(v3)))
        _ = fi19
    }
    v3 -= (^(v3))
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 float64 = 0.1
        _ = fi20
        _ = v21
    }
    var v22 float32 = (v2 * v2)
    var v23 int = -2147483648
    v2 += 0.1
    var v24 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
