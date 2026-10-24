package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int64; F1 int8 }
func genFn0() int {
    var v5 uint64 = 512
    switch "foo bar" {
    case "hello":
        switch ("" + "") {
        case "foo bar":
            v5 = 1
        default:
            var v6 uint64 = 836
            _ = v6
        }
    case "\n\t":
        var v7 uint16 = 32767
        _ = v7
    case "":
        v8 := GenS1{F0: 152, F1: 0}
        _ = v8
    default:
        if (true || (9223372036854775807 >= -9223372036854775808)) {
            var v9 string = (("\n\t" + (("" + (("\n\t" + "foo bar") + "foo bar")) + "hello")) + ("a" + ""))
            v10 := []int{1, -5}
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 uint8 = 216
                if (-32768 <= 1) {
                    _ = len(v10)
                    var v13 uint32 = 233
                    v13 -= v13
                    _ = v13
                }
                _ = len(v10)
                _ = fi11
                _ = v12
            }
            _ = v10
            _ = v9
        } else {
            fmt.Println(v5)
        }
    }
    _ = v5
    return 128
}

func main() {
    var v1 int8 = 127
    var v2 uint16 = 1
    v3 := []string{"hello", "hello", "hello", ""}
    v4 := []uint32{2, 2147483647}
    v14 := GenS2{F0: 514}
    v15 := GenS3{F0: 1}
    var v16 uint64 = 1
    if (42 < genFn0()) {
        v17 := GenS4{F0: 0, F1: v1}
        v18 := GenS3{F0: 32768}
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v1 += (^(v1))
        _ = fi19
    }
    v2 = 963
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS5{F0: true}
            _ = fi21
            _ = v22
        }
        v23 := GenS6{F0: 810, F1: 0}
        _ = fi20
        _ = v23
    }
    v16 = v16
    if ((!true) && (0.0 == 3.14159)) {
        v2 -= (^(v2))
        var v24 string = "foo bar"
        _ = v24
    }
    if false {
        fmt.Println(v2)
        v2 = (v2 | v2)
    } else {
        v16 -= v16
        if false {
            v16 = 0
            for fi25 := 0; fi25 < 16; fi25++ {
                v16 -= uint64(v2)
                v16 = (v16 / ((v16) | 1))
                _ = fi25
            }
        }
        v26 := GenS7{F0: -1, F1: (v1 << ((18) & 63))}
        _ = v26
    }
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v27 float32 = 1.5
    if (42 >= 911) {
        switch -2147483648 {
        case -883:
            if true {
                v2 = (v2 - (v2 % ((v2) | 1)))
                v2 -= (v2 & v2)
            } else {
                fmt.Println(v16)
            }
        }
        var v28 uint16 = v2
        var v29 int16 = 0
        _ = v28
        _ = v29
    }
    var v30 string = ""
    var v31 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
