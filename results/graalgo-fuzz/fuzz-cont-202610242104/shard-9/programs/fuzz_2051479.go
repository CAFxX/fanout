package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int8; F2 int16 }
type GenS6 struct { F0 int16; F1 int8 }
func genFn0(p0 int8) int32 {
    v4 := []int{-24, -1, 2147483647}
    var v5 float32 = (1e10 * (-2.25 / 1.5))
    var v6 string = "hello"
    var v7 int16 = -822
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 int8) uint {
    var v8 int = 32768
    v9 := []int{v8, v8}
    if true {
        fmt.Println(p0, v8)
        v10 := GenS1{F0: 2, F1: 541}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v10
    } else {
        v11 := GenS2{F0: 4294967295}
        _ = v11
    }
    if (("foo bar" == (("hello" + (("" + (("" + "\n\t") + "a")) + ("foo bar" + "hello"))) + "\n\t")) || false) {
        p0 -= p0
    } else {
        if len(v9) > 0 {
            v9[0] = v8
        }
    }
    _ = p0
    _ = v8
    _ = v9
    return 65535
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]string)
    var v3 uint32 = (v1 - v1)
    v12 := GenS3{F0: 2147483647, F1: -1}
    v13 := []int{42}
    switch (("a" + ("foo bar" + "")) + "a") {
    case "":
        switch "hello" {
        case "hello":
            v3 = (v3 / ((247) | 1))
        }
    case "foo bar":
        if ((23 < 0) || false) {
            if false {
                fmt.Println(v1, v1)
                var v14 uint64 = 42
                var v15 uint64 = 53
                _ = v14
                _ = v15
            } else {
                if true {
                    var v16 int16 = 42
                    if (!((3.14159 + float64(1e10)) == (0.0 * 3.14159))) {
                        v3 = (v1 - v1)
                        switch "" {
                        case "hello":
                            v16 += (v16 - v16)
                        default:
                            v17 := v2["foo bar"]
                            _ = v17
                        }
                        if (!(42 == 9223372036854775807)) {
                            v18 := GenS4{F0: 18446744073709551615}
                            _ = v18
                        } else {
                            var v19 int8 = -1
                            switch 0 {
                            case 32768:
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    var v21 uint16 = 2
                                    var v22 int64 = 0
                                    fmt.Println(v19, v22, fi20)
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            case 0:
                                v19 = (^(v19))
                            }
                            var v23 uint16 = 467
                            _ = v19
                            _ = v23
                        }
                    }
                    switch 0 {
                    case 127:
                        fmt.Println(v3, v3)
                    case 237:
                        var v24 uint = 42
                        _ = v24
                    }
                    _ = v16
                } else {
                    var v25 int16 = -1
                    for fi26 := 0; fi26 < 8; fi26++ {
                        switch "foo bar" {
                        case "hello":
                            if false {
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    var v28 float32 = 0.1
                                    _ = fi27
                                    _ = v28
                                }
                                v3 = v1
                            } else {
                                var v29 float32 = 1.5
                                var v30 int = 42
                                v31 := GenS5{F0: 985, F1: 1, F2: (v25 + 42)}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        case "a":
                            var v32 uint64 = 42
                            _ = v32
                        default:
                            _ = len(v2)
                        }
                        _ = fi26
                    }
                    _ = v25
                }
                var v33 uint16 = 431
                fmt.Println(v3, v1, v3)
                _ = v33
            }
        } else {
            var v34 string = "foo bar"
            v35 := GenS1{F0: 53, F1: 0}
            _ = v34
            _ = v35
        }
    }
    v3 += 1
    v36 := GenS6{F0: 42, F1: -1}
    var v37 uint = 224
    var v38 string = "hello"
    v39 := GenS6{F0: -1, F1: 1}
    _ = len(v13)
    var v40 uint64 = 2
    fmt.Println(v1, v37, v37)
    var v41 uint8 = 127
    v40 = v40
    var v42 int64 = -653
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
