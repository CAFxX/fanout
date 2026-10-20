package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint }
type GenS2 struct { F0 string }
func genFn0(p0 uint8) int64 {
    p0 = p0
    if true {
        if false {
            v6 := []int{32768, 1, 945}
            var v7 uint64 = 18446744073709551615
            switch 42 {
            case 4294967295:
                v7 = v7
            case 128:
                if len(v6) > 0 {
                    v6[0] = 419
                }
            case 42:
                switch "foo bar" {
                case "\n\t":
                    v8 := GenS1{F0: (^(p0)), F1: 0, F2: 65535}
                    _ = v8
                }
            default:
                if true {
                    switch 4294967295 {
                    case 128:
                        fmt.Println(v7, v7)
                    case 32768:
                        var v9 int32 = 13
                        _ = v9
                    case 127:
                        v7 += v7
                    default:
                        fmt.Println(v7, p0)
                    }
                    for fi10 := 0; fi10 < 3; fi10++ {
                        v6 = append(v6, 59)
                        v11 := GenS1{F0: uint8(v7), F1: 1, F2: 18446744073709551615}
                        _ = fi10
                        _ = v11
                    }
                    var v12 int = 1
                    _ = v12
                } else {
                    if true {
                        v7 += v7
                    }
                }
            }
            _ = v6
            _ = v7
        }
    } else {
        if true {
            v13 := []int{128, -2147483648, 4294967295, -299}
            _ = v13
        } else {
            var v14 float64 = 1.5
            if (!true) {
                v14 += (v14 / (v14 - 1e10))
                var v15 int32 = 42
                _ = v15
            }
            v16 := []int{-2147483648, -2147483648, 313, -9223372036854775808}
            _ = v14
            _ = v16
        }
    }
    v17 := []int{}
    _ = p0
    _ = v17
    return -650
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    var v3 int16 = 422
    var v4 int8 = -64
    var v5 string = "\n\t"
    var v18 uint32 = 2147483648
    fmt.Println(v5)
    v19 := GenS1{F0: 140, F1: -1, F2: 0}
    if (true && false) {
        if true {
            v4 = v4
            v20 := GenS1{F0: 2, F1: 1, F2: 368}
            v5 = "foo bar"
            _ = v20
        }
        var v21 int16 = v3
        _ = v21
    }
    _ = len(v1)
    var v22 int8 = 1
    switch 0 {
    case -9223372036854775808:
        v4 = v22
    }
    var v23 uint64 = 18446744073709551615
    if true {
        v24 := GenS1{F0: 1, F1: 357, F2: 4294967295}
        var v25 float32 = float32(0.0)
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: 32, F1: genFn0(uint8(v23)), F2: 0}
    v4 -= 0
    switch 9223372036854775807 {
    case 32768:
        switch 2147483647 {
        case 42:
            fmt.Println(v3, v3)
        case 32767:
            v22 = (v4 ^ v4)
        case -1:
            var v27 int64 = genFn0(0)
            _ = v27
        default:
            v23 = v23
        }
    case -980:
        v28 := []int{4294967295, 4294967295}
        _ = v28
    case 0:
        v29 := GenS2{F0: (("hello" + (v5 + "")) + v5)}
        _ = v29
    default:
        if true {
            if false {
                var v30 string = v5
                var v31 uint16 = 42
                _ = v30
                _ = v31
            }
            var v32 int16 = (v3 * v3)
            v33 := GenS1{F0: 42, F1: 1, F2: 4294967295}
            _ = v32
            _ = v33
        }
    }
    v34 := []int{-358, 42, -9223372036854775808}
    var v35 string = (v5 + ((v5 + (v5 + v5)) + v5))
    v36 := GenS1{F0: 127, F1: -1, F2: 255}
    fmt.Println(v35)
    switch 136 {
    case 9223372036854775807:
        var v37 uint = 465
        _ = v37
    case -530:
        switch v35 {
        case "foo bar":
            for fi38 := 0; fi38 < 10; fi38++ {
                v23 = v23
                v39 := GenS2{F0: ((v5 + "") + "foo bar")}
                _ = fi38
                _ = v39
            }
        case "":
            v35 = (("\n\t" + (("a" + v35) + v35)) + "hello")
        default:
            var v40 float64 = 0.1
            _ = v40
        }
    }
    var v41 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v5
}
