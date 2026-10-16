package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
func genFn0() uint32 {
    var v5 int8 = -6
    switch "a" {
    case "":
        var v6 uint16 = 0
        _ = v6
    }
    var v7 int32 = 85
    switch 4294967295 {
    case -962:
        var v8 int64 = int64(v5)
        _ = v8
    }
    _ = v5
    _ = v7
    return 280
}
func genFn1(p0 uint8) bool {
    if (!((false && (false && (!(42 < 65535)))) && false)) {
        var v9 int8 = 127
        switch "a" {
        case "\n\t":
            var v10 int32 = 249
            _ = v10
        case "a":
            switch 42 {
            case -9223372036854775808:
                v11 := GenS1{F0: 406, F1: "foo bar", F2: -101}
                _ = v11
            case 59:
                var v12 float32 = ((float32(0.0) * 1.5) - 1e10)
                _ = v12
            case 912:
                fmt.Println(v9)
            }
        case "foo bar":
            v13 := GenS1{F0: 1, F1: "a", F2: (v9 >> ((923) & 63))}
            _ = v13
        default:
            p0 += 127
        }
        var v14 int8 = (v9 + 42)
        _ = v14
        _ = v9
    }
    v15 := []int{-1, 1, 9223372036854775807}
    _ = p0
    _ = v15
    return (1000 > -623)
}

func main() {
    var v1 uint32 = 1
    var v2 uint = 255
    v3 := make(map[int]string)
    v4 := []int{128, 577, 128}
    v16 := GenS1{F0: 0, F1: "\n\t", F2: -59}
    switch "\n\t" {
    case "foo bar":
        v2 += (^(v2))
    case "hello":
        if false {
            v17 := GenS1{F0: 209, F1: "", F2: 100}
            _ = v17
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    if (!false) {
        v4 = append(v4, 334)
    }
    v1 = v1
    v18 := GenS2{F0: -1, F1: 439}
    var v19 string = "a"
    fmt.Println(v1, v2, v1)
    for fi20 := 0; fi20 < 9; fi20++ {
        if false {
            v21 := GenS1{F0: 2, F1: ("" + "foo bar"), F2: -13}
            for fi22 := 0; fi22 < 14; fi22++ {
                v4 = append(v4, (-(fi20)))
                v23 := GenS1{F0: 32767, F1: v19, F2: -128}
                _ = fi22
                _ = v23
            }
            var v24 string = "foo bar"
            _ = v21
            _ = v24
        } else {
            v25 := GenS1{F0: 42, F1: "a", F2: -6}
            v26 := GenS3{F0: -833}
            var v27 uint8 = 255
            _ = v25
            _ = v26
            _ = v27
        }
        v28 := GenS4{F0: 216}
        _ = fi20
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
