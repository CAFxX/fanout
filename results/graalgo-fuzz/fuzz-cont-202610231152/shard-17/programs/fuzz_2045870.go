package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint16, p1 int8, p2 int16) int64 {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6, p2)
        switch fi6 {
        case -9223372036854775808:
            if (true && true) {
                if true {
                    if true {
                        for fi7 := 0; fi7 < 13; fi7++ {
                            var v8 string = "foo bar"
                            _ = fi7
                            _ = v8
                        }
                        v9 := GenS1{F0: 501, F1: -540, F2: 1}
                        var v10 int32 = 2147483647
                        _ = v10
                        _ = v9
                    } else {
                        v11 := []int{}
                        v12 := GenS1{F0: -2147483648, F1: -66, F2: p2}
                        _ = v11
                        _ = v12
                    }
                } else {
                    p1 += p1
                }
                p1 = (-(p1))
            }
        }
        p1 -= 87
        _ = fi6
    }
    if false {
        var v13 int32 = 1
        fmt.Println(p1, p2, v13)
        v14 := GenS2{F0: ("" == "a"), F1: 0}
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}
func genFn1(p0 uint16, p1 uint) int64 {
    fmt.Println(p1)
    v15 := GenS1{F0: 2147483647, F1: -9223372036854775808, F2: 0}
    var v16 int64 = 42
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return v16
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 int64 = 867
    v5 := []int{4294967295, -9223372036854775808}
    fmt.Println(v4, v1)
    v1 -= 1.5
    var v17 bool = true
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v18 := GenS1{F0: 1, F1: 42, F2: 1000}
        _ = v18
    case "":
        var v19 bool = v17
        _ = v19
    case "a":
        var v20 uint32 = 4294967295
        _ = v20
    }
    var v21 uint32 = 2
    v5 = append(v5, -686)
    var v22 uint8 = 128
    v23 := v2[0]
    var v24 int32 = -1
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 int16 = -32768
        _ = fi25
        _ = v26
    }
    if v17 {
        for fi27 := 0; fi27 < 9; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 float32 = (v1 / float32(-2.25))
                _ = fi28
                _ = v29
            }
            var v30 float64 = 1e10
            _ = fi27
            _ = v30
        }
    }
    v1 -= (v1 + float32(1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
