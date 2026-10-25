package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint64 }

func main() {
    var v1 int16 = -713
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 int64 = -431
    delete(v3, "")
    for fi5 := 0; fi5 < 3; fi5++ {
        switch "hello" {
        case "hello":
            var v6 uint32 = 801
            _ = v6
        case "\n\t":
            v7 := GenS1{F0: (fi5 | fi5), F1: false, F2: (fi5 | 128)}
            _ = v7
        }
        _ = fi5
    }
    var v8 int = 4294967295
    if (412 < 312) {
        switch v8 {
        case -9223372036854775808:
            if false {
                var v9 uint16 = 65535
                _ = v9
            }
        case 0:
            fmt.Println(v4, v8, v1)
        case 842:
            v4 = v4
        default:
            for fi10 := 0; fi10 < 18; fi10++ {
                v4 -= (v4 + v4)
                if true {
                    v11 := GenS1{F0: v8, F1: (207 < 2), F2: 0}
                    fmt.Println(fi10, v8)
                    _ = v11
                }
                _ = fi10
            }
        }
    }
    v8 = (v8 | 984)
    v12 := GenS1{F0: -297, F1: false, F2: (^(v8))}
    v13 := GenS2{F0: 65535}
    v14 := GenS2{F0: 46}
    var v15 uint8 = 127
    v15 += v15
    var v16 int = v8
    for fi17 := 0; fi17 < 4; fi17++ {
        switch "" {
        case "hello":
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := []int{(-(fi18)), 32768, (v16 - (fi18 / ((v8) | 1))), fi17}
                _ = fi18
                _ = v19
            }
        case "a":
            v20 := GenS2{F0: 189}
            _ = v20
        case "":
            switch (fi17 + v8) {
            case 9223372036854775807:
                v21 := GenS1{F0: fi17, F1: ("" < (("hello" + "") + ("foo bar" + "hello"))), F2: (fi17 | (fi17 | (v8 * (fi17 << ((319) & 63)))))}
                _ = v21
            default:
                for fi22 := 0; fi22 < 11; fi22++ {
                    v4 = v4
                    if (-716 >= 0) {
                        v4 = v4
                    }
                    if true {
                        fmt.Println(v4, v8, v15)
                        v23 := []int{698}
                        fmt.Println(v15, v1)
                        fmt.Println(v15, v1)
                        _ = v23
                    } else {
                        v15 = (v15 & (v15 % (((^(v15))) | 1)))
                        v24 := []int{}
                        _ = v24
                    }
                    _ = fi22
                }
            }
        }
        _ = fi17
    }
    v15 = v15
    v25 := []int{v16}
    var v26 int = (v16 + 343)
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 int = fi27
        for fi29 := 0; fi29 < 17; fi29++ {
            if ((fi27 & (fi27 & v8)) < (v26 & fi29)) {
                v30 := GenS3{F0: 100, F1: 1}
                _ = len(v25)
                fmt.Println(fi27, v8)
                _ = v30
            } else {
                var v31 uint8 = (^(v15))
                fmt.Println(fi29, fi27)
                _ = v31
            }
            var v32 float32 = (0.1 * 1e10)
            _ = fi29
            _ = v32
        }
        v15 -= 255
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v8
}
