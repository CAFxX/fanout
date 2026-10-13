package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint16, p1 uint) int64 {
    v6 := []int{}
    p0 += p0
    _ = p0
    _ = p1
    _ = v6
    return 543
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 0
    v3 := make(map[string]int)
    v4 := make(map[int]int)
    v5 := []uint32{0, 4294967295, 75, 884}
    v2 -= (^(v2))
    switch 307 {
    case 1:
        var v7 int64 = -9223372036854775808
        _ = v7
    }
    if true {
        switch 4294967295 {
        case 128:
            v2 = v2
        }
    }
    v8 := GenS1{F0: v2, F1: 32767}
    var v9 uint64 = 4294967295
    var v10 float64 = 1.5
    v11 := GenS2{F0: false}
    if true {
        var v12 float64 = v10
        v9 = 42
        switch 42 {
        case -821:
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
        case 32767:
            switch -455 {
            case 32767:
                for fi13 := 0; fi13 < 19; fi13++ {
                    v5 = append(v5, 465)
                    var v14 float32 = 0.1
                    if (!(709 == 9223372036854775807)) {
                        v15 := GenS2{F0: (v14 != v14)}
                        if false {
                            v9 = 8
                            v2 = 255
                            var v16 float64 = v12
                            _ = v16
                        }
                        _ = v15
                    } else {
                        var v17 uint32 = 856
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                }
            case 128:
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 float64 = ((v12 / 1e10) - v10)
                    v20 := GenS1{F0: v2, F1: 42}
                    v19 = (-2.25 + v10)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case 307:
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 string = ((("\n\t" + "hello") + ((((("\n\t" + "foo bar") + "") + "\n\t") + "") + "\n\t")) + ("" + ""))
                    _ = fi21
                    _ = v22
                }
            default:
                v2 -= (v2 + 127)
            }
        case 128:
            switch 1 {
            case 32767:
                var v23 float32 = 1e10
                _ = v23
            case 9223372036854775807:
                switch "\n\t" {
                case "":
                    v9 = 933
                }
            }
        }
        _ = v12
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v24 int32 = 1
        _ = v24
    }
    _ = len(v3)
    v10 -= ((0.0 / 0.1) - v10)
    if true {
        v9 -= 782
        delete(v3, "")
    } else {
        v25 := GenS2{F0: (false && (false && false))}
        _ = v25
    }
    var v26 int = 360
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 uint16 = 1
        v10 -= ((v10 * float64(0.0)) / v10)
        v2 = v2
        _ = fi27
        _ = v28
    }
    var v29 uint32 = 0
    v2 = (v2 | (v2 | (v2 << ((1) & 63))))
    fmt.Println(v2, v2, v2)
    fmt.Println(v2, v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
