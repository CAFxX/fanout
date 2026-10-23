package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int32; F2 int16 }
type GenS5 struct { F0 int16; F1 uint; F2 string }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    var v4 int64 = -215
    v5 := GenS1{F0: "\n\t", F1: (v4 ^ v4), F2: 707}
    if (!(true || true)) {
        fmt.Println(v1, v4)
        var v6 float32 = 3.14159
        _ = v6
    }
    var v7 string = ("hello" + (("a" + ("foo bar" + "")) + "\n\t"))
    var v8 string = "hello"
    v9 := v3[(v8 + "a")]
    var v10 int16 = -709
    v9 = true
    for fi11 := 0; fi11 < 12; fi11++ {
        switch v8 {
        case "\n\t":
            v12 := GenS1{F0: ("hello" + v8), F1: v4, F2: 677}
            _ = v12
        case "":
            _ = len(v3)
        }
        if ((1 <= 1) && false) {
            var v13 uint64 = 42
            if (((((v1 - v1) - v1) * v1) + v1) <= v1) {
                var v14 string = v7
                _ = len(v3)
                v4 = v4
                _ = v14
            }
            _ = v13
        } else {
            if false {
                fmt.Println(v10)
                v15 := []int{(-(fi11)), 2147483647, (fi11 | (fi11 - fi11))}
                if (v9 && (65535 < 53)) {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 int32 = 0
                        _ = fi16
                        _ = v17
                    }
                    v7 = ("hello" + (v8 + ((v7 + v8) + "a")))
                    for fi18 := 0; fi18 < 10; fi18++ {
                        var v19 bool = false
                        v8 = v8
                        _ = fi18
                        _ = v19
                    }
                } else {
                    v20 := GenS1{F0: v8, F1: 1, F2: 2}
                    v21 := GenS2{F0: (v4 * 206), F1: (0 != 195), F2: 386}
                    _ = v20
                    _ = v21
                }
                _ = v15
            }
        }
        v22 := GenS1{F0: "foo bar", F1: -460, F2: 1}
        _ = fi11
        _ = v22
    }
    v7 = "hello"
    v23 := []int{-1, 2147483647, -9223372036854775808}
    if (v4 >= v4) {
        for fi24 := 0; fi24 < 2; fi24++ {
            switch "hello" {
            case "\n\t":
                var v25 int64 = v4
                _ = v25
            case "hello":
                v26 := GenS3{F0: 0}
                _ = v26
            default:
                v8 = v7
            }
            var v27 int32 = 182
            v28 := GenS4{F0: 4294967295, F1: 2147483647, F2: (v10 - (v10 >> ((18446744073709551615) & 63)))}
            _ = fi24
            _ = v27
            _ = v28
        }
    }
    var v29 string = v8
    _ = len(v3)
    switch -313 {
    case 2147483647:
        if true {
            if v9 {
                for fi30 := 0; fi30 < 2; fi30++ {
                    fmt.Println(v8)
                    v8 = (("foo bar" + "") + v29)
                    v8 = ("" + "a")
                    _ = fi30
                }
            }
            if v9 {
                v31 := GenS1{F0: v29, F1: v4, F2: 135}
                v9 = v9
                _ = v31
            }
            if (!(!(0 >= 32767))) {
                if v9 {
                    var v32 bool = true
                    _ = v32
                }
                if (v9 || (v9 && v9)) {
                    v33 := GenS5{F0: v10, F1: 42, F2: v29}
                    if v9 {
                        v34 := v2[127]
                        var v35 uint32 = 65535
                        _ = v34
                        _ = v35
                    }
                    _ = v33
                } else {
                    v9 = (0 > 284)
                }
                fmt.Println(v9, v7)
            }
        }
    case -826:
        v8 = "foo bar"
    case 0:
        switch 0 {
        case 9223372036854775807:
            _ = len(v3)
        case 612:
            fmt.Println(v9, v9, v10)
        case -9223372036854775808:
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        default:
            v10 -= v10
        }
    }
    fmt.Println(v10, v29)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
