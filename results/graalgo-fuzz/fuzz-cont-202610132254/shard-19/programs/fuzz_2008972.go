package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int16; F1 uint8; F2 bool }
type GenS4 struct { F0 uint32; F1 uint16; F2 int32 }
func genFn0() uint64 {
    var v5 uint16 = 32767
    v5 -= v5
    v6 := GenS1{F0: (v5 >> ((1) & 63)), F1: ((("hello" + "a") + ("a" + "a")) + "")}
    v7 := GenS1{F0: (v5 | v5), F1: "foo bar"}
    if (!true) {
        for fi8 := 0; fi8 < 9; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                if ("hello" >= "hello") {
                    var v10 uint8 = 127
                    var v11 int32 = 2147483647
                    v12 := []int{220, fi9, fi9, fi9}
                    _ = v10
                    _ = v11
                    _ = v12
                }
                v13 := []int{-2147483648, (^(fi9))}
                _ = fi9
                _ = v13
            }
            v5 += 42
            v5 = v5
            _ = fi8
        }
        switch -2147483648 {
        case 165:
            var v14 int16 = 42
            _ = v14
        case 127:
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := []int{}
                _ = fi15
                _ = v16
            }
        case 32768:
            fmt.Println(v5)
        default:
            var v17 int8 = 100
            _ = v17
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint = 141
    var v2 float64 = (1.5 * (float64(-2.25) / 1.0))
    var v3 int = -349
    var v4 int16 = -3
    if (42 == 42) {
        v18 := GenS1{F0: 1, F1: ""}
        _ = v18
    } else {
        v19 := GenS2{F0: 186, F1: 483}
        _ = v19
    }
    v20 := GenS3{F0: (v4 - (v4 * v4)), F1: 42, F2: ((v4 >> (((v1 & (^(v1)))) & 63)) > -32768)}
    switch "foo bar" {
    case "foo bar":
        switch v3 {
        case 32768:
            for fi21 := 0; fi21 < 5; fi21++ {
                fmt.Println(fi21, v4, v2)
                var v22 uint64 = 42
                _ = fi21
                _ = v22
            }
        case 42:
            v2 -= 1.5
        case 772:
            if false {
                v23 := GenS2{F0: 255, F1: -9223372036854775808}
                v24 := []int{v3, v3, 9223372036854775807, -685}
                _ = v23
                _ = v24
            } else {
                var v25 int16 = 1
                _ = v25
            }
        }
    default:
        if (0 > -90) {
            var v26 float32 = -2.25
            v27 := GenS4{F0: 411, F1: 0, F2: 780}
            _ = v26
            _ = v27
        } else {
            var v28 int64 = 42
            _ = v28
        }
    }
    switch "foo bar" {
    case "a":
        var v29 uint32 = 2147483648
        _ = v29
    case "hello":
        v30 := []int{}
        _ = v30
    case "foo bar":
        fmt.Println(v1)
    }
    switch (v3 | (v3 | 564)) {
    case 32768:
        var v31 uint16 = 507
        _ = v31
    case 782:
        var v32 uint8 = 127
        _ = v32
    }
    v33 := []int{}
    v34 := GenS2{F0: 1, F1: -9223372036854775808}
    var v35 int16 = 32767
    v36 := GenS2{F0: 187, F1: -74}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
