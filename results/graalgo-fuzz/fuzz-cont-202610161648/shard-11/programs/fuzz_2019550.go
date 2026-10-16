package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
func genFn0() int8 {
    if false {
        var v7 int8 = 127
        _ = v7
    }
    var v8 int32 = -1
    _ = v8
    return -1
}

func main() {
    var v1 uint16 = 2
    v2 := []uint32{}
    var v3 uint64 = 4294967295
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    if true {
        v1 += v1
    } else {
        var v9 string = "\n\t"
        v3 += (^(v3))
        _ = v9
    }
    switch ("" + "foo bar") {
    case "foo bar":
        var v10 int64 = 1
        _ = v10
    case "":
        switch 42 {
        case 1:
            fmt.Println(v3, v1)
        case 2147483647:
            v1 += v1
        case 32768:
            delete(v6, 414)
        default:
            var v11 uint8 = 127
            _ = v11
        }
    }
    switch "a" {
    case "":
        if false {
            v12 := GenS1{F0: 303, F1: 72, F2: -128}
            var v13 uint = 1
            v14 := GenS1{F0: 32768, F1: -1, F2: -38}
            _ = v12
            _ = v13
            _ = v14
        }
    case "foo bar":
        for fi15 := 0; fi15 < 17; fi15++ {
            v3 = (v3 ^ (v3 | v3))
            v6[-1] = ((("" + ("foo bar" + "a")) + "a") + (("\n\t" + "a") + "hello"))
            _ = fi15
        }
    }
    var v16 int64 = 42
    v1 += (v1 + (v1 / (((v1 % (((v1 >> ((0) & 63))) | 1))) | 1)))
    var v17 int32 = 2147483647
    fmt.Println(v3)
    if false {
        var v18 uint = 575
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v18
    } else {
        if false {
            switch 947 {
            case 0:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 42:
                v17 = (^(v17))
            default:
                fmt.Println(v16, v3, v16)
            }
        } else {
            if false {
                var v19 int64 = v16
                _ = v19
            }
        }
    }
    if (!((v1 + v1) > v1)) {
        v2 = append(v2, 42)
        v1 = v1
    } else {
        v17 = (v17 ^ v17)
        delete(v5, 185)
    }
    fmt.Println(v17, v1)
    v20 := GenS1{F0: 32768, F1: 45, F2: genFn0()}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
