package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 string; F1 uint16; F2 uint32 }
func genFn0(p0 uint64) int64 {
    var v5 int8 = 42
    v6 := []int{32768}
    _ = p0
    _ = v5
    _ = v6
    return int64(p0)
}
func genFn1(p0 uint32, p1 uint64, p2 int16) int8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 = 2
        p0 = p0
        _ = fi7
    }
    if (103 == 0) {
        p0 -= p0
        v8 := []int{9223372036854775807, 127, 4294967295}
        _ = v8
    }
    var v9 uint64 = (p1 / (((p1 << ((98) & 63))) | 1))
    v9 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 100
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]uint32)
    var v3 uint = 65535
    v4 := []uint32{4294967295, 42, 494}
    if true {
        if false {
            v10 := GenS1{F0: (true && false), F1: 127, F2: 1}
            _ = v10
        }
        switch "hello" {
        case "a":
            var v11 string = ""
            _ = v11
        case "foo bar":
            var v12 float32 = 1e10
            _ = v12
        case "hello":
            fmt.Println(v1, v1)
        default:
            v1 = -1
        }
    } else {
        var v13 uint = v3
        v14 := GenS1{F0: false, F1: -3, F2: 65535}
        _ = v13
        _ = v14
    }
    var v15 uint8 = 127
    var v16 uint16 = 42
    v4 = append(v4, 0)
    switch ("hello" + "") {
    case "":
        for fi17 := 0; fi17 < 19; fi17++ {
            v16 = v16
            _ = fi17
        }
    case "a":
        if false {
            v18 := GenS1{F0: (true && false), F1: 14, F2: uint16(v3)}
            _ = v18
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            if (true && false) {
                var v21 int16 = 658
                if false {
                    v1 = v1
                    if false {
                        fmt.Println(v1)
                        fmt.Println(fi20, v16)
                    }
                    v22 := GenS1{F0: true, F1: 4, F2: uint16(v15)}
                    _ = v22
                }
                _ = v21
            }
            _ = fi20
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            v16 = 1
            _ = fi23
        }
        for fi24 := 0; fi24 < 18; fi24++ {
            if true {
                if (3.14159 <= 1e10) {
                    v16 -= (^(v16))
                    v3 -= v3
                }
            }
            v3 = 2
            _ = fi24
        }
        _ = fi19
    }
    switch "\n\t" {
    case "\n\t":
        v3 = (v3 % (((v3 / (((v3 + v3)) | 1))) | 1))
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v1 = v1
        _ = fi25
    }
    var v26 float64 = -2.25
    v4 = append(v4, 0)
    v27 := GenS2{F0: "\n\t", F1: v16, F2: 2}
    v28 := GenS1{F0: (125 <= 127), F1: int8(v15), F2: v16}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
