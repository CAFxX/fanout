package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int32 }
func genFn0(p0 bool, p1 int16, p2 uint16) int16 {
    var v7 uint32 = 65535
    switch 128 {
    case -540:
        v8 := []int{9223372036854775807, -727, 127, 32767}
        _ = v8
    case -2147483648:
        v7 += v7
    case 707:
        v7 = 91
    default:
        v9 := GenS1{F0: 2147483647, F1: 4294967295}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 & (p1 / (((p1 << ((480) & 63))) | 1)))
}
func genFn1(p0 uint8, p1 uint64, p2 int32) uint16 {
    var v10 int = 42
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 float32 = 3.14159
        _ = fi11
        _ = v12
    }
    v13 := []int{(v10 ^ v10), (v10 | v10), (v10 >> ((255) & 63)), (v10 & (v10 ^ v10))}
    fmt.Println(p2, p1)
    var v14 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    _ = v14
    return 13
}

func main() {
    var v1 string = ("foo bar" + ("hello" + ("a" + (("hello" + ("foo bar" + "")) + ""))))
    var v2 uint = 1
    var v3 int16 = 1000
    v4 := make(map[int]bool)
    var v5 float32 = (1e10 + (float32(3.14159) - 0.1))
    var v6 uint64 = 246
    switch 1 {
    case -2147483648:
        var v15 uint16 = 981
        _ = v15
    case 42:
        fmt.Println(v5, v1, v1)
    default:
        var v16 uint = (^(v2))
        _ = v16
    }
    fmt.Println(v3, v3, v1)
    var v17 bool = false
    v18 := GenS2{F0: 1}
    if (42 <= 127) {
        var v19 int32 = -1
        var v20 string = v1
        _ = v19
        _ = v20
    } else {
        switch "a" {
        case "\n\t":
            v21 := []int{}
            _ = v21
        case "a":
            switch 722 {
            case -9223372036854775808:
                switch (("" + v1) + (v1 + (("foo bar" + (v1 + "a")) + (v1 + v1)))) {
                case "foo bar":
                    var v22 int16 = v3
                    _ = v22
                }
            case 1:
                v23 := GenS1{F0: -221, F1: 0}
                _ = v23
            default:
                fmt.Println(v3)
            }
        default:
            var v24 uint32 = 1
            _ = v24
        }
        var v25 bool = v17
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 uint64 = v6
            v28 := []int{(fi26 | (fi26 | int(v2)))}
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v25
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        fmt.Println(fi29, v6, v2)
        _ = fi29
    }
    v30 := GenS3{F0: (-(v3)), F1: 0, F2: 97}
    v31 := []int{int(v6)}
    switch 32768 {
    case 4294967295:
        var v32 float64 = float64((v5 * v5))
        _ = v32
    default:
        v33 := GenS1{F0: -9223372036854775808, F1: 127}
        _ = v33
    }
    fmt.Println(v2, v1)
    var v34 bool = (!true)
    v35 := GenS4{F0: 16, F1: -634}
    v5 -= float32(0.0)
    v36 := GenS1{F0: -9223372036854775808, F1: 959}
    for fi37 := 0; fi37 < 14; fi37++ {
        if (genFn0(v17, (v3 >> ((v2) & 63)), 2) >= v3) {
            var v38 int64 = -9223372036854775808
            var v39 bool = v17
            _ = v38
            _ = v39
        } else {
            v17 = v34
        }
        for fi40 := 0; fi40 < 2; fi40++ {
            fmt.Println(fi37, v5, v17)
            if v34 {
                v1 = "a"
                v1 = (v1 + "a")
                switch "foo bar" {
                case "":
                    fmt.Println(fi40, v34)
                }
            }
            _ = fi40
        }
        v41 := GenS1{F0: fi37, F1: fi37}
        _ = fi37
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
