package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int32 }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 uint32 }
func genFn0() uint {
    var v7 uint32 = 549
    v8 := []int{-2147483648, -2147483648, -1, 128}
    var v9 float32 = 0.1
    v10 := GenS1{F0: 190, F1: -1, F2: 1}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 string, p1 int, p2 int16) bool {
    if (!true) {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 bool = (float32(1e10) > float32(0.0))
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            if (-1 < -128) {
                var v14 uint64 = 4294967295
                _ = v14
            }
            var v15 int32 = -2
            _ = fi13
            _ = v15
        }
        p2 -= 1
    }
    p0 = ((("" + p0) + "\n\t") + "foo bar")
    p2 -= (^(p2))
    p0 = p0
    if ((145 >= 18446744073709551615) && true) {
        for fi16 := 0; fi16 < 7; fi16++ {
            if true {
                switch (^(p1)) {
                case 32767:
                    var v17 uint = genFn0()
                    _ = v17
                case -108:
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 uint8 = 2
                        _ = fi18
                        _ = v19
                    }
                case 32768:
                    v20 := GenS1{F0: 42, F1: 9223372036854775807, F2: 2147483647}
                    _ = v20
                default:
                    v21 := GenS1{F0: 664, F1: -1, F2: 2147483647}
                    _ = v21
                }
            }
            v22 := []int{}
            p2 = p2
            _ = fi16
            _ = v22
        }
        p0 = "hello"
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -589
    var v3 int32 = v2
    var v4 uint32 = 761
    v5 := []int{-706}
    v6 := []string{}
    switch 750 {
    case -207:
        if (32767 != 961) {
            switch ("foo bar" + ("" + "\n\t")) {
            case "":
                if len(v5) > 0 {
                    v5[0] = -963
                }
            case "\n\t":
                var v23 int8 = 0
                _ = v23
            }
            var v24 uint64 = 594
            switch 1 {
            case 127:
                if false {
                    switch 1 {
                    case 42:
                        var v25 uint32 = (v4 >> ((genFn0()) & 63))
                        _ = v25
                    }
                }
            }
            _ = v24
        }
    case 128:
        switch "\n\t" {
        case "hello":
            fmt.Println(v2, v3)
        default:
            _ = len(v6)
        }
    default:
        if false {
            v2 -= 1
            var v26 int64 = -1
            _ = v26
        }
    }
    fmt.Println(v3)
    fmt.Println(v3, v2)
    v2 -= (v3 & -939)
    v5 = append(v5, -2147483648)
    var v27 int16 = -32768
    var v28 uint = 681
    for fi29 := 0; fi29 < 9; fi29++ {
        v27 = v27
        v30 := GenS2{F0: 255, F1: 65535, F2: -128}
        _ = fi29
        _ = v30
    }
    v31 := GenS3{F0: (!true), F1: (v28 + v28), F2: (v4 & 426)}
    if (0.1 < 0.1) {
        v4 = v4
        switch "" {
        case "\n\t":
            _ = len(v6)
        case "":
            fmt.Println(v3, v27, v4)
        case "foo bar":
            fmt.Println(v2, v28)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
