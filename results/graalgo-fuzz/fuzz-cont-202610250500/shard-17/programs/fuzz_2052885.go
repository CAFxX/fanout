package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
type GenS4 struct { F0 string; F1 uint64 }
func genFn0(p0 int) uint16 {
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(p0, p0)
        for fi7 := 0; fi7 < 16; fi7++ {
            p0 = 127
            _ = fi7
        }
        v8 := GenS1{F0: (2 == 32767), F1: -1}
        _ = fi6
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := []int{}
        for fi11 := 0; fi11 < 3; fi11++ {
            v10 = append(v10, 69)
            var v12 string = "\n\t"
            fmt.Println(fi9, v12, fi9)
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    return 376
}

func main() {
    v1 := []int{-2147483648}
    var v2 int = 572
    var v3 uint = 255
    v4 := make(map[uint32]string)
    v5 := []int{v2, 32768, (v2 - (-(v2))), v2}
    _ = len(v4)
    v3 = (v3 - (v3 * v3))
    var v13 int16 = -1
    v2 -= (^(v2))
    v14 := GenS1{F0: false, F1: -402}
    switch ("foo bar" + "") {
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        var v15 int = v2
        _ = v15
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        if (!true) {
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 int32 = 0
                fmt.Println(v3)
                v1 = append(v1, 166)
                _ = fi16
                _ = v17
            }
            v2 -= v2
            switch (("a" + "foo bar") + "\n\t") {
            case "foo bar":
                switch v2 {
                case 4294967295:
                    if true {
                        v3 += 2
                        if true {
                            v18 := GenS1{F0: false, F1: -1}
                            fmt.Println(v3)
                            v19 := GenS2{F0: true, F1: 1}
                            _ = v18
                            _ = v19
                        }
                    } else {
                        v20 := GenS3{F0: "foo bar", F1: 42, F2: 928}
                        for fi21 := 0; fi21 < 2; fi21++ {
                            var v22 uint64 = 2
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    }
                case 127:
                    var v23 uint = (v3 + v3)
                    _ = v23
                case 2147483647:
                    _ = len(v4)
                }
            case "\n\t":
                var v24 int = v2
                _ = v24
            default:
                v2 -= (v2 + v2)
            }
        }
    }
    var v25 int = -668
    v3 -= v3
    v26 := GenS4{F0: ("" + "a"), F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
