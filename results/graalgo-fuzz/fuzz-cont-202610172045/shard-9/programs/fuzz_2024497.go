package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 1; fi5++ {
        switch fi5 {
        case 2147483647:
            v6 := []int{fi5, fi5, fi5, 32767}
            _ = v6
        default:
            var v7 string = "a"
            _ = v7
        }
        var v8 uint = 772
        _ = fi5
        _ = v8
    }
    v9 := GenS1{F0: 32767, F1: 42}
    if false {
        var v10 bool = false
        _ = v10
    } else {
        var v11 uint64 = 42
        var v12 string = "a"
        _ = v11
        _ = v12
    }
    v13 := []int{-9223372036854775808, -720, 414}
    var v14 int64 = 0
    _ = v13
    _ = v14
    _ = v9
    return -933
}
func genFn1() int {
    var v15 uint8 = 229
    v16 := GenS1{F0: 32767, F1: 2}
    var v17 bool = true
    _ = v15
    _ = v16
    _ = v17
    return -9223372036854775808
}

func main() {
    v1 := []string{"\n\t", "\n\t", ""}
    var v2 uint = 4
    var v3 uint8 = 1
    v4 := make(map[int]uint32)
    var v18 float64 = 1e10
    var v19 int32 = 351
    v3 += 42
    v2 += (v2 - v2)
    for fi20 := 0; fi20 < 8; fi20++ {
        if len(v1) > 0 {
            v1[0] = (("a" + ("foo bar" + ("foo bar" + "foo bar"))) + ("hello" + "foo bar"))
        }
        _ = fi20
    }
    v19 = 1
    var v21 uint64 = 2
    v3 -= 128
    v19 += (v19 % (((v19 * v19)) | 1))
    fmt.Println(v3, v21, v19)
    if (42 > genFn1()) {
        fmt.Println(v18)
        var v22 int16 = 16
        _ = v22
    } else {
        var v23 uint = (v2 ^ v2)
        if (!((v21 | (^(v21))) != (^(v21)))) {
            for fi24 := 0; fi24 < 18; fi24++ {
                v2 -= (v2 * v2)
                _ = fi24
            }
            v23 = (v2 - 1)
        }
        _ = v23
    }
    fmt.Println(v21)
    v25 := GenS1{F0: 0, F1: 65535}
    switch (("a" + "\n\t") + ("a" + ("\n\t" + "\n\t"))) {
    case "hello":
        if (409 != 255) {
            fmt.Println(v21, v2)
            for fi26 := 0; fi26 < 2; fi26++ {
                delete(v4, fi26)
                _ = fi26
            }
        }
    }
    if false {
        v27 := GenS2{F0: int64(v2), F1: 65535, F2: genFn0()}
        fmt.Println(v19)
        _ = v27
    } else {
        for fi28 := 0; fi28 < 2; fi28++ {
            v21 += v21
            _ = fi28
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
}
