package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int64; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
func genFn0(p0 uint8, p1 uint, p2 uint) uint32 {
    v4 := []int{657, 32768}
    var v5 uint = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 65535
}
func genFn1() int64 {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: (fi6 ^ (fi6 & 42)), F1: 597}
        v8 := GenS2{F0: 1, F1: -2147483648, F2: 65535}
        v9 := []int{}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float64 = 0.0
    fmt.Println(v10)
    var v11 int16 = 0
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 int32 = 42
    var v2 int = -244
    var v3 int64 = 42
    v12 := GenS3{F0: 2, F1: 0}
    v1 += (v1 & (-(v1)))
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 float64 = 0.0
        var v15 uint32 = 2
        if true {
            switch "a" {
            case "hello":
                v16 := GenS4{F0: true}
                _ = v16
            default:
                fmt.Println(v1, fi13, v1)
            }
            var v17 string = (("foo bar" + "") + "foo bar")
            switch v2 {
            case 127:
                fmt.Println(v15, v14)
            case 0:
                v2 += v2
            default:
                v3 = (-(v3))
            }
            _ = v17
        } else {
            var v18 int32 = 1
            v19 := GenS1{F0: v2, F1: 449}
            _ = v18
            _ = v19
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v20 bool = true
    v21 := GenS5{F0: 65535}
    var v22 int64 = v3
    var v23 int8 = 56
    v24 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
