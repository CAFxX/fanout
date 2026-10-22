package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS4 struct { F0 int; F1 uint }
func genFn0(p0 uint32, p1 uint, p2 int8) uint64 {
    p2 = -128
    if (false || (0 < 42)) {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint = p1
            if true {
                if false {
                    var v9 int = (fi7 | -460)
                    _ = v9
                } else {
                    var v10 float32 = (0.0 * float32(3.14159))
                    v11 := GenS1{F0: -233}
                    if true {
                        fmt.Println(v8, fi7)
                        v12 := []int{(^(fi7)), fi7, fi7, (fi7 - (fi7 / ((fi7) | 1)))}
                        v12 = append(v12, fi7)
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 42}
                        var v14 uint8 = 121
                        switch fi7 {
                        case 32767:
                            var v15 uint16 = 65535
                            _ = v15
                        case 4294967295:
                            v14 = (v14 + 128)
                        case -340:
                            var v16 uint = v8
                            _ = v16
                        default:
                            v17 := []int{(fi7 - 1), fi7, int(p1)}
                            _ = v17
                        }
                        _ = v13
                        _ = v14
                    }
                    _ = v10
                    _ = v11
                }
            }
            _ = fi7
            _ = v8
        }
        v18 := GenS2{F0: "a"}
        p2 -= (p2 % ((p2) | 1))
        _ = v18
    }
    if (false || true) {
        var v19 uint8 = 97
        _ = v19
    } else {
        v20 := []int{668, 42}
        v21 := GenS2{F0: "\n\t"}
        var v22 bool = true
        _ = v20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        switch "" {
        case "hello":
            p0 = 42
        case "\n\t":
            v24 := []int{}
            _ = v24
        }
        var v25 float64 = 0.1
        _ = fi23
        _ = v25
    }
    v26 := []int{-402}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 2
}

func main() {
    var v1 float64 = 1.5
    var v2 int64 = 9223372036854775807
    v3 := []string{(("foo bar" + "\n\t") + "foo bar"), "hello", ("\n\t" + "hello")}
    v4 := make(map[int]uint32)
    v5 := []int{295, -749}
    var v6 int64 = (v2 + (v2 + (v2 % (((-(v2))) | 1))))
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 uint = 42
        var v29 uint8 = 9
        _ = fi27
        _ = v28
        _ = v29
    }
    v6 += (v6 & 0)
    v30 := GenS1{F0: -1}
    v6 += (^(v2))
    var v31 float32 = 0.1
    fmt.Println(v6, v1)
    var v32 int16 = 862
    for fi33 := 0; fi33 < 12; fi33++ {
        if (1 < 128) {
            var v34 string = (("" + "") + "\n\t")
            _ = v34
        } else {
            v35 := GenS3{F0: 1, F1: ((374 <= 2) || (false || false)), F2: 1}
            v5 = append(v5, -2147483648)
            _ = v35
        }
        var v36 int64 = (v6 ^ v2)
        fmt.Println(v1, fi33, v1)
        _ = fi33
        _ = v36
    }
    v37 := GenS1{F0: -32768}
    var v38 int64 = v6
    fmt.Println(v32, v2, v6)
    v39 := GenS4{F0: 127, F1: 1}
    _ = len(v3)
    switch 1 {
    case -733:
        for fi40 := 0; fi40 < 4; fi40++ {
            for fi41 := 0; fi41 < 2; fi41++ {
                v31 -= v31
                v42 := GenS3{F0: 1, F1: false, F2: 1}
                _ = fi41
                _ = v42
            }
            var v43 int = 127
            _ = fi40
            _ = v43
        }
    default:
        v44 := GenS2{F0: "a"}
        _ = v44
    }
    if true {
        fmt.Println(v1, v1)
        var v45 uint64 = 519
        var v46 uint64 = (v45 & (^(v45)))
        _ = v45
        _ = v46
    }
    if (!(true && ((true || true) && (0 < 425)))) {
        if false {
            fmt.Println(v6)
            fmt.Println(v1, v2)
        } else {
            v38 = (v6 >> ((255) & 63))
            fmt.Println(v38)
        }
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
