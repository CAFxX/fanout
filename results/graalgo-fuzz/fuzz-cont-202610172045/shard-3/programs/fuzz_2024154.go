package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint; F1 int; F2 int32 }
type GenS5 struct { F0 uint16; F1 int32; F2 int16 }
type GenS6 struct { F0 uint32; F1 int8 }
func genFn0(p0 int32, p1 int32, p2 int32) bool {
    v5 := GenS1{F0: 42, F1: 42, F2: 100}
    v6 := GenS2{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (!(226 >= -9223372036854775808))
}
func genFn1(p0 int64) bool {
    var v7 uint = 2
    p0 -= (-(p0))
    if (-2.25 <= -2.25) {
        var v8 int16 = 32767
        var v9 uint8 = 162
        _ = v8
        _ = v9
    }
    v10 := GenS3{F0: p0, F1: 2147483647, F2: 258}
    _ = p0
    _ = v10
    _ = v7
    return false
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 255
    v3 := []string{"a", (("" + ("" + "\n\t")) + ("" + "a"))}
    var v4 string = (("" + "") + ("" + ""))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v11 uint16 = 65535
    var v12 float64 = 0.1
    v2 -= v2
    v12 -= v12
    if false {
        v13 := GenS3{F0: 0, F1: -1, F2: 1}
        _ = v13
    }
    if false {
        var v14 uint64 = 4294967295
        v11 = (v11 | 995)
        _ = v14
    } else {
        if true {
            var v15 int = 1
            var v16 int32 = -75
            v2 += v2
            _ = v15
            _ = v16
        }
        v1[""] = (127 == 100)
    }
    v1[("a" + "a")] = ((!(!false)) && false)
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 bool = (42 >= 1)
        v2 = v2
        var v19 int32 = -768
        _ = fi17
        _ = v18
        _ = v19
    }
    if false {
        var v20 bool = true
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := v1[("foo bar" + "hello")]
            var v23 int = 42
            _ = fi21
            _ = v22
            _ = v23
        }
        fmt.Println(v12, v12)
        _ = v20
    }
    var v24 uint16 = (v11 - 32767)
    switch -923 {
    case 719:
        for fi25 := 0; fi25 < 12; fi25++ {
            v26 := GenS1{F0: 2147483648, F1: 127, F2: 100}
            v27 := GenS4{F0: (v2 >> ((v2) & 63)), F1: (fi25 | fi25), F2: 0}
            _ = fi25
            _ = v26
            _ = v27
        }
    case 32768:
        if (330 != 334) {
            v4 = ("" + (v4 + v4))
            if true {
                for fi28 := 0; fi28 < 15; fi28++ {
                    fmt.Println(v24, fi28, v2)
                    v29 := GenS5{F0: v24, F1: -1, F2: 1}
                    _ = fi28
                    _ = v29
                }
            } else {
                if true {
                    switch 1000 {
                    case -1:
                        v2 += (v2 % ((v2) | 1))
                    case -9223372036854775808:
                        if true {
                            _ = len(v1)
                            fmt.Println(v2, v12, v4)
                            v30 := GenS6{F0: 2147483647, F1: -23}
                            _ = v30
                        } else {
                            fmt.Println(v11)
                            for fi31 := 0; fi31 < 8; fi31++ {
                                v4 = "hello"
                                _ = fi31
                            }
                        }
                    case 1:
                        v11 = 255
                    }
                } else {
                    var v32 uint64 = 821
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v34 := GenS5{F0: (v11 % (((v24 >> ((4294967295) & 63))) | 1)), F1: -1, F2: 42}
                        var v35 uint8 = 127
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    _ = v32
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    case -2147483648:
        v36 := GenS4{F0: v2, F1: -2147483648, F2: 42}
        _ = v36
    }
    var v37 int8 = -16
    var v38 uint16 = v24
    var v39 uint = 2
    var v40 float32 = 0.0
    var v41 uint = (^(v2))
    var v42 bool = (1.5 >= float32(v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
