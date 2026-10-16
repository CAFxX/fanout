package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
func genFn0(p0 uint16, p1 int) uint {
    v5 := GenS1{F0: 624, F1: 1}
    for fi6 := 0; fi6 < 19; fi6++ {
        if true {
            if (false && (18446744073709551615 < 0)) {
                for fi7 := 0; fi7 < 15; fi7++ {
                    var v8 uint8 = 2
                    _ = fi7
                    _ = v8
                }
                p1 -= (fi6 << ((980) & 63))
            } else {
                v9 := []int{(^(fi6))}
                _ = v9
            }
            var v10 int32 = 42
            v11 := []int{}
            _ = v10
            _ = v11
        }
        var v12 uint8 = 1
        v13 := GenS1{F0: fi6, F1: 4294967295}
        _ = fi6
        _ = v12
        _ = v13
    }
    switch 397 {
    case 127:
        var v14 float32 = 1e10
        _ = v14
    case -222:
        if false {
            p0 = (p0 - (p0 + (^(p0))))
            v15 := []int{(^(p1)), p1, (p1 + (p1 + p1)), p1}
            _ = len(v15)
            _ = v15
        } else {
            var v16 int8 = 127
            _ = v16
        }
    case 518:
        if true {
            var v17 uint8 = 2
            v17 += v17
            v18 := GenS1{F0: p1, F1: 1}
            _ = v17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    var v1 uint64 = 416
    v2 := []uint32{2147483648, 2147483648, 95, 121}
    v3 := make(map[string]uint32)
    var v4 float32 = ((float32(-2.25) * float32(3.14159)) + 1e10)
    if (!(!true)) {
        var v19 uint16 = 255
        v20 := GenS1{F0: -2147483648, F1: (v1 & 2)}
        _ = v19
        _ = v20
    }
    if len(v2) > 0 {
        v2[0] = 498
    }
    var v21 float32 = float32(0.1)
    var v22 bool = (2 != (v1 % ((18446744073709551615) | 1)))
    fmt.Println(v1, v21, v22)
    for fi23 := 0; fi23 < 11; fi23++ {
        v21 -= v4
        _ = fi23
    }
    var v24 uint32 = 453
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint64 = 18446744073709551615
        var v27 int8 = 56
        var v28 uint64 = 0
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        switch ("a" + "\n\t") {
        case "a":
            if v22 {
                for fi30 := 0; fi30 < 13; fi30++ {
                    fmt.Println(v24, v4, v4)
                    _ = fi30
                }
            } else {
                switch fi29 {
                case 1:
                    if (false || (42 != 1)) {
                        v1 = 42
                        _ = len(v2)
                    } else {
                        v3["hello"] = 2147483647
                        fmt.Println(v21, v22)
                    }
                default:
                    var v31 int32 = 0
                    _ = v31
                }
                var v32 int16 = -1
                _ = v32
            }
        default:
            var v33 bool = (!(true && (false || (612 != 0))))
            _ = v33
        }
        var v34 int = (fi29 | fi29)
        fmt.Println(v34, v21)
        _ = fi29
        _ = v34
    }
    var v35 uint16 = 255
    var v36 int32 = 2147483647
    switch 975 {
    case 0:
        _ = len(v3)
    case 4294967295:
        fmt.Println(v4, v22, v1)
    case -1:
        if v22 {
            v4 += v4
            fmt.Println(v36, v21)
        }
    default:
        v1 += 4294967295
    }
    fmt.Println(v21)
    if v22 {
        v2 = append(v2, (v24 >> ((42) & 63)))
        switch "foo bar" {
        case "foo bar":
            v37 := GenS1{F0: 0, F1: 42}
            _ = v37
        }
    } else {
        v38 := GenS1{F0: 127, F1: 1}
        _ = v38
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        v21 += v21
    }
    var v39 float32 = (v4 - 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
}
