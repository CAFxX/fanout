package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int32; F1 int; F2 int64 }
func genFn0() int64 {
    v4 := GenS1{F0: -748}
    var v5 int = -1
    _ = v4
    _ = v5
    return 459
}
func genFn1() int64 {
    var v6 bool = false
    v6 = true
    var v7 uint64 = 4294967295
    v7 += v7
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 0
    var v2 float64 = 1.5
    var v3 uint32 = 2
    var v8 int32 = 2147483647
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint8 = 128
        var v11 bool = true
        var v12 uint32 = v3
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        if true {
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 float32 = 3.14159
                v2 = v2
                if (2147483647 <= (^(v8))) {
                    var v16 float64 = v2
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
        } else {
            v17 := []int{fi13, (-(fi13)), (fi13 & (^(fi13))), fi13}
            fmt.Println(fi13)
            v18 := GenS1{F0: v1}
            _ = v17
            _ = v18
        }
        switch "hello" {
        case "":
            var v19 uint16 = 0
            _ = v19
        case "foo bar":
            v8 += v8
        case "a":
            v3 = v3
        default:
            switch "\n\t" {
            case "\n\t":
                for fi20 := 0; fi20 < 13; fi20++ {
                    if ((v2 + 0.0) < v2) {
                        v21 := []int{(fi13 & fi20), (fi13 / ((-2147483648) | 1)), fi13, (fi13 << ((255) & 63))}
                        if true {
                            fmt.Println(v3)
                            fmt.Println(fi13, fi13)
                        } else {
                            var v22 float64 = (v2 / -2.25)
                            var v23 uint8 = 0
                            if len(v21) > 0 {
                                v21[0] = (fi20 - -769)
                            }
                            _ = v22
                            _ = v23
                        }
                        v24 := GenS2{F0: 42, F1: fi20}
                        _ = v21
                        _ = v24
                    } else {
                        var v25 float64 = -2.25
                        _ = v25
                    }
                    var v26 uint = 2
                    _ = fi20
                    _ = v26
                }
            case "foo bar":
                var v27 uint32 = (v3 - 2147483647)
                _ = v27
            case "":
                v28 := GenS1{F0: v1}
                _ = v28
            }
        }
        if true {
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 float32 = 3.14159
                var v31 uint8 = 255
                _ = fi29
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 int = fi32
                v1 = (v1 << ((255) & 63))
                _ = fi32
                _ = v33
            }
        } else {
            for fi34 := 0; fi34 < 5; fi34++ {
                v8 += v8
                var v35 float32 = (((-2.25 * 3.14159) - 1e10) / 0.1)
                v36 := []int{fi34, 246}
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        _ = fi13
    }
    fmt.Println(v2, v8)
    v37 := []int{}
    for fi38 := 0; fi38 < 5; fi38++ {
        for fi39 := 0; fi39 < 5; fi39++ {
            v1 += v1
            _ = fi39
        }
        var v40 float64 = 0.1
        for fi41 := 0; fi41 < 9; fi41++ {
            if len(v37) > 0 {
                v37[0] = (-(fi38))
            }
            _ = fi41
        }
        _ = fi38
        _ = v40
    }
    switch "\n\t" {
    case "hello":
        v2 = v2
    case "\n\t":
        v42 := GenS3{F0: 2147483647, F1: 1, F2: 232}
        _ = v42
    case "a":
        if len(v37) >= 2 {
            v37 = v37[1:len(v37)]
        }
    default:
        fmt.Println(v2)
    }
    if len(v37) > 0 {
        v37[0] = 127
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v8
}
