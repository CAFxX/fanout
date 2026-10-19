package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 491
    var v5 bool = false
    var v6 float64 = 1.5
    if v5 {
        if (2147483647 == (v4 << ((65535) & 63))) {
            var v7 int32 = -761
            var v8 uint16 = 255
            if (-1 != 9223372036854775807) {
                v2 = (v2 - v2)
            }
            _ = v7
            _ = v8
        } else {
            if (((!(42 < -1)) && (1 <= 103)) && (-988 >= 1)) {
                v4 = v4
                v5 = v5
                var v9 uint64 = 4294967295
                _ = v9
            }
            v4 += v4
            if v5 {
                if (3 != 32) {
                    _ = len(v1)
                }
                var v10 float32 = (float32((1.5 + (float64((0.1 / (float32(1e10) * float32(((((-2.25 / 1e10) - 3.14159) / v6) + ((v6 - float64((1e10 + (3.14159 * (1e10 + 1e10))))) - v6)))))) - v6))) + -2.25)
                _ = v10
            }
        }
    } else {
        v11 := GenS1{F0: 300, F1: 42, F2: 127}
        v6 = 0.0
        v12 := GenS1{F0: 4294967295, F1: 825, F2: 127}
        _ = v11
        _ = v12
    }
    v6 -= 1e10
    var v13 uint8 = 2
    if false {
        if true {
            if v5 {
                fmt.Println(v2, v2)
            } else {
                v14 := []int{}
                v15 := GenS2{F0: 0}
                _ = v14
                _ = v15
            }
            var v16 uint8 = v13
            switch 128 {
            case 4294967295:
                var v17 int = 127
                _ = v17
            case 1:
                var v18 uint16 = 0
                _ = v18
            case 9223372036854775807:
                var v19 float32 = -2.25
                _ = v19
            }
            _ = v16
        }
        var v20 uint32 = 2147483648
        _ = v20
    }
    fmt.Println(v6)
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 string = ""
            var v24 float64 = v6
            v23 = v23
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    fmt.Println(v4, v6, v13)
    switch "foo bar" {
    case "hello":
        var v25 float32 = (3.14159 * (-2.25 * 1e10))
        _ = v25
    case "\n\t":
        if false {
            var v26 int64 = 42
            _ = v26
        }
    case "":
        var v27 int32 = 42
        _ = v27
    default:
        v28 := GenS2{F0: 127}
        _ = v28
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := []int{fi29, 0, fi29}
        _ = fi29
        _ = v30
    }
    v31 := GenS3{F0: 2, F1: true}
    var v32 uint = uint(v13)
    var v33 uint = 593
    fmt.Println(v5, v5, v5)
    var v34 int = -803
    v35 := GenS4{F0: int64(v34), F1: 255, F2: (^(v2))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
