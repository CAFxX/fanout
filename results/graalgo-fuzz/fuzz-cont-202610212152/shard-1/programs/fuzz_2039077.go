package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string }
func genFn0() bool {
    var v4 uint64 = 4294967295
    v4 = v4
    v5 := GenS1{F0: 1}
    var v6 float32 = (0.0 + 1.5)
    _ = v4
    _ = v5
    _ = v6
    return true
}
func genFn1() uint32 {
    var v7 int = 127
    var v8 int32 = -2147483648
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 344
}

func main() {
    v1 := []int{}
    v2 := []uint32{}
    var v3 uint = 42
    v9 := GenS2{F0: 644}
    switch 2147483647 {
    case 4294967295:
        if true {
            for fi10 := 0; fi10 < 19; fi10++ {
                v3 = 65535
                _ = fi10
            }
        }
    case 0:
        var v11 int32 = -2147483648
        _ = v11
    default:
        if true {
            fmt.Println(v3)
            v3 = (v3 / (((v3 / (((v3 * v3)) | 1))) | 1))
            v12 := GenS3{F0: 26}
            _ = v12
        }
    }
    switch "\n\t" {
    case "\n\t":
        v3 += v3
    case "foo bar":
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(v3, fi13)
            _ = fi13
        }
    default:
        v3 = (v3 + (v3 << ((442) & 63)))
    }
    v2 = append(v2, 2147483648)
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(fi14, fi14)
        v15 := GenS4{F0: "a"}
        _ = fi14
        _ = v15
    }
    v3 -= (v3 ^ v3)
    switch 4294967295 {
    case 42:
        v2 = append(v2, 0)
    default:
        v3 = (v3 + (v3 >> ((v3) & 63)))
    }
    v3 -= (^(v3))
    v3 += v3
    fmt.Println(v3)
    v3 += (v3 >> (((v3 % ((705) | 1))) & 63))
    if true {
        switch "hello" {
        case "hello":
            if false {
                var v16 uint64 = 713
                _ = v16
            }
        case "":
            var v17 string = "foo bar"
            _ = v17
        default:
            var v18 uint = v3
            _ = v18
        }
    }
    var v19 int8 = 127
    switch "hello" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 42
        }
    case "\n\t":
        v3 = v3
    }
    switch -9223372036854775808 {
    case 128:
        var v20 int32 = 324
        _ = v20
    case -2147483648:
        fmt.Println(v19, v19)
    default:
        v21 := GenS4{F0: "foo bar"}
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
