package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 uint, p1 int, p2 bool) int64 {
    v7 := GenS1{F0: p0}
    if p2 {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := []int{}
            _ = fi8
            _ = v9
        }
        fmt.Println(p1, p2, p0)
        var v10 float32 = 3.14159
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int16, p1 int8) int {
    fmt.Println(p0, p0)
    v11 := []int{32767, 2147483647}
    var v12 int64 = genFn0(293, 1, false)
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint = 42
        var v15 uint64 = 482
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 712
}

func main() {
    v1 := []int{127, -754, 2147483647}
    var v2 float32 = float32(0.1)
    var v3 uint64 = 2
    v4 := []string{"a", "\n\t", (("foo bar" + "") + "hello"), ("a" + ("hello" + "foo bar"))}
    v5 := []int{128}
    var v6 uint = 4294967295
    fmt.Println(v6)
    var v16 uint64 = 42
    if (uint64(v6) < (^(v3))) {
        switch "" {
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            fmt.Println(v3, v16, v3)
        }
        var v17 uint8 = 150
        _ = v17
    } else {
        var v18 string = "hello"
        var v19 int32 = 1
        v6 = (v6 << ((v6) & 63))
        _ = v18
        _ = v19
    }
    switch 127 {
    case -997:
        v6 -= (^(v6))
    }
    var v20 string = "hello"
    var v21 int32 = -844
    if true {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 bool = false
            if len(v5) > 0 {
                v5[0] = 32768
            }
            _ = fi22
            _ = v23
        }
        _ = len(v4)
    } else {
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 int8 = -128
            v25 = -15
            _ = fi24
            _ = v25
        }
        var v26 uint = 0
        _ = v26
    }
    v2 = 3.14159
    v21 += v21
    switch 128 {
    case 9223372036854775807:
        var v27 float32 = 3.14159
        _ = v27
    case 0:
        if true {
            for fi28 := 0; fi28 < 8; fi28++ {
                if len(v1) > 0 {
                    v1[0] = 941
                }
                if false {
                    if len(v1) > 0 {
                        v1[0] = genFn1(1000, -118)
                    }
                    var v29 uint32 = 144
                    _ = v29
                }
                _ = fi28
            }
            v6 += (^(v6))
        }
    case 42:
        v20 = v20
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v2 = ((0.0 / v2) * v2)
        fmt.Println(fi30, v21, v21)
        _ = fi30
    }
    if len(v5) > 0 {
        v5[0] = -949
    }
    var v31 uint32 = 109
    for fi32 := 0; fi32 < 9; fi32++ {
        v21 += (v21 * 1)
        fmt.Println(v21)
        _ = fi32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        switch ("hello" + "\n\t") {
        case "\n\t":
            v20 = ("foo bar" + "hello")
        case "":
            if ((-2.25 < 1e10) || (uint(v16) == 18446744073709551615)) {
                switch (v20 + ("" + v20)) {
                case "":
                    v3 += v3
                case "foo bar":
                    v34 := GenS2{F0: -9223372036854775808, F1: -32768}
                    _ = v34
                default:
                    fmt.Println(v6, v21, fi33)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 -= v2
            } else {
                var v35 uint8 = 56
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(v6)
                _ = v35
            }
        }
        _ = fi33
    }
    var v36 float32 = v2
    var v37 int32 = v21
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
