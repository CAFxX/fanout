package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 bool; F1 int; F2 bool }
type GenS4 struct { F0 uint8; F1 uint8; F2 int64 }
func genFn0(p0 bool, p1 string) uint16 {
    fmt.Println(p1)
    var v6 bool = p0
    if false {
        p0 = (1e10 == (0.1 * (1e10 - 3.14159)))
        p0 = (p1 != ("foo bar" + ""))
        var v7 int8 = 1
        _ = v7
    } else {
        v8 := []int{1}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 bool = ((1e10 * -2.25) <= 1e10)
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    var v5 float32 = 1.5
    for fi9 := 0; fi9 < 14; fi9++ {
        if false {
            v1 = (!(42 == 899))
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            fmt.Println(fi9, v5, v1)
            var v11 uint8 = 1
            fmt.Println(v1, v11)
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 float32 = float32(3.14159)
    var v13 uint16 = 596
    var v14 string = ""
    v15 := GenS1{F0: (v13 | genFn0(true, "hello"))}
    v14 = v14
    if true {
        v16 := []int{1, 234}
        var v17 float64 = float64(((-2.25 - (float32(((1.5 - float64(1e10)) + 1e10)) - -2.25)) - v5))
        v14 = v14
        _ = v16
        _ = v17
    }
    if (!((-45 == 1) && (65535 == 2147483648))) {
        v18 := GenS1{F0: (^(v13))}
        v19 := v4[-123]
        v20 := GenS2{F0: 147, F1: 344}
        _ = v18
        _ = v19
        _ = v20
    }
    v3[601] = "a"
    if (-9223372036854775808 == 1) {
        v21 := GenS3{F0: (!v1), F1: -1, F2: (!((v5 - v5) > v12))}
        _ = v21
    }
    v22 := GenS3{F0: false, F1: -2147483648, F2: (-128 < -1)}
    var v23 uint = 0
    var v24 int16 = -32768
    var v25 uint64 = 4294967295
    for fi26 := 0; fi26 < 20; fi26++ {
        v12 += -2.25
        _ = fi26
    }
    if (0 != 2147483647) {
        var v27 string = v14
        if v1 {
            for fi28 := 0; fi28 < 12; fi28++ {
                for fi29 := 0; fi29 < 20; fi29++ {
                    switch fi29 {
                    case -777:
                        v30 := GenS4{F0: 0, F1: 128, F2: 21}
                        _ = v30
                    default:
                        var v31 uint64 = v25
                        _ = v31
                    }
                    _ = fi29
                }
                _ = fi28
            }
            var v32 uint64 = (v25 * (v25 / ((773) | 1)))
            fmt.Println(v23)
            _ = v32
        } else {
            var v33 string = v14
            var v34 int64 = 42
            switch (v14 + v27) {
            case "\n\t":
                switch -2147483648 {
                case 128:
                    fmt.Println(v5, v23)
                case 127:
                    v35 := GenS4{F0: 168, F1: uint8(v25), F2: v34}
                    _ = v35
                case -315:
                    v24 = 1
                default:
                    fmt.Println(v23)
                }
            }
            _ = v33
            _ = v34
        }
        _ = v27
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v36 bool = v1
        _ = v36
    case 42:
        var v37 bool = v1
        _ = v37
    case 128:
        fmt.Println(v2)
    default:
        for fi38 := 0; fi38 < 9; fi38++ {
            v14 = v14
            var v39 uint = 18446744073709551615
            var v40 uint64 = 0
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
