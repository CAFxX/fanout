package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint8 }

func main() {
    var v1 float32 = float32(1e10)
    var v2 int = 32768
    var v3 float64 = 3.14159
    v4 := GenS1{F0: 255, F1: -845, F2: 42}
    switch (v2 * v2) {
    case 128:
        fmt.Println(v3)
    case 919:
        v5 := GenS1{F0: 42, F1: 9223372036854775807, F2: 778}
        _ = v5
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v3 += -2.25
        v7 := []int{9223372036854775807, fi6}
        v2 = (fi6 - (-(v2)))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            var v9 uint64 = 284
            v3 = v3
            _ = v9
        } else {
            var v10 uint16 = 65535
            var v11 int64 = 0
            _ = v10
            _ = v11
        }
        var v12 uint32 = 42
        var v13 int16 = -1
        _ = fi8
        _ = v12
        _ = v13
    }
    var v14 uint64 = 18446744073709551615
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    var v17 float64 = 0.0
    var v18 int32 = -1
    v19 := GenS2{F0: 0, F1: (-(v18)), F2: 255}
    switch v2 {
    case 857:
        v20 := []int{v2, (^(v2)), v2}
        _ = v20
    case 751:
        v21 := []int{128, 2147483647, -1, (v2 % ((v2) | 1))}
        _ = v21
    default:
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := []int{(-(fi22)), fi22}
            v18 -= (v18 / (((v18 * 42)) | 1))
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(v2)
    var v24 int32 = (v18 & 1)
    switch "foo bar" {
    case "hello":
        v25 := GenS2{F0: 2, F1: -1, F2: 255}
        _ = v25
    default:
        switch ((("" + "") + "foo bar") + "") {
        case "foo bar":
            if true {
                var v26 uint16 = 32767
                v27 := []int{127, 0, (v2 % ((v2) | 1)), 2147483647}
                _ = v26
                _ = v27
            }
        case "a":
            if (0 != 1) {
                switch "foo bar" {
                case "\n\t":
                    if false {
                        v1 += v1
                    } else {
                        v28 := GenS2{F0: 755, F1: 2147483647, F2: 235}
                        v29 := []int{4294967295, -9223372036854775808}
                        _ = v28
                        _ = v29
                    }
                case "foo bar":
                    switch 2147483647 {
                    case -162:
                        var v30 uint8 = 224
                        _ = v30
                    default:
                        var v31 float32 = 1e10
                        _ = v31
                    }
                default:
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 float32 = v1
                        v34 := []int{}
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                }
            }
        case "hello":
            v35 := []int{v2, v2, 44, 901}
            _ = v35
        }
    }
    v36 := GenS2{F0: 1, F1: 2147483647, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v4
}
