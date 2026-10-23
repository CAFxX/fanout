package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0() int8 {
    var v6 int64 = 9223372036854775807
    switch "foo bar" {
    case "foo bar":
        switch 42 {
        case -2147483648:
            var v7 uint16 = 143
            _ = v7
        case 127:
            v6 = v6
        default:
            v6 += v6
        }
    case "a":
        v6 = (v6 / ((v6) | 1))
    case "":
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS1{F0: 42, F1: 0}
                switch 4294967295 {
                case 32767:
                    v6 = (v6 % (((v6 / ((9223372036854775807) | 1))) | 1))
                }
                fmt.Println(fi8, v6, v6)
                _ = fi9
                _ = v10
            }
            switch "" {
            case "\n\t":
                var v11 uint8 = 2
                _ = v11
            default:
                switch (fi8 ^ int(v6)) {
                case -1:
                    v6 -= (v6 - 42)
                case 358:
                    v6 -= -1
                }
            }
            _ = fi8
        }
    default:
        v12 := []int{127, 42, 127}
        _ = v12
    }
    _ = v6
    return -55
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    v5 := []int{552, 619}
    if (v2 < v2) {
        v2 = (v2 % ((v2) | 1))
        switch "" {
        case "hello":
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: 42, F1: 127}
                var v15 int16 = 1000
                _ = fi13
                _ = v14
                _ = v15
            }
        case "foo bar":
            v2 += (^(v2))
        }
        switch -2147483648 {
        case -9223372036854775808:
            var v16 string = "hello"
            _ = v16
        case 4294967295:
            if false {
                v17 := GenS2{F0: 0, F1: 18446744073709551615}
                fmt.Println(v2)
                _ = v17
            } else {
                var v18 int16 = 42
                _ = v18
            }
        }
    }
    v2 -= -1
    var v19 int64 = (^(v2))
    if (0 > 1) {
        var v20 uint16 = 0
        v19 = (v19 << ((18446744073709551615) & 63))
        if false {
            v21 := GenS2{F0: 4294967295, F1: 65535}
            v2 += -1
            var v22 uint = 4294967295
            _ = v21
            _ = v22
        }
        _ = v20
    }
    _ = len(v5)
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int = (^(fi23))
        var v25 int16 = 1
        _ = fi23
        _ = v24
        _ = v25
    }
    v19 -= v2
    v26 := GenS1{F0: -32768, F1: 255}
    if true {
        var v27 float32 = -2.25
        _ = v27
    } else {
        switch "" {
        case "a":
            var v28 float64 = 3.14159
            _ = v28
        case "":
            var v29 string = ("hello" + "a")
            _ = v29
        case "foo bar":
            var v30 int8 = genFn0()
            _ = v30
        default:
            var v31 int64 = v2
            _ = v31
        }
        var v32 string = "hello"
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
