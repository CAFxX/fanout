package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 bool }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int64; F1 uint }

func main() {
    v1 := []uint32{4294967295}
    v2 := []string{"hello", ("" + "hello"), "hello", "foo bar"}
    v3 := make(map[int]string)
    var v4 uint16 = 1
    switch "foo bar" {
    case "\n\t":
        if false {
            if (-2147483648 <= 1) {
                fmt.Println(v4)
            }
        }
    case "hello":
        v5 := GenS1{F0: 4294967295, F1: -1, F2: false}
        _ = v5
    case "a":
        switch "a" {
        case "hello":
            v6 := GenS1{F0: 812, F1: 42, F2: false}
            _ = v6
        case "foo bar":
            if ((100 <= 34) && true) {
                for fi7 := 0; fi7 < 5; fi7++ {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        v9 := v3[fi7]
                        _ = fi8
                        _ = v9
                    }
                    _ = fi7
                }
            } else {
                var v10 bool = (!((0.0 * float32(-2.25)) < 1e10))
                _ = v10
            }
        case "\n\t":
            v4 = v4
        default:
            switch "foo bar" {
            case "":
                v4 += v4
            case "\n\t":
                for fi11 := 0; fi11 < 5; fi11++ {
                    v4 = v4
                    _ = fi11
                }
            }
        }
    default:
        var v12 uint8 = 255
        _ = v12
    }
    switch -2147483648 {
    case -2147483648:
        var v13 float64 = (3.14159 + 3.14159)
        _ = v13
    case -658:
        var v14 uint64 = 1
        _ = v14
    default:
        var v15 bool = false
        _ = v15
    }
    switch 52 {
    case -335:
        v4 = (v4 ^ (^(v4)))
    default:
        fmt.Println(v4)
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: 731, F1: -2147483648, F2: true}
        _ = fi16
        _ = v17
    }
    var v18 uint = 42
    var v19 bool = true
    v20 := GenS1{F0: 1, F1: 707, F2: (-2147483648 < 2147483647)}
    var v21 string = ("a" + "")
    v4 = v4
    switch 32767 {
    case 4294967295:
        if false {
            var v22 int8 = 64
            _ = len(v2)
            for fi23 := 0; fi23 < 13; fi23++ {
                _ = len(v1)
                _ = fi23
            }
            _ = v22
        } else {
            for fi24 := 0; fi24 < 11; fi24++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi24
            }
        }
    case -651:
        v21 = "a"
    }
    var v25 int32 = -439
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 3; fi27++ {
            v28 := GenS1{F0: 1, F1: (v25 + (^(v25))), F2: v19}
            if (v19 || (v19 && (9223372036854775807 < -9223372036854775808))) {
                v25 = (v25 & -2147483648)
            } else {
                fmt.Println(v18, v25)
            }
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v21 = "hello"
        if v19 {
            v18 = 42
            v30 := GenS1{F0: 735, F1: 367, F2: (982 >= 42)}
            fmt.Println(v25, v19, fi29)
            _ = v30
        } else {
            v25 = -2147483648
            fmt.Println(v21, v21)
        }
        _ = fi29
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        fmt.Println(v18, v21, v19)
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 uint = 0
            var v34 float32 = 3.14159
            _ = fi32
            _ = v33
            _ = v34
        }
        v35 := GenS1{F0: 65535, F1: v25, F2: (-1 < -9223372036854775808)}
        _ = fi31
        _ = v35
    }
    v36 := GenS2{F0: 9223372036854775807, F1: int32(v18)}
    if (v19 && v19) {
        v37 := GenS3{F0: 42, F1: v18}
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v36
    _ = v4
}
