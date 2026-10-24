package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int }
type GenS4 struct { F0 uint8; F1 uint }
func genFn0(p0 int8, p1 string, p2 string) uint64 {
    var v4 int8 = p0
    v5 := GenS1{F0: 42, F1: 42, F2: (-(v4))}
    var v6 int8 = p0
    v7 := []int{-736}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint32 = 2
    var v2 float64 = 0.1
    v3 := make(map[int]uint32)
    v2 = v2
    if (true || (989 >= 255)) {
        v2 = float64(((0.0 + -2.25) * -2.25))
        fmt.Println(v1, v1)
    }
    var v8 uint = 42
    v8 = (v8 + (^(v8)))
    for fi9 := 0; fi9 < 17; fi9++ {
        switch "\n\t" {
        case "a":
            v10 := GenS1{F0: (-(fi9)), F1: 42, F2: -128}
            _ = v10
        }
        _ = fi9
    }
    var v11 int8 = -115
    var v12 int = 0
    if false {
        v1 = v1
    } else {
        switch -2147483648 {
        case -169:
            var v13 int = v12
            _ = v13
        case -9223372036854775808:
            switch ("hello" + "\n\t") {
            case "\n\t":
                var v14 uint8 = 127
                _ = v14
            }
        }
    }
    var v15 uint16 = 65535
    switch (-(v12)) {
    case 32767:
        switch "a" {
        case "":
            delete(v3, v12)
        case "foo bar":
            for fi16 := 0; fi16 < 8; fi16++ {
                if true {
                    var v17 int = (-(fi16))
                    v18 := GenS1{F0: v17, F1: 32767, F2: 94}
                    var v19 int64 = 1
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        }
    case -9223372036854775808:
        if ((float32(0.0) <= 0.0) && true) {
            for fi20 := 0; fi20 < 14; fi20++ {
                for fi21 := 0; fi21 < 10; fi21++ {
                    v12 = (v12 / (((fi21 * 0)) | 1))
                    _ = fi21
                }
                _ = fi20
            }
            v22 := GenS1{F0: -506, F1: 455, F2: v11}
            _ = v22
        }
    case -2147483648:
        var v23 int32 = 0
        _ = v23
    default:
        v1 = v1
    }
    v24 := GenS2{F0: 42, F1: -9223372036854775808}
    if (false && false) {
        v2 = v2
    } else {
        v25 := GenS3{F0: 176, F1: (v15 / ((v15) | 1)), F2: -2147483648}
        v26 := GenS2{F0: -868, F1: 0}
        v27 := GenS4{F0: uint8(v1), F1: (^(v8))}
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v8
}
