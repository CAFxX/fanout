package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int8; F1 string; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int64 = 1
    var v2 uint16 = 65535
    var v3 int32 = 2147483647
    v4 := make(map[int]int)
    var v5 uint16 = v2
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: 0, F1: (v5 << ((4294967295) & 63))}
    v8 := []int{}
    if true {
        switch "" {
        case "a":
            var v9 int8 = -1
            _ = v9
        case "hello":
            v10 := GenS1{F0: -171, F1: v5}
            _ = v10
        case "":
            if false {
                var v11 uint = 4294967295
                switch "hello" {
                case "\n\t":
                    fmt.Println(v11)
                }
                _ = v11
            }
        default:
            v3 = (v3 & v3)
        }
        if false {
            v12 := GenS2{F0: "", F1: "foo bar"}
            v13 := GenS2{F0: "hello", F1: "\n\t"}
            v1 = (-(v1))
            _ = v12
            _ = v13
        } else {
            var v14 uint64 = 0
            _ = v14
        }
    }
    var v15 uint8 = 151
    if true {
        v16 := GenS3{F0: 104, F1: "hello", F2: 42}
        _ = v16
    } else {
        var v17 uint32 = 2147483647
        _ = v17
    }
    v18 := GenS4{F0: v15}
    v19 := GenS5{F0: ((v1 >> ((325) & 63)) > v1), F1: 1}
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS2{F0: "a", F1: ("a" + ("" + ("hello" + "a")))}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 12; fi23++ {
            if true {
                v24 := GenS5{F0: true, F1: 65535}
                var v25 string = "a"
                var v26 float64 = 0.1
                v6[v25] = 769
                var v27 uint32 = 4294967295
                _ = v24
                _ = v25
                _ = v26
                _ = v27
            } else {
                fmt.Println(v5, v1, v2)
                var v28 float64 = 3.14159
                var v29 float64 = 3.14159
                var v30 uint8 = (v15 / (((v15 | v15)) | 1))
                v15 = v30
                _ = v28
                _ = v29
                _ = v30
            }
            _ = len(v8)
            switch (fi22 / ((fi22) | 1)) {
            case 937:
                switch -9223372036854775808 {
                case -9223372036854775808:
                    fmt.Println(fi23, v2, fi23)
                case 127:
                    var v31 uint16 = v2
                    _ = v31
                default:
                    var v32 uint = 1
                    _ = v32
                }
            case 619:
                fmt.Println(fi22, v3, v1)
            case 4294967295:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            _ = fi23
        }
        v2 -= 456
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
