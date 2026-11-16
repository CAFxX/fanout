package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint16, p1 int64) string {
    v7 := []int{}
    var v8 int8 = 0
    var v9 uint32 = 2147483648
    switch -527 {
    case -9223372036854775808:
        v10 := GenS1{F0: 9223372036854775807, F1: ("a" + "foo bar")}
        _ = v10
    case 128:
        v11 := GenS1{F0: -794, F1: ("hello" + "\n\t")}
        _ = v11
    }
    v7 = append(v7, 551)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}
func genFn1() int64 {
    v12 := []int{943, 710, -894}
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 string = "hello"
        v14 = v14
        _ = fi13
        _ = v14
    }
    _ = v12
    return 1
}

func main() {
    var v1 uint64 = 23
    v2 := []int{9223372036854775807}
    v3 := []string{""}
    var v4 int8 = 28
    v5 := make(map[int]uint32)
    var v6 uint = 255
    switch -740 {
    case 121:
        v1 -= v1
    }
    v15 := GenS2{F0: 181, F1: 9223372036854775807}
    var v16 uint16 = 255
    fmt.Println(v1, v16, v1)
    v17 := GenS1{F0: -1, F1: "foo bar"}
    switch ("foo bar" + ("a" + "")) {
    case "":
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := GenS2{F0: 164, F1: (-(fi18))}
            _ = fi18
            _ = v19
        }
    case "hello":
        for fi20 := 0; fi20 < 10; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                fmt.Println(v16)
                _ = fi21
            }
            _ = fi20
        }
    case "a":
        v22 := GenS1{F0: int(v4), F1: "hello"}
        _ = v22
    default:
        _ = len(v5)
    }
    switch "a" {
    case "foo bar":
        for fi23 := 0; fi23 < 14; fi23++ {
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 uint = v6
                _ = fi24
                _ = v25
            }
            if true {
                v26 := GenS1{F0: (fi23 + fi23), F1: "a"}
                _ = len(v3)
                var v27 uint64 = (v1 & v1)
                _ = v26
                _ = v27
            }
            _ = fi23
        }
    case "hello":
        if (!true) {
            var v28 int32 = 1
            switch -493 {
            case 1:
                v4 -= int8(v16)
            default:
                var v29 string = ""
                _ = v29
            }
            _ = v28
        } else {
            v16 = 255
            v16 += (v16 % ((uint16(v1)) | 1))
            var v30 float64 = 1e10
            _ = v30
        }
    case "a":
        v6 += 65535
    }
    var v31 uint = 65535
    var v32 int8 = (v4 & v4)
    var v33 uint64 = (^(v1))
    switch "foo bar" {
    case "":
        v34 := GenS3{F0: 99, F1: uint8(v33)}
        _ = v34
    }
    v33 += (v33 & (v1 - (v1 ^ v33)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
