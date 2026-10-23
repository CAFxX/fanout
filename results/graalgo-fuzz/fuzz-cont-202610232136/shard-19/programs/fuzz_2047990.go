package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 uint8; F1 string; F2 int16 }
func genFn0(p0 int) bool {
    var v6 uint64 = 18446744073709551615
    if (((2 > 4294967295) || (32767 == -2147483648)) || true) {
        var v7 int16 = 32767
        if false {
            if true {
                switch "foo bar" {
                case "":
                    fmt.Println(v7, v7, v6)
                case "a":
                    if (42 <= 1) {
                        var v8 uint8 = uint8(v6)
                        switch p0 {
                        case -9223372036854775808:
                            v9 := []int{p0, (p0 - p0), p0}
                            _ = v9
                        }
                        _ = v8
                    }
                case "\n\t":
                    v7 = -1
                default:
                    v10 := GenS1{F0: 1}
                    _ = v10
                }
                var v11 int8 = -68
                v12 := []int{2147483647, (p0 % (((p0 << ((1) & 63))) | 1)), (p0 + p0)}
                _ = v11
                _ = v12
            }
            if (false || (269 != 1000)) {
                v13 := GenS2{F0: 127, F1: "hello"}
                _ = v13
            } else {
                switch "foo bar" {
                case "hello":
                    if false {
                        var v14 uint16 = 42
                        var v15 string = "hello"
                        _ = v14
                        _ = v15
                    }
                }
                v16 := GenS3{F0: 2, F1: (-(v7))}
                _ = v16
            }
        } else {
            var v17 int8 = 103
            _ = v17
        }
        _ = v7
    } else {
        v18 := []int{p0}
        _ = v18
    }
    var v19 int64 = -1
    switch "" {
    case "foo bar":
        var v20 uint16 = 69
        _ = v20
    case "hello":
        var v21 float32 = (3.14159 + float32(float64((3.14159 * -2.25))))
        _ = v21
    case "a":
        var v22 int64 = (-(v19))
        _ = v22
    default:
        var v23 uint16 = 255
        _ = v23
    }
    v24 := GenS4{F0: 9223372036854775807}
    _ = p0
    _ = v19
    _ = v24
    _ = v6
    return false
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{127, 789, 127, 2147483647}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v5 int64 = 0
    v25 := GenS5{F0: -742, F1: 0}
    var v26 uint8 = 128
    switch ("\n\t" + ("a" + "foo bar")) {
    case "foo bar":
        _ = len(v2)
    default:
        v27 := GenS6{F0: 206, F1: 255}
        _ = v27
    }
    fmt.Println(v5, v26)
    switch 9223372036854775807 {
    case 809:
        for fi28 := 0; fi28 < 11; fi28++ {
            v26 += (v26 % (((^(v26))) | 1))
            _ = fi28
        }
    case 128:
        if false {
            v26 += (v26 << ((973) & 63))
        }
    case 1:
        v29 := GenS7{F0: (v26 ^ (v26 << ((1) & 63))), F1: "", F2: 1000}
        _ = v29
    }
    fmt.Println(v5, v5)
    v5 -= (-(v5))
    switch "hello" {
    case "":
        v5 = (v5 << ((324) & 63))
    case "hello":
        _ = len(v2)
    case "\n\t":
        switch -591 {
        case -1:
            v26 += v26
        default:
            fmt.Println(v26)
        }
    }
    if (!(false && true)) {
        for fi30 := 0; fi30 < 11; fi30++ {
            _ = len(v1)
            _ = fi30
        }
    }
    _ = len(v1)
    var v31 string = "a"
    fmt.Println(v31, v5, v26)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v32 int8 = 127
    v4[((v31 + ((v31 + (((v31 + "foo bar") + ("\n\t" + v31)) + v31)) + v31)) + "\n\t")] = 924
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
