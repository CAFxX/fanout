package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 string; F1 uint8 }
func genFn0(p0 int8, p1 uint32) uint64 {
    for fi5 := 0; fi5 < 14; fi5++ {
        p0 = p0
        p1 = (p1 | 2147483648)
        var v6 uint64 = 543
        _ = fi5
        _ = v6
    }
    if (-164 >= 955) {
        if false {
            v7 := GenS1{F0: 18446744073709551615, F1: (p0 / ((p0) | 1))}
            _ = v7
        }
        fmt.Println(p1)
    } else {
        v8 := GenS2{F0: "a", F1: 255}
        fmt.Println(p0, p0)
        _ = v8
    }
    if (455 != 1) {
        if true {
            var v9 uint16 = 834
            _ = v9
        }
        v10 := []int{82, 32768}
        _ = v10
    }
    v11 := []int{42, 939, 707}
    _ = p0
    _ = p1
    _ = v11
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"foo bar"}
    v3 := make(map[int]string)
    var v4 int64 = 0
    switch 42 {
    case 4294967295:
        var v12 string = ""
        _ = v12
    case 128:
        v13 := GenS1{F0: 2, F1: 42}
        _ = v13
    case 32767:
        var v14 uint16 = 200
        _ = v14
    default:
        for fi15 := 0; fi15 < 7; fi15++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi15
        }
    }
    var v16 int8 = 127
    if len(v2) > 0 {
        v2[0] = "a"
    }
    fmt.Println(v4, v4)
    if false {
        v16 += v16
        var v17 float32 = -2.25
        v17 = float32((float64((0.0 - 0.1)) + 0.0))
        _ = v17
    }
    var v18 uint32 = 2147483647
    v19 := GenS1{F0: 4294967295, F1: 119}
    v20 := GenS2{F0: "\n\t", F1: 0}
    var v21 string = ("foo bar" + ("" + ""))
    if true {
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 int32 = 1
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 10; fi24++ {
            switch fi24 {
            case 9223372036854775807:
                v21 = (v21 + v21)
            case 336:
                v21 = "\n\t"
            }
            _ = fi24
        }
    }
    var v25 float64 = (3.14159 * (((-2.25 - float64(0.1)) + 1.5) - (0.1 + float64(1.5))))
    v21 = "foo bar"
    var v26 uint16 = 25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
