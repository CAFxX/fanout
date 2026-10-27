package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool; F1 int; F2 string }
func genFn0() int8 {
    var v7 uint = 938
    v8 := []int{}
    if (2 != 1) {
        fmt.Println(v7)
    } else {
        _ = len(v8)
        v7 = 650
        var v9 int16 = 0
        _ = v9
    }
    _ = v7
    _ = v8
    return 42
}
func genFn1() int32 {
    v10 := []int{32767, -241, 286, 655}
    var v11 float32 = (1.5 * (0.0 / float32(0.1)))
    var v12 uint32 = 0
    _ = v10
    _ = v11
    _ = v12
    return 234
}

func main() {
    var v1 uint16 = 0
    var v2 int8 = -128
    v3 := make(map[uint32]string)
    v4 := []uint32{111}
    v5 := make(map[int]int)
    var v6 uint8 = 63
    v2 -= v2
    v3[4294967295] = ""
    for fi13 := 0; fi13 < 14; fi13++ {
        v6 = (v6 >> ((18446744073709551615) & 63))
        switch "a" {
        case "a":
            v14 := GenS1{F0: 1, F1: (v6 | (v6 << ((988) & 63)))}
            _ = v14
        default:
            v2 -= (v2 - (v2 & v2))
        }
        _ = fi13
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 bool = (!false)
        var v17 int16 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    v5[42] = 2147483647
    v6 += (^(v6))
    for fi18 := 0; fi18 < 7; fi18++ {
        v2 = 127
        v4 = append(v4, 0)
        _ = fi18
    }
    var v19 bool = (!(32767 < 32767))
    var v20 uint64 = 4294967295
    if ((^(v1)) <= v1) {
        v21 := GenS2{F0: ((!v19) || (false || v19)), F1: 4294967295, F2: ("\n\t" + ((("" + "") + "") + ("a" + ("" + ((("hello" + ("foo bar" + "a")) + ("foo bar" + "hello")) + (("hello" + "hello") + (("foo bar" + "\n\t") + "foo bar")))))))}
        if false {
            var v22 int = 4294967295
            fmt.Println(v2)
            v22 -= 32767
            _ = v22
        }
        v1 = (v1 / (((v1 - 42)) | 1))
        _ = v21
    } else {
        v20 -= (v20 % (((v20 * v20)) | 1))
        fmt.Println(v1)
    }
    v23 := GenS1{F0: 1, F1: (v6 / ((255) | 1))}
    v6 += (v6 >> ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
