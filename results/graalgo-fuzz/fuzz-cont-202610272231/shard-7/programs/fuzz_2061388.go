package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 string; F1 int32 }
func genFn0(p0 int, p1 int16) uint64 {
    var v6 int32 = -1
    var v7 int32 = v6
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 bool) uint {
    var v8 uint16 = 32767
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: (^(v8)), F1: 0}
        var v11 bool = p0
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v8)
    for fi12 := 0; fi12 < 5; fi12++ {
        p0 = false
        _ = fi12
    }
    switch ("hello" + "a") {
    case "hello":
        if (!(-2.25 >= 0.1)) {
            switch 525 {
            case 254:
                var v13 float32 = float32((float64((1e10 + (1.5 + ((0.0 - ((0.0 * (float32((((((0.0 / -2.25) * 0.0) * (0.0 / 1e10)) * 0.0) / 1.0)) - 0.1)) / 0.1)) - 1e10)))) * 1.5))
                _ = v13
            case 528:
                v8 += (v8 % ((v8) | 1))
            }
            v8 = v8
        }
    default:
        var v14 uint64 = 1
        _ = v14
    }
    _ = p0
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{42}
    var v3 float64 = ((1e10 / -2.25) / -2.25)
    var v4 string = "\n\t"
    var v5 uint8 = 1
    for fi15 := 0; fi15 < 7; fi15++ {
        v5 = v5
        if false {
            v3 += v3
            fmt.Println(v4, v3, v4)
        } else {
            for fi16 := 0; fi16 < 17; fi16++ {
                v3 = 1.5
                _ = fi16
            }
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 bool = true
                v4 = v4
                v19 := GenS2{F0: "a", F1: 42}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = fi15
    }
    if true {
        var v20 bool = false
        fmt.Println(v5, v20)
        _ = v20
    } else {
        v2 = append(v2, 2147483647)
    }
    delete(v1, 831)
    if false {
        fmt.Println(v3)
    } else {
        var v21 int16 = -1
        if true {
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 uint64 = 86
                _ = fi22
                _ = v23
            }
            v24 := GenS2{F0: (v4 + "a"), F1: 0}
            fmt.Println(v5, v5)
            _ = v24
        }
        _ = v21
    }
    fmt.Println(v4, v5)
    switch -2147483648 {
    case 0:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -546:
        var v25 float64 = v3
        _ = v25
    }
    if false {
        v26 := GenS2{F0: ((v4 + v4) + v4), F1: -444}
        var v27 uint8 = v5
        for fi28 := 0; fi28 < 1; fi28++ {
            v4 = ("" + v4)
            _ = fi28
        }
        _ = v26
        _ = v27
    }
    fmt.Println(v4, v5, v5)
    for fi29 := 0; fi29 < 14; fi29++ {
        v30 := GenS2{F0: v4, F1: 0}
        v5 += v5
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 int8 = int8(fi31)
        _ = fi31
        _ = v32
    }
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = 32767
    }
    if true {
        fmt.Println(v5, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
