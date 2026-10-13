package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int8; F1 uint8; F2 int }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint) bool {
    v4 := GenS1{F0: 2, F1: p0}
    var v5 int64 = -1
    switch 127 {
    case 32768:
        var v6 uint64 = 313
        _ = v6
    default:
        v5 += v5
    }
    p0 = 2
    _ = p0
    _ = v4
    _ = v5
    return (59 > -128)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 983
    var v3 uint = 4294967295
    v7 := GenS1{F0: 4294967295, F1: (v3 / ((v3) | 1))}
    var v8 int32 = 0
    v2 = (v2 + v2)
    v8 = v8
    for fi9 := 0; fi9 < 9; fi9++ {
        if false {
            v10 := GenS2{F0: 100, F1: 66, F2: fi9}
            var v11 uint64 = 1
            _ = v10
            _ = v11
        } else {
            var v12 float64 = 0.1
            _ = v12
        }
        fmt.Println(v2, fi9, fi9)
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(v8, fi13)
            fmt.Println(fi13)
            v8 = v8
            _ = fi13
        }
        _ = fi9
    }
    switch 923 {
    case -9223372036854775808:
        switch "hello" {
        case "":
            v14 := GenS3{F0: (v3 & v3), F1: 489}
            _ = v14
        case "a":
            var v15 uint = v3
            _ = v15
        case "hello":
            switch 1 {
            case 128:
                v16 := GenS2{F0: 127, F1: 193, F2: 488}
                _ = v16
            case 726:
                v3 = v3
            case 32767:
                for fi17 := 0; fi17 < 5; fi17++ {
                    fmt.Println(v3)
                    v18 := GenS4{F0: 42}
                    v3 += uint(v2)
                    _ = fi17
                    _ = v18
                }
            }
        }
    case 1:
        var v19 uint16 = 833
        _ = v19
    case 32768:
        v20 := GenS2{F0: 79, F1: 127, F2: 2147483647}
        _ = v20
    }
    if true {
        if ((v3 | (v3 % (((v3 % ((v3) | 1))) | 1))) >= 65535) {
            var v21 uint32 = 65535
            v2 -= 18446744073709551615
            v2 = v2
            _ = v21
        }
    }
    v22 := v1[-2147483648]
    switch "foo bar" {
    case "a":
        v23 := []int{}
        _ = v23
    case "\n\t":
        v22 = v22
    }
    if false {
        var v24 int8 = -1
        _ = v24
    } else {
        var v25 int16 = -423
        _ = v25
    }
    if true {
        delete(v1, -711)
        fmt.Println(v2, v3)
    }
    v2 += v2
    switch "hello" {
    case "foo bar":
        switch ("hello" + "\n\t") {
        case "foo bar":
            v2 += 42
        case "":
            var v26 uint32 = v22
            _ = v26
        }
    case "\n\t":
        v22 = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v7
    _ = v8
}
