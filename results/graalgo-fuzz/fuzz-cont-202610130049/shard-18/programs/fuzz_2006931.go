package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int8; F1 int; F2 uint }
type GenS3 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    v3 := []int{}
    var v4 int16 = 0
    var v5 bool = ((-1 > -1) && ((905 >= 4294967295) && (true && ((42 >= 0) && false))))
    v4 = (v4 / ((v4) | 1))
    if true {
        for fi6 := 0; fi6 < 17; fi6++ {
            for fi7 := 0; fi7 < 14; fi7++ {
                v4 += v4
                var v8 string = "foo bar"
                v9 := GenS1{F0: -1, F1: v8}
                _ = fi7
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: 42, F1: ""}
            switch fi6 {
            case 1:
                var v11 uint = 18446744073709551615
                _ = v11
            case 9223372036854775807:
                v12 := GenS1{F0: 0, F1: ("a" + ("hello" + "\n\t"))}
                _ = v12
            case 884:
                var v13 int8 = -1
                _ = v13
            default:
                fmt.Println(v4)
            }
            _ = fi6
            _ = v10
        }
    } else {
        var v14 uint = 927
        _ = v14
    }
    delete(v2, 2)
    switch -495 {
    case 32768:
        if v5 {
            v4 += (v4 % ((0) | 1))
            var v15 int64 = 865
            var v16 int64 = v15
            _ = v15
            _ = v16
        }
    case 32767:
        v17 := GenS1{F0: 0, F1: "hello"}
        _ = v17
    case -2147483648:
        v18 := GenS1{F0: 9223372036854775807, F1: (((((("a" + "") + "a") + ("foo bar" + "a")) + "") + "foo bar") + ("" + (("\n\t" + "foo bar") + "\n\t")))}
        _ = v18
    default:
        switch 127 {
        case 32767:
            v5 = true
        case -1:
            v5 = (!v5)
        case 127:
            if (!v5) {
                v4 = (v4 & (^(v4)))
                if (!v5) {
                    if true {
                        var v19 uint32 = 4294967295
                        var v20 uint16 = 2
                        v5 = true
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int16 = v4
                        v22 := GenS2{F0: 0, F1: 127, F2: 2}
                        _ = v21
                        _ = v22
                    }
                }
            }
        }
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        if (0 == 147) {
            var v24 int32 = 42
            v25 := GenS2{F0: 42, F1: 42, F2: 342}
            v24 -= (v24 & v24)
            _ = v24
            _ = v25
        }
        v26 := GenS3{F0: 176, F1: (!false)}
        var v27 uint64 = 4294967295
        _ = fi23
        _ = v26
        _ = v27
    }
    var v28 uint = 42
    for fi29 := 0; fi29 < 11; fi29++ {
        v30 := GenS3{F0: 219, F1: false}
        _ = len(v3)
        _ = fi29
        _ = v30
    }
    v31 := GenS3{F0: 128, F1: v5}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
