package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8, p1 int) uint8 {
    v7 := []int{9223372036854775807, (p1 | 2147483647)}
    v8 := GenS1{F0: (p1 & 127)}
    v9 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1() uint {
    var v10 uint16 = 1
    v10 += v10
    if (2 > 0) {
        var v11 uint32 = 0
        v11 = (^(v11))
        if false {
            var v12 int8 = 0
            _ = v12
        }
        _ = v11
    }
    _ = v10
    return 841
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 100
    var v3 int16 = 1000
    var v4 uint8 = 145
    v5 := make(map[int]bool)
    var v6 float32 = (1.5 / 1.5)
    switch "foo bar" {
    case "foo bar":
        if (!true) {
            delete(v1, 453)
            v6 = ((v6 + (v6 - (v6 - float32(1e10)))) / (0.0 - float32(0.1)))
        } else {
            var v13 int32 = -976
            _ = v13
        }
    }
    switch 0 {
    case -2147483648:
        var v14 string = ("foo bar" + ("hello" + ("" + "\n\t")))
        _ = v14
    case 1:
        v15 := GenS1{F0: 9223372036854775807}
        _ = v15
    default:
        switch 9223372036854775807 {
        case -9223372036854775808:
            fmt.Println(v4, v4, v6)
        default:
            fmt.Println(v2, v4)
        }
    }
    var v16 int32 = 42
    var v17 uint = 1
    for fi18 := 0; fi18 < 5; fi18++ {
        v3 -= 1000
        var v19 uint16 = 42
        _ = fi18
        _ = v19
    }
    var v20 uint32 = 681
    v4 = genFn0(v4, -73)
    if (2 <= 121) {
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v3)
            v22 := GenS1{F0: fi21}
            _ = fi21
            _ = v22
        }
        if false {
            v23 := []int{0, 127, 4294967295}
            _ = v23
        }
    }
    var v24 int64 = 9223372036854775807
    switch "a" {
    case "foo bar":
        var v25 float32 = (v6 * (0.1 - (-2.25 * v6)))
        _ = v25
    default:
        fmt.Println(v16)
    }
    if ((!true) && ((v17 + v17) >= (v17 >> (((v17 + v17)) & 63)))) {
        var v26 bool = (true || (true && true))
        _ = v26
    } else {
        fmt.Println(v17)
        var v27 bool = true
        var v28 float64 = (1.5 / 1.5)
        _ = v27
        _ = v28
    }
    v20 = v20
    v2 += (^(v2))
    v24 = v24
    for fi29 := 0; fi29 < 8; fi29++ {
        fmt.Println(v24, fi29)
        v2 = int8(v3)
        _ = fi29
    }
    var v30 float64 = (0.0 + 0.0)
    for fi31 := 0; fi31 < 19; fi31++ {
        fmt.Println(v4, v20, v6)
        switch fi31 {
        case -411:
            var v32 uint8 = v4
            _ = v32
        case -1:
            if (true || (!true)) {
                fmt.Println(v17)
                v6 -= v6
                v17 = (v17 ^ genFn1())
            } else {
                v33 := GenS2{F0: 557, F1: v17, F2: (^(v16))}
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v3, v6)
                default:
                    v34 := GenS3{F0: (v24 / ((987) | 1)), F1: 0}
                    _ = v34
                }
                _ = v33
            }
        default:
            v35 := GenS4{F0: (fi31 / (((fi31 % (((^(fi31))) | 1))) | 1))}
            _ = v35
        }
        var v36 int16 = int16(v20)
        _ = fi31
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        var v38 uint = v17
        _ = fi37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
