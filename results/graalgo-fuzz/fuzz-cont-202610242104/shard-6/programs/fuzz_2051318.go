package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 uint64 }
func genFn0(p0 int16, p1 int16) uint16 {
    var v7 string = ""
    v8 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 65535
}
func genFn1() string {
    if false {
        var v9 string = "foo bar"
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS2{F0: fi10}
            switch -9223372036854775808 {
            case -532:
                v12 := []int{9223372036854775807}
                _ = v12
            default:
                v9 = v9
            }
            v9 = v9
            _ = fi10
            _ = v11
        }
        var v13 bool = true
        _ = v13
        _ = v9
    } else {
        var v14 int64 = 975
        fmt.Println(v14)
        if false {
            v14 -= v14
            v14 -= v14
        }
        _ = v14
    }
    var v15 int8 = -1
    v16 := GenS1{F0: 0}
    v17 := []int{181}
    v15 = (v15 / ((v15) | 1))
    _ = v15
    _ = v16
    _ = v17
    return "a"
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int16 = 1
    var v4 uint64 = 0
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 255
    if true {
        if (false && true) {
            for fi18 := 0; fi18 < 1; fi18++ {
                switch "hello" {
                case "\n\t":
                    if false {
                        v19 := GenS1{F0: (v4 - (v4 % ((v4) | 1)))}
                        _ = v19
                    } else {
                        v20 := GenS3{F0: 2, F1: (v4 | v4), F2: (v6 * 823)}
                        var v21 int8 = 42
                        v22 := GenS1{F0: (v4 / ((1) | 1))}
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                default:
                    v4 += (^(v4))
                }
                var v23 uint64 = (v4 | 552)
                var v24 int32 = 0
                _ = fi18
                _ = v23
                _ = v24
            }
            v25 := GenS3{F0: 4294967295, F1: 1, F2: 42}
            fmt.Println(v3, v3, v6)
            _ = v25
        }
    }
    switch "\n\t" {
    case "foo bar":
        var v26 int16 = 42
        _ = v26
    case "a":
        v6 += genFn0(-412, v3)
    case "":
        var v27 int16 = -1
        _ = v27
    default:
        v28 := GenS4{F0: 0}
        _ = v28
    }
    v29 := GenS4{F0: -2147483648}
    v6 = (v6 & v6)
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS3{F0: 18446744073709551615, F1: v4, F2: (v6 | genFn0(v3, (-(v3))))}
        _ = fi30
        _ = v31
    }
    _ = len(v1)
    if true {
        switch "foo bar" {
        case "hello":
            v3 = v3
        case "":
            for fi32 := 0; fi32 < 4; fi32++ {
                switch "\n\t" {
                case "foo bar":
                    var v33 int16 = (v3 - (v3 & v3))
                    _ = v33
                }
                var v34 float32 = (0.0 + 3.14159)
                _ = fi32
                _ = v34
            }
        case "foo bar":
            for fi35 := 0; fi35 < 8; fi35++ {
                var v36 int = 9223372036854775807
                fmt.Println(v3, v3, fi35)
                var v37 int16 = 1000
                _ = fi35
                _ = v36
                _ = v37
            }
        default:
            fmt.Println(v4, v4, v4)
        }
    }
    delete(v1, 512)
    switch ("" + "foo bar") {
    case "hello":
        fmt.Println(v3)
    case "\n\t":
        v38 := GenS5{F0: 1, F1: (^(v4))}
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
