package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0(p0 uint) string {
    v5 := []int{}
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(fi6)
        v7 := GenS1{F0: (1 == 42), F1: 210}
        _ = fi6
        _ = v7
    }
    switch ("hello" + "\n\t") {
    case "\n\t":
        p0 = p0
    case "foo bar":
        for fi8 := 0; fi8 < 13; fi8++ {
            p0 -= 0
            _ = fi8
        }
    case "hello":
        var v9 uint32 = 2147483647
        _ = v9
    default:
        var v10 uint64 = 969
        _ = v10
    }
    p0 += (p0 & (p0 >> ((p0) & 63)))
    var v11 string = "\n\t"
    _ = p0
    _ = v11
    _ = v5
    return v11
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    var v3 int8 = 127
    v4 := make(map[string]bool)
    var v12 uint32 = 42
    switch "a" {
    case "a":
        if (!true) {
            delete(v1, "foo bar")
        }
    case "\n\t":
        switch "hello" {
        case "a":
            v13 := GenS1{F0: (125 >= 42), F1: -9223372036854775808}
            _ = v13
        case "":
            v14 := GenS1{F0: true, F1: 1}
            _ = v14
        }
    default:
        var v15 int8 = 100
        _ = v15
    }
    v16 := GenS1{F0: false, F1: -9223372036854775808}
    v3 = (^(v3))
    v12 = 1
    if true {
        switch -454 {
        case -392:
            if false {
                v17 := v2[1]
                _ = v17
            } else {
                _ = len(v2)
            }
        case -2147483648:
            if true {
                switch "a" {
                case "foo bar":
                    v3 = v3
                case "":
                    v3 -= v3
                default:
                    v18 := GenS1{F0: (0 == 4294967295), F1: 1}
                    _ = v18
                }
            }
        case 2147483647:
            fmt.Println(v12, v12)
        default:
            var v19 int8 = (v3 - (v3 * v3))
            _ = v19
        }
        if (true || false) {
            var v20 int64 = 1
            if false {
                v3 = v3
            }
            if (32767 > uint16(v12)) {
                v1["\n\t"] = 4294967295
                v20 -= -709
            }
            _ = v20
        } else {
            var v21 int = -879
            v22 := v4["hello"]
            _ = v21
            _ = v22
        }
        switch ("foo bar" + "hello") {
        case "\n\t":
            var v23 float32 = 1.5
            _ = v23
        case "hello":
            _ = len(v2)
        case "":
            var v24 uint8 = 1
            _ = v24
        }
    }
    v25 := GenS1{F0: false, F1: 1}
    fmt.Println(v3, v12)
    v26 := []int{128, 777}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
