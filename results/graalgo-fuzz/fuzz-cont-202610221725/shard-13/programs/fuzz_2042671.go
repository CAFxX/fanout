package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint8 }
func genFn0() int32 {
    v4 := GenS1{F0: 2, F1: (0.1 <= 1e10)}
    var v5 int32 = -2147483648
    v6 := []int{-488, 32768, 0}
    var v7 int64 = 9223372036854775807
    if false {
        var v8 float32 = float32(((0.0 * 0.0) * 0.0))
        _ = v8
    } else {
        v9 := GenS1{F0: 42, F1: true}
        _ = v9
    }
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (v5 >> ((508) & 63))
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = -2.25
    var v3 bool = true
    v2 = (v2 * 0.1)
    v2 = v2
    var v10 bool = v3
    v3 = (v10 || false)
    if v3 {
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := []int{}
            _ = fi11
            _ = v12
        }
    } else {
        v13 := v1[42]
        v3 = v3
        _ = v13
    }
    switch 9223372036854775807 {
    case -2147483648:
        var v14 uint = 914
        _ = v14
    case -9223372036854775808:
        var v15 int = -1
        _ = v15
    default:
        for fi16 := 0; fi16 < 13; fi16++ {
            v2 = (1e10 / 1.5)
            var v17 uint64 = 847
            _ = fi16
            _ = v17
        }
    }
    v18 := []int{1, 32767, 128, 1}
    var v19 int8 = 115
    switch "hello" {
    case "\n\t":
        v20 := v1[-2147483648]
        _ = v20
    case "foo bar":
        v10 = false
    }
    if (genFn0() == -71) {
        var v21 int8 = 112
        _ = len(v1)
        _ = v21
    } else {
        v22 := GenS2{F0: 1}
        fmt.Println(v19, v19, v19)
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 string = "a"
            v25 := GenS1{F0: 4294967295, F1: v10}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v22
    }
    switch 9223372036854775807 {
    case 1:
        var v26 int8 = (v19 << ((2) & 63))
        _ = v26
    case -2147483648:
        var v27 int8 = (v19 / ((-38) | 1))
        _ = v27
    case 32767:
        var v28 int = -755
        _ = v28
    default:
        if true {
            v29 := v1[-2147483648]
            _ = len(v18)
            if true {
                v2 = v2
            } else {
                var v30 string = "foo bar"
                v31 := GenS2{F0: 1}
                _ = v30
                _ = v31
            }
            _ = v29
        } else {
            v3 = v10
            for fi32 := 0; fi32 < 1; fi32++ {
                v33 := GenS2{F0: 4294967295}
                v19 += (v19 + (v19 ^ v19))
                _ = fi32
                _ = v33
            }
        }
    }
    v19 -= (-(v19))
    var v34 int8 = v19
    var v35 bool = v3
    switch 0 {
    case 1:
        if (4294967295 > 2) {
            switch "a" {
            case "a":
                v36 := GenS2{F0: 4294967295}
                _ = v36
            case "":
                if ((!v10) || false) {
                    _ = len(v1)
                    var v37 uint8 = 53
                    v38 := GenS3{F0: 18446744073709551615, F1: v37, F2: v37}
                    _ = v37
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 2; fi39++ {
                        v19 += -15
                        _ = fi39
                    }
                    _ = len(v18)
                }
            }
        }
    case 2147483647:
        v40 := GenS2{F0: 0}
        _ = v40
    case -295:
        fmt.Println(v10)
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        var v42 int = (^(fi41))
        var v43 string = "hello"
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
}
