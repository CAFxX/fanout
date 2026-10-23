package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS3 struct { F0 string; F1 int8; F2 uint16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint64, p1 string, p2 int16) int8 {
    p0 += p0
    if true {
        var v5 float64 = 1e10
        if false {
            fmt.Println(v5, v5, v5)
            var v6 int32 = 2147483647
            _ = v6
        }
        if (!((float32(v5) + 1e10) > (0.1 + ((0.1 + float32(1e10)) + (1e10 + (-2.25 / -2.25)))))) {
            if false {
                fmt.Println(p0, p1, p2)
                v7 := []int{32768, 823, -1}
                v8 := GenS1{F0: p1, F1: false, F2: 572}
                _ = v7
                _ = v8
            } else {
                if false {
                    fmt.Println(p1, v5, p2)
                    var v9 float64 = (0.1 - v5)
                    _ = v9
                }
                if false {
                    if (!false) {
                        switch 0 {
                        case -874:
                            var v10 int = 290
                            _ = v10
                        case -3:
                            v11 := GenS2{F0: 128, F1: -1, F2: 1}
                            _ = v11
                        default:
                            v12 := []int{-9223372036854775808, -434}
                            _ = v12
                        }
                    }
                } else {
                    var v13 int64 = -423
                    _ = v13
                }
            }
        }
        _ = v5
    } else {
        switch ((((((p1 + ((p1 + (p1 + (p1 + "hello"))) + (p1 + (p1 + "")))) + "foo bar") + "") + p1) + p1) + ((((("a" + (p1 + ((p1 + (((p1 + p1) + p1) + (p1 + (("hello" + "a") + ("foo bar" + "\n\t"))))) + ""))) + p1) + p1) + ((("hello" + p1) + p1) + ((p1 + "foo bar") + "hello"))) + "")) {
        case "a":
            fmt.Println(p0, p0)
        case "hello":
            var v14 int16 = p2
            _ = v14
        }
        fmt.Println(p2, p2)
        var v15 uint64 = 570
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        fmt.Println(p0, p0, p1)
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := GenS3{F0: (p1 + p1), F1: 78, F2: 1}
            v19 := GenS4{F0: 9223372036854775807, F1: 100}
            var v20 int64 = -779
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    return -8
}

func main() {
    var v1 string = "foo bar"
    var v2 int16 = 510
    var v3 bool = (-51 > -9)
    v4 := make(map[int]bool)
    v3 = false
    if true {
        fmt.Println(v1, v2, v2)
    } else {
        delete(v4, -2147483648)
    }
    switch "hello" {
    case "hello":
        if (32768 > 2147483647) {
            v3 = v3
        }
    }
    v2 += v2
    v21 := GenS5{F0: -128}
    if v3 {
        _ = len(v4)
    } else {
        v3 = (("" + ((v1 + v1) + v1)) >= v1)
    }
    v3 = ((!v3) && (v3 || v3))
    v22 := GenS1{F0: "", F1: ((v3 && v3) && (v2 == v2)), F2: 952}
    v23 := GenS5{F0: genFn0(1, v1, v2)}
    var v24 int64 = 9223372036854775807
    for fi25 := 0; fi25 < 6; fi25++ {
        _ = len(v4)
        var v26 float64 = -2.25
        _ = fi25
        _ = v26
    }
    var v27 uint16 = 2
    v2 += v2
    v1 = ("" + "\n\t")
    fmt.Println(v3, v2, v27)
    if (922 == 4294967295) {
        if v3 {
            var v28 int32 = -288
            v29 := GenS4{F0: -9223372036854775808, F1: 125}
            _ = v28
            _ = v29
        } else {
            for fi30 := 0; fi30 < 12; fi30++ {
                fmt.Println(v3)
                _ = fi30
            }
        }
        var v31 int = 9223372036854775807
        v31 = 568
        _ = v31
    } else {
        fmt.Println(v27, v27)
        var v32 uint8 = 2
        v4[0] = (v3 || true)
        _ = v32
    }
    v33 := []int{42, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v33
    _ = v4
}
