package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint }
func genFn0() int64 {
    var v6 string = "a"
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{fi7, 302, (fi7 + (fi7 ^ fi7))}
        _ = fi7
        _ = v8
    }
    v9 := []int{-544, -877, 375, 720}
    if (!(!true)) {
        var v10 int = 4294967295
        _ = v10
    } else {
        var v11 uint32 = 2147483647
        v6 = v6
        _ = v11
    }
    v12 := GenS1{F0: 513}
    _ = v12
    _ = v6
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 1
    v3 := []string{}
    var v4 int32 = 0
    v5 := make(map[string]bool)
    var v13 uint64 = 0
    fmt.Println(v13, v2, v4)
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            v2 += uint8(v13)
            _ = fi15
        }
        _ = fi14
    }
    switch "a" {
    case "":
        v13 += (v13 - (v13 << ((65535) & 63)))
    case "foo bar":
        var v16 int8 = 26
        _ = v16
    }
    if (true && false) {
        var v17 uint32 = 854
        _ = v17
    }
    v4 = (v4 >> ((215) & 63))
    v4 = v4
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS1{F0: 449}
        for fi20 := 0; fi20 < 19; fi20++ {
            switch fi20 {
            case -37:
                switch fi18 {
                case 693:
                    _ = len(v1)
                default:
                    v21 := GenS2{F0: ("hello" + (("\n\t" + "\n\t") + "")), F1: 931}
                    _ = v21
                }
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v4, v4, v4)
    var v22 int64 = 9223372036854775807
    switch "" {
    case "":
        v13 += (v13 ^ (v13 / (((v13 * 989)) | 1)))
    case "a":
        if true {
            switch -1 {
            case -250:
                v22 = (v22 * (v22 | (v22 / ((genFn0()) | 1))))
            }
            if false {
                fmt.Println(v2)
                v3 = append(v3, "")
                var v23 uint = 4294967295
                _ = v23
            } else {
                var v24 int64 = -9223372036854775808
                _ = v24
            }
        }
    }
    fmt.Println(v4, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
