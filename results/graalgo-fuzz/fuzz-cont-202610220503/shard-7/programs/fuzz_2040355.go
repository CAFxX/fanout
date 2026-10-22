package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 string; F1 int8; F2 int64 }

func main() {
    var v1 int64 = 42
    var v2 uint = 255
    var v3 int64 = 244
    var v4 string = "\n\t"
    var v5 uint8 = 204
    for fi6 := 0; fi6 < 6; fi6++ {
        v1 -= v3
        fmt.Println(v2, v2, v5)
        _ = fi6
    }
    v7 := []int{2147483647, 128}
    if (v4 > (("foo bar" + v4) + v4)) {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: (v2 / (((v2 ^ (v2 - v2))) | 1)), F1: v1}
            if len(v7) > 0 {
                v7[0] = int(v3)
            }
            _ = fi8
            _ = v9
        }
        v2 += 121
    }
    if true {
        v4 = (v4 + v4)
    } else {
        var v10 uint16 = 42
        v11 := GenS2{F0: true}
        _ = v10
        _ = v11
    }
    if false {
        fmt.Println(v4)
        fmt.Println(v2)
    } else {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 int8 = -128
            _ = len(v7)
            _ = fi12
            _ = v13
        }
        var v14 int8 = -36
        _ = v14
    }
    if false {
        if (false || ((!true) && false)) {
            var v15 float64 = -2.25
            _ = v15
        } else {
            if ((-2.25 - 0.0) < 3.14159) {
                v7 = append(v7, 32767)
                if (true && true) {
                    v7 = append(v7, -9223372036854775808)
                    var v16 uint16 = 342
                    _ = v16
                }
            } else {
                v17 := GenS3{F0: v3, F1: 127}
                _ = v17
            }
            var v18 float32 = 3.14159
            _ = v18
        }
    } else {
        if (!((2 < 0) || false)) {
            for fi19 := 0; fi19 < 6; fi19++ {
                v5 = 2
                switch fi19 {
                case 869:
                    v1 += v3
                default:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
                _ = fi19
            }
        }
        fmt.Println(v2, v1, v4)
        switch -182 {
        case 0:
            var v20 int8 = 0
            _ = v20
        case 9223372036854775807:
            _ = len(v7)
        case -1:
            fmt.Println(v1)
        }
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        fmt.Println(v5)
        v1 = 9223372036854775807
        _ = fi21
    }
    switch -2147483648 {
    case 365:
        v22 := GenS3{F0: v3, F1: -22}
        _ = v22
    case 128:
        if false {
            v23 := GenS2{F0: (-127 > 1)}
            if (v5 <= v5) {
                for fi24 := 0; fi24 < 4; fi24++ {
                    v25 := GenS4{F0: true, F1: (true || false)}
                    _ = fi24
                    _ = v25
                }
                _ = len(v7)
                if len(v7) > 0 {
                    v7[0] = -812
                }
            } else {
                _ = len(v7)
            }
            switch (("hello" + v4) + v4) {
            case "\n\t":
                switch 4294967295 {
                case -9223372036854775808:
                    var v26 float64 = 1e10
                    _ = v26
                case 32767:
                    if true {
                        v27 := GenS3{F0: -250, F1: 1}
                        fmt.Println(v3, v2)
                        fmt.Println(v1, v1)
                        _ = v27
                    }
                case 32768:
                    var v28 uint16 = 2
                    _ = v28
                default:
                    if (42 <= 71) {
                        v2 -= uint(v5)
                        var v29 int32 = 441
                        _ = v29
                    }
                }
            case "foo bar":
                v30 := GenS5{F0: false}
                _ = v30
            case "a":
                if (-1 == -32768) {
                    if len(v7) > 0 {
                        v7[0] = -921
                    }
                    v31 := GenS6{F0: 18446744073709551615, F1: v3}
                    v32 := GenS7{F0: "a", F1: 1, F2: -1}
                    _ = v31
                    _ = v32
                }
            }
            _ = v23
        } else {
            var v33 string = v4
            if true {
                var v34 float64 = 0.0
                _ = v34
            }
            for fi35 := 0; fi35 < 19; fi35++ {
                fmt.Println(v1, v4, v4)
                v5 = (v5 + v5)
                v1 = v3
                _ = fi35
            }
            _ = v33
        }
    case 9223372036854775807:
        var v36 int64 = (v1 % ((v3) | 1))
        _ = v36
    default:
        v7 = append(v7, -488)
    }
    v37 := GenS2{F0: true}
    for fi38 := 0; fi38 < 1; fi38++ {
        v4 = (((v4 + v4) + "\n\t") + v4)
        switch "a" {
        case "foo bar":
            for fi39 := 0; fi39 < 14; fi39++ {
                v40 := GenS5{F0: false}
                v1 = v3
                _ = fi39
                _ = v40
            }
        case "a":
            if (true && (false && true)) {
                v7 = append(v7, -1)
                if len(v7) > 0 {
                    v7[0] = 32768
                }
            }
        case "hello":
            fmt.Println(fi38, v3, v3)
        }
        _ = fi38
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        v5 = v5
        var v42 uint64 = 42
        _ = fi41
        _ = v42
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v7 = append(v7, -702)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v7
}
