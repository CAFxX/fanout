package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint8; F2 bool }
func genFn0(p0 uint8, p1 string, p2 string) int32 {
    v7 := []int{798, 0, 42, 9223372036854775807}
    if (0 <= -826) {
        v8 := GenS1{F0: ("" + "a"), F1: -1}
        v9 := GenS2{F0: 1, F1: 4294967295}
        var v10 int8 = -70
        _ = v10
        _ = v8
        _ = v9
    } else {
        p0 -= (^(p0))
        p1 = "\n\t"
    }
    var v11 int32 = 0
    if (284 < uint64(p0)) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        v12 := GenS2{F0: v11, F1: 2147483647}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return v11
}

func main() {
    var v1 int16 = 1
    var v2 uint8 = 127
    var v3 int64 = 42
    var v4 int = 2147483647
    v5 := []uint32{2147483648}
    v6 := []string{("" + "\n\t"), "a", "\n\t", "\n\t"}
    v3 -= (v3 + -1)
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int64 = (v3 << ((1) & 63))
            fmt.Println(fi14)
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v6) > 0 {
        v6[0] = ""
    }
    var v16 uint8 = v2
    var v17 int16 = v1
    var v18 uint64 = 2
    var v19 int32 = 0
    var v20 uint32 = 42
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 string = "hello"
        if true {
            for fi23 := 0; fi23 < 3; fi23++ {
                if (true || false) {
                    v24 := GenS2{F0: (-(v19)), F1: fi21}
                    _ = v24
                }
                if false {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v19 = genFn0((v2 / ((v16) | 1)), ("\n\t" + "a"), (v22 + ("foo bar" + v22)))
                }
                fmt.Println(v1, v20, v17)
                _ = fi23
            }
        }
        v25 := GenS2{F0: (v19 ^ v19), F1: (^(fi21))}
        _ = fi21
        _ = v22
        _ = v25
    }
    v26 := GenS2{F0: v19, F1: v4}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, v20)
    var v27 int16 = int16(v4)
    v28 := GenS3{F0: v4}
    if false {
        for fi29 := 0; fi29 < 5; fi29++ {
            var v30 uint16 = 1
            if true {
                v31 := GenS4{F0: (v19 >> ((0) & 63)), F1: (v2 * 255), F2: (!((0.1 != (0.0 + 0.0)) || (!(!(1e10 <= 0.0)))))}
                _ = v31
            } else {
                _ = len(v6)
                _ = len(v5)
            }
            _ = fi29
            _ = v30
        }
        v32 := GenS4{F0: (v19 - (v19 % ((v19) | 1))), F1: (^(v16)), F2: false}
        _ = v32
    } else {
        _ = len(v5)
        fmt.Println(v18)
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 string = ("" + "foo bar")
            v27 = 879
            var v35 int64 = int64(v19)
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    v36 := GenS4{F0: (-(v19)), F1: 213, F2: true}
    var v37 int16 = (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
