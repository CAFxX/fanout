package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }
type GenS3 struct { F0 string; F1 uint32; F2 uint16 }
func genFn0() uint64 {
    var v6 string = ("" + "a")
    var v7 int64 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(v7)
        if true {
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := GenS1{F0: (v7 + 377)}
                if (-32768 >= -497) {
                    v11 := GenS2{F0: v6, F1: v7, F2: 2}
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
            v7 += v7
        }
        v12 := GenS1{F0: v7}
        _ = fi8
        _ = v12
    }
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 bool) bool {
    v13 := []int{42}
    if (3.14159 == 1e10) {
        fmt.Println(p1, p1)
        for fi14 := 0; fi14 < 2; fi14++ {
            p0 = (!(p0 && false))
            var v15 uint64 = 474
            _ = fi14
            _ = v15
        }
    } else {
        for fi16 := 0; fi16 < 16; fi16++ {
            if (!false) {
                var v17 int32 = 1
                _ = v17
            } else {
                var v18 uint64 = 1
                v18 = v18
                _ = v18
            }
            var v19 string = ((("hello" + ((("foo bar" + "") + ("" + ("a" + "foo bar"))) + (("\n\t" + (("hello" + "hello") + ("" + ""))) + (("hello" + "") + "a")))) + ("" + "hello")) + (("\n\t" + "hello") + ("\n\t" + "a")))
            _ = fi16
            _ = v19
        }
        if (255 >= 1) {
            for fi20 := 0; fi20 < 16; fi20++ {
                p0 = ((-32768 > 1) && p1)
                _ = fi20
            }
        }
    }
    p1 = ((p1 && p1) || false)
    _ = p0
    _ = p1
    _ = v13
    return (-563 >= -1)
}

func main() {
    var v1 int32 = -158
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 float64 = (3.14159 - (1e10 / 0.1))
    var v5 uint32 = 1
    v5 += v5
    var v21 uint16 = 42
    switch "foo bar" {
    case "hello":
        var v22 float32 = (0.1 + 1.5)
        _ = v22
    }
    var v23 bool = (float32(v4) == (0.0 / 1.0))
    if (42 >= 1) {
        var v24 uint8 = 128
        for fi25 := 0; fi25 < 10; fi25++ {
            v24 = 163
            if (127 <= fi25) {
                var v26 uint32 = v5
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 bool = (!(0.0 == 0.0))
                    v24 -= (v24 | (v24 & (v24 << ((562) & 63))))
                    v29 := v3[fi25]
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            } else {
                v21 -= 1
                v5 += v5
                v30 := GenS3{F0: "foo bar", F1: v5, F2: 0}
                _ = v30
            }
            _ = fi25
        }
        _ = v24
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        v23 = (-738 >= 397)
        v23 = false
        _ = fi31
    }
    var v32 uint32 = 0
    switch "" {
    case "\n\t":
        v21 = v21
    case "a":
        v32 -= v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 int16 = 1000
        _ = fi33
        _ = v34
    }
    switch -743 {
    case 128:
        if v23 {
            fmt.Println(v21, v32)
        }
    case -2147483648:
        for fi35 := 0; fi35 < 18; fi35++ {
            v21 += (v21 - (v21 + 604))
            var v36 string = "a"
            _ = fi35
            _ = v36
        }
    }
    if true {
        v5 += v5
        fmt.Println(v23, v32)
        v37 := GenS3{F0: "", F1: v5, F2: (v21 & 384)}
        _ = v37
    } else {
        var v38 uint16 = (v21 + v21)
        var v39 int32 = 1
        _ = v38
        _ = v39
    }
    var v40 int16 = 92
    v4 = v4
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v5
}
