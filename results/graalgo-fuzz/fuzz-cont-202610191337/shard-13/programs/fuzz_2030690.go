package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int32; F1 string; F2 int8 }
type GenS3 struct { F0 bool; F1 string; F2 bool }
func genFn0(p0 int8, p1 uint64) uint32 {
    switch 32768 {
    case 752:
        p1 = p1
    case 128:
        v7 := GenS1{F0: 32767, F1: 9223372036854775807}
        _ = v7
    case 4294967295:
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := GenS2{F0: 0, F1: "\n\t", F2: p0}
                _ = fi9
                _ = v10
            }
            var v11 int = (-(fi8))
            v12 := []int{(fi8 % (((v11 - v11)) | 1)), v11, fi8, fi8}
            _ = fi8
            _ = v11
            _ = v12
        }
    }
    v13 := GenS2{F0: 1, F1: ("hello" + "a"), F2: p0}
    if (-2.25 >= ((-2.25 / 1.0) + 1.5)) {
        var v14 uint64 = 246
        p1 += (p1 << ((2) & 63))
        v15 := []int{}
        _ = v14
        _ = v15
    } else {
        if (!true) {
            switch "a" {
            case "":
                if true {
                    var v16 float64 = -2.25
                    v17 := []int{}
                    v18 := GenS3{F0: (int32(p1) >= 42), F1: "\n\t", F2: (4294967295 < 225)}
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    var v19 int64 = -1
                    if (255 >= 1) {
                        var v20 int16 = -32768
                        v21 := GenS3{F0: true, F1: ((("hello" + "a") + ("a" + ("hello" + "foo bar"))) + "hello"), F2: false}
                        _ = v20
                        _ = v21
                    }
                    _ = v19
                }
            case "a":
                switch "foo bar" {
                case "a":
                    fmt.Println(p0, p1)
                }
            case "\n\t":
                v22 := GenS2{F0: -2147483648, F1: "hello", F2: 100}
                _ = v22
            }
        } else {
            var v23 int = 9223372036854775807
            var v24 uint64 = p1
            _ = v23
            _ = v24
        }
        if (-1 < 341) {
            var v25 bool = false
            _ = v25
        } else {
            var v26 int8 = p0
            _ = v26
        }
        p0 -= (-(p0))
    }
    switch "" {
    case "\n\t":
        p0 = (p0 & p0)
    }
    _ = p0
    _ = p1
    _ = v13
    return 512
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]int)
    var v3 int32 = 275
    var v4 uint32 = 252
    var v5 uint32 = (^(v4))
    var v6 float32 = 1.5
    if ((true && (523 > -1)) || (((v6 / ((v6 * (v6 / v6)) + 0.1)) - 1e10) > ((v6 - 0.0) * v6))) {
        fmt.Println(v4)
    }
    var v27 uint16 = 255
    var v28 int64 = 323
    for fi29 := 0; fi29 < 10; fi29++ {
        fmt.Println(v6, v27, v5)
        var v30 float64 = 1.5
        fmt.Println(v27, v27, fi29)
        _ = fi29
        _ = v30
    }
    v1 += v1
    for fi31 := 0; fi31 < 8; fi31++ {
        switch "hello" {
        case "foo bar":
            v27 = v27
        default:
            var v32 uint16 = (^(v27))
            _ = v32
        }
        v4 = genFn0(-56, 722)
        _ = fi31
    }
    var v33 uint16 = 516
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 bool = false
        _ = fi34
        _ = v35
    }
    fmt.Println(v4, v3)
    if false {
        v6 = ((v6 - v6) / 1.5)
        v5 = (v4 | 4294967295)
        fmt.Println(v27, v1, v6)
    } else {
        v1 += 255
        v36 := GenS2{F0: (v3 >> ((18446744073709551615) & 63)), F1: ("foo bar" + "\n\t"), F2: 24}
        v37 := []int{525, 128, 0, 32768}
        _ = v36
        _ = v37
    }
    var v38 uint = 946
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
