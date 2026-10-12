package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int32 }
func genFn0(p0 uint8, p1 bool) int16 {
    switch 128 {
    case 9223372036854775807:
        for fi6 := 0; fi6 < 14; fi6++ {
            switch ("foo bar" + "hello") {
            case "hello":
                p1 = p1
            case "a":
                for fi7 := 0; fi7 < 3; fi7++ {
                    v8 := GenS1{F0: fi6, F1: 2}
                    fmt.Println(fi7)
                    p0 -= (p0 | 221)
                    _ = fi7
                    _ = v8
                }
            }
            _ = fi6
        }
    case -2147483648:
        v9 := []int{9223372036854775807, 2147483647, 0, 42}
        _ = v9
    case 325:
        v10 := GenS2{F0: 1, F1: 32, F2: 1}
        _ = v10
    default:
        switch ("" + ("hello" + ("hello" + ("\n\t" + ("hello" + "\n\t"))))) {
        case "foo bar":
            var v11 uint32 = 4294967295
            _ = v11
        }
    }
    var v12 uint32 = 456
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            switch "\n\t" {
            case "\n\t":
                fmt.Println(p1, p1, fi13)
            case "hello":
                v15 := []int{fi13, fi14}
                _ = v15
            case "a":
                var v16 uint16 = 1
                _ = v16
            default:
                v17 := GenS2{F0: 0, F1: 255, F2: 0}
                _ = v17
            }
            _ = fi14
        }
        v18 := []int{fi13, 42}
        _ = fi13
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v12
    return 1000
}

func main() {
    v1 := []uint32{2147483647, 636, 2, 65535}
    v2 := []string{"\n\t", "", ((("a" + "hello") + "foo bar") + "\n\t")}
    var v3 bool = (((-2.25 / 1e10) + 1.5) == -2.25)
    var v4 int8 = 73
    var v5 bool = v3
    if v3 {
        var v19 int16 = 0
        var v20 string = ""
        _ = v19
        _ = v20
    } else {
        v4 = 127
        var v21 uint8 = 42
        _ = v21
    }
    v5 = ((genFn0(1, (-1 >= 911)) < 1000) || (-226 >= 683))
    _ = len(v1)
    v5 = (-1 < -2147483648)
    switch -583 {
    case 0:
        v4 = -1
    case 1:
        var v22 uint32 = 65535
        _ = v22
    }
    v3 = v3
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS2{F0: 65535, F1: 2, F2: -2147483648}
        v4 = v4
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: 6, F1: 23, F2: 0}
    fmt.Println(v4)
    fmt.Println(v5)
    var v26 uint64 = 1
    v4 += -23
    var v27 uint = 225
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
