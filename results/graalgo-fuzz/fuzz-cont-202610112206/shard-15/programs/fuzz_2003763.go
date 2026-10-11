package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 string; F2 uint16 }
type GenS5 struct { F0 bool; F1 int8; F2 int64 }
type GenS6 struct { F0 int16 }
func genFn0() int64 {
    var v5 int8 = 1
    if false {
        switch "a" {
        case "foo bar":
            fmt.Println(v5)
        default:
            if false {
                for fi6 := 0; fi6 < 20; fi6++ {
                    if (!true) {
                        v7 := GenS1{F0: 686, F1: false}
                        if (!false) {
                            v8 := []int{599}
                            _ = v8
                        } else {
                            for fi9 := 0; fi9 < 3; fi9++ {
                                var v10 bool = ((0.1 != (((-2.25 * ((-2.25 / 1.0) - 0.0)) + ((1e10 / ((0.1 - 3.14159) - 1.5)) - 0.1)) / 1.5)) && false)
                                _ = fi9
                                _ = v10
                            }
                            var v11 float64 = -2.25
                            _ = v11
                        }
                        var v12 bool = false
                        _ = v12
                        _ = v7
                    }
                    v5 = 1
                    switch 9223372036854775807 {
                    case 1:
                        v13 := GenS1{F0: 507, F1: (1 != -2147483648)}
                        _ = v13
                    default:
                        fmt.Println(v5)
                    }
                    _ = fi6
                }
            } else {
                v14 := GenS2{F0: 0}
                var v15 int = 2147483647
                fmt.Println(v15, v5)
                _ = v14
                _ = v15
            }
        }
        if true {
            v16 := []int{}
            _ = len(v16)
            _ = v16
        }
        fmt.Println(v5)
    }
    var v17 bool = (!true)
    v18 := GenS1{F0: 255, F1: v17}
    switch -638 {
    case -1:
        v19 := GenS2{F0: 0}
        _ = v19
    }
    _ = v17
    _ = v18
    _ = v5
    return 1
}
func genFn1(p0 int32, p1 uint32, p2 uint) bool {
    switch (((("" + "\n\t") + "a") + "hello") + "a") {
    case "":
        p0 = p0
    case "hello":
        p1 -= p1
    default:
        if true {
            var v20 uint = p2
            v20 += (v20 / ((470) | 1))
            v21 := GenS3{F0: 4294967295}
            _ = v20
            _ = v21
        }
    }
    v22 := GenS4{F0: 4294967295, F1: ("foo bar" + ("hello" + "\n\t")), F2: 105}
    if false {
        var v23 uint32 = (p1 >> ((440) & 63))
        _ = v23
    } else {
        var v24 int64 = 0
        if ((!(-1 <= 32768)) && false) {
            v25 := GenS1{F0: 1, F1: true}
            v26 := []int{int(p1), -16, -9223372036854775808, 9223372036854775807}
            if len(v26) > 0 {
                v26[0] = 678
            }
            _ = v25
            _ = v26
        } else {
            var v27 uint8 = 1
            if false {
                var v28 int = 0
                var v29 int32 = p0
                _ = v28
                _ = v29
            } else {
                p2 = p2
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 float32 = (0.1 - -2.25)
                    v31 = (v31 * v31)
                    for fi32 := 0; fi32 < 11; fi32++ {
                        p1 = (p1 & (p1 / ((356) | 1)))
                        var v33 int = 0
                        if ((0.1 - (1.5 * -2.25)) != 1e10) {
                            p2 = 42
                        }
                        _ = fi32
                        _ = v33
                    }
                    _ = fi30
                    _ = v31
                }
            }
            v34 := []int{-305}
            _ = v27
            _ = v34
        }
        _ = v24
    }
    fmt.Println(p0)
    v35 := GenS3{F0: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v35
    return false
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 986
    var v3 string = "hello"
    var v4 int32 = 0
    switch -210 {
    case 32767:
        v2 = (^(v2))
    default:
        switch 128 {
        case 127:
            v36 := GenS5{F0: (v2 <= v2), F1: 1, F2: 9223372036854775807}
            _ = v36
        default:
            var v37 int32 = v4
            _ = v37
        }
    }
    v2 -= v2
    v4 = v4
    fmt.Println(v4, v4)
    var v38 uint8 = 42
    v39 := GenS6{F0: 32767}
    var v40 float32 = (3.14159 - 3.14159)
    v41 := GenS1{F0: 682, F1: ("\n\t" < "")}
    v40 = 0.0
    v42 := []int{}
    var v43 int32 = (v4 | v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
