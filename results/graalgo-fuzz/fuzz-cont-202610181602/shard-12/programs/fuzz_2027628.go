package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int32 }
type GenS2 struct { F0 bool; F1 string; F2 int32 }
func genFn0() uint {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint32 = 4294967295
        fmt.Println(v6, fi5)
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 42
    var v8 uint16 = 243
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 float32 = float32(3.14159)
    v2 := make(map[int]int)
    var v3 float32 = v1
    var v4 bool = (4294967295 < 598)
    var v9 uint64 = 589
    if (1.5 == 0.1) {
        var v10 float32 = v1
        v11 := []int{175, 4294967295, 128}
        delete(v2, -825)
        _ = v10
        _ = v11
    }
    v1 = v3
    v1 = 0.0
    var v12 int = 2147483647
    v2[v12] = v12
    switch (v12 ^ 4294967295) {
    case -2147483648:
        switch "" {
        case "":
            if (127 == 120) {
                switch -1 {
                case 454:
                    fmt.Println(v4, v4, v1)
                case -428:
                    var v13 int8 = 42
                    _ = v13
                case 32767:
                    if (2 <= 255) {
                        v14 := v2[(v12 >> ((uint(v9)) & 63))]
                        _ = v14
                    } else {
                        if v4 {
                            v15 := v2[(^(v12))]
                            switch "" {
                            case "hello":
                                v16 := []int{}
                                _ = v16
                            case "a":
                                v17 := GenS1{F0: 65535, F1: "a", F2: -996}
                                _ = v17
                            case "\n\t":
                                if ((false && true) && false) {
                                    switch (("" + ("foo bar" + "a")) + "foo bar") {
                                    case "":
                                        var v18 bool = true
                                        _ = v18
                                    case "\n\t":
                                        v12 = (v15 % ((v12) | 1))
                                    }
                                    for fi19 := 0; fi19 < 11; fi19++ {
                                        fmt.Println(v3, v9)
                                        for fi20 := 0; fi20 < 10; fi20++ {
                                            var v21 int = v12
                                            _ = fi20
                                            _ = v21
                                        }
                                        if v4 {
                                            v1 -= (3.14159 + v3)
                                            v22 := []int{v12, (v15 + v12), (fi19 | v12)}
                                            switch 32767 {
                                            case 829:
                                                delete(v2, v12)
                                            }
                                            _ = v22
                                        }
                                        _ = fi19
                                    }
                                    var v23 int = v12
                                    _ = v23
                                }
                            default:
                                if (false || (v4 && (1e10 <= (v3 * 1.5)))) {
                                    v24 := GenS2{F0: (0 != 9223372036854775807), F1: "", F2: -2147483648}
                                    _ = v24
                                }
                            }
                            _ = v15
                        } else {
                            v2[v12] = (v12 / ((32768) | 1))
                        }
                    }
                default:
                    v3 += ((1.5 + (v1 * float32(((-2.25 / 0.1) + (0.1 * -2.25))))) * v1)
                }
                v2[(v12 | (v12 - v12))] = 32767
            } else {
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 int32 = 42
                    _ = fi25
                    _ = v26
                }
                var v27 uint8 = 127
                _ = v27
            }
        case "hello":
            var v28 int64 = 0
            _ = v28
        case "a":
            for fi29 := 0; fi29 < 19; fi29++ {
                v30 := []int{v12, (^(fi29)), 32768}
                var v31 uint32 = 4294967295
                delete(v2, 1)
                _ = fi29
                _ = v30
                _ = v31
            }
        default:
            v9 -= 0
        }
    case 2147483647:
        if false {
            fmt.Println(v3, v9)
            for fi32 := 0; fi32 < 18; fi32++ {
                v12 += -191
                v33 := GenS1{F0: 0, F1: "", F2: -1}
                _ = fi32
                _ = v33
            }
        }
    case 199:
        var v34 int8 = -82
        _ = v34
    default:
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 int16 = 42
            v37 := []int{v12, (^(v12)), (fi35 << ((genFn0()) & 63))}
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    if v4 {
        for fi38 := 0; fi38 < 20; fi38++ {
            v3 = (v1 - (-2.25 - 0.0))
            _ = fi38
        }
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        v40 := GenS1{F0: 2, F1: "hello", F2: 42}
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
