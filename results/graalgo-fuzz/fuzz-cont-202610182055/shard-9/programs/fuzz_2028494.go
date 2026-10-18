package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 bool; F1 int16; F2 int16 }
func genFn0(p0 bool) string {
    if p0 {
        switch -782 {
        case 0:
            var v6 uint32 = 65535
            _ = v6
        case 991:
            if p0 {
                for fi7 := 0; fi7 < 5; fi7++ {
                    var v8 bool = (!(!(2 >= 18446744073709551615)))
                    for fi9 := 0; fi9 < 10; fi9++ {
                        v8 = true
                        var v10 int16 = 42
                        v11 := GenS1{F0: 4294967295}
                        v8 = (3.14159 != 1.5)
                        p0 = ((!v8) && v8)
                        v12 := GenS1{F0: 2}
                        if true {
                            p0 = ((3.14159 + (1.5 * 0.0)) != 1e10)
                            fmt.Println(fi9, fi9)
                        } else {
                            v13 := GenS1{F0: 259}
                            _ = v13
                        }
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    _ = fi7
                    _ = v8
                }
                p0 = (-640 != 0)
                v14 := []int{-934, -388, -1}
                _ = v14
            }
        default:
            if p0 {
                p0 = (1 == 4294967295)
            }
        }
        v15 := []int{-1}
        _ = v15
    } else {
        fmt.Println(p0)
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        if p0 {
            var v17 bool = (false || true)
            v18 := []int{}
            fmt.Println(p0)
            _ = v17
            _ = v18
        } else {
            v19 := []int{-752}
            _ = v19
        }
        _ = fi16
    }
    if (!(255 >= 255)) {
        var v20 float32 = ((0.1 * 1e10) * (0.0 - (((-2.25 / 1.5) / 1e10) / 0.1)))
        var v21 bool = (1 <= -128)
        _ = v20
        _ = v21
    } else {
        v22 := []int{}
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        _ = v22
    }
    fmt.Println(p0)
    _ = p0
    return (("hello" + "foo bar") + "hello")
}
func genFn1(p0 uint8, p1 int8, p2 uint) int {
    var v23 int16 = 1000
    switch (("foo bar" + "foo bar") + ("hello" + "foo bar")) {
    case "foo bar":
        var v24 bool = (!true)
        _ = v24
    case "\n\t":
        v25 := []int{32767}
        _ = v25
    }
    var v26 int64 = -356
    p0 += (p0 << ((35) & 63))
    v27 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v26
    _ = v27
    return -1
}

func main() {
    v1 := []string{}
    v2 := make(map[int]int)
    v3 := []int{}
    var v4 int32 = -45
    var v5 uint = 255
    for fi28 := 0; fi28 < 9; fi28++ {
        v29 := GenS2{F0: (fi28 - (fi28 / ((fi28) | 1))), F1: false}
        switch "" {
        case "\n\t":
            v5 += v5
        case "foo bar":
            v5 += (v5 * 4294967295)
        }
        _ = fi28
        _ = v29
    }
    v5 -= 4294967295
    v30 := GenS3{F0: 1}
    for fi31 := 0; fi31 < 1; fi31++ {
        if len(v1) > 0 {
            v1[0] = (("a" + ("foo bar" + "")) + (((("a" + "\n\t") + ("foo bar" + "hello")) + "hello") + "hello"))
        }
        var v32 string = "\n\t"
        if false {
            if len(v3) > 0 {
                v3[0] = genFn1(3, 16, v5)
            }
            v4 += v4
        }
        _ = fi31
        _ = v32
    }
    _ = len(v2)
    if (-1 != 1) {
        fmt.Println(v4, v4)
        v4 = 1
    } else {
        v5 += (v5 | (^(v5)))
    }
    _ = len(v3)
    var v33 int = -2147483648
    if false {
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 float64 = -2.25
            fmt.Println(v4, v33, v35)
            _ = fi34
            _ = v35
        }
        if ((1.5 + -2.25) >= float64(-2.25)) {
            v36 := GenS1{F0: 2}
            _ = v36
        }
        v37 := GenS4{F0: 42}
        _ = v37
    }
    if (4 >= 1) {
        v5 -= (v5 | 42)
    }
    v5 = 2
    v38 := GenS5{F0: 65535, F1: "hello"}
    v4 = -2147483648
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v33 = v33
    for fi39 := 0; fi39 < 17; fi39++ {
        v40 := GenS6{F0: ((0 < 1) && true), F1: -1, F2: -325}
        var v41 int = 42
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v38
    _ = v4
    _ = v5
}
