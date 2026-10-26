package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 int; F2 string }
func genFn0(p0 uint, p1 int, p2 int32) uint32 {
    var v6 int32 = (-(p2))
    var v7 int16 = -914
    for fi8 := 0; fi8 < 1; fi8++ {
        if false {
            v9 := GenS1{F0: -81, F1: v7, F2: (65535 != 187)}
            _ = v9
        } else {
            fmt.Println(v6, p2, v6)
        }
        v10 := []int{p1, p1, 479}
        var v11 int32 = (p2 >> ((p0) & 63))
        _ = fi8
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 int = 127
    var v2 int8 = 1
    var v3 string = (("hello" + "a") + "a")
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    for fi12 := 0; fi12 < 17; fi12++ {
        v2 = v2
        v13 := GenS1{F0: 1, F1: 1000, F2: false}
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 uint16 = 288
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v2 += -1
    switch v1 {
    case 1:
        switch v1 {
        case -2147483648:
            fmt.Println(v1)
        default:
            for fi16 := 0; fi16 < 9; fi16++ {
                v1 -= (-(fi16))
                fmt.Println(v2, v1, v3)
                if true {
                    if false {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            var v18 int16 = 42
                            _ = fi17
                            _ = v18
                        }
                        v3 = ((("foo bar" + v3) + "hello") + "foo bar")
                        v19 := GenS2{F0: "foo bar"}
                        _ = v19
                    } else {
                        var v20 int = (-(fi16))
                        v1 = 32768
                        _ = v20
                    }
                    v21 := GenS3{F0: 42}
                    if ((((127 > 0) || false) || (genFn0(75, (v1 + fi16), 2147483647) <= 2147483647)) || false) {
                        v22 := GenS2{F0: ""}
                        _ = v22
                    }
                    _ = v21
                }
                _ = fi16
            }
        }
    case 42:
        v23 := GenS4{F0: 0}
        _ = v23
    case 128:
        v24 := []int{(v1 / ((v1) | 1)), 4294967295, v1}
        _ = v24
    default:
        var v25 int = 1
        _ = v25
    }
    v26 := []int{v1, v1}
    for fi27 := 0; fi27 < 3; fi27++ {
        delete(v5, "\n\t")
        if (1000 == -32768) {
            var v28 int32 = -1
            v1 += int(v28)
            v28 += (v28 / ((42) | 1))
            _ = v28
        } else {
            v3 = ""
        }
        fmt.Println(v2)
        _ = fi27
    }
    v29 := GenS5{F0: v2, F1: v1, F2: ("\n\t" + v3)}
    switch (v1 - (-(v1))) {
    case 733:
        v30 := GenS4{F0: 228}
        _ = v30
    default:
        fmt.Println(v2, v1)
    }
    var v31 uint = 65535
    if (!(4294967295 < v31)) {
        var v32 uint32 = 190
        _ = v32
    } else {
        v2 += v2
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        if false {
            var v34 float64 = 0.0
            var v35 int = (v1 | (v1 * (v1 + (fi33 << ((v31) & 63)))))
            _ = v34
            _ = v35
        } else {
            var v36 int = (fi33 ^ 925)
            if len(v26) > 0 {
                v26[0] = v1
            }
            fmt.Println(v3)
            _ = v36
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
