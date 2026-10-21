package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int8; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64; F2 int64 }
type GenS4 struct { F0 int16 }
func genFn0() string {
    if false {
        var v7 float32 = 1e10
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(fi8, v7, fi8)
                _ = fi9
            }
            var v10 uint64 = 18446744073709551615
            var v11 int32 = 1
            _ = fi8
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 0}
        _ = v12
        _ = v7
    } else {
        var v13 uint8 = 168
        v14 := GenS1{F0: 42}
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 42, F1: 42, F2: 124}
    v16 := GenS2{F0: 4294967295, F1: -1, F2: 100}
    _ = v15
    _ = v16
    return "a"
}
func genFn1(p0 int64, p1 bool, p2 uint32) uint {
    p2 = 1
    if (("hello" >= "hello") || p1) {
        var v17 uint64 = 2
        _ = v17
    } else {
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := []int{fi18, 4294967295}
            p1 = p1
            _ = fi18
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 int64 = 1
    var v2 string = ("foo bar" + (("" + "foo bar") + (("foo bar" + ("" + "a")) + "a")))
    v3 := []uint32{140, 1}
    v4 := []int{674}
    var v5 uint8 = 66
    v6 := make(map[string]string)
    v5 = v5
    _ = len(v4)
    fmt.Println(v1, v1)
    fmt.Println(v2)
    switch ("hello" + v2) {
    case "":
        var v20 uint16 = 997
        _ = v20
    case "hello":
        _ = len(v6)
    default:
        var v21 bool = true
        _ = v21
    }
    v22 := GenS1{F0: -1}
    switch v2 {
    case "":
        var v23 int16 = 1
        _ = v23
    case "\n\t":
        v1 = (v1 & v1)
    }
    var v24 int16 = 1000
    v25 := GenS3{F0: (0 > 2147483648), F1: 47, F2: v1}
    v2 = v2
    if (!(genFn1(v1, ((1.5 == (-2.25 - 0.1)) && (true || (false && false))), 42) >= 721)) {
        switch v2 {
        case "a":
            v1 = 42
        case "\n\t":
            v6["hello"] = v2
        case "foo bar":
            v2 = v2
        }
        v26 := GenS4{F0: v24}
        _ = v26
    } else {
        v4 = append(v4, -115)
        for fi27 := 0; fi27 < 1; fi27++ {
            for fi28 := 0; fi28 < 18; fi28++ {
                switch fi28 {
                case 128:
                    var v29 uint64 = 42
                    _ = v29
                }
                if (42 != 288) {
                    fmt.Println(fi28, fi28)
                    switch -2147483648 {
                    case -389:
                        var v30 float64 = float64(0.1)
                        _ = v30
                    case 1:
                        v31 := GenS3{F0: (!false), F1: 18446744073709551615, F2: -9223372036854775808}
                        _ = v31
                    case 42:
                        fmt.Println(v2)
                    default:
                        if true {
                            var v32 uint32 = 297
                            var v33 float64 = 1e10
                            _ = v32
                            _ = v33
                        } else {
                            v24 -= (-(v24))
                            v34 := v6[("hello" + ((("hello" + v2) + (v2 + "")) + ("foo bar" + ("" + ""))))]
                            v35 := GenS2{F0: 2147483648, F1: 42, F2: 1}
                            _ = v34
                            _ = v35
                        }
                    }
                } else {
                    switch "\n\t" {
                    case "":
                        v2 = v2
                    case "a":
                        var v36 int = (fi27 - -632)
                        _ = v36
                    }
                }
                _ = fi28
            }
            _ = fi27
        }
        var v37 int16 = (-(v24))
        _ = v37
    }
    if true {
        var v38 float64 = ((-2.25 / float64((1e10 * 0.1))) / 1.0)
        fmt.Println(v38, v1, v1)
        _ = v38
    }
    var v39 uint8 = (v5 & 0)
    v5 -= (v39 % ((v5) | 1))
    var v40 uint16 = 684
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
