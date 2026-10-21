package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint }
type GenS4 struct { F0 int64; F1 string; F2 bool }
type GenS5 struct { F0 uint64 }
func genFn0() int {
    v4 := []int{42, 218}
    var v5 bool = true
    for fi6 := 0; fi6 < 9; fi6++ {
        switch (fi6 - fi6) {
        case -9223372036854775808:
            if (!v5) {
                var v7 uint32 = 609
                var v8 int = (fi6 - fi6)
                _ = v7
                _ = v8
            }
        default:
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint = 68
                if v5 {
                    switch fi9 {
                    case 32768:
                        v4 = append(v4, fi6)
                    }
                    switch 588 {
                    case 42:
                        v10 -= v10
                    case 330:
                        if true {
                            if (!(!(88 >= -98))) {
                                v5 = ((!false) && (v5 || (-2147483648 < -755)))
                                v11 := GenS1{F0: 42, F1: 1, F2: 72}
                                fmt.Println(fi9, v5)
                                _ = v11
                            } else {
                                switch ("foo bar" + ("\n\t" + ("\n\t" + "hello"))) {
                                case "a":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "hello":
                                    var v12 uint64 = 349
                                    _ = v12
                                case "\n\t":
                                    var v13 int32 = 2147483647
                                    _ = v13
                                }
                                v14 := GenS2{F0: 0}
                                if (true || v5) {
                                    var v15 string = "a"
                                    _ = v15
                                }
                                var v16 uint = v10
                                v5 = v5
                                _ = v14
                                _ = v16
                            }
                            var v17 int8 = 100
                            fmt.Println(fi9, fi9)
                            _ = v17
                        } else {
                            v4 = append(v4, fi6)
                            v10 = v10
                        }
                    case 1:
                        if len(v4) > 0 {
                            v4[0] = -1
                        }
                    default:
                        v5 = v5
                    }
                    var v18 uint8 = 2
                    _ = v18
                }
                var v19 uint32 = 65535
                _ = fi9
                _ = v10
                _ = v19
            }
        }
        v4 = append(v4, (fi6 / (((fi6 & 32768)) | 1)))
        _ = fi6
    }
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 0
    v3 := make(map[int]bool)
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 uint16 = v2
        v21 = v2
        _ = fi20
        _ = v21
    }
    fmt.Println(v2)
    v22 := GenS2{F0: 42}
    if ((!(249 <= 42)) || (true || true)) {
        v23 := []int{9223372036854775807}
        v24 := GenS3{F0: 58, F1: -65, F2: 4294967295}
        if true {
            _ = len(v1)
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            var v25 float32 = 0.1
            _ = v25
        }
        _ = v23
        _ = v24
    } else {
        var v26 int8 = 42
        _ = v26
    }
    v27 := []int{-582, 683}
    v2 = 130
    v2 = (v2 & (v2 | v2))
    if false {
        v27 = append(v27, 1)
    } else {
        if true {
            v2 = v2
            fmt.Println(v2)
        }
    }
    delete(v3, 127)
    for fi28 := 0; fi28 < 8; fi28++ {
        v2 += v2
        _ = fi28
    }
    v29 := GenS2{F0: 0}
    if false {
        for fi30 := 0; fi30 < 14; fi30++ {
            v2 = v2
            for fi31 := 0; fi31 < 16; fi31++ {
                v2 = v2
                v2 = (v2 - v2)
                _ = fi31
            }
            _ = fi30
        }
        switch "hello" {
        case "a":
            var v32 int = -314
            _ = v32
        case "hello":
            var v33 int32 = -972
            _ = v33
        }
    }
    if false {
        v2 = 32767
        v34 := GenS1{F0: -1, F1: -1, F2: 11}
        _ = v34
    }
    var v35 float32 = -2.25
    var v36 uint64 = 266
    for fi37 := 0; fi37 < 16; fi37++ {
        v38 := GenS4{F0: -633, F1: "", F2: false}
        _ = fi37
        _ = v38
    }
    v39 := GenS4{F0: 0, F1: "hello", F2: true}
    switch (("foo bar" + (((("foo bar" + (("a" + ("\n\t" + "foo bar")) + ("hello" + "hello"))) + ("\n\t" + "a")) + ("\n\t" + "")) + ("foo bar" + "\n\t"))) + ("hello" + "foo bar")) {
    case "\n\t":
        v2 += 470
    case "":
        v40 := GenS5{F0: (v36 << ((69) & 63))}
        _ = v40
    default:
        v2 = 32767
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v39
}
