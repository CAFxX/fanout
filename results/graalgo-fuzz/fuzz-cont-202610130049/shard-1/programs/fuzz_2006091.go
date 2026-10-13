package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint; F2 int }
type GenS4 struct { F0 uint64; F1 bool }
func genFn0() string {
    var v7 uint16 = 2
    v8 := []int{-2147483648, 2147483647, 2147483647}
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: 32767, F1: 0}
        v11 := GenS2{F0: fi9}
        var v12 uint32 = 1
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v7 = v7
    var v13 int16 = 42
    _ = v13
    _ = v7
    _ = v8
    return ("\n\t" + "\n\t")
}
func genFn1(p0 int16, p1 uint) int {
    if false {
        v14 := GenS2{F0: -640}
        v15 := []int{42, 133}
        var v16 string = "a"
        _ = v14
        _ = v15
        _ = v16
    } else {
        v17 := []int{2147483647}
        _ = v17
    }
    fmt.Println(p0, p1)
    p0 += -32768
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 32767
    var v3 string = ""
    v4 := make(map[uint32]string)
    v5 := []string{}
    var v6 float32 = (0.1 + ((3.14159 / 0.1) * 1e10))
    if len(v5) > 0 {
        v5[0] = ((v3 + v3) + (v3 + v3))
    }
    if true {
        var v18 float32 = 3.14159
        _ = v18
    } else {
        fmt.Println(v2, v6)
        var v19 uint64 = 616
        _ = v19
    }
    v20 := GenS3{F0: false, F1: 381, F2: int(v2)}
    v21 := GenS2{F0: 128}
    v6 += v6
    switch 4294967295 {
    case -1:
        _ = len(v1)
    case 127:
        if false {
            var v22 int8 = 0
            _ = v22
        }
    default:
        var v23 int32 = 122
        _ = v23
    }
    var v24 int = 0
    var v25 int = v24
    var v26 uint = uint(v2)
    switch genFn1(-32768, (v26 - (v26 + 47))) {
    case -768:
        if (true || (!true)) {
            var v27 string = v3
            _ = v27
        } else {
            v24 = -9223372036854775808
            var v28 int16 = 42
            for fi29 := 0; fi29 < 1; fi29++ {
                v30 := GenS3{F0: true, F1: (v26 | (v26 + (v26 ^ (v26 | 18446744073709551615)))), F2: (^(fi29))}
                _ = fi29
                _ = v30
            }
            _ = v28
        }
    default:
        var v31 int8 = -13
        _ = v31
    }
    v24 = (v24 / ((v24) | 1))
    for fi32 := 0; fi32 < 4; fi32++ {
        v25 += (v25 + (-(v24)))
        _ = fi32
    }
    if (-9223372036854775808 >= 386) {
        var v33 uint64 = 2
        _ = v33
    } else {
        fmt.Println(v25, v2)
        v4[4294967295] = "\n\t"
        fmt.Println(v25, v24)
    }
    var v34 int16 = 0
    v26 += 1
    switch v3 {
    case "hello":
        fmt.Println(v24)
    case "a":
        var v35 int32 = 42
        _ = v35
    case "\n\t":
        for fi36 := 0; fi36 < 12; fi36++ {
            switch "\n\t" {
            case "foo bar":
                if (-128 > 1) {
                    var v37 int = (v24 ^ genFn1(v34, 256))
                    _ = v37
                }
            default:
                v38 := GenS4{F0: 2, F1: (!false)}
                _ = v38
            }
            _ = fi36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
