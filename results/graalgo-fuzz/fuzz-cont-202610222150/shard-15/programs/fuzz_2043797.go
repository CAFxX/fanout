package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int32 }
type GenS2 struct { F0 uint; F1 int32; F2 int32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int8 }
func genFn0(p0 int64) uint16 {
    v6 := GenS1{F0: 255, F1: (("hello" + "hello") + ""), F2: 42}
    switch -9223372036854775808 {
    case -2147483648:
        var v7 int16 = 1
        _ = v7
    case 889:
        p0 -= p0
    case 128:
        if true {
            p0 = -9223372036854775808
            v8 := []int{}
            if len(v8) > 0 {
                v8[0] = 128
            }
            _ = v8
        } else {
            fmt.Println(p0)
            if true {
                v9 := GenS1{F0: 211, F1: "foo bar", F2: 0}
                _ = v9
            }
        }
    default:
        v10 := []int{2147483647, 42, 434}
        _ = v10
    }
    p0 += p0
    _ = p0
    _ = v6
    return 884
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]int)
    var v3 float32 = (1.5 / 0.1)
    var v4 uint = 65535
    v5 := make(map[int]string)
    v2[32768] = -575
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS2{F0: v4, F1: (^(v1)), F2: v1}
        _ = fi11
        _ = v12
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        v3 = 0.1
    }
    var v13 int16 = 1
    v3 = ((v3 / ((((v3 + v3) / (3.14159 + ((float32(1.5) - v3) + v3))) / 1e10) * -2.25)) / (v3 + (((v3 / v3) + v3) - v3)))
    v14 := v5[int(v13)]
    switch 127 {
    case -751:
        v13 = (-(v13))
    }
    var v15 uint16 = genFn0(-1)
    v4 += (v4 ^ 1)
    v2[-350] = 32768
    var v16 int64 = -1
    var v17 int = -9223372036854775808
    switch v17 {
    case -1:
        v18 := GenS1{F0: genFn0((v16 % (((v16 << ((154) & 63))) | 1))), F1: v14, F2: v1}
        _ = v18
    }
    v14 = ""
    switch v17 {
    case 128:
        for fi19 := 0; fi19 < 16; fi19++ {
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := []int{}
                fmt.Println(fi20, v13, v13)
                var v22 string = (v14 + "a")
                _ = fi20
                _ = v21
                _ = v22
            }
            v14 = v14
            v1 += (v1 | v1)
            _ = fi19
        }
    default:
        var v23 int16 = (-(v13))
        _ = v23
    }
    v24 := GenS3{F0: 252, F1: v17}
    if false {
        if (!false) {
            var v25 uint32 = 0
            var v26 float64 = -2.25
            _ = v25
            _ = v26
        }
        switch v14 {
        case "":
            fmt.Println(v3)
        case "foo bar":
            v27 := []int{-9223372036854775808, (v17 - 32767), (v17 + v17)}
            _ = v27
        default:
            v4 += 2
        }
    } else {
        var v28 uint64 = 0
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        for fi30 := 0; fi30 < 10; fi30++ {
            v31 := GenS4{F0: -67}
            v3 = 0.0
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
