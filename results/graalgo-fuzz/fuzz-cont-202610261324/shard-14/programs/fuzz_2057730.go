package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint }
type GenS3 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint; F1 int16; F2 uint32 }
type GenS5 struct { F0 int8; F1 int64; F2 uint16 }
type GenS6 struct { F0 string; F1 int8 }
func genFn0(p0 int32, p1 int16, p2 uint16) uint64 {
    v7 := []int{1}
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 227, F1: 128}
        var v10 uint16 = p2
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1() uint16 {
    var v11 string = ("foo bar" + (("\n\t" + ("foo bar" + ("" + "\n\t"))) + "hello"))
    if (4294967295 < 4294967295) {
        v11 = "\n\t"
        var v12 float32 = 0.0
        v13 := []int{}
        _ = v12
        _ = v13
    }
    v11 = ((v11 + v11) + ((v11 + (v11 + v11)) + ((((v11 + ("\n\t" + ("foo bar" + "hello"))) + v11) + v11) + v11)))
    v11 = v11
    _ = v11
    return 0
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]bool)
    var v3 int32 = -1
    v4 := []int{32768, 42, 1, -9223372036854775808}
    v5 := make(map[int]bool)
    var v6 int = 42
    v4 = append(v4, v6)
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS2{F0: 42, F1: (v3 + (v3 & (v3 + (v3 ^ -1)))), F2: 255}
        _ = fi14
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := GenS3{F0: v3, F1: 158, F2: 281}
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 int32 = -2147483648
                _ = len(v4)
                if false {
                    if true {
                        switch ("foo bar" + ("" + "a")) {
                        case "foo bar":
                            var v20 string = ""
                            _ = v20
                        case "\n\t":
                            _ = len(v4)
                        case "":
                            switch "a" {
                            case "":
                                var v21 int64 = 9223372036854775807
                                _ = v21
                            default:
                                var v22 uint = 65535
                                _ = v22
                            }
                        }
                    }
                } else {
                    var v23 string = "foo bar"
                    if (!false) {
                        v24 := GenS4{F0: 255, F1: -32768, F2: 2}
                        _ = v24
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v23
                }
                _ = fi18
                _ = v19
            }
            _ = fi16
            _ = v17
        }
    }
    fmt.Println(v3, v1, v1)
    v25 := GenS4{F0: 2, F1: -1, F2: 623}
    var v26 uint = 2
    if (!(-605 >= 326)) {
        var v27 uint = 0
        _ = v27
    } else {
        var v28 bool = true
        _ = v28
    }
    v29 := GenS1{F0: 2, F1: 42}
    var v30 int64 = 9223372036854775807
    for fi31 := 0; fi31 < 18; fi31++ {
        if false {
            fmt.Println(v3, v3)
        }
        var v32 uint8 = 127
        v33 := GenS5{F0: -1, F1: (^(v30)), F2: 0}
        _ = fi31
        _ = v32
        _ = v33
    }
    v34 := GenS6{F0: "\n\t", F1: 1}
    v6 += v6
    var v35 float32 = 1.5
    _ = len(v4)
    fmt.Println(v30, v26)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
