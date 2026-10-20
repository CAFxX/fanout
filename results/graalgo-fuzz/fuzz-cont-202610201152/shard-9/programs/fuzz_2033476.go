package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0() int16 {
    v5 := []int{}
    var v6 uint16 = 255
    var v7 uint32 = 4294967295
    if (!(74 == 126)) {
        var v8 int32 = 61
        _ = v8
    }
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int8, p1 int) int64 {
    if true {
        switch p1 {
        case 153:
            if ((0.0 == -2.25) && false) {
                if true {
                    var v9 string = (((("" + "foo bar") + "\n\t") + "\n\t") + "")
                    fmt.Println(v9, p0)
                    var v10 uint16 = 224
                    _ = v10
                    _ = v9
                }
            } else {
                if false {
                    p1 -= (p1 * 4294967295)
                    fmt.Println(p0)
                    if (false || (32767 <= 42)) {
                        v11 := []int{(^(p1)), -549, p1, 4294967295}
                        _ = v11
                    } else {
                        v12 := []int{(p1 % (((p1 + p1)) | 1))}
                        if len(v12) > 0 {
                            v12[0] = p1
                        }
                        _ = v12
                    }
                } else {
                    for fi13 := 0; fi13 < 4; fi13++ {
                        var v14 uint16 = 255
                        _ = fi13
                        _ = v14
                    }
                }
                var v15 int32 = -593
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := GenS1{F0: v15}
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
        }
        v18 := GenS1{F0: 0}
        switch (("a" + "a") + "") {
        case "hello":
            v19 := GenS2{F0: 282, F1: 42}
            _ = v19
        case "foo bar":
            fmt.Println(p1, p0)
        case "a":
            fmt.Println(p1)
        default:
            if (!((!true) && false)) {
                v20 := []int{9223372036854775807, (p1 / (((p1 | p1)) | 1))}
                _ = v20
            }
        }
        _ = v18
    }
    var v21 uint16 = 326
    var v22 uint8 = 127
    v23 := []int{p1, p1, 1, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v23
    return 1
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint32 = 4294967295
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    if false {
        if (-128 <= -52) {
            for fi24 := 0; fi24 < 17; fi24++ {
                fmt.Println(v1, v2)
                _ = fi24
            }
        } else {
            if false {
                v1 = v1
                v1 -= v1
                v1 += 0
            }
            fmt.Println(v1, v1)
            if (!(2 > 18446744073709551615)) {
                var v25 string = "foo bar"
                v26 := []int{59, v1, (v1 ^ 259)}
                _ = v25
                _ = v26
            }
        }
        var v27 bool = false
        var v28 uint64 = 1
        _ = v27
        _ = v28
    } else {
        fmt.Println(v2, v2)
        switch "hello" {
        case "hello":
            var v29 uint8 = 1
            _ = v29
        case "a":
            if (!("" == ("a" + "hello"))) {
                v1 += 129
                fmt.Println(v1, v2)
            }
        }
    }
    fmt.Println(v2)
    var v30 uint64 = 667
    v1 += (v1 % ((v1) | 1))
    v31 := GenS1{F0: 2147483647}
    v32 := GenS1{F0: -1}
    v33 := []int{}
    v30 += (^(v30))
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            switch "hello" {
            case "foo bar":
                fmt.Println(fi35, v30, v30)
            case "a":
                fmt.Println(v1, fi34, fi34)
            case "hello":
                fmt.Println(v1, v30, v2)
            }
            delete(v3, v1)
            _ = fi35
        }
        delete(v4, "")
        _ = len(v3)
        _ = fi34
    }
    v1 -= (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
