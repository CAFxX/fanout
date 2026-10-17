package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 bool }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int16; F1 int64; F2 uint64 }
type GenS5 struct { F0 int8; F1 int16; F2 int64 }
type GenS6 struct { F0 int }

func main() {
    var v1 int64 = 0
    var v2 int64 = v1
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 16; fi4++ {
        v2 -= 9223372036854775807
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 uint64 = 622
            switch ("" + "") {
            case "\n\t":
                delete(v3, "\n\t")
            case "a":
                v7 := []int{}
                v2 -= 9223372036854775807
                switch -2147483648 {
                case 32768:
                    var v8 int = (-(fi4))
                    _ = v8
                case 127:
                    if false {
                        var v9 int8 = -1
                        v10 := GenS1{F0: ("" + "a"), F1: ("" + "\n\t"), F2: false}
                        switch fi5 {
                        case 32768:
                            v1 -= int64(v9)
                        case 661:
                            var v11 int64 = 9223372036854775807
                            switch (("\n\t" + "") + ("hello" + ("" + "foo bar"))) {
                            case "foo bar":
                                fmt.Println(v6, v1)
                            }
                            var v12 uint8 = 128
                            _ = v11
                            _ = v12
                        default:
                            fmt.Println(v1, fi4)
                        }
                        _ = v10
                        _ = v9
                    } else {
                        v13 := GenS1{F0: "\n\t", F1: "foo bar", F2: (!("a" == "foo bar"))}
                        v14 := GenS1{F0: "hello", F1: ("hello" + "hello"), F2: true}
                        _ = v13
                        _ = v14
                    }
                default:
                    switch -2147483648 {
                    case -1:
                        v15 := GenS2{F0: 928, F1: 0}
                        _ = v15
                    }
                }
                _ = v7
            default:
                var v16 float32 = 1.5
                _ = v16
            }
            var v17 uint32 = 65535
            _ = fi5
            _ = v17
            _ = v6
        }
        _ = fi4
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v1 -= v1
        _ = fi18
    }
    var v19 float32 = 3.14159
    var v20 int32 = -2147483648
    v1 -= v1
    v21 := GenS3{F0: 2, F1: 238, F2: 65535}
    fmt.Println(v20, v1)
    _ = len(v3)
    fmt.Println(v20)
    v20 += v20
    if true {
        v19 += float32(-2.25)
    } else {
        v3["hello"] = 127
        switch "hello" {
        case "hello":
            fmt.Println(v20)
        case "a":
            v22 := GenS3{F0: 0, F1: 900, F2: 0}
            _ = v22
        case "":
            v23 := GenS4{F0: 1, F1: -9223372036854775808, F2: 1}
            _ = v23
        }
    }
    v24 := []int{-659, 4294967295, 2147483647}
    v25 := GenS5{F0: 42, F1: 937, F2: (v1 * v1)}
    var v26 int32 = -2147483648
    var v27 uint8 = 255
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 int8 = -1
        v20 = v20
        _ = fi28
        _ = v29
    }
    if (!false) {
        v30 := GenS5{F0: -1, F1: 417, F2: v1}
        for fi31 := 0; fi31 < 3; fi31++ {
            if (32768 > (fi31 + (-(fi31)))) {
                v32 := GenS2{F0: 2, F1: 231}
                _ = v32
            } else {
                if ((float64(3.14159) + ((1e10 - 0.1) * -2.25)) <= (0.0 + (float64((1.5 + -2.25)) - 1e10))) {
                    var v33 int8 = -53
                    v34 := GenS6{F0: (^(fi31))}
                    _ = v33
                    _ = v34
                } else {
                    v35 := GenS1{F0: "", F1: "hello", F2: true}
                    var v36 int64 = v2
                    v37 := GenS6{F0: 32768}
                    _ = v35
                    _ = v36
                    _ = v37
                }
                v26 -= 2147483647
            }
            v27 = v27
            _ = fi31
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
