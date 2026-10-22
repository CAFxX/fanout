package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 int32; F1 int32; F2 int }

func main() {
    v1 := []int{}
    var v2 uint16 = 0
    v3 := []int{2147483647, 1, -2147483648}
    var v4 uint16 = (v2 ^ (v2 % ((v2) | 1)))
    if (42 <= 2147483648) {
        v4 = (v2 % (((v4 >> ((0) & 63))) | 1))
        if ((!false) && ((!(true && true)) || true)) {
            fmt.Println(v4, v4)
        }
        switch 128 {
        case -1:
            var v5 uint64 = 0
            _ = v5
        default:
            v4 -= v2
        }
    }
    v2 -= v4
    fmt.Println(v2, v2)
    for fi6 := 0; fi6 < 5; fi6++ {
        if (!true) {
            v4 = v4
        } else {
            fmt.Println(v4, fi6, v4)
            var v7 int8 = -72
            _ = v7
        }
        v8 := GenS1{F0: fi6, F1: 2, F2: 0}
        _ = fi6
        _ = v8
    }
    if (2 != 124) {
        for fi9 := 0; fi9 < 15; fi9++ {
            if true {
                var v10 string = "hello"
                var v11 int16 = -283
                _ = v10
                _ = v11
            }
            _ = fi9
        }
        v2 += 32767
    } else {
        if (((0 <= 42) && false) || (9223372036854775807 != 9223372036854775807)) {
            v12 := GenS1{F0: 32768, F1: 42, F2: 666}
            _ = v12
        } else {
            fmt.Println(v2)
        }
    }
    v4 -= v2
    var v13 uint8 = 1
    switch 1 {
    case 0:
        var v14 int64 = -889
        _ = v14
    case 2147483647:
        v13 = v13
    case 9223372036854775807:
        switch ("a" + "foo bar") {
        case "":
            v4 = 677
        default:
            var v15 uint32 = 2147483648
            _ = v15
        }
    }
    switch ("foo bar" + "hello") {
    case "":
        v4 = (v4 | (v4 >> ((18446744073709551615) & 63)))
    case "\n\t":
        v4 -= 32767
    default:
        if true {
            if (609 <= 32767) {
                v13 -= v13
                if false {
                    var v16 float32 = (-2.25 * 0.0)
                    _ = v16
                } else {
                    var v17 float64 = -2.25
                    v18 := GenS2{F0: -1, F1: 329, F2: -9223372036854775808}
                    _ = v17
                    _ = v18
                }
                v2 -= v4
            }
            v4 = 722
            if (false || (false && false)) {
                switch 916 {
                case 0:
                    var v19 int = 0
                    _ = v19
                case 1:
                    if false {
                        _ = len(v3)
                        fmt.Println(v2, v2, v13)
                    }
                }
                v13 -= uint8(v2)
            }
        }
    }
    switch 9223372036854775807 {
    case -1:
        fmt.Println(v4, v13, v2)
    case 127:
        if false {
            if true {
                var v20 int8 = 1
                v4 = (v4 & (v2 & v4))
                _ = v20
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                for fi22 := 0; fi22 < 10; fi22++ {
                    fmt.Println(fi22, v4, v4)
                    _ = fi22
                }
                var v23 bool = (1 < 1)
                _ = fi21
                _ = v23
            }
        } else {
            _ = len(v3)
        }
    default:
        fmt.Println(v2, v2, v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
