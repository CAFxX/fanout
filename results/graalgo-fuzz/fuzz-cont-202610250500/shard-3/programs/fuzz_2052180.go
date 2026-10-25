package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int64; F2 string }
type GenS4 struct { F0 int; F1 int; F2 int16 }
type GenS5 struct { F0 string; F1 uint64; F2 int16 }

func main() {
    v1 := []int{}
    v2 := make(map[int]int)
    var v3 int8 = 109
    v4 := make(map[uint32]uint32)
    v5 := []int{127, -2147483648}
    delete(v4, 2147483647)
    switch ("foo bar" + (((("\n\t" + ("\n\t" + "\n\t")) + ("hello" + ("foo bar" + ("" + ("a" + "\n\t"))))) + "") + "a")) {
    case "foo bar":
        var v6 int16 = 597
        _ = v6
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
    }
    var v7 int = -1
    var v8 uint32 = 42
    for fi9 := 0; fi9 < 10; fi9++ {
        v3 += int8(v7)
        var v10 int64 = 703
        switch "hello" {
        case "foo bar":
            var v11 int32 = -15
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        switch ("hello" + (("foo bar" + "foo bar") + "hello")) {
        case "hello":
            if (!(v8 >= (v8 | 970))) {
                var v13 int8 = (v3 + (v3 & (v3 | v3)))
                _ = len(v5)
                fmt.Println(v8, v7, v7)
                _ = v13
            }
        default:
            switch 32768 {
            case 32767:
                if (false && true) {
                    v14 := GenS1{F0: (v3 * 100), F1: (v3 % (((v3 & v3)) | 1))}
                    v15 := GenS2{F0: 1}
                    _ = v14
                    _ = v15
                } else {
                    var v16 int8 = (^(v3))
                    for fi17 := 0; fi17 < 15; fi17++ {
                        v1 = append(v1, 9223372036854775807)
                        v18 := GenS3{F0: false, F1: -1, F2: "foo bar"}
                        v19 := GenS2{F0: 93}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    var v20 int8 = (v3 % (((^(v3))) | 1))
                    _ = v16
                    _ = v20
                }
            case 2147483647:
                var v21 int = v7
                _ = v21
            }
        }
        v3 = v3
        _ = fi12
    }
    v22 := GenS4{F0: v7, F1: -9223372036854775808, F2: int16(v3)}
    for fi23 := 0; fi23 < 9; fi23++ {
        v24 := GenS5{F0: "", F1: 0, F2: -1}
        _ = fi23
        _ = v24
    }
    var v25 uint64 = 4294967295
    if ((!true) && true) {
        fmt.Println(v25, v25, v8)
        fmt.Println(v25)
    } else {
        var v26 int64 = 530
        fmt.Println(v25, v26, v8)
        v8 -= 42
        _ = v26
    }
    v4[v8] = (v8 / (((v8 - (v8 / (((v8 % (((v8 / (((v8 + v8)) | 1))) | 1))) | 1)))) | 1))
    switch v7 {
    case -355:
        if true {
            delete(v2, v7)
            v8 = (v8 << ((417) & 63))
            var v27 int16 = 0
            _ = v27
        } else {
            v7 = v7
            fmt.Println(v8)
        }
    case 128:
        fmt.Println(v25, v25)
    case -9223372036854775808:
        for fi28 := 0; fi28 < 14; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 bool = false
                var v31 bool = false
                v30 = false
                _ = fi29
                _ = v30
                _ = v31
            }
            v7 = (v7 << ((547) & 63))
            _ = fi28
        }
    }
    _ = len(v5)
    var v32 int32 = 2147483647
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    if true {
        for fi33 := 0; fi33 < 20; fi33++ {
            if false {
                fmt.Println(v3, v7)
            }
            v8 = v8
            _ = fi33
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
