package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int16; F2 string }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }

func main() {
    var v1 uint8 = 0
    var v2 int8 = 42
    var v3 float64 = 0.1
    v4 := []int{2147483647}
    var v5 int64 = -400
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: (v2 & (v2 & (v2 / ((v2) | 1)))), F1: ("" + "hello"), F2: (fi6 + fi6)}
        _ = fi6
        _ = v7
    }
    if true {
        if false {
            switch "hello" {
            case "hello":
                var v8 int64 = (v5 << ((0) & 63))
                _ = v8
            default:
                v9 := GenS2{F0: 0, F1: 18446744073709551615}
                _ = v9
            }
            var v10 int32 = -2147483648
            _ = v10
        } else {
            v2 = (v2 % (((^(v2))) | 1))
            var v11 int32 = 42
            _ = v11
        }
        v12 := GenS3{F0: true}
        _ = v12
    } else {
        var v13 int32 = -126
        switch "" {
        case "\n\t":
            v14 := GenS4{F0: (v1 % (((v1 % ((2) | 1))) | 1)), F1: 42, F2: "\n\t"}
            _ = v14
        default:
            v15 := GenS4{F0: 148, F1: -842, F2: "hello"}
            _ = v15
        }
        switch "" {
        case "a":
            var v16 uint64 = 18446744073709551615
            _ = v16
        default:
            var v17 float64 = v3
            _ = v17
        }
        _ = v13
    }
    fmt.Println(v5, v3)
    var v18 float32 = -2.25
    v18 = 3.14159
    v2 += v2
    var v19 string = (("a" + "hello") + "foo bar")
    v5 = 981
    v20 := GenS2{F0: -119, F1: 65535}
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v5 += (v5 & 9223372036854775807)
        if false {
            var v21 uint16 = 84
            _ = v21
        }
    } else {
        v1 = (v1 ^ (v1 * v1))
        if len(v4) > 0 {
            v4[0] = 127
        }
    }
    var v22 uint64 = 828
    if true {
        v23 := GenS5{F0: v22, F1: 47}
        v18 = (v18 * -2.25)
        _ = v23
    } else {
        v24 := GenS1{F0: 127, F1: v19, F2: 42}
        if true {
            var v25 int64 = -740
            v1 = (v1 * v1)
            if true {
                _ = len(v4)
            }
            _ = v25
        }
        _ = v24
    }
    v26 := GenS5{F0: 42, F1: 1}
    if ((v19 + (v19 + v19)) < (("\n\t" + "hello") + (("hello" + "\n\t") + v19))) {
        if len(v4) > 0 {
            v4[0] = int(v5)
        }
    }
    switch int(v22) {
    case -728:
        v18 = (float32((((float64(0.0) / ((0.0 - 1e10) * v3)) - (1e10 - 0.0)) + v3)) + v18)
    case 2147483647:
        v22 -= 42
    default:
        v27 := GenS5{F0: 18446744073709551615, F1: 100}
        _ = v27
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (v22 < v22) {
        if (42 != -9223372036854775808) {
            if ((197 < 1) || false) {
                v1 = v1
            } else {
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 float64 = 0.1
                    _ = fi28
                    _ = v29
                }
                switch "hello" {
                case "":
                    if true {
                        v30 := GenS6{F0: 2147483647}
                        _ = len(v4)
                        _ = v30
                    }
                case "hello":
                    v31 := GenS1{F0: 42, F1: v19, F2: 128}
                    _ = v31
                case "\n\t":
                    fmt.Println(v3, v1, v2)
                default:
                    v5 = (v5 & v5)
                }
            }
            if false {
                fmt.Println(v5)
                fmt.Println(v19, v5, v2)
                switch -687 {
                case 1:
                    fmt.Println(v5, v3)
                case 4294967295:
                    for fi32 := 0; fi32 < 19; fi32++ {
                        var v33 string = v19
                        _ = fi32
                        _ = v33
                    }
                case -324:
                    var v34 uint8 = v1
                    _ = v34
                default:
                    switch "\n\t" {
                    case "hello":
                        v3 = v3
                    case "a":
                        for fi35 := 0; fi35 < 17; fi35++ {
                            fmt.Println(v19, v2, v19)
                            for fi36 := 0; fi36 < 16; fi36++ {
                                v2 -= -1
                                switch (fi35 | 796) {
                                case 358:
                                    var v37 int8 = int8(v5)
                                    _ = v37
                                }
                                _ = fi36
                            }
                            _ = fi35
                        }
                    case "":
                        var v38 uint8 = 127
                        _ = v38
                    }
                }
            }
            v39 := GenS1{F0: -84, F1: "hello", F2: -2147483648}
            _ = v39
        } else {
            v40 := GenS7{F0: (v5 >> ((42) & 63))}
            _ = v40
        }
        v18 = v18
        v3 = float64(v18)
    } else {
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 int64 = int64(v22)
            _ = fi41
            _ = v42
        }
        var v43 int = 9223372036854775807
        var v44 int32 = -730
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
