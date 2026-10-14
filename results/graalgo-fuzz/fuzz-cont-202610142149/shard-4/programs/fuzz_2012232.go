package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int; F2 uint16 }
func genFn0(p0 uint16) uint {
    v5 := GenS1{F0: "a", F1: p0, F2: (("hello" + ("a" + "")) + "a")}
    p0 -= 32767
    var v6 int = -9223372036854775808
    if (true || true) {
        v7 := []int{v6, (v6 << ((18446744073709551615) & 63)), (-(v6)), 0}
        if true {
            if false {
                for fi8 := 0; fi8 < 19; fi8++ {
                    var v9 uint64 = 1
                    p0 += (p0 % ((p0) | 1))
                    _ = fi8
                    _ = v9
                }
                if (!false) {
                    for fi10 := 0; fi10 < 6; fi10++ {
                        for fi11 := 0; fi11 < 11; fi11++ {
                            v7 = append(v7, v6)
                            fmt.Println(v6, p0, fi10)
                            _ = fi11
                        }
                        _ = fi10
                    }
                    var v12 int16 = 107
                    _ = v12
                }
                var v13 float64 = (0.0 * 1e10)
                _ = v13
            } else {
                var v14 uint = 204
                _ = v14
            }
            fmt.Println(v6)
        }
        fmt.Println(p0)
        _ = v7
    }
    if (!true) {
        if false {
            v6 = v6
        }
        var v15 uint32 = 851
        _ = v15
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1() uint8 {
    switch "hello" {
    case "a":
        var v16 int = 231
        _ = v16
    case "hello":
        var v17 int8 = 127
        _ = v17
    case "":
        var v18 uint32 = 2147483648
        _ = v18
    }
    var v19 uint32 = 199
    switch 149 {
    case -9223372036854775808:
        var v20 int32 = -1
        _ = v20
    case -456:
        var v21 string = ""
        _ = v21
    case 901:
        v19 -= (v19 >> ((18446744073709551615) & 63))
    default:
        var v22 bool = (2 > 898)
        _ = v22
    }
    if true {
        fmt.Println(v19)
        v23 := []int{4294967295, 1, -9223372036854775808, 32768}
        _ = v23
    }
    v19 -= (^(v19))
    _ = v19
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = float64(-2.25)
    v3 := make(map[int]bool)
    var v4 int32 = 0
    switch 4294967295 {
    case 770:
        v24 := GenS2{F0: 1}
        _ = v24
    case 2147483647:
        switch "a" {
        case "foo bar":
            v25 := GenS2{F0: -355}
            _ = v25
        case "a":
            v3[32768] = false
        case "":
            var v26 uint8 = 0
            _ = v26
        default:
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v2, v4)
            case "a":
                if ((127 > -2147483648) || (!(true && false))) {
                    fmt.Println(v2, v2)
                    var v27 int8 = -1
                    _ = v27
                }
            case "":
                v28 := GenS2{F0: 1}
                _ = v28
            }
        }
    default:
        v4 = v4
    }
    fmt.Println(v4, v4)
    fmt.Println(v4, v4)
    if (42 < 2147483648) {
        for fi29 := 0; fi29 < 10; fi29++ {
            for fi30 := 0; fi30 < 10; fi30++ {
                v31 := []int{}
                _ = fi30
                _ = v31
            }
            var v32 float64 = ((0.1 + ((1.5 + 3.14159) + 3.14159)) / (float64(0.0) * v2))
            _ = fi29
            _ = v32
        }
        v33 := GenS3{F0: 673, F1: 65535}
        v2 = v2
        _ = v33
    } else {
        switch "hello" {
        case "a":
            v4 -= v4
        default:
            var v34 float64 = v2
            _ = v34
        }
    }
    v4 -= (v4 - 1)
    var v35 int = -9223372036854775808
    if false {
        fmt.Println(v35, v2, v35)
        fmt.Println(v4, v4)
    } else {
        v36 := []int{v35, (v35 + (v35 / (((-(v35))) | 1)))}
        var v37 int8 = -128
        var v38 int8 = (v37 * 100)
        _ = v36
        _ = v37
        _ = v38
    }
    var v39 int64 = -463
    v40 := GenS4{F0: 18446744073709551615, F1: (-(v35)), F2: 515}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
}
