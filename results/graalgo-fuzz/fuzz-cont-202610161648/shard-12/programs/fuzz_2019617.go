package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int32; F2 string }
func genFn0(p0 uint32, p1 int32) uint16 {
    if (false && (!false)) {
        var v5 int64 = -1
        v6 := []int{42, -762}
        var v7 bool = (false && (2 >= 2))
        _ = v5
        _ = v6
        _ = v7
    } else {
        p1 -= (p1 ^ p1)
        fmt.Println(p0, p0)
    }
    p0 = (p0 >> ((65535) & 63))
    v8 := []int{-703}
    p1 -= -58
    var v9 string = ("hello" + ("foo bar" + (((("hello" + "\n\t") + "foo bar") + ("foo bar" + "a")) + (("a" + "hello") + ("a" + (("foo bar" + "\n\t") + ("hello" + "a")))))))
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 492
}
func genFn1(p0 uint16, p1 int, p2 string) int {
    switch ("hello" + p2) {
    case "\n\t":
        switch p1 {
        case -778:
            var v10 int64 = -872
            _ = v10
        case -347:
            p2 = (p2 + p2)
        }
    case "":
        var v11 float64 = (1e10 / -2.25)
        _ = v11
    case "a":
        v12 := []int{}
        _ = v12
    }
    p2 = "hello"
    var v13 uint16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -1
}

func main() {
    v1 := []string{((("foo bar" + "") + (("a" + ("hello" + "foo bar")) + "hello")) + ("foo bar" + (("\n\t" + ("" + "\n\t")) + ("foo bar" + "foo bar"))))}
    var v2 bool = true
    var v3 int16 = -1
    var v4 int16 = 0
    if v2 {
        v2 = true
        fmt.Println(v4, v4)
        switch (("\n\t" + "a") + ((("foo bar" + "a") + ("foo bar" + (("foo bar" + (("hello" + ("foo bar" + "")) + "foo bar")) + "\n\t"))) + "foo bar")) {
        case "\n\t":
            var v14 int32 = -469
            _ = v14
        case "foo bar":
            if (42 == 0) {
                var v15 int64 = -9223372036854775808
                if (v2 && (v15 != (^(v15)))) {
                    fmt.Println(v15, v15)
                    switch -956 {
                    case -9223372036854775808:
                        if true {
                            v15 -= (v15 ^ (v15 ^ (-(v15))))
                        } else {
                            var v16 int32 = 1
                            _ = v16
                        }
                    default:
                        var v17 uint8 = 127
                        _ = v17
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    v2 = false
                }
                _ = v15
            }
        default:
            for fi18 := 0; fi18 < 18; fi18++ {
                v19 := GenS1{F0: 121}
                _ = fi18
                _ = v19
            }
        }
    } else {
        v1 = append(v1, "")
    }
    var v20 string = ""
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 bool = true
        var v23 uint8 = 74
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 uint = 420
    if (-2.25 == -2.25) {
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 uint8 = 127
            var v27 uint64 = 2
            for fi28 := 0; fi28 < 4; fi28++ {
                v29 := GenS2{F0: 1, F1: -394, F2: "foo bar"}
                _ = len(v1)
                fmt.Println(v24, v24, v20)
                _ = fi28
                _ = v29
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        fmt.Println(v24)
    }
    var v30 uint = 4294967295
    var v31 int32 = 1
    var v32 uint16 = 1
    switch ((("a" + ("foo bar" + ("\n\t" + "hello"))) + v20) + v20) {
    case "a":
        var v33 float64 = 1.5
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
