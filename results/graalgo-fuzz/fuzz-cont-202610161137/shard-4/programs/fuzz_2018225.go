package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint8; F1 string }
func genFn0() string {
    v6 := []int{}
    var v7 int8 = 42
    if ((!false) || (0 != 10)) {
        var v8 bool = false
        fmt.Println(v7, v8)
        _ = v8
    } else {
        fmt.Println(v7)
        v9 := GenS1{F0: "", F1: v7, F2: 255}
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: "\n\t", F1: (v7 * v7), F2: 977}
        _ = fi10
        _ = v11
    }
    switch 0 {
    case 0:
        v7 = (v7 >> ((193) & 63))
    case 91:
        v7 -= (v7 >> ((65535) & 63))
    case 27:
        v12 := GenS1{F0: "a", F1: (v7 % ((v7) | 1)), F2: 1}
        _ = v12
    default:
        v13 := GenS1{F0: "", F1: v7, F2: 1}
        _ = v13
    }
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    var v1 uint8 = 255
    var v2 float32 = -2.25
    var v3 int32 = 2147483647
    v4 := []uint32{}
    v5 := []string{"", (((("foo bar" + "foo bar") + "hello") + "a") + ("a" + "\n\t")), "a"}
    if (!false) {
        var v14 string = "a"
        fmt.Println(v1, v1)
        v14 = v14
        _ = v14
    }
    if (-46 != 28) {
        if ((("\n\t" + "hello") != ("\n\t" + "\n\t")) && true) {
            v15 := GenS2{F0: (v3 | (v3 | (v3 / ((v3) | 1)))), F1: 177}
            fmt.Println(v2, v3)
            switch 9223372036854775807 {
            case 127:
                for fi16 := 0; fi16 < 16; fi16++ {
                    if false {
                        for fi17 := 0; fi17 < 18; fi17++ {
                            fmt.Println(v1, v3)
                            var v18 int64 = -9223372036854775808
                            _ = fi17
                            _ = v18
                        }
                        v1 = (v1 * v1)
                        var v19 float32 = (0.0 + v2)
                        _ = v19
                    }
                    fmt.Println(v3, v3, v3)
                    for fi20 := 0; fi20 < 8; fi20++ {
                        switch "hello" {
                        case "\n\t":
                            switch 2147483647 {
                            case 0:
                                v21 := GenS3{F0: v1, F1: 100, F2: 42}
                                _ = v21
                            default:
                                fmt.Println(v3)
                            }
                        case "a":
                            v22 := GenS2{F0: v3, F1: 42}
                            _ = v22
                        case "":
                            switch "hello" {
                            case "a":
                                fmt.Println(v2)
                            case "":
                                var v23 uint8 = v1
                                _ = v23
                            }
                            v2 = -2.25
                            var v24 int8 = 2
                            _ = v24
                        }
                        _ = fi20
                    }
                    _ = fi16
                }
            case -1:
                v3 = v3
            case 32768:
                fmt.Println(v1, v2, v2)
            }
            _ = v15
        }
    } else {
        if (v3 < v3) {
            var v25 int32 = v3
            var v26 float64 = (-2.25 + 1.5)
            for fi27 := 0; fi27 < 4; fi27++ {
                if true {
                    var v28 bool = true
                    _ = v28
                } else {
                    v29 := GenS4{F0: 2, F1: (fi27 * (fi27 | fi27))}
                    v30 := GenS5{F0: 4294967295, F1: false}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
            _ = v25
            _ = v26
        }
        var v31 string = "hello"
        v32 := GenS4{F0: 4294967295, F1: 32767}
        _ = v31
        _ = v32
    }
    fmt.Println(v2)
    var v33 int = -227
    var v34 string = (("hello" + "\n\t") + "a")
    fmt.Println(v33, v34)
    v35 := GenS1{F0: "a", F1: 42, F2: 2}
    var v36 int16 = 1000
    if ((v1 - v1) <= (v1 ^ 128)) {
        if len(v4) > 0 {
            v4[0] = 65535
        }
        var v37 float64 = 3.14159
        _ = v37
    } else {
        var v38 uint8 = (v1 >> ((385) & 63))
        _ = v38
    }
    v33 = v33
    v3 -= (v3 * 42)
    v39 := GenS6{F0: (^(v1)), F1: (("foo bar" + v34) + v34)}
    v40 := GenS2{F0: (v3 / (((v3 >> ((2) & 63))) | 1)), F1: 2}
    for fi41 := 0; fi41 < 5; fi41++ {
        _ = len(v4)
        _ = fi41
    }
    v2 = (v2 + 3.14159)
    fmt.Println(v3)
    _ = len(v4)
    v1 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
