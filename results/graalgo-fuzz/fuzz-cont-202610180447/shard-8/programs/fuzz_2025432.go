package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int64 }
func genFn0(p0 uint) uint16 {
    p0 = p0
    p0 = 65535
    v5 := []int{42, -2147483648, 9223372036854775807}
    _ = p0
    _ = v5
    return 32767
}
func genFn1() string {
    var v6 int64 = 42
    v6 = (^(v6))
    switch 4294967295 {
    case 128:
        if ((151 < 255) || ((true || true) || (v6 <= (v6 * 1)))) {
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 uint = 1
                fmt.Println(v8)
                _ = fi7
                _ = v8
            }
            switch "foo bar" {
            case "hello":
                v6 -= 161
            case "a":
                if false {
                    var v9 string = ""
                    _ = v9
                } else {
                    var v10 string = "a"
                    v11 := GenS1{F0: -91, F1: 42, F2: v6}
                    v6 += v6
                    _ = v10
                    _ = v11
                }
            case "\n\t":
                fmt.Println(v6)
            }
            v12 := []int{-290, -9223372036854775808}
            _ = v12
        } else {
            switch "" {
            case "foo bar":
                switch "hello" {
                case "a":
                    for fi13 := 0; fi13 < 18; fi13++ {
                        v14 := []int{(fi13 << ((270) & 63)), 0}
                        fmt.Println(fi13, fi13)
                        v6 += (-(v6))
                        _ = fi13
                        _ = v14
                    }
                }
            case "a":
                var v15 uint32 = 71
                _ = v15
            default:
                for fi16 := 0; fi16 < 2; fi16++ {
                    v6 -= (^(v6))
                    if (!false) {
                        v17 := []int{}
                        _ = v17
                    }
                    _ = fi16
                }
            }
            if false {
                v6 = 9223372036854775807
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := GenS1{F0: -1, F1: 42, F2: (v6 + v6)}
                    _ = fi18
                    _ = v19
                }
            }
        }
    case -9223372036854775808:
        fmt.Println(v6)
    case 32767:
        v6 += 0
    default:
        fmt.Println(v6)
    }
    _ = v6
    return "hello"
}

func main() {
    var v1 int64 = 42
    var v2 int8 = 42
    var v3 int64 = (v1 | 1)
    v4 := make(map[int]string)
    fmt.Println(v2)
    v20 := v4[2147483647]
    v2 -= (v2 << ((255) & 63))
    v21 := []int{127, -562}
    var v22 float64 = (3.14159 - (3.14159 * -2.25))
    var v23 uint32 = 1
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    var v24 uint16 = genFn0(0)
    v25 := GenS1{F0: v2, F1: 18446744073709551615, F2: 1}
    var v26 uint32 = 0
    fmt.Println(v24, v1, v20)
    var v27 uint32 = (v26 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
