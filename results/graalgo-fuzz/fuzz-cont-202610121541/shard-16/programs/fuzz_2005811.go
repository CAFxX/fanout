package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint8, p1 int16) int8 {
    v7 := []int{-1}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 int16 = (p1 | int16(p0))
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int = (fi9 ^ 616)
        fmt.Println(fi9, fi9)
        var v11 int16 = (p1 & p1)
        _ = fi9
        _ = v10
        _ = v11
    }
    v7 = append(v7, 32767)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 107
}
func genFn1() int32 {
    var v12 uint64 = 2
    switch -2147483648 {
    case -9223372036854775808:
        v13 := []int{4294967295, 1, 538, 4294967295}
        _ = v13
    case 42:
        var v14 int32 = 0
        _ = v14
    default:
        v12 = 642
    }
    v15 := []int{}
    switch -1 {
    case 2147483647:
        fmt.Println(v12)
    default:
        var v16 uint32 = 1
        _ = v16
    }
    _ = v12
    _ = v15
    return -1
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    v4 := []string{("\n\t" + "foo bar")}
    var v5 uint8 = 127
    var v6 uint32 = 22
    if (true && true) {
        if false {
            var v17 float64 = (0.1 + 0.1)
            v18 := GenS1{F0: (v6 / ((266) | 1)), F1: "\n\t"}
            _ = v17
            _ = v18
        } else {
            switch 42 {
            case -816:
                for fi19 := 0; fi19 < 19; fi19++ {
                    v3[v6] = ("foo bar" + v1)
                    _ = len(v4)
                    if (9223372036854775807 > -9223372036854775808) {
                        _ = len(v4)
                    }
                    _ = fi19
                }
            }
            var v20 uint8 = 255
            _ = v20
        }
        if true {
            var v21 uint16 = 65535
            _ = v21
        }
        v1 = v1
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(fi22, v1)
        _ = fi22
    }
    if false {
        switch "\n\t" {
        case "hello":
            var v23 bool = true
            _ = v23
        case "foo bar":
            v5 = v5
        case "a":
            switch v1 {
            case "hello":
                v1 = v1
            case "":
                v1 = "\n\t"
            case "foo bar":
                _ = len(v2)
            default:
                fmt.Println(v5, v5, v1)
            }
        default:
            v4 = append(v4, ((v1 + "foo bar") + v1))
        }
    } else {
        v2[("hello" + "foo bar")] = false
        switch v1 {
        case "\n\t":
            v4 = append(v4, v1)
        case "":
            var v24 string = v1
            _ = v24
        default:
            for fi25 := 0; fi25 < 15; fi25++ {
                if true {
                    v1 = (v1 + v1)
                }
                _ = fi25
            }
        }
    }
    v5 += 233
    delete(v2, "hello")
    v26 := GenS2{F0: -9223372036854775808, F1: 2}
    var v27 int16 = 32767
    if true {
        if (true || ((true && true) || true)) {
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 uint = 1
                _ = fi28
                _ = v29
            }
        }
        var v30 uint8 = (v5 | v5)
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
