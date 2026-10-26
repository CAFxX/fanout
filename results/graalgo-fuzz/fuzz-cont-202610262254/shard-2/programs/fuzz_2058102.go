package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
func genFn0() int16 {
    var v6 uint64 = 0
    var v7 int8 = 42
    v7 -= v7
    v7 = 127
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 bool) uint32 {
    switch 32768 {
    case -780:
        var v8 int16 = 188
        _ = v8
    }
    var v9 string = ("" + "foo bar")
    _ = p0
    _ = v9
    return 574
}

func main() {
    var v1 bool = (889 >= 1)
    var v2 uint32 = 65535
    v3 := make(map[int]int)
    var v4 int32 = 306
    v5 := make(map[string]bool)
    var v10 uint16 = 32767
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: 128}
        _ = fi11
        _ = v12
    }
    switch "hello" {
    case "a":
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(v10, v4, v1)
            if (!((0.0 < 3.14159) && true)) {
                v1 = (-32768 != -1)
                var v14 string = "\n\t"
                var v15 uint = 1000
                _ = v14
                _ = v15
            }
            switch fi13 {
            case 42:
                v10 -= v10
            case -2147483648:
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 float64 = 1e10
                    _ = fi16
                    _ = v17
                }
            default:
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 float32 = 0.1
                    _ = fi18
                    _ = v19
                }
            }
            _ = fi13
        }
    default:
        fmt.Println(v10, v1)
    }
    delete(v3, -287)
    if (((1.5 + (1e10 - -2.25)) - 1.5) != (((-2.25 * 3.14159) + (1e10 - 0.0)) + -2.25)) {
        for fi20 := 0; fi20 < 15; fi20++ {
            _ = len(v5)
            for fi21 := 0; fi21 < 17; fi21++ {
                v22 := GenS1{F0: 0}
                v23 := []int{(^(fi20)), (^(fi21))}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        v24 := v3[1]
        _ = v24
    } else {
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v4, v10, v2)
            _ = fi25
        }
        var v26 int32 = v4
        _ = v26
    }
    fmt.Println(v4)
    if true {
        for fi27 := 0; fi27 < 11; fi27++ {
            switch "hello" {
            case "hello":
                v28 := []int{}
                _ = v28
            case "foo bar":
                for fi29 := 0; fi29 < 14; fi29++ {
                    fmt.Println(fi27, fi29, v4)
                    fmt.Println(v1, fi27, fi27)
                    v10 -= (v10 % ((v10) | 1))
                    _ = fi29
                }
            default:
                var v30 int8 = -77
                _ = v30
            }
            _ = fi27
        }
        fmt.Println(v1, v4)
    } else {
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 uint8 = uint8(v10)
            _ = fi31
            _ = v32
        }
    }
    v33 := []int{}
    v34 := v3[-407]
    v10 += v10
    for fi35 := 0; fi35 < 11; fi35++ {
        fmt.Println(v2, v10, v1)
        v36 := GenS2{F0: -9223372036854775808}
        fmt.Println(v34, fi35)
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
