package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 bool, p1 uint8, p2 string) int {
    fmt.Println(p1)
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int8 = 127
        _ = fi8
        _ = v9
    }
    var v10 int8 = 127
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -629
}
func genFn1(p0 uint, p1 int8, p2 bool) int8 {
    if (true && (593 >= 4294967295)) {
        for fi11 := 0; fi11 < 18; fi11++ {
            if (p2 || (1 <= 256)) {
                var v12 float64 = (1e10 - 0.0)
                p2 = (96 < 65535)
                var v13 int = (fi11 * genFn0((((!(2 != 255)) && (53 < 2)) && p2), 53, "a"))
                _ = v12
                _ = v13
            } else {
                var v14 bool = false
                _ = v14
            }
            _ = fi11
        }
        p0 -= 459
        v15 := []int{-1, -215}
        _ = v15
    } else {
        if p2 {
            if false {
                switch "foo bar" {
                case "\n\t":
                    switch 32767 {
                    case 4294967295:
                        p0 = 4294967295
                    case 0:
                        var v16 uint64 = 18446744073709551615
                        _ = v16
                    default:
                        v17 := GenS1{F0: 1000, F1: ((("" + "a") + "\n\t") + "a"), F2: 62}
                        _ = v17
                    }
                case "":
                    var v18 uint = (p0 | (^(p0)))
                    _ = v18
                default:
                    p0 += (p0 + p0)
                }
            } else {
                var v19 uint16 = 682
                _ = v19
            }
            var v20 uint16 = 42
            fmt.Println(p1, p0, p2)
            _ = v20
        }
        if false {
            fmt.Println(p2)
            v21 := GenS1{F0: -32768, F1: "\n\t", F2: -1}
            _ = v21
        }
        var v22 uint8 = 42
        _ = v22
    }
    if p2 {
        var v23 int8 = p1
        p1 += (v23 ^ 123)
        if p2 {
            switch -218 {
            case -753:
                v24 := GenS1{F0: 32767, F1: (("foo bar" + "") + ""), F2: 9223372036854775807}
                _ = v24
            }
            if true {
                v25 := []int{-637, 127}
                var v26 int64 = 42
                for fi27 := 0; fi27 < 2; fi27++ {
                    switch "a" {
                    case "a":
                        v26 = int64(p1)
                    case "":
                        var v28 int = 9223372036854775807
                        _ = v28
                    default:
                        v29 := GenS1{F0: -671, F1: (("a" + "") + "\n\t"), F2: v26}
                        _ = v29
                    }
                    p0 += p0
                    _ = fi27
                }
                _ = v25
                _ = v26
            }
        }
        _ = v23
    }
    if (("" < "foo bar") || (!(41 < 128))) {
        var v30 int32 = 42
        p0 -= p0
        _ = v30
    }
    v31 := GenS1{F0: -349, F1: "", F2: -300}
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return 1
}

func main() {
    var v1 int16 = -32768
    var v2 float64 = 0.0
    v3 := make(map[int]int)
    v4 := []uint32{}
    var v5 int64 = 42
    v2 = v2
    v32 := GenS2{F0: 800}
    if (false || true) {
        var v33 uint8 = 128
        var v34 string = "\n\t"
        _ = v33
        _ = v34
    }
    var v35 uint = 18446744073709551615
    v36 := v3[32768]
    for fi37 := 0; fi37 < 8; fi37++ {
        v38 := GenS3{F0: (v36 + v36)}
        var v39 uint8 = 42
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 uint = 896
    var v41 int16 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
