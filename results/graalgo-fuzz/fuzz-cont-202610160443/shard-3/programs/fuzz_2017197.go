package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint64; F1 int8; F2 bool }
type GenS3 struct { F0 uint }
func genFn0() int32 {
    if ((!true) || (!(-464 <= 2147483647))) {
        for fi4 := 0; fi4 < 1; fi4++ {
            for fi5 := 0; fi5 < 1; fi5++ {
                if (("\n\t" < "") || (3.14159 >= float64(1.5))) {
                    var v6 float64 = 1.5
                    _ = v6
                } else {
                    v7 := []int{(fi5 ^ fi4)}
                    if true {
                        _ = len(v7)
                        switch (fi5 % (((-(fi4))) | 1)) {
                        case -1:
                            if (-128 >= 42) {
                                v8 := GenS1{F0: 0, F1: (fi5 ^ 2147483647)}
                                var v9 uint16 = 42
                                _ = v8
                                _ = v9
                            } else {
                                var v10 int8 = 60
                                v11 := GenS1{F0: 32767, F1: 509}
                                _ = v10
                                _ = v11
                            }
                        case 128:
                            var v12 bool = false
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = v12
                        }
                    }
                    switch "hello" {
                    case "hello":
                        fmt.Println(fi4)
                    case "\n\t":
                        var v13 int = (fi4 + fi4)
                        _ = v13
                    default:
                        if false {
                            if len(v7) > 0 {
                                v7[0] = fi5
                            }
                            fmt.Println(fi4, fi4)
                            v14 := GenS1{F0: -633, F1: (fi5 - fi5)}
                            _ = v14
                        } else {
                            if false {
                                var v15 uint = 18446744073709551615
                                _ = v15
                            }
                        }
                    }
                    _ = v7
                }
                switch "hello" {
                case "a":
                    fmt.Println(fi4)
                case "\n\t":
                    v16 := []int{(fi4 + fi4)}
                    _ = v16
                case "":
                    var v17 int8 = -125
                    _ = v17
                }
                var v18 float32 = -2.25
                _ = fi5
                _ = v18
            }
            var v19 uint8 = 128
            v19 += v19
            _ = fi4
            _ = v19
        }
        if (42 >= -32768) {
            switch ("" + "hello") {
            case "hello":
                var v20 int64 = 564
                _ = v20
            case "a":
                switch ("" + (((((("a" + "foo bar") + "hello") + ("\n\t" + "\n\t")) + "") + (("\n\t" + "a") + (("\n\t" + "hello") + ("hello" + "a")))) + "")) {
                case "a":
                    var v21 uint32 = 42
                    _ = v21
                default:
                    var v22 int16 = 625
                    _ = v22
                }
            default:
                var v23 int64 = 0
                _ = v23
            }
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 uint16 = 255
                v26 := []int{9223372036854775807, (fi24 / ((fi24) | 1)), fi24, 1}
                _ = fi24
                _ = v25
                _ = v26
            }
            v27 := []int{}
            _ = v27
        } else {
            v28 := []int{}
            var v29 int8 = 100
            _ = v28
            _ = v29
        }
    } else {
        var v30 int8 = 127
        _ = v30
    }
    var v31 bool = true
    var v32 int16 = -32768
    _ = v31
    _ = v32
    return 42
}
func genFn1() bool {
    var v33 int64 = 42
    var v34 int8 = 77
    _ = v33
    _ = v34
    return (v33 >= 9223372036854775807)
}

func main() {
    var v1 string = (("a" + "") + ("" + ("foo bar" + "hello")))
    var v2 int8 = 42
    var v3 bool = true
    v3 = true
    v35 := []int{-923, 4294967295, -1}
    v36 := GenS1{F0: 1, F1: 0}
    v3 = v3
    var v37 uint16 = 255
    switch 42 {
    case -9223372036854775808:
        v38 := GenS2{F0: 42, F1: v2, F2: (v3 || (!v3))}
        _ = v38
    }
    var v39 uint = 179
    v3 = v3
    v40 := GenS3{F0: uint(v37)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v40
}
