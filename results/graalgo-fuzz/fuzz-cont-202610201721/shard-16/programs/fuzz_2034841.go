package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
func genFn0() uint {
    var v7 bool = false
    switch 127 {
    case -756:
        var v8 float32 = (1e10 - (3.14159 / 1.5))
        _ = v8
    default:
        v7 = true
    }
    v9 := []int{32768, -9223372036854775808}
    _ = v7
    _ = v9
    return 1
}
func genFn1() int8 {
    var v10 uint64 = 534
    if false {
        if false {
            v11 := GenS1{F0: 1}
            v10 -= 18446744073709551615
            _ = v11
        } else {
            v10 -= 18446744073709551615
            v12 := GenS2{F0: "\n\t", F1: 449, F2: 0}
            _ = v12
        }
    }
    _ = v10
    return -102
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -128
    v3 := make(map[int]uint32)
    var v4 float32 = (((1.5 / 1.0) * (0.0 / (0.0 - ((1.5 - 1e10) / (0.1 * 0.1))))) * -2.25)
    var v5 int32 = -995
    var v6 uint8 = 75
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := []int{}
        _ = fi13
        _ = v14
    }
    if true {
        fmt.Println(v6, v2, v2)
    } else {
        v15 := GenS3{F0: 1, F1: 42}
        v3[-2147483648] = 42
        v2 += v2
        _ = v15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            switch (fi16 >> ((0) & 63)) {
            case 127:
                v18 := GenS4{F0: 0}
                _ = v18
            default:
                var v19 uint = 1
                _ = v19
            }
            _ = fi17
        }
        v20 := GenS5{F0: (v6 ^ v6)}
        _ = fi16
        _ = v20
    }
    v21 := GenS4{F0: 1000}
    for fi22 := 0; fi22 < 2; fi22++ {
        v1[42] = ""
        for fi23 := 0; fi23 < 7; fi23++ {
            v4 = v4
            _ = fi23
        }
        _ = fi22
    }
    v6 -= 121
    for fi24 := 0; fi24 < 5; fi24++ {
        v2 = -81
        _ = fi24
    }
    var v25 int64 = 9223372036854775807
    v25 = v25
    if (1 != 255) {
        v26 := GenS4{F0: -1}
        fmt.Println(v2, v2, v25)
        _ = v26
    } else {
        v2 = (v2 + (v2 ^ v2))
        switch int(v5) {
        case 127:
            switch "\n\t" {
            case "a":
                if true {
                    switch ("foo bar" + ("a" + "foo bar")) {
                    case "foo bar":
                        fmt.Println(v2)
                    }
                    if false {
                        v3[32767] = 488
                    }
                }
            default:
                v27 := []int{0, 475, 128, 42}
                _ = v27
            }
        case 0:
            var v28 int32 = v5
            _ = v28
        }
    }
    fmt.Println(v4, v5, v4)
    v25 += (v25 * v25)
    var v29 int16 = -1
    fmt.Println(v5, v6, v25)
    var v30 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
