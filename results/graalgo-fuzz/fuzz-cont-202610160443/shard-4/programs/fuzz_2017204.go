package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 bool; F1 int64; F2 int16 }
type GenS4 struct { F0 string }
func genFn0() int32 {
    if false {
        var v7 int8 = 1
        _ = v7
    } else {
        var v8 bool = (0.1 > 1e10)
        v9 := GenS1{F0: (("" + "\n\t") + (("" + "foo bar") + "hello")), F1: ("" + "hello")}
        if (false && true) {
            switch "foo bar" {
            case "\n\t":
                if (v8 && v8) {
                    fmt.Println(v8)
                    if false {
                        v10 := GenS2{F0: 65535, F1: 42}
                        v11 := GenS3{F0: (true && (v8 && (!v8))), F1: 42, F2: 1000}
                        v12 := []int{128}
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    v8 = ((!true) && v8)
                }
            }
        }
        _ = v8
        _ = v9
    }
    v13 := GenS4{F0: "foo bar"}
    v14 := []int{4294967295, 0, 32768, 32768}
    var v15 float32 = ((3.14159 * 1.5) + 0.1)
    _ = v13
    _ = v14
    _ = v15
    return 2147483647
}
func genFn1(p0 uint) uint64 {
    var v16 int8 = 127
    v17 := []int{-546, -2147483648, 2147483647}
    fmt.Println(p0, v16)
    _ = p0
    _ = v16
    _ = v17
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -1
    v3 := make(map[string]string)
    var v4 float64 = 1e10
    var v5 string = ((("a" + "foo bar") + ("a" + "hello")) + ("foo bar" + ("hello" + "foo bar")))
    v6 := make(map[int]string)
    _ = len(v6)
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 uint32 = 939
        for fi20 := 0; fi20 < 3; fi20++ {
            v19 = v19
            if (false && false) {
                if false {
                    var v21 uint = 677
                    var v22 uint = 330
                    v2 -= v2
                    _ = v21
                    _ = v22
                }
            } else {
                var v23 int8 = 127
                var v24 int = fi18
                if false {
                    v25 := GenS4{F0: (v5 + "hello")}
                    v2 += (v2 % ((-1) | 1))
                    if (!false) {
                        switch v5 {
                        case "hello":
                            var v26 int64 = -916
                            _ = v26
                        case "":
                            if (v23 < (v23 - v23)) {
                                if true {
                                    fmt.Println(v5, v4, v4)
                                } else {
                                    var v27 string = "hello"
                                    _ = v27
                                }
                            } else {
                                v28 := []int{(fi20 | v24)}
                                v29 := GenS1{F0: "hello", F1: "foo bar"}
                                _ = v28
                                _ = v29
                            }
                        }
                    } else {
                        v30 := []int{}
                        v31 := GenS1{F0: ("" + v5), F1: "\n\t"}
                        v23 = v23
                        _ = v30
                        _ = v31
                    }
                    _ = v25
                }
                switch ("a" + v5) {
                case "\n\t":
                    v19 = (^(v19))
                case "hello":
                    v24 -= 127
                default:
                    fmt.Println(v5, v2)
                }
                _ = v23
                _ = v24
            }
            v4 = v4
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v4, v2)
    v4 += 3.14159
    _ = len(v1)
    switch "\n\t" {
    case "hello":
        v4 += 1.5
    }
    v2 += -1
    var v32 uint64 = 160
    v5 = "a"
    var v33 float64 = 1.5
    var v34 uint32 = 2147483647
    v5 = (v5 + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
