package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }
func genFn0(p0 uint32, p1 uint64) int8 {
    v6 := GenS1{F0: 213, F1: -843, F2: 0}
    var v7 int32 = -1
    var v8 uint64 = (^(p1))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1() int16 {
    var v9 uint64 = 59
    var v10 uint8 = 1
    fmt.Println(v9)
    v10 -= v10
    fmt.Println(v9, v9)
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 uint32 = 316
    var v4 int = 9223372036854775807
    v5 := make(map[uint32]uint32)
    v11 := GenS1{F0: 1, F1: -32768, F2: (v4 - 128)}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 float32 = 3.14159
        _ = fi12
        _ = v13
    }
    _ = len(v5)
    var v14 float64 = -2.25
    if true {
        var v15 uint = 992
        switch v4 {
        case -2147483648:
            fmt.Println(v4, v14)
        }
        _ = v15
    } else {
        fmt.Println(v14, v4)
    }
    var v16 uint64 = 679
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := v5[65535]
        delete(v2, v4)
        fmt.Println(fi17)
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 14, F1: genFn1(), F2: v4}
    if ((true || (true || true)) && (3.14159 == float32(v14))) {
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS1{F0: 24, F1: -916, F2: v4}
            for fi22 := 0; fi22 < 3; fi22++ {
                v14 = -2.25
                switch fi22 {
                case 1:
                    switch fi22 {
                    case -870:
                        var v23 string = "a"
                        _ = v23
                    }
                case 559:
                    fmt.Println(v3, v3)
                case 128:
                    v16 -= v16
                default:
                    fmt.Println(v16)
                }
                var v24 int64 = -768
                _ = fi22
                _ = v24
            }
            _ = fi20
            _ = v21
        }
        var v25 int = v4
        _ = v25
    } else {
        var v26 int32 = 0
        _ = v26
    }
    v27 := GenS1{F0: 2, F1: 704, F2: -1}
    var v28 int64 = -9223372036854775808
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 int64 = -9223372036854775808
        var v31 int = fi29
        _ = fi29
        _ = v30
        _ = v31
    }
    v28 = -579
    for fi32 := 0; fi32 < 11; fi32++ {
        v33 := GenS1{F0: 74, F1: 306, F2: (v4 & -714)}
        var v34 int64 = v28
        switch ("a" + "") {
        case "":
            v4 = 606
        default:
            v4 = (v4 >> ((65535) & 63))
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
