package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 int8, p1 bool, p2 uint16) int32 {
    if p1 {
        if (p0 >= 1) {
            v6 := []int{-466, -1}
            fmt.Println(p0, p0)
            _ = len(v6)
            _ = v6
        }
        p2 = p2
        switch ("" + "foo bar") {
        case "a":
            var v7 uint32 = 2147483648
            _ = v7
        case "foo bar":
            var v8 int16 = 987
            _ = v8
        default:
            var v9 int8 = (p0 << ((255) & 63))
            _ = v9
        }
    } else {
        p1 = (p2 == 65535)
        v10 := GenS1{F0: 303}
        _ = v10
    }
    var v11 uint32 = 265
    v12 := []int{593, 103, 127, -1}
    var v13 int16 = -1
    var v14 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 2147483647
}

func main() {
    var v1 string = "a"
    var v2 int64 = 9223372036854775807
    var v3 string = (v1 + v1)
    v4 := make(map[int]int)
    var v5 float64 = (3.14159 + 1e10)
    if true {
        v15 := []int{}
        switch -1 {
        case 32768:
            if (true || (false || ((-1 >= 42) && ("hello" >= (v3 + "a"))))) {
                var v16 uint32 = 528
                v17 := GenS2{F0: 1, F1: (v16 | 2147483647)}
                var v18 int = -845
                _ = v16
                _ = v17
                _ = v18
            } else {
                v1 = ("" + "foo bar")
                for fi19 := 0; fi19 < 7; fi19++ {
                    fmt.Println(v3)
                    var v20 uint8 = 2
                    v21 := v4[-2147483648]
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        case 9223372036854775807:
            switch (v1 + v3) {
            case "":
                var v22 uint64 = 4294967295
                _ = v22
            case "a":
                v23 := GenS1{F0: 1}
                _ = v23
            default:
                for fi24 := 0; fi24 < 15; fi24++ {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v5 += v5
                        _ = fi25
                    }
                    _ = fi24
                }
            }
        }
        fmt.Println(v2, v2, v2)
        _ = v15
    }
    v26 := []int{-705, 9223372036854775807}
    switch v1 {
    case "":
        var v27 uint16 = 14
        _ = v27
    case "foo bar":
        var v28 uint32 = 2
        _ = v28
    case "a":
        fmt.Println(v2)
    }
    v3 = "hello"
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 int8 = 6
        if true {
            v1 = "hello"
            var v31 uint = 2
            _ = v31
        }
        _ = fi29
        _ = v30
    }
    var v32 uint = 1
    v26 = append(v26, 32767)
    var v33 bool = true
    var v34 uint32 = 285
    switch -9223372036854775808 {
    case 717:
        v3 = v1
    case 143:
        if (!(620 < -32768)) {
            var v35 uint64 = 18446744073709551615
            _ = v35
        }
    default:
        v4[127] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
