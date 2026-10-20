package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 string }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 uint, p1 uint32) int8 {
    if (1 > 2) {
        var v4 bool = (false || false)
        if (65535 > 790) {
            v5 := []int{}
            v4 = true
            _ = v5
        } else {
            var v6 int16 = 1
            _ = v6
        }
        _ = v4
    }
    v7 := []int{9223372036854775807, -2147483648, -2147483648, 9223372036854775807}
    var v8 uint16 = 0
    p1 -= 1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -109
}
func genFn1() int16 {
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: "", F1: 49, F2: 127}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 1, F1: 412, F2: ("foo bar" + ("" + "foo bar"))}
    switch "\n\t" {
    case "a":
        v12 := []int{}
        _ = v12
    case "foo bar":
        if (((0 < 255) || false) && true) {
            switch 649 {
            case 0:
                v13 := []int{-2147483648}
                _ = v13
            case 9223372036854775807:
                v14 := []int{32767, -1, 42, 1}
                _ = v14
            case 4294967295:
                var v15 float32 = 0.0
                _ = v15
            }
        }
    case "":
        var v16 uint8 = 255
        _ = v16
    default:
        if false {
            var v17 uint = 0
            fmt.Println(v17)
            v17 += v17
            _ = v17
        } else {
            var v18 int32 = 0
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := []int{805, (-(fi19)), fi19}
                v21 := GenS1{F0: ("a" + ("\n\t" + ("" + "hello"))), F1: -35, F2: 116}
                fmt.Println(fi19, fi19)
                _ = fi19
                _ = v20
                _ = v21
            }
            v22 := GenS3{F0: -810, F1: ("a" >= ("a" + "hello"))}
            _ = v18
            _ = v22
        }
    }
    if (4294967295 != 4294967295) {
        var v23 uint = 255
        _ = v23
    } else {
        var v24 int = 127
        _ = v24
    }
    _ = v11
    return -918
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]string)
    var v3 int = -739
    _ = len(v2)
    v25 := []int{v3}
    v3 += v3
    switch v3 {
    case 128:
        var v26 int16 = (v1 << ((255) & 63))
        _ = v26
    default:
        v2[42] = (("foo bar" + ("a" + "")) + "")
    }
    v2[v3] = "hello"
    switch v3 {
    case -949:
        v27 := GenS3{F0: 2147483647, F1: false}
        _ = v27
    default:
        v3 = (v3 + v3)
    }
    if len(v25) > 0 {
        v25[0] = 32767
    }
    v1 += (v1 << ((796) & 63))
    v28 := GenS4{F0: ("hello" + ("hello" + "foo bar")), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
}
