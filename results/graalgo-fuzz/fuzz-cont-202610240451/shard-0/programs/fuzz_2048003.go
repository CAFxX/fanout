package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int32, p2 uint32) string {
    if (32768 <= 128) {
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{(fi7 & fi7), (fi7 % (((fi7 - (^(fi7)))) | 1)), (fi7 ^ fi7), -1}
            v9 := GenS1{F0: 0, F1: 1}
            _ = fi7
            _ = v8
            _ = v9
        }
        p2 = 2147483648
    } else {
        v10 := GenS2{F0: (((42 >= 1) || false) || (183 != 1))}
        for fi11 := 0; fi11 < 7; fi11++ {
            p2 = 2147483647
            switch int(p2) {
            case -1:
                switch "\n\t" {
                case "":
                    p1 = -2147483648
                }
            default:
                if false {
                    fmt.Println(p0)
                }
            }
            _ = fi11
        }
        switch "hello" {
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                fmt.Println(p0, p1)
            default:
                var v12 int32 = (-(p1))
                _ = v12
            }
        }
        _ = v10
    }
    v13 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return ("\n\t" + (("" + "hello") + "a"))
}

func main() {
    var v1 float64 = 0.0
    var v2 uint8 = 2
    v3 := []int{-774, 1, 2147483647}
    v4 := []int{7, 2147483647, 9223372036854775807}
    v5 := make(map[int]string)
    v6 := []string{}
    v1 = v1
    delete(v5, 1)
    if (true || (194 == 28)) {
        switch ("hello" + ((("foo bar" + "\n\t") + "foo bar") + "hello")) {
        case "a":
            var v14 float32 = -2.25
            _ = v14
        case "foo bar":
            var v15 uint64 = 791
            _ = v15
        case "\n\t":
            v16 := GenS3{F0: 2}
            _ = v16
        }
    } else {
        v17 := GenS2{F0: (1 > 4294967295)}
        v1 -= 0.0
        _ = v17
    }
    fmt.Println(v2)
    _ = len(v3)
    v2 += (v2 + v2)
    fmt.Println(v2, v2)
    switch -497 {
    case 67:
        fmt.Println(v1)
    default:
        v18 := GenS3{F0: 32767}
        _ = v18
    }
    var v19 uint64 = 4294967295
    if false {
        var v20 int16 = -1
        _ = v20
    }
    v1 += v1
    if false {
        switch "hello" {
        case "a":
            var v21 float64 = 1e10
            _ = v21
        case "hello":
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 int8 = 110
                fmt.Println(v1)
                _ = fi22
                _ = v23
            }
        case "":
            v1 += (1e10 / (float64(1e10) + 1.5))
        }
        if false {
            v1 -= (((v1 * (v1 + (0.1 / (0.0 - 1e10)))) / (0.1 - (v1 * v1))) - v1)
        }
        if len(v4) > 0 {
            v4[0] = 323
        }
    }
    v24 := GenS1{F0: 9223372036854775807, F1: 0}
    v25 := GenS2{F0: (false && false)}
    fmt.Println(v19)
    var v26 int = 677
    fmt.Println(v19, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
