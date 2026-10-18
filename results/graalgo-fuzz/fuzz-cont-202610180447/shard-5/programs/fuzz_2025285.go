package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 int16 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint16 }
func genFn0() int64 {
    v6 := GenS1{F0: 9223372036854775807, F1: false, F2: 83}
    if true {
        var v7 int64 = -697
        fmt.Println(v7)
        _ = v7
    }
    v8 := []int{}
    for fi9 := 0; fi9 < 5; fi9++ {
        if len(v8) > 0 {
            v8[0] = 2147483647
        }
        var v10 string = "a"
        var v11 uint8 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v6
    _ = v8
    return 16
}
func genFn1(p0 uint32, p1 int16, p2 uint16) int64 {
    p1 = p1
    var v12 uint8 = uint8(p2)
    for fi13 := 0; fi13 < 7; fi13++ {
        if (498 == 892) {
            v14 := []int{32767, fi13, (fi13 + fi13), 32768}
            _ = v14
        }
        v12 -= (v12 / ((v12) | 1))
        var v15 int = fi13
        _ = fi13
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        if false {
            var v17 uint16 = 1
            _ = v17
        }
        var v18 int16 = (p1 << ((0) & 63))
        if true {
            if (false && (4294967295 == 65535)) {
                if (1 < 192) {
                    v12 = 0
                    p0 = (^(p0))
                } else {
                    v19 := GenS1{F0: genFn0(), F1: (0.1 < (0.1 * (-2.25 + ((-2.25 - 0.0) - 0.1)))), F2: p2}
                    _ = v19
                }
                v12 -= v12
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                if ((^(p0)) == p0) {
                    p2 -= (p2 << ((337) & 63))
                    p1 = (p1 - 42)
                    var v21 int8 = 100
                    _ = v21
                } else {
                    v22 := []int{}
                    fmt.Println(p1, v12)
                    fmt.Println(fi20)
                    switch ("hello" + "a") {
                    case "hello":
                        switch "foo bar" {
                        case "a":
                            _ = len(v22)
                        case "\n\t":
                            switch fi16 {
                            case -581:
                                if true {
                                    v18 -= (^(v18))
                                    p2 = 0
                                }
                                fmt.Println(fi20)
                            default:
                                fmt.Println(v18)
                            }
                        }
                    case "a":
                        var v23 int8 = 42
                        _ = v23
                    }
                    _ = v22
                }
                fmt.Println(p0, p1)
                if ((1.5 * (-2.25 + (1.5 - (0.0 / float32(1e10))))) == (3.14159 - -2.25)) {
                    v24 := GenS2{F0: "\n\t", F1: p1, F2: -32768}
                    _ = v24
                } else {
                    var v25 int8 = -33
                    if true {
                        var v26 int16 = 1
                        _ = v26
                    }
                    _ = v25
                }
                _ = fi20
            }
        }
        _ = fi16
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 1
    v2 := []int{-9223372036854775808}
    var v3 int8 = -128
    v4 := make(map[string]uint32)
    var v5 int = -1
    var v27 int32 = int32(v3)
    v3 = 1
    v3 -= int8(v1)
    if true {
        if (!(genFn0() < -1)) {
            var v28 uint = 594
            _ = v28
        }
    } else {
        v29 := GenS2{F0: "foo bar", F1: 1000, F2: 32767}
        _ = v29
    }
    v5 += v5
    v30 := GenS3{F0: v5, F1: 1000}
    v4[(("a" + "a") + ("a" + ("" + "a")))] = 0
    for fi31 := 0; fi31 < 5; fi31++ {
        delete(v4, "foo bar")
        v32 := GenS4{F0: 0}
        v33 := GenS4{F0: 1}
        _ = fi31
        _ = v32
        _ = v33
    }
    switch (v5 % (((v5 * 32767)) | 1)) {
    case 938:
        for fi34 := 0; fi34 < 8; fi34++ {
            for fi35 := 0; fi35 < 14; fi35++ {
                fmt.Println(fi34, fi34, v1)
                var v36 bool = false
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    default:
        v3 += (v3 >> ((386) & 63))
    }
    if (949 <= 0) {
        var v37 float32 = 0.1
        var v38 uint64 = 42
        _ = v37
        _ = v38
    }
    v39 := GenS3{F0: (v5 % (((v5 << ((1) & 63))) | 1)), F1: -333}
    v3 -= v3
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v5
}
