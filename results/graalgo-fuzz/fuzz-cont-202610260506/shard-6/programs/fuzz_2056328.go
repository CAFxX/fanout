package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
func genFn0() uint16 {
    switch -348 {
    case 106:
        var v4 uint16 = 0
        _ = v4
    case -9223372036854775808:
        var v5 uint32 = 0
        _ = v5
    case -2147483648:
        var v6 uint = 881
        _ = v6
    }
    var v7 uint = 255
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 uint16 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    v3 -= 0
    switch -2147483648 {
    case 128:
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 bool = false
            if true {
                v11 := GenS1{F0: "foo bar", F1: 1}
                _ = v11
            } else {
                var v12 int = fi9
                v12 = 32767
                _ = v12
            }
            v3 = (v3 | genFn0())
            _ = fi9
            _ = v10
        }
    case 0:
        switch "\n\t" {
        case "hello":
            fmt.Println(v3)
        case "":
            fmt.Println(v3)
        case "\n\t":
            if (v3 > (v3 >> ((65535) & 63))) {
                fmt.Println(v3)
            } else {
                var v13 uint64 = 0
                v14 := []int{2147483647, 370, 2147483647, 371}
                if (v3 >= v3) {
                    if false {
                        v2[("hello" + (((("foo bar" + ("" + "hello")) + "a") + "foo bar") + (("a" + "\n\t") + "foo bar")))] = (!((9223372036854775807 < 42) || (951 == 661)))
                    } else {
                        v15 := GenS1{F0: "foo bar", F1: 42}
                        _ = v15
                    }
                }
                _ = v13
                _ = v14
            }
        }
    case 127:
        fmt.Println(v3)
    default:
        for fi16 := 0; fi16 < 4; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 uint = 4294967295
                v3 += 32767
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    }
    v3 = (v3 >> ((0) & 63))
    fmt.Println(v3)
    var v19 int32 = 0
    switch 127 {
    case 4294967295:
        fmt.Println(v19, v19)
    case 207:
        v20 := []int{9223372036854775807, -1}
        _ = v20
    }
    var v21 bool = (-59 <= 1)
    var v22 int16 = 976
    v3 += v3
    v22 = (v22 >> ((4294967295) & 63))
    v23 := []int{}
    v3 = v3
    var v24 int16 = (v22 >> ((0) & 63))
    var v25 uint16 = (v3 & (^(v3)))
    var v26 int = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
