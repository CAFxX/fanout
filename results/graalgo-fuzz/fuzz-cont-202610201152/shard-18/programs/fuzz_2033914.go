package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 int }
func genFn0() int32 {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{1, 846, 2147483647, (fi7 ^ (fi7 / ((1) | 1)))}
        _ = fi7
        _ = v8
    }
    var v9 string = "a"
    fmt.Println(v9)
    _ = v9
    return -1
}

func main() {
    var v1 uint = 255
    var v2 bool = ("hello" <= "\n\t")
    var v3 int16 = 973
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 0
    v6 := []uint32{319}
    var v10 int64 = 9223372036854775807
    if len(v6) > 0 {
        v6[0] = 401
    }
    switch "\n\t" {
    case "\n\t":
        v11 := GenS1{F0: v1, F1: 857}
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint8 = (v5 / ((109) | 1))
        _ = fi12
        _ = v13
    }
    if ("" != "foo bar") {
        v14 := GenS1{F0: v1, F1: 18446744073709551615}
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(v1, v1, fi15)
            _ = fi15
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint8 = v5
            var v18 uint64 = 1
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v14
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(v3, fi19, fi19)
            for fi20 := 0; fi20 < 13; fi20++ {
                v10 = (v10 % ((v10) | 1))
                var v21 int64 = v10
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    }
    v10 += v10
    v1 -= (v1 ^ (v1 - 65535))
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS2{F0: v5, F1: 2147483648}
        switch (((("" + "a") + "") + "foo bar") + "a") {
        case "a":
            v2 = v2
        default:
            v10 = -9223372036854775808
        }
        var v24 int32 = -263
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 float64 = 3.14159
    var v26 int8 = 127
    v25 -= -2.25
    v27 := GenS3{F0: (("" + ("\n\t" + ("\n\t" + ""))) + "foo bar"), F1: v2, F2: 132}
    var v28 int64 = 42
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 uint64 = 211
        var v31 uint16 = 0
        _ = fi29
        _ = v30
        _ = v31
    }
    v2 = (!(v25 != (0.0 / (v25 + v25))))
    var v32 uint16 = 363
    v33 := GenS3{F0: "hello", F1: v2, F2: 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
