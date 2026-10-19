package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint16; F2 uint }
func genFn0(p0 string) int {
    var v6 uint = 355
    if (v6 >= v6) {
        var v7 int8 = -97
        v8 := GenS1{F0: p0, F1: 1}
        for fi9 := 0; fi9 < 18; fi9++ {
            v7 = v7
            v7 += 0
            var v10 uint64 = 60
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        v11 := []int{42, 128, -1}
        v6 = 83
        _ = v11
    }
    v12 := GenS1{F0: "a", F1: 65535}
    v13 := GenS2{F0: 1, F1: 1000}
    _ = p0
    _ = v12
    _ = v13
    _ = v6
    return 32767
}
func genFn1(p0 uint) int32 {
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS3{F0: 32767, F1: ((fi14 + fi14) > genFn0("a")), F2: 1}
        var v16 uint64 = 2
        p0 = p0
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int8 = -89
    var v18 uint = p0
    fmt.Println(v18)
    _ = p0
    _ = v17
    _ = v18
    return -27
}

func main() {
    var v1 int = 0
    var v2 uint32 = 65535
    var v3 bool = (((v2 * v2) >= v2) && false)
    var v4 bool = ((4294967295 <= 1) || (9223372036854775807 > (v1 % ((-2147483648) | 1))))
    v5 := make(map[uint32]string)
    if false {
        for fi19 := 0; fi19 < 2; fi19++ {
            v2 -= 65535
            _ = fi19
        }
    }
    v20 := GenS4{F0: 100, F1: 42, F2: 138}
    v1 = v1
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 int16 = 1
        v22 = v22
        _ = fi21
        _ = v22
    }
    var v23 uint64 = 2
    if (float32(0.1) >= ((1.5 / 0.1) - 0.1)) {
        switch "foo bar" {
        case "foo bar":
            for fi24 := 0; fi24 < 3; fi24++ {
                switch "foo bar" {
                case "a":
                    var v25 uint32 = 2147483647
                    _ = v25
                case "hello":
                    var v26 int64 = 0
                    _ = v26
                case "foo bar":
                    _ = len(v5)
                }
                var v27 uint = 0
                fmt.Println(v2, v27)
                _ = fi24
                _ = v27
            }
        case "hello":
            v28 := []int{v1, genFn0("hello"), 0}
            _ = v28
        case "a":
            v3 = false
        default:
            var v29 bool = v4
            _ = v29
        }
        for fi30 := 0; fi30 < 5; fi30++ {
            fmt.Println(v3)
            _ = fi30
        }
        fmt.Println(v1)
    } else {
        fmt.Println(v4)
        v4 = false
    }
    v3 = (((((!(362 < 42)) || v4) || v3) && v3) || v3)
    fmt.Println(v2, v2)
    v31 := GenS1{F0: "a", F1: (v2 | (^(v2)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
