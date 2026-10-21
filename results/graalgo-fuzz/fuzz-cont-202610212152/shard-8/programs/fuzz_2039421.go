package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }
func genFn0() int64 {
    if ((!false) || false) {
        var v6 uint64 = 18446744073709551615
        _ = v6
    } else {
        for fi7 := 0; fi7 < 5; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{(fi8 * fi8), 42}
                v10 := GenS1{F0: -9223372036854775808, F1: 1}
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 int = fi7
            _ = fi7
            _ = v11
        }
    }
    var v12 float32 = 0.1
    _ = v12
    return -9223372036854775808
}
func genFn1() uint64 {
    var v13 uint64 = 554
    if (1 == -1) {
        fmt.Println(v13)
    } else {
        v13 = 632
        var v14 float32 = (0.1 * ((1e10 + 3.14159) + 0.1))
        _ = v14
    }
    var v15 bool = false
    var v16 int = 9223372036854775807
    _ = v13
    _ = v15
    _ = v16
    return 42
}

func main() {
    var v1 int16 = -32768
    var v2 bool = false
    v3 := make(map[uint32]bool)
    var v4 float64 = (((3.14159 / float64(1e10)) + 1e10) * -2.25)
    var v5 string = ""
    v17 := GenS1{F0: int64(v1), F1: 42}
    if v2 {
        switch "" {
        case "hello":
            v1 = v1
        case "a":
            var v18 int32 = -787
            _ = v18
        case "":
            var v19 int = 128
            _ = v19
        default:
            v2 = false
        }
    }
    if false {
        var v20 bool = v2
        v4 -= v4
        _ = v20
    }
    if (0.1 != 0.1) {
        v21 := GenS2{F0: -9223372036854775808, F1: 902, F2: 18446744073709551615}
        var v22 uint = 464
        for fi23 := 0; fi23 < 18; fi23++ {
            fmt.Println(v2, v2)
            _ = fi23
        }
        _ = v21
        _ = v22
    } else {
        v24 := v3[4294967295]
        v25 := GenS3{F0: 511, F1: 32767, F2: 0}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 uint8 = 0
        for fi28 := 0; fi28 < 12; fi28++ {
            v5 = ((("a" + (v5 + "a")) + "hello") + v5)
            fmt.Println(v27)
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    switch (v5 + v5) {
    case "\n\t":
        switch 32767 {
        case -579:
            v4 = v4
        default:
            switch 4294967295 {
            case 128:
                var v29 uint64 = 1
                _ = v29
            case 1:
                var v30 int8 = 42
                _ = v30
            case 42:
                v5 = (("a" + ("hello" + "\n\t")) + v5)
            }
        }
    }
    fmt.Println(v2)
    v31 := GenS3{F0: 32767, F1: -2147483648, F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
