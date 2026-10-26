package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 bool; F1 int32; F2 int16 }
type GenS4 struct { F0 uint8; F1 string; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int16 }
func genFn0() bool {
    var v5 float32 = 0.0
    v6 := []int{474, 128}
    var v7 int = -1
    var v8 uint32 = 2147483647
    switch "foo bar" {
    case "a":
        v9 := GenS1{F0: (^(v8)), F1: 128}
        _ = v9
    case "foo bar":
        v8 = v8
    default:
        if true {
            if false {
                v10 := GenS1{F0: v8, F1: 2}
                if false {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    switch "hello" {
                    case "foo bar":
                        var v11 int64 = -244
                        _ = v11
                    case "\n\t":
                        var v12 float64 = 1e10
                        _ = v12
                    case "hello":
                        v7 -= v7
                    default:
                        var v13 float32 = 1.5
                        _ = v13
                    }
                    switch -9223372036854775808 {
                    case 2147483647:
                        switch "" {
                        case "hello":
                            v5 += v5
                        case "a":
                            v7 += 32767
                        case "":
                            for fi14 := 0; fi14 < 19; fi14++ {
                                if (!(2 != 525)) {
                                    fmt.Println(v5)
                                    switch ("" + "a") {
                                    case "foo bar":
                                        v7 -= fi14
                                    case "":
                                        v15 := GenS1{F0: 4294967295, F1: 0}
                                        _ = v15
                                    case "\n\t":
                                        for fi16 := 0; fi16 < 5; fi16++ {
                                            v6 = append(v6, (v7 << ((760) & 63)))
                                            _ = fi16
                                        }
                                    default:
                                        v7 += (-(v7))
                                    }
                                } else {
                                    var v17 uint32 = 4294967295
                                    v18 := GenS2{F0: true, F1: 756}
                                    for fi19 := 0; fi19 < 17; fi19++ {
                                        fmt.Println(v17)
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v7 = fi19
                                        _ = fi19
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                var v20 int32 = -901
                                _ = fi14
                                _ = v20
                            }
                        }
                    }
                }
                _ = v10
            } else {
                var v21 bool = true
                if v21 {
                    v22 := GenS3{F0: v21, F1: 731, F2: 42}
                    var v23 int32 = -1
                    var v24 float32 = 1.5
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    if (-1 >= -2147483648) {
                        var v25 string = ("hello" + "a")
                        _ = len(v6)
                        v26 := GenS1{F0: (v8 + (v8 % ((42) | 1))), F1: 129}
                        _ = v25
                        _ = v26
                    }
                }
                v27 := GenS4{F0: 50, F1: "", F2: v8}
                _ = v21
                _ = v27
            }
        }
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (984 <= 255)
}

func main() {
    var v1 uint16 = 639
    var v2 uint = 982
    v3 := make(map[uint32]string)
    v4 := []int{32768, 4294967295}
    v1 += 12
    var v28 int16 = 42
    v3[1] = "\n\t"
    v28 = (v28 & v28)
    var v29 uint64 = 799
    if len(v4) > 0 {
        v4[0] = 32768
    }
    var v30 uint8 = uint8(v29)
    delete(v3, 2147483648)
    v4 = append(v4, -9223372036854775808)
    v31 := GenS5{F0: 646, F1: uint8(v29), F2: v28}
    switch "a" {
    case "hello":
        if ((^(v30)) > 1) {
            if (2 >= 539) {
                switch ("a" + "hello") {
                case "hello":
                    if (true && (!(42 <= -9223372036854775808))) {
                        v30 = 1
                        v3[65535] = ("a" + "hello")
                    } else {
                        fmt.Println(v28, v1)
                        fmt.Println(v1, v29, v29)
                    }
                default:
                    if ("a" != (("hello" + ("foo bar" + "\n\t")) + ("hello" + "hello"))) {
                        var v32 int16 = v28
                        var v33 int = 4294967295
                        _ = v32
                        _ = v33
                    }
                }
                for fi34 := 0; fi34 < 19; fi34++ {
                    var v35 int32 = 2147483647
                    v29 = (v29 / ((v29) | 1))
                    _ = fi34
                    _ = v35
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v36 := GenS1{F0: 65535, F1: 255}
            v28 = v28
            _ = v36
        }
    }
    if (264 < v29) {
        var v37 string = ""
        _ = v37
    } else {
        fmt.Println(v29, v28, v28)
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 uint64 = v29
        _ = fi38
        _ = v39
    }
    var v40 uint16 = v1
    var v41 float64 = 0.0
    _ = len(v3)
    if true {
        fmt.Println(v28, v40, v40)
    }
    var v42 uint8 = (^(v30))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
