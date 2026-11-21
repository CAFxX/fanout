package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint32 }
type GenS2 struct { F0 int; F1 bool; F2 int16 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int; F1 int64; F2 bool }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int }

func main() {
    var v1 string = "foo bar"
    var v2 float64 = ((((((1.5 / 1.0) + (1e10 + float64((1.5 * 0.1)))) + float64(1e10)) - ((0.0 - 0.1) * 0.1)) / float64(((1.5 / (1e10 / (3.14159 + 0.0))) + -2.25))) + 1.5)
    var v3 uint64 = 18446744073709551615
    var v4 float32 = 3.14159
    var v5 float32 = v4
    if true {
        fmt.Println(v5, v3)
        v6 := []int{}
        fmt.Println(v5)
        _ = v6
    } else {
        v7 := GenS1{F0: 0, F1: (((("foo bar" + (v1 + "\n\t")) + ("" + "hello")) + ((((v1 + (("foo bar" + v1) + v1)) + "hello") + v1) + v1)) + ""), F2: 2147483648}
        var v8 float32 = ((1.5 * 3.14159) - (v5 - 1e10))
        v9 := []int{-804, 1}
        _ = v7
        _ = v8
        _ = v9
    }
    v3 = v3
    var v10 uint8 = 127
    fmt.Println(v1, v5)
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            var v12 uint16 = 65535
            fmt.Println(v10, v1)
            v5 = v5
            _ = v12
        } else {
            v4 = (v5 + v5)
        }
        v3 = 4294967295
        var v13 int8 = -77
        _ = fi11
        _ = v13
    }
    var v14 int8 = -128
    fmt.Println(v14, v14, v10)
    switch -935 {
    case 9223372036854775807:
        var v15 float64 = v2
        _ = v15
    case -187:
        v16 := []int{1, 2147483647, 1}
        _ = v16
    default:
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := []int{fi17, (fi17 % ((fi17) | 1)), 9223372036854775807}
            v1 = "a"
            _ = fi17
            _ = v18
        }
    }
    if false {
        v4 = (((3.14159 * (v4 / v5)) * 3.14159) + float32((3.14159 * v2)))
        v19 := []int{42, -9223372036854775808}
        v20 := GenS1{F0: 0, F1: v1, F2: 1}
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := GenS1{F0: 98, F1: v1, F2: 2}
            fmt.Println(v4, v1, fi21)
            _ = fi21
            _ = v22
        }
        v23 := []int{-820}
        v4 -= (v5 / v5)
        _ = v23
    }
    v24 := GenS2{F0: 42, F1: false, F2: 1}
    var v25 bool = (4294967295 > 436)
    v3 = (v3 & v3)
    if (v3 <= v3) {
        var v26 int = 0
        var v27 uint64 = v3
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 10; fi28++ {
            v29 := GenS3{F0: v1, F1: (v10 | (^(v10)))}
            var v30 int8 = -31
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    v14 -= v14
    switch "hello" {
    case "":
        v10 = v10
    default:
        if (609 == 0) {
            var v31 uint64 = (v3 | 4294967295)
            _ = v31
        } else {
            switch 0 {
            case 127:
                if false {
                    if (!v25) {
                        v32 := []int{9223372036854775807, 698, -9223372036854775808}
                        var v33 uint = 4294967295
                        _ = v32
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 1; fi34++ {
                        v5 = (v4 / (v4 * (v4 * 3.14159)))
                        v35 := GenS4{F0: fi34, F1: 1, F2: ((634 != 0) || true)}
                        _ = fi34
                        _ = v35
                    }
                } else {
                    v36 := GenS5{F0: 247, F1: v14}
                    _ = v36
                }
            case 0:
                v37 := GenS6{F0: -458}
                _ = v37
            case 32767:
                v38 := GenS7{F0: 127}
                _ = v38
            }
        }
    }
    v3 += 938
    if (uint64(v10) >= (v3 & (v3 >> ((2) & 63)))) {
        v39 := []int{983, 2147483647}
        _ = v39
    } else {
        if true {
            v14 = (^(v14))
            var v40 int16 = 1
            v2 = (1e10 / ((0.1 - (((1.5 * (v2 + float64(float32((((v2 - (v2 / float64((0.1 - (v4 * v4))))) * v2) / v2))))) - v2) * ((-2.25 * float64((v4 * 1e10))) - ((0.1 * v2) * v2)))) / 1.0))
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
