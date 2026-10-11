package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 uint8 }
func genFn0(p0 uint16, p1 int8, p2 int32) int {
    if true {
        if false {
            if false {
                v5 := []int{-990, -797, -9223372036854775808, 471}
                for fi6 := 0; fi6 < 16; fi6++ {
                    v7 := GenS1{F0: 2}
                    var v8 int64 = 294
                    var v9 uint32 = 4294967295
                    _ = fi6
                    _ = v7
                    _ = v8
                    _ = v9
                }
                _ = v5
            }
        }
    } else {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: 255}
            _ = fi10
            _ = v11
        }
    }
    v12 := []int{-1, -9223372036854775808, -777, -184}
    var v13 int = 32767
    p2 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 646
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 2147483647
    v3 := []int{305, 128, 127, -2147483648}
    var v4 string = "a"
    var v14 int32 = (v2 * -2147483648)
    v15 := GenS2{F0: 699, F1: "a", F2: 1}
    for fi16 := 0; fi16 < 9; fi16++ {
        v1[v4] = genFn0(1, 1, (v2 * v2))
        v4 = v4
        _ = fi16
    }
    var v17 uint64 = 42
    if ((675 <= 0) || false) {
        for fi18 := 0; fi18 < 10; fi18++ {
            switch fi18 {
            case 32767:
                if ((32767 > 0) && (42 == 643)) {
                    fmt.Println(v4)
                    switch fi18 {
                    case -1:
                        switch "\n\t" {
                        case "hello":
                            if false {
                                var v19 uint16 = 0
                                v4 = (v4 + v4)
                                if false {
                                    v20 := GenS1{F0: (v19 / ((v19) | 1))}
                                    v19 = v19
                                    var v21 uint16 = (v19 >> ((170) & 63))
                                    _ = v20
                                    _ = v21
                                } else {
                                    _ = len(v1)
                                    v19 += (^(v19))
                                }
                                _ = v19
                            } else {
                                var v22 uint16 = 42
                                v23 := GenS2{F0: 1, F1: v4, F2: 127}
                                _ = v22
                                _ = v23
                            }
                        default:
                            for fi24 := 0; fi24 < 7; fi24++ {
                                v2 = (v14 % ((-834) | 1))
                                fmt.Println(fi18)
                                if (0 <= -9223372036854775808) {
                                    switch fi18 {
                                    case 42:
                                        var v25 bool = false
                                        var v26 uint32 = 2147483647
                                        _ = v25
                                        _ = v26
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v27 := GenS2{F0: 9223372036854775807, F1: ((("hello" + "") + v4) + ((v4 + v4) + v4)), F2: 42}
                                    _ = v27
                                }
                                _ = fi24
                            }
                        }
                    case 732:
                        var v28 int16 = -496
                        _ = v28
                    default:
                        v29 := GenS1{F0: 2}
                        _ = v29
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            _ = fi18
        }
        fmt.Println(v2, v4)
    }
    var v30 int8 = 100
    for fi31 := 0; fi31 < 15; fi31++ {
        fmt.Println(v30, v4, v4)
        if (!(127 < 1)) {
            var v32 uint16 = 687
            if len(v3) > 0 {
                v3[0] = 9223372036854775807
            }
            _ = v32
        } else {
            if ((42 == 42) && false) {
                v14 = v2
            }
            switch "a" {
            case "hello":
                v2 = (^(v14))
            case "\n\t":
                v14 += (v14 & -1)
            default:
                var v33 int16 = 42
                _ = v33
            }
            _ = len(v3)
        }
        _ = fi31
    }
    if (!(4294967295 == 42)) {
        if len(v3) > 0 {
            v3[0] = -244
        }
        fmt.Println(v17, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
