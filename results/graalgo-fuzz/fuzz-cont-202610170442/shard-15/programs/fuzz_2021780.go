package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint; F2 string }

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 18446744073709551615
    var v3 bool = (!(false && true))
    var v4 uint = 179
    var v5 int16 = -218
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int16 = v5
        _ = fi6
        _ = v7
    }
    fmt.Println(v4)
    var v8 uint32 = 65535
    fmt.Println(v8, v4, v8)
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: -9223372036854775808, F1: 65535, F2: "foo bar"}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: -128}
    if true {
        v2 += v2
        var v12 bool = (2147483647 == 1)
        _ = v12
    }
    fmt.Println(v8, v2)
    var v13 int = 32768
    v14 := GenS2{F0: int8(v2)}
    fmt.Println(v13, v13)
    fmt.Println(v13, v5)
    if (v3 || (uint8(v4) != 2)) {
        var v15 bool = (v2 > uint64(v4))
        _ = v15
    } else {
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(v2, v3)
            for fi17 := 0; fi17 < 9; fi17++ {
                switch v13 {
                case -474:
                    v18 := GenS2{F0: -1}
                    _ = v18
                case -9223372036854775808:
                    _ = len(v1)
                case -2147483648:
                    var v19 int64 = 1
                    _ = v19
                default:
                    v3 = ((v8 | 929) < v8)
                }
                v20 := []int{(fi16 - (-(fi17))), fi17, -828}
                fmt.Println(v2, fi16, v13)
                _ = fi17
                _ = v20
            }
            _ = fi16
        }
    }
    fmt.Println(v4, v5, v4)
    if false {
        if (-9223372036854775808 > v13) {
            switch "" {
            case "foo bar":
                var v21 uint8 = 138
                _ = v21
            case "\n\t":
                delete(v1, v13)
            default:
                var v22 int64 = 42
                _ = v22
            }
            v23 := v1[127]
            if ((397 < 682) && (float64(1e10) <= -2.25)) {
                var v24 int16 = v5
                v25 := GenS2{F0: 127}
                _ = v24
                _ = v25
            } else {
                v2 -= (v2 & 398)
            }
            _ = v23
        }
    }
    switch "foo bar" {
    case "":
        v4 += v4
    case "a":
        var v26 uint64 = v2
        _ = v26
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            var v27 uint = (^(v4))
            _ = v27
        case "a":
            fmt.Println(v13, v3, v8)
        case "foo bar":
            for fi28 := 0; fi28 < 5; fi28++ {
                _ = len(v1)
                if v3 {
                    for fi29 := 0; fi29 < 20; fi29++ {
                        v2 = 186
                        _ = fi29
                    }
                } else {
                    v3 = ("hello" != (("foo bar" + "") + "hello"))
                }
                v2 -= 846
                _ = fi28
            }
        }
    default:
        v30 := GenS3{F0: 117, F1: (v4 & uint(v2)), F2: (("hello" + "foo bar") + "\n\t")}
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
