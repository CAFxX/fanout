package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }
type GenS2 struct { F0 string; F1 bool; F2 int16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 uint) int32 {
    var v5 int64 = -9223372036854775808
    if (740 >= 652) {
        fmt.Println(p0, p1)
        p1 = (p1 << ((p1) & 63))
    } else {
        switch "\n\t" {
        case "a":
            var v6 bool = (1 != 128)
            _ = v6
        case "foo bar":
            v7 := []int{(p0 + (p0 % ((-514) | 1))), (p0 ^ p0)}
            _ = v7
        case "\n\t":
            fmt.Println(p0, p1)
        default:
            v8 := []int{-9223372036854775808, p0, 127, (p0 * p0)}
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 uint8, p1 uint8) string {
    if (18446744073709551615 != 1) {
        switch ((("hello" + "foo bar") + "foo bar") + ("hello" + "hello")) {
        case "\n\t":
            var v9 int = 32768
            _ = v9
        default:
            v10 := []int{}
            _ = v10
        }
        var v11 uint64 = 0
        _ = v11
    }
    if (!(18446744073709551615 < 1)) {
        v12 := GenS1{F0: ("a" + ""), F1: -1, F2: 616}
        fmt.Println(p1)
        if (!(!false)) {
            v13 := []int{42, -2147483648, 127, 0}
            var v14 uint64 = 620
            _ = v13
            _ = v14
        }
        _ = v12
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            p0 = 1
            _ = fi15
        }
        v16 := GenS2{F0: "foo bar", F1: ((2 != 42) && false), F2: -1}
        _ = v16
    }
    p0 = 121
    _ = p0
    _ = p1
    return "\n\t"
}

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 4
    var v3 string = "a"
    var v4 uint16 = (v1 * (^(v1)))
    v17 := []int{127, 264}
    var v18 uint16 = (^(v1))
    v2 += v2
    v17 = append(v17, 900)
    v4 += (^(v18))
    v19 := GenS1{F0: "a", F1: genFn0(-9223372036854775808, 877), F2: 1}
    fmt.Println(v2, v4, v3)
    switch 303 {
    case 128:
        var v20 int = -2147483648
        _ = v20
    case 127:
        v4 = uint16(v2)
    default:
        v17 = append(v17, 0)
    }
    v21 := GenS3{F0: -1}
    var v22 string = ""
    if (((2147483648 <= 2147483648) && (true && false)) && false) {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        _ = len(v17)
        switch ((v22 + "\n\t") + "a") {
        case "\n\t":
            var v24 string = "a"
            _ = v24
        case "a":
            switch "" {
            case "\n\t":
                var v25 string = v3
                _ = v25
            case "a":
                var v26 int8 = 79
                _ = v26
            }
        }
        v3 = v3
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
