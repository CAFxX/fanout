package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int8; F1 int8; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := []string{("a" + "foo bar"), "a"}
    var v2 int64 = -377
    v3 := []int{-676}
    var v4 string = ((("" + "\n\t") + "\n\t") + "")
    v5 := GenS1{F0: 954, F1: (v4 + (v4 + ((v4 + ((v4 + v4) + v4)) + (v4 + ""))))}
    fmt.Println(v2, v2)
    v6 := GenS1{F0: 0, F1: ("" + "\n\t")}
    if true {
        v7 := GenS1{F0: -1, F1: (v4 + v4)}
        _ = v7
    } else {
        switch 4294967295 {
        case -200:
            fmt.Println(v2)
        case 32767:
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := GenS2{F0: 42, F1: -1, F2: true}
                _ = fi8
                _ = v9
            }
        }
        v2 = v2
        if ((0 <= 611) && (!(809 == 42))) {
            var v10 uint64 = 18446744073709551615
            if false {
                v11 := GenS2{F0: 1, F1: 10, F2: (32767 < 32767)}
                switch -962 {
                case 434:
                    fmt.Println(v2, v10, v2)
                case 1:
                    if (-26 != 94) {
                        for fi12 := 0; fi12 < 10; fi12++ {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            v13 := GenS3{F0: 1}
                            v14 := GenS2{F0: 0, F1: -90, F2: true}
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    } else {
                        var v15 float32 = (0.1 * 0.0)
                        v16 := GenS4{F0: 2}
                        _ = v15
                        _ = v16
                    }
                case 42:
                    v4 = (("\n\t" + ((("" + v4) + ("\n\t" + (v4 + v4))) + "")) + v4)
                default:
                    fmt.Println(v10, v10, v4)
                }
                if (0 > 0) {
                    if true {
                        if ((v10 | v10) <= (v10 % (((v10 + (^(v10)))) | 1))) {
                            v2 = v2
                            for fi17 := 0; fi17 < 6; fi17++ {
                                if true {
                                    fmt.Println(v10, fi17, v4)
                                    var v18 bool = false
                                    _ = v18
                                }
                                _ = fi17
                            }
                            v4 = "a"
                        }
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v4 = (v4 + (v4 + ((((("\n\t" + (v4 + v4)) + v4) + "foo bar") + v4) + "a")))
                            var v20 uint32 = 2147483648
                            v10 += (v10 + 2)
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        var v21 int8 = -23
                        switch "hello" {
                        case "":
                            v10 = 511
                        case "\n\t":
                            var v22 float32 = float32((3.14159 - 0.1))
                            _ = v22
                        }
                        v23 := GenS3{F0: -1}
                        _ = v21
                        _ = v23
                    }
                    v24 := GenS5{F0: 65535, F1: 0}
                    _ = v24
                } else {
                    var v25 uint8 = 89
                    fmt.Println(v4, v4)
                    _ = v25
                }
                _ = v11
            }
            _ = v10
        } else {
            v4 = ("foo bar" + ((v4 + v4) + v4))
        }
    }
    var v26 uint32 = 732
    if true {
        fmt.Println(v26, v4)
    }
    v4 = ((((v4 + v4) + "foo bar") + "") + v4)
    for fi27 := 0; fi27 < 10; fi27++ {
        fmt.Println(fi27, v2)
        _ = fi27
    }
    v26 = v26
    if len(v1) > 0 {
        v1[0] = ("" + v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
