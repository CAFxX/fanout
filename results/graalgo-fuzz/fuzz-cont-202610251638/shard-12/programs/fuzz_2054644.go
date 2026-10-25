package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int32 {
    var v6 int32 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        if false {
            v6 = v6
            fmt.Println(v6, v6)
        }
        fmt.Println(fi7)
        _ = fi7
    }
    _ = v6
    return v6
}
func genFn1() uint {
    var v8 int32 = 0
    var v9 int64 = 0
    v9 -= (v9 >> ((42) & 63))
    v9 = v9
    _ = v8
    _ = v9
    return 72
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = (("hello" + "foo bar") + "a")
    var v3 int = 23
    var v4 uint32 = 2147483648
    v5 := []uint32{}
    var v10 int8 = 127
    v11 := GenS1{F0: 4294967295}
    var v12 int32 = -2147483648
    v13 := GenS1{F0: 280}
    var v14 int32 = (-(v12))
    var v15 uint64 = 18446744073709551615
    fmt.Println(v4, v10, v12)
    switch v3 {
    case 32768:
        fmt.Println(v12, v4, v15)
    case 768:
        var v16 int = v3
        _ = v16
    default:
        v17 := GenS1{F0: 88}
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        if true {
            fmt.Println(v3)
            _ = len(v1)
        }
        var v19 uint32 = 332
        if false {
            var v20 uint = 18446744073709551615
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    var v21 uint64 = v15
    delete(v1, (v2 + ((v2 + "a") + (("" + v2) + v2))))
    v21 += (^(v21))
    for fi22 := 0; fi22 < 9; fi22++ {
        if ((1 == 100) || ((1 > (v3 - 2147483647)) || false)) {
            var v23 uint8 = 99
            _ = v23
        } else {
            switch "foo bar" {
            case "\n\t":
                v24 := GenS1{F0: genFn1()}
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            if (1000 < 608) {
                v10 = (v10 / (((v10 >> ((0) & 63))) | 1))
                v26 := GenS1{F0: 0}
                _ = v26
            }
            if (false || ((0.1 == ((-2.25 / 1.0) - 0.0)) && false)) {
                switch "foo bar" {
                case "\n\t":
                    if true {
                        fmt.Println(v2, v4)
                        fmt.Println(v10, v4)
                    } else {
                        v27 := GenS1{F0: 277}
                        var v28 int32 = v12
                        _ = v27
                        _ = v28
                    }
                }
            } else {
                v12 = (^(v12))
                v10 = (-(v10))
            }
            _ = fi25
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 uint8 = 42
            v10 -= (v10 >> ((42) & 63))
            var v31 int8 = (v10 / ((127) | 1))
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi22
    }
    if (false || (918 < 42)) {
        v1[v2] = true
        if (int(v10) <= v3) {
            var v32 uint8 = uint8(v21)
            var v33 uint16 = 32767
            v5 = append(v5, (v4 + (v4 - v4)))
            _ = v32
            _ = v33
        }
    } else {
        var v34 int8 = v10
        fmt.Println(v2)
        _ = v34
    }
    v35 := v1[(v2 + v2)]
    if (!false) {
        v36 := GenS1{F0: 255}
        var v37 float32 = 3.14159
        _ = v36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
