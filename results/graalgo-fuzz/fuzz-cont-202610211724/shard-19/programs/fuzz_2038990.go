package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 bool }
func genFn0(p0 uint, p1 uint8, p2 uint) int32 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint8 = 128
        fmt.Println(p0, v7, p2)
        switch fi6 {
        case 1:
            var v8 int16 = 1
            _ = v8
        case -198:
            v9 := []int{-738, fi6, fi6}
            _ = v9
        default:
            p0 -= p2
        }
        _ = fi6
        _ = v7
    }
    switch "hello" {
    case "hello":
        v10 := GenS1{F0: 949, F1: -435}
        _ = v10
    case "":
        fmt.Println(p2)
    case "\n\t":
        var v11 uint = 1
        _ = v11
    }
    v12 := GenS2{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -2147483648
}

func main() {
    var v1 int16 = 425
    v2 := make(map[uint32]string)
    var v3 int32 = 2147483647
    v4 := make(map[string]int)
    var v5 int64 = 933
    v13 := GenS2{F0: 649}
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    fmt.Println(v1, v5)
    switch "foo bar" {
    case "a":
        v16 := GenS3{F0: "hello", F1: (!(v5 == (v5 * (v5 | v5))))}
        _ = v16
    case "":
        var v17 int8 = -57
        _ = v17
    case "hello":
        var v18 int32 = v3
        _ = v18
    }
    var v19 float32 = (1e10 * -2.25)
    var v20 float64 = 0.1
    v2[42] = ((("a" + "hello") + (("hello" + "a") + (("hello" + (("" + "a") + "\n\t")) + ("" + "")))) + "\n\t")
    var v21 uint32 = 42
    var v22 uint = 18446744073709551615
    for fi23 := 0; fi23 < 15; fi23++ {
        v21 = v21
        fmt.Println(v22, v3)
        _ = fi23
    }
    var v24 int64 = (v5 + v5)
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS3{F0: ("\n\t" + "\n\t"), F1: (fi25 <= (^(fi25)))}
        fmt.Println(v20, v20, v24)
        v20 = -2.25
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
