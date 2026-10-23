package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    v3 := make(map[int]int)
    v2 = false
    var v4 uint16 = 42
    v5 := GenS1{F0: 2, F1: 902}
    var v6 float32 = (((0.0 + (-2.25 * ((0.0 * float32(((0.1 / (-2.25 / 1.0)) / (0.1 * 0.1)))) + 3.14159))) - (-2.25 / float32(-2.25))) + -2.25)
    switch "\n\t" {
    case "\n\t":
        var v7 string = ("foo bar" + ("" + "a"))
        _ = v7
    case "foo bar":
        v8 := GenS2{F0: 208, F1: 1, F2: 18446744073709551615}
        _ = v8
    case "":
        v9 := GenS2{F0: 54, F1: 1000, F2: 1}
        _ = v9
    default:
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(v6, v4, v6)
            _ = fi10
        }
    }
    v11 := GenS1{F0: 1, F1: 2147483647}
    for fi12 := 0; fi12 < 11; fi12++ {
        if true {
            v2 = (0 == 1)
        } else {
            delete(v1, 681)
            fmt.Println(v2, v4, v2)
        }
        v6 = (1e10 * 3.14159)
        var v13 uint32 = 2
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: 169, F1: 1, F2: 18446744073709551615}
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := []int{1, (^(fi15)), fi15}
        _ = fi15
        _ = v16
    }
    switch 0 {
    case 560:
        v17 := v1[55]
        _ = v17
    default:
        var v18 uint32 = 2147483648
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS1{F0: 127, F1: 65535}
        v21 := GenS3{F0: 65535}
        _ = fi19
        _ = v20
        _ = v21
    }
    switch ((("a" + ((("a" + ("\n\t" + "hello")) + "a") + ("foo bar" + ("" + "hello")))) + "hello") + "") {
    case "foo bar":
        v22 := GenS3{F0: 2}
        _ = v22
    }
    v4 += v4
    for fi23 := 0; fi23 < 4; fi23++ {
        if v2 {
            var v24 int64 = 0
            fmt.Println(v6, v6, v4)
            _ = v24
        }
        fmt.Println(v2)
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
