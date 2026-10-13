package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int32 }
func genFn0(p0 int16, p1 int16) int {
    var v5 uint16 = 507
    v6 := []int{4294967295}
    v6 = append(v6, 128)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 183
}

func main() {
    var v1 uint = 42
    var v2 uint8 = 1
    var v3 float64 = float64((-2.25 - -2.25))
    v4 := make(map[string]string)
    if true {
        switch -402 {
        case 32768:
            switch 441 {
            case -399:
                for fi7 := 0; fi7 < 1; fi7++ {
                    v8 := GenS1{F0: 42, F1: -104, F2: (fi7 * (fi7 + -535))}
                    _ = fi7
                    _ = v8
                }
            }
        case 4294967295:
            var v9 float64 = v3
            _ = v9
        default:
            switch -560 {
            case 437:
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 uint = (v1 % ((v1) | 1))
                    var v12 int8 = 100
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            default:
                for fi13 := 0; fi13 < 6; fi13++ {
                    if (127 <= -128) {
                        for fi14 := 0; fi14 < 12; fi14++ {
                            v3 += v3
                            v15 := []int{fi14, fi14, (fi13 & -605), (fi13 & genFn0(1000, -411))}
                            _ = fi14
                            _ = v15
                        }
                        v16 := v4["a"]
                        _ = v16
                    }
                    var v17 uint64 = 175
                    _ = fi13
                    _ = v17
                }
            }
        }
        var v18 uint64 = 683
        v19 := GenS2{F0: 1, F1: false}
        _ = v18
        _ = v19
    } else {
        var v20 int8 = 1
        for fi21 := 0; fi21 < 13; fi21++ {
            switch ("a" + "a") {
            case "foo bar":
                switch "" {
                case "foo bar":
                    fmt.Println(v3, v3, v2)
                case "hello":
                    fmt.Println(v1, v3)
                }
            case "\n\t":
                for fi22 := 0; fi22 < 9; fi22++ {
                    fmt.Println(v3)
                    _ = fi22
                }
            }
            v1 += 0
            _ = fi21
        }
        var v23 uint = v1
        _ = v20
        _ = v23
    }
    var v24 uint32 = 70
    var v25 uint64 = 18446744073709551615
    v26 := v4[(("a" + "\n\t") + ("a" + ("foo bar" + "a")))]
    v2 += v2
    v27 := GenS3{F0: v2, F1: 112, F2: (v24 << ((0) & 63))}
    var v28 uint8 = v2
    if (false || false) {
        v29 := GenS4{F0: 18446744073709551615, F1: -31}
        v30 := v4[("" + v26)]
        v3 -= v3
        _ = v29
        _ = v30
    }
    v31 := GenS3{F0: v2, F1: (^(v28)), F2: (v24 / ((v24) | 1))}
    v26 = "\n\t"
    for fi32 := 0; fi32 < 8; fi32++ {
        v33 := GenS1{F0: 0, F1: 106, F2: (fi32 / ((0) | 1))}
        for fi34 := 0; fi34 < 2; fi34++ {
            switch fi32 {
            case 4294967295:
                v35 := GenS4{F0: v25, F1: 0}
                _ = v35
            case -349:
                v36 := GenS3{F0: (v28 + v28), F1: v28, F2: 42}
                _ = v36
            default:
                v37 := GenS1{F0: -1, F1: 127, F2: (fi34 + fi32)}
                _ = v37
            }
            v3 += v3
            _ = fi34
        }
        _ = fi32
        _ = v33
    }
    switch int(v2) {
    case 2147483647:
        delete(v4, v26)
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        v39 := GenS2{F0: 42, F1: false}
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
