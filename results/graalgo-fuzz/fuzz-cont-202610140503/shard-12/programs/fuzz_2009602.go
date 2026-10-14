package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint16 }
func genFn0() int32 {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{(fi6 + 924), (fi6 ^ 32767), -965, fi6}
        var v8 float64 = 1e10
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 166
    var v10 int16 = 586
    fmt.Println(v9, v10)
    v11 := []int{127, -9223372036854775808}
    _ = v10
    _ = v11
    _ = v9
    return -222
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 2147483647
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    v5 := []int{-2147483648, 359, 32768, 1}
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            v2 = (v2 & (v2 >> ((65535) & 63)))
            _ = fi13
        }
        if true {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint8 = 128
                v2 -= int32(v15)
                v16 := GenS1{F0: 2147483648, F1: 0, F2: 255}
                _ = fi14
                _ = v15
                _ = v16
            }
            v2 = v2
        }
        _ = fi12
    }
    switch ("a" + "a") {
    case "\n\t":
        var v17 float64 = 1e10
        _ = v17
    default:
        fmt.Println(v2)
    }
    if (9223372036854775807 > 21) {
        if ((32767 != -1) || true) {
            var v18 int64 = 515
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            v2 -= v2
            _ = v18
        } else {
            var v19 int64 = 9223372036854775807
            var v20 int64 = v19
            _ = v19
            _ = v20
        }
    } else {
        v2 = 755
        var v21 float64 = 0.0
        for fi22 := 0; fi22 < 7; fi22++ {
            if false {
                switch "foo bar" {
                case "hello":
                    var v23 uint16 = 65535
                    _ = v23
                case "":
                    v24 := GenS1{F0: 757, F1: 390, F2: 0}
                    _ = v24
                case "\n\t":
                    _ = len(v5)
                }
            } else {
                var v25 float64 = v21
                var v26 int8 = 19
                _ = v25
                _ = v26
            }
            v21 += v21
            fmt.Println(v21, fi22, fi22)
            _ = fi22
        }
        _ = v21
    }
    v27 := GenS2{F0: 42}
    v28 := GenS3{F0: 255}
    v29 := GenS1{F0: 4294967295, F1: 65535, F2: 255}
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 uint = 316
        _ = fi30
        _ = v31
    }
    v2 -= (v2 ^ 0)
    if ((0.0 < ((3.14159 * (-2.25 / 1.0)) + float64(1.5))) && ((false && (!false)) && (!(!(1 != 258))))) {
        switch "\n\t" {
        case "foo bar":
            var v32 int8 = -1
            _ = v32
        case "a":
            v33 := GenS1{F0: 377, F1: 65535, F2: 32767}
            _ = v33
        }
        switch 128 {
        case 127:
            var v34 uint64 = 0
            _ = v34
        case -1:
            v35 := v1[4294967295]
            _ = v35
        case 212:
            if (322 >= 2147483648) {
                v2 -= v2
                v36 := GenS4{F0: -1, F1: 720}
                var v37 uint = 0
                _ = v36
                _ = v37
            }
        default:
            fmt.Println(v2)
        }
    }
    var v38 float32 = (3.14159 + -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
