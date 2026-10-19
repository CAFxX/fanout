package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint; F2 int16 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]int)
    var v3 uint8 = 42
    var v4 uint = 65535
    var v5 int16 = -424
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: ("a" + "\n\t"), F1: false, F2: 2}
        _ = len(v2)
        if false {
            var v8 uint = v4
            var v9 uint = v4
            v10 := GenS1{F0: "a", F1: (2147483648 >= 2147483648), F2: 573}
            _ = v10
            _ = v8
            _ = v9
        } else {
            v11 := []int{v1, -570}
            v12 := GenS1{F0: ((("\n\t" + "") + (((("a" + (("foo bar" + (("\n\t" + "\n\t") + (("hello" + (("" + "a") + "hello")) + ("\n\t" + "")))) + "hello")) + "") + "foo bar") + "hello")) + "hello"), F1: (1e10 < (0.1 + 1.5)), F2: 0}
            _ = v11
            _ = v12
        }
        _ = fi6
        _ = v7
    }
    v13 := v2[4294967295]
    v14 := GenS2{F0: int32(v5)}
    if false {
        var v15 float64 = -2.25
        if (2 != 479) {
            v16 := v2[4294967295]
            _ = len(v2)
            _ = v16
        } else {
            v3 = (v3 | 1)
        }
        _ = v15
    } else {
        delete(v2, 4294967295)
    }
    var v17 uint64 = 18446744073709551615
    fmt.Println(v3, v3, v3)
    var v18 string = ("hello" + "hello")
    if (!(143 == 1)) {
        v1 += (v1 / (((v13 >> (((v4 >> ((v4) & 63))) & 63))) | 1))
        var v19 int64 = 1
        var v20 uint16 = 2
        _ = v19
        _ = v20
    } else {
        v21 := GenS1{F0: v18, F1: true, F2: 4294967295}
        if true {
            var v22 float64 = 1e10
            for fi23 := 0; fi23 < 3; fi23++ {
                if (!true) {
                    v24 := v2[uint32(v17)]
                    fmt.Println(fi23, v1)
                    v25 := []int{v13}
                    _ = v24
                    _ = v25
                } else {
                    if (4294967295 <= v17) {
                        v26 := GenS2{F0: 42}
                        var v27 string = (v18 + "a")
                        v5 -= int16(fi23)
                        _ = v26
                        _ = v27
                    } else {
                        fmt.Println(fi23)
                    }
                    switch -677 {
                    case 128:
                        switch 42 {
                        case 0:
                            var v28 uint8 = 128
                            _ = v28
                        case 212:
                            fmt.Println(v3, v1)
                        case 1:
                            v29 := GenS3{F0: (42 < 0), F1: v4, F2: (v5 / ((v5) | 1))}
                            _ = v29
                        default:
                            v3 = 1
                        }
                    }
                }
                v2[65535] = v1
                _ = fi23
            }
            _ = v22
        }
        _ = v21
    }
    v30 := GenS4{F0: -9223372036854775808, F1: v4}
    var v31 int32 = 1
    v32 := GenS5{F0: uint32(v17)}
    if (42 == 462) {
        v33 := GenS6{F0: int32(v4), F1: (v17 >> ((v4) & 63))}
        v34 := v2[2]
        _ = v33
        _ = v34
    }
    var v35 int = (v1 / (((v13 / ((v13) | 1))) | 1))
    v36 := []int{(v1 >> (((^(v4))) & 63))}
    var v37 bool = (!false)
    fmt.Println(v18)
    if len(v36) > 0 {
        v36[0] = (^(v35))
    }
    var v38 uint32 = 438
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
