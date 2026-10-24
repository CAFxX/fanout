package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
type GenS2 struct { F0 string; F1 uint8; F2 string }
type GenS3 struct { F0 int64; F1 int }
func genFn0() uint {
    var v6 string = ("a" + "hello")
    v7 := GenS1{F0: 42, F1: 1, F2: 128}
    fmt.Println(v6)
    v6 = ""
    _ = v6
    _ = v7
    return 778
}
func genFn1(p0 int, p1 uint16, p2 int) int8 {
    var v8 uint32 = 2147483647
    var v9 int16 = -32768
    var v10 float64 = -2.25
    for fi11 := 0; fi11 < 9; fi11++ {
        switch ("foo bar" + "foo bar") {
        case "a":
            v8 -= v8
        case "\n\t":
            v12 := []int{p0}
            _ = v12
        case "":
            fmt.Println(v8, p1)
        default:
            for fi13 := 0; fi13 < 11; fi13++ {
                fmt.Println(v9)
                var v14 uint32 = (v8 - (v8 - v8))
                _ = fi13
                _ = v14
            }
        }
        var v15 int16 = 0
        switch 1 {
        case -570:
            fmt.Println(v10)
        case 127:
            fmt.Println(v9, v8)
        default:
            fmt.Println(p1)
        }
        _ = fi11
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 100
}

func main() {
    var v1 uint = 1
    var v2 uint16 = 2
    v3 := make(map[string]bool)
    var v4 uint = (v1 - v1)
    v5 := make(map[uint32]uint32)
    switch -9223372036854775808 {
    case -114:
        v16 := []int{}
        _ = v16
    case 127:
        fmt.Println(v2, v2)
    default:
        v4 = v4
    }
    v1 = genFn0()
    v2 -= 0
    var v17 uint8 = 71
    v18 := GenS2{F0: ("hello" + "a"), F1: v17, F2: ("a" + ("a" + ""))}
    var v19 uint = (v4 * (v1 | v1))
    v1 -= v4
    var v20 bool = (-128 < 25)
    var v21 uint8 = (v17 * v17)
    if v20 {
        switch -7 {
        case 32767:
            for fi22 := 0; fi22 < 4; fi22++ {
                switch fi22 {
                case 1:
                    v23 := GenS2{F0: (("hello" + "\n\t") + ""), F1: (v17 ^ v21), F2: ("" + ("foo bar" + "a"))}
                    _ = v23
                default:
                    v24 := GenS2{F0: "foo bar", F1: v21, F2: ("" + ("" + ("a" + "")))}
                    _ = v24
                }
                var v25 int = (fi22 ^ -961)
                _ = fi22
                _ = v25
            }
        case 128:
            switch 128 {
            case 692:
                var v26 string = ("" + (("a" + ((("a" + "hello") + "") + "\n\t")) + ("hello" + (("hello" + (("a" + (("" + "") + "hello")) + "")) + ""))))
                _ = v26
            default:
                var v27 int = -62
                _ = v27
            }
        case 712:
            v28 := GenS1{F0: -771, F1: (v2 << ((genFn0()) & 63)), F2: 4294967295}
            _ = v28
        default:
            v29 := GenS3{F0: 798, F1: -1}
            _ = v29
        }
    } else {
        var v30 uint64 = 4294967295
        _ = v30
    }
    fmt.Println(v17, v4)
    v31 := []int{-1}
    fmt.Println(v2, v2)
    v5[2147483648] = 2147483648
    v21 = (v21 - 2)
    var v32 string = "\n\t"
    v33 := GenS3{F0: 9223372036854775807, F1: -822}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
