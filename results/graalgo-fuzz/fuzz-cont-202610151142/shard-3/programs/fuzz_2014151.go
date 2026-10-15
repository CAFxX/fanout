package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 int16 }
type GenS4 struct { F0 uint; F1 string; F2 int8 }

func main() {
    var v1 uint64 = 2
    var v2 bool = (float64(0.0) == (-2.25 + float64((-2.25 - 0.0))))
    var v3 string = ""
    switch (v3 + "hello") {
    case "":
        var v4 uint64 = v1
        _ = v4
    default:
        v2 = false
    }
    v5 := GenS1{F0: ("\n\t" + v3), F1: 116, F2: "hello"}
    var v6 int64 = 1
    switch (v3 + v3) {
    case "foo bar":
        var v7 string = v3
        _ = v7
    case "\n\t":
        v3 = "\n\t"
    default:
        var v8 int8 = -128
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int = (fi9 / (((fi9 | (^(fi9)))) | 1))
        v11 := []int{(^(v10)), (^(fi9))}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{-2147483648, 42, 0}
    v2 = true
    v3 = ""
    var v13 uint16 = 42
    if false {
        v14 := GenS2{F0: v13, F1: 1}
        _ = len(v12)
        var v15 uint = 0
        _ = v14
        _ = v15
    } else {
        v16 := GenS2{F0: (^(v13)), F1: 255}
        v17 := GenS3{F0: 4294967295, F1: v3, F2: 1000}
        var v18 uint = 65535
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 int = 9223372036854775807
    switch (v3 + "hello") {
    case "\n\t":
        fmt.Println(v2, v13)
    default:
        if v2 {
            for fi20 := 0; fi20 < 14; fi20++ {
                for fi21 := 0; fi21 < 1; fi21++ {
                    if (!v2) {
                        v3 = "hello"
                        if true {
                            v3 = (v3 + "\n\t")
                            if ((1 != 0) || (3.14159 <= float64(1e10))) {
                                _ = len(v12)
                                if v2 {
                                    var v22 uint32 = 2
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(fi21, v13)
                                    case "a":
                                        v2 = v2
                                    case "":
                                        fmt.Println(v22)
                                    default:
                                        var v23 uint = 301
                                        _ = v23
                                    }
                                    var v24 int64 = v6
                                    _ = v22
                                    _ = v24
                                } else {
                                    v2 = v2
                                    v1 -= (v1 + (v1 - 4294967295))
                                    v25 := GenS4{F0: 18446744073709551615, F1: ("hello" + v3), F2: -44}
                                    _ = v25
                                }
                            } else {
                                v26 := GenS3{F0: 65535, F1: ("a" + (((v3 + ("hello" + ((v3 + "foo bar") + ("\n\t" + v3)))) + "hello") + ((v3 + v3) + "hello"))), F2: 0}
                                _ = v26
                            }
                        }
                        switch (-(fi20)) {
                        case -534:
                            v2 = (v1 != 0)
                        case 143:
                            var v27 uint64 = 18446744073709551615
                            _ = v27
                        default:
                            v28 := GenS4{F0: 0, F1: v3, F2: 100}
                            _ = v28
                        }
                        switch v3 {
                        case "foo bar":
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        case "":
                            v1 += v1
                        }
                        v6 = (v6 << ((65535) & 63))
                        fmt.Println(v1, v2, fi20)
                        fmt.Println(v19)
                    } else {
                        v13 += 167
                        v29 := GenS1{F0: ("\n\t" + v3), F1: 4294967295, F2: v3}
                        var v30 int = (v19 & v19)
                        _ = v29
                        _ = v30
                    }
                    fmt.Println(v19, v13)
                    _ = fi21
                }
                fmt.Println(v6)
                v31 := GenS4{F0: 1, F1: (("foo bar" + "foo bar") + (("\n\t" + (v3 + v3)) + v3)), F2: 1}
                _ = fi20
                _ = v31
            }
            var v32 string = "hello"
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
