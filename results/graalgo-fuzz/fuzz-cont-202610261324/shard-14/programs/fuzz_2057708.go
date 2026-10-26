package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
func genFn0(p0 uint64, p1 uint16) int {
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := []int{(fi4 | (fi4 % (((fi4 ^ (fi4 | fi4))) | 1))), fi4, (fi4 ^ 32768)}
        _ = fi4
        _ = v5
    }
    var v6 int16 = 32767
    _ = p0
    _ = p1
    _ = v6
    return -441
}

func main() {
    var v1 uint32 = 65535
    var v2 string = (("a" + "hello") + "\n\t")
    v3 := []string{v2, v2, (v2 + v2), v2}
    if len(v3) > 0 {
        v3[0] = v2
    }
    var v7 float32 = (1.5 / 0.1)
    switch 4294967295 {
    case -9223372036854775808:
        switch v2 {
        case "a":
            v1 = (^(v1))
        case "":
            switch 32768 {
            case -9223372036854775808:
                var v8 uint32 = 695
                _ = v8
            case 1:
                var v9 uint16 = 255
                _ = v9
            }
        case "\n\t":
            if true {
                v1 = (^(v1))
            } else {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v1 -= (^(v1))
                    fmt.Println(v7, v1, fi10)
                    var v11 string = (v2 + "")
                    _ = fi10
                    _ = v11
                }
                _ = len(v3)
                _ = len(v3)
            }
        }
    }
    var v12 uint = 203
    if (((18446744073709551615 >= 424) || (true && false)) && (!false)) {
        var v13 uint8 = 0
        _ = v13
    } else {
        if false {
            var v14 int32 = 1
            _ = v14
        } else {
            var v15 int64 = 0
            fmt.Println(v1, v2, v2)
            v1 -= v1
            _ = v15
        }
        var v16 float32 = -2.25
        _ = v16
    }
    if true {
        if (("a" + v2) < (v2 + v2)) {
            var v17 float32 = v7
            _ = v17
        }
    }
    var v18 string = v2
    v19 := GenS1{F0: 32768, F1: 2}
    fmt.Println(v18)
    fmt.Println(v2, v1, v1)
    switch 32767 {
    case 127:
        var v20 int64 = 42
        _ = v20
    case 32767:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi21 := 0; fi21 < 9; fi21++ {
            v7 -= v7
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 int32 = 0
                _ = fi22
                _ = v23
            }
            var v24 int32 = 2147483647
            _ = fi21
            _ = v24
        }
    }
    v25 := GenS1{F0: -2147483648, F1: 255}
    if len(v3) > 0 {
        v3[0] = "a"
    }
    if true {
        v26 := GenS1{F0: -861, F1: 259}
        var v27 uint = 0
        _ = v26
        _ = v27
    }
    var v28 uint64 = 4294967295
    if len(v3) > 0 {
        v3[0] = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v7
}
