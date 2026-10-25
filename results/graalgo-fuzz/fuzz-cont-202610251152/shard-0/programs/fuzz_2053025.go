package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 int }
type GenS3 struct { F0 uint32; F1 int; F2 string }
type GenS4 struct { F0 uint32; F1 int32 }
func genFn0() int {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int32 = 629
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 283
    switch ((("\n\t" + "hello") + "foo bar") + (("a" + "a") + "foo bar")) {
    case "\n\t":
        var v9 float32 = 3.14159
        _ = v9
    default:
        v8 = 636
    }
    fmt.Println(v8)
    _ = v8
    return 42
}
func genFn1() uint8 {
    v10 := GenS1{F0: 0, F1: 255}
    var v11 int = genFn0()
    v11 = (-(v11))
    _ = v10
    _ = v11
    return 127
}

func main() {
    var v1 uint64 = 596
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 bool = (!((!true) && (!true)))
    var v5 uint8 = uint8(v1)
    switch "a" {
    case "":
        var v12 uint32 = 990
        _ = v12
    case "hello":
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := []int{(fi13 / (((^(fi13))) | 1)), -1}
            var v15 float32 = float32(0.1)
            _ = fi13
            _ = v14
            _ = v15
        }
    case "foo bar":
        v16 := GenS1{F0: (v1 / ((0) | 1)), F1: 65535}
        _ = v16
    default:
        v17 := GenS1{F0: v1, F1: 1}
        _ = v17
    }
    v18 := []int{9223372036854775807, 4294967295, 4294967295}
    v4 = v4
    v5 += v5
    v4 = true
    if v4 {
        v5 += (v5 * (^(v5)))
    } else {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 18; fi20++ {
                _ = len(v18)
                v1 = v1
                switch fi19 {
                case 127:
                    fmt.Println(v1, fi19)
                }
                if v4 {
                    _ = len(v3)
                    fmt.Println(v5)
                    var v21 uint32 = 658
                    _ = v21
                }
                _ = fi20
            }
            if true {
                var v22 int8 = -128
                if (-629 != 0) {
                    if len(v18) > 0 {
                        v18[0] = (-(fi19))
                    }
                    var v23 int64 = -1
                    v4 = v4
                    _ = v23
                } else {
                    v2[genFn0()] = v4
                    if v4 {
                        v5 = genFn1()
                        v5 -= v5
                        var v24 int8 = -118
                        _ = v24
                    }
                }
                _ = v22
            }
            _ = fi19
        }
        v25 := v3[596]
        _ = v25
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v26 int64 = 9223372036854775807
    fmt.Println(v26)
    switch 105 {
    case 883:
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    default:
        if v4 {
            var v27 uint32 = 0
            v26 += 9223372036854775807
            switch "foo bar" {
            case "":
                var v28 uint8 = v5
                _ = v28
            case "a":
                v3[316] = ((("a" + "") + ((("hello" + "\n\t") + "a") + "\n\t")) + "hello")
            }
            _ = v27
        }
    }
    fmt.Println(v4)
    v29 := GenS1{F0: 1, F1: 633}
    var v30 uint16 = 255
    if (1e10 <= (1.5 * 0.0)) {
        v26 -= -9223372036854775808
        v31 := GenS2{F0: v30, F1: v5, F2: -982}
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = v31
    } else {
        if (true && true) {
            v32 := GenS3{F0: 4294967295, F1: 42, F2: "hello"}
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            v1 = v1
            _ = v32
        } else {
            fmt.Println(v26, v4)
            delete(v2, -174)
            switch "foo bar" {
            case "hello":
                var v33 string = ""
                _ = v33
            case "a":
                v34 := GenS4{F0: 2, F1: 0}
                _ = v34
            }
        }
        fmt.Println(v26, v1, v4)
    }
    v26 -= 859
    v26 += (v26 / (((v26 + (v26 ^ -1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
