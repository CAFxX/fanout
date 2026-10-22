package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint32; F1 string; F2 uint }
type GenS5 struct { F0 uint; F1 bool }
func genFn0() int32 {
    v7 := []int{1, 1, 32767, -892}
    switch "\n\t" {
    case "a":
        var v8 int16 = 0
        _ = v8
    case "\n\t":
        var v9 float32 = 1.5
        _ = v9
    }
    var v10 float32 = -2.25
    switch "foo bar" {
    case "foo bar":
        v10 -= v10
    default:
        var v11 int16 = 42
        _ = v11
    }
    _ = v10
    _ = v7
    return -291
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    var v4 uint = 97
    var v5 bool = false
    v6 := make(map[int]uint32)
    fmt.Println(v4, v4)
    v12 := GenS1{F0: 127, F1: 166, F2: 4294967295}
    for fi13 := 0; fi13 < 20; fi13++ {
        v4 = v4
        var v14 float64 = ((float64(float32(((-2.25 + 1.5) * 0.1))) + -2.25) + 0.1)
        _ = fi13
        _ = v14
    }
    if false {
        v2["a"] = ("" + "foo bar")
    } else {
        v15 := GenS2{F0: (^(v4))}
        v16 := GenS1{F0: 2, F1: -753, F2: 4294967295}
        _ = v15
        _ = v16
    }
    v17 := []int{32768}
    v18 := GenS1{F0: uint8(v4), F1: -408, F2: 18446744073709551615}
    switch 32768 {
    case 4294967295:
        for fi19 := 0; fi19 < 16; fi19++ {
            v4 -= 295
            v5 = (1 < 0)
            _ = fi19
        }
    }
    v5 = false
    v20 := GenS3{F0: 128, F1: (42 < 227)}
    if false {
        switch "\n\t" {
        case "foo bar":
            var v21 int64 = 0
            _ = v21
        case "hello":
            var v22 int32 = genFn0()
            _ = v22
        }
        if (-150 >= 1000) {
            v5 = v5
            fmt.Println(v5)
        } else {
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(fi23)
                if (-684 == 1) {
                    for fi24 := 0; fi24 < 9; fi24++ {
                        v25 := GenS1{F0: 1, F1: fi23, F2: 1}
                        v26 := GenS3{F0: 1, F1: (true || false)}
                        v27 := GenS4{F0: 2147483648, F1: "foo bar", F2: (v4 << (((v4 | v4)) & 63))}
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                }
                for fi28 := 0; fi28 < 18; fi28++ {
                    v5 = (v5 || (!(499 != 570)))
                    if (v5 || (42 != 255)) {
                        if (100 > 127) {
                            v4 += v4
                            v4 = (v4 | 936)
                            v29 := GenS3{F0: 45, F1: (v5 && v5)}
                            _ = v29
                        }
                        v5 = (!((fi28 | fi23) < 2147483647))
                        fmt.Println(fi28, fi28)
                    }
                    if (127 > 212) {
                        if (v5 || (!true)) {
                            var v30 uint16 = 554
                            _ = v30
                        }
                        v31 := GenS4{F0: 42, F1: "foo bar", F2: v4}
                        v4 = 680
                        var v32 int = (fi23 - (-(fi28)))
                        var v33 float32 = 1e10
                        fmt.Println(v33, v32, v4)
                        if v5 {
                            if ((!(!(9223372036854775807 != 42))) || v5) {
                                switch -2147483648 {
                                case 128:
                                    v34 := GenS2{F0: v4}
                                    _ = v34
                                default:
                                    if true {
                                        v35 := GenS1{F0: 42, F1: fi28, F2: 596}
                                        var v36 int = fi23
                                        var v37 uint64 = uint64(v4)
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    } else {
                                        var v38 int32 = 42
                                        _ = v38
                                    }
                                }
                            }
                            var v39 uint16 = 32767
                            _ = v39
                        } else {
                            var v40 int32 = 0
                            _ = v40
                        }
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    var v41 uint8 = 1
                    _ = fi28
                    _ = v41
                }
                _ = fi23
            }
            var v42 int32 = 510
            _ = v42
        }
    } else {
        v43 := GenS4{F0: 2, F1: "hello", F2: v4}
        v44 := GenS5{F0: v4, F1: ("" > ("foo bar" + (("\n\t" + "") + ((("" + ("a" + "a")) + ("hello" + (("a" + "foo bar") + "foo bar"))) + "hello"))))}
        fmt.Println(v5, v5)
        _ = v43
        _ = v44
    }
    v4 += 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
