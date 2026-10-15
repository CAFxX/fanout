package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 int32; F1 uint64; F2 uint64 }

func main() {
    var v1 bool = (0 >= 748)
    var v2 float32 = 0.1
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    var v5 float32 = v2
    v6 := make(map[uint32]bool)
    switch 42 {
    case 128:
        switch -2147483648 {
        case 0:
            if ((32767 == 127) || (2 >= 2)) {
                v7 := GenS1{F0: 1000, F1: 4294967295, F2: 1}
                for fi8 := 0; fi8 < 10; fi8++ {
                    fmt.Println(v5, v1)
                    _ = fi8
                }
                fmt.Println(v5)
                _ = v7
            }
        case 389:
            var v9 int64 = 9223372036854775807
            _ = v9
        default:
            fmt.Println(v2, v5)
        }
    case 127:
        fmt.Println(v5)
    case -2147483648:
        v1 = (170 >= 248)
    }
    v5 = ((((1.5 + float32(((-2.25 - (3.14159 - -2.25)) - 0.1))) / v2) / 1.0) - v5)
    v1 = true
    v10 := []int{9223372036854775807, -1, 32767, 127}
    v11 := GenS2{F0: 128, F1: 1, F2: "hello"}
    v5 = v2
    v5 = v5
    for fi12 := 0; fi12 < 7; fi12++ {
        v1 = (0 < -1)
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 uint8 = 255
            _ = fi13
            _ = v14
        }
        if len(v10) > 0 {
            v10[0] = (-(fi12))
        }
        _ = fi12
    }
    switch "\n\t" {
    case "a":
        v5 = float32((-2.25 - (1e10 / 1.0)))
    case "":
        fmt.Println(v5)
    default:
        if (4294967295 >= 42) {
            switch 32768 {
            case 42:
                var v15 int8 = -1
                _ = v15
            default:
                v2 = float32(((0.0 + 0.1) / (float64(((1.5 * ((0.1 / (-2.25 * v5)) - (v2 - v2))) * v5)) / 1.5)))
            }
        }
    }
    if v1 {
        var v16 string = "\n\t"
        _ = v16
    } else {
        v17 := GenS2{F0: 128, F1: 42, F2: ("foo bar" + "foo bar")}
        fmt.Println(v5, v2, v1)
        v18 := GenS1{F0: -1, F1: 65535, F2: 1}
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        switch "hello" {
        case "hello":
            var v20 uint32 = 152
            _ = v20
        case "\n\t":
            for fi21 := 0; fi21 < 19; fi21++ {
                if len(v10) > 0 {
                    v10[0] = 9223372036854775807
                }
                v22 := GenS3{F0: 42, F1: 42, F2: 42}
                _ = fi21
                _ = v22
            }
        default:
            if (true && false) {
                _ = len(v10)
            } else {
                for fi23 := 0; fi23 < 18; fi23++ {
                    if v1 {
                        fmt.Println(fi23, v5, fi23)
                    }
                    _ = fi23
                }
                v1 = (4294967295 < 2)
            }
        }
        _ = fi19
    }
    switch 774 {
    case 2147483647:
        fmt.Println(v5)
    }
    fmt.Println(v5, v2)
    fmt.Println(v2, v2, v2)
    if (18446744073709551615 >= 598) {
        v2 = (3.14159 + v2)
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS1{F0: 1000, F1: 18446744073709551615, F2: 0}
            v26 := GenS3{F0: 0, F1: 2, F2: 1}
            _ = fi24
            _ = v25
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 11; fi27++ {
            if len(v10) > 0 {
                v10[0] = fi27
            }
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 uint = 830
                if v1 {
                    if (9223372036854775807 == 32768) {
                        v30 := GenS1{F0: -741, F1: v29, F2: 255}
                        _ = v30
                    }
                }
                _ = fi28
                _ = v29
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi27
        }
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
