package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int8 {
    if true {
        var v5 uint64 = 4294967295
        _ = v5
    }
    if false {
        var v6 int16 = 628
        _ = v6
    } else {
        var v7 uint64 = 2
        _ = v7
    }
    var v8 uint32 = 2147483648
    for fi9 := 0; fi9 < 5; fi9++ {
        switch fi9 {
        case 32767:
            for fi10 := 0; fi10 < 10; fi10++ {
                v8 = v8
                _ = fi10
            }
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            if false {
                v8 = 2147483647
                v8 = 645
                if (-283 <= -345) {
                    var v12 int = fi9
                    v12 += fi11
                    var v13 string = ("hello" + "\n\t")
                    fmt.Println(fi9, fi11, v12)
                    v14 := GenS1{F0: 18446744073709551615}
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    v8 = (v8 ^ (v8 + (^(v8))))
                    var v15 uint16 = 42
                    _ = v15
                }
            } else {
                var v16 bool = false
                v16 = v16
                _ = v16
            }
            var v17 int64 = 42
            _ = fi11
            _ = v17
        }
        switch "foo bar" {
        case "\n\t":
            v18 := []int{(fi9 & (fi9 >> ((113) & 63))), fi9, fi9}
            _ = v18
        case "foo bar":
            var v19 uint32 = v8
            _ = v19
        default:
            var v20 int8 = 127
            _ = v20
        }
        _ = fi9
    }
    _ = v8
    return 43
}
func genFn1(p0 bool, p1 int, p2 uint64) uint32 {
    var v21 bool = (!false)
    v22 := GenS1{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 float64 = 0.0
    var v4 float64 = float64(1.5)
    if false {
        _ = len(v2)
        switch 4294967295 {
        case 9223372036854775807:
            for fi23 := 0; fi23 < 1; fi23++ {
                fmt.Println(v3, fi23)
                for fi24 := 0; fi24 < 9; fi24++ {
                    delete(v2, 342)
                    _ = fi24
                }
                for fi25 := 0; fi25 < 1; fi25++ {
                    v26 := GenS1{F0: 4294967295}
                    delete(v1, (fi23 / (((^(fi23))) | 1)))
                    var v27 string = "hello"
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = fi23
            }
        default:
            for fi28 := 0; fi28 < 16; fi28++ {
                v4 += float64(1e10)
                fmt.Println(v4, v4, v3)
                _ = fi28
            }
        }
    } else {
        var v29 uint16 = 255
        _ = v29
    }
    v4 -= (v4 / ((((1e10 - (float64(3.14159) / v3)) / ((v4 / 0.1) * v4)) + v3) - float64(((1e10 / 1e10) / -2.25))))
    fmt.Println(v3)
    switch -9223372036854775808 {
    case -392:
        delete(v2, 42)
    }
    var v30 int = 127
    v4 -= float64(3.14159)
    fmt.Println(v30, v30)
    v3 = 1.5
    fmt.Println(v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
