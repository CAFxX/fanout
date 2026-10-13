package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int8; F1 uint16; F2 int64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint8) int32 {
    p0 = 1
    switch 42 {
    case -410:
        var v6 int = 42
        _ = v6
    case 503:
        fmt.Println(p0)
    case 0:
        var v7 int8 = 0
        _ = v7
    }
    p0 += p0
    var v8 int16 = 1000
    p0 -= (p0 + 2)
    _ = p0
    _ = v8
    return -2147483648
}

func main() {
    var v1 int32 = 42
    var v2 float32 = 0.0
    v3 := make(map[uint32]int)
    v4 := []int{-1, 1, -9223372036854775808}
    var v5 float64 = 0.1
    v1 = v1
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int8 = 0
        var v11 uint64 = 42
        v5 = 0.0
        _ = fi9
        _ = v10
        _ = v11
    }
    if false {
        var v12 int = int(v1)
        _ = v12
    } else {
        switch "" {
        case "foo bar":
            for fi13 := 0; fi13 < 17; fi13++ {
                fmt.Println(v2, v5)
                var v14 int32 = v1
                v15 := GenS1{F0: 763}
                _ = fi13
                _ = v14
                _ = v15
            }
        case "\n\t":
            v2 = v2
        default:
            switch "\n\t" {
            case "a":
                fmt.Println(v5, v5)
            case "hello":
                var v16 uint16 = 838
                _ = v16
            case "\n\t":
                if (0 != 486) {
                    var v17 string = "hello"
                    var v18 int16 = 32767
                    _ = v17
                    _ = v18
                }
            default:
                var v19 string = ((((("a" + "a") + ("\n\t" + "hello")) + ("" + "\n\t")) + ("foo bar" + ("a" + "hello"))) + "a")
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            v2 += v2
            if (0 >= 9223372036854775807) {
                var v21 int = (^(fi20))
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 int = -1
                    _ = fi22
                    _ = v23
                }
                _ = v21
            } else {
                v1 = 952
                var v24 bool = (100 <= 1)
                var v25 uint32 = 2147483647
                _ = v24
                _ = v25
            }
            _ = fi20
        }
    }
    fmt.Println(v2, v5, v1)
    v26 := GenS2{F0: 127, F1: 42}
    switch "hello" {
    case "":
        fmt.Println(v1, v1, v1)
    case "hello":
        v5 += 1.5
    default:
        for fi27 := 0; fi27 < 18; fi27++ {
            for fi28 := 0; fi28 < 13; fi28++ {
                v29 := GenS3{F0: 42, F1: 1, F2: -1}
                _ = fi28
                _ = v29
            }
            v30 := GenS1{F0: 328}
            _ = fi27
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        v5 = v5
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 uint64 = 4294967295
            switch "" {
            case "foo bar":
                v34 := GenS2{F0: 73, F1: 820}
                _ = v34
            case "a":
                fmt.Println(v5, v5, v5)
            case "":
                v1 = 0
            }
            if ((65535 != 65535) || true) {
                v35 := GenS4{F0: v1}
                var v36 int16 = 42
                _ = v35
                _ = v36
            } else {
                var v37 uint64 = v33
                _ = v37
            }
            _ = fi32
            _ = v33
        }
        var v38 bool = false
        _ = fi31
        _ = v38
    }
    var v39 float32 = (v2 / 1.5)
    var v40 int32 = (v1 | 3)
    v41 := GenS4{F0: (-(v1))}
    var v42 int8 = -1
    v43 := GenS1{F0: 2147483648}
    v3[1] = 32767
    v5 = (v5 + v5)
    v42 = v42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
