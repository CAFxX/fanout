package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint }
func genFn0(p0 int16, p1 uint64, p2 int64) int64 {
    v7 := GenS1{F0: 247, F1: 4294967295, F2: p2}
    var v8 uint32 = 216
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -997
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{1, 128, 4294967295}
    v3 := make(map[string]string)
    v4 := []string{(("hello" + "") + ((("hello" + "hello") + "\n\t") + "\n\t")), (("" + ("\n\t" + "a")) + ((("hello" + ((("\n\t" + ("foo bar" + "")) + ("\n\t" + (((("hello" + "a") + "hello") + ("hello" + "hello")) + ""))) + "\n\t")) + "a") + ("\n\t" + "a"))), ("" + "\n\t")}
    v5 := make(map[uint32]string)
    var v6 string = "foo bar"
    v6 = v6
    var v9 uint16 = 42
    switch "a" {
    case "\n\t":
        var v10 int32 = 42
        _ = v10
    default:
        var v11 int = 32767
        _ = v11
    }
    var v12 int32 = -1
    for fi13 := 0; fi13 < 14; fi13++ {
        if false {
            v6 = v6
        }
        _ = fi13
    }
    if true {
        if (!true) {
            if (!true) {
                v6 = v6
            } else {
                if true {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        if true {
                            v6 = (("\n\t" + v6) + "")
                        } else {
                            if ((v12 + v12) != v12) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v15 int16 = 1
                                var v16 uint32 = 2147483647
                                _ = v15
                                _ = v16
                            }
                            if (v6 >= v6) {
                                v9 -= v9
                            }
                        }
                        v17 := GenS1{F0: 0, F1: 0, F2: 9223372036854775807}
                        _ = fi14
                        _ = v17
                    }
                } else {
                    v6 = (v6 + (("\n\t" + (((v6 + v6) + v6) + v6)) + v6))
                    for fi18 := 0; fi18 < 19; fi18++ {
                        var v19 float64 = -2.25
                        fmt.Println(v19)
                        _ = fi18
                        _ = v19
                    }
                }
                var v20 uint32 = 0
                _ = v20
            }
            v2 = append(v2, 0)
            var v21 int8 = -1
            _ = v21
        } else {
            var v22 int32 = v12
            if true {
                v5[409] = ("hello" + ("" + (v6 + v6)))
                v12 = (-(v12))
            } else {
                v6 = v6
                v12 += v12
            }
            switch 9223372036854775807 {
            case 999:
                if true {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v9 = v9
                        fmt.Println(v12, v6, fi23)
                        _ = fi23
                    }
                    v24 := GenS2{F0: 193, F1: 10, F2: (v9 != (v9 >> ((907) & 63)))}
                    var v25 uint32 = 65535
                    _ = v24
                    _ = v25
                }
            default:
                fmt.Println(v22)
            }
            _ = v22
        }
    }
    v12 = 0
    var v26 uint32 = 317
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 uint8 = 1
        _ = fi27
        _ = v28
    }
    _ = len(v1)
    v6 = ("" + ("\n\t" + "hello"))
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 uint = 255
        v6 = "hello"
        _ = fi29
        _ = v30
    }
    v31 := v5[(v26 + 42)]
    var v32 int8 = 35
    v33 := GenS3{F0: (v26 * 904), F1: 657, F2: 42}
    v34 := GenS1{F0: 2, F1: (v26 | v26), F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
