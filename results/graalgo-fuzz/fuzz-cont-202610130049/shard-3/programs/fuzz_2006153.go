package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64, p1 uint64) int32 {
    if true {
        p0 -= p0
        if ("a" == "hello") {
            if false {
                fmt.Println(p1, p1)
                var v4 uint8 = 0
                fmt.Println(v4)
                _ = v4
            }
        } else {
            var v5 uint16 = 42
            if false {
                var v6 uint64 = p1
                v6 -= (p1 & 1)
                _ = v6
            }
            _ = v5
        }
        fmt.Println(p1, p0)
    } else {
        p0 += p0
        switch 32768 {
        case -2147483648:
            var v7 uint = 18446744073709551615
            _ = v7
        default:
            var v8 uint = 419
            _ = v8
        }
        v9 := GenS1{F0: 2147483647, F1: -2147483648, F2: 367}
        _ = v9
    }
    v10 := GenS2{F0: p0}
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 uint = 255
        p0 += p0
        switch fi11 {
        case 1:
            fmt.Println(fi11, p0, p1)
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return 1
}
func genFn1(p0 uint16, p1 uint8, p2 int32) int32 {
    p1 = (p1 >> ((255) & 63))
    var v13 uint64 = 262
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (p2 + p2)
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 uint32 = 2147483648
    var v14 float64 = (-2.25 / 1e10)
    v14 -= -2.25
    v15 := []int{97, 32768, -1}
    v3 += (v3 >> ((0) & 63))
    fmt.Println(v3, v14)
    var v16 string = ""
    v16 = (v16 + v16)
    v17 := GenS1{F0: 0, F1: 1, F2: 2}
    var v18 int16 = 1000
    var v19 int64 = 1
    switch -754 {
    case 117:
        v14 += (v14 - 1e10)
    default:
        v3 = v3
    }
    v20 := GenS2{F0: v19}
    var v21 string = (v16 + (v16 + "\n\t"))
    if true {
        switch 2147483647 {
        case 648:
            v3 = (v3 | v3)
        default:
            fmt.Println(v21)
        }
        var v22 uint16 = 65535
        switch ("a" + "a") {
        case "foo bar":
            _ = len(v1)
        case "hello":
            if true {
                v18 = v18
            } else {
                for fi23 := 0; fi23 < 10; fi23++ {
                    fmt.Println(fi23, v21, v16)
                    _ = fi23
                }
            }
        case "":
            v21 = v16
        default:
            var v24 int = 32768
            _ = v24
        }
        _ = v22
    }
    v16 = ("hello" + v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
