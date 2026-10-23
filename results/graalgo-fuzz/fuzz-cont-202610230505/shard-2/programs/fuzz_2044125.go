package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 string }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint, p2 uint16) bool {
    if false {
        var v5 uint64 = p0
        var v6 bool = (((1.5 / 1.0) + 0.0) != 0.1)
        _ = v5
        _ = v6
    } else {
        p0 += (p0 + (p0 ^ (p0 * p0)))
    }
    var v7 uint8 = 128
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{fi8}
        switch "\n\t" {
        case "a":
            switch "a" {
            case "":
                if false {
                    var v10 uint64 = (^(p0))
                    v11 := GenS1{F0: 42, F1: 42, F2: "a"}
                    v12 := GenS2{F0: (0 < 4294967295), F1: 985}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
            case "foo bar":
                v13 := GenS3{F0: 1}
                _ = v13
            case "hello":
                var v14 bool = (!(false && (729 == -2147483648)))
                _ = v14
            }
        }
        _ = fi8
        _ = v9
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(p1)
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (("a" < ("foo bar" + "foo bar")) && false)
}

func main() {
    var v1 string = "hello"
    var v2 int8 = 42
    var v3 int16 = 0
    v4 := make(map[uint32]string)
    if ((-2147483648 >= -1) || (!(!true))) {
        var v16 uint = 4294967295
        switch 4294967295 {
        case 128:
            var v17 int8 = (-(v2))
            _ = v17
        default:
            if true {
                var v18 uint = v16
                v3 -= (v3 | v3)
                v19 := []int{1, -425, 685, 249}
                _ = v18
                _ = v19
            } else {
                v2 += (-(v2))
                for fi20 := 0; fi20 < 15; fi20++ {
                    v21 := GenS1{F0: 2147483647, F1: -1, F2: v1}
                    _ = fi20
                    _ = v21
                }
                var v22 float64 = 3.14159
                _ = v22
            }
        }
        _ = v16
    }
    v23 := GenS4{F0: (v3 - (v3 ^ v3))}
    v2 += 42
    v24 := v4[0]
    switch -642 {
    case 0:
        v25 := GenS4{F0: (v3 / (((v3 % ((1000) | 1))) | 1))}
        _ = v25
    case 424:
        v24 = v24
    default:
        if (!false) {
            v1 = v24
        }
    }
    v24 = v24
    for fi26 := 0; fi26 < 3; fi26++ {
        fmt.Println(v3, v1)
        _ = fi26
    }
    v27 := []int{-1}
    var v28 int64 = -9223372036854775808
    var v29 int8 = v2
    v1 = v24
    var v30 uint32 = 700
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
