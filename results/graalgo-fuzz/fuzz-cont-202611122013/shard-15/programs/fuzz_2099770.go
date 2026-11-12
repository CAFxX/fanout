package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
func genFn0(p0 uint) int8 {
    var v7 uint16 = 513
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(v7)
        switch "foo bar" {
        case "\n\t":
            v9 := []int{4294967295, -742}
            _ = v9
        default:
            if false {
                var v10 string = ("hello" + "foo bar")
                v10 = v10
                _ = v10
            }
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 int16) bool {
    v11 := GenS1{F0: 65535, F1: (-(p1))}
    p0 = (p1 | (p1 & -909))
    _ = p0
    _ = p1
    _ = v11
    return (1 > -1)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 42
    v3 := make(map[string]int)
    var v4 float32 = (1.5 + (0.0 * 0.0))
    v5 := []int{32768, -545, 128}
    v6 := make(map[string]int)
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int8 = -77
        fmt.Println(v13, fi12, fi12)
        _ = fi12
        _ = v13
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v4)
    for fi14 := 0; fi14 < 20; fi14++ {
        _ = len(v3)
        var v15 string = ("a" + ("\n\t" + "a"))
        var v16 uint = 2
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = len(v5)
    var v17 uint32 = 478
    v18 := GenS1{F0: 65535, F1: 0}
    if true {
        if true {
            var v19 int64 = v2
            switch ("a" + "") {
            case "a":
                v20 := GenS2{F0: 65535, F1: 127, F2: 1}
                _ = v20
            case "hello":
                var v21 uint = 0
                _ = v21
            case "":
                if false {
                    v22 := GenS1{F0: 825, F1: -378}
                    var v23 int32 = 0
                    _ = v22
                    _ = v23
                }
            default:
                var v24 uint64 = 0
                _ = v24
            }
            var v25 float32 = float32((1.5 + 1e10))
            _ = v19
            _ = v25
        } else {
            var v26 int16 = -32768
            v2 += 1
            _ = len(v5)
            _ = v26
        }
        var v27 uint16 = 2
        var v28 int8 = 100
        _ = v27
        _ = v28
    } else {
        var v29 int32 = -2147483648
        var v30 uint8 = 42
        fmt.Println(v17, v29)
        _ = v29
        _ = v30
    }
    fmt.Println(v2, v2, v17)
    v1[174] = false
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 uint = 4294967295
        var v33 float64 = 0.0
        if len(v5) > 0 {
            v5[0] = fi31
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 int8 = -104
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
