package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int; F1 int16; F2 int8 }
type GenS4 struct { F0 uint32; F1 string; F2 uint }
type GenS5 struct { F0 int64; F1 uint64 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]bool)
    var v3 string = (("hello" + ("foo bar" + "")) + "")
    v4 := make(map[string]uint32)
    fmt.Println(v3, v1)
    v5 := GenS1{F0: -488, F1: 332, F2: -32768}
    var v6 uint8 = (^(v1))
    var v7 int8 = 100
    fmt.Println(v3)
    switch (v3 + "\n\t") {
    case "a":
        if true {
            var v8 int = 1
            _ = v8
        } else {
            fmt.Println(v6)
            v9 := GenS1{F0: 42, F1: uint32(v6), F2: -32768}
            _ = v9
        }
    default:
        var v10 bool = true
        _ = v10
    }
    var v11 uint64 = 0
    v12 := []int{-1, 9223372036854775807}
    v13 := GenS2{F0: 42, F1: -1}
    v7 = 121
    fmt.Println(v7)
    if true {
        v12 = append(v12, 32768)
        var v14 int32 = 2147483647
        _ = v14
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS3{F0: fi16, F1: 42, F2: v7}
                if false {
                    if false {
                        var v18 string = v3
                        v19 := GenS4{F0: 42, F1: "a", F2: 65535}
                        _ = len(v12)
                        _ = v18
                        _ = v19
                    } else {
                        v7 += v7
                    }
                }
                _ = fi16
                _ = v17
            }
            if false {
                switch -898 {
                case 662:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 int = 32768
                        _ = fi20
                        _ = v21
                    }
                case -609:
                    var v22 float32 = 3.14159
                    _ = v22
                }
                var v23 uint16 = 2
                v24 := GenS5{F0: 1, F1: v11}
                _ = v23
                _ = v24
            } else {
                fmt.Println(fi15)
                switch (fi15 << ((1) & 63)) {
                case -1:
                    v25 := GenS1{F0: fi15, F1: 2, F2: 282}
                    _ = v25
                case 656:
                    switch (fi15 & (^(fi15))) {
                    case -9223372036854775808:
                        var v26 int32 = -2147483648
                        _ = v26
                    case -957:
                        delete(v2, 2147483648)
                    default:
                        var v27 float32 = 1e10
                        _ = v27
                    }
                default:
                    var v28 int8 = 117
                    _ = v28
                }
                v29 := GenS3{F0: (fi15 - fi15), F1: 32767, F2: v7}
                _ = v29
            }
            _ = fi15
        }
        switch v3 {
        case "\n\t":
            if false {
                fmt.Println(v1, v11)
                var v30 float32 = (((0.0 - ((0.1 + 0.1) + 0.0)) + 0.0) + 0.1)
                v4[("a" + v3)] = 42
                _ = v30
            }
        case "":
            v1 = 0
        case "foo bar":
            for fi31 := 0; fi31 < 7; fi31++ {
                _ = len(v4)
                for fi32 := 0; fi32 < 9; fi32++ {
                    if (false || (272 < -1)) {
                        fmt.Println(v6, fi32, v6)
                        fmt.Println(v6, fi31, v6)
                        _ = len(v2)
                    } else {
                        v1 = (^(v6))
                        var v33 uint8 = 128
                        var v34 int64 = 42
                        _ = v33
                        _ = v34
                    }
                    v12 = append(v12, (-(fi32)))
                    var v35 uint16 = 593
                    _ = fi32
                    _ = v35
                }
                _ = fi31
            }
        }
    }
    fmt.Println(v11)
    for fi36 := 0; fi36 < 3; fi36++ {
        v7 = (v7 * v7)
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
