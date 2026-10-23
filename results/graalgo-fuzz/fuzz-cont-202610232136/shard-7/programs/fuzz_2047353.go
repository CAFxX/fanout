package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int16; F2 int64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0(p0 uint64) int16 {
    v6 := []int{}
    fmt.Println(p0)
    p0 = (p0 >> ((453) & 63))
    switch ("" + "a") {
    case "\n\t":
        _ = len(v6)
    default:
        var v7 int32 = 2147483647
        _ = v7
    }
    if (309 <= 65535) {
        p0 = p0
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 float64 = (-2.25 + float64(-2.25))
            _ = fi8
            _ = v9
        }
        if true {
            for fi10 := 0; fi10 < 10; fi10++ {
                p0 += (^(p0))
                _ = fi10
            }
            var v11 int64 = 0
            var v12 uint32 = 1
            _ = v11
            _ = v12
        }
    } else {
        v6 = append(v6, 335)
        var v13 uint = 18446744073709551615
        if ((false || false) && true) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = v13
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1() int {
    v14 := []int{593, 9223372036854775807, 2147483647}
    for fi15 := 0; fi15 < 6; fi15++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        var v16 bool = (float64(((0.1 * -2.25) * 0.0)) == 0.0)
        v14 = append(v14, (-(fi15)))
        _ = fi15
        _ = v16
    }
    if (39 >= -93) {
        var v17 uint = 4294967295
        var v18 uint16 = 894
        switch 83 {
        case 790:
            fmt.Println(v17, v17)
        }
        _ = v17
        _ = v18
    } else {
        var v19 float32 = (1.5 / (1e10 * float32(float64(3.14159))))
        if (float32(((0.0 * 1e10) - (-2.25 / float64((0.1 + ((0.1 + (v19 / v19)) + ((1.5 + ((0.0 - 1e10) - 1.5)) + v19))))))) != -2.25) {
            var v20 int = -2147483648
            _ = v20
        } else {
            for fi21 := 0; fi21 < 20; fi21++ {
                fmt.Println(fi21, fi21)
                _ = fi21
            }
        }
        var v22 string = "hello"
        _ = v19
        _ = v22
    }
    _ = v14
    return 166
}

func main() {
    var v1 string = ("hello" + "foo bar")
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 float64 = 1e10
    v5 := []int{}
    v23 := GenS1{F0: -354, F1: 1, F2: 32768}
    v4 += 0.0
    v4 = (v4 + (-2.25 + v4))
    v24 := v3[-9223372036854775808]
    v4 -= v4
    for fi25 := 0; fi25 < 8; fi25++ {
        if (42 == -2147483648) {
            var v26 int = fi25
            _ = v26
        } else {
            v4 -= 0.0
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 int8 = 42
                _ = fi27
                _ = v28
            }
        }
        var v29 uint32 = 678
        _ = fi25
        _ = v29
    }
    if (true && (!true)) {
        v30 := GenS2{F0: 2147483647}
        _ = v30
    }
    v3[genFn1()] = (v24 - 0)
    v31 := GenS3{F0: (v1 <= "\n\t")}
    if false {
        switch v24 {
        case 128:
            if (true && (2 > 42)) {
                var v32 uint64 = 0
                var v33 float64 = v4
                _ = v32
                _ = v33
            }
        case 9223372036854775807:
            if true {
                v34 := GenS4{F0: 517, F1: -396, F2: 845}
                v5 = append(v5, (v24 >> ((0) & 63)))
                var v35 float32 = -2.25
                _ = v34
                _ = v35
            }
        default:
            var v36 int32 = 0
            _ = v36
        }
        v37 := GenS5{F0: 1, F1: 0}
        _ = v37
    } else {
        v24 = v24
    }
    if (0.1 < (v4 + 3.14159)) {
        var v38 bool = (genFn1() >= -789)
        _ = v38
    } else {
        v39 := GenS5{F0: -17, F1: 128}
        _ = v39
    }
    var v40 int64 = -1
    v41 := GenS6{F0: 151, F1: 42, F2: 92}
    v42 := GenS3{F0: ((2 > 2) || (float64(-2.25) != (3.14159 / v4)))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
