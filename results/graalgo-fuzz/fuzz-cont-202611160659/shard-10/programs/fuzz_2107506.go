package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int8 }
func genFn0() int16 {
    v4 := GenS1{F0: ((float32(1.5) - (float32((1.5 + -2.25)) - -2.25)) != 1.5)}
    var v5 uint64 = 0
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 float32 = (0.0 - 0.0)
    var v2 uint32 = 42
    var v3 bool = false
    v6 := []int{-676, 9223372036854775807}
    var v7 string = "foo bar"
    v8 := GenS1{F0: (!(!true))}
    if (4294967295 == 645) {
        for fi9 := 0; fi9 < 13; fi9++ {
            _ = len(v6)
            _ = fi9
        }
        v10 := GenS1{F0: (!false)}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 uint32 = 42
            _ = fi11
            _ = v12
        }
        if v3 {
            if (!(1 != 42)) {
                v7 = ("a" + (v7 + ("" + ("\n\t" + (v7 + v7)))))
                if v3 {
                    switch -724 {
                    case -9223372036854775808:
                        var v13 float64 = (1.5 / ((0.1 / float64(0.1)) + 1e10))
                        _ = v13
                    }
                    var v14 int8 = 127
                    _ = v14
                }
                fmt.Println(v2, v7, v3)
            } else {
                for fi15 := 0; fi15 < 20; fi15++ {
                    if (v7 > "a") {
                        _ = len(v6)
                        if v3 {
                            for fi16 := 0; fi16 < 7; fi16++ {
                                fmt.Println(v1, v3)
                                v7 = v7
                                _ = fi16
                            }
                        } else {
                            var v17 uint32 = v2
                            _ = v17
                        }
                    } else {
                        var v18 int = fi15
                        var v19 int16 = 42
                        _ = v18
                        _ = v19
                    }
                    var v20 int8 = 127
                    v21 := GenS1{F0: v3}
                    _ = fi15
                    _ = v20
                    _ = v21
                }
            }
            v22 := GenS2{F0: 42, F1: 1}
            v2 += (v2 * v2)
            _ = v22
        } else {
            switch -2147483648 {
            case 32768:
                if true {
                    v7 = v7
                    for fi23 := 0; fi23 < 9; fi23++ {
                        if (254 >= 103) {
                            _ = len(v6)
                        }
                        v24 := GenS3{F0: 18446744073709551615, F1: 2}
                        _ = fi23
                        _ = v24
                    }
                }
            default:
                for fi25 := 0; fi25 < 14; fi25++ {
                    var v26 uint32 = 0
                    _ = fi25
                    _ = v26
                }
            }
            v27 := GenS2{F0: 9223372036854775807, F1: 127}
            if v3 {
                v3 = (612 <= -1)
            } else {
                v6 = append(v6, 628)
                v2 -= 2
            }
            _ = v27
        }
        v2 += (v2 >> ((842) & 63))
    }
    switch -241 {
    case 32767:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case 127:
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS3{F0: 4294967295, F1: 10}
            if (v3 || (int32(v2) >= 2147483647)) {
                var v30 uint64 = 18446744073709551615
                _ = v30
            }
            _ = fi28
            _ = v29
        }
    }
    v7 = v7
    fmt.Println(v7, v1, v3)
    v31 := GenS3{F0: 4294967295, F1: 430}
    var v32 bool = (v3 && false)
    v6 = append(v6, -1)
    v33 := GenS4{F0: 53, F1: 127}
    switch "" {
    case "foo bar":
        v2 += (v2 / ((248) | 1))
    default:
        v6 = append(v6, 4294967295)
    }
    v34 := GenS3{F0: 598, F1: 255}
    for fi35 := 0; fi35 < 15; fi35++ {
        var v36 uint8 = 1
        if v32 {
            var v37 string = "hello"
            _ = v37
        } else {
            var v38 int8 = 84
            _ = len(v6)
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v6
    _ = v7
    _ = v8
}
