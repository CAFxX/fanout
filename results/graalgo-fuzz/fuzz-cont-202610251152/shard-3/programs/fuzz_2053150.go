package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int64; F1 bool; F2 uint8 }
type GenS4 struct { F0 int32; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint8 }
func genFn0() string {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: 42, F1: (false && (224 != 65535))}
        if true {
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 float32 = 1e10
                _ = fi6
                _ = v7
            }
        }
        var v8 int8 = 1
        _ = fi4
        _ = v5
        _ = v8
    }
    switch "hello" {
    case "":
        switch "a" {
        case "hello":
            var v9 uint64 = 2
            _ = v9
        case "foo bar":
            switch ("hello" + (("hello" + (("\n\t" + "hello") + "foo bar")) + "foo bar")) {
            case "foo bar":
                var v10 bool = (((!("hello" <= "hello")) || false) && true)
                _ = v10
            default:
                var v11 int64 = -488
                _ = v11
            }
        case "\n\t":
            v12 := GenS2{F0: 65535, F1: "a"}
            _ = v12
        default:
            var v13 uint32 = 0
            _ = v13
        }
    default:
        v14 := []int{869, 748}
        _ = v14
    }
    switch -317 {
    case 399:
        var v15 int = 1
        _ = v15
    case -513:
        v16 := GenS3{F0: 604, F1: false, F2: 0}
        _ = v16
    case 2147483647:
        var v17 bool = true
        _ = v17
    default:
        if (1 <= 255) {
            var v18 int8 = 102
            _ = v18
        }
    }
    var v19 int8 = -1
    var v20 int32 = int32(v19)
    _ = v19
    _ = v20
    return "\n\t"
}

func main() {
    var v1 int64 = -133
    var v2 uint16 = 143
    v3 := []int{1, 42}
    v1 -= -99
    switch "a" {
    case "a":
        fmt.Println(v2, v1)
    default:
        var v21 uint64 = 42
        _ = v21
    }
    v22 := GenS2{F0: (v2 % (((v2 & v2)) | 1)), F1: "a"}
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 uint64 = 4294967295
        _ = fi23
        _ = v24
    }
    v25 := GenS4{F0: 973, F1: 2147483647, F2: 255}
    var v26 int64 = v1
    if (false && false) {
        v26 -= -9223372036854775808
        if false {
            if true {
                for fi27 := 0; fi27 < 2; fi27++ {
                    v26 = -9223372036854775808
                    _ = fi27
                }
                switch 0 {
                case -200:
                    for fi28 := 0; fi28 < 20; fi28++ {
                        v29 := GenS1{F0: (v1 + v26), F1: false}
                        _ = fi28
                        _ = v29
                    }
                case 539:
                    v30 := GenS3{F0: v1, F1: true, F2: 239}
                    _ = v30
                }
            } else {
                v31 := GenS1{F0: v1, F1: false}
                var v32 int32 = 1
                _ = v31
                _ = v32
            }
            v2 += (^(v2))
        } else {
            for fi33 := 0; fi33 < 9; fi33++ {
                var v34 uint8 = 255
                v35 := GenS5{F0: v34, F1: 90, F2: v34}
                _ = fi33
                _ = v34
                _ = v35
            }
            var v36 int = -2147483648
            _ = v36
        }
        if (v2 == v2) {
            var v37 uint64 = 0
            _ = len(v3)
            var v38 int64 = v26
            _ = v37
            _ = v38
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
}
