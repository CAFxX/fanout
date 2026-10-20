package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 int32) int64 {
    switch -275 {
    case 32767:
        for fi4 := 0; fi4 < 4; fi4++ {
            v5 := []int{(fi4 * 42), fi4, -839}
            for fi6 := 0; fi6 < 13; fi6++ {
                if len(v5) > 0 {
                    v5[0] = (^(fi6))
                }
                p0 = (p0 | 42)
                _ = fi6
            }
            _ = fi4
            _ = v5
        }
    }
    v7 := GenS1{F0: 255, F1: 42}
    fmt.Println(p0)
    var v8 int64 = -9223372036854775808
    v9 := GenS2{F0: 42}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 uint, p1 int64) int8 {
    var v10 bool = (127 > 100)
    switch ("a" + "hello") {
    case "foo bar":
        v11 := GenS1{F0: 2, F1: 42}
        _ = v11
    case "":
        var v12 float64 = 3.14159
        _ = v12
    case "\n\t":
        v13 := []int{42, 128, -1, 749}
        _ = v13
    default:
        switch "a" {
        case "\n\t":
            v14 := GenS2{F0: 1}
            _ = v14
        case "hello":
            p1 -= genFn0(2147483647)
        }
    }
    var v15 uint8 = 1
    v16 := GenS3{F0: ("\n\t" + ("a" + ("hello" + "\n\t"))), F1: -474}
    var v17 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    return 65
}

func main() {
    v1 := []uint32{2, 2147483647, 34}
    v2 := make(map[uint32]bool)
    var v3 int32 = 2147483647
    v3 += (v3 + v3)
    if true {
        delete(v2, 2)
        var v18 uint8 = 42
        _ = v18
    }
    v3 = 1
    switch -964 {
    case 1:
        if false {
            for fi19 := 0; fi19 < 7; fi19++ {
                for fi20 := 0; fi20 < 10; fi20++ {
                    fmt.Println(fi19, v3)
                    v21 := GenS1{F0: 2, F1: 1}
                    _ = fi20
                    _ = v21
                }
                var v22 bool = false
                _ = fi19
                _ = v22
            }
        } else {
            var v23 int8 = 100
            _ = v23
        }
    }
    fmt.Println(v3)
    _ = len(v1)
    if (-9223372036854775808 >= 0) {
        fmt.Println(v3)
    }
    if (-119 < 1) {
        if false {
            fmt.Println(v3)
        }
        v3 = (v3 | v3)
    }
    switch "a" {
    case "":
        v3 = 42
    case "a":
        v24 := v2[2]
        _ = v24
    }
    var v25 int32 = 0
    if false {
        _ = len(v2)
        v1 = append(v1, 694)
    } else {
        if len(v1) > 0 {
            v1[0] = 0
        }
        v25 -= -1
        v25 += -1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
}
