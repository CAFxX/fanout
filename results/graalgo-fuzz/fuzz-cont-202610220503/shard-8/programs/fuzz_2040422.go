package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 uint64 = (v1 % ((v1) | 1))
    var v5 int32 = -994
    v6 := make(map[string]uint32)
    var v7 int8 = -118
    if true {
        v5 -= -2147483648
        v1 -= 42
    } else {
        v8 := GenS1{F0: (v7 * (v7 << ((65535) & 63))), F1: ""}
        _ = v8
    }
    v9 := GenS2{F0: true, F1: 0}
    v10 := []int{127, -2147483648, 4294967295}
    fmt.Println(v1, v4, v4)
    v11 := GenS1{F0: v7, F1: ("a" + "\n\t")}
    var v12 int64 = 0
    v10 = append(v10, int(v5))
    var v13 int64 = 42
    if false {
        switch ("a" + "foo bar") {
        case "a":
            var v14 uint32 = 27
            _ = v14
        default:
            if (!true) {
                v1 -= (v4 & v1)
            } else {
                if false {
                    v13 = (v12 + (v13 & v13))
                } else {
                    v15 := GenS3{F0: (("foo bar" + "foo bar") + ("a" + "")), F1: ("" + "hello")}
                    if false {
                        switch "a" {
                        case "":
                            v4 += v1
                        case "\n\t":
                            if false {
                                if ("hello" == (("hello" + ("a" + "foo bar")) + ("foo bar" + ""))) {
                                    switch ("hello" + "\n\t") {
                                    case "hello":
                                        var v16 int16 = -135
                                        _ = v16
                                    case "foo bar":
                                        for fi17 := 0; fi17 < 11; fi17++ {
                                            var v18 int = (fi17 * (fi17 | -2147483648))
                                            v6["a"] = 587
                                            _ = fi17
                                            _ = v18
                                        }
                                    case "a":
                                        var v19 float64 = (1.5 + (0.1 + -2.25))
                                        _ = v19
                                    }
                                    var v20 uint = 255
                                    _ = v20
                                }
                            } else {
                                fmt.Println(v5, v12)
                            }
                        }
                    }
                    _ = v15
                }
                for fi21 := 0; fi21 < 5; fi21++ {
                    fmt.Println(v5, fi21)
                    v22 := GenS2{F0: false, F1: -737}
                    v23 := GenS4{F0: 42}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                var v24 int8 = v7
                _ = v24
            }
        }
        if false {
            v13 = v12
            var v25 float32 = (1e10 / (float32((3.14159 * 1e10)) * 1.5))
            var v26 uint = 2
            _ = v25
            _ = v26
        }
    }
    v5 -= (v5 * v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
