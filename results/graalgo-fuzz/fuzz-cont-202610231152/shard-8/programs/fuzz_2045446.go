package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32; F1 int; F2 int64 }
type GenS8 struct { F0 int8; F1 uint8; F2 bool }
type GenS9 struct { F0 int32; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint }
func genFn0(p0 uint32) int8 {
    v4 := []int{}
    p0 = (p0 % ((p0) | 1))
    p0 = (p0 * p0)
    if ("foo bar" == "hello") {
        v5 := GenS1{F0: (p0 % ((p0) | 1)), F1: 242}
        v6 := GenS2{F0: 42, F1: 42}
        _ = v5
        _ = v6
    } else {
        _ = len(v4)
        p0 -= 2147483648
    }
    _ = p0
    _ = v4
    return 127
}
func genFn1() string {
    var v7 int16 = -1
    var v8 int16 = (v7 + v7)
    switch "" {
    case "":
        var v9 uint8 = 128
        _ = v9
    default:
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS3{F0: 0, F1: false, F2: (fi10 % (((fi10 | (fi10 & 42))) | 1))}
            _ = fi10
            _ = v11
        }
    }
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    var v1 uint = 255
    var v2 int = 132
    v3 := make(map[int]string)
    if true {
        v1 -= v1
        var v12 int64 = 0
        _ = v12
    }
    fmt.Println(v1)
    var v13 uint8 = 0
    v14 := GenS4{F0: ((!true) || true), F1: -1}
    fmt.Println(v1, v13, v2)
    if false {
        var v15 uint = (v1 + v1)
        if (456 > 1) {
            v1 = v15
            v16 := GenS5{F0: 42, F1: "\n\t"}
            _ = v16
        } else {
            switch "a" {
            case "foo bar":
                for fi17 := 0; fi17 < 6; fi17++ {
                    if true {
                        v1 = 42
                    }
                    var v18 float64 = -2.25
                    _ = fi17
                    _ = v18
                }
            case "\n\t":
                v13 = (v13 * v13)
            default:
                for fi19 := 0; fi19 < 1; fi19++ {
                    v20 := []int{v2, (fi19 & (fi19 / ((32767) | 1))), fi19, (v2 % ((v2) | 1))}
                    v15 = v1
                    fmt.Println(v13, v1)
                    _ = fi19
                    _ = v20
                }
            }
            var v21 uint16 = 715
            if true {
                v13 += v13
                v22 := GenS6{F0: 32767}
                _ = v22
            }
            _ = v21
        }
        fmt.Println(v15, v15)
        _ = v15
    } else {
        for fi23 := 0; fi23 < 12; fi23++ {
            v3[(v2 << ((0) & 63))] = "hello"
            switch -3 {
            case 4294967295:
                v2 = -870
            case -2147483648:
                v24 := GenS7{F0: 2147483648, F1: (fi23 | -384), F2: 1}
                _ = v24
            case 128:
                v13 = v13
            }
            _ = fi23
        }
    }
    var v25 int8 = 1
    v26 := []int{(v2 - v2), (v2 - v2), v2}
    switch (^(v2)) {
    case 2147483647:
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    case 32768:
        v27 := GenS2{F0: 0, F1: uint64(v13)}
        _ = v27
    case -1:
        for fi28 := 0; fi28 < 2; fi28++ {
            fmt.Println(v25)
            _ = fi28
        }
    default:
        v25 = (v25 | v25)
    }
    v1 = 65535
    v2 -= (v2 - (^(v2)))
    if false {
        switch v2 {
        case 1:
            if false {
                var v29 uint = v1
                _ = v29
            }
        default:
            var v30 int8 = (v25 & (^(v25)))
            _ = v30
        }
        if false {
            _ = len(v3)
            v31 := GenS8{F0: v25, F1: v13, F2: false}
            v32 := GenS9{F0: 719, F1: genFn0(92), F2: 32767}
            _ = v31
            _ = v32
        }
    } else {
        v33 := GenS10{F0: v1}
        v34 := v3[-9223372036854775808]
        switch ((v34 + v34) + v34) {
        case "\n\t":
            var v35 uint32 = 2147483647
            _ = v35
        default:
            fmt.Println(v1)
        }
        _ = v33
        _ = v34
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        var v37 int8 = v25
        v37 -= v25
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        v2 = (^(fi38))
        var v39 uint64 = 18446744073709551615
        _ = fi38
        _ = v39
    }
    v26 = append(v26, 127)
    var v40 float64 = float64(float32(1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v40
}
