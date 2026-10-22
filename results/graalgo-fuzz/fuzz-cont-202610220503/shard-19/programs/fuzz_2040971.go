package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint8; F2 int }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int16, p1 uint32, p2 int32) bool {
    p0 = (-(p0))
    if (!false) {
        fmt.Println(p2)
        switch "foo bar" {
        case "\n\t":
            p2 = 146
        case "hello":
            var v4 int = 42
            _ = v4
        }
    } else {
        v5 := GenS1{F0: 9223372036854775807, F1: (p1 << ((683) & 63))}
        if (((1.5 + (3.14159 - 1e10)) - (1.5 + float64(1.5))) >= -2.25) {
            v6 := []int{728, -9223372036854775808, 859, 0}
            var v7 int64 = 0
            _ = v6
            _ = v7
        } else {
            p1 = p1
            fmt.Println(p0, p2, p0)
        }
        _ = v5
    }
    v8 := GenS2{F0: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return false
}
func genFn1() bool {
    switch ("" + (("\n\t" + "") + "foo bar")) {
    case "\n\t":
        v9 := GenS1{F0: -1, F1: 2}
        _ = v9
    case "a":
        var v10 float32 = 1e10
        _ = v10
    default:
        var v11 bool = (42 <= 127)
        _ = v11
    }
    var v12 string = "foo bar"
    switch (v12 + "\n\t") {
    case "\n\t":
        if (!(!true)) {
            for fi13 := 0; fi13 < 4; fi13++ {
                v12 = v12
                for fi14 := 0; fi14 < 19; fi14++ {
                    if (true && false) {
                        v15 := []int{}
                        switch (fi14 | (fi13 << ((900) & 63))) {
                        case 127:
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        case -422:
                            if false {
                                var v16 bool = false
                                if (v16 || (247 >= 18446744073709551615)) {
                                    var v17 int16 = 0
                                    var v18 uint16 = 601
                                    _ = v17
                                    _ = v18
                                } else {
                                    if len(v15) > 0 {
                                        v15[0] = fi13
                                    }
                                }
                                switch fi14 {
                                case 4294967295:
                                    v12 = ""
                                case 128:
                                    v16 = (v16 || v16)
                                default:
                                    switch fi14 {
                                    case 32767:
                                        var v19 string = v12
                                        _ = v19
                                    case 1:
                                        var v20 uint8 = 255
                                        _ = v20
                                    default:
                                        var v21 uint = 42
                                        _ = v21
                                    }
                                }
                                v22 := GenS3{F0: 42, F1: 405, F2: 2}
                                _ = v16
                                _ = v22
                            }
                        case 2147483647:
                            fmt.Println(fi13, v12)
                        default:
                            var v23 float32 = float32(1e10)
                            _ = v23
                        }
                        _ = v15
                    }
                    _ = fi14
                }
                v24 := GenS4{F0: 548, F1: 127, F2: (fi13 - fi13)}
                _ = fi13
                _ = v24
            }
        }
    default:
        switch v12 {
        case "":
            v12 = (v12 + v12)
        default:
            v12 = v12
        }
    }
    _ = v12
    return true
}

func main() {
    v1 := []uint32{}
    var v2 string = ("\n\t" + "hello")
    v3 := make(map[string]uint32)
    v2 = v2
    v25 := GenS5{F0: 65535}
    if (-245 <= -593) {
        var v26 int = 469
        v2 = v2
        _ = v26
    } else {
        _ = len(v1)
    }
    switch -9223372036854775808 {
    case 0:
        var v27 string = "\n\t"
        _ = v27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := GenS4{F0: -2147483648, F1: 1, F2: 220}
        var v30 int16 = 1
        v30 = 0
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        v2 = ((("hello" + v2) + ("\n\t" + v2)) + "hello")
        _ = fi31
    }
    if (42 != -662) {
        if true {
            var v32 int64 = -9223372036854775808
            v2 = v2
            _ = v32
        } else {
            for fi33 := 0; fi33 < 20; fi33++ {
                v2 = v2
                for fi34 := 0; fi34 < 17; fi34++ {
                    v2 = ((("\n\t" + "hello") + "") + "a")
                    v2 = ("" + v2)
                    v2 = (("foo bar" + "hello") + v2)
                    _ = fi34
                }
                _ = fi33
            }
        }
        v3[("" + v2)] = 2147483647
        v2 = v2
    }
    v35 := GenS2{F0: 42}
    switch v2 {
    case "a":
        var v36 uint = 915
        _ = v36
    case "":
        v2 = "a"
    default:
        v37 := GenS5{F0: 2147483648}
        _ = v37
    }
    var v38 uint64 = 676
    switch -1 {
    case 128:
        fmt.Println(v2, v38)
    }
    v3["a"] = 0
    var v39 uint16 = 32767
    var v40 uint = 0
    fmt.Println(v40)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v40
}
