package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint64; F2 int64 }
func genFn0() int8 {
    v6 := GenS1{F0: 1}
    var v7 string = ""
    if (3.14159 >= 3.14159) {
        if true {
            v8 := []int{}
            _ = v8
        }
        switch -1 {
        case 463:
            v9 := GenS2{F0: 496, F1: 604, F2: 408}
            _ = v9
        default:
            var v10 int16 = -563
            _ = v10
        }
    }
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint64) int32 {
    switch 9223372036854775807 {
    case 304:
        v11 := GenS1{F0: 1}
        _ = v11
    case 42:
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                fmt.Println(fi12, fi13)
                _ = fi13
            }
            v14 := []int{42}
            var v15 float32 = float32((-2.25 * 3.14159))
            _ = fi12
            _ = v14
            _ = v15
        }
    case 128:
        fmt.Println(p0)
    }
    v16 := []int{128, -996, -338, 396}
    _ = p0
    _ = v16
    return 0
}

func main() {
    var v1 int64 = 1
    var v2 bool = (!true)
    v3 := []string{"", (("" + ("" + "foo bar")) + "a"), "\n\t", ""}
    var v4 uint8 = 255
    v5 := make(map[string]int)
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS1{F0: 911}
        switch "\n\t" {
        case "a":
            _ = len(v3)
        }
        _ = fi17
        _ = v18
    }
    var v19 string = "hello"
    if (42 <= -9223372036854775808) {
        switch 32767 {
        case 2147483647:
            delete(v5, v19)
        case -1:
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 int8 = 42
                fmt.Println(v2, fi20)
                if (!((^(v4)) < v4)) {
                    if ((-2.25 / 1.5) >= 3.14159) {
                        switch fi20 {
                        case 880:
                            for fi22 := 0; fi22 < 18; fi22++ {
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                }
                                _ = fi22
                            }
                        case 128:
                            if v2 {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v19, v4)
                                if (((0.1 * float32(((-2.25 * 1.5) - ((1.5 * (1.5 + 0.1)) + 0.1)))) + 3.14159) == float32(0.0)) {
                                    var v23 uint8 = 0
                                    _ = v23
                                }
                            }
                        default:
                            v19 = v19
                        }
                        var v24 bool = (!((65535 != 4294967295) && (!(1e10 != 1.5))))
                        _ = v24
                    }
                    v21 += (v21 - genFn0())
                    var v25 int64 = (v1 / ((-9223372036854775808) | 1))
                    _ = v25
                } else {
                    var v26 int64 = 694
                    fmt.Println(v2, v1, v19)
                    var v27 int = fi20
                    _ = v26
                    _ = v27
                }
                _ = fi20
                _ = v21
            }
        }
    }
    var v28 uint8 = v4
    _ = len(v3)
    for fi29 := 0; fi29 < 3; fi29++ {
        if (!(3.14159 <= 1e10)) {
            var v30 uint32 = 165
            var v31 int32 = 1
            v2 = (v2 && (v2 && false))
            _ = v30
            _ = v31
        }
        _ = fi29
    }
    v28 += (v4 >> ((365) & 63))
    fmt.Println(v1)
    switch -749 {
    case 81:
        fmt.Println(v1, v19)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
